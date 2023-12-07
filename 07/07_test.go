package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []hand
	}{
		{
			name: "1",
			args: args{input: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
				"22552 201",
				"66336 201",
				"77337 201",
				"88338 201",
				"99449 201",
			}},
			want: []hand{
				{bet: 765, values: []int{3, 2, 10, 3, 13}, rank: Pair},
				{bet: 684, values: []int{10, 5, 5, 11, 5}, rank: ThreeOfKind},
				{bet: 28, values: []int{13, 13, 6, 7, 7}, rank: TwoPairs},
				{bet: 220, values: []int{13, 10, 11, 11, 10}, rank: TwoPairs},
				{bet: 483, values: []int{12, 12, 12, 11, 14}, rank: ThreeOfKind},
				{bet: 201, values: []int{2, 2, 5, 5, 2}, rank: FullHouse},
				{bet: 201, values: []int{6, 6, 3, 3, 6}, rank: FullHouse},
				{bet: 201, values: []int{7, 7, 3, 3, 7}, rank: FullHouse},
				{bet: 201, values: []int{8, 8, 3, 3, 8}, rank: FullHouse},
				{bet: 201, values: []int{9, 9, 4, 4, 9}, rank: FullHouse},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bid(t *testing.T) {
	type args struct {
		hands []hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				hands: []hand{
					{bet: 765, values: []int{3, 2, 10, 3, 13}, rank: Pair},
					{bet: 684, values: []int{10, 5, 5, 11, 5}, rank: ThreeOfKind},
					{bet: 28, values: []int{13, 13, 6, 7, 7}, rank: TwoPairs},
					{bet: 220, values: []int{13, 10, 11, 11, 10}, rank: TwoPairs},
					{bet: 483, values: []int{12, 12, 12, 11, 14}, rank: ThreeOfKind},
				},
			},
			want: 6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bid(tt.args.hands); got != tt.want {
				t.Errorf("bid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sort(t *testing.T) {
	type args struct {
		hands []hand
	}
	tests := []struct {
		name string
		args args
		want []hand
	}{
		{
			name: "1",
			args: args{
				hands: []hand{
					{bet: 28, values: []int{13, 13, 6, 7, 7}, rank: TwoPairs},
					{bet: 28, values: []int{15, 13, 6, 7, 7}, rank: TwoPairs},
					{bet: 28, values: []int{21, 13, 6, 7, 7}, rank: TwoPairs},
					{bet: 28, values: []int{21, 21, 6, 7, 7}, rank: TwoPairs},
					{bet: 220, values: []int{13, 10, 11, 11, 10}, rank: TwoPairs},
					{bet: 220, values: []int{13, 10, 11, 11, 15}, rank: TwoPairs},
					{bet: 220, values: []int{13, 10, 11, 18, 15}, rank: TwoPairs},
				},
			},
			want: []hand{
				{bet: 28, values: []int{21, 21, 6, 7, 7}, rank: TwoPairs},
				{bet: 28, values: []int{21, 13, 6, 7, 7}, rank: TwoPairs},
				{bet: 28, values: []int{15, 13, 6, 7, 7}, rank: TwoPairs},
				{bet: 28, values: []int{13, 13, 6, 7, 7}, rank: TwoPairs},
				{bet: 220, values: []int{13, 10, 11, 18, 15}, rank: TwoPairs},
				{bet: 220, values: []int{13, 10, 11, 11, 15}, rank: TwoPairs},
				{bet: 220, values: []int{13, 10, 11, 11, 10}, rank: TwoPairs},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.hands, len(tt.args.hands)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
