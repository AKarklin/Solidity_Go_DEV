package main

import (
	"net/http"

	"github.com/AKarklin/Solidity_Go_DEV/contracts_go_api/Test01/api" // this would be your generated smart contract bindings
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	conn, err := api.NewApi(common.HexToAddress("0x7c9F32bF623525aa4Ec2Af9bC8A53FbA12FfdF0B"), client)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/greet/:message", func(c echo.Context) error {
		message := c.Param("message")
		reply, err := conn.Greet(&bind.CallOpts{}, message)
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, reply)
		return nil
	})
	e.GET("/hello", func(c echo.Context) error {
		reply, err := conn.Hello(&bind.CallOpts{})
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, reply) // Hello World
		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
