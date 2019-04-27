package rp

import "thread-pool/model"

const CR_POOL_SIZE = 10000

var pool chan *model.ClientReq

func init() {
	pool = make(chan *model.ClientReq, CR_POOL_SIZE)
}

func Alloc() *model.ClientReq {
	select {
	case cr := <-pool:
		return cr
	default:
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
