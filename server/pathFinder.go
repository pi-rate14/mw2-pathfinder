package main

const START_NODE = "Player"

type Node struct {
	Source string
	Type   string
	Level  int
}

type PathFinder struct {
	objects set
	graph   map[string][]*Node
}

func NewPathFinder(graph map[string][]*Node, list set) *PathFinder {
	return &PathFinder{
		objects: list,
		graph:   graph,
	}
}

func (response *Response) getPath(startNode string, nodeType string, endNode string, vis map[string]bool, level int) bool {

	// fmt.Printf("START NODE: %s\n", startNode)
	// fmt.Printf("END NODE: %s\n", endNode)
	vis[startNode] = true
	n := &Node{
		Source: startNode,
		Level:  level,
		Type:   nodeType,
	}
	response.Path = append(response.Path, n)

	// for _, key := range response.Path {
	// 	fmt.Println(key.Level, key.Source, key.Type)
	// }

	if startNode == endNode {
		return true
	}

	for _, node := range response.pathFinder.graph[startNode] {
		adjNode := node.Source
		level := node.Level
		nodeType := node.Type
		if !vis[adjNode] {
			// fmt.Printf("VISITING: %s\n", adjNode)
			if response.getPath(adjNode, nodeType, endNode, vis, level) {
				return true
			}
		}
	}

	response.Path = response.Path[:len(response.Path)-1]
	return false
}
