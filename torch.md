go test -bench . -cpuprofile=cpu.prof

go test -bench=BenchmarkOptest -cpuprofile cpu2.prof -memprofile mem2.prof

#分析cpu
go-torch optest.test cpu2.prof
#分析内存
go-torch --colors=mem -alloc_space optest.test mem2.prof
go-torch --colors=mem -inuse_space optest.test mem.prof

go-torch --alloc_objects optest.test mem.prof