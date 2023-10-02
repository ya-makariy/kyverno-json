package v1alpha1

// ContextEntry adds variables and data sources to a rule Context.
type ContextEntry struct {
	// Name is the variable name.
	Name string `json:"name"`

	// Variable defines an arbitrary JMESPath context variable that can be defined inline.
	Variable *Variable `json:"variable,omitempty"`
}