package main

import (
	"encoding/json"
	"math"
	"math/big"
	"net/http"

	token "main/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Result struct {
	UserAddr string
	Balance  *big.Float
	Contract string
	Error    string
}

func checkBalance(contract string, addr string) Result {
	result := Result{}
	client, err := ethclient.Dial("ENDPOINT")
	if err != nil {
		result.Error = err.Error()
		return result
	}

	tokenAddr := common.HexToAddress(contract)
	instance, err := token.NewToken(tokenAddr, client)
	if err != nil {
		result.Error = err.Error()
		return result
	}
	userAddr := common.HexToAddress(addr)

	bal, err := instance.BalanceOf(&bind.CallOpts{}, userAddr)
	if err != nil {
		result.Error = err.Error()
		return result
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
	result.UserAddr = userAddr.String()
	result.Balance = value
	result.Contract = tokenAddr.String()
	return result
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		contract := c.QueryParam("contract")
		addr := c.QueryParam("addr")
		if len(contract) == 0 || len(addr) == 0 {
			contract = "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"
			addr = "0xc2132d05d31c914a87c6611c10748aeb04b58e8f"
		}

		result := checkBalance(contract, addr)

		data, err := json.Marshal(result)
		if err != nil {
			logrus.Errorln(err)
			return c.String(http.StatusBadGateway, "")
		}
		return c.String(http.StatusOK, string(data))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
