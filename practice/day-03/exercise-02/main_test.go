package main

import (
	"reflect"
	"testing"
)

// Проверки наставника; порядок проверяется внутри каждой группы.
func TestGroupByDepartment(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []User
		want map[string][]string
	}{
		{"nil", nil, map[string][]string{}},
		{"empty", []User{}, map[string][]string{}},
		{"order and empty department", []User{{"Ann", "backend"}, {"Bob", "frontend"}, {"Cat", "backend"}, {"Dan", ""}}, map[string][]string{"backend": {"Ann", "Cat"}, "frontend": {"Bob"}, "": {"Dan"}}},
		{"duplicates and empty name", []User{{"Ann", "backend"}, {"Ann", "backend"}, {"", "backend"}}, map[string][]string{"backend": {"Ann", "Ann", ""}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := groupByDepartment(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %#v, want %#v", got, tc.want)
			}
		})
	}
}
