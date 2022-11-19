package jwtHelper

import (
	"testing"
)

func TestConvertToJsonMap_EmptyByteInput(t *testing.T) {
	var empty []byte
	kvp, err := ConvertToJsonMap(&empty)

	if kvp != nil || err == nil {
		t.Fatal()
	}
}

func TestConvertToJsonMap_NotJsonByteInput(t *testing.T) {
	notJson := []byte("random-string-not-json")
	kvp, err := ConvertToJsonMap(&notJson)

	if kvp != nil || err == nil {
		t.Fatal()
	}
}

func TestConvertToJsonMap_JsonByteInput_OneKVP(t *testing.T) {
	json := []byte(`{"name":"krilin"}`)
	want := "krilin"
	kvp, err := ConvertToJsonMap(&json)

	if kvp == nil || want != kvp["name"] || err != nil {
		t.Fatal()
	}
}
