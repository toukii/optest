package optest

import (
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

type Road struct {
	AIds   []string
	VIds   []int64
	Reduce int
}

var (
	vs     []*V
	As     []*A
	length int
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

func Search(a *A) ([]string, []int64, int) {
	length++
	if length > 100 {
		panic(length)
	}
	if a.VUsed {
		return nil, nil, 0
	}
	leftfee := a.LeftFee()
	if leftfee <= 0 {
		return nil, nil, 0
	}
	if len(a.Vs) <= 0 {
		return nil, nil, 0
	}
	// len 可以斟酌
	roads := make([]*Road, 0, len(a.Vs))
	for _, v := range a.Vs {
		if v.UsedBy != "" {
			continue
		}
		realReduce := min(leftfee, v.Reduce)
		// 使用voucher
		v.UsedBy = a.Id
		a.VUsed = true

		road := &Road{
			AIds:   []string{a.Id},
			VIds:   []int64{v.Id},
			Reduce: realReduce,
		}
		roads = append(roads, road)
		for _, _a := range As {
			aids, vids, reduce := Search(_a)
			road.AIds = append(road.AIds, aids...)
			road.VIds = append(road.VIds, vids...)
			road.Reduce += reduce
		}

		// 清空使用voucher
		v.UsedBy = ""
		a.VUsed = false
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
		return nil, nil, 0
	}
	for _, a := range As {
		if a.Id == roads[0].AIds[0] {
			a.VUsed = true
		}
	}
	for _, v := range vs {
		if v.Id == roads[0].VIds[0] {
			v.UsedBy = roads[0].AIds[0]
		}
	}

	return roads[0].AIds, roads[0].VIds, roads[0].Reduce
}
