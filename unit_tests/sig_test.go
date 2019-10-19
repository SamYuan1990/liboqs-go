// oqs.Signature unit testing
package testing

import (
    "fmt"
    "oqs"
    "testing"
)

func TestSignature(t *testing.T) {
    var signer, verifier oqs.Signature
    msg := []byte("This is our favourite message to sign")

    for _, sigName := range oqs.GetEnabledSIGs() {
        fmt.Println(sigName)
        signer.Init(sigName, []byte{})
        verifier.Init(sigName, []byte{})
        pubKey := signer.GenerateKeypair()
        signature := signer.Sign(msg)
        isValid := verifier.Verify(msg, signature, pubKey)
        if !isValid {
            t.Fatal("Signature verification failed")
        }
        signer.Release()
        verifier.Release()
    }
}

func TestUnsupportedSignature(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Unsupported signature should have generated a panic")
        }
    }()
    signer := oqs.Signature{}
    signer.Init("unsupported_sig", []byte{})
}
