package main

import (
//    "strings"
//"fmt"
    "github.com/ArturC03/shellcade/functions"
    "github.com/eiannone/keyboard"
    "log"
)

func RunSnake() {
  cursor.Clear()
  width,height,_ := cursor.GetScreenSize()
  cursor.Position(width/2,height/2)

  if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()


  for { // Read key
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
    
    if char == 's'{
      
    }

  }


    keyboard.Close()
}
