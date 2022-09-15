package config

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/dotbitHQ/das-lib/common"
	"github.com/scorpiotzh/mylog"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

var (
	log = mylog.NewLogger("config", mylog.LevelDebug)
)

type CountryTimeZone struct {
	CountryCode string   `json:"countryCode"`
	CountryName string   `json:"countryName"`
	TimeZones   []string `json:"timeZones"`
	UTCOffset   []string `json:"UTCOffset"`
}

var (
	Cfg                    CfgServer
	AccountCharSetEmoji    string
	MapReservedAccounts    map[string]struct{}
	MapUnAvailableAccounts map[string]struct{}
	MapCountryTimeZone     map[string]CountryTimeZone
	AccountCharSetNumber   = "0123456789"
	AccountCharSetEn       = "abcdefghijklmnopqrstuvwxyz."
	AppEnv                 string
	//go:embed static
	StaticDir embed.FS
)

func init() {
	// 兼容 go test, go test使用了原生的 flag.Parse()
	if len(os.Args) >= 2 && strings.HasPrefix(os.Args[1], "-test.") {
		log.Info("run in go test, do not read flags")
		AppEnv = "test"
	} else {
		app := &cli.App{
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config",
					Aliases: []string{"c"},
					Usage:   "Load configuration from `FILE`",
				},
			},
			Action: func(context *cli.Context) error {
				AppEnv = context.String("config")
				if AppEnv == "" {
					AppEnv = "test"
				}
				return nil
			},
		}
		// parse flags
		if err := app.Run(os.Args); err != nil {
			log.Fatal(err.Error())
			panic(err)
		}
	}
	println("AppEnv=", AppEnv)
	// init config
	ymlData, err := StaticDir.ReadFile("static/config_" + AppEnv + ".yml")
	if err != nil {
		log.Fatal("parse yml fail, err=" + err.Error())
		panic(err)
	}
	yaml.Unmarshal(ymlData, &Cfg)
	// init data
	initDataConfig()
}

func initDataConfig() error {
	dataFilePath := "static/data/"
	// preserved_accounts.txt
	preservedData, err := StaticDir.ReadFile(dataFilePath + "preserved_accounts.txt")
	if err != nil {
		return fmt.Errorf("read preserved_accounts.txt err:%s", err.Error())
	}
	preservedArray := strings.Split(string(preservedData), "\n")
	MapReservedAccounts = make(map[string]struct{})
	for _, value := range preservedArray {
		MapReservedAccounts[value] = struct{}{}
	}
	//log.Info("MapReservedAccounts：", MapReservedAccounts)

	//unavailable_account_hashes.txt
	unavailableData, err := StaticDir.ReadFile(dataFilePath + "unavailable_account_hashes.txt")
	if err != nil {
		return fmt.Errorf("read unavailable_account_hashes.txt err:%s", err.Error())
	}
	unavailableArray := strings.Split(string(unavailableData), "\n")
	MapUnAvailableAccounts = make(map[string]struct{})
	for _, value := range unavailableArray {
		MapUnAvailableAccounts[value] = struct{}{}
	}
	//log.Info("MapUnAvailableAccounts：", MapUnAvailableAccounts)

	//char_set_emoji.txt
	emojiData, err := StaticDir.ReadFile(dataFilePath + "char_set_emoji.txt")
	if err != nil {
		return fmt.Errorf("read char_set_emoji.txt err:%s", err.Error())
	}
	AccountCharSetEmoji = strings.Replace(string(emojiData), "\n", "", -1)
	//log.Info("AccountCharSetEmoji：", AccountCharSetEmoji)

	//country
	countryTimeZoneData, err := StaticDir.ReadFile(dataFilePath + "country_timezone.json")
	if err != nil {
		return fmt.Errorf("read country_timezone.json err:%s", err.Error())
	}
	countryTimeZoneDataMap := make(map[string]CountryTimeZone)
	json.Unmarshal(countryTimeZoneData, &countryTimeZoneDataMap)
	MapCountryTimeZone = countryTimeZoneDataMap

	return nil
}

