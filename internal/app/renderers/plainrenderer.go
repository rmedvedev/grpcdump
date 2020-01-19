package renderers

import (
	"fmt"

	"github.com/rmedvedev/grpcdump/internal/app/models"
)

//PlainRenderer ...
type PlainRenderer struct{}

//Render renders model
func (PlainRenderer) Render(model models.RenderModel) string {
	return fmt.Sprintf(
		"%s:%s -> %s:%s %s %+v %s",
		model.GetSrcHost(),
		model.GetSrcPort(),
		model.GetDstHost(),
		model.GetDstPort(),
		model.GetPath(),
		model.GetBody(),
		model.GetHeaders(),
	)
}
