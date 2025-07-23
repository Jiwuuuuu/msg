package tests

import (
	"database/sql"
	"testing"

	"github.com/Jiwuuuuu/msg/models"
	"github.com/Jiwuuuuu/msg/models/messages"
	"github.com/Jiwuuuuu/msg/models/rooms"
	_ "modernc.org/sqlite"
)

// Only for queries
const DEFAULT_LIMIT = 1

// helperSetupDB is used to create a temporary test database
func helperSetupDB(t *testing.T) *sql.DB {
	t.Helper()

	return models.InitDB("sqlite", "file:test.db?mode=memory&cache=shared")
}

func TestMessages(t *testing.T) {
	db := helperSetupDB(t)
	defer db.Close()

	t.Run("Messages: Test inserting", func(t *testing.T) {
		message := messages.Message{
			Content:  "foo bar",
			Checksum: "fbc1a9f858ea9e177916964bd88c3d37b91a1e84412765e29950777f265c4b75",
			Username: "user",
		}

		err := message.Add(db)
		if err != nil {
			t.Fatalf("failed during adding a message: %v", err)
		}
	})

	t.Run("Messages: Test queries", func(t *testing.T) {
		_, err := messages.Query(db, DEFAULT_LIMIT)
		if err != nil {
			t.Fatalf("failed during executing a query: %v", err)
		}
	})
}

func TestRooms(t *testing.T) {
	db := helperSetupDB(t)
	defer db.Close()

	t.Run("Rooms: Test inserting", func(t *testing.T) {
		room := rooms.Room{
			Name: "test room",
		}

		err := room.Add(db)
		if err != nil {
			t.Fatalf("failed during adding a room: %v", err)
		}
	})

	t.Run("Rooms: Test queries", func(t *testing.T) {
		_, err := rooms.Query(db, DEFAULT_LIMIT)
		if err != nil {
			t.Fatalf("failed during executing a query: %v", err)
		}
	})
}
