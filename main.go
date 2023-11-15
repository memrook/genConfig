package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	var tmplFile, listFile, variable string

	flag.StringVar(&tmplFile, "t", "template.json", "name of the file containing the template")
	flag.StringVar(&listFile, "l", "list.txt", "name of the file containing the variable list")
	flag.StringVar(&variable, "v", "%variable%", "pattern for replacement")
	flag.Parse()

	template, err := os.ReadFile(tmplFile)
	if err != nil {
		log.Fatal("can't open the template file: ", err)
	}

	list, err := os.ReadFile(listFile)
	if err != nil {
		log.Fatal("can't open the list file: ", err)
	}

	vars := strings.Split(strings.Trim(string(list), "\n"), "\n")

	var res string

	for i, v := range vars {
		prefix := ",\n"
		if i == 0 {
			prefix = ""
		}
		res += prefix + strings.ReplaceAll(string(template), variable, strings.Trim(v, " "))
	}

	log.Println(res)

	resFile, err := os.Create("output.json")
	defer resFile.Close()

	_, err = resFile.WriteString(res)
	if err != nil {
		log.Fatal("can't write a file: ", err)
	}
}
