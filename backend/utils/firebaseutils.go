package utils

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func FirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SERVICE")))
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, err
}
