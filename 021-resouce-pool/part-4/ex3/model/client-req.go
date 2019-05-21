package model

const (
	ReqAdd        = iota
	ReqAvg        = iota
	ReqRandom     = iota
	ReqSpellCheck = iota
	ReqSearch     = iota
)

// ReqDataSize is the max bytes per ClentReq.Data byte array
const ReqDataSize = 1 * 1024 // 1kb

type (
	// ClientReq represents a request from a client with work to
	// to be done by the server
	ClientReq struct {
		ID      uint
		ReqType int               // one of ReqX defined above
		Data    [ReqDataSize]byte // request specific encoded data
		Size    int               // how many byte in Data
	}
)
