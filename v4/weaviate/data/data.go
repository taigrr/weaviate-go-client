package data

import (
	"github.com/semi-technologies/weaviate-go-client/v4/weaviate/connection"
)

// API Contains all the builders required to access the weaviate data API
type API struct {
	connection *connection.Connection
	version    string
}

// New {semanticKind} api group from connection
func New(con *connection.Connection, version string) *API {
	return &API{connection: con, version: version}
}

// Creator get a builder to create a data object
func (data *API) Creator() *Creator {
	return &Creator{
		connection: data.connection,
		version:    data.version,
	}
}

// ObjectsGetter get a builder to get an Action
func (data *API) ObjectsGetter() *ObjectsGetter {
	return &ObjectsGetter{
		connection:           data.connection,
		additionalProperties: []string{},
		version:              data.version,
	}
}

// Deleter get a builder to delete data objects
func (data *API) Deleter() *Deleter {
	return &Deleter{
		connection: data.connection,
		version:    data.version,
	}
}

// Updater get a builder to update a data object
func (data *API) Updater() *Updater {
	return &Updater{
		connection: data.connection,
		withMerge:  false,
		version:    data.version,
	}
}

// Validator get a builder to validate a data object definition
func (data *API) Validator() *Validator {
	return &Validator{
		connection: data.connection,
	}
}

// Checker get a builder to check data object existence
func (data *API) Checker() *Checker {
	return &Checker{
		connection: data.connection,
		version:    data.version,
	}
}

// ReferencePayloadBuilder get a builder to create the payloads that reference an object
func (data *API) ReferencePayloadBuilder() *ReferencePayloadBuilder {
	return &ReferencePayloadBuilder{
		connection: data.connection,
	}
}

// ReferenceCreator get a builder to add references to data objects
func (data *API) ReferenceCreator() *ReferenceCreator {
	return &ReferenceCreator{
		connection: data.connection,
	}
}

// ReferenceReplacer get a builder to replace references on a data object
func (data *API) ReferenceReplacer() *ReferenceReplacer {
	return &ReferenceReplacer{
		connection: data.connection,
	}
}

// ReferenceDeleter get a builder to delete references on a data object
func (data *API) ReferenceDeleter() *ReferenceDeleter {
	return &ReferenceDeleter{
		connection: data.connection,
	}
}
