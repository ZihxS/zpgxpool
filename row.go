package zpgxpool

import (
	"fmt"
	"reflect"

	"github.com/jackc/pgx/v5/pgconn"
)

func scanRow(row []any, defs []pgconn.FieldDescription, dest ...any) error {
	for i, col := range row {
		if dest[i] == nil {
			// behave compatible with pgx
			continue
		}
		destVal := reflect.ValueOf(dest[i])
		if destVal.Kind() != reflect.Ptr {
			return fmt.Errorf("destination argument must be a pointer for column %s", defs[i].Name)
		}
		if col == nil {
			dest[i] = nil
			continue
		}
		val := reflect.ValueOf(col)

		destKind := destVal.Elem().Kind()
		if destKind == val.Kind() || destKind == reflect.Interface {
			if destElem := destVal.Elem(); destElem.CanSet() {
				destElem.Set(val)
			} else {
				return fmt.Errorf("cannot set destination value for column %s", string(defs[i].Name))
			}
		} else {
			return fmt.Errorf("destination kind '%v' not supported for value kind '%v' of column '%s'",
				destKind, val.Kind(), string(defs[i].Name))
		}
	}
	return nil
}

// Row is a mocked row that can be returned from QueryRow
type Row struct {
	defs    []pgconn.FieldDescription
	row     []any
	nextErr error
}

func (r *Row) Scan(dest ...any) error {
	if len(dest) != len(r.defs) {
		return fmt.Errorf("incorrect argument number %d for columns %d", len(dest), len(r.defs))
	}
	err := scanRow(r.row, r.defs, dest...)
	if err != nil {
		return err
	}
	return r.nextErr
}

func NewRow(columns []string, values ...any) *Row {
	var coldefs []pgconn.FieldDescription
	for _, column := range columns {
		coldefs = append(coldefs, pgconn.FieldDescription{Name: column})
	}
	return &Row{
		defs: coldefs,
		row:  values,
	}
}

// RowError allows to set an error
// which will be returned when the row is read
func (r *Row) RowError(row int, err error) *Row {
	r.nextErr = err
	return r
}
