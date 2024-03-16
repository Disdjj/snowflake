package snowflake

import (
	"testing"
)

func BenchmarkParallel10(b *testing.B) {
	idGene := NewIDGenerator(WithRent())
	b.SetParallelism(10)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}

func BenchmarkParallel100(b *testing.B) {
	idGene := NewIDGenerator(WithRent())
	b.SetParallelism(100)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}

func BenchmarkParallel1000(b *testing.B) {
	idGene := NewIDGenerator(WithRent())
	b.SetParallelism(1000)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}

func BenchmarkParallel10000(b *testing.B) {
	idGene := NewIDGenerator(WithRent())
	b.SetParallelism(10000)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}

func BenchmarkParallel100000(b *testing.B) {
	idGene := NewIDGenerator(WithRent())
	b.SetParallelism(100000)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}
