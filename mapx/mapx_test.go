package mapx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MergeStringMapsT(t *testing.T) {
	v := map[string]string{}
	MergeStringMapsT(map[string]string{
		"username":    "test",
		"displayName": "test_only",
	}, v)
	username, ok := v["username"]
	if !ok || username != "test" {
		t.Fatalf("MergeStringMapsT merge fail")
	}

	MergeStringMapsT(map[string]string{
		"userName": "modifyed",
		"id":       "test",
	}, v, MergeConfig{
		KeyInsensitivise: false,
	})
	t.Logf("v:%+v", v)
	username, ok = v["username"]
	if !ok || username != "modifyed" {
		t.Fatalf("MergeStringMapsT merge fail")
	}

}

func Test_GetKeyValueAs(t *testing.T) {
	type tType struct {
		Id string
	}
	v := map[string]interface{}{
		"username":    "test",
		"displayName": "test_only",
		"age":         "30",
		"pValue": &tType{
			Id: "amy",
		},
	}

	pValue := GetKeyValueAs[*tType](v, "pValue")
	require.NotNil(t, pValue)
	require.Equal(t, "amy", pValue.Id)
	require.Nil(t, GetKeyValueAs[*tType](v, "pValueNil"))

	require.Nil(t, GetKeyValueAs[*int32](v, "age"))
	require.Equal(t, int32(30), GetKeyValueAs[int32](v, "age"))
	require.Equal(t, 30, GetKeyValueAs[int](v, "age"))
	require.Equal(t, float64(30), GetKeyValueAs[float64](v, "age"))
	require.Equal(t, "30", GetKeyValueAs[string](v, "age"))
}
