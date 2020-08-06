package utils

import (
	"github.com/projectgela/gela/eth"
	"github.com/projectgela/gela/eth/downloader"
	"github.com/projectgela/gela/ethstats"
	"github.com/projectgela/gela/gelx"
	gelxlending "github.com/projectgela/gela/gelxlending"
	"github.com/projectgela/gela/les"
	"github.com/projectgela/gela/node"
	whisper "github.com/projectgela/gela/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var gelXServ *gelx.GelX
			ctx.Service(&gelXServ)
			var lendingServ *gelxlending.Lending
			ctx.Service(&lendingServ)
			fullNode, err := eth.New(ctx, cfg, gelXServ, lendingServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterGelXService(stack *node.Node, cfg *gelx.Config) {
	gelX := gelx.New(cfg)
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return gelX, nil
	}); err != nil {
		Fatalf("Failed to register the GelX service: %v", err)
	}

	// register gelxlending service
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return gelxlending.New(gelX), nil
	}); err != nil {
		Fatalf("Failed to register the GelXLending service: %v", err)
	}
}
