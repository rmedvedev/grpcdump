package models

import "github.com/rmedvedev/grpcdump/internal/pkg/config"

//Http2Response represents http2 response
type Http2Response struct {
	SrcHost     string
	SrcPort     string
	DstHost     string
	DstPort     string
	Message     interface{}
	MetaHeaders map[string]string
}

//NewHttp2Response creates new Http2Request
func NewHttp2Response(packet *Packet, stream *Stream, grpcMessage interface{}) *Http2Response {
	return &Http2Response{
		SrcHost:     packet.NetSrc,
		SrcPort:     packet.TransportSrc,
		DstHost:     packet.NetDst,
		DstPort:     packet.TransportDst,
		Message:     grpcMessage,
		MetaHeaders: stream.MetaHeaders,
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
	renderMetaHeaders := make(map[string]string)
	logMetaHeaders := config.GetConfig().GetLogMetaHeaders()
	if len(logMetaHeaders) > 0 {
		if _, ok := logMetaHeaders["*"]; ok {
			renderMetaHeaders = r.MetaHeaders
		} else {
			for name, val := range r.MetaHeaders {
				if _, ok := logMetaHeaders[name]; ok {
					renderMetaHeaders[name] = val
				}
			}
		}
	}

	return renderMetaHeaders
}
