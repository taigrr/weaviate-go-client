package schema

import (
	"context"
	"fmt"
	"net/http"

	"github.com/weaviate/weaviate-go-client/v4/weaviate/connection"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/except"
)

// TenantsDeleter builder object to delete tenants
type TenantsDeleter struct {
	connection *connection.Connection
	className  string
	tenants    []string
}

// WithClassName specifies the class that tenants will be added to
func (td *TenantsDeleter) WithClassName(className string) *TenantsDeleter {
	td.className = className
	return td
}

// WithTenants specifies tenants that will be added to the class
func (td *TenantsDeleter) WithTenants(tenants ...string) *TenantsDeleter {
	td.tenants = tenants
	return td
}

// Add tenants to the class specified in the builder
func (td *TenantsDeleter) Do(ctx context.Context) error {
	path := fmt.Sprintf("/schema/%v/tenants", td.className)
	responseData, err := td.connection.RunREST(ctx, path, http.MethodDelete, td.tenants)
	return except.CheckResponseDataErrorAndStatusCode(responseData, err, 200)
}
