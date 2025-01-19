package util

import "github.com/abmpio/libx/json"

func StringToInterfaceArray(array []string) []interface{} {
	var (
		interfaceArray []interface{}
		elem           interface{}
	)
	for _, elem = range array {
		jStruct, err := json.TryJsonToAnonymousStruct(elem.(string))
		if err == nil {
			elem = jStruct
		}
		interfaceArray = append(interfaceArray, elem)
	}
	return interfaceArray
}

func StringToInterfaceArray2d(arrays [][]string) [][]interface{} {
	var interfaceArrays [][]interface{}
	for _, req := range arrays {
		var (
			interfaceArray []interface{}
			elem           interface{}
		)
		for _, elem = range req {
			jStruct, err := json.TryJsonToAnonymousStruct(elem.(string))
			if err == nil {
				elem = jStruct
			}
			interfaceArray = append(interfaceArray, elem)
		}
		interfaceArrays = append(interfaceArrays, interfaceArray)
	}
	return interfaceArrays
}
