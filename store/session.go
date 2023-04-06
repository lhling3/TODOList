package store

// SessionMap 直接在内存中使用Map存放用户会话的容器
type SessionMap struct {
	data map[string]*User
}

func NewSessionMap() *SessionMap {
	return &SessionMap{
		data: make(map[string]*User),
	}
}

func (s *SessionMap) Put(sessionId string, user *User) {
	s.data[sessionId] = user
}

func (s *SessionMap) Get(sessionId string) *User {
	return s.data[sessionId]
}

func (s *SessionMap) Del(sessionId string) {
	delete(s.data, sessionId)
}

func (s *SessionMap) Exist(sessionId string) bool {
	return s.Get(sessionId) != nil
}
