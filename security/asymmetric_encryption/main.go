package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)


const (
  keySize = 512
)

func WritePrivateKeyFile(privateKey *rsa.PrivateKey) error{
  privateKeyFile := "key.private"
  f,err := os.Create(privateKeyFile)
  if err != nil{
    return err
  }
  defer f.Close()
  w := gob.NewEncoder(f)
  w.Encode(privateKey)
  return nil
}
func WritePublicKeyFile(privateKey *rsa.PrivateKey) error{
  privateKeyFile := "key.public"
  f,err := os.Create(privateKeyFile)
  if err != nil{
    return err
  }
  defer f.Close()
  w := gob.NewEncoder(f)
  w.Encode(privateKey.Public())
  return nil
}
func main(){
  rnd := rand.Reader
  privateKey, err := rsa.GenerateKey(rnd,keySize)
  if err != nil {
    log.Fatalf("error generation key %v", err)
  }
  WritePublicKeyFile(privateKey)
  WritePrivateKeyFile(privateKey)
  fmt.Printf("private key: %v\n",privateKey)
  fmt.Printf("public key: %v\n",privateKey.Public())

}
