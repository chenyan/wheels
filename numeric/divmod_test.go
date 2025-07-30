package numeric

import "testing"

func TestDivmod(t *testing.T) {
	a, b := Divmod(10, 3)
	if a != 3 {
		t.Errorf("expected 3, got %d", a)
	}
	if b != 1 {
		t.Errorf("expected 1, got %d", b)
	}

	a, b = Divmod(-10, 2)
	if a != -5 {
		t.Errorf("expected -5, got %d", a)
	}
	if b != 0 {
		t.Errorf("expected 0, got %d", b)
	}
}
