package queue

// Policy wraps the JSON policy
type Policy struct {
	Version   string      `json:"Version"`
	ID        string      `json:"Id"`
	Statement []Statement `json:"Statement"`
}

// Statement defines the QueuePolicy Statement
type Statement struct {
	Sid       string    `json:"Sid"`
	Effect    string    `json:"Effect"`
	Principal string    `json:"Principal"`
	Action    []string  `json:"Action"`
	Resource  string    `json:"Resource"`
	Condition Condition `json:"Condition"`
}

// Condition defines the Condition for Statments
type Condition struct {
	ArnEquals ArnEquals `json:"ArnEquals"`
}

// ArnEquals is a mapping for the SourceArn
type ArnEquals struct {
	AwsSourceArn string `json:"aws:SourceArn"`
}
