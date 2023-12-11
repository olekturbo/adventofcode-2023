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
		want [][]int
	}{
		{
			name: "1",
			args: args{
				s: []string{
					"...#......",
					".......#..",
					"#.........",
					"..........",
					"......#...",
					".#........",
					".........#",
					"..........",
					".......#..",
					"#...#.....",
				},
			},
			want: [][]int{
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0},
				{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0},
				{0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 0},
				{8, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0},
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

func Test_points(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{
			name: "1",
			args: args{
				graph: [][]int{
					{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0},
					{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0},
					{0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 0},
					{8, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: []point{
				{x: 4, y: 0, v: 1},
				{x: 9, y: 1, v: 2},
				{x: 0, y: 2, v: 3},
				{x: 8, y: 5, v: 4},
				{x: 1, y: 6, v: 5},
				{x: 12, y: 7, v: 6},
				{x: 9, y: 10, v: 7},
				{x: 0, y: 11, v: 8},
				{x: 5, y: 11, v: 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := points(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("points() = %v, want %v", got, tt.want)
			}
		})
	}
}
