package cryptonight

import "testing"

func TestEcho(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "echo test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Echo()
		})
	}
}
