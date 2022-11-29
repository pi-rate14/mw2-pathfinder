package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func makeGraph(filePath string) (map[string][]*Node, set, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	graph := map[string][]*Node{}

	list := NewSet()

	csvReader := csv.NewReader(f)

	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return graph, *list, err
		}

		r := &Node{}
		r.Source = row[2]

		r.Type = row[3]
		r.Level, _ = strconv.Atoi(row[4])

		if !list.Contains(row[2]) {
			list.Add(*r)
		}

		graph[row[0]] = append(graph[row[0]], r)
	}

}

func (response *Response) buildResponse(endNode string) []ResponseBody {
	var responseBody []ResponseBody

	numResponses := len(response.Path)

	for idx := 0; idx < numResponses-1; idx++ {
		var r ResponseBody
		r.Object = response.Path[idx].Source
		r.Type = response.Path[idx].Type
		r.Level = response.Path[idx+1].Level
		responseBody = append(responseBody, r)
	}

	finalOb := ResponseBody{
		Object: endNode,
		Level:  1,
		Type:   response.pathFinder.objects.m[endNode].Type,
	}

	responseBody = append(responseBody, finalOb)

	return responseBody
}
