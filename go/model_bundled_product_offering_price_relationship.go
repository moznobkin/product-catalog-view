/*
 * Product Catalog View
 *
 * Product Catalog View
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// This represents a bundling pricing relationship, allowing a price to be composed of multiple other prices (e.g. a recurring charge and a onetime charge).
type BundledProductOfferingPriceRelationship struct {
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
	// hyperlink reference of the bundled product offering price
	Href string `json:"href,omitempty"`
	// Unique identifier of the bundled product offering price
	Id string `json:"id"`
	// Name of the bundled product offering price
	Name string `json:"name,omitempty"`
}
