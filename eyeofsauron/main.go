package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/machinebox/graphql"
)

// IntrospectionResult represents the schema structure received from the GraphQL introspection query
type IntrospectionResult struct {
	Schema struct {
		Types []Type `json:"types"`
	} `json:"__schema"`
}

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

func main() {

	client := graphql.NewClient("http://localhost:1000/public/gql")
	req := graphql.NewRequest(introspectionQuery)

	ctx := context.Background()

	var respData IntrospectionResult

	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatalf("Error making introspection query: %v", err)
	}

	sdl := convertToSDL(respData)
	writeToFile("schema.graphql", sdl)

	typesMap := buildtypesMap(respData)
	interfaceImplementations := buildInterfaceImplementationsMap(respData)

	var stringBuilder *strings.Builder = &strings.Builder{}
	generateFragments(stringBuilder, typesMap, interfaceImplementations)
	generateQueries(stringBuilder, respData, typesMap, interfaceImplementations)
	generateMutations(stringBuilder, respData, typesMap, interfaceImplementations)

	writeToFile("genqlient.graphql", stringBuilder.String())

	fmt.Println("Schema, fragments, queries, and mutations dumped to files.")
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

func convertToSDL(introspection IntrospectionResult) string {
	var sb strings.Builder
	implementsMap := buildImplementsMap(introspection)

	for _, t := range introspection.Schema.Types {
		if strings.HasPrefix(t.Name, "__") {
			continue // Skip introspection types
		}
		switch t.Kind {
		case "OBJECT":
			writeObjectType(&sb, t, implementsMap)
		case "ENUM":
			writeEnumType(&sb, t)
		case "SCALAR":
			writeScalarType(&sb, t)
		case "INTERFACE":
			writeInterfaceType(&sb, t)
		case "INPUT_OBJECT":
			writeInputObjectType(&sb, t)
		case "UNION":
			writeUnionType(&sb, t)
		}
	}

	return sb.String()
}

func buildImplementsMap(introspection IntrospectionResult) map[string][]string {
	implementsMap := make(map[string][]string)
	for _, t := range introspection.Schema.Types {
		for _, iface := range t.Interfaces {
			implementsMap[t.Name] = append(implementsMap[t.Name], iface.Name)
		}
	}
	return implementsMap
}

func writeObjectType(sb *strings.Builder, t Type, implementsMap map[string][]string) {
	sb.WriteString(fmt.Sprintf("type %s", t.Name))
	if implements, ok := implementsMap[t.Name]; ok {
		sb.WriteString(fmt.Sprintf(" implements %s", strings.Join(implements, " & ")))
	}
	sb.WriteString(" {\n")
	for _, field := range t.Fields {
		writeField(sb, field)
	}
	sb.WriteString("}\n\n")
}

func writeEnumType(sb *strings.Builder, t Type) {
	sb.WriteString(fmt.Sprintf("enum %s {\n", t.Name))
	for _, value := range t.EnumValues {
		sb.WriteString(fmt.Sprintf("  %s\n", value.Name))
	}
	sb.WriteString("}\n\n")
}

func writeScalarType(sb *strings.Builder, t Type) {
	sb.WriteString(fmt.Sprintf("scalar %s\n\n", t.Name))
}

func writeInterfaceType(sb *strings.Builder, t Type) {
	sb.WriteString(fmt.Sprintf("interface %s {\n", t.Name))
	for _, field := range t.Fields {
		writeField(sb, field)
	}
	sb.WriteString("}\n\n")
}

func writeInputObjectType(sb *strings.Builder, t Type) {
	sb.WriteString(fmt.Sprintf("input %s {\n", t.Name))
	for _, inputField := range t.InputFields {
		sb.WriteString(fmt.Sprintf("  %s: %s\n", inputField.Name, formatType(inputField.Type)))
	}
	sb.WriteString("}\n\n")
}

func writeUnionType(sb *strings.Builder, t Type) {
	sb.WriteString(fmt.Sprintf("union %s = ", t.Name))
	for i, pt := range t.PossibleTypes {
		if i > 0 {
			sb.WriteString(" | ")
		}
		sb.WriteString(pt.Name)
	}
	sb.WriteString("\n\n")
}

func writeField(sb *strings.Builder, field Field) {
	sb.WriteString(fmt.Sprintf("  %s", field.Name))
	if len(field.Args) > 0 {
		sb.WriteString("(")
		for i, arg := range field.Args {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%s: %s", arg.Name, formatType(arg.Type)))
		}
		sb.WriteString(")")
	}
	sb.WriteString(fmt.Sprintf(": %s\n", formatType(field.Type)))
}

func formatType(t Type) string {
	if t.OfType != nil {
		if t.Kind == "LIST" {
			return fmt.Sprintf("[%s]", formatType(*t.OfType))
		} else if t.Kind == "NON_NULL" {
			return fmt.Sprintf("%s!", formatType(*t.OfType))
		}
		return formatType(*t.OfType)
	}
	return t.Name
}

func buildtypesMap(introspection IntrospectionResult) map[string]Type {
	typesMap := make(map[string]Type)
	for _, t := range introspection.Schema.Types {
		if strings.HasPrefix(t.Name, "__") || t.Name == "Query" {
			continue // Skip introspection types
		}

		typesMap[t.Name] = t
	}

	return typesMap
}

