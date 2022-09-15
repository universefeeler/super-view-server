package timer

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"sync"
)

var (
	randomLock sync.RWMutex
)

func (t *TxTimer) doRandomPrivateKeyView() error {
	randomLock.Lock()
	defer randomLock.Unlock()
	return preExecuteRandomView()
}

func preExecuteRandomView() error {
	counter := uint16(0)
	for {
		// 生成随机私钥
		privateKey, err1 := crypto.GenerateKey()
		if err1 != nil {
			log.Errorf("err:", err1.Error())
		}
		// 使用FromECDSA方法将其转换为字节
		privateKeyBytes := crypto.FromECDSA(privateKey)
		// 转换为可读16进制
		privateKeyHex := fmt.Sprintf(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
		log.Info("privateKeyHex:", privateKeyHex)
		// 私钥派生出公钥
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Error("err:", "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		publicKeyRaw := fmt.Sprintf(hexutil.Encode(publicKeyBytes)[4:])
		log.Info("publicKeyRaw:", publicKeyRaw)
		// 公共地址其实就是公钥的Keccak-256哈希，然后我们取最后40个字符（20个字节）并用“0x”作为前缀
		// keccak256算法则可以将任意长度的输入压缩成64位16进制的数，且哈希碰撞的概率近乎为0.
		// Keccak-256 is the winning SHA-3 proposal hash function from the Keccak team
		// https://keccak-256.cloxy.net/
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		log.Info("address:", address)

		line := fmt.Sprintf("%v#%v", privateKeyHex, address)

		ExecuteLineByDebankOpenApiWithRandom(line, counter)

		//ExecuteLineByDebankOpenApi(line, counter)
	}
}
