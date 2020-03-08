package gtrends

import "testing"

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
