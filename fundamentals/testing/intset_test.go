package testing

import (
	"reflect"
	"testing"

	m "github.com/Artemides/problems/fundamentals/oop/methods"
)

type insetTestCase struct {
	receiver *m.IntSet
	params   interface{}
	method   interface{}
	expected interface{}
}

var intsetTestCases = []insetTestCase{
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
	{
		receiver: m.NewInset(12, 22),
		params:   22,
		method:   (*m.IntSet).Has,
		expected: true,
	},
}

func TestIntset(t *testing.T) {
	for _, testCase := range intsetTestCases {

		fn := reflect.TypeOf(testCase.method)

		if unionFn := (*m.IntSet).Union; fn == reflect.TypeOf(unionFn) {
			insetUnion(unionFn, testCase, t)
			continue
		}

		if addFn := (*m.IntSet).Add; fn == reflect.TypeOf(addFn) {
			intsetAdd(addFn, testCase, t)
			continue
		}

		if hasFunc := (*m.IntSet).Has; fn == reflect.TypeOf(hasFunc) {
			intsetHas(hasFunc, testCase, t)
			continue
		}
	}

}

func insetUnion(unionFn func(*m.IntSet, *m.IntSet), testCase insetTestCase, t *testing.T) {
	set1 := testCase.receiver
	set1Str := set1.String()
	set2, ok := testCase.params.(*m.IntSet)
	if !ok {
		t.Errorf(`invalid Union Params %T`, set2)
		return
	}

	unionFn(set1, set2)
	expected, ok := testCase.expected.(*m.IntSet)
	if !ok {
		t.Errorf(`invalid expected result %T `, expected)
		return
	}

	if set1.String() != expected.String() {
		t.Errorf(`%+v Union %+v = %+v ; expected: %+v `, set1Str, set2, set1.String(), expected.String())
		return
	}
}
func intsetAdd(addFn func(*m.IntSet, int), testCase insetTestCase, t *testing.T) {
	set1 := testCase.receiver
	set1Str := set1.String()
	param, ok := testCase.params.(int)
	if !ok {
		t.Errorf(`invalid Union Params %T`, param)
		return
	}

	addFn(set1, param)
	expected, ok := testCase.expected.(*m.IntSet)
	if !ok {
		t.Errorf(`invalid expected result %T `, expected)
		return
	}

	if set1.String() != expected.String() {
		t.Errorf(`%+v Add %d = %+v ; expected: %+v `, set1Str, param, set1.String(), expected.String())
	}
}
func intsetHas(hasFunc func(*m.IntSet, int) bool, testCase insetTestCase, t *testing.T) {
	set1 := testCase.receiver
	param, ok := testCase.params.(int)
	if !ok {
		t.Errorf(`invalid Union Params %T`, param)
		return
	}

	has := hasFunc(set1, param)
	expected, ok := testCase.expected.(bool)
	if !ok {
		t.Errorf(`invalid expected result %T `, expected)
		return
	}

	if has != expected {
		t.Errorf(`%+v Has %v = %v; expected: %v`, set1.String(), param, has, expected)
	}
}
