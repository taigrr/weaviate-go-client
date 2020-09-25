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

type C11yVectorBasedQuestionInner struct {
	// Vectorized classname.
	ClassVectors []float32 `json:"classVectors,omitempty"`
	// Vectorized properties.
	ClassProps []C11yVectorBasedQuestionInnerClassProps `json:"classProps,omitempty"`
}
