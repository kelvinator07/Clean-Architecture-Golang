package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/kelvinator05/clean-architecture-go/internal/entity"
	"github.com/kelvinator05/clean-architecture-go/internal/usecase"
)

var (
	UserRe              = regexp.MustCompile(`^/users/*$`)
	UserReWithIDOrEmail = regexp.MustCompile(`^/users/([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}|[a-zA-Z0-9-]+)$`)
	EmailRegex          = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type HTTPServer struct {
	UserUseCase usecase.UserUseCase
}

func NewHTTPServer(userUseCase usecase.UserUseCase) *HTTPServer {
	return &HTTPServer{UserUseCase: userUseCase}
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path =", r.URL.Path, "Method =", r.Method)
	switch {
	case r.Method == http.MethodPost && UserRe.MatchString(r.URL.Path):
		s.CreateUserHandler(w, r)
		return
	case r.Method == http.MethodGet && UserRe.MatchString(r.URL.Path):
		s.GetAllUsersHandler(w, r)
		return
	case r.Method == http.MethodGet && UserReWithIDOrEmail.MatchString(r.URL.Path):
		s.GetUserHandler(w, r)
		return
	default:
		return
	}
}

func (s *HTTPServer) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := s.UserUseCase.CreateUser(req.Name, req.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (s *HTTPServer) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the resource ID/slug using a regex
	matches := UserReWithIDOrEmail.FindStringSubmatch(r.URL.Path)
	// Expect matches to be length >= 2 (full string + 1 matching group)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	var user *entity.User
	var err error

	idOrEmail := matches[1]

	if EmailRegex.MatchString(idOrEmail) {
		user, err = s.UserUseCase.GetUserByEmail(idOrEmail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		id, err := strconv.Atoi(idOrEmail)
		if err != nil {
			InternalServerErrorHandler(w, r)
			return
		}
		user, err = s.UserUseCase.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (s *HTTPServer) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := s.UserUseCase.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
