package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"time"
)

type OrderSide uint8
type OrderType uint8

const (
	Buy  OrderSide = 0
	Sell OrderSide = 1

	Market OrderType = 0
	Limit  OrderType = 1
)

type WireOrder struct {
	OrderID   uint64
	ClientID  uint32
	Symbol    [8]byte
	Side      OrderSide
	Type      OrderType
	Price     float64
	Quantity  float64
	Timestamp uint64
}

func sendOrders(conn net.Conn) {
	orderID := uint64(1)
	clientID := uint32(123)
	symbol := [8]byte{'A', 'A', 'P', 'L', 0, 0, 0, 0}
	side := Buy

	for {
		order := WireOrder{
			OrderID:   orderID,
			ClientID:  clientID,
			Symbol:    symbol,
			Side:      side,
			Type:      Market,
			Price:     100.0,
			Quantity:  1.0,
			Timestamp: uint64(time.Now().UnixMilli()),
		}

		buf := new(bytes.Buffer)
		if err := binary.Write(buf, binary.LittleEndian, &order); err != nil {
			log.Println("Failed to serialize order:", err)
			continue
		}

		if _, err := conn.Write(buf.Bytes()); err != nil {
			log.Println("Failed to send order:", err)
			return
		}

		if side == Buy {
			log.Println("Placed buy order.")
			side = Sell
		} else {
			log.Println("Placed sell order.")
			side = Buy
		}

		orderID++
		time.Sleep(5 * time.Second)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5005")
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	go sendOrders(conn)

	select {}
}
