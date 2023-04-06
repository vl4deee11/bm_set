package bm_set

import "sync"

type BMSyncSet struct {
	masks []uint64
	zero  bool
	sync.RWMutex
}

type SyncSetI interface {
	sync.Locker
	Set(i int)
	Get(i int) bool
	Delete(i int)
	Intersect(oth *BMSyncSet) SyncSetI
	Union(oth *BMSyncSet) SyncSetI
}

func NewSync(size uint64) SyncSetI {
	c := size / 64
	if size%64 != 0 {
		c++
	}
	return &BMSyncSet{
		masks: make([]uint64, c),
	}
}

func (s *BMSyncSet) Set(i int) {
	s.Lock()
	defer s.Unlock()
	if i == 0 {
		s.zero = true
		return
	}
	bn, k := s.getSettings(i)
	s.masks[bn] = s.masks[bn] | k
}

func (s *BMSyncSet) Get(i int) bool {
	s.RLock()
	defer s.RUnlock()
	if i == 0 {
		return s.zero
	}
	bn, k := s.getSettings(i)
	return s.masks[bn]&k != 0
}

func (s *BMSyncSet) Delete(i int) {
	s.Lock()
	defer s.Unlock()
	if i == 0 {
		s.zero = false
		return
	}
	bn, k := s.getSettings(i)
	s.masks[bn] = s.masks[bn] & (^k)
}

func (s *BMSyncSet) getSettings(i int) (int, uint64) {
	bn := i / 64
	if i%64 != 0 {
		bn++
	}
	if bn > 0 {
		bn--
	}

	return bn, uint64(1 << (i % 64))
}

func (s *BMSyncSet) Intersect(oth *BMSyncSet) SyncSetI {
	oth.RLock()
	s.RLock()
	defer func() {
		oth.RUnlock()
		s.RUnlock()
	}()

	ll := len(oth.masks)
	if len(s.masks) > ll {
		ll = len(s.masks)
	}
	masks := make([]uint64, ll)
	for i := 0; i < ll; i++ {
		if i < len(s.masks) && i < len(oth.masks) {
			masks[i] = s.masks[i] | oth.masks[i]
		} else if i < len(s.masks) {
			masks[i] = s.masks[i]
		} else if i < len(oth.masks) {
			masks[i] = oth.masks[i]
		}
	}
	if s.zero || oth.zero {
		return &BMSyncSet{zero: true, masks: masks}
	}
	return &BMSyncSet{zero: false, masks: masks}
}

func (s *BMSyncSet) Union(oth *BMSyncSet) SyncSetI {
	oth.RLock()
	s.RLock()
	defer func() {
		oth.RUnlock()
		s.RUnlock()
	}()

	ll := len(oth.masks)
	if len(s.masks) < ll {
		ll = len(s.masks)
	}
	masks := make([]uint64, ll)
	for i := 0; i < ll; i++ {
		masks[i] = s.masks[i] & oth.masks[i]
	}
	if s.zero && oth.zero {
		return &BMSyncSet{zero: true, masks: masks}
	}
	return &BMSyncSet{zero: false, masks: masks}
}
