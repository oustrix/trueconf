package json

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"refactoring/internal/entity"
	"strconv"
)

var (
	errorNotFound = errors.New("user not found")
)

type UsersStore struct {
	Increment int             `json:"increment"`
	List      entity.UserList `json:"list"`
	path      string
}

func (s *UsersStore) save() {
	b, _ := json.Marshal(s)
	_ = os.WriteFile(s.path, b, fs.ModePerm)
}

func openUsersStore(path string) *UsersStore {
	f, _ := os.ReadFile(path)
	s := &UsersStore{}
	_ = json.Unmarshal(f, s)

	s.path = path

	return s
}

type UsersRepository struct {
	storePath string
}

func NewUsersRepository(storePath string) *UsersRepository {
	store := openUsersStore(storePath)
	if store.List == nil {
		store.List = make(entity.UserList)
		store.save()
	}

	return &UsersRepository{
		storePath: storePath,
	}
}

func (r *UsersRepository) GetUsers() *entity.UserList {
	s := openUsersStore(r.storePath)
	defer s.save()

	return &s.List
}

func (r *UsersRepository) CreateUser(user *entity.User) string {
	s := openUsersStore(r.storePath)
	defer s.save()

	s.Increment++
	id := strconv.Itoa(s.Increment)
	s.List[id] = *user

	return id
}

func (r *UsersRepository) GetUser(id string) (entity.User, error) {
	s := openUsersStore(r.storePath)
	defer s.save()

	if _, ok := s.List[id]; !ok {
		return entity.User{}, errorNotFound
	}

	return s.List[id], nil
}

func (r *UsersRepository) UpdateUser(id string, user *entity.User) error {
	s := openUsersStore(r.storePath)
	defer s.save()

	if _, ok := s.List[id]; !ok {
		return errorNotFound
	}

	s.List[id] = *user

	return nil
}

func (r *UsersRepository) DeleteUser(id string) error {
	s := openUsersStore(r.storePath)
	defer s.save()

	if _, ok := s.List[id]; !ok {
		return errorNotFound
	}

	delete(s.List, id)

	return nil
}
