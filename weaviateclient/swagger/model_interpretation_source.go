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

// This underscore property contains additional info about the how the class was vectorized
type InterpretationSource struct {
	Concept string `json:"concept,omitempty"`
	Weight float32 `json:"weight,omitempty"`
	Occurrence float32 `json:"occurrence,omitempty"`
}
