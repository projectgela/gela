package simulation

import (
	"math/big"
	"os"

	"github.com/projectgela/gela/common"
	"github.com/projectgela/gela/crypto"
)

var (
	BaseGEL     = big.NewInt(0).Mul(big.NewInt(10), big.NewInt(100000000000000000)) // 1 GEL
	RpcEndpoint = "http://127.0.0.1:8501/"
	MainKey, _  = crypto.HexToECDSA(os.Getenv("MAIN_ADDRESS_KEY"))
	MainAddr    = crypto.PubkeyToAddress(MainKey.PublicKey) //0x17F2beD710ba50Ed27aEa52fc4bD7Bda5ED4a037

	// GRC21 Token
	MinGRC21Apply  = big.NewInt(0).Mul(big.NewInt(10), BaseGEL) // 10 GEL
	GRC21TokenCap  = big.NewInt(0).Mul(big.NewInt(1000000000000), BaseGEL)
	GRC21TokenFee  = big.NewInt(100)
	GelXListingFee = big.NewInt(0).Mul(big.NewInt(1000), BaseGEL) // 1000 GEL

	// GELX
	MaxRelayers               = big.NewInt(200)
	MaxTokenList              = big.NewInt(200)
	MinDeposit                = big.NewInt(0).Mul(big.NewInt(25000), BaseGEL) // 25000 GEL
	CollateralDepositRate     = big.NewInt(150)
	CollateralLiquidationRate = big.NewInt(110)
	CollateralRecallRate      = big.NewInt(200)
	TradeFee                  = uint16(10)  // trade fee decimals 10^4
	LendingTradeFee           = uint16(100) // lending trade fee decimals 10^4
	// 1m , 1d,7d,30d
	Terms                 = []*big.Int{big.NewInt(60), big.NewInt(86400), big.NewInt(7 * 86400), big.NewInt(30 * 86400)}
	RelayerCoinbaseKey, _ = crypto.HexToECDSA(os.Getenv("RELAYER_COINBASE_KEY")) //
	RelayerCoinbaseAddr   = crypto.PubkeyToAddress(RelayerCoinbaseKey.PublicKey) // 0x0D3ab14BBaD3D99F4203bd7a11aCB94882050E7e

	OwnerRelayerKey, _ = crypto.HexToECDSA(os.Getenv("RELAYER_OWNER_KEY"))
	OwnerRelayerAddr   = crypto.PubkeyToAddress(OwnerRelayerKey.PublicKey) //0x703c4b2bD70c169f5717101CaeE543299Fc946C7

	GELANative = common.HexToAddress("0x0000000000000000000000000000000000000001")

	TokenNameList = []string{"BTC", "ETH", "XRP", "LTC", "BNB", "ADA", "ETC", "BCH", "EOS", "USDT"}
	TeamAddresses = []common.Address{
		common.HexToAddress("0x8fB1047e874d2e472cd08980FF8383053dd83102"), // MM team
		common.HexToAddress("0x9ca1514E3Dc4059C29a1608AE3a3E3fd35900888"), // MM team
		common.HexToAddress("0x15e08dE16f534c890828F2a0D935433aF5B3CE0C"), // CTO
		common.HexToAddress("0xb68D825655F2fE14C32558cDf950b45beF18D218"), // DEX team
		common.HexToAddress("0xF7349C253FF7747Df661296E0859c44e974fb52E"), // HaiDV
		common.HexToAddress("0x9f6b8fDD3733B099A91B6D70CDC7963ebBbd2684"), // Can
		common.HexToAddress("0x06605B28aab9835be75ca242a8aE58f2e15F2F45"), // Nien
		common.HexToAddress("0x33c2E732ae7dce8B05F37B2ba0CFe14c980c4Dbe"), // Vu Pham
		common.HexToAddress("0x16a73f3a64eca79e117258e66dfd7071cc8312a9"), // BTCGEL
		common.HexToAddress("0xac177441ac2237b2f79ecff1b8f6bca39e27ef9f"), // ETHGEL
		common.HexToAddress("0x4215250e55984c75bbce8ae639b86a6cad8ec126"), // XRPGEL
		common.HexToAddress("0x6b70ca959814866dd5c426d63d47dde9cc6c32d2"), // LTCGEL
		common.HexToAddress("0x33df079fe9b9cd7fb23a1085e4eaaa8eb6952cb3"), // BNBGEL
		common.HexToAddress("0x3cab8292137804688714670640d19f9d7a60c472"), // ADAGEL
		common.HexToAddress("0x9415d953d47c5f155cac9de7b24a756f352eafbf"), // ETCGEL
		common.HexToAddress("0xe32d2e7c8e8809e45c8e2332830b48d9e231e3f2"), // BCHGEL
		common.HexToAddress("0xf76ddbda664ea47088937e1cf9ff15036714dee3"), // EOSGEL
		common.HexToAddress("0xc465ee82440dada9509feb235c7cd7d896acf13c"), // ETHBTC
		common.HexToAddress("0xb95bdc136c579dc3fd2b2424a8e925a90228d2c2"), // XRPBTC
		common.HexToAddress("0xe36c1842365595D44854eEcd64B11c8115E133EF"), // USDGEL
		common.HexToAddress("0xaaC1959F6F0fb539F653409079Ec4146267B7555"), // BTCUSD
		common.HexToAddress("0x726DA688e2e09f01A2e1aB4c10F25B7CEdD4a0f3"), // ETHUSD
	}

	Required = big.NewInt(2)
	Owners   = []common.Address{
		common.HexToAddress("0x244e17B2141288a6F00E79E8feC2341f827d156f"),
		common.HexToAddress("0xd106159eC58BD2EAf5B62eF4e9cDb286170B0Bb9"),
		common.HexToAddress("0x0197BE034Bf0Bd2b3adDC84366a5681Bb7545888"),
	}
)
