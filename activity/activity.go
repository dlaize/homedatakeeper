package activity

import (
	"database/sql"
	"errors"
)

type activity struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

func (a *activity) getActivity(db *sql.DB) error {
	return db.QueryRow("SELECT name, value, unit FROM activities WHERE id=$1",
		a.ID).Scan(&a.Name, &a.Value, &a.Unit)
}

func getListActivities(db *sql.DB, start, count int) ([]activity, error) {
	rows, err := db.Query(
		"SELECT id, name, value, unit FROM activities LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	activities := []activity{}

	for rows.Next() {
		var a activity
		if err := rows.Scan(&a.ID, &a.Name, &a.Value, &a.Unit); err != nil {
			return nil, err
		}
		activities = append(activities, a)
	}

	return activities, nil
}

func (a *activity) updateActivity(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (a *activity) deleteActivity(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (a *activity) createActivity(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO activities(name, value, unit) VALUES($1, $2, $3) RETURNING id",
		a.Name, a.Value, a.Unit).Scan(&a.ID)

	if err != nil {
		return err
	}

	return nil
}
