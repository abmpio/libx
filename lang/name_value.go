package lang

type NameValue struct {
	Name  string      `json:"name" bson:"name"`
	Value interface{} `json:"value" bson:"value"`
}

type NameValueWith[T any] struct {
	Name  string `json:"name" bson:"name"`
	Value T      `json:"value" bson:"value"`
}

func NewNameWithT[T any](name string, v T) *NameValueWith[T] {
	return &NameValueWith[T]{
		Name:  name,
		Value: v,
	}
}
