package collections

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

func saveCollection(app *pocketbase.PocketBase, collection *models.Collection) {
	dao := app.Dao()
	if dao == nil {
		log.Fatalf("Dao is not initialized")
		return
	}

	existingCollection, err := dao.FindCollectionByNameOrId(collection.Name)
	if err != nil {
		log.Printf("Error finding collection by name or ID: %v", err)
	}

	if existingCollection == nil {
		if err := dao.SaveCollection(collection); err != nil {
			log.Fatalf("Failed to create '%s' collection: %v", collection.Name, err)
		} else {
			log.Printf("Collection '%s' initialized successfully!", collection.Name)
		}
	} else {
		log.Printf("Collection '%s' already exists.", collection.Name)
	}
}

func DropFieldFromCollection(app *pocketbase.PocketBase, collectionName, fieldName string) error {
	// Load the collection by name
	collection, err := app.Dao().FindCollectionByNameOrId(collectionName)
	if err != nil {
		return err
	}

	// Retrieve the existing fields and filter out the field to be removed
	newFields := []*schema.SchemaField{}
	for _, field := range collection.Schema.Fields() {
		if field.Name != fieldName {
			newFields = append(newFields, field)
		}
	}

	// Update the schema with the modified fields
	collection.Schema = schema.NewSchema(newFields...)

	// Save the updated collection
	if err := app.Dao().SaveCollection(collection); err != nil {
		log.Fatalf("Failed to update collection '%s': %v", collectionName, err)
		return err
	}

	log.Printf("Field '%s' dropped from collection '%s' successfully.", fieldName, collectionName)
	return nil
}
