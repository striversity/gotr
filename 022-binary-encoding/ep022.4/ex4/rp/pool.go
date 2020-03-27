package rp

import "benc/model"
import "fmt"

const CR_POOL_SIZE = 300

var pool chan *model.ClientReq
var total, alloced, reused uint

func init() {
	pool = make(chan *model.ClientReq, CR_POOL_SIZE)
}

func Alloc() *model.ClientReq {
	total++
	select {
	case cr := <-pool:
		reused++
		return cr
	default:
		alloced++
		cr := &model.ClientReq{}
		return cr
	}
}

func Release(cr *model.ClientReq) {
	select {
	case pool <- cr:
	default:
	}
}

func Stats() {
	fmt.Printf("Total: %v, Allocated: %v, Reused: %v\n", total, alloced, reused)
}
