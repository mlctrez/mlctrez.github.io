package compo

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

var _ app.AppUpdater = (*Version)(nil)
var _ app.Mounter = (*Version)(nil)

type Version struct {
	app.Compo
	updateAvailable bool
}

func (v *Version) OnMount(ctx app.Context) {
	ctx.Async(ctx.TryUpdateApp)
}

func (v *Version) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		v.updateAvailable = true
	}
}

func (v *Version) Render() app.UI {
	if v.updateAvailable {
		return app.Div().Class("version").Text("Update Available").OnClick(func(ctx app.Context, e app.Event) {
			ctx.Reload()
		})
	}
	return app.Div().Class("version").Text(app.Getenv("GOAPP_VERSION")).OnClick(func(ctx app.Context, e app.Event) {
		ctx.Dispatch(func(context app.Context) {
			context.TryUpdateApp()
		})
	})
}
