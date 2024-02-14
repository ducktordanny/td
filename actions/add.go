package actions

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/ducktordanny/td/config"
	"github.com/ducktordanny/td/flags"
	"github.com/google/uuid"
)

func generateUniqueSha() string {
	uuidStr := uuid.New().String()
	hash := sha1.New()
	hash.Write([]byte(uuidStr))

	hashed := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashed)

	return hashStr[:7]
}

func Add(scope flags.Scope, value *string) {
	fmt.Println("add", *value)
	todo := config.TodoModel{
		Id:       generateUniqueSha(),
		Content:  *value,
		Resolved: false,
	}
	fmt.Println(todo, scope)
}
