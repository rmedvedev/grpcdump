package filter

import "golang.org/x/net/bpf"

//PacketFilter represents wrapper for bpf filter
type PacketFilter struct {
	bpfFilter []bpf.RawInstruction
}

//New creates new filter
func New() *PacketFilter {
	return &PacketFilter{}
}

//SetPort creates new bpf filter for src and dst port
func (filter *PacketFilter) SetPort(port uint32) {

	filter.bpfFilter = []bpf.RawInstruction{
		{Op: 40, Jt: 0, Jf: 0, K: 12},
		{Op: 21, Jt: 0, Jf: 6, K: 34525},
		{Op: 48, Jt: 0, Jf: 0, K: 20},
		{Op: 21, Jt: 0, Jf: 15, K: 6},
		{Op: 40, Jt: 0, Jf: 0, K: 54},
		{Op: 21, Jt: 12, Jf: 0, K: port},
		{Op: 40, Jt: 0, Jf: 0, K: 56},
		{Op: 21, Jt: 10, Jf: 11, K: port},
		{Op: 21, Jt: 0, Jf: 10, K: 2048},
		{Op: 48, Jt: 0, Jf: 0, K: 23},
		{Op: 21, Jt: 0, Jf: 8, K: 6},
		{Op: 40, Jt: 0, Jf: 0, K: 20},
		{Op: 69, Jt: 6, Jf: 0, K: 8191},
		{Op: 177, Jt: 0, Jf: 0, K: 14},
		{Op: 72, Jt: 0, Jf: 0, K: 14},
		{Op: 21, Jt: 2, Jf: 0, K: port},
		{Op: 72, Jt: 0, Jf: 0, K: 16},
		{Op: 21, Jt: 0, Jf: 1, K: port},
		{Op: 6, Jt: 0, Jf: 0, K: 1000},
		{Op: 6, Jt: 0, Jf: 0, K: 0},
	}
}

//GetBpfFilter return bpf filter
func (filter *PacketFilter) GetBpfFilter() []bpf.RawInstruction {
	return filter.bpfFilter
}
