package framereader

import (
	"sync"

	"github.com/rmedvedev/grpcdump/internal/app/models"
)

//Streams ...
type Streams struct {
	collection map[string]map[uint32]*models.Stream
	mutex      sync.RWMutex
}

//NewStreams ...
func NewStreams() *Streams {
	return &Streams{
		collection: make(map[string]map[uint32]*models.Stream),
		mutex:      sync.RWMutex{},
	}
}

//Add ...
func (streams *Streams) Add(connectionKey string, stream *models.Stream) {
	streams.mutex.Lock()
	defer streams.mutex.Unlock()

	if _, ok := streams.collection[connectionKey]; !ok {
		streams.collection[connectionKey] = make(map[uint32]*models.Stream)
	}

	streams.collection[connectionKey][stream.ID] = stream
}

//Get ...
func (streams *Streams) Get(connectionKey string, streamID uint32) (*models.Stream, bool) {
	streams.mutex.RLock()
	defer streams.mutex.RUnlock()

	if _, ok := streams.collection[connectionKey][streamID]; !ok {
		return nil, false
	}

	return streams.collection[connectionKey][streamID], true
}
