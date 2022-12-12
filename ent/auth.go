// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mofu/ent/auth"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Auth is the model entity for the Auth schema.
type Auth struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Auth) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case auth.FieldID:
			values[i] = new(sql.NullInt64)
		case auth.FieldToken, auth.FieldUser:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Auth", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Auth fields.
func (a *Auth) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case auth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case auth.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				a.Token = value.String
			}
		case auth.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				a.User = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Auth.
// Note that you need to call Auth.Unwrap() before calling this method if this Auth
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Auth) Update() *AuthUpdateOne {
	return (&AuthClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Auth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Auth) Unwrap() *Auth {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Auth is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Auth) String() string {
	var builder strings.Builder
	builder.WriteString("Auth(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("token=")
	builder.WriteString(a.Token)
	builder.WriteString(", ")
	builder.WriteString("user=")
	builder.WriteString(a.User)
	builder.WriteByte(')')
	return builder.String()
}

// Auths is a parsable slice of Auth.
type Auths []*Auth

func (a Auths) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
