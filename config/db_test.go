package config

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInMemoryDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to in-memory SQLite: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("sql.DB error: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Ping failed: %v", err)
	}
}
