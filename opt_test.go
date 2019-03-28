package optest

import (
	"testing"
)

func TestOpt(t *testing.T) {
	Init()
	vm := make(map[int64]*V, len(Vouchers))
	for _, v := range Vouchers {
		vm[v.Id] = v
	}

	road := Search(Acts[0], Acts, vm)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", road.AIds, road.VIds, road.Reduce, length)

	Init()
	road = Search(Acts[1], Acts, vm)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", road.AIds, road.VIds, road.Reduce, length)

	Init()
	road = Search(Acts[2], Acts, vm)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", road.AIds, road.VIds, road.Reduce, length)
}

func TestA(t *testing.T) {
	Init()

	road := SearchOpt(Acts, Vouchers)
	t.Logf("road:%+v, length:%d", road, length)

	road1 := SearchOpt(Acts, Vouchers)
	t.Logf("road:%+v, length:%d", road1, length)
}

func BenchmarkOptest(b *testing.B) {
	Init()
	for i := 0; i < b.N; i++ {
		road := SearchOpt(Acts, Vouchers)
		if road.Reduce != 600 {
			b.Errorf("road:%+v", road)
		}
	}
}
