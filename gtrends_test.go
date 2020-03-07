package gtrends

import "testing"

func TestIsGtrends(t *testing.T) {
	for _, c := range testCases {
		if TrendsIn(c.input) != c.expected {
			t.Fatalf("FAIL: %s\nWord %q, expected %t, got %t", c.description, c.input, c.expected, !c.expected)
		}

		t.Logf("PASS: Word %q", c.input)
	}
}

func BenchmarkIsGtrends(b *testing.B) {
	b.StopTimer()
	for _, c := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			TrendsIn(c.input)
		}

		b.StopTimer()
	}
}
