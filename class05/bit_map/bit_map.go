package bit_map

type BitMap struct {
	Bits []uint64
}

func NewBitMap(max int) *BitMap {
	return &BitMap{
		Bits: make([]uint64, (max+64)>>6),
	}
}

func (bm *BitMap) Add(num int) {
	bm.Bits[num>>6] |= 1 << (num & 63)
}

func (bm *BitMap) Delete(num int) {
	bm.Bits[num>>6] &= ^(1 << (num & 63))
}

func (bm *BitMap) Contains(num int) bool {
	return bm.Bits[num>>6]&(1<<(num&63)) != 0
}
