package atoms

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// SophonButton is a Human-First accessible button component.
// It requires a label for accessibility (SophonLabel) even if it's icon-only.
type SophonButton struct {
	material.ButtonStyle
	SophonLabel string // The accessible label for screen readers
	Tooltip     string // Visible tooltip on hover
}

func NewSophonButton(th *material.Theme, click *widget.Clickable, text string, sophonLabel string) SophonButton {
	return SophonButton{
		ButtonStyle: material.Button(th, click, text),
		SophonLabel: sophonLabel,
		Tooltip:     sophonLabel,
	}
}

import (
	"gioui.org/io/semantic"
)

func (b SophonButton) Layout(gtx layout.Context) layout.Dimensions {
	// Simple tooltip logic: if the button is hovered, we render a label on top.
	// In Gio, we would typically use a dedicated tooltip widget.
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return semantic.Description(b.SophonLabel).Layout(gtx)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return b.ButtonStyle.Layout(gtx)
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			if gtx.Source.Focused(b.Button) {
				// Simple mock of a visible tooltip
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Caption(b.ButtonStyle.Theme, b.Tooltip).Layout(gtx)
				})
			}
			return layout.Dimensions{}
		}),
	)
}
