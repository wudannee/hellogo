// Xml parsing
package ch2

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Student struct {
	XMLName    xml.Name
	Name       string
	Age        int
	Department string
}

type Mode int

const (
	QUICK_MODE Mode = iota
	NORMAL_MODE
)

// DemoXml parses the student.xml file and prints out the student's information.
func DemoXml(mode Mode) {
	file, err := os.Open("student.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var student Student

	switch mode {
	case QUICK_MODE:
		if err = xml.NewDecoder(file).Decode(&student); err != nil {
			log.Fatalln(err)
		}
	case NORMAL_MODE:
		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		err = xml.Unmarshal(data, &student)
		if err != nil {
			fmt.Println("Error unmarshalling XML:", err)
			return
		}
	default:
		log.Fatalln("Invalid mode to parse xml")
	}

	fmt.Println("Parsed Student Info:")
	fmt.Printf("Name: %s\n", student.Name)
	fmt.Printf("Age: %d\n", student.Age)
	fmt.Printf("Department: %s\n", student.Department)

}
