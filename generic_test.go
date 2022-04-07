package genericBench

import (
	"testing"
)

func sum(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		res := a.(int) + b.(int)
		return &res
	case float32:
		res := a.(float32) + b.(float32)
		return &res
	case float64:
		res := a.(float64) + b.(float64)
		return &res
	default:
		return nil
	}
}

func sumInts(a, b int) *int {
	res := a + b
	return &res
}

func sumFloats32(a, b float32) *float32 {
	res := a + b
	return &res
}

func sumFloats64(a, b float64) *float64 {
	res := a + b
	return &res
}

func sumGeneric[T int | float32 | float64](a, b T) *T {
	res := a + b
	return &res
}

type typesSet interface {
	int | float32 | float64
}

func BenchmarkSum(b *testing.B) {
	integers := make([]int, 100_000_001)
	floats32 := make([]float32, 100_000_001)
	floats64 := make([]float64, 100_000_001)

	for i := 0; i < b.N; i++ {
		integers[i] = 1
		floats32[i] = float32(1)
		floats64[i] = float64(1)
	}

	b.Run("SumSwitchInt", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sum(integers[i], integers[i])
			_ = r
		}
	})

	b.Run("SumInts", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sumInts(integers[i], integers[i])
			_ = r
		}
	})

	b.Run("SumGenericInts", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sumGeneric(integers[i], integers[i])
			_ = r
		}
	})

	b.Run("SumSwitchFloat32", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sum(floats32[i], floats32[i])
			_ = r
		}
	})

	b.Run("SumFloats32", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sumFloats32(floats32[i], floats32[i])
			_ = r
		}
	})

	b.Run("SumGenericFloats32", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sumGeneric(floats32[i], floats32[i])
			_ = r
		}
	})

	b.Run("SumSwitchFloat64", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sum(floats64[i], floats64[i])
			_ = r
		}
	})

	b.Run("SumFloats64", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sumFloats64(floats64[i], floats64[i])
			_ = r
		}
	})

	b.Run("SumGenericFloats64", func(b *testing.B) {
		var r interface{}
		for i := 0; i < b.N; i++ {
			r = sumGeneric(floats64[i], floats64[i])
			_ = r
		}
	})

}
