package main

const gqlgenConfig = `
schema: schema.graphql
operations:
- genqlient.graphql
generated: generated.go
bindings:
  Int64:
    type: github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/graphqlTypes.Int64
  Date:
    type: github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/graphqlTypes.Date
  UUID:
    type: github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/graphqlTypes.UUID
  Time:
    type: time.Time
`

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
