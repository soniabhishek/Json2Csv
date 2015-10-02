package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
)

func Json2csv(fields []string, jfile string) {
	// for _, v := range fields {
	// 	fmt.Println(v)
	// }
	a := make([]map[string]interface{}, 0)

	data, err := ioutil.ReadFile(jfile)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(data))

	err = json.Unmarshal(data, &a)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(a)
	f, err := os.Create("./people.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)

	// fmt.Println(a)
	record := make([]string, 0)
	for _, v := range fields {
		record = append(record, v)
	}

	if err := w.Write(record); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	for _, v := range a {
		record = make([]string, 0)
		for _, v2 := range fields {
			if value, ok := v[v2]; ok == true {
				switch reflect.TypeOf(value).Name() {
				case "string":
					record = append(record, value.(string))

				case "float64":
					record = append(record, strconv.Itoa(int(value.(float64))))

				}

			}

		}
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

var (
	inputFile   = flag.String("i", "", "/path/to/input.json (optional; default is stdin)")
	outputFile  = flag.String("o", "", "/path/to/output.json (optional; default is stdout)")
	outputDelim = flag.String("d", ",", "delimiter used for output values")
	verbose     = flag.Bool("v", false, "verbose output (to stderr)")
	showVersion = flag.Bool("version", false, "print version string")
	printHeader = flag.Bool("p", false, "prints header to output")
	//keys        = StringArray{}
)

func main() {
	fields := []string{"name", "value", "role"}
	//m := make(map[string]interface{}, 0)
	Json2csv(fields, "./jsonfile.json")
	flag.Parse()
	if inputFile != nil {
		fmt.Println(*inputFile)
	}
}
