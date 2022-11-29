package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	graph, list, _ := makeGraph("map.csv")

	// PrintGraph(graph)

	// PrintList(list)

	pathFinder := NewPathFinder(graph, list)

	// resp := NewResponse(*pathFinder)
	// vis := make(map[string]bool)

	// resp.getPath("Player", "Player", "Minibak", vis, 1)

	// for _, key := range resp.Path {
	// 	fmt.Println(key.Level, key.Source, key.Type)
	// }

	http.HandleFunc("/getAdjList", pathFinder.PrintAdjList)

	http.HandleFunc("/findPath", pathFinder.findPath)

	log.Default().Println("SERVER STARTED ON PORT 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func PrintGraph(graph map[string][]*Node) {
	for key, node := range graph {
		fmt.Printf("KEY: %s\n", key)
		for _, n := range node {
			fmt.Printf("DESTINATION: %s->", n.Source)
			fmt.Printf("TYPE: %s->", n.Type)
			fmt.Printf("LEVEL: %d\n", n.Level)
		}
		fmt.Println()
	}
}

func PrintList(st set) {
	for _, node := range st.m {
		fmt.Println(node.Source, node.Type, node.Level)

	}
	fmt.Println(len(st.m))
}
