package mapx

import "strings"

func KeyExists[V any](k string, m map[string]V, keyInsensitivise bool) string {
	lk := k
	if keyInsensitivise {
		lk = strings.ToLower(k)
	}
	for mk := range m {
		lmk := mk
		if keyInsensitivise {
			lmk = strings.ToLower(mk)
		}
		if lmk == lk {
			return mk
		}
	}
	return ""
}
