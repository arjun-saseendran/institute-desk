package session

import "gorm.io/gorm"

type SessionService interface {
	CreateSession(sessionData *InputCreateSession) (*Session, error)
	GetSessions() ([]Session, error)
	GetSession(id string) (*Session, error)
	UpdateSession(id string, sessionData *InputUpdateSession) (*Session, error)
	DeleteSession(id string) error
}

type sessionService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) SessionService {
	return &sessionService{db: db}
}

func (ss *sessionService) CreateSession(sessionData *InputCreateSession) (*Session, error) {
	newSession := &Session{Title: sessionData.Title, StartTime: sessionData.StartTime, EndTime: sessionData.EndTime}
	result := ss.db.Create(newSession)
	if result.Error != nil {
		return nil, result.Error
	}
	return newSession, nil
}

func (ss *sessionService) GetSessions() ([]Session, error) {
	var sessions []Session
	result := ss.db.Find(&sessions)
	if result.Error != nil {
		return nil, result.Error
	}
	return sessions, nil
}

func (ss *sessionService) GetSession(id string) (*Session, error) {
	singleSession := NewSession()
	result := ss.db.First(&singleSession, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return singleSession, nil
}

func (ss *sessionService) UpdateSession(id string, sessionData *InputUpdateSession) (*Session, error) {
	updateSession := NewSession()

	result := ss.db.First(updateSession, id)
	if result.Error != nil {
		return nil, result.Error
	}
	ss.db.Model(&updateSession).Updates(Session{Title: sessionData.Title, StartTime: sessionData.StartTime, EndTime: sessionData.EndTime})
	return updateSession, nil
}

func (ss *sessionService) DeleteSession(id string) error {
	deleteSession := NewSession()
	result := ss.db.First(deleteSession, id)
	if result.Error != nil {
		return result.Error
	}
	ss.db.Delete(deleteSession)
	return nil
}
