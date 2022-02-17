package userService

import (
	"log"
	"microservice/api/common"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func isLoggedIn(db *sqlx.DB, w http.ResponseWriter, r *http.Request, username string) {
	user, err := getUserFromName(db, username)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Failed to retrieve login stauts")
		return
	}

	session, err := getSession(db, user.ID)

	// TODO: Check if session expired

	if err != nil {
		log.Printf("Error: %v", err)

		common.TryWriteResponse(w, "User not logged in")
	} else {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "User logged in with id "+strconv.FormatInt(int64(session.ID), 10))
	}
}

func IsLoginHandler(w http.ResponseWriter, r *http.Request) {
}
