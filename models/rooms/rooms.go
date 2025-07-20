package rooms

import (
	"database/sql"
	"fmt"
)

// Room represents a single record in the rooms table
type Room struct {
	// RoomId is the primary key
	RoomID int    `json:"room_id"`
	Name   string `json:"room_name"`
}

// Add is used to insert a new room to the database
func (r *Room) Add(db *sql.DB) error {
	const query = `
		INSERT INTO rooms ( name )
		VALUES ( ? )
	`

	_, err := db.Exec(query, r.Name)
	if err != nil {
		return fmt.Errorf("failed to add the room: %w", err)
	}

	return nil
}

// Query is used to retrieve information about rooms from the database
func Query(db *sql.DB, limit ...int) ([]Room, error) {
	rows, err := db.Query("SELECT * FROM rooms LIMIT ?", limit)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	rooms := []Room{}
	for rows.Next() {
		var room Room
		if err := rows.Scan(&room.RoomID, &room.Name); err != nil {
			return nil, fmt.Errorf("failed to extract the results: %w", err)
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}
