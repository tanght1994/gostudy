package thttools

import (
	"fmt"
	"reflect"
	"testing"
)

var data1 = []struct {
	in  int32
	out int32
}{
	{-1, 1},
	{0, 0},
	{1, 1},
}

func TestFun1(t *testing.T) {
	for i, d := range data1 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			out := fun1(d.in)
			if !reflect.DeepEqual(out, d.out) {
				t.Errorf("expected:%v, got:%v", d.out, out)
			}
		})
	}
}

func TestFun2(t *testing.T) {
	g := fun2(1)
	w := int32(1)
	if !reflect.DeepEqual(w, g) {
		t.Errorf("expected:%v, got:%v", w, g)
	}

	g = fun2(0)
	w = int32(0)
	if !reflect.DeepEqual(w, g) {
		t.Errorf("expected:%v, got:%v", w, g)
	}

	g = fun2(-1)
	w = int32(1)
	if !reflect.DeepEqual(w, g) {
		t.Errorf("expected:%v, got:%v", w, g)
	}
}
