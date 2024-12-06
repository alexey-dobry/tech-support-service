package session

import (
	"sync"
)

var managerSessions = sync.Map{}

type ManagerSession struct {
	Authorized   bool  // Флаг авторизации
	ActiveClient int64 // ID активного клиента
}

func GetSession(managerID int64) *ManagerSession {
	session, exists := managerSessions.Load(managerID)
	if exists {
		return session.(*ManagerSession)
	}

	newSession := &ManagerSession{
		Authorized:   false,
		ActiveClient: 0,
	}
	managerSessions.Store(managerID, newSession)
	return newSession
}

func SetActiveClient(managerID int64, clientID int64) {
	session := GetSession(managerID)
	session.ActiveClient = clientID
}

func AuthorizeManager(managerID int64) {
	session := GetSession(managerID)
	session.Authorized = true
}

func IsAuthorized(managerID int64) bool {
	session := GetSession(managerID)
	return session.Authorized
}
