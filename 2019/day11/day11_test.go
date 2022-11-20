package day11

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRobot_handleOutput(t *testing.T) {
	type want struct {
		Position Point
		Facing   int
		Hull     map[Point]int
	}
	tests := []struct {
		name   string
		output []int
		want   want
	}{
		{"example step 1", []int{1, 0}, want{Point{-1, 0}, Left, map[Point]int{{0, 0}: 1}}},
		{"example step 2", []int{1, 0, 0, 0}, want{Point{-1, -1}, Down, map[Point]int{{0, 0}: 1, {-1, 0}: 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			robot := newRobot("99")
			for _, v := range tt.output {
				robot.handleOutput(v)
			}
			got := want{
				Position: robot.Position,
				Facing:   robot.Facing,
				Hull:     robot.Hull,
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Robot.handleOutput() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
