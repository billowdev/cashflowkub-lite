package collections

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var CashflowinCollection = "cashflowin"

func NewCashflowinCollection(app *pocketbase.PocketBase) {
	collection := &models.Collection{
		Name: CashflowinCollection,
		Schema: schema.NewSchema(
			&schema.SchemaField{Name: "amount", Type: schema.FieldTypeNumber, Required: true},
			&schema.SchemaField{Name: "description", Type: schema.FieldTypeText, Required: false},
			&schema.SchemaField{Name: "user_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "user"}, Required: true},
			&schema.SchemaField{Name: "pocket_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "pocket"}, Required: true},
			&schema.SchemaField{Name: "category_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "category"}, Required: true},
			&schema.SchemaField{Name: "receipt", Type: schema.FieldTypeFile, Required: false},
			&schema.SchemaField{Name: "created_at", Type: schema.FieldTypeDate, Required: false},
			&schema.SchemaField{Name: "updated_at", Type: schema.FieldTypeDate, Required: false},
		),
	}

	saveCollection(app, collection)

}
