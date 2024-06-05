package connect

import (
	"context"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"net/url"
	"github.com/skip2/go-qrcode"

	graphqlClient "github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/graphql"
)

const APP_CLIP_BASE_URL = "https://appclip.apple.com/id?p=network.gandalf.connect.Clip"
const SAURON_BASE_URL = "https://sauron.gandalf.network/public/gql"

const (
	InvalidService GandalfErrorCode = iota
	InvalidPublicKey
	InvalidRedirectURL
	QRCodeGenNotSupported
	QRCodeNotGenerated
)

func (e *GandalfError) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.Message, e.Code)
}

func GenerateURL(publicKey string, redirectURL string, input Services) (string, error) {
	services, err := runValidation(publicKey, redirectURL, input)
	if err != nil {
		return "", err
	}

	servicesJSON := servicesToJSON(services)

	return encodeComponents(servicesJSON, publicKey, redirectURL), nil
}

func GenerateQRCode(publicKey string, redirectURL string, input Services) (string, error) {
	if publicKey == "" || redirectURL == "" || input == nil {
		return "", &GandalfError{
			Message: "Invalid input parameters",
			Code:    QRCodeGenNotSupported,
		}
	}

	services, err := runValidation(publicKey, redirectURL, input)
	if err != nil {
		return "", err
	}

	servicesJSON := servicesToJSON(services)
	appClipURL := encodeComponents(servicesJSON, redirectURL, publicKey)

	qrCode, err := qrcode.New(appClipURL, qrcode.Medium)
	if err != nil {
		return "", &GandalfError{
			Message: "QRCode Generation Error",
			Code:    QRCodeNotGenerated,
		}
	}

	qrCodeData, err := qrCode.PNG(256)
	if err != nil {
		return "", &GandalfError{
			Message: "QRCode Generation Error",
			Code:    QRCodeNotGenerated,
		}
	}
	qrCodeURL := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(qrCodeData))
	return qrCodeURL, nil
}

func introspectSauron() IntrospectionResult {
	client := graphqlClient.NewClient(SAURON_BASE_URL)
	req := graphqlClient.NewRequest(introspectionQuery)

	ctx := context.Background()

	var respData IntrospectionResult

	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatalf("Error making introspection query: %v", err)
	}
	return respData
}

func validateRedirectURL(rawURL string) error {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return &GandalfError{
			Message: "Invalid redirect URL",
			Code:    InvalidRedirectURL,
		}
	}
	return nil
}

func validatePublicKey(publicKey string) bool {
	return publicKeyRequest(publicKey)
}

func getSupportedServices() []Value {
	gqlSchema := introspectSauron()
	for _, val := range gqlSchema.Schema.Types {
		if val.Kind == "ENUM" && val.Name == "Source" {
			return val.EnumValues
		}
	}
	return nil
}

func validateInputServices(input Services) (map[string]Service, error) {
	// supportedServicesInterface := getSupportedServices()
	supportedServices := getSupportedServices()
	
	serviceMap := make(map[string]Value)
	for _, val := range supportedServices {
		serviceMap[val.Name] = val
	}

	var unsupportedServices []string
	requiredServices := 0
	cleanedServices :=  make(map[string]Service)
	for _, val := range input {
		key := strings.ToUpper(val.Name)
		if _, found := serviceMap[key]; !found {
			unsupportedServices = append(unsupportedServices, key)
			continue
		}
		
		if val.Name != "" {
			requiredServices++
		}
		cleanedServices[key] = val
	}

	if len(unsupportedServices) > 0 {
		return nil, &GandalfError{
			Message: fmt.Sprintf("These services %s are unsupported", strings.Join(unsupportedServices, " ")),
			Code:    InvalidService,
		}
	}

	if requiredServices < 1 {
		return nil, &GandalfError{
			Message: "At least one service has to be required",
			Code:    InvalidService,
		}
	}

	return cleanedServices, nil

}

func publicKeyRequest(publicKey string) bool {
	graphqlRequest := graphqlClient.NewRequest(`
	query GetAppByPublicKey($publicKey: String!) {
	  getAppByPublicKey(
		publicKey: $publicKey
		) {
		appName
		gandalfID
	  }
	}
  `)
  
	graphqlRequest.Var("publicKey", publicKey)
  	client := graphqlClient.NewClient(SAURON_BASE_URL)

	ctx := context.Background()

	var graphqlResponse map[string]interface{}

	if err := client.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		log.Fatalf("Error making publicKey request query: %v", err)
	}

	responseData, ok := graphqlResponse["getAppByPublicKey"].(map[string]interface{})
	if !ok {
		log.Fatalf("Unexpected response structure: %v", graphqlResponse)
	}

	body, err := json.Marshal(responseData)
	if err != nil {
		return false
	}

	var respData Application
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return false
	}
	return respData.GandalfID > 0
}

func runValidation(publicKey string, redirectURL string, input Services) (map[string]Service, error) {
	isPublicKeyValid := validatePublicKey(publicKey)
	if !isPublicKeyValid {
		return nil, &GandalfError{
			Message: "Invalid public key",
			Code:    InvalidPublicKey,
		}
	}

	err := validateRedirectURL(redirectURL)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	services, err := validateInputServices(input)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func encodeComponents(servicesJSON []byte, publicKey string, redirectURL string) string {
	encodedServices := url.QueryEscape(string(servicesJSON))
	encodedRedirectURL := url.QueryEscape(redirectURL)
	encodedPublicKey := url.QueryEscape(publicKey)

	return fmt.Sprintf("%s&services=%s&redirectUrl=%s&publicKey=%s", APP_CLIP_BASE_URL, encodedServices, encodedRedirectURL, encodedPublicKey)
}

func servicesToJSON(services map[string]Service) []byte {
	var servicesSlice []Service
	for _, service := range services {
		servicesSlice = append(servicesSlice, service)
	}

	servicesJSON, err := json.Marshal(servicesSlice)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	return servicesJSON
}