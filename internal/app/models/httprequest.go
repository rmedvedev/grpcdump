package models

//HttpRequest represents http request
type HttpRequest struct {
	SrcHost string
	SrcPort string
	DstHost string
	DstPort string
	Method  string
	URL     string
}

//NewHttpRequest creates new HttpRequest model
func NewHttpRequest(packet *Packet) *HttpRequest {
	return &HttpRequest{
		SrcHost: packet.NetSrc,
		SrcPort: packet.TransportSrc,
		DstHost: packet.NetDst,
		DstPort: packet.TransportDst,
	}
}

//GetSrcHost ...
func (r *HttpRequest) GetSrcHost() string {
	return r.SrcHost
}

//GetSrcPort ...
func (r *HttpRequest) GetSrcPort() string {
	return r.SrcPort
}

//GetDstHost ...
func (r *HttpRequest) GetDstHost() string {
	return r.DstHost
}

//GetDstPort ...
func (r *HttpRequest) GetDstPort() string {
	return r.DstPort
}

//GetPath ...
func (r *HttpRequest) GetPath() string {
	return r.Method + ": " + r.URL
}

//GetBody ...
func (r *HttpRequest) GetBody() interface{} {
	return nil
}

//GetHeaders ...
func (r *HttpRequest) GetHeaders() map[string]string {
	return make(map[string]string)
}
