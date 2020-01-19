package httpparser

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/rmedvedev/grpcdump/internal/app/framereader"
	"github.com/rmedvedev/grpcdump/internal/app/models"
	"github.com/sirupsen/logrus"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

//HttpParser ...
type HttpParser struct {
	assembler *tcpassembly.Assembler
}

// httpStreamFactory implements tcpassembly.StreamFactory
type httpStreamFactory struct {
	modelsCh *chan models.RenderModel
	paths    *sync.Map
}

// httpStream will handle the actual decoding of http requests.
type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
	modelsCh       *chan models.RenderModel
	Paths          *sync.Map
}

//New creates new instance of HttpParser
func New(modelsCh *chan models.RenderModel) *HttpParser {
	streamFactory := &httpStreamFactory{modelsCh, &sync.Map{}}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	return &HttpParser{assembler}
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
		modelsCh:  h.modelsCh,
		Paths:     h.paths,
	}
	go hstream.run()

	return &hstream.r
}

//Parse ...
func (parser *HttpParser) Parse(packet gopacket.Packet) error {
	if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
		return errors.New("Unusable packet")
	}

	tcp := packet.TransportLayer().(*layers.TCP)
	parser.assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)

	return nil
}

func tryReadHttpRequest(packet *models.Packet, prefix string, buf *bufio.Reader) (models.RenderModel, error) {
	prefix = strings.ToUpper(prefix)
	if strings.HasPrefix(prefix, "GET") ||
		strings.HasPrefix(prefix, "POST") ||
		strings.HasPrefix(prefix, "PUT") ||
		strings.HasPrefix(prefix, "DELETE") ||
		strings.HasPrefix(prefix, "OPTIONS") ||
		strings.HasPrefix(prefix, "HEAD") {

		httpRequest := models.NewHttpRequest(packet)

		r, err := http.ReadRequest(buf)
		if err != nil {
			return nil, err
		}

		httpRequest.URL = r.URL.String()
		httpRequest.Method = r.Method

		_, err = buf.Discard(int(r.ContentLength))
		if err != nil {
			return nil, err
		}

		err = r.Body.Close()
		if err != nil {
			return nil, err
		}

		return httpRequest, nil
	}

	return nil, nil
}

func tryReadHttpResponse(packet *models.Packet, prefix string, buf *bufio.Reader) (models.RenderModel, error) {
	prefix = strings.ToUpper(prefix)
	if strings.HasPrefix(prefix, "HTTP") {

		httpResponse := models.NewHttpResponse(packet)

		resp, err := http.ReadResponse(buf, nil)
		if err != nil {
			return nil, err
		}

		//httpResponse.Body = resp.Body

		_, err = buf.Discard(int(resp.ContentLength))
		if err != nil {
			return nil, err
		}

		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		return httpResponse, nil
	}

	return nil, nil
}

func tryReadHttp2(packet *models.Packet, prefix string, buf *bufio.Reader, frameReader *framereader.FrameReader) (models.RenderModel, error) {

	if strings.HasPrefix(prefix, "PRI") {
		_, err := buf.Discard(len(http2.ClientPreface))
		if err != nil {
			return nil, err
		}
	}

	http2model, err := frameReader.Read(packet)
	if err != nil {
		return nil, err
	}

	return http2model, nil
}

func (h *httpStream) run() {
	// 1 request, 2 response, 0 unkonwn

	buf := bufio.NewReader(&h.r)
	framer := http2.NewFramer(ioutil.Discard, buf)
	framer.ReadMetaHeaders = hpack.NewDecoder(4096, nil)
	frameReader := framereader.New(framer, h.Paths)

	packet := models.NewPacket(h.net, h.transport)

	defer func() {
		h.Paths.Delete(packet.GetConnectionKey())
		h.Paths.Delete(packet.GetRevConnectionKey())
	}()

	for {
		peekBuf, err := buf.Peek(9)
		if err == io.EOF {
			return
		} else if err != nil {
			logrus.Error("Error reading frame", h.net, h.transport, ":", err)
			continue
		}

		prefix := string(peekBuf)

		//if http request then continue
		if model, err := tryReadHttpRequest(packet, prefix, buf); model != nil || err != nil {
			if err != nil {
				logrus.Errorf("Error in try to read http request: %s", err.Error())
			} else {
				*h.modelsCh <- model
			}

			continue
		}

		//if http response then continue
		if model, err := tryReadHttpResponse(packet, prefix, buf); model != nil || err != nil {
			if err != nil {
				logrus.Errorf("Error in try to read http response: %s", err.Error())
			} else {
				*h.modelsCh <- model
			}
			continue
		}

		if model, err := tryReadHttp2(packet, prefix, buf, frameReader); model != nil || err != nil {
			if err != nil {
				logrus.Errorf("Error in try to read http2: %s", err.Error())
			} else {
				*h.modelsCh <- model
			}
			continue
		}
	}
}
