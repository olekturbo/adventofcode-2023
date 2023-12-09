package main

import (
	"reflect"
	"testing"
)

func Test_extrapolate(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				s: []int{0, 3, 6, 9, 12, 15},
			},
			want: [][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
				{0, 0, 0},
				{0, 0},
				{0},
			},
		},
		{
			name: "2",
			args: args{
				s: []int{10, 13, 16, 21, 30, 45},
			},
			want: [][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
				{0},
			},
		},
		{
			name: "3",
			args: args{
				s: []int{1, 3, 6, 10, 15, 21},
			},
			want: [][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
				{0, 0},
				{0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extrapolate(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extrapolate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_takeLast(t *testing.T) {
	type args struct {
		s [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				s: [][]int{
					{10, 13, 16, 21, 30, 45},
					{3, 3, 5, 9, 15},
					{0, 2, 4, 6},
					{2, 2, 2},
					{0, 0},
					{0},
				},
			},
			want: []int{45, 15, 6, 2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := takeLast(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("takeLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_history(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: []int{
					45, 15, 6, 2, 0,
				},
			},
			want: 68,
		},
		{
			name: "2",
			args: args{
				s: []int{
					21, 6, 1, 0,
				},
			},
			want: 28,
		},
		{
			name: "3",
			args: args{
				s: []int{
					15, 3, 0,
				},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := history(tt.args.s); got != tt.want {
				t.Errorf("history() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
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
					"0 3 6 9 12 15",
					"1 3 6 10 15 21",
					"10 13 16 21 30 45",
				},
			},
			want: [][]int{
				{0, 3, 6, 9, 12, 15},
				{1, 3, 6, 10, 15, 21},
				{10, 13, 16, 21, 30, 45},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
