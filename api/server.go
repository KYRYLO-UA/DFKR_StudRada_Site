package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudTicket struct {
	ID int `json:"id"`
}

type Student struct {
	// ID uuid.UUID `json:"id"`
	ID int `json:"id"`

	Ticket StudTicket `json:"ticket"`

	Firstname  string `json:"firstname"`
	Secondname string `json:"secondname"`
	Middlename string `json:"middlename"`
}

type Server struct {
	*mux.Router
	studentsList []Student
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		studentsList: []Student{
			{
				ID:         1,
				Ticket:     StudTicket{ID: 1},
				Firstname:  "Кирило",
				Secondname: "Біцай",
				Middlename: "Борисович",
			},
		},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	// s.HandleFunc("/students", s.getStudentList()).Methods("GET")
	// s.HandleFunc("/students/{id}", s.getStudentList()).Methods("GET")
	// s.HandleFunc("/students/{id}/delete", s.deleteStudent()).Methods("POST")
	// s.HandleFunc("/students/{id}", s.deleteStudent()).Methods("DELETE")
	// s.HandleFunc("/register", s.registerStudent()).Methods("POST")

	s.HandleFunc("/students", s.getStudentList()).Methods("GET")
	s.HandleFunc("/students", s.registerStudent()).Methods("POST")
	s.HandleFunc("/students/{id}", s.deleteStudent()).Methods("DELETE")
}

func (s *Server) registerStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Student
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Register")

		i.ID = 2 //len(s.studentsList) + 1 //uuid.New()

		s.studentsList = append(s.studentsList, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getStudentList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.studentsList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) deleteStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr) //uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, student := range s.studentsList {
			if student.ID == id {
				s.studentsList = append(s.studentsList[:i], s.studentsList[i+1:]...)
				break
			}
		}
	}
}
