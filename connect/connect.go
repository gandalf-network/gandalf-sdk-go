package connect

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/constants"
	"github.com/skip2/go-qrcode"
	graphqlClient "github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/graphql"
)

var (
	IOS_APP_CLIP_BASE_URL = "https://appclip.apple.com/id?p=network.gandalf.connect.Clip"
	ANDROID_APP_CLIP_BASE_URL   =  "https://auth.gandalf.network"
	UNIVERSAL_APP_CLIP_BASE_URL = "https://auth.gandalf.network"
	SAURON_BASE_URL = "https://sauron.gandalf.network/public/gql"
)

const (
	InvalidService GandalfErrorCode = iota
	InvalidPublicKey
	InvalidRedirectURL
	QRCodeGenNotSupported
	QRCodeNotGenerated
	EncodingError
)

func (e *GandalfError) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.Message, e.Code)
}


func NewConnect(config Config) (*Connect, error) {
	if config.PublicKey == "" || config.RedirectURL == "" {
		return nil, fmt.Errorf("invalid parameters")
	}

	if config.Platform == "" {
		config.Platform = PlatformTypeIOS
	}
	return &Connect{PublicKey: config.PublicKey, RedirectURL: config.RedirectURL, Data: config.Data, Platform: config.Platform}, nil
}

func (c *Connect) GenerateURL() (string, error) {
	services, err := runValidation(c.PublicKey, c.RedirectURL, c.Data, c.VerificationStatus)
	if err != nil {
		return "", err
	}

	servicesJSON := servicesToJSON(services)

	url, err := c.encodeComponents(string(servicesJSON), c.RedirectURL, c.PublicKey)
	if err != nil {
		return "", &GandalfError{
			Message: "Encoding Error",
			Code: EncodingError,
		}
	}
	return url, nil
}

func (c *Connect) GenerateQRCode() (string, error) {
	if c.Data == nil {
		return "", &GandalfError{
			Message: "Invalid input parameters",
			Code:    QRCodeGenNotSupported,
		}
	}

	services, err := runValidation(c.PublicKey, c.RedirectURL, c.Data, c.VerificationStatus)
	if err != nil {
		return "", err
	}

	servicesJSON := servicesToJSON(services)
	appClipURL, err := c.encodeComponents(string(servicesJSON), c.RedirectURL, c.PublicKey)
	if err != nil {
		return "", &GandalfError{
			Message: "Encoding Error",
			Code: EncodingError,
		}
	}

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
	req := graphqlClient.NewRequest(constants.IntrospectionQuery)

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

func validateInputData(input InputData) (InputData, error) {
	services := getSupportedServices()
	
	cleanServices := make(InputData)
	unsupportedServices := []string{}

	keys := make([]string, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}

	if len(keys) > 1 {
		return nil, &GandalfError{
			Message: "Only one service is supported per Connect URL",
			Code:    InvalidService,
		}
	}

	for _, key := range keys {
		lowerKey := strings.ToUpper(key)
		if !contains(services, lowerKey) {
			unsupportedServices = append(unsupportedServices, key)
			continue
		}

		value := input[key]
		switch v := value.(type) {
		case bool:
			if !v {
				return nil, &GandalfError{
					Message: "At least one service has to be required",
					Code:    InvalidService,
				}
			}
			cleanServices[lowerKey] = v
		case Service:
			if err := validateInputService(v); err != nil {
				return nil, err
			}
			cleanServices[lowerKey] = v
		default:
			return nil, &GandalfError{
				Message: fmt.Sprintf("Unsupported value type for key %s", key),
				Code:    InvalidService,
			}
		}
	}

	if len(unsupportedServices) > 0 {
		return nil, &GandalfError{
			Message: fmt.Sprintf("These services %s are unsupported", strings.Join(unsupportedServices, " ")),
			Code:    InvalidService,
		}
	}

	return cleanServices, nil
}

func contains(slice []Value, item string) bool {
	for _, v := range slice {
		if v.Name == item {
			return true
		}
	}
	return false
}

func validateInputService(input Service) error {
	if (len(input.Activities) < 1) && (len(input.Traits) < 1) {
		return &GandalfError{
			Message: "At least one trait or activity is required",
			Code:    InvalidService,
		}
	}
	return nil
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
		log.Printf("Error making publicKey request query: %v", err)
		return false
	}

	responseData, ok := graphqlResponse["getAppByPublicKey"].(map[string]interface{})
	if !ok {
		log.Printf("Unexpected response structure: %v", graphqlResponse)
		return false
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

func runValidation(publicKey string, redirectURL string, input InputData, verificationStatus bool) (InputData, error) {
	if !verificationStatus {
		isPublicKeyValid := validatePublicKey(publicKey)
		if !isPublicKeyValid {
			return nil, &GandalfError{
				Message: "Invalid public key",
				Code:    InvalidPublicKey,
			}
		}

		err := validateRedirectURL(redirectURL)
		if err != nil {
			return nil, err
		}

		services, err := validateInputData(input)
		if err != nil {
			return nil, err
		}
		return services, nil
	}
	return nil, nil
}


func (c *Connect) encodeComponents(data, redirectUrl string, publicKey string) (string, error) {
	var baseURL string
	switch c.Platform {
	case PlatformTypeAndroid:
		baseURL = ANDROID_APP_CLIP_BASE_URL
	case PlatformUniversal:
		baseURL = UNIVERSAL_APP_CLIP_BASE_URL
	default:
		baseURL = IOS_APP_CLIP_BASE_URL
	}

	base64Data := base64.StdEncoding.EncodeToString([]byte(data))

	encodedServices := url.QueryEscape(string(base64Data))
	encodedRedirectURL := url.QueryEscape(redirectUrl)
	encodedPublicKey := url.QueryEscape(publicKey)

	return fmt.Sprintf("%s?data=%s&redirectUrl=%s&publicKey=%s", baseURL, encodedServices, encodedRedirectURL, encodedPublicKey), nil
}

func servicesToJSON(services InputData) []byte {
	servicesJSON, err := json.Marshal(services)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	return servicesJSON
}
