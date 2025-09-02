package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Print("Enter the name of the text file: ")
	fmt.Scanln(&filename)

	var sli []Name

	f, err := os.Open(filename)
	if err == nil {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			name := strings.Split(line, " ")

			if len(name[0]) > 20 {
				name[0] = name[0][:20]
			}
			if len(name[1]) > 20 {
				name[1] = name[1][:20]
			}

			sli = append(sli, Name{fname: name[0], lname: name[1]})
		}

		for name := range sli {
			fmt.Println(sli[name].fname, sli[name].lname)
		}

		f.Close()
	}

}
