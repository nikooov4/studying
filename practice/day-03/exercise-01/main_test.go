package main

import (
	"reflect"
	"testing"
)

// Проверки наставника для кода, перенесённого из отчёта ученика.
func TestWordFrequency(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		want map[string]int
	}{
		{"nil", nil, map[string]int{}},
		{"empty", []string{}, map[string]int{}},
		{"repeated and empty word", []string{"go", "map", "go", "", "map", "go"}, map[string]int{"": 1, "go": 3, "map": 2}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := wordFrequency(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}
