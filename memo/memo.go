package memo

import (
	"github.com/golang/glog"
	"github.com/shirou/gopsutil/mem"
	"sync"
)

var oldMem float64

//获取内存占用百分比
func memPercent() (float64, error) {
	vm, err := mem.VirtualMemory()
	return vm.UsedPercent / 100, err
}

//计算当前内存与之前内存占用情况的差异
func MemDiff() (float64, error) {
	newM, err := memPercent()
	return newM - oldMem, err
}

//懒汉模式，保证oldmem每次执行只在最开始计算一次。
func init() {
	o := new(sync.Once)
	o.Do(func() {
		var err error
		oldMem, err = memPercent()
		if err != nil {
			glog.Error(err)
		}
	})
}
