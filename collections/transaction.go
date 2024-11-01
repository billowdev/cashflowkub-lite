package collections

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var TransactionCollection = "transaction"

func NewTransactionCollection(app *pocketbase.PocketBase) {
	collection := &models.Collection{
		Name: TransactionCollection,
		Schema: schema.NewSchema(
			&schema.SchemaField{Name: "type", Type: schema.FieldTypeSelect, Options: schema.SelectOptions{Values: []string{"in", "out", "transfer"}}, Required: true},
			&schema.SchemaField{Name: "cashflowin_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "cashflowin"}, Required: false},
			&schema.SchemaField{Name: "cashflowout_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "cashflowout"}, Required: false},
			&schema.SchemaField{Name: "transfer_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "transfer"}, Required: false},
			&schema.SchemaField{Name: "user_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "user"}, Required: true},
			&schema.SchemaField{Name: "created_at", Type: schema.FieldTypeDate, Required: false},
			&schema.SchemaField{Name: "updated_at", Type: schema.FieldTypeDate, Required: false},
		),
	}

	saveCollection(app, collection)

}
