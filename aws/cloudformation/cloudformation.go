package cloudformation

// Parameter defines a param for a goformation template
type Parameter struct {
	Type                  string   `json:"Type"`
	Description           string   `json:"Description,omitempty"`
	Default               string   `json:"Default,omitempty"`
	AllowedPattern        string   `json:"AllowedPattern,omitempty"`
	AllowedValues         []string `json:"AllowedValues,omitempty"`
	ConstraintDescription string   `json:"ConstraintDescription,omitempty"`
	MaxLength             int      `json:"MaxLength,omitempty"`
	MinLength             int      `json:"MinLength,omitempty"`
	MaxValue              float64  `json:"MaxValue,omitempty"`
	MinValue              float64  `json:"MinValue,omitempty"`
	NoEcho                bool     `json:"NoEcho,omitempty"`
}

// Output defines a output for a goformation template
type Output struct {
	Value       string            `json:"Value"`
	Description string            `json:"Description,omitempty"`
	Export      map[string]string `json:"Export,omitempty"`
}
