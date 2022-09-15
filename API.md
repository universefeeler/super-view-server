 * [Query API LIST](#query-api-list)
    * [Token List](#token-list)
 * [OPERATE API LIST](#operate-api-list)
    * [Reverse Declare](#reverse-declare)

### Query API LIST

#### Token List

**Request**

* path: /token/list
* param: none

**Response**

```json
{
  "err_no": 0,
  "err_msg": "",
  "data": {
    "token_list": [
      {
        "token_id": "ckb_ckb",
        "chain_type": 0,
        "contract": "",
        "name": "Nervos Network",
        "symbol": "CKB",
        "decimals": 8,
        "logo": "https://app.da.systems/images/components/portal-wallet.svg",
        "price": "0.01850608"
      },
      {
        "token_id": "polygon_matic",
        "chain_type": 1,
        "contract": "",
        "name": "Polygon",
        "symbol": "MATIC",
        "decimals": 18,
        "logo": "https://app.da.systems/images/components/polygon.svg",
        "price": "2.15"
      },
      {
        "token_id": "bsc_bnb",
        "chain_type": 5,
        "contract": "",
        "name": "Binance",
        "symbol": "BNB",
        "decimals": 18,
        "logo": "https://app.da.systems/images/components/binance-smart-constant.svg",
        "price": "435.85"
      },
      {
        "token_id": "wx_cny",
        "chain_type": 4,
        "contract": "",
        "name": "WeChat Pay",
        "symbol": "¥",
        "decimals": 2,
        "logo": "https://app.da.systems/images/components/wechat_pay.png",
        "price": "0.1569"
      },
      {
        "token_id": "tron_trx",
        "chain_type": 3,
        "contract": "",
        "name": "TRON",
        "symbol": "TRX",
        "decimals": 6,
        "logo": "https://app.da.systems/images/components/tron.svg",
        "price": "0.064233"
      },
      {
        "token_id": "btc_btc",
        "chain_type": 2,
        "contract": "",
        "name": "Bitcoin",
        "symbol": "BTC",
        "decimals": 8,
        "logo": "https://app.da.systems/images/components/bitcoin.svg",
        "price": "42161"
      },
      {
        "token_id": "eth_eth",
        "chain_type": 1,
        "contract": "",
        "name": "Ethereum",
        "symbol": "ETH",
        "decimals": 18,
        "logo": "https://app.da.systems/images/components/ethereum.svg",
        "price": "3115.47"
      }
    ]
  }
}
```

**Usage**

```curl
curl -X POST http://127.0.0.1:8120/v1/token/list
```

### OPERATE API LIST

#### Reverse Declare

**Request**

* path: /reverse/declare
* param:
    * evm_chain_id: eth-1/5 bsc-56/97 polygon-137/8001

```json
{
  "chain_type": 1,
  "address": "0xc9f53b1d85356b60453f867610888d89a0b667ad",
  "account": "aaaaa.bit",
  "evm_chain_id": 5
}
```

**Response**

```json
{
  "err_no": 0,
  "err_msg": "",
  "data": {
    "sign_key": "",
    "sign_list": [
      {
        "sign_type": 5,
        "sign_msg": ""
      }
    ],
    "mm_json": {}
  }
}
```

**Usage**

```curl
curl -X POST http://127.0.0.1:8120/v1/reverse/declare -d'{"chain_type":1,"address":"0xc9f53b1d85356b60453f867610888d89a0b667ad","account":"aaaa.bit","evm_chain_id":5}'
```

