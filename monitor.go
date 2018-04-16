package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var account string

func monitor() {
	for _, v := range Conf.Accounts {
		account = v
		Acc := `{"uid":` + `"` + account + `"` + "}"
		for {
			<-CRechargeOK
			body, err := HttpPostJsonReq(Conf.Api.WalletInfo, Acc)
			if err != nil {
				log.Println(err)
			}
			err = Parsing(body)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func Parsing(body []byte) error {
	conf := InitConfig()
	h := make(map[string][]string)
	str := make(map[string]float64)
	var r interface{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return err
	}
	data, ok := r.(map[string]interface{})
	if ok {
		for k, v := range data {
			switch v1 := v.(type) {
			case interface{}:
				for m1, n1 := range v1.(map[string]interface{}) {
					switch v2 := n1.(type) {
					case interface{}:
						for m2, n2 := range v2.(map[string]interface{}) {
							switch v3 := n2.(type) {
							case interface{}:
								switch v4 := v3.(type) {
								case float64:
									switch m1 {
									case "1":
										if m2 == "active" {
											str["CNY"] = v4
											if v4/1e8 < float64((*conf).Charges["CNY"].MinActiveAllowed) {
												h[account] = append(h[account], "CNY")
											}
										}
									case "2":
										if m2 == "active" {
											str["BTC"] = v4
											if v4/1e8 < float64((*conf).Charges["BTC"].MinActiveAllowed) {
												h[account] = append(h[account], "BTC")
											}
										}
									case "3":
										if m2 == "active" {
											str["BTY"] = v4
											if v4/1e8 < float64((*conf).Charges["BTY"].MinActiveAllowed) {
												h[account] = append(h[account], "BTY")
											}
										}
									case "4":
										if m2 == "active" {
											str["ETH"] = v4
											if v4/1e8 < float64((*conf).Charges["ETH"].MinActiveAllowed) {
												h[account] = append(h[account], "ETH")
											}
										}
									case "5":
										if m2 == "active" {
											str["ETC"] = v4
											if v4/1e8 < float64((*conf).Charges["ETC"].MinActiveAllowed) {
												h[account] = append(h[account], "ETC")
											}
										}
									case "7":
										if m2 == "active" {
											str["SC"] = v4
											if v4/1e8 < float64((*conf).Charges["SC"].MinActiveAllowed) {
												h[account] = append(h[account], "SC")
											}
										}
									case "8":
										if m2 == "active" {
											str["ZEC"] = v4
											if v4/1e8 < float64((*conf).Charges["ZEC"].MinActiveAllowed) {
												h[account] = append(h[account], "ZEC")
											}
										}
									case "9":
										if m2 == "active" {
											str["BTS"] = v4
											if v4/1e8 < float64((*conf).Charges["BTS"].MinActiveAllowed) {
												h[account] = append(h[account], "BTS")
											}
										}
									case "10":
										if m2 == "active" {
											str["LTC"] = v4
											if v4/1e8 < float64((*conf).Charges["LTC"].MinActiveAllowed) {
												h[account] = append(h[account], "LTC")
											}
										}
									case "11":
										if m2 == "active" {
											str["BCC"] = v4
											if v4/1e8 < float64((*conf).Charges["BCC"].MinActiveAllowed) {
												h[account] = append(h[account], "BCC")
											}
										}
									case "15":
										if m2 == "active" {
											str["USDT"] = v4
											if v4/1e8 < float64((*conf).Charges["USDT"].MinActiveAllowed) {
												h[account] = append(h[account], "USDT")
											}
										}
									case "17":
										if m2 == "active" {
											str["DCR"] = v4
											if v4/1e8 < float64((*conf).Charges["DCR"].MinActiveAllowed) {
												h[account] = append(h[account], "DCR")
											}
										}
									}
								}
							}
						}
					}
				}
			default:
				err = fmt.Errorf(k + "is another type not handle yet")
				return err
				// fmt.Println(k, "is another type not handle yet")
			}
		}
	}
	CData <- h
	return nil
}

type Asset_ struct {
	Coin map[string]Number_ `json:"data"`
}

type Number_ struct {
	Active   int `json:"active"`
	Frozen   int `json:"frozen"`
	Poundage int `json:"poundage"`
}

func ParsingWeb_(body []byte) map[string]float64 {
	CoinIdName := make(map[string]string, 0)
	CoinIdName["1"] = "CNY"
	CoinIdName["2"] = "BTC"
	CoinIdName["3"] = "BTY"
	CoinIdName["4"] = "ETH"
	CoinIdName["5"] = "ETC"
	CoinIdName["7"] = "SC"
	CoinIdName["8"] = "ZEC"
	CoinIdName["9"] = "BTS"
	CoinIdName["10"] = "LTC"
	CoinIdName["11"] = "BCC"
	CoinIdName["15"] = "USDT"
	CoinIdName["17"] = "DCR"

	var asset Asset_
	err := json.Unmarshal(body, &asset)
	if err != nil {
		log.Println("json Unmarshal failed")
	}
	assetMap := make(map[string]float64, 0)
	for coinId, CoinName := range CoinIdName {
		assetMap[CoinName] = float64(asset.Coin[coinId].Active) / 1e+8
	}
	return assetMap
}
