package slicex

// 将一个切片分割成指定数量的切片,每个子切片的长度都是chunkSize
func ChunkSlice[SV any](input []SV, chunkSize int) [][]SV {
	var result [][]SV

	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize

		if end > len(input) {
			end = len(input)
		}

		result = append(result, input[i:end])
	}

	return result
}

// 将一个切片均匀地分配到多个子切片中，子切片的数量由参数 partitionCount 决定
// 如果原始切片的长度小于 partitionCount，则每个子切片只包含一个元素
func PartitionSlice[SV any](input []SV, partitionCount int) [][]SV {
	var partitionSlice [][]SV
	count := 0
	if len(input) < partitionCount {
		count = len(input)
	} else {
		count = partitionCount
	}
	partitionSlice = make([][]SV, count)

	for i := 0; i < count; i++ {
		partitionSlice[i] = make([]SV, 0)
	}

	for i, v := range input {
		index := i % count
		partitionSlice[index] = append(partitionSlice[index], v)
	}

	return partitionSlice
}

// 将一个新项目插入到slice的指定索引处，并返回新的索引
// 如果index小于或等于0,则插入到最前面
// 如果index大于或等于slice的长度,则插入到最后面
func InsertItemAtIndex[SV any](slice []SV, index int, value SV) []SV {
	// 确保索引有效
	if index <= 0 {
		newArray := make([]SV, 0)
		newArray = append(newArray, value)
		newArray = append(newArray, slice...)
		return newArray
	}
	if index >= len(slice) {
		return append(slice, value)
	}
	// 插入元素
	return append(slice[:index], append([]SV{value}, slice[index:]...)...)
}
