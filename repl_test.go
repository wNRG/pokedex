package main

import (
	"reflect"
	"testing"
) 

func TestCleanInput(t *testing.T) {
	tests := map[string] struct {
		input string
		want []string
	}{
		"simple": {input: "hello world", want: []string{"hello", "world"}},
		"single word": {input: "  hello ", want: []string{"hello"}},
		"empty": {input: "  ", want: []string{}},
		"capitalized": {input: "Gengar Darkrai Rayquaza", want: []string{"gengar", "darkrai", "rayquaza"}},
		"multi-whitespace": {input: "Gengar   Darkrai   rayquaza", want: []string {"gengar", "darkrai", "rayquaza"}},
	}

	for name, tc := range tests {
		got := cleanInput(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}


