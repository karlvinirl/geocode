package models

import (
	"fmt"
	"time"

	"example.com/go_practice/restapi/db"
)

type Event struct {
	ID          int64
	NAME        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e Event) Register(userId int64) error {
	query := "INSERT into registrations(event_id, user_id) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) DeleteRegistration(userId int64) err  {
	query := "DELETE from registrations WHERE event_id = ? and user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.NAME, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func DeleteEvent(eventId int64) error {

	query := "DELETE FROM events where ID = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(eventId)

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * from events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.NAME, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (event Event) Update() error {
	query :=
		`UPDATE events SET name = ? , description = ? , location = ? , dateTime = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.NAME, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.NAME, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Returning an event from GetEvent by id %v", id)
	return &event, nil
}
