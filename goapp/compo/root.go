package compo

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Root struct {
	app.Compo
}

func (r *Root) Render() app.UI {
	return app.Div().Body(&Version{})
}
