
package main

import (
		"fmt"
		"os"
		//"encoding/xml"
)


func openXML(){
	
	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Printf("Error opening file:", err )
		createXML()
		
	}
	defer xmlFile.Close()

}

func createXML(){
	var n string
	
	fmt.Printf("\nDo you wish to create it? y/n : ")
		
		_, err := fmt.Scanf("%s", &n)
		if err != nil {
			panic(err)
		}
		if n == "y" {
			_, err := os.Create("data.xml")
			if err != nil{
				fmt.Printf("Error creating file:", err)
			}
		} else {
			return
		}

}


func main(){
	openXML()
}