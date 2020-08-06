package gelx

import (
	"github.com/projectgela/gela/accounts/abi/bind"
	"github.com/projectgela/gela/common"
	"github.com/projectgela/gela/contracts/gelx/contract"
	"math/big"
)

type MyGRC21 struct {
	*contract.MyGRC21Session
	contractBackend bind.ContractBackend
}

func NewGRC21(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*MyGRC21, error) {
	smartContract, err := contract.NewMyGRC21(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &MyGRC21{
		&contract.MyGRC21Session{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployGRC21(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, owners []common.Address, required *big.Int, name string, symbol string, decimals uint8, cap, fee, depositFee, withdrawFee *big.Int) (common.Address, *MyGRC21, error) {
	contractAddr, _, _, err := contract.DeployMyGRC21(transactOpts, contractBackend, owners, required, name, symbol, decimals, cap, fee, depositFee, withdrawFee)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewGRC21(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
