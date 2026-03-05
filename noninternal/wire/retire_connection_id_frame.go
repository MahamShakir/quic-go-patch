package wire

import (
	"github.com/MahamShakir/quic-go-patch/noninternal/protocol"
	"github.com/MahamShakir/quic-go-patch/quicvarint"
)

// A RetireConnectionIDFrame is a RETIRE_CONNECTION_ID frame
type RetireConnectionIDFrame struct {
	SequenceNumber uint64
}

func parseRetireConnectionIDFrame(b []byte, _ protocol.Version) (*RetireConnectionIDFrame, int, error) {
	seq, l, err := quicvarint.Parse(b)
	if err != nil {
		return nil, 0, replaceUnexpectedEOF(err)
	}
	return &RetireConnectionIDFrame{SequenceNumber: seq}, l, nil
}

func (f *RetireConnectionIDFrame) Append(b []byte, _ protocol.Version) ([]byte, error) {
	b = append(b, byte(FrameTypeRetireConnectionID))
	b = quicvarint.Append(b, f.SequenceNumber)
	return b, nil
}

// Length of a written frame
func (f *RetireConnectionIDFrame) Length(protocol.Version) protocol.ByteCount {
	return 1 + protocol.ByteCount(quicvarint.Len(f.SequenceNumber))
}
