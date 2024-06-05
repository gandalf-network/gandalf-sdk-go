package connect


type GandalfErrorCode int

// GandalfError is a custom error type for validation errors
type GandalfError struct {
	Message string
	Code    GandalfErrorCode
}

type Application struct {
	// The human-readable name of the application.
	AppName string `json:"appName"`
	// A public key associated with the application, used for cryptographic operations such as
	// verifying the identity of the application.
	PublicKey string `json:"publicKey"`
	// The URL pointing to the icon graphic for the application. This URL should link to an image
	// that visually represents the application, aiding in its identification and branding.
	IconURL string `json:"iconURL"`
	// A unique identifier assigned to the application upon registration.
	GandalfID int64 `json:"gandalfID"`
	// The address of the user who registered the application.
	AppRegistrar string `json:"appRegistrar"`
}
// type Application map[string]interface{}

type SupportedService struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason"`
}

type SupportedServices []Value

type Service struct {
	Name string
	Status bool
}

type Services []Service

// Type represents a GraphQL type with various properties like kind, name, description, etc.
type Type struct {
	Kind          string  `json:"kind"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Fields        []Field `json:"fields"`
	InputFields   []Field `json:"inputFields"`
	Interfaces    []Type  `json:"interfaces"`
	EnumValues    []Value `json:"enumValues"`
	PossibleTypes []Type  `json:"possibleTypes"`
	OfType        *Type   `json:"ofType"`
}

// Field represents a field in a GraphQL type with various properties
type Field struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Type              Type    `json:"type"`
	DefaultValue      string  `json:"defaultValue"`
	IsDeprecated      bool    `json:"isDeprecated"`
	DeprecationReason string  `json:"deprecationReason"`
	Args              []Field `json:"args"`
}

// Value represents a value in an enum type
type Value struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason"`
}

// IntrospectionResult represents the schema structure received from the GraphQL introspection query
type IntrospectionResult struct {
	Schema struct {
		Types []Type `json:"types"`
	} `json:"__schema"`
}

const introspectionQuery = `
	query {
		__schema {
			types {
				kind
				name
				description
				fields(includeDeprecated: true) {
					name
					description
					args {
						name
						description
						type {
							kind
							name
							ofType {
								kind
								name
								ofType {
									kind
									name
									ofType {
										kind
										name
									}
								}
							}
						}
						defaultValue
					}
					type {
						kind
						name
						ofType {
							kind
							name
							ofType {
								kind
								name
								ofType {
									kind
									name
								}
							}
						}
					}
					isDeprecated
					deprecationReason
				}
				inputFields {
					name
					description
					type {
						kind
						name
						ofType {
							kind
							name
							ofType {
								kind
								name
							}
						}
					}
					defaultValue
				}
				interfaces {
					kind
					name
					ofType {
						kind
						name
					}
				}
				enumValues(includeDeprecated: true) {
					name
					description
					isDeprecated
					deprecationReason
				}
				possibleTypes {
					kind
					name
					ofType {
						kind
						name
					}
				}
			}
		}
	}
`