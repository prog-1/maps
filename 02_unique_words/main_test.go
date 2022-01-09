package main

import (
	"reflect"
	"testing"
)

func TestUniqueWords(t *testing.T) {
	tests := []struct {
		name string
		args string
		want map[string]int
	}{
		{
			name: "empty",
			args: "",
			want: map[string]int{},
		},
		{
			name: "separators",
			args: " @1!`~",
			want: map[string]int{},
		},
		{
			name: "single",
			args: "abc",
			want: map[string]int{"abc": 1},
		},
		{
			name: "single repeated",
			args: "+++abc123abc!@#abc+++",
			want: map[string]int{"abc": 3},
		},
		{
			name: "multiple",
			args: "abc cba",
			want: map[string]int{"abc": 1, "cba": 1},
		},
		{
			name: "multiple repeated",
			args: "abc cba abc cba abc",
			want: map[string]int{"abc": 3, "cba": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueWords(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
