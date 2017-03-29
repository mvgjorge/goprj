package api

import (
	"fmt"
	"github.com/dimfeld/httptreemux"
	//	"log"
	"net/http"
)

type CreateMusicHandler struct{}

func (h *CreateMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Criar registro: %s!", params["id"])
}

type DeleteMusicHandler struct{}

func (h *DeleteMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Delete registro: %s!", params["id"])
}

type UpdateMusicHandler struct{}

func (h *UpdateMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Update registro registro: %s!", params["id"])
}

type GetMusicHandler struct{}

func (h *GetMusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Get registro registro: %s!", params["id"])
}

//
//
// type UpsertCarHandler struct{}
//
// func (h *UpsertCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	params := httptreemux.ContextParams(r.Context())
// 	fmt.Fprintf(w, "Eu deveria criar um carro chamado: %s!", params["id"])
// 	fmt.Fprintln(w, "Não crio por que sou mal!")
// }
//
// type GetCarHandler struct{}
//
// func (h *GetCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	params := httptreemux.ContextParams(r.Context())
// 	fmt.Fprintf(w, "Eu deveria busca um carro chamado: %s!", params["id"])
// 	fmt.Fprintln(w, "Não busco por que estou com preguiça!")
// }
