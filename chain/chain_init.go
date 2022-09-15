package chain

var (
	Evm_Chain_Map_Prod   map[string]string
	Chain_Map_For_Debank map[string]string
	Chain_Name_To_Id     map[string]string
)

const (
	EthereumMainnetId          = "1"
	PolygonMainnetId           = "137"
	BinanceSmartChainMainnetId = "56"

	EthereumChainForDebank          = "eth"
	PolygonChainForDebank           = "matic"
	BinanceSmartChainChainForDebank = "bsc"
)

// https://github.com/ethereum-lists/chains
func init() {
	Evm_Chain_Map_Prod = make(map[string]string)
	//Eth
	Evm_Chain_Map_Prod[EthereumMainnetId] = "{\"name\":\"Ethereum Mainnet\",\"constant\":\"ETH\",\"icon\":\"ethereum\",\"rpc\":[\"https://mainnet.infura.io/v3/${INFURA_API_KEY}\",\"wss://mainnet.infura.io/ws/v3/${INFURA_API_KEY}\",\"https://api.mycryptoapi.com/eth\",\"https://cloudflare-eth.com\"],\"faucets\":[],\"nativeCurrency\":{\"name\":\"Ether\",\"symbol\":\"ETH\",\"decimals\":18},\"infoURL\":\"https://ethereum.org\",\"shortName\":\"eth\",\"chainId\":1,\"networkId\":1,\"slip44\":60,\"ens\":{\"registry\":\"0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e\"},\"explorers\":[{\"name\":\"etherscan\",\"url\":\"https://etherscan.io\",\"standard\":\"EIP3091\"}]}"
	//Polyogn
	Evm_Chain_Map_Prod[PolygonMainnetId] = "{\"name\":\"Polygon Mainnet\",\"constant\":\"Polygon\",\"rpc\":[\"https://polygon-rpc.com/\",\"https://rpc-mainnet.matic.network\",\"https://matic-mainnet.chainstacklabs.com\",\"https://rpc-mainnet.maticvigil.com\",\"https://rpc-mainnet.matic.quiknode.pro\",\"https://matic-mainnet-full-rpc.bwarelabs.com\"],\"faucets\":[],\"nativeCurrency\":{\"name\":\"MATIC\",\"symbol\":\"MATIC\",\"decimals\":18},\"infoURL\":\"https://polygon.technology/\",\"shortName\":\"MATIC\",\"chainId\":137,\"networkId\":137,\"slip44\":966,\"explorers\":[{\"name\":\"polygonscan\",\"url\":\"https://polygonscan.com\",\"standard\":\"EIP3091\"}]}{\"name\":\"Polygon Mainnet\",\"constant\":\"Polygon\",\"rpc\":[\"https://polygon-rpc.com/\",\"https://rpc-mainnet.matic.network\",\"https://matic-mainnet.chainstacklabs.com\",\"https://rpc-mainnet.maticvigil.com\",\"https://rpc-mainnet.matic.quiknode.pro\",\"https://matic-mainnet-full-rpc.bwarelabs.com\"],\"faucets\":[],\"nativeCurrency\":{\"name\":\"MATIC\",\"symbol\":\"MATIC\",\"decimals\":18},\"infoURL\":\"https://polygon.technology/\",\"shortName\":\"MATIC\",\"chainId\":137,\"networkId\":137,\"slip44\":966,\"explorers\":[{\"name\":\"polygonscan\",\"url\":\"https://polygonscan.com\",\"standard\":\"EIP3091\"}]}"
	//Bsc
	Evm_Chain_Map_Prod[BinanceSmartChainMainnetId] = "{\"name\":\"Binance Smart Chain Mainnet\",\"chain\":\"BSC\",\"rpc\":[\"https://bsc-dataseed1.binance.org\",\"https://bsc-dataseed2.binance.org\",\"https://bsc-dataseed3.binance.org\",\"https://bsc-dataseed4.binance.org\",\"https://bsc-dataseed1.defibit.io\",\"https://bsc-dataseed2.defibit.io\",\"https://bsc-dataseed3.defibit.io\",\"https://bsc-dataseed4.defibit.io\",\"https://bsc-dataseed1.ninicoin.io\",\"https://bsc-dataseed2.ninicoin.io\",\"https://bsc-dataseed3.ninicoin.io\",\"https://bsc-dataseed4.ninicoin.io\",\"wss://bsc-ws-node.nariox.org\"],\"faucets\":[\"https://free-online-app.com/faucet-for-eth-evm-chains/\"],\"nativeCurrency\":{\"name\":\"Binance Chain Native Token\",\"symbol\":\"BNB\",\"decimals\":18},\"infoURL\":\"https://www.binance.org\",\"shortName\":\"bnb\",\"chainId\":56,\"networkId\":56,\"slip44\":714,\"explorers\":[{\"name\":\"bscscan\",\"url\":\"https://bscscan.com\",\"standard\":\"EIP3091\"}]}"

	/**
	 * Just for Debank
	 */
	Chain_Map_For_Debank = make(map[string]string)
	//Eth
	Chain_Map_For_Debank[EthereumChainForDebank] = "{\"name\":\"Ethereum Mainnet\",\"constant\":\"ETH\",\"icon\":\"ethereum\",\"rpc\":[\"https://mainnet.infura.io/v3/${INFURA_API_KEY}\",\"wss://mainnet.infura.io/ws/v3/${INFURA_API_KEY}\",\"https://api.mycryptoapi.com/eth\",\"https://cloudflare-eth.com\"],\"faucets\":[],\"nativeCurrency\":{\"name\":\"Ether\",\"symbol\":\"ETH\",\"decimals\":18},\"infoURL\":\"https://ethereum.org\",\"shortName\":\"eth\",\"chainId\":1,\"networkId\":1,\"slip44\":60,\"ens\":{\"registry\":\"0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e\"},\"explorers\":[{\"name\":\"etherscan\",\"url\":\"https://etherscan.io\",\"standard\":\"EIP3091\"}]}"
	//Matic
	Chain_Map_For_Debank[PolygonChainForDebank] = "{\"name\":\"Polygon Mainnet\",\"constant\":\"Polygon\",\"rpc\":[\"https://polygon-rpc.com/\",\"https://rpc-mainnet.matic.network\",\"https://matic-mainnet.chainstacklabs.com\",\"https://rpc-mainnet.maticvigil.com\",\"https://rpc-mainnet.matic.quiknode.pro\",\"https://matic-mainnet-full-rpc.bwarelabs.com\"],\"faucets\":[],\"nativeCurrency\":{\"name\":\"MATIC\",\"symbol\":\"MATIC\",\"decimals\":18},\"infoURL\":\"https://polygon.technology/\",\"shortName\":\"MATIC\",\"chainId\":137,\"networkId\":137,\"slip44\":966,\"explorers\":[{\"name\":\"polygonscan\",\"url\":\"https://polygonscan.com\",\"standard\":\"EIP3091\"}]}{\"name\":\"Polygon Mainnet\",\"constant\":\"Polygon\",\"rpc\":[\"https://polygon-rpc.com/\",\"https://rpc-mainnet.matic.network\",\"https://matic-mainnet.chainstacklabs.com\",\"https://rpc-mainnet.maticvigil.com\",\"https://rpc-mainnet.matic.quiknode.pro\",\"https://matic-mainnet-full-rpc.bwarelabs.com\"],\"faucets\":[],\"nativeCurrency\":{\"name\":\"MATIC\",\"symbol\":\"MATIC\",\"decimals\":18},\"infoURL\":\"https://polygon.technology/\",\"shortName\":\"MATIC\",\"chainId\":137,\"networkId\":137,\"slip44\":966,\"explorers\":[{\"name\":\"polygonscan\",\"url\":\"https://polygonscan.com\",\"standard\":\"EIP3091\"}]}"
	//Bsc
	Chain_Map_For_Debank[BinanceSmartChainChainForDebank] = "{\"name\":\"Binance Smart Chain Mainnet\",\"chain\":\"BSC\",\"rpc\":[\"https://bsc-dataseed1.binance.org\",\"https://bsc-dataseed2.binance.org\",\"https://bsc-dataseed3.binance.org\",\"https://bsc-dataseed4.binance.org\",\"https://bsc-dataseed1.defibit.io\",\"https://bsc-dataseed2.defibit.io\",\"https://bsc-dataseed3.defibit.io\",\"https://bsc-dataseed4.defibit.io\",\"https://bsc-dataseed1.ninicoin.io\",\"https://bsc-dataseed2.ninicoin.io\",\"https://bsc-dataseed3.ninicoin.io\",\"https://bsc-dataseed4.ninicoin.io\",\"wss://bsc-ws-node.nariox.org\"],\"faucets\":[\"https://free-online-app.com/faucet-for-eth-evm-chains/\"],\"nativeCurrency\":{\"name\":\"Binance Chain Native Token\",\"symbol\":\"BNB\",\"decimals\":18},\"infoURL\":\"https://www.binance.org\",\"shortName\":\"bnb\",\"chainId\":56,\"networkId\":56,\"slip44\":714,\"explorers\":[{\"name\":\"bscscan\",\"url\":\"https://bscscan.com\",\"standard\":\"EIP3091\"}]}"

	/**
	 * Name mapping Id
	 */
	Chain_Name_To_Id = make(map[string]string)
	//Eth
	Chain_Name_To_Id[EthereumChainForDebank] = EthereumMainnetId
	//Polyogn
	Chain_Name_To_Id[PolygonChainForDebank] = PolygonMainnetId
	//Bsc
	Chain_Name_To_Id[BinanceSmartChainChainForDebank] = BinanceSmartChainMainnetId

}
