package util

import (
	"context"
	"errors"
	"time"

	"github.com/cenkalti/backoff/v5"
)

// RetryConfig 配置
type RetryConfig struct {
	InitialInterval     time.Duration // 初始间隔
	Multiplier          float64       // 间隔倍增
	MaxInterval         time.Duration // 最大间隔
	RandomizationFactor float64       // 抖动因子
	MaxRetries          int           // 最大重试次数（0 = 无限）
	NonRetryableErrors  []error       // 遇到这些错误 不重试，直接返回

	// 用于输出每次重试的错误信息
	// 没有设置时则不输出
	// retryIndex:当前重试索引,从 1 开始
	// nextRetryTime:下一次重试间隔
	// err:本次操作返回的错误
	ErrorRetryFn func(retryIndex int, nextRetryTime time.Duration, err error)
}

// 默认配置：无限重试 + 指数增长 + jitter
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		InitialInterval:     1 * time.Second,
		Multiplier:          1.5,
		MaxInterval:         30 * time.Second,
		RandomizationFactor: 0.5,
		MaxRetries:          0, // 0=无限
		NonRetryableErrors:  nil,

		ErrorRetryFn: nil,
	}
}

// 判断 err 是否属于不可重试错误
func isNonRetryable(err error, list []error) bool {
	if err == nil {
		return false
	}
	for _, e := range list {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}

// 重试操作，直到成功或达到最大重试次数或上下文取消
func RetryWithBackoffWithCtx(op func() error, cfg *RetryConfig) error {
	return RetryWithBackoff(context.Background(), op, cfg)
}

// 重试操作，直到成功或达到最大重试次数或上下文取消
func RetryWithBackoff(ctx context.Context, op func() error, cfg *RetryConfig) error {

	bo := backoff.NewExponentialBackOff()
	bo.InitialInterval = cfg.InitialInterval
	bo.Multiplier = cfg.Multiplier
	bo.MaxInterval = cfg.MaxInterval
	bo.RandomizationFactor = cfg.RandomizationFactor
	bo.Reset()

	retryCount := 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err := op()
		if err == nil {
			return nil
		}

		// 不可重试错误 → 直接返回
		if isNonRetryable(err, cfg.NonRetryableErrors) {
			return err
		}

		retryCount++

		// 超过最大重试次数（非 0 才生效）
		if cfg.MaxRetries > 0 && retryCount > cfg.MaxRetries {
			return err
		}

		// 计算下一次重试间隔
		next := bo.NextBackOff()

		if cfg.ErrorRetryFn != nil {
			cfg.ErrorRetryFn(retryCount, next, err)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(next):
		}
	}
}
