package optest

import (
	"testing"

	"git.ezbuy.me/ezbuy/base/misc/log"
)

func TestOpt(t *testing.T) {
	Init()
	vm := make(map[int64]*V, len(vs))
	for _, v := range vs {
		vm[v.Id] = v
	}

	aids, vids, r := Search(As[0], As, vm)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)

	Init()
	aids, vids, r = Search(As[1], As, vm)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)

	Init()
	aids, vids, r = Search(As[2], As, vm)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)
}

func TestA(*testing.T) {
	Init()
	SearchOpt(As, vs)
}
