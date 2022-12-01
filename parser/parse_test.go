package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	p := Parse{}
	p.AppendCodeGragment(`package pages

	import (
		"git.philo.top/zettavisor/rulesAndAlert/wasm/comps"
		"github.com/maxence-charriere/go-app/v9/pkg/app"
	)
	
	type Homepage struct {
		app.Compo
	
		userID            string
		userNickName      string
		SideBarExpandedLG string
	}
	
	func NewHomepage() *Homepage {
		return &Homepage{}
	}
	
	func (h *Homepage) OnMount(ctx app.Context) {
		ctx.ObserveState(comps.StorageSideBarExpandedLG).Value(&h.SideBarExpandedLG)
		comps.LoginCheckAndRedirectSignin(ctx)
		h.userID = comps.GetCurrentUserID(ctx)
		h.userNickName = comps.GetCurrentNickName(ctx)
	}
	
	func (h *Homepage) Render() app.UI {
		return comps.FrameWithSidebar(comps.NewSidebar(), comps.NewHeader(&comps.HeaderState{
			UserID:       h.userID,
			UserNickName: h.userNickName,
		}), h.SideBarExpandedLG, app.Div().Class("px-4 sm:px-6 lg:px-8 py-8 w-full max-w-9xl mx-auto").Body(
			app.Text("welcome to zettavisor!"),
		))
	}
	
	func (h *Homepage) OnAppUpdate(ctx app.Context) {
		comps.Upgrade(ctx)
	}
	`)
	a, b := p.Generate()
	fmt.Println(a, b)
}
