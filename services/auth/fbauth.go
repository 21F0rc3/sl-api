package fbauth

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
)

var (
	firebaseConfigFile = os.Getenv("FIREBASE_CONFIG_FILE")
)

var firebaseClient *auth.Client

func InitAuth() error {
	log.Println(firebaseConfigFile)

	opt := option.WithCredentialsFile(firebaseConfigFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return errors.Wrap(err, "error initializing firebase auth")
	}

	client, errAuth := app.Auth(context.Background())
	if errAuth != nil {
		return errors.Wrap(errAuth, "error initializing firebase auth (creating client)")
	}

	firebaseClient = client
	return nil
}
