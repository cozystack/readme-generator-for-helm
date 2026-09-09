package main

import (
	"reflect"
	"testing"
)

func TestInferType(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  string
	}{
		{"nil", nil, "nil"},
		{"string", "hello", "string"},
		{"bool", true, "boolean"},
		{"int", 1, "number"},
		{"int64", int64(1), "number"},
		{"float64", 1.5, "number"},
		{"array", []interface{}{1, 2}, "array"},
		{"object", map[string]interface{}{"a": 1}, "object"},
		{"unknown", struct{}{}, "unknown"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := inferType(c.input); got != c.want {
				t.Errorf("inferType(%v) = %q, want %q", c.input, got, c.want)
			}
		})
	}
}

func TestGetArrayPrefix(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"no brackets", "a.b.c", "a.b.c"},
		{"single index", "a.b[0]", "a.b"},
		{"index not at start", "a[0].b", "a"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := getArrayPrefix(c.input); got != c.want {
				t.Errorf("getArrayPrefix(%q) = %q, want %q", c.input, got, c.want)
			}
		})
	}
}

func TestSanitizeProperty(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"no brackets", "a.b.c", "a.b.c"},
		{"with index", "a.b[0].c", "a.b"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := sanitizeProperty(c.input); got != c.want {
				t.Errorf("sanitizeProperty(%q) = %q, want %q", c.input, got, c.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		name string
		a, b []string
		want []string
	}{
		{"no overlap", []string{"a", "b"}, []string{"c"}, []string{"a", "b"}},
		{"full overlap", []string{"a", "b"}, []string{"a", "b"}, nil},
		{"partial overlap", []string{"a", "b", "c"}, []string{"b"}, []string{"a", "c"}},
		{"empty a", nil, []string{"a"}, nil},
		{"empty b", []string{"a"}, nil, []string{"a"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := difference(c.a, c.b)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("difference(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
			}
		})
	}
}
