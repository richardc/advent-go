package day08

import (
	_ "embed"

	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year: 2019,
			Day:  8,
			Input: func() any {
				image := newImage(25, 6, puzzle)
				return &image
			},
			Part1: func(i any) any { return i.(*Image).checksum() },
		},
	)
}

type Layer struct {
	data string
}

func newLayer(s string) Layer {
	return Layer{s}
}

func (l Layer) Count(digit byte) int {
	return len(slices.Filter([]byte(l.data), func(b byte) bool { return b == digit }))
}

type Image struct {
	width  int
	height int
	layers []Layer
}

func newImage(width, height int, data string) Image {
	var layers []Layer
	size := width * height
	for len(data) > size {
		layers = append(layers, newLayer(data[:size]))
		data = data[size:]
	}
	return Image{
		width:  width,
		height: height,
		layers: layers,
	}
}

func (i *Image) checksum() int {
	layer := slices.MinBy(i.layers, func(l Layer) int { return l.Count('0') })
	return layer.Count('1') * layer.Count('2')
}
