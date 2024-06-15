package main

import (
	"os"
	"fmt"
	"log"

	"encoding/base64"
	"github.com/gandalf-network/gandalf-sdk-go/connect"
	"strings"
)

const publicKey = "0x036518f1c7a10fc77f835becc0aca9916c54505f771c82d87dd5943bb01ba5ca08";
const redirectURL = "https://example.com"


func main() {
	services := connect.InputData{
		"netflix": connect.Service{
			Traits:     []string{"rating"},
			Activities: []string{"watch"},
		},
	}

	config := connect.Config{
		PublicKey:   publicKey,
		RedirectURL: redirectURL,
		Data:    	 services,
	}
	
	conn, err := connect.NewConnect(config)
	if err != nil {
		log.Fatalf("An error occurred with initializing connect: %v", err)
	}


	url, err := conn.GenerateURL()
	if err != nil {
		log.Fatalf("An error occurred generating url: %v", err)
	}
	fmt.Println("URL => ", url)
	
	
	qrCode, err := conn.GenerateQRCode()
	if err != nil {
		log.Fatalf("An error occurred generating QR Code url: %v", err)
	}
	fmt.Println("Base64 QR Code => ", qrCode)

	var base64QRCodeWithPrefix string
	if strings.HasPrefix(qrCode, "data:image/png;base64,") {
		base64QRCodeWithPrefix = strings.TrimPrefix(qrCode, "data:image/png;base64,")
	}
	
	decodedQRCode, err := base64.StdEncoding.DecodeString(base64QRCodeWithPrefix)
	if err != nil {
		log.Fatalf("Failed to decode base64 QR code: %v", err)
	}

	outputFile := "decoded_qrcode.png"
	err = os.WriteFile(outputFile, decodedQRCode, 0644)
	if err != nil {
		log.Fatalf("Failed to save decoded QR code as PNG: %v", err)
	}

	fmt.Printf("Decoded QR code saved to %s\n", outputFile)


	// Define the input data
	servicess := connect.InputData{
		"netflix": connect.Service{
			Traits:     []string{"rating"},
			Activities: []string{"watch"},
		},
	}

	// Define the config parameters
	configg := connect.Config{
		PublicKey:   publicKey,
		RedirectURL: redirectURL,
		Data:    	 servicess,
		Platform: 	 connect.PlatformTypeAndroid,
	}

	conn, err = connect.NewConnect(configg)
	if err != nil {
		log.Fatalf("An error occurred with initializing connect: %v", err)
	}

	// Call the GenerateURL method for Android
	androidUrl, err := conn.GenerateURL()
	if err != nil {
		log.Fatalf("An error occurred generating url: %v", err)
	}
	fmt.Println("Android URL => ", androidUrl)
}