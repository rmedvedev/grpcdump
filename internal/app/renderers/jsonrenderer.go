package renderers

import (
	"encoding/json"
	"fmt"

	"github.com/rmedvedev/grpcdump/internal/app/models"
	"github.com/sirupsen/logrus"
)

//JSONRenderer ...
type JSONRenderer struct{}

type jsonView struct {
	Src     string            `json:"src"`
	Dst     string            `json:"dst"`
	Path    string            `json:"path"`
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
}

//Render renders model
func (JSONRenderer) Render(model models.RenderModel) string {
	bytes, err := json.Marshal(jsonView{
		Src:     model.GetSrcHost() + ":" + model.GetSrcPort(),
		Dst:     model.GetDstHost() + ":" + model.GetDstPort(),
		Path:    model.GetPath(),
		Body:    fmt.Sprintf("%v", model.GetBody()),
		Headers: model.GetHeaders(),
	})

	if err != nil {
		logrus.Errorf("Error to marshal model: %s", err.Error())
	}

	return string(bytes)
}
