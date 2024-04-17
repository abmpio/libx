package slicex

// paginate list
func Paginate[V any](list []V, currentPage int, pageSize int) []V {
	startIndex := (currentPage - 1) * pageSize
	endIndex := currentPage * pageSize

	// check start index
	if startIndex >= len(list) {
		return []V{}
	}

	// normalize end index
	if endIndex > len(list) {
		endIndex = len(list)
	}
	return list[startIndex:endIndex]
}
