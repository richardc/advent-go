package day10

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestField_project(t *testing.T) {
	type fields struct {
		rows int
		cols int
	}
	type args struct {
		origin Point
		point  Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Point
	}{
		{"A", fields{10, 10}, args{Point{}, Point{1, 3}}, []Point{{2, 6}, {3, 9}}},
		{"B", fields{10, 10}, args{Point{}, Point{2, 3}}, []Point{{4, 6}, {6, 9}}},
		{"C", fields{10, 10}, args{Point{}, Point{3, 3}}, []Point{{4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}}},
		{"D", fields{10, 10}, args{Point{}, Point{3, 2}}, []Point{{6, 4}, {9, 6}}},
		{"E", fields{10, 10}, args{Point{}, Point{3, 1}}, []Point{{6, 2}, {9, 3}}},
		{"F", fields{10, 10}, args{Point{}, Point{4, 2}}, []Point{{6, 3}, {8, 4}}},
		{"G", fields{10, 10}, args{Point{}, Point{3, 4}}, []Point{{6, 8}}},
		{"Vertical", fields{10, 10}, args{Point{}, Point{3, 0}}, []Point{{4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}, {9, 0}}},
		{"Horizontal", fields{10, 10}, args{Point{}, Point{0, 3}}, []Point{{0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}}},
		{"Vertical Mid Up", fields{10, 10}, args{Point{5, 0}, Point{3, 0}}, []Point{{2, 0}, {1, 0}, {0, 0}}},
		{"Vertical Mid Down", fields{10, 10}, args{Point{5, 0}, Point{7, 0}}, []Point{{8, 0}, {9, 0}}},
		{"Horizontal", fields{10, 10}, args{Point{}, Point{0, 3}}, []Point{{0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}}},
		{"Reverse C", fields{10, 10}, args{Point{9, 9}, Point{7, 7}}, []Point{{6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}, {0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			got := f.project(tt.args.origin, tt.args.point)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Field.project() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

//go:embed "example0.txt"
var example0 string

//go:embed "example1.txt"
var example1 string

//go:embed "example2.txt"
var example2 string

//go:embed "example3.txt"
var example3 string

//go:embed "example4.txt"
var example4 string

func Test_mostAsteroids(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 0", args{example0}, 8},
		{"example 1", args{example1}, 33},
		{"example 2", args{example2}, 35},
		{"example 3", args{example3}, 41},
		{"example 4", args{example4}, 210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostAsteroids(tt.args.puzzle); got != tt.want {
				t.Errorf("mostAsteroids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_laserBeam(t *testing.T) {
	type args struct {
		puzzle string
		nth    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 4", args{example4, 200}, 802},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := laserBeam(tt.args.puzzle, tt.args.nth); got != tt.want {
				t.Errorf("laserBeam() = %v, want %v", got, tt.want)
			}
		})
	}
}
