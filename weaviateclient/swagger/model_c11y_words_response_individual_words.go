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

type C11yWordsResponseIndividualWords struct {
	Word string                 `json:"word,omitempty"`
	InC11y bool                 `json:"inC11y,omitempty"`
	Info *C11yWordsResponseInfo `json:"info,omitempty"`
}
