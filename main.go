package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("gobatis.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := GoBatis{}
	err = xml.Unmarshal(data, &v)
	fmt.Printf("value:%#v\n", v)
	fmt.Println("\n##################\nanalyzing....\n###########################\n")
	AnalyzeSource(v)
}
