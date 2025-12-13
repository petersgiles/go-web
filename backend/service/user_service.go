package service

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/pete/go-web/graph/model"
)

// UserService handles user profile operations
type UserService struct {
	users     []*model.User
	usersById map[string]*model.User
	mu        sync.RWMutex
}

// NewUserService creates a new user service and loads users from JSON
func NewUserService() (*UserService, error) {
	svc := &UserService{
		usersById: make(map[string]*model.User),
	}

	if err := svc.loadUsers(); err != nil {
		return nil, err
	}

	return svc, nil
}

// loadUsers reads user data from JSON file
func (s *UserService) loadUsers() error {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = filepath.Join("..", "data")
	}

	dataPath := filepath.Join(dataDir, "users.json")
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return err
	}

	var users []*model.User
	if err := json.Unmarshal(data, &users); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.users = users
	s.usersById = make(map[string]*model.User)
	for _, user := range users {
		s.usersById[user.ID] = user
	}

	return nil
}

// GetUserByID returns a user by their ID
func (s *UserService) GetUserByID(id string) (*model.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, found := s.usersById[id]
	return user, found
}

// GetUserByEmail returns a user by their email
func (s *UserService) GetUserByEmail(email string) (*model.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		if user.Email == email {
			return user, true
		}
	}

	return nil, false
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() []*model.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.users
}
