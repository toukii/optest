package optest

import (
	"testing"

	"git.ezbuy.me/ezbuy/base/misc/log"
)

func TestOpt(t *testing.T) {
	Init()
	aids, vids, r := Search(As[0])
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)

	Init()
	aids, vids, r = Search(As[1])
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)

	Init()
	aids, vids, r = Search(As[2])
	t.Logf("aids:%+v, vids:%+v, reduce:%d, length:%d", aids, vids, r, length)
	log.JSON(As)
}
