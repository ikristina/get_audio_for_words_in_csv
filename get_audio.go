package main

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"encoding/csv"
	"io"
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

func ReadCsvFile(filePath string)  {
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

	apiKey := GetAPIKey()
	fmt.Println(apiKey)
	url := GetEndpoint()
	fmt.Println(url)

	// if _, err := os.Stat(AUDIO_DIR_PATH); os.IsNotExist(err) {
	//     os.Mkdir(AUDIO_DIR_PATH, os.ModeDir)
	// }

	// Create a directory. Ignore any issues raised in case the dir exists
	_ = os.Mkdir(AUDIO_DIR_PATH, os.ModeDir)

	ReadCsvFile(CSV_FILE)

}
