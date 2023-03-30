package firebase

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/smalake/kakebo-api/utils/logging"
	"google.golang.org/api/option"
)

var Auth *auth.Client

// Firebase SDKの初期化処理
func InitFirebase() {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_FILE"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		os.Exit(1)
	}
	Auth, err = app.Auth(context.Background())
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		os.Exit(1)
	}
}
