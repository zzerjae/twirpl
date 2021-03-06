package main_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/rynop/twirpl/rpc/publicservices"
	"github.com/twitchtv/twirp"
)

func TestHandler(t *testing.T) {
	//URL for APIG
	urlPrefix := "https://YOURAPIG.execute-api.us-east-1.amazonaws.com/prod/"

	//URL for local testing
	// urlPrefix := "http://localhost:8080"

	client := publicservices.NewImageProtobufClient(urlPrefix, &http.Client{})

	//Attach the required accept header for APIG binary support
	header := make(http.Header)
	header.Set("Accept", "application/protobuf")
	ctx := context.Background()
	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		log.Printf("twirp error setting headers: %s", err)
		return
	}

	resp, err := client.CreateGiphy(ctx, &publicservices.Search{Term: "wahooo"})
	if err == nil {
		fmt.Println(resp)
	} else {
		if twerr, ok := err.(twirp.Error); ok {
			switch twerr.Code() {
			case twirp.InvalidArgument:
				fmt.Println("Oh no " + twerr.Error())
			default:
				fmt.Println(twerr.Error())
			}
		}
		return
	}
}
