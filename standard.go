package bloomfilter

import (
	"github.com/spaolacci/murmur3"
	"hash"
	"math"
	"time"
)

type StandardBloom struct {
	elems uint          // currently inserted elements
	set   []byte        // Bitset that holds the elements
	m     uint          // Size of bloom filter biteset
	k     uint          // Number of hash functions
	fpr   float64       // False positive rate
	n     uint          // Expected number of items in the collection
	h     []hash.Hash32 // Hash functions
}

func New(n uint, fpr float64) *StandardBloom {
	m := uint(math.Ceil(float64(n) * math.Abs(math.Log(fpr)) / math.Pow(math.Log(2), float64(2))))
	k := uint(math.Ceil((float64(m) / float64(n)) * math.Log(2)))
	return &StandardBloom{
		n:     n,
		m:     m,
		k:     k,
		fpr:   fpr,
		elems: 0,
		h:     fns(k),
		set:   make([]byte, m),
	}
}

func fns(k uint) []hash.Hash32 {
	h := []hash.Hash32{}
	ts := uint(time.Now().Unix())
	for i := uint(0); i < k; i++ {
		h = append(h, murmur3.New32WithSeed(uint32(ts+1)))
	}
	return h
}

func prepare(hfn hash.Hash32, key string, size uint) uint32 {
	hfn.Write([]byte(key))
	idx := hfn.Sum32() % uint32(size)
	hfn.Reset()
	return idx
}

func (sbf *StandardBloom) Add(key string) {
	for _, hfn := range sbf.h {
		idx := prepare(hfn, key, sbf.m)
		sbf.set[idx] = 1
	}
	sbf.elems++ // add to number of elements in the set
}

func (sbf *StandardBloom) Test(key string) bool {
	for _, hfn := range sbf.h {
		idx := prepare(hfn, key, sbf.m)
		if sbf.set[idx] != 1 {
			return false
		}
	}
	return true
}

func (sbf *StandardBloom) Data() []byte {
	return sbf.set
}
