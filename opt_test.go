package optest

import (
	"testing"

	"github.com/onsi/ginkgo"
	. "github.com/smartystreets/goconvey/convey"
)

type tCase struct {
	vouchers []*V
	acts     []*A
	road     *Road
}

var (
	vouchers0 = []*V{
		&V{1, 100, ""},
		&V{2, 200, ""},
		&V{3, 300, ""},
	}
	tcase0 = tCase{
		vouchers: vouchers0,
		acts: []*A{
			&A{Id: "A1", Vs: []*V{vouchers0[0], vouchers0[1], vouchers0[2]}, Fee: 800},
			&A{Id: "A2", Vs: []*V{vouchers0[1], vouchers0[2]}, Fee: 800},
			&A{Id: "A3", Vs: []*V{vouchers0[1], vouchers0[2]}, Fee: 200},
		},
		road: &Road{
			AIds:   []string{"A3", "A1", "A2"},
			VIds:   []int64{2, 1, 3},
			Reduce: 600,
		},
	}

	tcases = []tCase{
		tcase0,
		tCase{
			vouchers: vouchers0,
			acts: []*A{
				&A{Id: "A1", Vs: []*V{vouchers0[0], vouchers0[1], vouchers0[2]}, Fee: 800},
				&A{Id: "A2", Vs: []*V{vouchers0[1], vouchers0[2]}, Fee: 800},
				&A{Id: "A3", Vs: []*V{vouchers0[2]}, Fee: 200},
			},
			road: &Road{
				AIds:   []string{"A2", "A1", "A3"},
				VIds:   []int64{2, 1, 3},
				Reduce: 500,
			},
		},

		tCase{
			vouchers: vouchers0,
			acts: []*A{
				&A{Id: "A1", Vs: []*V{vouchers0[0], vouchers0[1], vouchers0[2]}, Fee: 800},
				&A{Id: "A2", Vs: []*V{vouchers0[1], vouchers0[2]}, Fee: 800},
				&A{Id: "A3", Vs: []*V{vouchers0[2]}, Fee: 300},
			},
			road: &Road{
				AIds:   []string{"A1", "A2", "A3"},
				VIds:   []int64{1, 2, 3},
				Reduce: 600,
			},
		},

		tCase{
			vouchers: vouchers0,
			acts: []*A{
				&A{Id: "A1", Vs: []*V{vouchers0[0], vouchers0[1], vouchers0[2]}, Fee: 800},
				&A{Id: "A2", Vs: []*V{vouchers0[1], vouchers0[2]}, Fee: 800},
				&A{Id: "A3", Vs: []*V{vouchers0[2]}, Fee: 0},
			},
			road: &Road{
				AIds:   []string{"A1", "A2"},
				VIds:   []int64{3, 2},
				Reduce: 500,
			},
		},
	}
)

func equalRoad(r1, r2 *Road) {
	So(r1.Reduce, ShouldEqual, r2.Reduce)
	So(len(r1.AIds), ShouldResemble, len(r2.AIds))
	r1m := r1.MapAV()
	r2m := r2.MapAV()

	for _, actId := range r1.AIds {
		So(r1m[actId], ShouldEqual, r2m[actId])
	}
}

func TestOpt(t *testing.T) {
	vm := make(map[int64]*V, len(tcase0.vouchers))
	for _, v := range tcase0.vouchers {
		vm[v.Id] = v
	}

	length = 0
	road := Search(tcase0.acts[0], tcase0.acts, vm)
	t.Logf("road:%+v, length:%d", road, length)

	length = 0
	Clear(tcase0.acts, tcase0.vouchers)
	road = Search(tcase0.acts[1], tcase0.acts, vm)
	t.Logf("road:%+v, length:%d", road, length)

	length = 0
	Clear(tcase0.acts, tcase0.vouchers)
	road = Search(tcase0.acts[2], tcase0.acts, vm)
	t.Logf("road:%+v, length:%d", road, length)
}

func TestSearchOptDup(t *testing.T) {
	ginkgo.Context("SearchOptDup tcase0 first", func() {
		road := SearchOpt(tcase0.acts, tcase0.vouchers)
		Convey("SearchOptDup tcase0 first", t, func() {
			equalRoad(road, tcase0.road)
		})
	})

	ginkgo.Context("SearchOptDup tcase0 second", func() {
		road := SearchOpt(tcase0.acts, tcase0.vouchers)
		Convey("SearchOptDup tcase0 second", t, func() {
			equalRoad(road, tcase0.road)
		})
	})

	ginkgo.Context("SearchOptDup tcase0 third", func() {
		road := SearchOpt(tcase0.acts, tcase0.vouchers)
		Convey("SearchOptDup tcase0 third", t, func() {
			equalRoad(road, tcase0.road)
		})
	})
}
func TestSearchOpt(t *testing.T) {
	for _, tcase := range tcases {
		ginkgo.Context("SearchOpt", func() {
			road := SearchOpt(tcase.acts, tcase.vouchers)
			Convey("SearchOpt", t, func() {
				// So(road, ShouldResemble, tcase.road)
				equalRoad(road, tcase.road)
			})
		})
	}
}

func BenchmarkOptest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchOpt(tcase0.acts, tcase0.vouchers)
	}
}
