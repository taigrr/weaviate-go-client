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

// An array of available words and contexts.
type C11yWordsResponse struct {
	ConcatenatedWord *C11yWordsResponseConcatenatedWord `json:"concatenatedWord,omitempty"`
	// Weighted results for per individual word
	IndividualWords []C11yWordsResponseIndividualWords `json:"individualWords,omitempty"`
}
