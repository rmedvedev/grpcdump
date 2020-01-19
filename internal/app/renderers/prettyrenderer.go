package renderers

import (
	"fmt"

	c "github.com/logrusorgru/aurora"
	"github.com/rmedvedev/grpcdump/internal/app/models"
)

//PrettyRenderer ...
type PrettyRenderer struct{}

//Render renders model
func (PrettyRenderer) Render(model models.RenderModel) string {
	return fmt.Sprintf(
		"%s:%s -> %s:%s %s %+v %s",
		c.Green(model.GetSrcHost()),
		c.Green(model.GetSrcPort()),
		c.Yellow(model.GetDstHost()),
		c.Yellow(model.GetDstPort()),
		c.BgBrightBlue(c.White(model.GetPath())),
		model.GetBody(),
		renderHeaders(model.GetHeaders()),
	)
}

func renderHeaders(headers map[string]string) string {
	result := "["

	for name, val := range headers {
		result += c.Cyan(name).String() + ":" + val + ", "
	}

	return result + "]"
}
