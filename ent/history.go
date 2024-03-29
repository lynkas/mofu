// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mofu/ent/history"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// History is the model entity for the History schema.
type History struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatorID holds the value of the "creator_id" field.
	CreatorID string `json:"creator_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// LastUpdate holds the value of the "last_update" field.
	LastUpdate time.Time `json:"last_update,omitempty"`
	// ContentFlag holds the value of the "content_flag" field.
	ContentFlag int `json:"content_flag,omitempty"`
	// SentFlag holds the value of the "sent_flag" field.
	SentFlag int `json:"sent_flag,omitempty"`
	// MentionedCount holds the value of the "mentioned_count" field.
	MentionedCount int `json:"mentioned_count,omitempty"`
	// TakeEffectTime holds the value of the "take_effect_time" field.
	TakeEffectTime time.Time `json:"take_effect_time,omitempty"`
	// SendingContent holds the value of the "sending_content" field.
	SendingContent []byte `json:"sending_content,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case history.FieldSendingContent:
			values[i] = new([]byte)
		case history.FieldContentFlag, history.FieldSentFlag, history.FieldMentionedCount:
			values[i] = new(sql.NullInt64)
		case history.FieldID, history.FieldCreatorID:
			values[i] = new(sql.NullString)
		case history.FieldCreateAt, history.FieldLastUpdate, history.FieldTakeEffectTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type History", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				h.ID = value.String
			}
		case history.FieldCreatorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				h.CreatorID = value.String
			}
		case history.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				h.CreateAt = value.Time
			}
		case history.FieldLastUpdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_update", values[i])
			} else if value.Valid {
				h.LastUpdate = value.Time
			}
		case history.FieldContentFlag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field content_flag", values[i])
			} else if value.Valid {
				h.ContentFlag = int(value.Int64)
			}
		case history.FieldSentFlag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sent_flag", values[i])
			} else if value.Valid {
				h.SentFlag = int(value.Int64)
			}
		case history.FieldMentionedCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mentioned_count", values[i])
			} else if value.Valid {
				h.MentionedCount = int(value.Int64)
			}
		case history.FieldTakeEffectTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field take_effect_time", values[i])
			} else if value.Valid {
				h.TakeEffectTime = value.Time
			}
		case history.FieldSendingContent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field sending_content", values[i])
			} else if value != nil {
				h.SendingContent = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this History.
// Note that you need to call History.Unwrap() before calling this method if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return (&HistoryClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the History entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("creator_id=")
	builder.WriteString(h.CreatorID)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(h.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_update=")
	builder.WriteString(h.LastUpdate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("content_flag=")
	builder.WriteString(fmt.Sprintf("%v", h.ContentFlag))
	builder.WriteString(", ")
	builder.WriteString("sent_flag=")
	builder.WriteString(fmt.Sprintf("%v", h.SentFlag))
	builder.WriteString(", ")
	builder.WriteString("mentioned_count=")
	builder.WriteString(fmt.Sprintf("%v", h.MentionedCount))
	builder.WriteString(", ")
	builder.WriteString("take_effect_time=")
	builder.WriteString(h.TakeEffectTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sending_content=")
	builder.WriteString(fmt.Sprintf("%v", h.SendingContent))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History

func (h Histories) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
