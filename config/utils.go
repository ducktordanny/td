package config

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

/** Generates a 7 character ID. */
func GenerateUniqueSha() string {
	uuidStr := uuid.New().String()
	hash := sha1.New()
	hash.Write([]byte(uuidStr))

	hashed := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashed)

	return hashStr[:7]
}

/** Returns a `LocalModel` from based on path. Creates one if there isn't one with the given path. */
func GetLocalModelByPath(conf *TdConfig, path string) *LocalModel {
	localListIndex := slices.IndexFunc[[]LocalModel](conf.Locals, func(local LocalModel) bool {
		return local.Path == path
	})

	// Create a `LocalModel` if there is none with the given path.
	if localListIndex == -1 {
		localModel := LocalModel{
			Path:  path,
			Items: []TodoModel{},
		}
		conf.Locals = append(conf.Locals, localModel)
		localListIndex = len(conf.Locals) - 1
	}

	return &conf.Locals[localListIndex]
}
