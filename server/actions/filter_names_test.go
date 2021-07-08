package actions

import (
	"os"
	"reflect"
	"testing"
)

func TestFilterNamesWithValue(t *testing.T) {
	t.Run("convert env variable to array of filter names", func(t *testing.T) {
		filterNames := "a,b,c"
		os.Setenv("FILTER_NAMES", filterNames)

		got := GetFilterNames()
		want := []string{"a", "b", "c"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given, \"%v\"", got, want, os.Getenv("FILTER_NAMES"))
		}
	})

	t.Run("check if singleton is working", func(t *testing.T) {
		filterNames := "a"
		os.Setenv("FILTER_NAMES", filterNames)

		got := GetFilterNames()
		want := []string{"a", "b", "c"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given, \"%v\"", got, want, os.Getenv("FILTER_NAMES"))
		}
	})
}

func TestFilterNamesWithEmpty(t *testing.T) {
	t.Run("return array with len equal 0 when FILTER_NAMES is empty", func(t *testing.T) {
		os.Setenv("FILTER_NAMES", "")

		got := GetFilterNames()
		want := []string{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given, \"%v\"", got, want, os.Getenv("FILTER_NAMES"))
		}
	})
}

func TestCheckNameInList(t *testing.T) {
	t.Run("name in list name", func(t *testing.T) {
		name := "abc"
		list := []string{"abc", "def", "ghi"}

		got := CheckNameInList(list, name)
		want := true

		if got != want {
			t.Errorf("got %v want %v given, name: \"%v\" and list %v", got, want, name, list)
		}
	})

	t.Run("name not in list name", func(t *testing.T) {
		name := "abc"
		list := []string{"def", "ghi"}

		got := CheckNameInList(list, name)
		want := false

		if got != want {
			t.Errorf("got %v want %v given, name: \"%v\" and list %v", got, want, name, list)
		}
	})
}
