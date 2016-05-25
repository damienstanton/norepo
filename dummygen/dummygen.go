package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// FakeData tells us how to marshal/unmarshal the correct type for each column.
type FakeData struct {
	duns                   int
	contractNumber         string
	orderID                string
	orderDateSk            string
	productDescription     string
	productManufacturer    string
	manufacturerPartNumber string
	quantity               int
	upc                    int
	unitSalesPrice         float32
	totalSalesPrice        float32
	psc                    int
	airBillNumber          int
	pickupDateSk           string
	actualWeightUnit       string
	numOfPiecesInShipment  int
	shipmentCost           float32
	primeCarrierScac       string
	invoiceNumber          int
	invoiceDate            string
	lineItem               string
	description            string
	unitOfMeasure          string
	lineItemQuantitiy      int
	lineItemSellingPrince  float32
	lineItemTotalPrice     float32
	awardVehicleName       string
	agencyName             string
	bureauName             string
	supplierName           string
	sysSourceType          string
}

// Generate creates the fake dataset and writes it to a csv file.
func build() {
	// Create the csv file
	csvFile, err := os.Create("dummyset.csv")
	if err != nil {
		println("Error in creating a file. Check permissions?")
	}
	defer csvFile.Close()

	data := []FakeData{
		{rand.Int(), randStr(), randStr(), randStr(), randStr(), randStr(),
			randStr(), rand.Int(), rand.Int(), rand.Float32(), rand.Float32(),
			rand.Int(), rand.Int(), randStr(), randStr(), rand.Int(),
			rand.Float32(), randStr(), rand.Int(), randStr(), randStr(),
			randStr(), randStr(), rand.Int(), rand.Float32(), rand.Float32(),
			randStr(), randStr(), randStr(), randStr(), randStr()},
	}
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	for _, j := range data {
		w.Write(data)
	}

}

func randStr() string {
	s := strconv.Itoa(rand.Int())
	return s
}

func main() {
	v := "test val"
	fmt.Printf("Test val is set to: %v", v)
	build()
}
