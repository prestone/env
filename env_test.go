package env

//testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)
	_ = a

	a.Equal("ok", String("nice", "ok"))
	a.Equal(
		[]string{"1", "2", "3"},
		Strings("nice", ",", "1", "2", "3"),
	)
	a.Equal(9999, Int("port", 9999))
	a.Equal(float64(-1.4235), Float("long", -1.4235))
	a.True(Bool("dev", true))
	a.True(Bool("dev", 1))
	a.False(Bool("dev", 0))
	a.False(Bool("dev", false))

}

func BenchmarkString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		String("nice", "ok")
	}
}

func BenchmarkInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Int("nice", 42)
	}
}

func BenchmarkStrings(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Strings("nice", ",", "1", "2", "3")
	}
}

func BenchmarkBool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Bool("dev", true)
	}
}
