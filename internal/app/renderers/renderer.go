package renderers

import (
	"github.com/rmedvedev/grpcdump/internal/pkg/config"
	"github.com/rmedvedev/grpcdump/internal/app/models"
)

//Renderer ...
type Renderer interface {
	Render(model models.RenderModel) string
}

//GetApplicationRenderer ...
func GetApplicationRenderer() Renderer {
	cfg := config.GetConfig()
	if cfg.ColorOutput {
		return &PrettyRenderer{}
	} else if cfg.JSONOutput {
		return &JSONRenderer{}
	}

	return &PlainRenderer{}
}
