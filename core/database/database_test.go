package database

import (
	"testing"
)

func TestGetStorage(t *testing.T) {

	_, err := GetStorage()

	if err != nil {
		t.Error("Get Storage Error: ", err)
	}
}
