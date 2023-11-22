package sortmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	var sMap SortedMap[string, int] 
	var expMap SortedMap[string, int] = SortedMap[string, int]{
		elems: []mapElement[string, int]{
			{key: 1, data: "first"},
			{key: 2, data: "second"},
			{key: 3, data: "third"},
		},
	} 

	sMap.Init()

	sMap.Insert(1, "first")
	sMap.Insert(3, "third")
	sMap.Insert(2, "second")
	
	assert.Equal(t, expMap, sMap)
}
