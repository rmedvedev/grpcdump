package models

const (
	//RequestType ...
	RequestType = 1
	//ResponseType ...
	ResponseType = 2
)

//Stream ...
type Stream struct {
	ID                  uint32
	MetaHeaders         map[string]string
	Path                string
	Type                int
	GrpcState           GrpcState
	ResponseGrpcMessage interface{}
}

type GrpcState struct {
	IsPartialRead bool
	Buf           []byte
}
