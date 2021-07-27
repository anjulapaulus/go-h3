package H3Go

import (
	"math"
	"testing"

	"github.com/uber/h3-go"
)

func TestH3Index(t *testing.T) {
	result := "87283472bffffff"
	index := H3Index(37.3615593, -122.0553238, 7)
	if h3.ToString(index) != result {
		t.Error("Failed TestH3Index")
	}
}

func TestKRing(t *testing.T) {
	origin := H3Index(37.3615593, -122.0553238, 8)
	resultArray := KRing(origin, 8, 3)

	h3_array := h3.KRing(origin, int(math.Floor(float64(3)/(cellEdgeLength[8]*2))))

	if len(resultArray) != len(h3_array) {
		t.Error("Failed TestKRing", len(resultArray))
	}
}

func BenchmarkH3Index(b *testing.B) {
	for i := 0; i < b.N; i++ {
		H3Index(37.3615593, -122.0553238, 8)
	}
}

func BenchmarkKRing(b *testing.B) {
	origin := H3Index(37.3615593, -122.0553238, 8)
	for i := 0; i < b.N; i++ {
		KRing(origin, 8, 4)
	}
}
