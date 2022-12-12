// Code generated by ent, DO NOT EDIT.

package auth

const (
	// Label holds the string label denoting the auth type in the database.
	Label = "auth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// Table holds the table name of the auth in the database.
	Table = "auths"
)

// Columns holds all SQL columns for auth fields.
var Columns = []string{
	FieldID,
	FieldToken,
	FieldUser,
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
