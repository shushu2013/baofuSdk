# Baofu Payment SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/shushu2013/baofuSdk.svg)](https://pkg.go.dev/github.com/shushu2013/baofuSdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/shushu2013/baofuSdk)](https://goreportcard.com/report/github.com/shushu2013/baofuSdk)

A Go SDK for integrating with Baofu payment services.

## Installation

```bash
go get github.com/shushu2013/baofuSdk@v0.0.1
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/shushu2013/baofuSdk"
)

func main() {
	// Initialize SDK client
	config := &baofuSdk.Config{
		MerchantID: "your_merchant_id",
		TerminalID: "your_terminal_id",
		SecretKey:  "your_secret_key",
		BaseURL:    "https://api.baofu.com",
	}

	client := baofuSdk.NewClient(config)

	// Create payment order
	req := &baofuSdk.PaymentRequest{
		Amount:      "100.00",
		OrderNo:     "ORDER123456",
		ProductName: "Test Product",
	}

	resp, err := client.CreatePayment(req)
	if err != nil {
		panic(err)
	}

	if resp.Success {
		fmt.Printf("Payment created successfully. Pay URL: %s\n", resp.PayURL)
	} else {
		fmt.Printf("Payment failed: %s\n", resp.Message)
	}
}
```
