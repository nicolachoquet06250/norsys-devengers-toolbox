package router

import (
	"encoding/json"
	"net/http"
	"norsys-devengers-toolbox/api"
	"norsys-devengers-toolbox/arrays"
	"norsys-devengers-toolbox/parser"
	"strconv"
)

func ApiToolsPage(w http.ResponseWriter, r *http.Request) {
	tools := arrays.ArrayMerge(
		categories[0].Tools,
		categories[1].Tools,
		categories[2].Tools,
		categories[3].Tools,
	)
	_tools, err := json.Marshal(tools)

	if r.URL.Query().Get("id") != "" {
		var tool api.Tool
		for item := range _tools {
			var id, _ = strconv.Atoi(r.URL.Query().Get("id"))
			if tools[item].Id == id {
				tool = tools[item]
				break
			}
		}

		_tool, err := json.Marshal(tool)

		if err != nil {
			w.WriteHeader(500)
			parser.Text(&w, "[]")
		} else {
			parser.Text(&w, string(_tool))
		}
		return
	}

	if err != nil {
		w.WriteHeader(500)
		parser.Text(&w, "[]")
	} else {
		parser.Text(&w, string(_tools))
	}
}

func ApiCategoriesPage(w http.ResponseWriter, r *http.Request) {
	_categories, err := json.Marshal(categories)
	if err != nil {
		w.WriteHeader(500)
		parser.Text(&w, "[]")
	} else {
		parser.Text(&w, string(_categories))
	}
}
