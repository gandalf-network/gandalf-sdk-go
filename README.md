# Gandalf SDK for Go

The Gandalf SDK for Go provides two main libraries: `EyeOfSauron` for interacting with GraphQL APIs and `Connect` for generating valid Connect URLs for account linking.

## EyeOfSauron

`EyeOfSauron` completely abstracts away the complexity of authentication and interacting with Gandalf GraphQL APIs.

### Installation

To install the `EyeOfSauron` package, use the following command:

```bash
go get github.com/gandalf-network/gandalf-sdk-go/eyeofsauron
```

### Usage

To generate the necessary files, use the following command:

```bash
go run github.com/gandalf-network/gandalf-sdk-go/eyeofsauron -f ./example/generated
```

#### Flags

- `-f, --folder [folder]`: Set the destination folder for the generated files.

### Using the Generated Files

Once you have successfully generated the necessary files and installed the required dependencies using `EyeOfSauron`, you can proceed to use these files to interact with the API.

#### Initialization

```go
package main

import (
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
)

func main() {
	eye, err := generated.NewEyeOfSauron("<YOUR_GANDALF_PRIVATE_KEY")
	if err != nil {
		log.Fatalf("failed to run gandalf client: %s", err)
	}
}
```

#### Get Activity

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
)

func main() {
    // Initialization
    eye, err := generated.NewEyeOfSauron("<YOUR_GANDALF_PRIVATE_KEY")
	if err != nil {
		log.Fatalf("failed to initialize gandalf client: %s", err)
	}

    // Get activity
    response, err := eye.GetActivity(
		context.Background(),
		"MY_DATA_KEY",
		generated.SourceNetflix,
		10,
		1,
	)
	if err != nil {
		log.Fatalf("failed to get activity: %s", err)
	}

    for _, activity := range response.GetGetActivity().Data {
		switch meta := activity.Metadata.(type) {
        case *generated.GetActivityActivityResponseDataActivityMetadataAmazonActivityMetadata:
            fmt.Println("Amazon Activity Metadata:")
            printJSON(meta.AmazonActivityMetadata)
        case *generated.GetActivityActivityResponseDataActivityMetadataInstacartActivityMetadata:
            fmt.Println("Instacart Activity Metadata:")
            printJSON(meta.InstacartActivityMetadata)
        case *generated.GetActivityActivityResponseDataActivityMetadataNetflixActivityMetadata:
            fmt.Println("Netflix Activity Metadata:")
            printJSON(meta.NetflixActivityMetadata)
        case *generated.GetActivityActivityResponseDataActivityMetadataPlaystationActivityMetadata:
            fmt.Println("Playstation Activity Metadata:")
            printJSON(meta.PlaystationActivityMetadata)
        case *generated.GetActivityActivityResponseDataActivityMetadataUberActivityMetadata:
            fmt.Println("Uber Activity Metadata:")
            printJSON(meta.UberActivityMetadata)
        case *generated.GetActivityActivityResponseDataActivityMetadataYoutubeActivityMetadata:
            fmt.Println("YouTube Activity Metadata:")
            printJSON(meta.YoutubeActivityMetadata)
        default:
            log.Printf("Unknown metadata type: %T\n", meta)
        }
	}
}

func printJSON(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("failed to marshal metadata: %v", err)
		return
	}
	fmt.Println(string(jsonData))
}
```

#### Lookup Activity

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
)

func main() {
    // Initialization
    eye, err := generated.NewEyeOfSauron("<YOUR_GANDALF_PRIVATE_KEY")
	if err != nil {
		log.Fatalf("failed to initialize gandalf client: %s", err)
	}

    // Lookup activity
    response, err := eye.LookupActivity(
		context.Background(),
		"MY_DATA_KEY",
		"ACTIVITY_ID",
	)
	if err != nil {
		log.Fatalf("failed to lookup activity: %s", err)
	}

    metadata := response.LookupActivity.GetMetadata()
    switch meta := metadata.(type) {
	case *generated.LookupActivityMetadataAmazonActivityMetadata:
		fmt.Println("Amazon Activity Metadata:")
		printJSON(meta.AmazonActivityMetadata)
	case *generated.LookupActivityMetadataInstacartActivityMetadata:
		fmt.Println("Instacart Activity Metadata:")
		printJSON(meta.InstacartActivityMetadata)
	case *generated.LookupActivityMetadataNetflixActivityMetadata:
		fmt.Println("Netflix Activity Metadata:")
		printJSON(meta.NetflixActivityMetadata)
	case *generated.LookupActivityMetadataPlaystationActivityMetadata:
		fmt.Println("Playstation Activity Metadata:")
		printJSON(meta.PlaystationActivityMetadata)
	case *generated.LookupActivityMetadataUberActivityMetadata:
		fmt.Println("Uber Activity Metadata:")
		printJSON(meta.UberActivityMetadata)
	case *generated.LookupActivityMetadataYoutubeActivityMetadata:
		fmt.Println("YouTube Activity Metadata:")
		printJSON(meta.YoutubeActivityMetadata)
	default:
		log.Printf("Unknown metadata type: %T\n", meta)
	}
}

func printJSON(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("failed to marshal metadata: %v", err)
		return
	}
	fmt.Println(string(jsonData))
}
```


