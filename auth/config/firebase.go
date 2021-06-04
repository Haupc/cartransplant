package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var (
	app             *firebase.App
	firestoreClient *firestore.Client
	authClient      *auth.Client
	fcmClient       *messaging.Client
)

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
	var err error
	if authClient == nil {
		authClient, err = GetFirebaseApp().Auth(context.Background())
		if err != nil {
			log.Fatalf("error initializing firebase auth client: %v\n", err)
		}
	}
	return authClient
}

func GetFireStoreClient() *firestore.Client {
	var err error
	if firestoreClient == nil {
		firestoreClient, err = GetFirebaseApp().Firestore(context.Background())
		if err != nil {
			log.Fatalf("error initializing firestore client: %v\n", err)
		}
	}
	return firestoreClient
}

func GetFcmClient() *messaging.Client {
	var err error
	if fcmClient == nil {
		fcmClient, err = GetFirebaseApp().Messaging(context.Background())
		if err != nil {
			log.Fatalf("error initializing fcm client: %v\n", err)
		}
	}
	return fcmClient
}
