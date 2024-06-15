package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
	"github.com/google/uuid"
)

func main() {
	eye, err := generated.NewEyeOfSauron("0x366a5ae7c7575f8cb0e3832ad53e668061e0ad800b94ffb75fd5b6d241a83e56")
	if err != nil {
		log.Fatalf("failed to run gandalf client: %s", err)
	}

	response, err := eye.GetActivity(
		context.Background(),
		"BG7u85FMLGnYnUv2ZsFTAXrGT2Xw3TikrBHm2kYz31qq",
		[]generated.ActivityType{generated.ActivityTypePlay},
		generated.SourceNetflix,
		100,
		1,
	)

	if err != nil {
		log.Fatalf("failed to run gandalf client: %s", err)
	}

	for _, activity := range response.GetGetActivity().Data {
		activityID, _ := uuid.Parse(activity.Id)
		response, err := eye.LookupActivity(
			context.Background(),
			"BG7u85FMLGnYnUv2ZsFTAXrGT2Xw3TikrBHm2kYz31qq",
			activityID,
		)

		if err != nil {
			log.Fatalf("unable to look up activity: %s", err)
		}
		printLookupActivityMetadata(response.LookupActivity.GetMetadata())
	}
}

func printGetActivityMetadata(metadata interface{}) {
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

func printLookupActivityMetadata(metadata interface{}) {
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
