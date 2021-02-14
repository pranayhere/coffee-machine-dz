# Coffee Machine

Project simulates coffee-machine

## Usage
To run the code please execute following command in the directory code is copied.
` go run cmd/main.go`

## Overview

A coffee-machine can have N outlets, and can serve N drinks simultaneously. We are using worker-pool concurrency model of golang to simulate the same.

When coffee machine is started, pool of worker of the size of coffee machine is initialized, and it listen to `orders` channel, once the order is received, one of the worker picks the order process it.