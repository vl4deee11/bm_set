# Bitmask set 

#### small memory usage => ((max_value / 64) + 1) * 8 bytes
```go
// usage

package main

import (
	"fmt"
	"github.com/vl4deee11/bm_set"
)

func main() {
	sz := 131
	bms := bm_set.New(uint64(sz))
	for i := 0; i <= sz; i++ {

		bms.Set(i)
		ok := bms.Get(i)
		fmt.Printf("after set, element = %d, ok = %v", i, ok)

		bms.Delete(i)
		ok = bms.Get(i)
		fmt.Printf("after delete, element = %d, ok = %v", i, ok)
	}
}
```