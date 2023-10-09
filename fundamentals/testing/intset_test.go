package testing

import (
	"reflect"
	"testing"

	m "github.com/Artemides/problems/fundamentals/oop/methods"
)

func TestIntset(t *testing.T) {

	testCases := []struct {
		receiver *m.IntSet
		params   interface{}
		method   interface{}
		expected interface{}
	}{
		{
			receiver: m.NewInset(1, 2),
			params:   m.NewInset(3, 4),
			method:   (*m.IntSet).Union,
			expected: m.NewInset(1, 2, 3, 4),
		},
		{
			receiver: m.NewInset(1, 2),
			params:   112,
			method:   (*m.IntSet).Add,
			expected: m.NewInset(1, 2, 112),
		},
	}

	for _, testCase := range testCases {

		fn := reflect.TypeOf(testCase.method)
		unionFn := (*m.IntSet).Union

		if fn == reflect.TypeOf(unionFn) {
			set1 := testCase.receiver
			set1Str := set1.String()
			set2, ok := testCase.params.(*m.IntSet)
			if !ok {
				t.Errorf(`invalid Union Params %T`, set2)
				continue
			}

			unionFn(set1, set2)
			expected, ok := testCase.expected.(*m.IntSet)
			if !ok {
				t.Errorf(`invalid expected result %T `, expected)
				continue
			}

			if set1.String() != expected.String() {
				t.Errorf(`%+v Union %+v = %+v ; expected: %+v `, set1Str, set2, set1.String(), expected.String())
			}

			continue
		}
		addFn := (*m.IntSet).Add
		if fn == reflect.TypeOf(addFn) {
			set1 := testCase.receiver
			set1Str := set1.String()
			param, ok := testCase.params.(int)
			if !ok {
				t.Errorf(`invalid Union Params %T`, param)
				continue
			}

			addFn(set1, param)
			expected, ok := testCase.expected.(*m.IntSet)
			if !ok {
				t.Errorf(`invalid expected result %T `, expected)
				continue
			}

			if set1.String() != expected.String() {
				t.Errorf(`%+v Add %d = %+v ; expected: %+v `, set1Str, param, set1.String(), expected.String())
			}
			continue
		}
	}

}
