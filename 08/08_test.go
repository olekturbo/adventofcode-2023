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
		name  string
		args  args
		want  map[string]collection
		want1 []string
		want2 string
		want3 string
	}{
		{
			name: "1",
			args: args{
				s: []string{
					"RL",
					"",
					"AAA = (BBB, CCC)",
					"BBB = (DDD, EEE)",
					"CCC = (ZZZ, GGG)",
					"DDD = (DDD, DDD)",
					"EEE = (EEE, EEE)",
					"GGG = (GGG, GGG)",
					"ZZZ = (ZZZ, ZZZ)",
				},
			},
			want: map[string]collection{
				"AAA": {left: "BBB", right: "CCC"},
				"BBB": {left: "DDD", right: "EEE"},
				"CCC": {left: "ZZZ", right: "GGG"},
				"DDD": {left: "DDD", right: "DDD"},
				"EEE": {left: "EEE", right: "EEE"},
				"GGG": {left: "GGG", right: "GGG"},
				"ZZZ": {left: "ZZZ", right: "ZZZ"},
			},
			want1: []string{"R", "L"},
		},
		{
			name: "2",
			args: args{
				s: []string{
					"LLR",
					"",
					"AAA = (BBB, BBB)",
					"BBB = (AAA, ZZZ)",
					"ZZZ = (ZZZ, ZZZ)",
				},
			},
			want: map[string]collection{
				"AAA": {left: "BBB", right: "BBB"},
				"BBB": {left: "AAA", right: "ZZZ"},
				"ZZZ": {left: "ZZZ", right: "ZZZ"},
			},
			want1: []string{"L", "L", "R"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parse(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_steps(t *testing.T) {
	type args struct {
		collections map[string]collection
		ways        []string
		start       string
		end         string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				collections: map[string]collection{
					"AAA": {left: "BBB", right: "BBB"},
					"BBB": {left: "AAA", right: "ZZZ"},
					"ZZZ": {left: "ZZZ", right: "ZZZ"},
				},
				ways:  []string{"L", "L", "R"},
				start: "AAA",
				end:   "ZZZ",
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				collections: map[string]collection{
					"AAA": {left: "BBB", right: "CCC"},
					"BBB": {left: "DDD", right: "EEE"},
					"CCC": {left: "ZZZ", right: "GGG"},
					"DDD": {left: "DDD", right: "DDD"},
					"EEE": {left: "EEE", right: "EEE"},
					"GGG": {left: "GGG", right: "GGG"},
					"ZZZ": {left: "ZZZ", right: "ZZZ"},
				},
				ways:  []string{"R", "L"},
				start: "AAA",
				end:   "ZZZ",
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				collections: map[string]collection{
					"AAA": {left: "BBB", right: "BBB"},
					"BBB": {left: "CCC", right: "CCC"},
					"CCC": {left: "DDD", right: "DDD"},
					"DDD": {left: "EEE", right: "EEE"},
					"EEE": {left: "GGG", right: "GGG"},
					"GGG": {left: "YYY", right: "YYY"},
					"YYY": {left: "ZZZ", right: "ZZZ"},
					"ZZZ": {left: "ZZZ", right: "ZZZ"},
				},
				ways:  []string{"R", "L"},
				start: "AAA",
				end:   "ZZZ",
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := steps(tt.args.collections, tt.args.ways, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("steps() = %v, want %v", got, tt.want)
			}
		})
	}
}
