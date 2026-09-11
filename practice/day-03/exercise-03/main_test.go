package main

import (
	"reflect"
	"testing"
)

// Проверки наставника: считаются кодовые точки, а не байты.
func TestRuneFrequency(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want map[rune]int
	}{
		{"empty", "", map[rune]int{}},
		{"ASCII", "aba", map[rune]int{'a': 2, 'b': 1}},
		{"Cyrillic and emoji", "мама🙂м", map[rune]int{'м': 3, 'а': 2, '🙂': 1}},
		{"emoji repeats", "🙂🙂", map[rune]int{'🙂': 2}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := runeFrequency(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}
