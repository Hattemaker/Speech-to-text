package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := speech.NewClient(ctx, option.WithCredentialsFile("/Users/magnusneergaard/Documents/Skole/My-First-Project-89ab7741a22f.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name of the audio file to transcribe.
	filename := "/Users/magnusneergaard/Documents/Skole/audio-file.flac"

	// Reads the audio file into memory.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Detects speech in the audio file.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_FLAC,
			SampleRateHertz: 44100,
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}

	// Prints the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
}