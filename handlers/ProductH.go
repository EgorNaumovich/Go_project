package handlers

import (
	"awesomeProject2/Go_project/structure"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func ProductsLog(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) Get(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Test")
	pl := structure.GetPList()
	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error with json", http.StatusInternalServerError)
	}
}

func (p *Products) Add(rw http.ResponseWriter, r *http.Request) {
	pr := &structure.Product{}
	err := pr.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Decoding error", http.StatusBadRequest)
	}
	p.l.Printf("Pr: %#v", pr)
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.Get(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.Add(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
