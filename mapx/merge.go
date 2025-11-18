package mapx

import (
	"github.com/abmpio/libx/stringslice"
)

type MergeConfig struct {
	OnlyReplaceExist bool
	OnlyAdd          bool
	// Deprecated: KeyInsensitivise is a misspelling and will be removed in a future release.
	// Please use KeyInsensitive instead.
	KeyInsensitivise bool

	// KeyInsensitive controls whether key matching is case-insensitive.
	KeyInsensitive bool
	ExcludeKey     []string
}

func defaultMergeConfig() *MergeConfig {
	return &MergeConfig{
		OnlyReplaceExist: false,
		OnlyAdd:          false,
		KeyInsensitivise: false,
		KeyInsensitive:   false,
		ExcludeKey:       make([]string, 0),
	}
}

func (c *MergeConfig) IsCaseInsensitive() bool {
	// first used KeyInsensitive
	if c.KeyInsensitive {
		return true
	}

	// second old KeyInsensitivise
	if c.KeyInsensitivise {
		return true
	}
	return false
}

// merge map
func MergeMaps(src map[string]interface{}, dst map[string]interface{}, opts ...MergeConfig) {
	currentOpt := defaultMergeConfig()
	if len(opts) > 0 {
		currentOpt = &opts[0]
	}

	caseInsensitive := currentOpt.IsCaseInsensitive()
	for sk, sv := range src {
		if len(currentOpt.ExcludeKey) > 0 {
			ignore := false
			if caseInsensitive {
				ignore = stringslice.ContainsIgnoreLowercase(currentOpt.ExcludeKey, sk)
			} else {
				ignore = stringslice.Contains(currentOpt.ExcludeKey, sk)
			}
			if ignore {
				continue
			}
		}
		tk := KeyExists(sk, dst, caseInsensitive)
		if tk == "" {
			// not exist
			if currentOpt.OnlyReplaceExist {
				continue
			}
			dst[sk] = sv
			continue
		}
		if currentOpt.OnlyAdd {
			continue
		}
		dst[tk] = sv
	}
}

// merge map
func MergeStringMapsT[V any](src map[string]V, dst map[string]V, opts ...MergeConfig) {
	currentOpt := defaultMergeConfig()
	if len(opts) > 0 {
		currentOpt = &opts[0]
	}

	caseInsensitive := currentOpt.IsCaseInsensitive()
	for sk, sv := range src {
		if len(currentOpt.ExcludeKey) > 0 {
			ignore := false
			if caseInsensitive {
				ignore = stringslice.ContainsIgnoreLowercase(currentOpt.ExcludeKey, sk)
			} else {
				ignore = stringslice.Contains(currentOpt.ExcludeKey, sk)
			}
			if ignore {
				continue
			}
		}
		tk := KeyExists(sk, dst, caseInsensitive)
		if tk == "" {
			// not exist
			if currentOpt.OnlyReplaceExist {
				continue
			}
			dst[sk] = sv
			continue
		}
		if currentOpt.OnlyAdd {
			continue
		}
		dst[tk] = sv
	}
}

func MergeMapsTWith[TS comparable, VS any, TD comparable, VD any](srcMap map[TS]VS, toMap map[TD]VD, itemFunc func(srcKey TS, srcValue VS) (TD, VD)) {
	for sk, sv := range srcMap {
		toKey, toValue := itemFunc(sk, sv)
		toMap[toKey] = toValue
	}
}
