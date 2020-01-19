package models

import (
	"github.com/rmedvedev/grpcdump/internal/pkg/config"
)

//Http2Request represents http request
type Http2Request struct {
	SrcHost     string
	SrcPort     string
	DstHost     string
	DstPort     string
	GrpcPath    string
	Message     interface{}
	MetaHeaders map[string]string
}

//NewHttp2Request creates new Http2Request
func NewHttp2Request(packet *Packet, stream *Stream, grpcMessage interface{}) *Http2Request {
	return &Http2Request{
		SrcHost:     packet.NetSrc,
		SrcPort:     packet.TransportSrc,
		DstHost:     packet.NetDst,
		DstPort:     packet.TransportDst,
		GrpcPath:    stream.Path,
		Message:     grpcMessage,
		MetaHeaders: stream.MetaHeaders,
	}
}

//GetSrcHost ...
func (r *Http2Request) GetSrcHost() string {
	return r.SrcHost
}

//GetSrcPort ...
func (r *Http2Request) GetSrcPort() string {
	return r.SrcPort
}

//GetDstHost ...
func (r *Http2Request) GetDstHost() string {
	return r.DstHost
}

//GetDstPort ...
func (r *Http2Request) GetDstPort() string {
	return r.DstPort
}

//GetPath ...
func (r *Http2Request) GetPath() string {
	return r.GrpcPath
}

//GetBody ...
func (r *Http2Request) GetBody() interface{} {
	return r.Message
}

//GetHeaders ...
func (r *Http2Request) GetHeaders() map[string]string {
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
