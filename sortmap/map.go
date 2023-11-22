package sortmap

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type mapElement[A any, O constraints.Ordered] struct {
	key  O
	data A
}

type SortedMap[A any, O constraints.Ordered] struct {
	elems []mapElement[A, O]
}

func mapCmp[A any, O constraints.Ordered](a mapElement[A, O], b mapElement[A, O]) int {
	if a.key < b.key {
		return -1
	}

	if a.key == b.key {
		return 0
	}
	return 1
}

func mapCmpKey[A any, O constraints.Ordered](a mapElement[A, O], key O) int {
	if a.key < key {
		return -1
	}

	if a.key == key {
		return 0
	}
	return 1
}

func (s *SortedMap[A, O]) Init() {
	s.elems = make([]mapElement[A,O], 0)
}

func (s *SortedMap[A, O]) Insert(key O, data A) {
	elem := mapElement[A, O]{key, data}
	s.elems = append(s.elems, elem)
	slices.SortFunc(s.elems, mapCmp[A, O])
}

func (s SortedMap[A, O]) SubMap(begin O, end O) {
	slices.BinarySearchFunc(s.elems, begin, mapCmpKey[A, O])
}
