package components

import (
	"fmt"

	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

type TableRow struct {
	Values map[string]string
}

type TableColumn struct {
	Name     string
	Label    string
	Sortable bool
}

type TableData struct {
	Title         string
	Columns       []TableColumn
	Rows          []TableRow
	Editable      bool
	Deletable     bool
	EditRoute     string
	DeleteRoute   string
	IDField       string // Field name to use as ID for edit/delete operations
	RefreshTarget string // HTMX target to refresh after operations
}

func DataTable(data *TableData) gomponents.Node {
	// Add Actions column if editable or deletable
	columns := data.Columns
	if data.Editable || data.Deletable {
		columns = append(columns, TableColumn{Name: "actions", Label: "Actions"})
	}

	return html.Div(
		html.ID("data-table-container"),
		html.Class("overflow-x-auto"),
		html.Table(
			html.Class("table table-zebra"),
			TableHeader(columns),
			TableBody(data.Rows, columns, data),
		),
	)
}

func TableHeader(columns []TableColumn) gomponents.Node {
	headerCells := []gomponents.Node{}
	for _, col := range columns {
		headerCells = append(headerCells, html.Th(gomponents.Text(col.Label)))
	}
	return html.THead(
		html.Tr(headerCells...),
	)
}

func TableBody(rows []TableRow, columns []TableColumn, data *TableData) gomponents.Node {
	bodyRows := []gomponents.Node{}
	for _, row := range rows {
		cells := []gomponents.Node{}
		for _, col := range columns {
			if col.Name == "actions" {
				// Render action buttons
				cells = append(cells, html.Td(
					html.Class("flex gap-2"),
					gomponents.If(data.Editable,
						html.Button(
							html.Class("btn btn-xs btn-primary"),
							htmx.Get(fmt.Sprintf("%s/%s", data.EditRoute, row.Values[data.IDField])),
							htmx.Target("#modal-container"),
							html.I(html.Class("ri-edit-line")),
						),
					),
					gomponents.If(data.Deletable,
						html.Button(
							html.Class("btn btn-xs btn-error"),
							htmx.Delete(fmt.Sprintf("%s/%s", data.DeleteRoute, row.Values[data.IDField])),
							htmx.Target(data.RefreshTarget),
							htmx.Confirm(fmt.Sprintf("Are you sure you want to delete this item?")),
							html.I(html.Class("ri-delete-bin-line")),
						),
					),
				))
			} else {
				cells = append(cells, html.Td(gomponents.Text(row.Values[col.Name])))
			}
		}
		bodyRows = append(bodyRows, html.Tr(cells...))
	}
	return html.TBody(bodyRows...)
}
