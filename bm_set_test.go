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

func TestIntersect(t *testing.T) {
	bm1 := New(128)
	bm1.Set(3)
	bm1.Set(65)
	bm1.Set(120)
	bm1.Set(0)
	bm1.Set(4)

	bm2 := New(67)
	bm2.Set(3)
	bm2.Set(66)

	bm3 := bm1.Intersect(bm2.(*BMSet))
	vals := []int{3, 65, 0, 4, 66, 120}
	for i := 0; i < len(vals); i++ {
		if !bm3.Get(vals[i]) {
			t.Errorf("test bms.Get(i) fail on = %d", vals[i])
			return
		}
		bm3.Delete(vals[i])
	}

	for i := 0; i < 128; i++ {
		if bm3.Get(i) {
			t.Errorf("test bms.Get(i) fail on = %d", i)
			return
		}
	}
}

func TestUnion(t *testing.T) {
	bm1 := New(128)
	bm1.Set(3)
	bm1.Set(66)
	bm1.Set(120)
	bm1.Set(0)
	bm1.Set(4)

	bm2 := New(67)
	bm2.Set(3)
	bm2.Set(66)

	bm3 := bm1.Union(bm2.(*BMSet))
	vals := []int{3, 66}
	for i := 0; i < len(vals); i++ {
		if !bm3.Get(vals[i]) {
			t.Errorf("test bms.Get(i) fail on = %d", vals[i])
			return
		}
		bm3.Delete(vals[i])
	}

	for i := 0; i < 67; i++ {
		if bm3.Get(i) {
			t.Errorf("test bms.Get(i) fail on = %d", i)
			return
		}
	}
}
