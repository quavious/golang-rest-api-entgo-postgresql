// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/quavious/golang-docker-forum/ent/category"
	"github.com/quavious/golang-docker-forum/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// NetPrice holds the value of the "net_price" field.
	NetPrice int `json:"net_price,omitempty"`
	// SalePrice holds the value of the "sale_price" field.
	SalePrice int `json:"sale_price,omitempty"`
	// DiscountRate holds the value of the "discount_rate" field.
	DiscountRate int `json:"discount_rate,omitempty"`
	// LinkURL holds the value of the "link_url" field.
	LinkURL string `json:"link_url,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// Company holds the value of the "company" field.
	Company string `json:"company,omitempty"`
	// ExpiredAt holds the value of the "expired_at" field.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges             ProductEdges `json:"edges"`
	category_contains *int
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Belongs holds the value of the belongs edge.
	Belongs *Category
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BelongsOrErr returns the Belongs value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) BelongsOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Belongs == nil {
			// The edge belongs was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Belongs, nil
	}
	return nil, &NotLoadedError{edge: "belongs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldNetPrice, product.FieldSalePrice, product.FieldDiscountRate:
			values[i] = &sql.NullInt64{}
		case product.FieldID, product.FieldName, product.FieldLinkURL, product.FieldImageURL, product.FieldCompany:
			values[i] = &sql.NullString{}
		case product.FieldExpiredAt:
			values[i] = &sql.NullTime{}
		case product.ForeignKeys[0]: // category_contains
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Product", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldNetPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field net_price", values[i])
			} else if value.Valid {
				pr.NetPrice = int(value.Int64)
			}
		case product.FieldSalePrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sale_price", values[i])
			} else if value.Valid {
				pr.SalePrice = int(value.Int64)
			}
		case product.FieldDiscountRate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field discount_rate", values[i])
			} else if value.Valid {
				pr.DiscountRate = int(value.Int64)
			}
		case product.FieldLinkURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link_url", values[i])
			} else if value.Valid {
				pr.LinkURL = value.String
			}
		case product.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				pr.ImageURL = value.String
			}
		case product.FieldCompany:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company", values[i])
			} else if value.Valid {
				pr.Company = value.String
			}
		case product.FieldExpiredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expired_at", values[i])
			} else if value.Valid {
				pr.ExpiredAt = value.Time
			}
		case product.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_contains", value)
			} else if value.Valid {
				pr.category_contains = new(int)
				*pr.category_contains = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBelongs queries the "belongs" edge of the Product entity.
func (pr *Product) QueryBelongs() *CategoryQuery {
	return (&ProductClient{config: pr.config}).QueryBelongs(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", net_price=")
	builder.WriteString(fmt.Sprintf("%v", pr.NetPrice))
	builder.WriteString(", sale_price=")
	builder.WriteString(fmt.Sprintf("%v", pr.SalePrice))
	builder.WriteString(", discount_rate=")
	builder.WriteString(fmt.Sprintf("%v", pr.DiscountRate))
	builder.WriteString(", link_url=")
	builder.WriteString(pr.LinkURL)
	builder.WriteString(", image_url=")
	builder.WriteString(pr.ImageURL)
	builder.WriteString(", company=")
	builder.WriteString(pr.Company)
	builder.WriteString(", expired_at=")
	builder.WriteString(pr.ExpiredAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