type CfgServer struct {
	Server struct {
		IsUpdate               bool              `json:"is_update" yaml:"is_update"`
		Net                    common.DasNetType `json:"net" yaml:"net"`
		HttpServerAddr         string            `json:"http_server_addr" yaml:"http_server_addr"`
		HttpServerInternalAddr string            `json:"http_server_internal_addr" yaml:"http_server_internal_addr"`
		PayServerAddress       string            `json:"pay_server_address" yaml:"pay_server_address"`
		PayPrivate             string            `json:"pay_private" yaml:"pay_private"`
		RemoteSignApiUrl       string            `json:"remote_sign_api_url" yaml:"remote_sign_api_url"`
		PushLogUrl             string            `json:"push_log_url" yaml:"push_log_url"`
		PushLogIndex           string            `json:"push_log_index" yaml:"push_log_index"`
		ParserUrl              string            `json:"parser_url" yaml:"parser_url"`
		TxToolSwitch           bool              `json:"tx_tool_switch" yaml:"tx_tool_switch"`
	} `json:"server" yaml:"server"`
	Origins []string `json:"origins" yaml:"origins"`
	Notify  struct {
		LarkErrorKey      string `json:"lark_error_key" yaml:"lark_error_key"`
		LarkRegisterKey   string `json:"lark_register_key" yaml:"lark_register_key"`
		LarkRegisterOkKey string `json:"lark_register_ok_key" yaml:"lark_register_ok_key"`
		DiscordWebhook    string `json:"discord_webhook" yaml:"discord_webhook"`
	} `json:"notify" yaml:"notify"`
	SendEmailConfig struct {
		From    string `json:"from" yaml:"from"`
		Charset string `json:"charset" yaml:"charset"`
		Subject string `json:"subject" yaml:"subject"`
	} `json:"send_email_config" yaml:"send_email_config"`
	Chain struct {
		CkbUrl             string `json:"ckb_url" yaml:"ckb_url"`
		IndexUrl           string `json:"index_url" yaml:"index_url"`
		CurrentBlockNumber uint64 `json:"current_block_number" yaml:"current_block_number"`
		ConfirmNum         uint64 `json:"confirm_num" yaml:"confirm_num"`
		ConcurrencyNum     uint64 `json:"concurrency_num" yaml:"concurrency_num"`
	} `json:"constant" yaml:"constant"`
	DB struct {
		Mysql       DbMysql `json:"mysql" yaml:"mysql"`
		ParserMysql DbMysql `json:"parser_mysql" yaml:"parser_mysql"`
	} `json:"db" yaml:"db"`
	Cache struct {
		Redis struct {
			Addr     string `json:"addr" yaml:"addr"`
			Password string `json:"password" yaml:"password"`
			DbNum    int    `json:"db_num" yaml:"db_num"`
		} `json:"redis" yaml:"redis"`
	} `json:"cache" yaml:"cache"`
	NOSQL struct {
		Elasticsearch struct {
			Addr string `json:"addr" yaml:"addr"`
		} `json:"elasticsearch" yaml:"elasticsearch"`
	} `json:"nosql" yaml:"nosql"`
	APOLLO struct {
		EndPoint string `json:"endpoint" yaml:"endpoint"`
	} `json:"apollo" yaml:"apollo"`
	Super struct {
		AccountMinLength     uint8           `json:"account_min_length" yaml:"account_min_length"`
		AccountMaxLength     uint8           `json:"account_max_length" yaml:"account_max_length"`
		OpenAccountMinLength uint8           `json:"open_account_min_length" yaml:"open_account_min_length"`
		OpenAccountMaxLength uint8           `json:"open_account_max_length" yaml:"open_account_max_length"`
		OpenAccountHashNum   uint32          `json:"open_account_hash_num" yaml:"open_account_hash_num"`
		MaxRegisterYears     int             `json:"max_register_years" yaml:"max_register_years"`
		Premium              decimal.Decimal `json:"premium" yaml:"premium"`
	} `json:"super" yaml:"super"`
	SuperLib struct {
		SuperArgs             string                            `json:"super_args" yaml:"super_args"`
		THQCodeHash           string                            `json:"thq_code_hash" yaml:"thq_code_hash"`
		SuperContractArgs     string                            `json:"super_contract_args" yaml:"super_contract_args"`
		SuperContractCodeHash string                            `json:"super_contract_code_hash" yaml:"super_contract_code_hash"`
		MapSuperContract      map[common.DasContractName]string `json:"map_super_contract" yaml:"map_super_contract"`
	} `json:"super_lib" yaml:"super_lib"`
	SupportedChains []ChainInfo `json:"supported_chains" yaml:"supported_chains"`
}

type DbMysql struct {
	Addr        string `json:"addr" yaml:"addr"`
	User        string `json:"user" yaml:"user"`
	Password    string `json:"password" yaml:"password"`
	DbName      string `json:"db_name" yaml:"db_name"`
	MaxOpenConn int    `json:"max_open_conn" yaml:"max_open_conn"`
	MaxIdleConn int    `json:"max_idle_conn" yaml:"max_idle_conn"`
}

type ChainInfo struct {
	ChainId    string `json:"chain_id" yaml:"chain_id"`
	GethUrl    string `json:"geth_url" yaml:"geth_url"`
	AlchemyUrl string `json:"alchemy_url" yaml:"alchemy_url"`
	Weth       string `json:"weth" yaml:"weth"`
}
