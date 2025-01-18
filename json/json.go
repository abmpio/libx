package json

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
)

// 将一个结构体序列化成json字符串并忽略错误
func ObjectToJson(v any) string {
	s, _ := json.Marshal(v)
	return string(s)
}

// 使用json的方式将一个对象转换为另一个对象
func ConvertObjectTo(src interface{}, dest interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, dest)
	if err != nil {
		return err
	}
	return nil
}

// 将一个json字符串转化为一个map，如果无法转换为map[string]interface{}则直接返回一个空的map
func JsonStringToMap(s string) map[string]interface{} {
	var jsonValue map[string]interface{}
	err := json.Unmarshal([]byte(s), &jsonValue)
	if err != nil {
		return jsonValue
	}
	return jsonValue
}

// 将一个json字符串反序列化为一个对象，如果转换失败，则返回nil
func JsonStringToObject(s string, v interface{}) interface{} {
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		return nil
	}
	return v
}

// 将一个byte数组反序列化为一个对象，如果转换失败，则返回nil
func RawMessageToObject(data []byte, v interface{}) interface{} {
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil
	}
	return v
}

// 读取文件中的数据并转换为json对象
func ReadJson(file string, v interface{}) error {
	jsonFile, err := os.Open(file)
	if err != nil {
		return err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &v)
	if err != nil {
		return err
	}
	return nil
}

func TryJsonToAnonymousStruct(j string) (interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(j), &data); err != nil {
		return nil, err
	}

	// Create a slice of StructFields
	fields := make([]reflect.StructField, 0, len(data))
	for k, v := range data {
		fields = append(fields, reflect.StructField{
			Name: k,
			Type: reflect.TypeOf(v),
		})
	}

	// Create the struct type
	t := reflect.StructOf(fields)

	// Unmarshal again, this time to the new struct type
	val := reflect.New(t)
	i := val.Interface()
	if err := json.Unmarshal([]byte(j), &i); err != nil {
		return nil, err
	}
	return i, nil
}
