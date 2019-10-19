// key encapsulation Go example
package main

import (
    "bytes"
    "fmt"
    "oqs"
)

func main() {
    fmt.Println("Supported KEMs:")
    fmt.Println(oqs.GetSupportedKEMs())

    fmt.Println("\nEnabled KEMs:")
    fmt.Println(oqs.GetEnabledKEMs())

    kemName := "DEFAULT"
    client := oqs.KeyEncapsulation{}
    client.Init(kemName, []byte{})

    clientPublicKey := client.GenerateKeypair()
    fmt.Printf("\nKEM details:\n%#v\n", client.GetDetails())

    server := oqs.KeyEncapsulation{}
    server.Init(kemName, []byte{})
    ciphertext, sharedSecretServer := server.EncapSecret(clientPublicKey)

    sharedSecretClient := client.DecapSecret(ciphertext)

    fmt.Printf("\nClient shared secret:\n% X ... % X\n",
        sharedSecretClient[0:8], sharedSecretClient[len(sharedSecretClient)-8:])
    fmt.Printf("\nServer shared secret:\n% X ... % X\n",
        sharedSecretServer[0:8], sharedSecretServer[len(sharedSecretServer)-8:])

    isValid := bytes.Compare(sharedSecretClient, sharedSecretServer) == 0
    fmt.Println("\nShared secrets coincide? ", isValid)

    client.Release()
    server.Release()
}