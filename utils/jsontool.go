package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int8
}

//ToJson 对生成json
func ToJson() []byte {
	user := User{
		Name: "Tab",
		Age:  18,
	}
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

//FromJson 读取json
func FromJson(data []byte) {
	//var user User
	var a []*string
	result := make(map[string]interface{})
	err := json.Unmarshal(data, &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func testRead() []byte {
	fp, err := os.OpenFile("./info.json", os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data[:n]))
	return data[:n]
}

func testWrite(data []byte) {
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func Testmain() {
	var data []byte
	/*data = testMarshal()
	fmt.Println(string(data))
	testWrite(data)
	*/
	data = testRead()
	FromJson(data)
}
