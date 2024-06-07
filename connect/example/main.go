package main

import (
	"fmt"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/connect"
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
	
	conn, err := connect.NewConnect(publicKey, redirectURL, connect.PlatformTypeIOS, services)
	if err != nil {
		log.Fatalf("An error occurred with initializing connect: %v", err)
	}


	url, err := conn.GenerateURL()
	if err != nil {
		log.Fatalf("An error occurred generating url: %v", err)
	}
	fmt.Println("URL => ", url)
	
	
	qrCodeURL, err := conn.GenerateQRCode(services)
	if err != nil {
		log.Fatalf("An error occurred generating QR Code url: %v", err)
	}
	fmt.Println("QR Code URL => ", qrCodeURL)


}