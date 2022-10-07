package main

import (
	"context"
	"encoding/json"

	"aws_lambda/internal/sns"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

func handler(ctx context.Context, snsEvent events.SNSEvent) (err error) {
	for _, message := range snsEvent.Records {
		log.Infof("received sns event `%+v`", message)
		event := &sns.SnsMessage{}
		err = json.Unmarshal([]byte(message.SNS.Message), event)
		if err != nil {
			log.Errorf("unable to parse sns payload, err: %s", err.Error())
			return
		}
		log.Printf("%+v", event)
	}
	return
}

func main() {
	lambda.Start(handler)
}
