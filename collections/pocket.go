package collections

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var PocketCollection = "pocket"

func NewPocketCollection(app *pocketbase.PocketBase) {
	collection := &models.Collection{
		Name: PocketCollection,
		Schema: schema.NewSchema(
			&schema.SchemaField{Name: "name", Type: schema.FieldTypeText, Required: true},
			&schema.SchemaField{Name: "balance", Type: schema.FieldTypeNumber, Required: true},
			&schema.SchemaField{Name: "user_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "user"}, Required: true},
			&schema.SchemaField{Name: "created_at", Type: schema.FieldTypeDate, Required: false},
			&schema.SchemaField{Name: "updated_at", Type: schema.FieldTypeDate, Required: false},
		),
	}

	saveCollection(app, collection)

}
