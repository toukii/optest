package optest

import (
	"testing"

	"git.ezbuy.me/ezbuy/base/misc/log"
)

func TestOpt(t *testing.T) {
	Init()
	aids, vids, r := Search(As[0], As)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)

	Init()
	aids, vids, r = Search(As[1], As)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)

	Init()
	aids, vids, r = Search(As[2], As)
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)
}
