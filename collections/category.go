package collections

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var CategoryCollection = "category"

func NewCategoryCollection(app *pocketbase.PocketBase) {
	collection := &models.Collection{
		Name: CategoryCollection,
		Schema: schema.NewSchema(
			&schema.SchemaField{Name: "name", Type: schema.FieldTypeText, Required: true},
			&schema.SchemaField{Name: "description", Type: schema.FieldTypeText, Required: false},
			&schema.SchemaField{Name: "type", Type: schema.FieldTypeSelect, Options: schema.SelectOptions{Values: []string{"income", "expense", "investment", "saving"}}, Required: true},
			&schema.SchemaField{Name: "user_id", Type: schema.FieldTypeRelation, Options: schema.RelationOptions{CollectionId: "user"}, Required: false},
			&schema.SchemaField{Name: "created_at", Type: schema.FieldTypeDate, Required: false},
			&schema.SchemaField{Name: "updated_at", Type: schema.FieldTypeDate, Required: false},
		),
	}
	saveCollection(app, collection)
}
