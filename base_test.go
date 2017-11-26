package gonfig

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type testSource struct {
	loadCount int
}

func (s *testSource) Load() {
	s.loadCount++
}

func (s *testSource) Get(key string) (string, bool) {
	return "", false
}

func randomString(strlen int) string {
	rand.Seed(time.Now().UnixNano())

	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func randomFile() (*os.File, error) {
	name := randomString(10)
	path := filepath.Join(os.TempDir(), name)

	return os.Create(path)
}
