package day10

import (
	_ "embed"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/maths"
	"github.com/richardc/advent-go/runner"
	slcs "github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   10,
			Part1: func(any) any { return mostAsteroids(puzzle) },
		},
	)
}

type Point struct {
	Row int
	Col int
}

func (p *Point) Distance(other Point) int {
	return maths.AbsDiff(p.Row, other.Row) + maths.AbsDiff(p.Col, other.Col)
}

type Field struct {
	rows, cols int
	asteroids  []Point
}

func newField(s string) Field {
	var asteroids []Point
	for r, line := range input.Lines(s) {
		for c, char := range line {
			if char == '#' {
				asteroids = append(asteroids, Point{r, c})
			}
		}
	}

	return Field{
		rows:      len(input.Lines(s)),
		cols:      len(input.Lines(s)[0]),
		asteroids: asteroids,
	}
}

func (f *Field) mostAsteroids() int {
	visible := map[Point]int{}
	for _, asteroid := range f.asteroids {
		visible[asteroid] = f.visibleFrom(asteroid)
	}
	// for r := 0; r < f.rows; r++ {
	// 	for c := 0; c < f.cols; c++ {
	// 		fmt.Print(visible[Point{r, c}])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println(visible)
	best := slcs.MaxBy(maps.Keys(visible), func(p Point) int { return visible[p] })
	return visible[best]
}

func (f *Field) contains(p Point) bool {
	if p.Col < 0 || p.Row < 0 || p.Col >= f.cols || p.Row >= f.rows {
		return false
	}
	return true
}

func (f *Field) visibleFrom(asteroid Point) int {
	others := slcs.Filter(f.asteroids, func(p Point) bool { return p != asteroid })
	slices.SortFunc(others, func(a, b Point) bool { return asteroid.Distance(a) < asteroid.Distance(b) })
	visible := map[Point]struct{}{}
	for _, v := range others {
		visible[v] = struct{}{}
	}

	for _, point := range others {
		if _, ok := visible[point]; !ok {
			continue
		}
		// delete everything behind this
		for _, hidden := range f.project(asteroid, point) {
			delete(visible, hidden)
		}
	}

	return len(visible)
}

func (f *Field) project(origin, point Point) []Point {
	vector := Point{
		Row: point.Row - origin.Row,
		Col: point.Col - origin.Col,
	}
	gcd := maths.GCD(vector.Row, vector.Col)
	vector.Row /= gcd
	vector.Col /= gcd

	var generate func(int) Point
	switch {
	case vector.Row == 0:
		generate = func(i int) Point {
			return Point{
				Row: point.Row,
				Col: point.Col + i*maths.Signum(vector.Col),
			}
		}
	case vector.Col == 0:
		generate = func(i int) Point {
			return Point{
				Row: point.Row + i*maths.Signum(vector.Row),
				Col: point.Col,
			}
		}
	default:
		generate = func(i int) Point {
			return Point{
				Row: point.Row + vector.Row*i,
				Col: point.Col + vector.Col*i,
			}
		}
	}

	var points []Point
	i := 1
	for {
		hide := generate(i)
		if !f.contains(hide) {
			break
		}
		points = append(points, hide)
		i++
	}
	return points
}

func mostAsteroids(puzzle string) int {
	field := newField(puzzle)
	return field.mostAsteroids()
}
