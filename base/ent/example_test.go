// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleMeta() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the meta's edges.

	// create meta vertex with its edges.
	m := client.Meta.
		Create().
		SetBase("string").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetDeletedAt(time.Now()).
		SaveX(ctx)
	log.Println("meta created:", m)

	// query edges.

	// Output:
}