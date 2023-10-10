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
	{
		receiver: m.NewInset(12, 22),
		params:   m.NewInset(12, 11, 23),
		method:   (*m.IntSet).IntersectWith,
		expected: m.NewInset(12),
	},
}

func TestIntset(t *testing.T) {
	for _, testCase := range intsetTestCases {

		fn := reflect.TypeOf(testCase.method)

		if unionFn := (*m.IntSet).Union; fn == reflect.TypeOf(unionFn) {
			union(unionFn, testCase, t)
			continue
		}

		if addFn := (*m.IntSet).Add; fn == reflect.TypeOf(addFn) {
			add(addFn, testCase, t)
			continue
		}

		if hasFunc := (*m.IntSet).Has; fn == reflect.TypeOf(hasFunc) {
			has(hasFunc, testCase, t)
			continue
		}
		if intersectWithFunc := (*m.IntSet).IntersectWith; fn == reflect.TypeOf(intersectWithFunc) {
			intersects(intersectWithFunc, testCase, t)
			continue
		}
	}

}

func union(unionFn func(*m.IntSet, *m.IntSet), testCase insetTestCase, t *testing.T) {
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
func add(addFn func(*m.IntSet, int), testCase insetTestCase, t *testing.T) {
	set1 := testCase.receiver
	set1Str := set1.String()
	param, ok := testCase.params.(int)
	if !ok {
		t.Errorf(`invalid Add Params %T`, param)
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
func has(hasFunc func(*m.IntSet, int) bool, testCase insetTestCase, t *testing.T) {
	set1 := testCase.receiver
	param, ok := testCase.params.(int)
	if !ok {
		t.Errorf(`invalid Has Params %T`, param)
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
func intersects(intersectWithFunc func(*m.IntSet, *m.IntSet), testCase insetTestCase, t *testing.T) {
	set1 := testCase.receiver
	set1Str := set1.String()
	param, ok := testCase.params.(*m.IntSet)
	if !ok {
		t.Errorf(`invalid intersectWithFunc Params %T`, param)
		return
	}

	intersectWithFunc(set1, param)
	expected, ok := testCase.expected.(*m.IntSet)
	if !ok {
		t.Errorf(`invalid expected result %T `, expected)
		return
	}

	if set1 != expected {
		t.Errorf(`%+v Intersect %+v = %+v; expected: %+v`, set1Str, param.String(), set1.String(), expected.String())
	}
}
