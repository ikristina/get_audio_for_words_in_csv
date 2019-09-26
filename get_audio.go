package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/watson-developer-cloud/go-sdk/texttospeechv1"
)

const AUDIO_DIR_PATH string = "audio"
const CSV_FILE string = "words.csv"

func GetAPIKey() string {
	apiKey, exists := os.LookupEnv("TEXT_TO_SPEECH_IAM_APIKEY")
	if exists == false {
		log.Fatal("TEXT_TO_SPEECH_IAM_APIKEY could not be found in ibm-credentials.env")
	}
	return apiKey
}

func GetEndpoint() string {
	url, exists := os.LookupEnv("TEXT_TO_SPEECH_URL")
	if exists == false {
		log.Fatal("TEXT_TO_SPEECH_URL could not be found in ibm-credentials.env")
	}
	return url
}

func GetAndSaveAudio(text string) {
	textToSpeech, textToSpeechErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
		URL: GetEndpoint(),
		Authenticator: &core.IamAuthenticator{
			ApiKey: GetAPIKey(), // <--- found in the docs somewhere like Apikey, another mistake
		},
		// IAMApiKey: "{apiKey}",  <--- this is from official docs, doesn't work
	})

	if textToSpeechErr != nil {
		panic(textToSpeechErr)
	}

	response, responseErr := textToSpeech.Synthesize(
		&texttospeechv1.SynthesizeOptions{
			Text:   core.StringPtr(text),
			Accept: core.StringPtr("audio/mp3"),
			Voice:  core.StringPtr("fr-FR_ReneeV3Voice"),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	result := textToSpeech.GetSynthesizeResult(response)
	if result != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(result)

		pwd, _ := os.Getwd()

		fileName := filepath.Join(pwd, "audio", text+".mp3")
		fmt.Println(fileName)
		file, _ := os.Create(fileName)
		file.Write(buff.Bytes())
		file.Close()
	}
}

func ReadCsvFile(filePath string) {
	// Load a csv file.
	f, _ := os.Open(filePath)

	// Create a new reader.
	r := csv.NewReader(f)
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		// Display record.
		// ... Display record length.
		fmt.Println(record)
		fmt.Println(len(record))

		// Display first element of each row

		GetAndSaveAudio(record[0])
		fmt.Println(record[0])

		// Display all elements of a row (record)
		// for value := range record {
		//     fmt.Printf("  %v\n", record[value])
		// }
	}
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load("ibm-credentials.env"); err != nil {
		log.Fatal("No .env file found")
	}

}

func main() {
	// if _, err := os.Stat(AUDIO_DIR_PATH); os.IsNotExist(err) {
	//     os.Mkdir(AUDIO_DIR_PATH, os.ModeDir)
	// }

	// Create a directory. Ignore any issues raised in case the dir exists
	_ = os.Mkdir(AUDIO_DIR_PATH, 0700)

	ReadCsvFile(CSV_FILE)

}
