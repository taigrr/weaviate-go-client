package data

import (
	"context"
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate-go-client/v4/weaviate/connection"
	"github.com/semi-technologies/weaviate-go-client/v4/weaviate/except"
	"github.com/semi-technologies/weaviate/entities/models"
)

// Updater builder to update property values in a data object
type Updater struct {
	connection     *connection.Connection
	id             string
	className      string
	propertySchema models.PropertySchema
	withMerge      bool
	version        string
}

// WithID specifies the uuid of the object about to be  updated
func (updater *Updater) WithID(uuid string) *Updater {
	updater.id = uuid
	return updater
}

// WithClassName specifies the class of the object about to be updated
func (updater *Updater) WithClassName(className string) *Updater {
	updater.className = className
	return updater
}

// WithProperties specifies the property schema of the class about to be updated
func (updater *Updater) WithProperties(propertySchema models.PropertySchema) *Updater {
	updater.propertySchema = propertySchema
	return updater
}

// WithMerge indicates that the object should be merged with the existing object instead of replacing it
func (updater *Updater) WithMerge() *Updater {
	updater.withMerge = true
	return updater
}

// Do update the data object specified in the builder
func (updater *Updater) Do(ctx context.Context) error {
	path := buildObjectsPath(updater.id, updater.className, updater.className)
	httpMethod := http.MethodPut
	expectedStatuscode := 200
	if updater.withMerge {
		httpMethod = http.MethodPatch
		expectedStatuscode = 204
	}
	responseData, responseErr := updater.runUpdate(ctx, path, httpMethod)
	return except.CheckResponseDataErrorAndStatusCode(responseData, responseErr, expectedStatuscode)
}

func (updater *Updater) runUpdate(ctx context.Context, path string, httpMethod string) (*connection.ResponseData, error) {
	object := models.Object{
		Class:      updater.className,
		ID:         strfmt.UUID(updater.id),
		Properties: updater.propertySchema,
	}
	return updater.connection.RunREST(ctx, path, httpMethod, object)
}
