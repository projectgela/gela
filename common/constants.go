package common

import (
	"math/big"
)

const (
	RewardMasterPercent        = 60
	RewardVoterPercent         = 30
	RewardFoundationPercent    = 10
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 800
	EpocBlockOpening           = 850
	EpocBlockRandomize         = 900
	MaxMasternodes             = 150
	LimitPenaltyEpoch          = 4
	BlocksPerYear              = uint64(15768000)
	LimitThresholdNonceInQueue = 10
	DefaultMinGasPrice         = 250000000
	MergeSignRange             = 15
	RangeReturnSigner          = 150
	MinimunMinerBlockPerEpoch  = 1
	IgnoreSignerCheckBlock     = uint64(14458500)
	OneYear                    = uint64(365 * 86400)
	LiquidateLendingTradeBlock = uint64(100)
)

var Rewound = uint64(0)

// hardforks
var TIP2019Block = big.NewInt(0)
var TIPSigning = big.NewInt(0)
var TIPRandomize = big.NewInt(0)
var BlackListHFNumber = uint64(0)
var TIPGelX = big.NewInt(0)
var TIPGelXTestnet = big.NewInt(0)
var TIPGelXLending = big.NewInt(0)

var IsTestnet bool = false
var StoreRewardFolder string
var RollbackHash Hash
var BasePrice = big.NewInt(1000000000000000000)                         // 1
var RelayerLockedFund = big.NewInt(20000)                               // 20000 GEL
var RelayerFee = big.NewInt(1000000000000000)                           // 0.001
var GelXBaseFee = big.NewInt(10000)                                    // 1 / GelXBaseFee
var RelayerCancelFee = big.NewInt(100000000000000)                      // 0.0001
var GelXBaseCancelFee = new(big.Int).Mul(GelXBaseFee, big.NewInt(10)) // 1/ (GelXBaseFee *10)
var RelayerLendingFee = big.NewInt(10000000000000000)                   // 0.01
var RelayerLendingCancelFee = big.NewInt(1000000000000000)              // 0.001
var BaseLendingInterest = big.NewInt(100000000)                         // 1e8

var MinGasPrice = big.NewInt(DefaultMinGasPrice)
var RelayerRegistrationSMC = ""
var RelayerRegistrationSMCTestnet = ""
var LendingRegistrationSMC = ""
var LendingRegistrationSMCTestnet = ""
var GRC21IssuerSMCTestNet = HexToAddress("")
var GRC21IssuerSMC = HexToAddress("")
var GelXListingSMC = HexToAddress("")
var GelXListingSMCTestNet = HexToAddress("")
var GRC21GasPriceBefore = big.NewInt(2500)
var GRC21GasPrice = big.NewInt(250000000)
var RateTopUp = big.NewInt(90) // 90%
var BaseTopUp = big.NewInt(100)
var BaseRecall = big.NewInt(100)
var Blacklist = map[Address]bool{
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
	HexToAddress(""): true,
}
var TIPGRC21Fee = big.NewInt(0)
var LimitTimeFinality = uint64(30) // limit in 30 block
