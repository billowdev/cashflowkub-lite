package main

import (
	"log"
	"os"

	"github.com/billowdev/cashflowkub-lite/collections"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// Move the collection initialization to OnBeforeServe
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		collections.NewUserCollection(app)
		collections.NewCategoryCollection(app)
		collections.NewCashflowinCollection(app)
		collections.NewCashflowoutCollection(app)
		collections.NewPocketCollection(app)
		collections.NewTransferCollection(app)
		collections.NewTransactionCollection(app)
	

		if err := collections.DropFieldFromCollection(app, collections.CategoryCollection, "is_custom"); err != nil {
			log.Printf("Failed to drop field: %v", err)
		}

		// Serve static files from the provided public directory (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
