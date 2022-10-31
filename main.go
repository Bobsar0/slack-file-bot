package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_OAUTH_TOKEN", "")
	os.Setenv("CHANNEL_ID", "")

	api := slack.New(os.Getenv("SLACK_BOT_OAUTH_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}

	fileArr := []string{"CUX Guide Microsoft.pdf", "THE BAPTISM IN THE HOLY SPIRIT.docx"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URLPrivate)
	}
}
