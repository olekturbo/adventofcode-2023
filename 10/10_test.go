package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want points
	}{
		{
			name: "1",
			args: args{
				s: []string{
					"..F7.",
					".FJ|.",
					"SJ.L7",
					"|F--J",
					"LJ...",
				},
			},
			want: points{
				{value: ".", x: 0, y: 0},
				{value: ".", x: 1, y: 0},
				{value: "F", x: 2, y: 0},
				{value: "7", x: 3, y: 0},
				{value: ".", x: 4, y: 0},
				{value: ".", x: 0, y: 1},
				{value: "F", x: 1, y: 1},
				{value: "J", x: 2, y: 1},
				{value: "|", x: 3, y: 1},
				{value: ".", x: 4, y: 1},
				{value: "S", x: 0, y: 2},
				{value: "J", x: 1, y: 2},
				{value: ".", x: 2, y: 2},
				{value: "L", x: 3, y: 2},
				{value: "7", x: 4, y: 2},
				{value: "|", x: 0, y: 3},
				{value: "F", x: 1, y: 3},
				{value: "-", x: 2, y: 3},
				{value: "-", x: 3, y: 3},
				{value: "J", x: 4, y: 3},
				{value: "L", x: 0, y: 4},
				{value: "J", x: 1, y: 4},
				{value: ".", x: 2, y: 4},
				{value: ".", x: 3, y: 4},
				{value: ".", x: 4, y: 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
