package main

import (
	"github.com/erkylima/golab/pkg/webapi"
)

func main() {
	webapi.ListenAndServe()
}

/*
func init() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("assets/golab-af2f4-firebase-adminsdk-87lbh-50001a2dee.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}
*/
