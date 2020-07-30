/*
 * Product Catalog View
 *
 * Product Catalog View
 *
 * API version: 0.0.3-oas3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A condition under which a ProductOffering is made available to Customers. For instance, a productOffering can be offered with multiple commitment periods.
type ProductOfferingTerm struct {
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	//  hyperlink reference to the schema describing this resource.
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
	// Description of the productOfferingTerm
	Description string `json:"description,omitempty"`

	Duration *Quantity `json:"duration,omitempty"`
	// Name of the productOfferingTerm
	Name string `json:"name,omitempty"`

	ValidFor *TimePeriod `json:"validFor,omitempty"`
}
