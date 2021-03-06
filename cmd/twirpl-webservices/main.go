package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rynop/twirpl/pkg/blogserver"
	"github.com/rynop/twirpl/pkg/imageserver"
	"github.com/rynop/twirpl/rpc/publicservices"
)

func main() {
	//Dump all env vars
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	// You can use any mux you like
	mux := http.NewServeMux()

	//&blogserver.Server{} implements Blog interface
	blogHandler := publicservices.NewBlogServer(&blogserver.Server{}, nil)
	// The generated code includes a const, <ServiceName>PathPrefix, which
	// can be used to mount your service on a mux.
	mux.Handle(publicservices.BlogPathPrefix, blogHandler)

	imageHandler := publicservices.NewImageServer(&imageserver.Server{}, nil)
	mux.Handle(publicservices.ImagePathPrefix, imageHandler)

	appStage, _ := os.LookupEnv("APP_STAGE")
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong. stage:"+appStage)
		fmt.Fprintln(w, r.Header)
	})

	// Un-comment below to test locally
	listenPort, exists := os.LookupEnv("LISTEN_PORT")
	if !exists {
		listenPort = "8080"
	}

	log.Print("Listening on " + listenPort + " in stage " + appStage + " docker image: --CodeImage--")
	//Comment out below if serving over lambda
	http.ListenAndServe(":"+listenPort, mux)

	// Un-comment below before deploying to Lambda
	// log.Fatal(gateway.ListenAndServe("", mux))
}