### Get Traits
```go
func main() {
	// Initialization
    eye, err := generated.NewEyeOfSauron("8c48ad0e5892d51d8e2e411a77a1d73ebe764b619c846d1cab3dc45ee172e8ca")
	if err != nil {
		log.Fatalf("failed to run gandalf client: %s", err)
	}

	response, err := eye.GetTraits(context.Background(), "3pLT1hCieyPQQb876i24D34Qf8y6Yyke5m4rhPRhV67D", generated.SourceNetflix, []generated.TraitLabel{generated.TraitLabelPlan})
	if err != nil {
		log.Fatalf("failed to get traits: %s", err)
	}

	fmt.Println("Get Traits", response.GetGetTraits())
}
```

### Lookup Traits
```go
func main() {
	eye, err := generated.NewEyeOfSauron("8c48ad0e5892d51d8e2e411a77a1d73ebe764b619c846d1cab3dc45ee172e8ca")
	if err != nil {
		log.Fatalf("failed to run gandalf client: %s", err)
	}

	traitID, err := uuid.Parse("e55bf3a6-66a5-4902-b7f2-34e352b65d52")
	if err != nil {
		log.Fatalf("failed to parse string to uuid")
	}
	response, err := eye.LookupTrait(context.Background(), "3pLT1hCieyPQQb876i24D34Qf8y6Yyke5m4rhPRhV67D", traitID)
	if err != nil {
		log.Fatalf("failed to lookup trait: %s", err)
	}

	fmt.Println("Lookup Trait", response.GetLookupTrait())
}
```


## Connect

`Connect` is a library in Go that makes it easier to generate valid Connect URLs that let your users link their accounts to Gandalf. To use this library, follow the installation and usage instructions provided in the documentation.


### Connect installation

To install the `Connect` package, use the following command:

```bash
go get github.com/gandalf-network/gandalf-sdk-go/connect
```

```go
const publicKey = "0x036518f1c7a10fc77f835becc0aca9916c54505f771c82d87dd5943bb01ba5ca08";
const redirectURL = "https://example.com"


func main() {
	// Define the input data
	services := connect.InputData{
		"netflix": connect.Service{
			Traits:     []string{"rating"},
			Activities: []string{"watch"},
		},
	}

	// Define the config parameters
	config := connect.Config{
		PublicKey:   publicKey,
		RedirectURL: redirectURL,
		Data:    	 services,
	}
	
	// Initialization
	conn, err := connect.NewConnect(config)
	if err != nil {
		log.Fatalf("An error occurred with initializing connect: %v", err)
	}

	// Call the GenerateURL method
	url, err := conn.GenerateURL()
	if err != nil {
		log.Fatalf("An error occurred generating url: %v", err)
	}
	fmt.Println("URL => ", url)
	
	
	// Call the GenerateQRCode method
	qrCode, err := conn.GenerateQRCode()
	if err != nil {
		log.Fatalf("An error occurred generating QR Code url: %v", err)
	}
	fmt.Println("Base64 QR Code => ", qrCode)
}
```

#### Generate URL for Android
```go
func main() {
	// Define the input data
	services := connect.InputData{
		"netflix": connect.Service{
			Traits:     []string{"rating"},
			Activities: []string{"watch"},
		},
	}

	// Define the config parameters
	config := connect.Config{
		PublicKey:   publicKey,
		RedirectURL: redirectURL,
		Data:    	 services,
		Platform: 	 PlatformTypeAndroid,
	}

	// Call the GenerateURL method for Android
	androidUrl, err := conn.GenerateURL()
	if err != nil {
		log.Fatalf("An error occurred generating url: %v", err)
	}
	fmt.Println("URL => ", androidUrl)
}

```