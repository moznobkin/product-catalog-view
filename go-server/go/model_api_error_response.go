/*
 * Product Catalog View
 *
 * Product Catalog View
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type ApiErrorResponse struct {

	Errors []ModelError `json:"errors,omitempty"`

	// Api Exception Error
	Message string `json:"message,omitempty"`

	// HTTP Status
	Status string `json:"status,omitempty"`

	// Current date timestamp
	Timestamp time.Time `json:"timestamp,omitempty"`
}
