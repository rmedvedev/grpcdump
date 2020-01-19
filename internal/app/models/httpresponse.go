package models

//HttpResponse represents http response
type HttpResponse struct {
	SrcHost string
	SrcPort string
	DstHost string
	DstPort string
	Body    string
}

//NewHttpResponse creates new HttpResponse model
func NewHttpResponse(packet *Packet) *HttpResponse {
	return &HttpResponse{
		packet.NetSrc,
		packet.TransportSrc,
		packet.NetDst,
		packet.TransportDst,
		"",
	}
}

//GetSrcHost ...
func (r *HttpResponse) GetSrcHost() string {
	return r.SrcHost
}

//GetSrcPort ...
func (r *HttpResponse) GetSrcPort() string {
	return r.SrcPort
}

//GetDstHost ...
func (r *HttpResponse) GetDstHost() string {
	return r.DstHost
}

//GetDstPort ...
func (r *HttpResponse) GetDstPort() string {
	return r.DstPort
}

//GetPath ...
func (r *HttpResponse) GetPath() string {
	return ""
}

//GetBody ...
func (r *HttpResponse) GetBody() interface{} {
	return r.Body
}

//GetHeaders ...
func (r *HttpResponse) GetHeaders() map[string]string {
	return make(map[string]string)
}
