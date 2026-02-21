package templates

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type KanbanTemplate struct {
	Columns []KanbanColumn
}

type KanbanColumn struct {
	Title string
	Tasks []string
}

func (t KanbanTemplate) Layout(th *material.Theme, gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			// Render ToDo column
			return layout.Dimensions{}
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			// Render In Progress column
			return layout.Dimensions{}
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			// Render Done column
			return layout.Dimensions{}
		}),
	)
}
