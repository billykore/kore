package db

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/billykore/kore/pkg/config"
	"github.com/billykore/kore/pkg/log"
	"github.com/billykore/kore/pkg/path"
	"google.golang.org/api/option"
)

func New(appCfg *config.Config) *firestore.Client {
	ctx := context.Background()
	logger := log.NewLogger()

	opt := option.WithCredentialsFile(path.GetAbsolutePath(appCfg.Firestore.SDKFile))
	conf := &firebase.Config{ProjectID: appCfg.Firestore.ProjectId}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		logger.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	return client
}
