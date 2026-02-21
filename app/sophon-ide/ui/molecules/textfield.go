package molecules

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

// SophonTextField is a molecule that pairs a label with a text input,
// following the "Human-First" philosophy that every input must have a clear label.
type SophonTextField struct {
	material.EditorStyle
	Label string
}

func (t SophonTextField) Layout(th *material.Theme, gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, th.TextSize, t.Label).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.EditorStyle.Layout(gtx)
		}),
	)
}
