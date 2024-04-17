package lang

import "fmt"

type KeyValuePair[K, V any] struct {
	Key   K
	Value V
}

func NewKeyValuePair[K, V any](key K, v V) KeyValuePair[K, V] {
	return KeyValuePair[K, V]{
		Key:   key,
		Value: v,
	}
}

func (p *KeyValuePair[K, V]) String() string {
	return fmt.Sprintf("key:%v,value:%v", p.Key, p.Value)
}
