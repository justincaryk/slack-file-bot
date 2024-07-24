package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"5E_Hardcore_Mode.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.UploadFileV2Parameters{
			Channel:  channelArr[0],
			File:     fileArr[i],
			Filename: fileArr[i],
			FileSize: 116,
		}

		fmt.Printf("Uploading file %s", fileArr[i])

		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Title: %s, ID: %s\n", file.Title, file.ID)
	}
}
