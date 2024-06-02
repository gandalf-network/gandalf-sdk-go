package main

import (
	"context"
	"log"

	"github.com/gandalf-network/gandalf-sdk-go/eyeofsauron/example/generated"
)

func main() {
	eye, err := generated.NewEyeOfSauron("0x366a5ae7c7575f8cb0e3832ad53e668061e0ad800b94ffb75fd5b6d241a83e56")
	if err != nil {
		log.Fatalf("failed to run gandalf client: %s", err)
	}

	_, err = eye.GetActivity(
		context.Background(),
		"BG7u85FMLGnYnUv2ZsFTAXrGT2Xw3TikrBHm2kYz31qq",
		generated.SourceNetflix,
		100,
		10,
	)
	if err != nil {
		log.Fatalf("failed to run gandalf client: %s", err)
	}
}
