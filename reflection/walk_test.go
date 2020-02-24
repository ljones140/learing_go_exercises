package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Lewis"},
			[]string{"Lewis"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Lewis", "London"},
			[]string{"Lewis", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Lewis", 39},
			[]string{"Lewis"},
		},
		{
			"Nested fields",
			Person{
				"Lewis",
				Profile{39, "London"},
			},
			[]string{"Lewis", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
