package session

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

var sessionMap sync.Map

func init() {
	sessionMap = sync.Map{}
}

// creat
func GenerateNewSession() {
	id, _ := uuid.NewV4

	//ss := &
}

// delete

// load

// is expire
