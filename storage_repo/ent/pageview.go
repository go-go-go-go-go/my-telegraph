// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"telegraph/storage_repo/ent/pageview"

	"entgo.io/ent/dialect/sql"
)

// PageView is the model entity for the PageView schema.
type PageView struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PageID holds the value of the "page_id" field.
	PageID int `json:"page_id,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Year holds the value of the "year" field.
	Year int `json:"year,omitempty"`
	// Month holds the value of the "month" field.
	Month int `json:"month,omitempty"`
	// Day holds the value of the "day" field.
	Day int `json:"day,omitempty"`
	// Hour holds the value of the "hour" field.
	Hour int `json:"hour,omitempty"`
	// Views holds the value of the "views" field.
	Views int `json:"views,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PageView) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pageview.FieldID, pageview.FieldPageID, pageview.FieldYear, pageview.FieldMonth, pageview.FieldDay, pageview.FieldHour, pageview.FieldViews:
			values[i] = new(sql.NullInt64)
		case pageview.FieldPath:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PageView", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PageView fields.
func (pv *PageView) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pageview.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pv.ID = int(value.Int64)
		case pageview.FieldPageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field page_id", values[i])
			} else if value.Valid {
				pv.PageID = int(value.Int64)
			}
		case pageview.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				pv.Path = value.String
			}
		case pageview.FieldYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field year", values[i])
			} else if value.Valid {
				pv.Year = int(value.Int64)
			}
		case pageview.FieldMonth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field month", values[i])
			} else if value.Valid {
				pv.Month = int(value.Int64)
			}
		case pageview.FieldDay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field day", values[i])
			} else if value.Valid {
				pv.Day = int(value.Int64)
			}
		case pageview.FieldHour:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hour", values[i])
			} else if value.Valid {
				pv.Hour = int(value.Int64)
			}
		case pageview.FieldViews:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field views", values[i])
			} else if value.Valid {
				pv.Views = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PageView.
// Note that you need to call PageView.Unwrap() before calling this method if this PageView
// was returned from a transaction, and the transaction was committed or rolled back.
func (pv *PageView) Update() *PageViewUpdateOne {
	return (&PageViewClient{config: pv.config}).UpdateOne(pv)
}

// Unwrap unwraps the PageView entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pv *PageView) Unwrap() *PageView {
	_tx, ok := pv.config.driver.(*txDriver)
	if !ok {
		panic("ent: PageView is not a transactional entity")
	}
	pv.config.driver = _tx.drv
	return pv
}

// String implements the fmt.Stringer.
func (pv *PageView) String() string {
	var builder strings.Builder
	builder.WriteString("PageView(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pv.ID))
	builder.WriteString("page_id=")
	builder.WriteString(fmt.Sprintf("%v", pv.PageID))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(pv.Path)
	builder.WriteString(", ")
	builder.WriteString("year=")
	builder.WriteString(fmt.Sprintf("%v", pv.Year))
	builder.WriteString(", ")
	builder.WriteString("month=")
	builder.WriteString(fmt.Sprintf("%v", pv.Month))
	builder.WriteString(", ")
	builder.WriteString("day=")
	builder.WriteString(fmt.Sprintf("%v", pv.Day))
	builder.WriteString(", ")
	builder.WriteString("hour=")
	builder.WriteString(fmt.Sprintf("%v", pv.Hour))
	builder.WriteString(", ")
	builder.WriteString("views=")
	builder.WriteString(fmt.Sprintf("%v", pv.Views))
	builder.WriteByte(')')
	return builder.String()
}

// PageViews is a parsable slice of PageView.
type PageViews []*PageView

func (pv PageViews) config(cfg config) {
	for _i := range pv {
		pv[_i].config = cfg
	}
}
