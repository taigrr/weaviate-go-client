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

type Action struct {
	// Type of the Action, defined in the schema.
	Class string                 `json:"class,omitempty"`
	VectorWeights *VectorWeights `json:"vectorWeights,omitempty"`
	Schema *PropertySchema       `json:"schema,omitempty"`
	Meta *UnderscoreProperties   `json:"meta,omitempty"`
	// ID of the Action.
	Id string `json:"id,omitempty"`
	// Timestamp of creation of this Action in milliseconds since epoch UTC.
	CreationTimeUnix int64 `json:"creationTimeUnix,omitempty"`
	// Timestamp of the last update made to the Action since epoch UTC.
	LastUpdateTimeUnix int64 `json:"lastUpdateTimeUnix,omitempty"`
	// If this object was subject of a classificiation, additional meta info about this classification is available here. (Underscore properties are optional, include them using the ?include=_<propName> parameter)
	Classification *UnderscorePropertiesClassification `json:"_classification,omitempty"`
	// This object's position in the Contextionary vector space. (Underscore properties are optional, include them using the ?include=_<propName> parameter)
	Vector *C11yVector `json:"_vector,omitempty"`
	// Additional information about how this property was interpreted at vectorization. (Underscore properties are optional, include them using the ?include=_<propName> parameter)
	Interpretation *Interpretation `json:"_interpretation,omitempty"`
	// Additional information about the neighboring concepts of this element
	NearestNeighbors *NearestNeighbors `json:"_nearestNeighbors,omitempty"`
	// A feature projection of the object's vector into lower dimensions for visualization
	FeatureProjection *FeatureProjection `json:"_featureProjection,omitempty"`
}
