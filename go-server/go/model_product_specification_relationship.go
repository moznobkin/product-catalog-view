/*
 * Product Catalog View
 *
 * Product Catalog View
 *
 * API version: 0.0.3-oas3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A migration, substitution, dependency or exclusivity relationship between/among product specifications.
type ProductSpecificationRelationship struct {
	// Reference of the productSpecification
	Href string `json:"href,omitempty"`
	// Unique identifier of the productSpecification
	Id string `json:"id,omitempty"`

	RelationshipType string `json:"relationshipType,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
}