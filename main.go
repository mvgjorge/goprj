package main

import (
	"github.com/dimfeld/httptreemux"
	"github.com/mvgjorge/goprj/api"
	"github.com/mvgjorge/goprj/music"
	"log"
	"net/http"
)

func main() {

	m := &music.Song{
		Id:     1,
		Name:   "somewhere over the rainbow",
		Singer: "Israel kamakawiwoole",
	}

	log.Printf(m.Name, m.Id)

	addr := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()
	router.Handler(http.MethodGet, "/music/:id", &api.GetMusicHandler{})
	router.Handler(http.MethodPost, "/music/:id", &api.CreateMusicHandler{})
	router.Handler(http.MethodPut, "/music/:id", &api.UpdateMusicHandler{})
	router.Handler(http.MethodDelete, "/music/:id", &api.DeleteMusicHandler{})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))

	// execute
	// curl http://localhost:8081/cars/gol
	// curl -XPUT http://localhost:8081/cars/fusca -d'{"name": 1}'
}
