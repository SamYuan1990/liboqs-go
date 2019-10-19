// oqs.KeyEncapsulation unit testing
package testing

import (
    "bytes"
    "fmt"
    "oqs"
    "testing"
)

func TestKeyEncapsulation(t *testing.T) {
    var client, server oqs.KeyEncapsulation

    for _, kemName := range oqs.GetEnabledKEMs() {
        fmt.Println(kemName)
        client.Init(kemName, []byte{})
        server.Init(kemName, []byte{})
        clientPublicKey := client.GenerateKeypair()
        ciphertext, sharedSecretServer := server.EncapSecret(clientPublicKey)
        sharedSecretClient := client.DecapSecret(ciphertext)
        isValid := bytes.Compare(sharedSecretClient, sharedSecretServer) == 0
        if !isValid {
            t.Fatal("Shared secrets do not coincide")
        }
        client.Release()
        server.Release()
    }
}

func TestUnsupportedKeyEncapsulation(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Unsupported KEM should have generated a panic")
        }
    }()
    client := oqs.KeyEncapsulation{}
    client.Init("unsupported_kem", []byte{})
}