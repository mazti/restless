// Code generated by entc, DO NOT EDIT.

package metacolumn

import (
	"time"

	"github.com/mazti/restless/base/ent/schema"
)

const (
	// Label holds the string label denoting the metacolumn type in the database.
	Label = "meta_column"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type vertex property in the database.
	FieldType = "type"
	// FieldDefault holds the string denoting the default vertex property in the database.
	FieldDefault = "default"
	// FieldTypeOption holds the string denoting the type_option vertex property in the database.
	FieldTypeOption = "type_option"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at vertex property in the database.
	FieldDeletedAt = "deleted_at"

	// Table holds the table name of the metacolumn in the database.
	Table = "meta_columns"
	// TableTable is the table the holds the table relation/edge.
	TableTable = "meta_columns"
	// TableInverseTable is the table name for the MetaTable entity.
	// It exists in this package in order to avoid circular dependency with the "metatable" package.
	TableInverseTable = "meta_tables"
	// TableColumn is the table column denoting the table relation/edge.
	TableColumn = "table_id"
)

// Columns holds all SQL columns are metacolumn fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldDefault,
	FieldTypeOption,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

var (
	fields = schema.MetaColumn{}.Fields()

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = fields[4].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)

	// descUpdatedAt is the schema descriptor for updated_at field.
	descUpdatedAt = fields[5].Descriptor()
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt = descUpdatedAt.Default.(func() time.Time)
)
