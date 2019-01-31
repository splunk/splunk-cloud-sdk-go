/*
 * Action Service
 *
 * A service that receives incoming notifications and uses pre-defined templates (action objects) to turn those notifications into meaningful actions.
 *
 * API version: v1beta2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package generated

type MutableDeprecatedEmailAction struct {
	// Human readable name title for the action. Must be less than 128 characters.
	Title *string `json:"title,omitempty"`
	HtmlPart *string `json:"htmlPart,omitempty"`
	SubjectPart *string `json:"subjectPart,omitempty"`
	TextPart *string `json:"textPart,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
}
