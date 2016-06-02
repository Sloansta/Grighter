package main

import(
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

func main() {
	fileToWrite := os.Args
	if fileToWrite[1] == "create" {
		CreateTextFile(fileToWrite[2])
	}else if fileToWrite[1] == "ow" {
		OpenAndWrite(fileToWrite[2])
	}
}

//Creates a new file of which the user can write to
func CreateTextFile(file string) {
	fi, err  := os.Create(file)
	if err != nil {
		panic(err)
	}
	

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Start Writing your file (Hitting <ENTER> will write to the file):")
	text, _ := reader.ReadString('\n')
	
	fi.WriteString(text)
	fi.Close()

	fmt.Println("Successfully created, and wrote to:", file)
}

//Opens the file, and shows the contents, as well as allows the user to write to the file
func OpenAndWrite(file string) {
	fi, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("\n", string(fi))

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Start Writing your file (Hitting <ENTER> will write to the file):")
	text, _ := reader.ReadString('\n')

	fo, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	
	finalText := string(fi) + "\n" + text
	fo.WriteString(finalText)
	fo.Close()

	fmt.Println("Successfully wrote to file:", file)
}
