package timer

import (
	"encoding/json"
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/scorpiotzh/mylog"
	"github.com/shopspring/decimal"
	"github.com/tyler-smith/go-bip39"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"super-view-server/leveldb"
	"sync"
	"time"
)

var (
	log                = mylog.NewLogger("config", mylog.LevelDebug)
	randomMnemonicLock sync.RWMutex
)

func (t *TxTimer) doRandomMnemonicView() error {
	randomMnemonicLock.Lock()
	defer randomMnemonicLock.Unlock()
	return preExecuteRandomMnemonicView()
}

func preExecuteRandomMnemonicView() error {
	counter := uint16(0)
	flag, _ := leveldb.LevelDb.Has("view_data_counter")
	if flag {
		counterByte, _ := leveldb.LevelDb.Get("view_data_counter")
		counterObject := &Counter{}
		json.Unmarshal(counterByte, counterObject)
		counter = counterObject.ViewCounter
	}
	for {
		entropy, err1 := bip39.NewEntropy(128)
		if err1 != nil {
			log.Error(err1)
		}

		mnemonic, _ := bip39.NewMnemonic(entropy)
		//var mnemonic = "pepper hair process town say voyage exhibit over carry property follow define"
		//log.Info("mnemonic:", mnemonic)
		seed := bip39.NewSeed(mnemonic, "") //这里可以选择传入指定密码或者空字符串，不同密码生成的助记词不同

		wallet, err2 := hdwallet.NewFromSeed(seed)
		if err2 != nil {
			log.Error(err2)
		}

		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0") //最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址
		account, err3 := wallet.Derive(path, false)
		if err3 != nil {
			log.Error(err3)
		}

		privateKey, _ := wallet.PrivateKeyHex(account)
		//publicKey, _ := wallet.PublicKeyHex(account)
		address0 := account.Address.Hex()

		//log.Info("privateKey:", privateKey) // 私钥
		//log.Info("publicKey:", publicKey)   // 公钥
		//log.Info("address0:", address0)     // id为0的钱包地址

		/*
			path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") //生成id为1的钱包地址
			account1, err4 := wallet.Derive(path, false)
			if err4 != nil {
				log.Error(err4)
			}
			log.Info("address1:", account1.Address.Hex())
		*/

		line := fmt.Sprintf("%v#%v#%v", privateKey, address0, mnemonic)
		log.Info("line:", line)
		counter++
		randValue := RangeRand(1, 10)
		if randValue%2 == 0 {
			code := ExecuteLineByDebankApiWithRandom(line, counter)
			if code == 429 {
				time.Sleep(time.Duration(10) * time.Second)
			}
		} else {
			code := ExecuteLineByDebankOpenApiWithRandom(line, counter)
			if code == 429 {
				time.Sleep(time.Duration(10) * time.Second)
			}
		}
		counter := Counter{
			ViewCounter: counter,
		}
		leveldb.LevelDb.Put("view_data_counter", counter)
	}
}

func ExecuteLineByDebankApiWithRandom(line string, counter uint16) int {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	pattern := "https://api.debank.com/user/total_balance?addr=%v"
	address := strings.Split(line, "#")[1]
	//address = "0x9f4A156c93E95636A6Cf00f974828BE47956e8F8"
	url := fmt.Sprintf(pattern, address)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	res, err1 := client.Do(req)
	log.Info("Req url:", url)
	if err1 != nil {
		log.Error("err:", err1.Error())
		return 0
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Errorf("status code error: %#v %#v", res.StatusCode, res.Status)
		return res.StatusCode
	}
	//{"total_usd_value": 0, "chain_list": []}	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	resultInfo := ViewResultInfo{}
	err2 := json.Unmarshal(body, &resultInfo)
	if err2 != nil {
		log.Errorf("json.Unmarshal error: %#v %#v", body, err2.Error())
		return 0
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", resultInfo.Data.TotalUsdValue), 64)
	decimalValue := decimal.NewFromFloat(value).Mul(decimal.NewFromInt(100)).BigInt()
	flag, _ := leveldb.LevelDb.Has("hit_data_with_random")
	if decimalValue.Sign() > 0 {
		log.Infof("●●●●●counter:%d hasFlag:%#v url:%s result:%#v", counter, flag, url, resultInfo)
		hitList := make([]HitData, 0)
		if flag {
			itemByte, _ := leveldb.LevelDb.Get("hit_data_with_random")
			json.Unmarshal(itemByte, &hitList)
			data := HitData{
				Address:    address,
				Line:       line,
				ResultJson: fmt.Sprintf("%v", decimalValue.Sign()),
			}
			hitList = append(hitList, data)
			leveldb.LevelDb.Put("hit_data_with_random", hitList)
		} else {
			hitList := make([]HitData, 0)
			data := HitData{
				Address:    address,
				Line:       line,
				ResultJson: fmt.Sprintf("%v", decimalValue.Sign()),
			}
			hitList = append(hitList, data)
			leveldb.LevelDb.Put("hit_data_with_random", hitList)
		}
	} else {
		log.Infof("○○○○○counter:%d hasFlag:%#v url:%s result:%s", counter, flag, url, body)
	}
	return 1
}

func ExecuteLineByDebankOpenApiWithRandom(line string, counter uint16) int {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	randValue := RangeRand(1, 5)
	time.Sleep(time.Duration(randValue) * time.Second)
	pattern := "https://openapi.debank.com/v1/user/total_balance?id=%v"
	address := strings.Split(line, "#")[1]
	//address = "0x9f4A156c93E95636A6Cf00f974828BE47956e8F8"
	url := fmt.Sprintf(pattern, address)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	res, err1 := client.Do(req)
	log.Info("Req url:", url)
	if err1 != nil {
		log.Error("err:", err1.Error())
		return 0
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Errorf("status code error: %#v %#v", res.StatusCode, res.Status)
		return res.StatusCode
	}
	//{"_cache_seconds":0,"_seconds":0.15242624282836914,"_use_cache":false,"data":{"total_usd_value":4.006184853456298},"error_code":0}
	body, _ := ioutil.ReadAll(res.Body)
	resultInfo := OpenApiViewResultInfo{}
	err2 := json.Unmarshal(body, &resultInfo)
	if err2 != nil {
		log.Errorf("json.Unmarshal error: %#v %#v", body, err2.Error())
		return 0
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", resultInfo.TotalUsdValue), 64)
	decimalValue := decimal.NewFromFloat(value).Mul(decimal.NewFromInt(100)).BigInt()
	flag, _ := leveldb.LevelDb.Has("hit_data_with_random")
	if decimalValue.Sign() > 0 {
		log.Infof("●●●●●counter:%d hasFlag:%#v url:%s result:%#v", counter, flag, url, resultInfo)
		hitList := make([]HitData, 0)
		if flag {
			itemByte, _ := leveldb.LevelDb.Get("hit_data_with_random")
			json.Unmarshal(itemByte, &hitList)
			data := HitData{
				Address:    address,
				Line:       line,
				ResultJson: fmt.Sprintf("%v", decimalValue.Sign()),
			}
			hitList = append(hitList, data)
			leveldb.LevelDb.Put("hit_data_with_random", hitList)
		} else {
			hitList := make([]HitData, 0)
			data := HitData{
				Address:    address,
				Line:       line,
				ResultJson: fmt.Sprintf("%v", decimalValue.Sign()),
			}
			hitList = append(hitList, data)
			leveldb.LevelDb.Put("hit_data_with_random", hitList)
		}
	} else {
		log.Infof("○○○○○counter:%d hasFlag:%#v url:%s result:%s", counter, flag, url, body)
	}
	return 1
}
