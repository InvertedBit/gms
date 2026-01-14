package adminviews

import (
	"github.com/invertedbit/gms/html/components"
	"maragu.dev/gomponents"
)

func UserListPage(userTableData *components.TableData) gomponents.Node {
	return components.DataTable(userTableData)
}
