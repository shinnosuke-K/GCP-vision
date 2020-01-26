package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

func main() {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client %v", err)
	}

	filename := "file Path"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	texts, err := client.DetectTexts(ctx, image, nil, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, text := range texts {
		if strings.Contains(text.GetDescription(), "豚") {
			fmt.Println("NG")
			break
		}
		//fmt.Println(text.GetDescription())
	}
}
