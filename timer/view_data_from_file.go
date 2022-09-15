package timer

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	mathRand "math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"super-view-server/leveldb"
	"sync"
	"time"
)

type ViewResultInfo struct {
	Data      Data   `json:"data"`
	ErrorCode uint16 `json:"error_code"`
}

type Data struct {
	TotalUsdValue float64 `json:"total_usd_value"`
}

type OpenApiViewResultInfo struct {
	TotalUsdValue float64 `json:"total_usd_value"`
}

type HitList struct {
	HitItems []HitData `json:"hit_list"`
}
type HitData struct {
	Address    string `json:"address"`
	Line       string `json:"line"`
	ResultJson string `json:"result_json"`
}

type Counter struct {
	ViewCounter uint16 `json:"view_counter"`
}

var (
	tokenLock sync.RWMutex
)

func (t *TxTimer) doViewFromFile() error {
	tokenLock.Lock()
	defer tokenLock.Unlock()
	return PreExecuteScrape()
}

func ClearLevelDbData() {
	leveldb.LevelDb.Delete("view_data_counter")
	leveldb.LevelDb.Delete("hit_data")
}

func CheckLevelDbData() {
	hitByte, _ := leveldb.LevelDb.Get("hit_data")
	hitList := make([]HitData, 0)
	json.Unmarshal(hitByte, &hitList)
	log.Info(string(hitByte))
}

func PreExecuteScrape() error {
	CheckLevelDbData()
	//ClearLevelDbData()
	counter := uint16(0)
	startView := uint16(0)
	filePath := "./file/reval.txt"
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	flag, _ := leveldb.LevelDb.Has("view_data_counter")
	if flag {
		counterByte, _ := leveldb.LevelDb.Get("view_data_counter")
		counterObject := &Counter{}
		json.Unmarshal(counterByte, counterObject)
		startView = counterObject.ViewCounter
	}
	for {
		lineByte, _, err1 := buf.ReadLine()
		line := strings.TrimSpace(string(lineByte))
		log.Infof("Data counter:%d, line: %s ", counter, line)
		counter++
		if counter > startView {
			ExecuteLineByDebankOpenApi(line, counter)
		}
		if err1 != nil {
			if err1 == io.EOF {
				return nil
			}
			return err
		}
		counter := Counter{
			ViewCounter: counter,
		}
		leveldb.LevelDb.Put("view_data_counter", counter)
	}
}

func ExecuteLineByDebankScrape(line string, counter uint16) int {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	randValue := RangeRand(5, 10)
	time.Sleep(time.Duration(randValue) * time.Second)
	pattern := "https://api.debank.com/user/total_balance?addr=%v"
	address := strings.Split(line, "#")[1]
	//address = "0x9f4A156c93E95636A6Cf00f974828BE47956e8F8"
	url := fmt.Sprintf(pattern, address)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	res, err1 := client.Do(req)
	if err1 != nil {
		log.Error("err:", err1.Error())
		return 0
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Errorf("status code error: %#v %#v", res.StatusCode, res.Status)
		return 0
	}
	//{"_cache_seconds":0,"_seconds":0.15242624282836914,"_use_cache":false,"data":{"total_usd_value":4.006184853456298},"error_code":0}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	resultInfo := ViewResultInfo{}
	err2 := json.Unmarshal(body, &resultInfo)
	if err2 != nil {
		log.Errorf("json.Unmarshal error: %#v %#v", body, err2.Error())
		return 0
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", resultInfo.Data.TotalUsdValue), 64)
	decimalValue := decimal.NewFromFloat(value).Mul(decimal.NewFromInt(100)).BigInt()
	if decimalValue.Sign() > 0 {
		log.Infof("★★★★★counter:%d, url:%s result:%s", counter, url, string(body))
		flag, _ := leveldb.LevelDb.Has("hit_data")
		hitList := make([]HitData, 0)
		if flag {
			itemByte, _ := leveldb.LevelDb.Get("hit_data")
			json.Unmarshal(itemByte, &hitList)
			data := HitData{
				Line:       line,
				ResultJson: string(body),
			}
			hitList = append(hitList, data)
			leveldb.LevelDb.Put("hit_data", hitList)
		} else {
			hitList := make([]HitData, 0)
			data := HitData{
				Line:       line,
				ResultJson: string(body),
			}
			hitList = append(hitList, data)
			leveldb.LevelDb.Put("hit_data", hitList)
		}
	} else {
		log.Infof("☆☆☆☆☆counter:%d, url:%s result:%s", counter, url, body)
	}
	return 1
}

func ExecuteLineByDebankOpenApi(line string, counter uint16) int {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	randValue := RangeRand(5, 10)
	time.Sleep(time.Duration(randValue) * time.Second)
	pattern := "https://openapi.debank.com/v1/user/total_balance?id=%v"
	address := strings.Split(line, "#")[1]
	//address = "0x9f4A156c93E95636A6Cf00f974828BE47956e8F8"
	url := fmt.Sprintf(pattern, address)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	res, err1 := client.Do(req)
	if err1 != nil {
		log.Error("err:", err1.Error())
		return 0
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Errorf("status code error: %#v %#v", res.StatusCode, res.Status)
		return 0
	}
	//{"_cache_seconds":0,"_seconds":0.15242624282836914,"_use_cache":false,"data":{"total_usd_value":4.006184853456298},"error_code":0}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	resultInfo := OpenApiViewResultInfo{}
	err2 := json.Unmarshal(body, &resultInfo)
	if err2 != nil {
		log.Errorf("json.Unmarshal error: %#v %#v", body, err2.Error())
		return 0
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", resultInfo.TotalUsdValue), 64)
	decimalValue := decimal.NewFromFloat(value).Mul(decimal.NewFromInt(100)).BigInt()
	if decimalValue.Sign() > 0 {
		log.Infof("★★★★★counter:%d, url:%s result:%#v", counter, url, resultInfo)
		flag, _ := leveldb.LevelDb.Has("hit_data")
		hitList := make([]HitData, 0)
		if flag {
			itemByte, _ := leveldb.LevelDb.Get("hit_data")
			json.Unmarshal(itemByte, &hitList)
			data := HitData{
				Address:    address,
				Line:       line,
				ResultJson: fmt.Sprintf("%v", decimalValue.Sign()),
			}
			hitList = append(hitList, data)
			leveldb.LevelDb.Put("hit_data", hitList)
		} else {
			hitList := make([]HitData, 0)
			data := HitData{
				Address:    address,
				Line:       line,
				ResultJson: fmt.Sprintf("%v", decimalValue.Sign()),
			}
			hitList = append(hitList, data)
			leveldb.LevelDb.Put("hit_data", hitList)
		}
	} else {
		log.Infof("☆☆☆☆☆counter:%d, url:%s result:%s", counter, url, body)
	}
	return 1
}

func Shuffle(arr interface{}) {
	contentType := reflect.TypeOf(arr)
	if contentType.Kind() != reflect.Slice {
		panic("expects a slice type")
	}
	contentValue := reflect.ValueOf(arr)
	source := mathRand.NewSource(time.Now().UnixNano())
	random := mathRand.New(source)
	len := contentValue.Len()
	for i := len - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		x, y := contentValue.Index(i).Interface(), contentValue.Index(j).Interface()
		contentValue.Index(i).Set(reflect.ValueOf(y))
		contentValue.Index(j).Set(reflect.ValueOf(x))
	}
}

func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))
		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
