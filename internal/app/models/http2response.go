package models

//Http2Response represents http2 response
type Http2Response struct {
	SrcHost string
	SrcPort string
	DstHost string
	DstPort string
	Message interface{}
}

//NewHttp2Response creates new Http2Request
func NewHttp2Response(packet *Packet, stream *Stream, grpcMessage interface{}) *Http2Response {
	return &Http2Response{
		packet.NetSrc,
		packet.TransportSrc,
		packet.NetDst,
		packet.TransportDst,
		grpcMessage,
	}
}

//GetSrcHost ...
func (r *Http2Response) GetSrcHost() string {
	return r.SrcHost
}

//GetSrcPort ...
func (r *Http2Response) GetSrcPort() string {
	return r.SrcPort
}

//GetDstHost ...
func (r *Http2Response) GetDstHost() string {
	return r.DstHost
}

//GetDstPort ...
func (r *Http2Response) GetDstPort() string {
	return r.DstPort
}

//GetPath ...
func (r *Http2Response) GetPath() string {
	return ""
}

//GetBody ...
func (r *Http2Response) GetBody() interface{} {
	return r.Message
}

//GetHeaders ...
func (r *Http2Response) GetHeaders() map[string]string {
	return make(map[string]string)
}
