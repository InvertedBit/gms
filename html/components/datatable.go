package components

import (
	"maragu.dev/gomponents"
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
	Title       string
	Columns     []TableColumn
	Rows        []TableRow
	Editable    bool
	Deletable   bool
	EditRoute   string
	DeleteRoute string
}

func DataTable(data *TableData) gomponents.Node {
	return html.Div(
		html.Class("overflow-x-auto"),
		html.Table(
			html.Class("table table-zebra"),
			TableHeader(data.Columns),
			TableBody(data.Rows, data.Columns),
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

func TableBody(rows []TableRow, columns []TableColumn) gomponents.Node {
	bodyRows := []gomponents.Node{}
	for _, row := range rows {
		cells := []gomponents.Node{}
		for _, col := range columns {
			cells = append(cells, html.Td(gomponents.Text(row.Values[col.Name])))
		}
		bodyRows = append(bodyRows, html.Tr(cells...))
	}
	return html.TBody(bodyRows...)
}
