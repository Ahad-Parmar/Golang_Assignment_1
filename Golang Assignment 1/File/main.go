package main 

import ( 
	"fmt"
	"io/ioutil"
	"log"
	"os"
) 

func CreateFile() { 

	fmt.Printf("wassup file? \n") 

	file, error := os.Create("wassup.txt") 
	
	if error != nil { 
		log.Fatalf("failed: %s", error) 
	} 

	defer file.Close() 
	
	len, error := file.WriteString("Welcome wassup"+ 
			" how are you? "+ " hope you are doing well!!") 

	if error != nil { 
		log.Fatalf("failed: %s", error) 
	} 
	fmt.Printf("\nName of the file: %s", file.Name()) 
	fmt.Printf("\nLength of the file: %d bytes", len) 
} 

func ReadFile() { 

	fmt.Printf("\n\nreading chat\n") 
	fileName := "chat.txt"
	
	data, error := ioutil.ReadFile("chat.txt") 
	if error != nil { 
		log.Panicf("failed: %s", error) 
	} 
	fmt.Printf("\nName of the file: %s", fileName) 
	fmt.Printf("\nSize of the file: %d bytes", len(data)) 
	fmt.Printf("\nData of the file: %s", data) 

} 

func main() { 

	CreateFile() 
	ReadFile() 
} 
