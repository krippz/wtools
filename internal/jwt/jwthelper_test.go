package jwthelper_test

//goland:noinspection ALL

import (
	"testing"

	jwthelper "github.com/krippz/wtools/internal/jwt"
)

func TestConvertToJSONMap_EmptyByteInput(t *testing.T) {
	var empty []byte
	kvp, err := jwthelper.ConvertToJSONMap(&empty)

	if kvp != nil || err == nil {
		t.Fatal()
	}
}

func TestConvertToJSONMap_NotJSONByteInput(t *testing.T) {
	notJSON := []byte("random-string-not-json")
	kvp, err := jwthelper.ConvertToJSONMap(&notJSON)

	if kvp != nil || err == nil {
		t.Fatal()
	}
}

func TestConvertToJSONMap_JSONByteInput_OneKVP(t *testing.T) {
	json := []byte(`{"name":"krilin"}`)
	want := "krilin"
	kvp, err := jwthelper.ConvertToJSONMap(&json)

	if kvp == nil || want != kvp["name"] || err != nil {
		t.Fatal()
	}
}
