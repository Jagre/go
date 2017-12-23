/*
Package myJson Reademe :
1. how to package ?
cd gopath (cd f:\go)
go install packageDirectory (go install jagre/myJson)*/
package myJson

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// JSONModel is entity about json
type JSONModel struct {
	Data    []byte      //source
	jsonObj interface{} //it's json object that's created by unmarshal
}

// Get a field's value from json string
// the "nodePath" simple:
func (myjson *JSONModel) Get(nodePath string) (interface{}, error) {
	if len(nodePath) == 0 {
		panic("The parameter \"nodePath\" is nil or empty")
	}
	if e := myjson.toJSONObj(); e != nil {
		return nil, e
	}
	if myjson.jsonObj == nil {
		return nil, errors.New("Parse json to the type \"interface{}\" is nil ")
	}

	nodes := strings.Split(nodePath, ".")
	return adaptor(myjson.jsonObj, nodes)
}

func adaptor(jsonObj interface{}, nodes []string) (interface{}, error) {
	if len(nodes) == 0 {
		return jsonObj, nil
	}

	switch jsonObj.(type) {
	case map[string]interface{}:
		return parse3(jsonObj.(map[string]interface{}), nodes)
	case []interface{}:
		return parse4(jsonObj.([]interface{}), nodes)
	default:
		return jsonObj, nil
	}

}

func parse3(jsonObj map[string]interface{}, nodes []string) (interface{}, error) {
	if len(nodes) == 0 {
		return jsonObj, nil
	}
	node := nodes[0]
	data := jsonObj[node]

	return adaptor(data, nodes[1:])
}

func parse4(jsonObj []interface{}, nodes []string) (interface{}, error) {

	if len(nodes) == 0 {
		return jsonObj, nil
	}
	node := nodes[0]
	if strings.HasPrefix(node, "(") && strings.HasSuffix(node, ")") {
		index, _ := strconv.Atoi(node[1 : len(node)-1])
		if index < 0 {
			index = 0
		} else if index > len(jsonObj) {
			index = len(jsonObj) - 1
		}
		data := jsonObj[index]
		return adaptor(data, nodes[1:])
	}

	return nil, fmt.Errorf("node(\"%s\") is invalid in the nodePath", node)
}

//will parse every key-value and save to map[string]interface{}
func (myjson *JSONModel) toJSONObj() error {
	if len(myjson.Data) == 0 {
		return errors.New("Json content is nil or empty")
	}
	if myjson.jsonObj == nil {
		return json.Unmarshal(myjson.Data, &myjson.jsonObj)
	}
	return nil
}
