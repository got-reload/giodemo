// SPDX-License-Identifier: Unlicense OR MIT

package main

// A simple Gio program. See https://gioui.org for more information.

// This code is not reloadable. Only functions that exit and are called again
// can be replaced. Moreover, we don't filter the "main" package. See the
// reloadable.Layout function.

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/got-reload/giodemo/reloadable"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	for {
		e := w.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// Ensure that we invalidate every frame so that new versions
			// of the reloadable.Layout function are used immediately.
			w.Invalidate()
			reloadable.Layout(gtx, th)

			e.Frame(gtx.Ops)
		}
	}
}
