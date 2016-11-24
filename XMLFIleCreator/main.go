
package main

import (
		"fmt"
		"os"
		//"encoding/xml"
)


func openXML(name string){
	
	xmlFile, err := os.Open(name + ".xml")
	if err != nil {
		fmt.Printf("Error opening file:", err )
		createXML(name)
		
	}
	defer xmlFile.Close()

}

func createXML(name string){
	var n string
	
	fmt.Printf("\nDo you wish to create it? y/n : ")
		
		_, err := fmt.Scan(&n)
		if err != nil {
			panic(err)
		}
		if n == "y" {
			_, err := os.Create(name + ".xml")
			if err != nil{
				fmt.Printf("Error creating file:", err)
			}
		} else {
			return
		}

}


func main(){
	var name string
	
	fmt.Printf("Which file do you want to open?")
	_, err := fmt.Scan(&name)
	if err != nil{
		fmt.Printf("Error on input:" , err)
	}
	openXML(name)
}