package grpc

import (
	"bytes"
	"encoding/binary"

	"github.com/rmedvedev/grpcdump/internal/app/protoprovider"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

const (
	maxMsgSize   = 1024 * 1024 * 1024
	emptyMessage = ""
)

//Decode to decode proto messages
func Decode(path string, frame *http2.DataFrame, side int) (interface{}, error) {
	var dataBuf bytes.Buffer
	buf := frame.Data()
	if len(buf) == 0 {
		return nil, nil
	}

	streamID := frame.Header().StreamID
	length := int(binary.BigEndian.Uint32(buf[1:5]))

	compress := buf[0]

	if compress == 1 {
		logrus.Warningf("%d use compression, msg %q\n", streamID, buf[5:])
		return nil, nil
	}

	dataBuf.Write(buf[5:])

	if length > maxMsgSize || dataBuf.Len() > maxMsgSize {
		dataBuf.Truncate(0)
		return nil, nil
	}

	if length != dataBuf.Len() {
		return nil, nil
	}

	data := dataBuf.Bytes()

	defer func() {
		dataBuf.Truncate(0)
	}()

	if method, ok := protoprovider.GetProtoByPath(path); ok {
		switch side {
		case 1:
			msg := *(method.Request)
			if err := msg.Unmarshal(data); err != nil {
				logrus.Errorf("Error unmarshal request: %s", err.Error())
			}
			return msg.String(), nil
		case 2:
			msg := *(method.Response)
			if err := msg.Unmarshal(data); err != nil {
				logrus.Errorf("Error unmarshal response: %s", err.Error())
			}
			return msg.String(), nil
		}
	}

	return emptyMessage, nil
}
