package messages

import (
	"database/sql"
	"fmt"
)

// Message represents a single record in the messages table
type Message struct {
	// MessageID is the primary key in the messages table
	MessageID int `json:"message_id"`

	// RoomID is used to link messages with rooms
	RoomID int `json:"room_id"`

	// Timestamp stores the UNIX time
	// of when the message was sent
	Timestamp int `json:"timestamp"`

	// Content stores the AES encrypted message
	Content string `json:"content"`

	// Checksum stores the SHA256 hash of the content
	Checksum string `json:"checksum"`
	Username string `json:"username"`
}

// Add is used to insert a new message into the database
func (m *Message) Add(db *sql.DB) error {
	const query = `
		INSERT INTO messages (
			RoomID, Timestamp, Content, Checksum, Username
		)

		VALUES (
			?, ?, ?, ?, ?
		);
	`

	_, err := db.Exec(
		query, m.RoomID, m.Timestamp, m.Content, m.Checksum, m.Username,
	)

	if err != nil {
		return fmt.Errorf("failed to add the message: %w", err)
	}

	return nil
}

// Query is used to retrieve all of the messages from the database
func Query(db *sql.DB) ([]Message, error) {
	rows, err := db.Query("SELECT * FROM messages")
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	messages := []Message{}
	for rows.Next() {
		var msg Message
		err := rows.Scan(
			&msg.MessageID, &msg.RoomID, &msg.Timestamp, &msg.Content, &msg.Checksum, &msg.Username,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to extract the results: %w", err)
		}

		messages = append(messages, msg)
	}

	return messages, nil
}
