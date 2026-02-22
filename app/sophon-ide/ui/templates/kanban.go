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
	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return t.renderColumn(th, gtx, "TODO", []string{"Refactor core", "Add tests"})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return t.renderColumn(th, gtx, "DOING", []string{"Implement Kanban"})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return t.renderColumn(th, gtx, "DONE", []string{"Rebrand to Sophon"})
		}),
	)
}

func (t KanbanTemplate) renderColumn(th *material.Theme, gtx layout.Context, title string, tasks []string) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(th, title).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			list := layout.List{Axis: layout.Vertical}
			return list.Layout(gtx, len(tasks), func(gtx layout.Context, i int) layout.Dimensions {
				return material.Body1(th, tasks[i]).Layout(gtx)
			})
		}),
	)
}
