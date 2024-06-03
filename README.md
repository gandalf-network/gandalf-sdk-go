# gandalf-sdk-go
Gandalf eyeofsauron and connect library


## EyeOfSauron 
eyeofsauron completely abstracts away the complexity of authentication and interacting with the GraphQL APIs.

### Installation

```bash

go get github.com/gandalf-network/gandalf-sdk-go/eyeofsauron

```

### Usage

```bash

go get github.com/gandalf-network/gandalf-sdk-go/eyeofsauron -f ./example/generated

```

#### Flags

- -f, --folder [folder]: Set the destination folder for the generated files


### Using the Generated Files

Once you have successfully generated the necessary files and installed the required dependencies using eyeofsauron, you can proceed to use these files to interact with the API.

#### Initialization

```go
package main

import (
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
)

func main() {
	eye, err := generated.NewEyeOfSauron()
	if err != nil {
		log.Fatalf("failed to run gandalf client: %s", err)
	}
}

```

#### Get Activity

```go
// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
)

func getActivity() {
    // initiatization
    // ....

    // get activity
    response, err := eye.GetActivity(
		context.Background(),
		"MY_DATA_KEY",,
		generated.SourceNetflix,
		10,
		1,
	)
	if err != nil {
		log.Fatalf("failed to get activity: %s", err)
	}

    for _, activity := range response.GetGetActivity().Data {
		switch meta := metadata.(type) {
        case *generated.GetActivityActivityResponseDataActivityMetadataAmazonActivityMetadata:
            fmt.Println("Amazon Activity Metadata:")
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
// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
)

func loopActivity() {
    // initiatization
    // ....

    // lookup activity
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


func printJSON(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("failed to marshal metadata: %v", err)
		return
	}
	fmt.Println(string(jsonData))
}

```
## Connect

connect is a library in go that makes it easier to generate valid Connect URLs that lets your users to link their accounts to Gandalf.