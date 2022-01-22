package main

import (
  "log"
  "os"
)

func main {
  emptyFile, err := os.Create("emptyFile.txt")   //  os. pkg allows for interfacing with operating systems
  if err != nil {
      log.Fatal(err)
  }
  log.Println(emptyFile)
  emptyFile.Close()       
}
