package pool_app

import (
	"testing"
)

//
//func BenchmarkLogWithPool(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			poolValue := Get()
//			poolValue.Dto.Min = utilities.RandomRange(1, 100)
//			poolValue.Dto.Max = utilities.RandomRange(1, 100)
//		}
//	})
//}

func BenchmarkLogWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		obj := new(Dto)
		// use obj
		obj.Max = 1000000000
	}
}
