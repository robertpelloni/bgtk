package core

import (
	"testing"
)

func TestPropertyAnimationLinearInterpolation(t *testing.T) {
	anim := NewPropertyAnimation(0.0, 100.0, 1000)

	val0, fin0 := anim.Step(0)
	if val0 != 0.0 || fin0 {
		t.Errorf("expected 0.0 / false, got %f / %v", val0, fin0)
	}

	valMid, finMid := anim.Step(500)
	if valMid != 50.0 || finMid {
		t.Errorf("expected 50.0 / false at midpoint, got %f / %v", valMid, finMid)
	}

	valEnd, finEnd := anim.Step(1000)
	if valEnd != 100.0 || !finEnd {
		t.Errorf("expected 100.0 / true at completion, got %f / %v", valEnd, finEnd)
	}
}
