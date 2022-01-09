package main

import (
	"reflect"
	"testing"
)

func TestUniqueRunes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			name: "empty",
			args: args{""},
			want: map[rune]int{},
		},
		{
			name: "one letter",
			args: args{"a"},
			want: map[rune]int{'a': 1},
		},
		{
			name: "one letter repeated",
			args: args{"aaaaa"},
			want: map[rune]int{'a': 5},
		},
		{
			name: "multiple letters",
			args: args{"abc"},
			want: map[rune]int{'a': 1, 'b': 1, 'c': 1},
		},
		{
			name: "multiple letters repeated",
			args: args{"abcabcabcd"},
			want: map[rune]int{'a': 3, 'b': 3, 'c': 3, 'd': 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueRunes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
