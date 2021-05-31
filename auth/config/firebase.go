package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var app *firebase.App

func GetFirebaseApp() *firebase.App {
	var err error
	if app == nil {
		opt := option.WithCredentialsFile("firebase-adminsdk.json")
		config := &firebase.Config{ProjectID: "cartransplant"}
		app, err = firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
	}
	return app
}
func GetFirebaseAuthClient() *auth.Client {
	client, err := GetFirebaseApp().Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing firebase auth client: %v\n", err)
	}
	return client
}