func buildInterfaceImplementationsMap(introspection IntrospectionResult) map[string][]string {
	interfaceImplementations := make(map[string][]string)
	for _, t := range introspection.Schema.Types {
		if strings.HasPrefix(t.Name, "__") || t.Name == "Query" {
			continue // Skip introspection types
		}

		for _, possibleType := range t.PossibleTypes {
			interfaceImplementations[t.Name] = append(interfaceImplementations[t.Name], possibleType.Name)
		}
	}
	return interfaceImplementations
}

func findInnermostType(t *Type) *Type {
	if t == nil {
		return nil
	}
	if t.Kind != "LIST" && t.Kind != "NON_NULL" {
		return t
	}
	if t.OfType != nil {
		return findInnermostType(t.OfType)
	}
	return t
}

func generateFragments(sb *strings.Builder, typesMap map[string]Type, interfaceImplementations map[string][]string) string {
	for typeName, typeValue := range typesMap {
		fields := typeValue.Fields
		if len(typeValue.Interfaces) == 0 {
			continue
		}

		sb.WriteString(fmt.Sprintf("fragment %sFragment on %s {\n", typeName, typeName))
		for _, field := range fields {
			sb.WriteString(fmt.Sprintf("  %s", field.Name))

			innermostType := findInnermostType(&field.Type)
			if innermostType.Kind == "INTERFACE" {
				sb.WriteString(fmt.Sprintf(" {\n    ...%sFragment\n  }\n", innermostType.Name))
			} else if innermostType.Kind == "OBJECT" {
				sb.WriteString(" {\n")
				writeFieldSelection(sb, *innermostType, typesMap, interfaceImplementations)
				sb.WriteString("  }\n")
			} else {
				sb.WriteString("\n")
			}

		}
		sb.WriteString("}\n\n")

	}

	return sb.String()
}

func generateQueries(sb *strings.Builder, introspection IntrospectionResult, typesMap map[string]Type, interfaceImplementations map[string][]string) string {
	for _, t := range introspection.Schema.Types {
		if t.Kind == "OBJECT" && t.Name == "Query" {
			for _, field := range t.Fields {
				writeOperation(sb, "query", field, typesMap, interfaceImplementations)
			}
		}
	}

	return sb.String()
}

func generateMutations(sb *strings.Builder, introspection IntrospectionResult, typesMap map[string]Type, interfaceImplementations map[string][]string) string {
	for _, t := range introspection.Schema.Types {
		if t.Kind == "OBJECT" && t.Name == "Mutation" {
			for _, field := range t.Fields {
				writeOperation(sb, "mutation", field, typesMap, interfaceImplementations)
			}
		}
	}

	return sb.String()
}

func writeOperation(sb *strings.Builder, opType string, field Field, typesMap map[string]Type, interfaceImplementations map[string][]string) {
	sb.WriteString(fmt.Sprintf("%s %s(", opType, field.Name))
	for i, arg := range field.Args {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("$%s: %s", arg.Name, formatType(arg.Type)))
	}
	sb.WriteString(") {\n")
	sb.WriteString(fmt.Sprintf("  %s(", field.Name))
	for i, arg := range field.Args {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%s: $%s", arg.Name, arg.Name))
	}
	sb.WriteString(") {\n")
	writeFieldSelection(sb, field.Type, typesMap, interfaceImplementations)
	sb.WriteString("  }\n")
	sb.WriteString("}\n\n")
}

func writeFieldSelection(sb *strings.Builder, t Type, typesMap map[string]Type, interfaceImplementations map[string][]string) {
	if t.OfType != nil {
		writeFieldSelection(sb, *t.OfType, typesMap, interfaceImplementations)
		return
	}

	if t.Kind == "OBJECT" || t.Kind == "INTERFACE" || t.Kind == "UNION" {
		sb.WriteString("    ... on ")
		sb.WriteString(t.Name)
		sb.WriteString(" {\n")
		typeValue := typesMap[t.Name]
		for _, field := range typeValue.Fields {
			sb.WriteString(fmt.Sprintf("      %s", field.Name))
			if len(field.Args) > 0 {
				sb.WriteString("(")
				for i, arg := range field.Args {
					if i > 0 {
						sb.WriteString(", ")
					}
					sb.WriteString(fmt.Sprintf("%s: %s", arg.Name, formatType(arg.Type)))
				}
				sb.WriteString(")")
			}

			if field.Type.OfType != nil {
				if field.Type.OfType.Kind == "INTERFACE" {
					sb.WriteString(" {\n")
					for _, impl := range interfaceImplementations[field.Type.OfType.Name] {
						sb.WriteString(fmt.Sprintf("        ...%sFragment\n", impl))
					}
					sb.WriteString("      }\n")
				} else if field.Type.OfType.Kind == "OBJECT" || field.Type.OfType.Kind == "UNION" || field.Type.OfType.Kind == "LIST" {
					sb.WriteString(" {\n")
					writeFieldSelection(sb, *field.Type.OfType, typesMap, interfaceImplementations)
					sb.WriteString("      }\n")
				} else {
					sb.WriteString("\n")
				}
			} else {
				sb.WriteString("\n")
			}
		}
		sb.WriteString("    }\n")
	} else {
		sb.WriteString(fmt.Sprintf("      %s\n", t.Name))
	}
}

func writeToFile(filename, content string) {
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		log.Fatalf("Error writing to file %s: %v", filename, err)
	}
}
