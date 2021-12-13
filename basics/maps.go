package main

import "fmt"





func main(){

  var users = make(map[string]string)
  users["Rob"] = "16"
  users["Bob"] = "20"
  users["Dod"] = "29"
  users["Brob"] = "23"

  for key,e := range users{
    fmt.Printf("Key: %s  Element: %s\n",key,e)
  }
  var usersArr = make([]string,5)
  usersArr[0]="Rob"
  usersArr[1]="Bob"
  for key,e := range usersArr{
    fmt.Printf("key: %s  Element: %s\n",key,e)
  }
  
  
  
}
