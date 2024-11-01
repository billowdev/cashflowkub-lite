package collections

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var UserCollection = "users"

func NewUserCollection(app *pocketbase.PocketBase) {
	collection := &models.Collection{
		Name: UserCollection,
		Schema: schema.NewSchema(
			&schema.SchemaField{Name: "username", Type: schema.FieldTypeText, Required: true},
			&schema.SchemaField{Name: "hash_password", Type: schema.FieldTypeText, Required: true},
			&schema.SchemaField{Name: "first_name", Type: schema.FieldTypeText, Required: false},
			&schema.SchemaField{Name: "last_name", Type: schema.FieldTypeText, Required: false},
			&schema.SchemaField{Name: "phone", Type: schema.FieldTypeText, Required: false},
			&schema.SchemaField{Name: "email", Type: schema.FieldTypeEmail, Required: true, Unique: true},
			&schema.SchemaField{Name: "role", Type: schema.FieldTypeSelect, Options: schema.SelectOptions{Values: []string{"user", "premium", "admin"}}, Required: true},
			&schema.SchemaField{Name: "is_active", Type: schema.FieldTypeBool, Required: true},
			&schema.SchemaField{Name: "created_at", Type: schema.FieldTypeDate, Required: false},
			&schema.SchemaField{Name: "updated_at", Type: schema.FieldTypeDate, Required: false},
		),
	}

	saveCollection(app, collection)
}
