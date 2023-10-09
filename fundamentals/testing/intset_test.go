package testing

import (
	"reflect"
	"testing"

	m "github.com/Artemides/problems/fundamentals/oop/methods"
)

func TestIntset(t *testing.T) {
	var set1 m.IntSet
	var set2 m.IntSet

	set1.AddAll(4, 5, 6, 3, 10, 11)
	set2.AddAll(2, 6, 7, 5, 3, 2, 3)

	testCases := []struct {
		sets     []*m.IntSet
		method   interface{}
		expected interface{}
	}{
		{
			sets:     []*m.IntSet{m.NewInset(1, 2), m.NewInset(3, 4)},
			method:   (*m.IntSet).Union,
			expected: m.NewInset(1, 2, 3, 4),
		},
	}

	for _, testCase := range testCases {
		set1 := testCase.sets[0]
		set2 := testCase.sets[1]

		fn := reflect.TypeOf(testCase.method)
		unionFn := (*m.IntSet).Union
		if fn == reflect.TypeOf(unionFn) {
			set1Str := set1.String()

			unionFn(set1, set2)
			expected, ok := testCase.expected.(*m.IntSet)
			if !ok {
				t.Errorf(``)
				continue
			}

			if set1.String() != expected.String() {
				t.Errorf(`%+v Union %+v = %+v ; expected: %+v `, set1Str, set2, set1.String(), expected.String())
			}
		}
	}

}
