package day08

import (
	_ "embed"
	"testing"
)

func TestImage_checksum(t *testing.T) {
	type fields struct {
		width  int
		height int
		data   string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"example 1", fields{3, 2, "123456789012"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := newImage(tt.fields.width, tt.fields.height, tt.fields.data)
			if got := i.checksum(); got != tt.want {
				t.Errorf("Image.checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
