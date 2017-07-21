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
	Data  []byte      //source
	JData interface{} //it was setted after unmarshal
}

//Get a field's value from json string
func (myjson *JSONModel) Get(fieldName string) (interface{}, error) {
	if len(fieldName) == 0 {
		panic("The parameter \"fieldName\" is nil or empty")
	}

	e := myjson.Parse()
	if e != nil {
		return nil, e
	}
	if myjson.JData == nil {
		return nil, errors.New("Parse json to the type \"interface{}\" is nil ")
	}

	//parse json's type: [{}] | {}
	// var jd interface{}
	// switch myjson.JData.(type) {
	// case map[string]interface{}: //{}
	// 	jd = myjson.JData.(map[string]interface{})
	// case []interface{}: //[{}]
	// 	jd = myjson.JData.([]interface{})
	// }
	jd := myjson.JData.(map[string]interface{})
	//parse field's path
	fns := strings.Split(fieldName, ".")

	for i := 0; i < len(fns); i++ {
		jdata := jd[fns[i]]
		if jdata == nil {
			break
		}
		if _, ok := jdata.(map[string]interface{}); ok {
			//{"x": 1, "y": 2, "z": {"m": 1, "n": 2}}
			if i != len(fns)-1 {
				jd = jdata.(map[string]interface{})
			} else {
				return jdata, nil
			}
		} else if _, ok := jdata.([]interface{}); ok {
			//[{"x": 1, "y": 1}, {"x": 2, "y": 2}, {"x": 3, "y": 3}]
			i++ //go to next fieldName : (index)
			si := fns[i]
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
					jd = temp.(map[string]interface{})
				} else {
					//specified one
					temp := temps[ii]
					jd = temp.(map[string]interface{})
				}
			}
		} else {
			return jdata, nil
		}
	}

	// for k, v := range jd {
	// 	if k == fieldName {
	// 		return v, nil
	// 	}
	// }
	return nil, errors.New("Not found \"" + fieldName + "\"")
}

//Parse I wonna parse every key-value and save to map[string]interface{}
func (myjson *JSONModel) Parse() error {
	if ok, e := myjson.isValidate(); !ok {
		return e
	}
	if myjson.JData == nil {
		return json.Unmarshal(myjson.Data, &myjson.JData)
	}
	return nil
}

func (myjson *JSONModel) isValidate() (bool, error) {
	if len(myjson.Data) == 0 {
		return false, errors.New("Json content is nil or empty")
	}
	return true, nil
}
