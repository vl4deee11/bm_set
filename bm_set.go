package bm_set

type BMSet struct {
	masks []uint64
	zero  bool
}

type SetI interface {
	Set(i int)
	Get(i int) bool
	Delete(i int)
}

func New(size uint64) SetI {
	c := size / 64
	if size%64 != 0 {
		c++
	}
	return &BMSet{
		masks: make([]uint64, c),
	}
}

func (s *BMSet) Set(i int) {
	if i == 0 {
		s.zero = true
		return
	}
	bn, k := s.getSettings(i)
	s.masks[bn] = s.masks[bn] | k
}

func (s *BMSet) Get(i int) bool {
	if i == 0 {
		return s.zero
	}
	bn, k := s.getSettings(i)
	return s.masks[bn]&k != 0
}

func (s *BMSet) Delete(i int) {
	if i == 0 {
		s.zero = false
		return
	}
	bn, k := s.getSettings(i)
	s.masks[bn] = s.masks[bn] & (^k)
}

func (s *BMSet) getSettings(i int) (int, uint64) {
	bn := i / 64
	if i%64 != 0 {
		bn++
	}
	if bn > 0 {
		bn--
	}

	return bn, uint64(1 << (i % 64))
}
