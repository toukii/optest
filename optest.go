package optest

import (
	"fmt"
	"sort"
)

type V struct {
	Id     int64
	Reduce int
	UsedBy string
}

type A struct {
	Id    string
	Vs    []*V
	Fee   int
	VUsed bool
}

func (a *A) LeftFee() int {
	fee := a.Fee
	for _, v := range a.Vs {
		if v.UsedBy != a.Id {
			continue
		}
		fee -= v.Reduce
	}
	if fee >= 0 {
		return fee
	}
	return fee
}

func (a *A) NotUsedVouchers() []*V {
	if len(a.Vs) <= 0 {
		return nil
	}
	vs := make([]*V, 0, len(a.Vs))
	for _, v := range a.Vs {
		if v.UsedBy != "" {
			continue
		}
		vs = append(vs, v)
	}
	return vs
}

type Road struct {
	AIds   []string
	VIds   []int64
	Reduce int
}

var (
	vs     []*V
	As     []*A
	length int
	_      = fmt.Sprint()
)

func Init() {
	vs = []*V{
		&V{1, 100, ""},
		&V{2, 200, ""},
		&V{3, 300, ""},
	}
	As = []*A{
		&A{Id: "A1", Vs: []*V{vs[0], vs[1], vs[2]}, Fee: 800},
		&A{Id: "A2", Vs: []*V{vs[1], vs[2]}, Fee: 800},
		&A{Id: "A3", Vs: []*V{vs[1], vs[2]}, Fee: 200},
	}

	length = 0
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func Search(act *A, acts []*A, voucherM map[int64]*V) *Road {
	length++
	if length > 100 {
		panic(length)
	}
	if act.VUsed {
		return nil
	}
	leftfee := act.LeftFee()
	if leftfee <= 0 {
		return nil
	}
	leftVouchers := act.NotUsedVouchers()
	if len(leftVouchers) <= 0 {
		return nil
	}

	act.VUsed = true

	size := len(acts) - 1
	var leftActs []*A
	if size > 0 {
		leftActs = make([]*A, 0, size)
		for _, _a := range acts {
			if _a.VUsed {
				continue
			}
			leftActs = append(leftActs, _a)
		}
	}
	roads := make([]*Road, 0, len(leftVouchers))
	for _, v := range leftVouchers {
		if v.UsedBy != "" {
			continue
		}
		realReduce := min(leftfee, v.Reduce)
		// 使用voucher
		v.UsedBy = act.Id

		road := &Road{
			AIds:   []string{act.Id},
			VIds:   []int64{v.Id},
			Reduce: realReduce,
		}
		roads = append(roads, road)
		if len(acts) > 1 {
			for _, _a := range leftActs {
				subroad := Search(_a, leftActs, voucherM)
				if subroad != nil {
					road.AIds = append(road.AIds, subroad.AIds...)
					road.VIds = append(road.VIds, subroad.VIds...)
					road.Reduce += subroad.Reduce
				}
			}
		}

		// 清空使用voucher
		v.UsedBy = ""
	}
	sort.Slice(roads, func(i, j int) bool {
		if roads[i] == nil {
			return false
		}
		if roads[j] == nil {
			return true
		}
		return roads[i].Reduce > roads[j].Reduce
	})
	if roads[0] == nil {
		act.VUsed = false
		return nil
	}

	// 全局voucher
	voucherM[roads[0].VIds[0]].UsedBy = act.Id

	return roads[0]
}

func SearchOpt(acts []*A, vouchers []*V) *Road {
	vm := make(map[int64]*V, len(vouchers))
	for _, v := range vouchers {
		vm[v.Id] = v
	}

	roads := make([]*Road, 0, len(acts))
	for _, act := range acts {
		Clear(acts, vouchers)
		road := Search(act, acts, vm)
		if road == nil {
			continue
		}
		roads = append(roads, road)
		// aids, vids, r := road.AIds, road.VIds, road.Reduce
		// fmt.Printf("aids:%+v, vids:%+v, reduce:%d, length:%d\n", aids, vids, r, length)
	}
	if len(roads) <= 0 {
		return nil
	}
	sort.Slice(roads, func(i, j int) bool {
		return roads[i].Reduce > roads[j].Reduce
	})

	return roads[0]
}

func Clear(acts []*A, vouchers []*V) {
	length = 0
	for _, act := range acts {
		act.VUsed = false
	}
	for _, v := range vouchers {
		v.UsedBy = ""
	}
}
