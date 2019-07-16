package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/saveav/pelacore/bindings"
	"github.com/saveav/pelacore/renderings"
)

func errPostQuery(c echo.Context, resp renderings.DatabaseQuery, message string) error {
	resp.Execution.Success = false
	resp.Execution.Message = message
	return c.JSON(http.StatusBadRequest, resp)
}

// PostQuery is used to execute a SQL query for a database
func PostQuery(c echo.Context) error {
	req := new(bindings.DatabaseQuery)
	resp := renderings.DatabaseQuery{}

	resp.Execution.Success = true

	if err := c.Bind(req); err != nil {
		return errPostQuery(c, resp, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return errPostQuery(c, resp, err.Error())
	}

	conn := req.ToDatabaseConnection()

	conStr, err := conn.String()
	if err != nil {
		return errPostQuery(c, resp, err.Error())
	}

	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		return errPostQuery(c, resp, err.Error())
	}
	defer db.Close()

	rows, err := db.Queryx(req.Query)
	if err != nil {
		return errPostQuery(c, resp, err.Error())
	}

	for rows.Next() {
		row := make(renderings.Row)
		err = rows.MapScan(row)
		if err != nil {
			return errPostQuery(c, resp, err.Error())
		}
		resp.Rows = append(resp.Rows, row)

		if len(resp.Columns) == 0 {
			columns, err := scanColumns(rows)
			if err != nil {
				return errPostQuery(c, resp, err.Error())
			}
			resp.Columns = columns
		}
	}

	c.JSON(http.StatusOK, resp)
	return nil
}

func scanColumns(rows *sqlx.Rows) ([]renderings.DatabaseColumns, error) {
	cols, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	var columns []renderings.DatabaseColumns

	for i := 0; i < len(cols); i++ {
		col := cols[i]
		column := renderings.DatabaseColumns{}
		column.Name = col.Name()
		column.Type = col.DatabaseTypeName()
		if nullable, ok := col.Nullable(); ok {
			// FIXME Postgres cannot detect Nullable. Consider submitting a PR?
			column.Nullable = &nullable
		}
		// VARCHAR length can be minus for Postgres
		if length, ok := col.Length(); ok && length > 0 {
			fmt.Println("LENGTH", length)
			val := strconv.FormatInt(length, 10)
			column.Length = &val
		}
		if precision, scale, ok := col.DecimalSize(); ok {
			val := fmt.Sprintf("%d, %d", precision, scale)
			column.Length = &val
		}
		columns = append(columns, column)
	}

	return columns, nil
}
