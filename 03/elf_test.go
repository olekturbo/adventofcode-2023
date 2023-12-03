package main

import (
	"reflect"
	"testing"
)

func Test_findSymbols(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{
			name: "1",
			args: args{
				s: []string{
					"467..114..",
					"...*......",
					"..35..633.",
					"......#...",
					"617*......",
					".....+.58.",
					"..592.....",
					"......755.",
					"...$.*....",
					".664.598..",
				},
			},
			want: []point{
				{x: 3, y: 1},
				{x: 6, y: 3},
				{x: 3, y: 4},
				{x: 5, y: 5},
				{x: 3, y: 8},
				{x: 5, y: 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSymbols(tt.args.s, isSymbol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAll() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_findStars(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{
			name: "1",
			args: args{
				s: []string{
					"467..114..",
					"...*......",
					"..35..633.",
					"......#...",
					"617*......",
					".....+.58.",
					"..592.....",
					"......755.",
					"...$.*....",
					".664.598..",
				},
			},
			want: []point{
				{x: 3, y: 1},
				{x: 3, y: 4},
				{x: 5, y: 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSymbols(tt.args.s, isStar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAll() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_findNumbers(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []number
	}{
		{
			name: "1",
			args: args{
				s: []string{
					"467..114..",
					"...*......",
					"..35..633.",
					"......#...",
					"617*......",
					".....+.58.",
					"..592.....",
					"......755.",
					"...$.*....",
					".664.598..",
				},
			},
			want: []number{
				{value: 467, points: []point{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}}},
				{value: 114, points: []point{{x: 5, y: 0}, {x: 6, y: 0}, {x: 7, y: 0}}},
				{value: 35, points: []point{{x: 2, y: 2}, {x: 3, y: 2}}},
				{value: 633, points: []point{{x: 6, y: 2}, {x: 7, y: 2}, {x: 8, y: 2}}},
				{value: 617, points: []point{{x: 0, y: 4}, {x: 1, y: 4}, {x: 2, y: 4}}},
				{value: 58, points: []point{{x: 7, y: 5}, {x: 8, y: 5}}},
				{value: 592, points: []point{{x: 2, y: 6}, {x: 3, y: 6}, {x: 4, y: 6}}},
				{value: 755, points: []point{{x: 6, y: 7}, {x: 7, y: 7}, {x: 8, y: 7}}},
				{value: 664, points: []point{{x: 1, y: 9}, {x: 2, y: 9}, {x: 3, y: 9}}},
				{value: 598, points: []point{{x: 5, y: 9}, {x: 6, y: 9}, {x: 7, y: 9}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumContactPoints(t *testing.T) {
	type args struct {
		numbers []number
		symbols []point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				symbols: []point{
					{x: 3, y: 1},
					{x: 6, y: 3},
					{x: 3, y: 4},
					{x: 5, y: 5},
					{x: 3, y: 8},
					{x: 5, y: 8},
				},
				numbers: []number{
					{value: 467, points: []point{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}}},
					{value: 114, points: []point{{x: 5, y: 0}, {x: 6, y: 0}, {x: 7, y: 0}}},
					{value: 35, points: []point{{x: 2, y: 2}, {x: 3, y: 2}}},
					{value: 633, points: []point{{x: 6, y: 2}, {x: 7, y: 2}, {x: 8, y: 2}}},
					{value: 617, points: []point{{x: 0, y: 4}, {x: 1, y: 4}, {x: 2, y: 4}}},
					{value: 58, points: []point{{x: 7, y: 5}, {x: 8, y: 5}}},
					{value: 592, points: []point{{x: 2, y: 6}, {x: 3, y: 6}, {x: 4, y: 6}}},
					{value: 755, points: []point{{x: 6, y: 7}, {x: 7, y: 7}, {x: 8, y: 7}}},
					{value: 664, points: []point{{x: 1, y: 9}, {x: 2, y: 9}, {x: 3, y: 9}}},
					{value: 598, points: []point{{x: 5, y: 9}, {x: 6, y: 9}, {x: 7, y: 9}}},
				},
			},
			want: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumContactPoints(tt.args.numbers, tt.args.symbols); got != tt.want {
				t.Errorf("sumContactPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplyContactPoints(t *testing.T) {
	type args struct {
		numbers []number
		symbols []point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				numbers: []number{
					{value: 467, points: []point{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}}},
					{value: 114, points: []point{{x: 5, y: 0}, {x: 6, y: 0}, {x: 7, y: 0}}},
					{value: 35, points: []point{{x: 2, y: 2}, {x: 3, y: 2}}},
					{value: 633, points: []point{{x: 6, y: 2}, {x: 7, y: 2}, {x: 8, y: 2}}},
					{value: 617, points: []point{{x: 0, y: 4}, {x: 1, y: 4}, {x: 2, y: 4}}},
					{value: 58, points: []point{{x: 7, y: 5}, {x: 8, y: 5}}},
					{value: 592, points: []point{{x: 2, y: 6}, {x: 3, y: 6}, {x: 4, y: 6}}},
					{value: 755, points: []point{{x: 6, y: 7}, {x: 7, y: 7}, {x: 8, y: 7}}},
					{value: 664, points: []point{{x: 1, y: 9}, {x: 2, y: 9}, {x: 3, y: 9}}},
					{value: 598, points: []point{{x: 5, y: 9}, {x: 6, y: 9}, {x: 7, y: 9}}},
				},
				symbols: []point{
					{x: 3, y: 1},
					{x: 3, y: 4},
					{x: 5, y: 8},
				},
			},
			want: 467835,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplyContactPoints(tt.args.numbers, tt.args.symbols); got != tt.want {
				t.Errorf("multiplyContactPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
