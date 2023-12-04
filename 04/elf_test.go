package main

import (
	"reflect"
	"testing"
)

func Test_toScratches(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []scratch
	}{
		{
			name: "1",
			args: args{
				s: []string{
					"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
					"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
					"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
					"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
					"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
					"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
				},
			},
			want: []scratch{
				{id: 1, winning: []int{41, 48, 83, 86, 17}, owned: []int{83, 86, 6, 31, 17, 9, 48, 53}},
				{id: 2, winning: []int{13, 32, 20, 16, 61}, owned: []int{61, 30, 68, 82, 17, 32, 24, 19}},
				{id: 3, winning: []int{1, 21, 53, 59, 44}, owned: []int{69, 82, 63, 72, 16, 21, 14, 1}},
				{id: 4, winning: []int{41, 92, 73, 84, 69}, owned: []int{59, 84, 76, 51, 58, 5, 54, 83}},
				{id: 5, winning: []int{87, 83, 26, 28, 32}, owned: []int{88, 30, 70, 12, 93, 22, 82, 36}},
				{id: 6, winning: []int{31, 18, 13, 56, 72}, owned: []int{74, 77, 10, 23, 35, 67, 36, 11}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toScratches(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toScratches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_count(t *testing.T) {
	type args struct {
		w []int
		o []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				w: []int{41, 48, 83, 86, 17},
				o: []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				w: []int{13, 32, 20, 16, 61},
				o: []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				w: []int{1, 21, 53, 59, 44},
				o: []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				w: []int{41, 92, 73, 84, 69},
				o: []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				w: []int{87, 83, 26, 28, 32},
				o: []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			want: 0,
		},
		{
			name: "6",
			args: args{
				w: []int{31, 18, 13, 56, 72},
				o: []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.w, tt.args.o); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worth(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				c: 0,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				c: 1,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				c: 2,
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				c: 3,
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				c: 4,
			},
			want: 8,
		},
		{
			name: "6",
			args: args{
				c: 5,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := worth(tt.args.c); got != tt.want {
				t.Errorf("worth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		scratches []scratch
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				scratches: []scratch{
					{id: 1, winning: []int{41, 48, 83, 86, 17}, owned: []int{83, 86, 6, 31, 17, 9, 48, 53}},
					{id: 2, winning: []int{13, 32, 20, 16, 61}, owned: []int{61, 30, 68, 82, 17, 32, 24, 19}},
					{id: 3, winning: []int{1, 21, 53, 59, 44}, owned: []int{69, 82, 63, 72, 16, 21, 14, 1}},
					{id: 4, winning: []int{41, 92, 73, 84, 69}, owned: []int{59, 84, 76, 51, 58, 5, 54, 83}},
					{id: 5, winning: []int{87, 83, 26, 28, 32}, owned: []int{88, 30, 70, 12, 93, 22, 82, 36}},
					{id: 6, winning: []int{31, 18, 13, 56, 72}, owned: []int{74, 77, 10, 23, 35, 67, 36, 11}},
				},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.scratches); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
