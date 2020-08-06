package gelx

import (
	"github.com/projectgela/gela/accounts/abi/bind"
	"github.com/projectgela/gela/common"
	"github.com/projectgela/gela/contracts/gelx/contract"
)

type GELXListing struct {
	*contract.GELXListingSession
	contractBackend bind.ContractBackend
}

func NewMyGELXListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*GELXListing, error) {
	smartContract, err := contract.NewGELXListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &GELXListing{
		&contract.GELXListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployGELXListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *GELXListing, error) {
	contractAddr, _, _, err := contract.DeployGELXListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyGELXListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
