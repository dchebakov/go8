// Code generated by entc, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gmhafiz/go8/ent/gen/book"
)

// Book is the model entity for the Book schema.
type Book struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// PublishedDate holds the value of the "published_date" field.
	PublishedDate time.Time `json:"published_date,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"-"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"-"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookQuery when eager-loading is set.
	Edges BookEdges `json:"edges"`
}

// BookEdges holds the relations/edges for other nodes in the graph.
type BookEdges struct {
	// Authors holds the value of the authors edge.
	Authors []*Author `json:"authors,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AuthorsOrErr returns the Authors value or an error if the edge
// was not loaded in eager-loading.
func (e BookEdges) AuthorsOrErr() ([]*Author, error) {
	if e.loadedTypes[0] {
		return e.Authors, nil
	}
	return nil, &NotLoadedError{edge: "authors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Book) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			values[i] = new(sql.NullInt64)
		case book.FieldTitle, book.FieldDescription:
			values[i] = new(sql.NullString)
		case book.FieldPublishedDate, book.FieldCreatedAt, book.FieldUpdatedAt, book.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Book", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Book fields.
func (b *Book) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = uint(value.Int64)
		case book.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case book.FieldPublishedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field published_date", values[i])
			} else if value.Valid {
				b.PublishedDate = value.Time
			}
		case book.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		case book.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case book.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case book.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				b.DeletedAt = new(time.Time)
				*b.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryAuthors queries the "authors" edge of the Book entity.
func (b *Book) QueryAuthors() *AuthorQuery {
	return (&BookClient{config: b.config}).QueryAuthors(b)
}

// Update returns a builder for updating this Book.
// Note that you need to call Book.Unwrap() before calling this method if this Book
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Book) Update() *BookUpdateOne {
	return (&BookClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Book entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Book) Unwrap() *Book {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("gen: Book is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Book) String() string {
	var builder strings.Builder
	builder.WriteString("Book(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", title=")
	builder.WriteString(b.Title)
	builder.WriteString(", published_date=")
	builder.WriteString(b.PublishedDate.Format(time.ANSIC))
	builder.WriteString(", description=<sensitive>")
	builder.WriteString(", created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	if v := b.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Books is a parsable slice of Book.
type Books []*Book

func (b Books) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
