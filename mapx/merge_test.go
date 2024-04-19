package mapx

import (
	"testing"
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
