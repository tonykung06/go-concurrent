package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// synchronizedProcessing()
	// firstLevelParallelism()
	secondLevelParallelism()
}

func secondLevelParallelism() {
	// runtime.GOMAXPROCS(4)
	start := time.Now()

	extractChannel := make(chan *Order)
	transformChannel := make(chan *Order)
	doneChannel := make(chan bool)

	go extractSecond(extractChannel)
	go transformSecond(extractChannel, transformChannel)
	go loadSecond(transformChannel, doneChannel)

	<-doneChannel
	fmt.Println(time.Since(start))
}

func firstLevelParallelism() {
	start := time.Now()

	extractChannel := make(chan *Order)
	transformChannel := make(chan *Order)
	doneChannel := make(chan bool)

	go extractFirst(extractChannel)
	go transformFirst(extractChannel, transformChannel)
	go loadFirst(transformChannel, doneChannel)

	<-doneChannel
	fmt.Println(time.Since(start))
}

func synchronizedProcessing() {
	start := time.Now()
	orders := extractSync()
	orders = transformSync(orders)
	loadSync(orders)
	fmt.Println(time.Since(start))
}

type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

type Order struct {
	CustomerNumber int
	PartNumber     string
	Quantity       int
	UnitCost       float64
	UnitPrice      float64
}

func extractSync() []*Order {
	result := []*Order{}
	f, _ := os.Open("./orders.txt")
	defer f.Close()
	r := csv.NewReader(f)
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		result = append(result, order)
	}
	return result
}

func transformSync(orders []*Order) []*Order {
	f, _ := os.Open("./productList.txt")
	defer f.Close()
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}

	for idx, _ := range orders {
		time.Sleep(3 * time.Millisecond)
		o := orders[idx]
		o.UnitCost = productList[o.PartNumber].UnitCost
		o.UnitPrice = productList[o.PartNumber].UnitPrice
	}

	return orders
}

func loadSync(orders []*Order) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n",
		"Part Number", "Quantity", "Unit Cost",
		"Unit Price", "Total Cost", "Total Price")

	for _, o := range orders {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
			o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice,
			o.UnitCost*float64(o.Quantity),
			o.UnitPrice*float64(o.Quantity))
	}
}

func extractFirst(ch chan *Order) {
	f, _ := os.Open("./orders.txt")
	defer f.Close()
	r := csv.NewReader(f)
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		ch <- order
	}
	close(ch)
}

func transformFirst(extractChannel, transformChannel chan *Order) {
	f, _ := os.Open("./productList.txt")
	defer f.Close()
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}

	for o := range extractChannel {
		time.Sleep(3 * time.Millisecond)
		o.UnitCost = productList[o.PartNumber].UnitCost
		o.UnitPrice = productList[o.PartNumber].UnitPrice
		transformChannel <- o
	}
	close(transformChannel)
}

func loadFirst(transformChannel chan *Order, doneChannel chan bool) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n",
		"Part Number", "Quantity", "Unit Cost",
		"Unit Price", "Total Cost", "Total Price")

	for o := range transformChannel {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
			o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice,
			o.UnitCost*float64(o.Quantity),
			o.UnitPrice*float64(o.Quantity))
	}
	doneChannel <- true
}

func extractSecond(ch chan *Order) {
	f, _ := os.Open("./orders.txt")
	defer f.Close()
	r := csv.NewReader(f)
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		ch <- order
	}
	close(ch)
}

func transformSecond(extractChannel, transformChannel chan *Order) {
	f, _ := os.Open("./productList.txt")
	defer f.Close()
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}

	numMessage := 0
	for o := range extractChannel {
		numMessage++
		go func(o *Order) {
			time.Sleep(3 * time.Millisecond)
			o.UnitCost = productList[o.PartNumber].UnitCost
			o.UnitPrice = productList[o.PartNumber].UnitPrice
			transformChannel <- o
			numMessage--
		}(o)
	}
	for numMessage > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	close(transformChannel)
}

func loadSecond(transformChannel chan *Order, doneChannel chan bool) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n",
		"Part Number", "Quantity", "Unit Cost",
		"Unit Price", "Total Cost", "Total Price")

	numMessage := 0
	for o := range transformChannel {
		numMessage++
		go func(o *Order) {
			time.Sleep(1 * time.Millisecond)
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
				o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice,
				o.UnitCost*float64(o.Quantity),
				o.UnitPrice*float64(o.Quantity))

			numMessage--
		}(o)
	}
	for numMessage > 0 {
		time.Sleep(1 * time.Millisecond)
	}
	doneChannel <- true
}
