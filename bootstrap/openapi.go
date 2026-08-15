package bootstrap

import (
	"encoding/json"
	"fmt"
)

// SecuritySchemeName e o nome do esquema referenciado pelas rotas protegidas.
const SecuritySchemeName = "authorization"

// WithSecurityScheme injeta components.securitySchemes no documento gerado e
// marca as operacoes indicadas como protegidas.
//
// O openapi-builder ainda nao modela securitySchemes, e a especificacao OpenAPI
// manda ignorar parametros de header chamados Accept, Content-Type e
// Authorization -- o swagger-ui implementa essa regra e descarta o valor
// digitado. Declarar o esquema e a unica forma de o botao Authorize enviar o
// header.
//
// O esquema usado e apiKey (e nao http/basic) porque o guard deste projeto
// espera o header inteiro no formato "Basic <token>", com o token sendo
// base64(login+senha) sem separador -- diferente do base64(login:senha) que o
// swagger-ui montaria sozinho num esquema http/basic.
func WithSecurityScheme(spec []byte, operations map[string][]string) ([]byte, error) {
	document := map[string]any{}

	if err := json.Unmarshal(spec, &document); err != nil {
		return nil, fmt.Errorf("spec invalida: %w", err)
	}

	components, _ := document["components"].(map[string]any)

	if components == nil {
		components = map[string]any{}
	}

	components["securitySchemes"] = map[string]any{
		SecuritySchemeName: map[string]any{
			"type":        "apiKey",
			"in":          "header",
			"name":        "Authorization",
			"description": `Cole o header completo, ex: "Basic <token>"`,
		},
	}

	document["components"] = components

	paths, _ := document["paths"].(map[string]any)

	for path, methods := range operations {
		item, _ := paths[path].(map[string]any)

		if item == nil {
			return nil, fmt.Errorf("path %q nao existe na spec", path)
		}

		for _, method := range methods {
			operation, _ := item[method].(map[string]any)

			if operation == nil {
				return nil, fmt.Errorf("operacao %s %s nao existe na spec", method, path)
			}

			operation["security"] = []any{
				map[string]any{SecuritySchemeName: []any{}},
			}
		}
	}

	return json.MarshalIndent(document, "", "  ")
}
