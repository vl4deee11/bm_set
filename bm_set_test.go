package bm_set

import (
	"testing"
)

// BenchmarkBMSet        100000000               10.6 ns/op             0 B/op          0 allocs/op
func BenchmarkBMSet(b *testing.B) {
	sz := uint64(b.N)
	bms := New(sz)
	for i := 0; i < b.N; i++ {
		bms.Set(i)
		v := bms.Get(i)
		if v {

		}
		bms.Delete(i)
	}
}

func TestBMSet1(t *testing.T) {
	sz := 131
	bms := New(uint64(sz))
	for i := 0; i <= sz; i++ {
		bms.Set(i)
		if !bms.Get(i) {
			t.Errorf("test !bms.Get(i) fail on = %d", i)
			return
		}
	}

	// Delete x & 1 == 0
	for i := 0; i <= sz; i++ {
		if !bms.Get(i) {
			t.Errorf("test !bms.Get(i) fail on = %d", i)
			return
		}
		if i&1 == 0 {
			bms.Delete(i)
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		}
	}

	// Delete x & 1 != 0
	for i := 0; i <= sz; i++ {
		if i&1 != 0 {
			if !bms.Get(i) {
				t.Errorf("test !bms.Get(i) fail on = %d", i)
				return
			}
			bms.Delete(i)
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		} else {
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		}
	}

	for i := 0; i <= sz; i++ {
		if bms.Get(i) {
			t.Errorf("test bms.Get(i) fail on = %d", i)
			return
		}
	}
}

func TestBMSet2(t *testing.T) {
	sz := 128
	bms := New(uint64(sz))
	for i := 0; i <= sz; i++ {
		bms.Set(i)
		if !bms.Get(i) {
			t.Errorf("test !bms.Get(i) fail on = %d", i)
			return
		}
	}

	// Delete x & 1 == 0
	for i := 0; i <= sz; i++ {
		if !bms.Get(i) {
			t.Errorf("test !bms.Get(i) fail on = %d", i)
			return
		}
		if i&1 == 0 {
			bms.Delete(i)
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		}
	}

	// Delete x & 1 != 0
	for i := 0; i <= sz; i++ {
		if i&1 != 0 {
			if !bms.Get(i) {
				t.Errorf("test !bms.Get(i) fail on = %d", i)
				return
			}
			bms.Delete(i)
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		} else {
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		}
	}

	for i := 0; i <= sz; i++ {
		if bms.Get(i) {
			t.Errorf("test bms.Get(i) fail on = %d", i)
			return
		}
	}
}
