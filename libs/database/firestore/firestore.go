package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/billykore/todolist/libs/config"
	"github.com/billykore/todolist/libs/pkg/log"
	"github.com/billykore/todolist/libs/pkg/path"
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
