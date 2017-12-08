/*
Package myJson Reademe :
1. how to package ?
cd gopath (cd f:\go)
go install packageDirectory (go install jagre/myJson)*/
package myJson

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// JSONModel is entity about json
type JSONModel struct {
	Data    []byte      //source
	jsonObj interface{} //it's json object that's created by unmarshal
}

// Get a field's value from json string
// the "fieldPath" simple:
//
func (myjson *JSONModel) Get(fieldPath string) (interface{}, error) {
	if len(fieldPath) == 0 {
		panic("The parameter \"fieldPath\" is nil or empty")
	}
	if e := myjson.convertToJSONObj(); e != nil {
		return nil, e
	}
	if myjson.jsonObj == nil {
		return nil, errors.New("Parse json to the type \"interface{}\" is nil ")
	}

	// if jsonobj, ok := myjson.jsonObj.(map[string]interface{}); ok {
	// 	return parse1(jsonobj, fieldPath)
	// }
	//parse json's type: {} | [{}]
	//type assert and type convert

	switch myjson.jsonObj.(type) {
	case map[string]interface{}:
		return parse1(myjson.jsonObj.(map[string]interface{}), fieldPath)
	case []interface{}:
		return parse2(myjson.jsonObj.([]interface{}), fieldPath)
	default:
		return nil, errors.New("Mybe your json data's pattern is not right, try to use standand json format")
	}

}

// parse2 type []interface{} like this structure: [{}]
func parse2(jsonObj []interface{}, fieldPath string) (interface{}, error) {
	//parse field's path
	fields := strings.Split(fieldPath, ".")
	f1 := fields[0]
	if !(strings.HasPrefix(f1, "(") && strings.HasSuffix(f1, ")")) {
		f1 = "(0)"
		//fieldPath needn't update
	} else {
		fieldPath = fieldPath[len(f1)+1:]
	}
	//get source str index
	si := f1[1 : len(f1)-1]
	//get source index
	i, e := strconv.Atoi(si)
	if e != nil {
		i = 0
	}

	//pattern {}: {"x": 1, "y": 2, "z": {"m": 1, "n": 2}}
	var normalJsonobj map[string]interface{}
	if i <= -1 || i >= len(jsonObj) {
		//last one
		temp := jsonObj[len(jsonObj)-1]
		normalJsonobj = temp.(map[string]interface{})
	} else {
		//specified one
		temp := jsonObj[i]
		normalJsonobj = temp.(map[string]interface{})
	}

	return parse1(normalJsonobj, fieldPath)
}

//parse1 type map[string]interface{} like this structure: {}
func parse1(jsonObj map[string]interface{}, fieldPath string) (interface{}, error) {
	//parse field's path
	fields := strings.Split(fieldPath, ".")

	for i := 0; i < len(fields); i++ {
		jdata := jsonObj[fields[i]]
		if jdata == nil {
			break
		}
		if _, ok := jdata.(map[string]interface{}); ok {
			//{"x": 1, "y": 2, "z": {"m": 1, "n": 2}}
			if i != len(fields)-1 {
				jsonObj = jdata.(map[string]interface{})
			} else {
				return jdata, nil
			}
		} else if _, ok := jdata.([]interface{}); ok {
			//[{"x": 1, "y": 1}, {"x": 2, "y": 2}, {"x": 3, "y": 3}]
			i++ //go to next fieldName : (index)
			si := fields[i]
			temps := jdata.([]interface{})
			if strings.HasPrefix(si, "(") && strings.HasSuffix(si, ")") {
				sii := si[1 : len(si)-1]
				ii, e := strconv.Atoi(sii)
				if e != nil {
					ii = 0
				}
				if ii <= -1 || ii >= len(temps) {
					//last one
					temp := temps[len(temps)-1]
					jsonObj = temp.(map[string]interface{})
				} else {
					//specified one
					temp := temps[ii]
					jsonObj = temp.(map[string]interface{})
				}
			}
		} else {
			return jdata, nil
		}
	}
	return nil, errors.New("Not found \"" + fieldPath + "\"")
}

//will parse every key-value and save to map[string]interface{}
func (myjson *JSONModel) convertToJSONObj() error {
	if len(myjson.Data) == 0 {
		return errors.New("Json content is nil or empty")
	}
	if myjson.jsonObj == nil {
		return json.Unmarshal(myjson.Data, &myjson.jsonObj)
	}
	return nil
}

func (myjson *JSONModel) isValidate() error {
	if len(myjson.Data) == 0 {
		return errors.New("Json content is nil or empty")
	}
	return nil
}
