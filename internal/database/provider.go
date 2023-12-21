package database

import (
	"github.com/billykore/todolist/internal/database/firestore"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(firestore.New)
