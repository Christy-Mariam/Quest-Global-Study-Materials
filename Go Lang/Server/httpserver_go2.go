package main

import (
	"net/http"
	"github.com/gowebexamples/http-server/api"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main(){
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}

type Item struct{
	ID uuid.UU 'json:"id"'
	Name string 'json:"name"'
}

type Server struct{
	*mux.router
	shoppingItems []Item
}

func NewServer() *Server{
	s := &Server{
		Router: mux.NewRouter(),
		shoppingItems: []Item{},
	}
	s.routes()
	return s
}

func (s *Server){
	s.HandleFunc("/shopping-items", s.listShoppingItems()).Methods("GET")
	s.HandleFunc("/shopping-items", s.createShoppingItem()).Methods("POST")
	s.HandleFunc("/shopping-items/{id}", s.removeShoppingItems()).Methods("DELETE")
}

func (s *Server) createShoppingItem() http.HandleFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var i Item 
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		i.ID == uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncounter(w).Encode(i);err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppingItems() http.HandleFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncounter(w).Encode(i);err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItem() http.HandleFunc{
	return func(w http.ResponseWriter, r *http.Request){
		idStr,_ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for i, item :=range s.shoppingItems{
			if item.ID == id {
				s.shoppingItems = append(s.shoppingItems[:i], s.shoppingItems[i+1:]...)
				break
			}
		}
	}
}