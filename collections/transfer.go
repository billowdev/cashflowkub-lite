package collections

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var TransferCollection = "transfer"

func NewTransferCollection(app *pocketbase.PocketBase) {
	collection := &models.Collection{
		Name: TransferCollection,
		Schema: schema.NewSchema(
			&schema.SchemaField{Name: "amount", Type: schema.FieldTypeNumber, Required: true},
			&schema.SchemaField{Name: "to_pocket_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "pocket"}, Required: true},
			&schema.SchemaField{Name: "from_pocket_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "pocket"}, Required: true},
			&schema.SchemaField{Name: "created_at", Type: schema.FieldTypeDate, Required: false},
			&schema.SchemaField{Name: "updated_at", Type: schema.FieldTypeDate, Required: false},
		),
	}
	saveCollection(app, collection)
}
