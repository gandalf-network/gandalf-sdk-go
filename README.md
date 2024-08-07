<h1 align="center">Gandalf SDK for Go</h1>

<p align="center">
  <img src="https://mintlify.s3-us-west-1.amazonaws.com/gandalf/logo/gandalf.svg" alt="gandalf-logo" width="30px" height="30px"/>
  <br>
  The Gandalf SDK for Go provides two main libraries: <code>EyeOfSauron</code> for interacting with GraphQL APIs and <code>Connect</code> for generating valid Connect URLs for account linking.
  <br>
</p>


## Documentation

Visit the official Gandalf documentation [here](https://docs.gandalf.network/).

- [Introduction](https://docs.gandalf.network/get-started/introduction)
- [Getting Started](https://docs.gandalf.network/get-started/quickstart)
- [Examples](https://docs.gandalf.network/examples/whoami.tv)

## Directory

| Package                                                                                                    | Description                                                                                                            | Links                                                           |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------- |
| [`EyeofSauron`](https://github.com/gandalf-network/gandalf-sdk-go/tree/main/eyeofsauron)                        | A command line tool for generating client code for [Gandalf eyeofsauron ](https://docs.gandalf.network/concepts/sauron).                                | [Docs](https://docs.gandalf.network/concepts/sauron)     |
| [`Connect`](https://github.com/gandalf-network/gandalf-sdk-go/tree/main/connect) | A package for generating connect URL for Gandalf                                                     | [Docs](https://docs.gandalf.network/concepts/connect/intro) |
                                                               

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
go run github.com/gandalf-network/gandalf-sdk-go/eyeofsauron -f generated
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

	"github.com/gandalf-network/gandalf-sdk-go/connect"
	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/generated"
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

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/generated"
)

func main() {
	// initialize eyeofsauron object
	...

    // Get activity
    response, err := eye.GetActivity(
		context.Background(),
		"MY_DATA_KEY",
		[]generated.ActivityType{generated.ActivityTypeWatch},
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

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/generated"
)

func main() {
	// initialize eyeofsauron object
	...

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
	// initialize eyeofsauron object
	...

	response, err := eye.GetTraits(context.Background(), "MY_DATA_KEY", generated.SourceNetflix, []generated.TraitLabel{generated.TraitLabelPlan})
	if err != nil {
		log.Fatalf("failed to get traits: %s", err)
	}

	fmt.Println("Get Traits", response.GetGetTraits())
}
```

### Lookup Traits
```go
func main() {
	// initialize eyeofsauron object
	...
	
	traitID, err := uuid.Parse("MY_TRAIT_ID")
	if err != nil {
		log.Fatalf("failed to parse string to uuid")
	}
	response, err := eye.LookupTrait(context.Background(), "MY_DATA_KEY", traitID)
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
import (
	...
	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/generated"
)

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
		Platform: 	connect.PlatformTypeAndroid,
	}

	// Call the GenerateURL method for Android
	androidUrl, err := conn.GenerateURL()
	if err != nil {
		log.Fatalf("An error occurred generating url: %v", err)
	}
	fmt.Println("URL => ", androidUrl)
}

```