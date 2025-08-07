package slicex

import "sync"

type SafeList[T any] struct {
	mu   sync.RWMutex
	data []T
}

func NewSafeList[T any]() *SafeList[T] {
	return &SafeList[T]{data: make([]T, 0)}
}

// add
func (s *SafeList[T]) Add(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, v)
}

// add list
func (s *SafeList[T]) AddList(list []T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, list...)
}

// get by index
func (s *SafeList[T]) Get(index int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if index < 0 || index >= len(s.data) {
		var zero T
		return zero, false
	}
	return s.data[index], true
}

// find by fn
func (s *SafeList[T]) FindIndex(fn func(T) bool) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return FindIndex(s.data, fn)
}

// remove by fn
func (s *SafeList[T]) RemoveBy(fn func(T) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := FindIndex(s.data, fn)
	if index > -1 {
		s.data = append(s.data[:index], s.data[index+1:]...)
	}
}

// remote at index
func (s *SafeList[T]) RemoveAt(index int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.data) {
		return false
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return true
}

func (s *SafeList[T]) RemoveAll() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make([]T, 0)
}

// get length
func (s *SafeList[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// iterate over the list
func (s *SafeList[T]) ForEach(fn func(i int, v T)) {
	otherList := s.Snapshot()
	for i, v := range otherList {
		fn(i, v)
	}
}

// get snapshot of the list
// This returns a copy of the current list, ensuring that the caller gets a consistent view of
// the data without being affected by concurrent modifications.
// It is important to note that the returned slice is a copy, so modifications to it will
// not affect the original list, and vice versa.
func (s *SafeList[T]) Snapshot() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cp := make([]T, len(s.data))
	copy(cp, s.data)
	return cp
}
