// Code generated by ent, DO NOT EDIT.

package history

import (
	"time"
)

const (
	// Label holds the string label denoting the history type in the database.
	Label = "history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatorID holds the string denoting the creator_id field in the database.
	FieldCreatorID = "creator_id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldLastUpdate holds the string denoting the last_update field in the database.
	FieldLastUpdate = "last_update"
	// FieldContentFlag holds the string denoting the content_flag field in the database.
	FieldContentFlag = "content_flag"
	// FieldSentFlag holds the string denoting the sent_flag field in the database.
	FieldSentFlag = "sent_flag"
	// FieldMentionedCount holds the string denoting the mentioned_count field in the database.
	FieldMentionedCount = "mentioned_count"
	// FieldTakeEffectTime holds the string denoting the take_effect_time field in the database.
	FieldTakeEffectTime = "take_effect_time"
	// FieldSendingContent holds the string denoting the sending_content field in the database.
	FieldSendingContent = "sending_content"
	// Table holds the table name of the history in the database.
	Table = "histories"
)

// Columns holds all SQL columns for history fields.
var Columns = []string{
	FieldID,
	FieldCreatorID,
	FieldCreateAt,
	FieldLastUpdate,
	FieldContentFlag,
	FieldSentFlag,
	FieldMentionedCount,
	FieldTakeEffectTime,
	FieldSendingContent,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultLastUpdate holds the default value on creation for the "last_update" field.
	DefaultLastUpdate func() time.Time
	// UpdateDefaultLastUpdate holds the default value on update for the "last_update" field.
	UpdateDefaultLastUpdate func() time.Time
	// DefaultMentionedCount holds the default value on creation for the "mentioned_count" field.
	DefaultMentionedCount int
	// DefaultTakeEffectTime holds the default value on creation for the "take_effect_time" field.
	DefaultTakeEffectTime time.Time
)
