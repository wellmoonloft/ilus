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

func testMarshal() []byte {
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

func testUnmarshal(data []byte) {
	var user Settings
	err := json.Unmarshal(data, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func testRead() []byte {
	fp, err := os.OpenFile("./info.json", os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 10000)
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
	data = testMarshal()
	fmt.Println(string(data))
	testWrite(data)
	data = testRead()
	testUnmarshal(data)
}
