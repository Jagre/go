package main

import (
	"fmt"
	"github.com/aaa/go/lib/myJson"
)

func main() {
	// d := myJson.JSONModel{Data: []byte(`{"code":0,"message":"",
	// 	"result":"/success.html?ticket=ST-35072-5jQz4UWTANeDRNzoORO5-sso01-example-org",
	// 	"haha": [{"name": "j", "age": 30}, {"name": "s", "age": 20}],
	// 	"hehe": {"sex": "frame", "tall": 180},
	// 	"xixi": [{"point": [{"x": 1, "y": "2"}, {"x": 2, "y": "3"}], "direct": "east"}]
	// }`)}
	//v, e := d.Get("result")
	//v, e := d.Get("hehe.sex")
	//v, e := d.Get("haha.(1).name")
	//v, e := d.Get("xixi.(0).point.(1).y")
	//v, e := d.Get("xixi.(0).direct")
	//v, e := d.Get("result")
	//v, e := d.Get("(1).x") //nil

	//Now can't support
	d2 := myJson.JSONModel{Data: []byte(`
		[{"x": "1", "y": "1"}, {"x": "2", "y": [{"Z1":"21", "Z2": "221"}, {"Z1":"22", "Z2": "222"}, {"Z1":"23", "Z2": "223"}]}, {"x": "3", "y": "3"}]
		`)}
	v, e := d2.Get("(1).x") //1
	//v, e := d2.Get("x") //raise error
	//v, e := d2.Get("(-1).x") //1
	//v, e := d2.Get("(1).y.(1).Z2") //222
	//v, e := d2.Get("(1).y.(1)") //map[Z1:22 Z2:222]
	//v, e := d2.Get("(1).y.(-1).Z1") //21
	if e != nil {
		fmt.Printf("%v", e)
		return
	}
	//v2 := v.(string)
	fmt.Printf("%v", v)
}
