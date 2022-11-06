package arithmetic

import "testing"

func TestAdd(t *testing.T) {
  if Add(1, 2) != 3 {
    t.Error("1 + 2 did not equal 3")
  }
}

func TestSubtract(t *testing.T) {
  if Subtract(8, 4) != 4 {
    t.Error("8 - 4 did not equal 4")
  }
}

func TestMult(t *testing.T) {
  if Mult(5, 4) != 20 {
    t.Error("5 * 4 did not equal 20")
  }
}
