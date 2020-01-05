// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// MetaColumns holds the columns for the "meta" table.
	MetaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "base", Type: field.TypeString},
		{Name: "schema", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// MetaTable holds the schema information for the "meta" table.
	MetaTable = &schema.Table{
		Name:        "meta",
		Columns:     MetaColumns,
		PrimaryKey:  []*schema.Column{MetaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MetaTable,
	}
)

func init() {
}
