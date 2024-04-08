package eldup

import (
	"reflect"
	"testing"
)

// type TstCases[T any] struct {
// 	origin        []T
// 	eliminatedDup []T
// }

// func TestEliminateDup(t *testing) {

// }

type StrCases struct {
	origin        []string
	eliminatedDup []string
}

func TestEliminateDup(t *testing.T) {
	strCases := []StrCases{
		{
			origin:        []string{"a", "a", "a", "b", "b", "b", "c", "c", "c"},
			eliminatedDup: []string{"a", "b", "c"},
		},
		{
			origin:        []string{"m", "a", "a", "a", "m", "m", "a"},
			eliminatedDup: []string{"m", "a", "m", "a"},
		},
	}

	t.Run("string slices", func(t *testing.T) {
		for _, strCase := range strCases {
			want := strCase.eliminatedDup
			got := EliminateDup(strCase.origin)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}

type TstCasesGen[T comparable] struct {
	origin        []T
	eliminatedDup []T
}

func TestEliminateDupGenericComparable(t *testing.T) {

	strCases := []TstCasesGen[string]{
		{
			origin:        []string{"a", "a", "a", "b", "b", "b", "c", "c", "c"},
			eliminatedDup: []string{"a", "b", "c"},
		},
		{
			origin:        []string{"m", "a", "a", "a", "m", "m", "a"},
			eliminatedDup: []string{"m", "a", "m", "a"},
		},
	}

	t.Run("string slices", func(t *testing.T) {
		for _, strCase := range strCases {
			want := strCase.eliminatedDup
			got := EliminateDupGenericComparable(strCase.origin)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})

	intCases := []TstCasesGen[int]{
		{
			origin:        []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			eliminatedDup: []int{1, 2, 3},
		},
		{
			origin:        []int{300, 300, 300, 300, 300},
			eliminatedDup: []int{300},
		},
	}

	t.Run("int slices", func(t *testing.T) {
		for _, intCase := range intCases {
			want := intCase.eliminatedDup
			got := EliminateDupGenericComparable(intCase.origin)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}
