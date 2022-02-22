package repository

import (
	"errors"
	"log"
	"microservice/internal/domain"
	"time"
)

type PostgresDatabaseSessionRepository struct {
	PostgresDatabase
}

func (repo *PostgresDatabaseUserRepository) CreateSession(userID int) (int, error) {

	user, err := repo.GetUserFromID(userID)
	session := domain.Session{ID: -1, UserID: user.ID, CreationTime: time.Now()}

	db, err := repo.GetConnection()
	if err != nil {
		return 0, err
	}

	err = db.QueryRow("INSERT INTO sessions(user_ID, creation_time) VALUES($1, $2) RETURNING ID", session.UserID, session.CreationTime).Scan(&session.ID)
	if err != nil {
		log.Printf("Error: %v", err)

		return 0, err
	}

	return session.ID, nil
}

func (repo *PostgresDatabaseUserRepository) GetSession(userID int) (domain.Session, error) {
	session := domain.Session{}

	db, err := repo.GetConnection()
	if err != nil {
		return session, err
	}

	err = db.Get(&session, "SELECT * FROM sessions WHERE user_ID = $1", userID)
	if err != nil {
		return domain.Session{}, errors.New("Failed to retrieve session")
	}

	return session, nil
}

func (repo *PostgresDatabaseUserRepository) DeleteSession(sessionID int) error {
	db, err := repo.GetConnection()
	if err != nil {
		return err
	}

	res, err := db.Exec("DELETE FROM sessions WHERE ID = $1", sessionID)
	if err != nil {
		return err
	}

	numDeletedSessions, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if numDeletedSessions == 0 {
		return errors.New("Session does not exist")
	}

	return nil
}
