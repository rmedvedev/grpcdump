package models

//RenderModel ...
type RenderModel interface {
	GetSrcHost() string
	GetSrcPort() string
	GetDstHost() string
	GetDstPort() string
	GetPath() string
	GetBody() interface{}
	GetHeaders() map[string]string
}
