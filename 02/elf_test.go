package main

import (
	"reflect"
	"testing"
)

func Test_extractGames(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want games
	}{
		{
			name: "1",
			args: args{
				s: []string{
					"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
					"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
					"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
					"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
					"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
				},
			},
			want: games{
				1: {
					subset{"blue": 3, "green": 0, "red": 4},
					subset{"blue": 6, "green": 2, "red": 1},
					subset{"blue": 0, "green": 2, "red": 0},
				},
				2: {
					subset{"blue": 1, "green": 2, "red": 0},
					subset{"blue": 4, "green": 3, "red": 1},
					subset{"blue": 1, "green": 1, "red": 0},
				},
				3: {
					subset{"blue": 6, "green": 8, "red": 20},
					subset{"blue": 5, "green": 13, "red": 4},
					subset{"blue": 0, "green": 5, "red": 1},
				},
				4: {
					subset{"blue": 6, "green": 1, "red": 3},
					subset{"blue": 0, "green": 3, "red": 6},
					subset{"blue": 15, "green": 3, "red": 14},
				},
				5: {
					subset{"blue": 1, "green": 3, "red": 6},
					subset{"blue": 2, "green": 2, "red": 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractGames(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_count(t *testing.T) {
	type args struct {
		g games
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				g: games{
					1: {
						subset{"blue": 3, "green": 0, "red": 4},
						subset{"blue": 6, "green": 2, "red": 1},
						subset{"blue": 0, "green": 2, "red": 0},
					},
					2: {
						subset{"blue": 1, "green": 2, "red": 0},
						subset{"blue": 4, "green": 3, "red": 1},
						subset{"blue": 1, "green": 1, "red": 0},
					},
					3: {
						subset{"blue": 6, "green": 8, "red": 20},
						subset{"blue": 5, "green": 13, "red": 4},
						subset{"blue": 0, "green": 5, "red": 1},
					},
					4: {
						subset{"blue": 6, "green": 1, "red": 3},
						subset{"blue": 0, "green": 3, "red": 6},
						subset{"blue": 15, "green": 3, "red": 14},
					},
					5: {
						subset{"blue": 1, "green": 3, "red": 6},
						subset{"blue": 2, "green": 2, "red": 1},
					},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.g); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countB(t *testing.T) {
	type args struct {
		g games
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				g: games{
					1: {
						subset{"blue": 3, "green": 0, "red": 4},
						subset{"blue": 6, "green": 2, "red": 1},
						subset{"blue": 0, "green": 2, "red": 0},
					},
					2: {
						subset{"blue": 1, "green": 2, "red": 0},
						subset{"blue": 4, "green": 3, "red": 1},
						subset{"blue": 1, "green": 1, "red": 0},
					},
					3: {
						subset{"blue": 6, "green": 8, "red": 20},
						subset{"blue": 5, "green": 13, "red": 4},
						subset{"blue": 0, "green": 5, "red": 1},
					},
					4: {
						subset{"blue": 6, "green": 1, "red": 3},
						subset{"blue": 0, "green": 3, "red": 6},
						subset{"blue": 15, "green": 3, "red": 14},
					},
					5: {
						subset{"blue": 1, "green": 3, "red": 6},
						subset{"blue": 2, "green": 2, "red": 1},
					},
				},
			},
			want: 2286,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countB(tt.args.g); got != tt.want {
				t.Errorf("countB() = %v, want %v", got, tt.want)
			}
		})
	}
}
