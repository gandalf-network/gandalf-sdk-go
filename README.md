# gandalf-sdk-go
The Gandalf SDK for Go provides two main libraries: `EyeOfSauron` for interacting with GraphQL APIs and `Connect` for generating valid Connect URLs for account linking.


## EyeOfSauron 
`EyeOfSauron` completely abstracts away the complexity of authentication and interacting with GraphQL APIs.

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
	eye, err := generated.NewEyeOfSauron("<YOUR_GANDALF_PRIVATE_KEY>")
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

func getActivity() {
    // Initialization
    eye, err := generated.NewEyeOfSauron("<YOUR_GANDALF_PRIVATE_KEY>")
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
// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
)

func lookupActivity() {
    // Initialization
    eye, err := generated.NewEyeOfSauron("<YOUR_GANDALF_PRIVATE_KEY>")
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
## Connect

# Gandalf SDK for Go

The Gandalf SDK for Go provides two main libraries: `EyeOfSauron` for interacting with GraphQL APIs and `Connect` for generating valid Connect URLs for account linking.

## EyeOfSauron

`EyeOfSauron` completely abstracts away the complexity of authentication and interacting with GraphQL APIs.

### Installation

To install the `EyeOfSauron` package, use the following command:

```bash
go get github.com/gandalf-network/gandalf-sdk-go/eyeofsauron
```

### Usage

To generate the necessary files, use the following command:

```bash
go get github.com/gandalf-network/gandalf-sdk-go/eyeofsauron -f ./example/generated
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
	eye, err := generated.NewEyeOfSauron()
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

func getActivity() {
    // Initialization
    eye, err := generated.NewEyeOfSauron()
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

func lookupActivity() {
    // Initialization
    eye, err := generated.NewEyeOfSauron()
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

## Connect

`Connect` is a library in Go that makes it easier to generate valid Connect URLs that let your users link their accounts to Gandalf. To use this library, follow the installation and usage instructions provided in the documentation.