/*
 * Weaviate
 *
 * Open Source Search Graph (GraphQL/RESTful/P2P)
 *
 * API version: 0.22.17
 * Contact: hello@semi.technology
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// List of Actions.
type ActionsListResponse struct {
	// The actual list of Actions.
	Actions []Action           `json:"actions,omitempty"`
	Deprecations []Deprecation `json:"deprecations,omitempty"`
	// The total number of Actions for the query. The number of items in a response may be smaller due to paging.
	TotalResults int64 `json:"totalResults,omitempty"`
}
