package gelx

import (
	"github.com/projectgela/gela/accounts/abi/bind"
	"github.com/projectgela/gela/common"
	"github.com/projectgela/gela/contracts/gelx/contract"
	"math/big"
)

type GRC21Issuer struct {
	*contract.GRC21IssuerSession
	contractBackend bind.ContractBackend
}

func NewGRC21Issuer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*GRC21Issuer, error) {
	contractObject, err := contract.NewGRC21Issuer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &GRC21Issuer{
		&contract.GRC21IssuerSession{
			Contract:     contractObject,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployGRC21Issuer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, minApply *big.Int) (common.Address, *GRC21Issuer, error) {
	contractAddr, _, _, err := contract.DeployGRC21Issuer(transactOpts, contractBackend, minApply)
	if err != nil {
		return contractAddr, nil, err
	}
	contractObject, err := NewGRC21Issuer(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, contractObject, nil
}
