package main

import (
	"reflect"
	"testing"
)

func Test_calibrate(t *testing.T) {
	type args struct {
		slice    []string
		extended bool
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				slice: []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
				},
				extended: false,
			},
			want: [][]int{
				{1, 2},
				{3, 8},
				{1, 5},
				{7, 7},
			},
		},
		{
			name: "2",
			args: args{
				slice: []string{
					"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen",
				},
				extended: true,
			},
			want: [][]int{
				{2, 9},
				{8, 3},
				{1, 3},
				{2, 4},
				{4, 2},
				{1, 4},
				{7, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calibrate(tt.args.slice, tt.args.extended); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calibrate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateCalibration(t *testing.T) {
	type args struct {
		slice    []string
		extended bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				slice: []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
				},
				extended: false,
			},
			want: 142,
		},
		{
			name: "2",
			args: args{
				slice: []string{
					"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen",
				},
				extended: true,
			},
			want: 281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCalibration(tt.args.slice, tt.args.extended); got != tt.want {
				t.Errorf("calculateCalibration() = %v, want %v", got, tt.want)
			}
		})
	}
}
