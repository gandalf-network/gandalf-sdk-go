package main

import (
	"fmt"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/connect"
)

const publicKey = "0x036518f1c7a10fc77f835becc0aca9916c54505f771c82d87dd5943bb01ba5ca08";
const redirectURL = "https://example.com"


func main() {
	qrCodeServices := connect.Services{
		{Name: "Netflix", Status: true},
		{Name: "Playstation", Status: false},
	}
	qrCodeURL, err := connect.GenerateQRCode(publicKey, redirectURL, qrCodeServices)
	if err != nil {
		log.Fatalf("An error occurred generating QR Code url: %v", err)
	}

	fmt.Println("QR Code URL => ", qrCodeURL)

	services := connect.Services{
		{Name: "Netflix", Status: false},
		{Name: "Playstation", Status: true},
	}
	url, err := connect.GenerateURL(publicKey, redirectURL, services)
	if err != nil {
		log.Fatalf("An error occurred url: %v", err)
	}

	fmt.Println("URL => ", url)
}