package optest

import (
	"testing"

	"git.ezbuy.me/ezbuy/base/misc/log"
)

// func TestOpt(t *testing.T) {
// 	Init()
// 	vm := make(map[int64]*V, len(vs))
// 	for _, v := range vs {
// 		vm[v.Id] = v
// 	}

// 	road := Search(As[0], As, vm)
// 	aids, vids, r := road.AIds, road.VIds, road.Reduce
// 	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
// 	log.JSON(As)

// 	Init()
// 	road = Search(As[1], As, vm)
// 	aids, vids, r = road.AIds, road.VIds, road.Reduce
// 	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
// 	log.JSON(As)

// 	Init()
// 	road = Search(As[2], As, vm)
// 	aids, vids, r = road.AIds, road.VIds, road.Reduce
// 	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
// 	log.JSON(As)
// }

func TestA(*testing.T) {
	Init()
	road := SearchOpt(As, vs)
	log.JSON(road)

	// road1 := SearchOpt(As, vs)
	// log.JSON(road1)
}

func BenchmarkOptest(b *testing.B) {
	Init()
	for i := 0; i < b.N; i++ {
		road := SearchOpt(As, vs)
		if road.Reduce != 600 {
			b.Errorf("road:%+v", road)
		}
	}
}
