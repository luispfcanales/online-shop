package storage

import (
	"testing"
)

func TestConnection(t *testing.T) {
	repository := New()
	err := repository.Ping()
	if err != nil {
		t.Errorf("can not connect %v", err)
	}
}
