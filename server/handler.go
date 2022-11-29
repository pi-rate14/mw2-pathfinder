package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"sort"
// )

type Response struct {
	pathFinder PathFinder
	Path       []*Node
}

func NewResponse(pf PathFinder) *Response {
	return &Response{
		Path:       []*Node{},
		pathFinder: pf,
	}
}

type ResponseBody struct {
	Object string `json:"object"`
	Level  int    `json:"level"`
	Type   string `json:"type"`
}

func (pathFinder *PathFinder) findPath(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	response := NewResponse(*pathFinder)

	endNode := r.URL.Query().Get("endNode")

	log.Default().Printf("RECEIVED REQUEST FOR: %s", endNode)

	vis := make(map[string]bool)

	if response.getPath("Player", "You", endNode, vis, 1) {

		// for _, key := range response.Path {
		// 	fmt.Println(key.Level, key.Source, key.Type)
		// }

		log.Default().Print("---FOUND PATH---")

		resp := response.buildResponse(endNode)

		for _, node := range resp {
			fmt.Printf("%s : %s : %d -> ", node.Object, node.Type, node.Level)
		}

		fmt.Printf("\n\n")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	} else {
		fmt.Println("noob")
	}

}

func (pathFinder *PathFinder) PrintAdjList(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var resp []string

	for k := range pathFinder.objects.m {
		resp = append(resp, k)
	}

	sort.Strings(resp)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}
