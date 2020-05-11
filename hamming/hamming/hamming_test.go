package hamming

import "testing"

func TestHamming(t *testing.T) {
	type test struct {
		strand1 string
		strand2 string
		answer  int
	}

	tests := []test{
		test{"ABCDE", "ABCDE", 0},
		test{"ABCDE", "ABCDD", 1},
		test{"ABCDE", "ABCCC", 2},
		test{"ABCDE", "FFFFF", 5},
	}

	for _, v := range tests {
		x := Hamming(v.strand1, v.strand2)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}

func BenchmarkHamming(b *testing.B) {
	Hamming("ABCDE", "ABCDE")
}
