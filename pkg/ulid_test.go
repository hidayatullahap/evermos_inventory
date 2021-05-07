package pkg

import (
	"fmt"
	"testing"
)

func TestNewULID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"TestNewULID1", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewULID()
			fmt.Println(got)
		})
	}
}
