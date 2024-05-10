package reloadable

// This code is relodable. Change some stuff in Layout and save and got-reload
// should do the rest!

import (
	"math"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

var (
	b1, b2, b3 widget.Clickable
	fl         widget.Float

	// These need to be pointers for now. I believe this is a Yaegi bug.
	offset      = new(float64)
	button1Dims = new(D)
)

func Layout(gtx C, th *material.Theme) D {
	// Try changing this inset, or redefine it as a literal with different
	//  values for each side:

	// sharedInset := layout.Inset{
	// 	Left:   unit.Dp(4),
	// 	Right:  unit.Dp(8),
	// 	Top:    unit.Dp(1),
	// 	Bottom: unit.Dp(13),
	// }
	sharedInset := layout.UniformInset(unit.Dp(8))
	return layout.Center.Layout(gtx, func(gtx C) D {
		/*
		   Try uncommenting these and clicking the buttons.
		*/
		// if b1.Clicked(gtx) {
		// 	log.Printf("b1 clicked")
		// }
		// if b2.Clicked(gtx) {
		// 	log.Printf("b2 clicked")
		// }
		// if b3.Clicked(gtx) {
		// 	log.Printf("b3 clicked")
		// }

		return layout.Flex{
			/*
				Try changing these properties!
			*/
			Axis:      layout.Vertical,
			Spacing:   layout.SpaceAround,
			Alignment: layout.Baseline,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				gtx.Constraints.Min.X = button1Dims.Size.X
				gtx.Constraints.Min.Y = button1Dims.Size.Y
				fl.Value = float32(math.Abs(math.Sin(*offset)))
				// log.Printf("fl.Value: %v", fl.Value)
				*offset += 0.01
				return material.Slider(th, &fl).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				*button1Dims = sharedInset.Layout(gtx, func(gtx C) D {
					btn := material.Button(th, &b1, "Click me 1")
					/*
						Play with the inset dimensions!
					*/
					btn.Inset = layout.UniformInset(unit.Dp(30))
					return btn.Layout(gtx)
				})
				return *button1Dims
			}),
			/*
				Play with the first parameter here!
				Alternatively, make this a layout.Rigid
			*/
			layout.Flexed(.75, func(gtx C) D {
				return sharedInset.Layout(gtx, func(gtx C) D {
					btn := material.Button(th, &b2, "Click me 2")
					return btn.Layout(gtx)
				})
			}),
			layout.Flexed(.25, func(gtx C) D {
				return sharedInset.Layout(gtx, func(gtx C) D {
					btn := material.Button(th, &b3, "Click me 3")
					return btn.Layout(gtx)
				})
			}),
		)
	})
}
