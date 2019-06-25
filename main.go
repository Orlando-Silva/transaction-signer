package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"C"

	"context"
	"github.com/btcsuite/btcd/btcec"
	"github.com/aws/aws-lambda-go/lambda"
)

type SignEvent struct {
	KeySigner string `json:"key"`
	DataToSign string `json:"data"`
}


func Sign(private, data []byte) ([]byte, error) {
	privkey, _ := btcec.PrivKeyFromBytes(btcec.S256(), private)
	sig, err := privkey.Sign(data)
	if err != nil {
		return nil, err
	}
	return sig.Serialize(), nil
}

func HandleRequest(ctx context.Context, receivedObject SignEvent) (string, error) {

	returnedData, errr := SignData(receivedObject.KeySigner, receivedObject.DataToSign)

	if errr == nil {
		return fmt.Sprintf(returnedData), nil
	} 

	return "Error", nil
}


func SignData(keyReceived, dataReceived string) (string, error) {


	keyDecoded,err :=  hex.DecodeString(keyReceived)
	
	if err != nil {
		log.Fatal(err)
	}

	dataDecoded, err :=  hex.DecodeString(dataReceived)

	if err != nil {
		log.Fatal(err)
	}

	signedString, errors := Sign(keyDecoded, dataDecoded)

	if errors == nil {
		return hex.EncodeToString(signedString), nil
	} 

	return "Not well", nil
	
} 

func main() {
	lambda.Start(HandleRequest)
}
