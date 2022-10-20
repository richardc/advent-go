package day20

import (
	_ "embed"
	"testing"
)

func Test_staysClosest(t *testing.T) {
	type args struct {
		particles []Particle
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]Particle{
			NewParticle("p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>"),
			NewParticle("p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>"),
		}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := staysClosest(tt.args.particles); got != tt.want {
				t.Errorf("staysClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
