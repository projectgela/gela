package gelx

import (
	"github.com/projectgela/gela/contracts/gelx/contract"
	"github.com/projectgela/gela/log"
	"math/big"
	"strings"

	"github.com/projectgela/gela"
	"github.com/projectgela/gela/accounts/abi"
	"github.com/projectgela/gela/accounts/abi/bind/backends"
	"github.com/projectgela/gela/common"
	"github.com/projectgela/gela/consensus"
	"github.com/projectgela/gela/core/state"
)


// GetTokenAbi return token abi
func GetTokenAbi() (*abi.ABI, error) {
	contractABI, err := abi.JSON(strings.NewReader(contract.GRC21ABI))
	if err != nil {
		return nil, err
	}
	return &contractABI, nil
}

// RunContract run smart contract
func RunContract(chain consensus.ChainContext, statedb *state.StateDB, contractAddr common.Address, abi *abi.ABI, method string, args ...interface{}) (interface{}, error) {
	input, err := abi.Pack(method)
	if err != nil {
		return nil, err
	}
	backend := (*backends.SimulatedBackend)(nil)
	fakeCaller := common.HexToAddress("0x0000000000000000000000000000000000000001")
	msg := gela.CallMsg{To: &contractAddr, Data: input, From: fakeCaller}
	result, err := backend.CallContractWithState(msg, chain, statedb)
	if err != nil {
		return nil, err
	}
	var unpackResult interface{}
	err = abi.Unpack(&unpackResult, method, result)
	if err != nil {
		return nil, err
	}
	return unpackResult, nil
}

func (gelx *GelX) GetTokenDecimal(chain consensus.ChainContext, statedb *state.StateDB, tokenAddr common.Address) (*big.Int, error) {
	if tokenDecimal, ok := gelx.tokenDecimalCache.Get(tokenAddr); ok {
		return tokenDecimal.(*big.Int), nil
	}
	if tokenAddr.String() == common.GelaNativeAddress {
		gelx.tokenDecimalCache.Add(tokenAddr, common.BasePrice)
		return common.BasePrice, nil
	}
	var decimals uint8
	defer func() {
		log.Debug("GetTokenDecimal from ", "relayerSMC", common.RelayerRegistrationSMC, "tokenAddr", tokenAddr.Hex(), "decimals", decimals)
	}()
	contractABI, err := GetTokenAbi()
	if err != nil {
		return nil, err
	}
	stateCopy := statedb.Copy()
	result, err := RunContract(chain, stateCopy, tokenAddr, contractABI, "decimals")
	if err != nil {
		return nil, err
	}
	decimals = result.(uint8)

	tokenDecimal := new(big.Int).SetUint64(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	gelx.tokenDecimalCache.Add(tokenAddr, tokenDecimal)
	return tokenDecimal, nil
}
