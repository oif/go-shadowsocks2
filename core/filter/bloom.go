package filter

import (
	"sync"

	"github.com/willf/bloom"
)

var (
	filterInstance *bloom.BloomFilter
	itemCount      = 0
	mux            = sync.Mutex{}
)

const itemSize = 1e6

func init() {
	bloom.NewWithEstimates(itemSize, 0.003)
}

func Add(data []byte) {
	mux.Lock()
	if itemCount > itemSize {
		filterInstance.ClearAll()
	}
	itemCount++
	mux.Unlock()
	filterInstance.Add(data)
}

func Check(data []byte) bool {
	return filterInstance.Test(data)
}
