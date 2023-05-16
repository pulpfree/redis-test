package rdb

import (
	"testing"

	"github.com/joho/godotenv"
)

const (
	envFp = "../.env"
	stage = "development"
)

func TestNewDB(t *testing.T) {
	godotenv.Load(envFp)
	rdb, err := New(stage)
	if err != nil {
		t.Fatalf("Expected null error received: %s", err)
	}
	if rdb == nil {
		t.Fatalf("Expected redis to not be nil")
	}

	/* err = rdb.Ping()
	if err != nil {
		t.Fatalf("Expected null error received: %s", err)
	}

	defer func() {
		rdb.Close()
	}() */
}
