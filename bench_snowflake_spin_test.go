package snowflake

import (
	"testing"
)

func BenchmarkParallelSpin10(b *testing.B) {
	idGene := NewIDGenerator()
	b.SetParallelism(10)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}

func BenchmarkParallelSpin100(b *testing.B) {
	idGene := NewIDGenerator()
	b.SetParallelism(100)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}

func BenchmarkParallelSpin1000(b *testing.B) {
	idGene := NewIDGenerator()
	b.SetParallelism(1000)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}

func BenchmarkParallelSpin10000(b *testing.B) {
	idGene := NewIDGenerator()
	b.SetParallelism(10000)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}

func BenchmarkParallelSpin100000(b *testing.B) {
	idGene := NewIDGenerator()
	b.SetParallelism(100000)
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = idGene.NextID()
			}
		},
	)
}
