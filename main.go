package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/unifra20/l2scan-indexer/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/unifra20/l2scan-indexer/db"
	"github.com/unifra20/l2scan-indexer/ethclient"
	"github.com/unifra20/l2scan-indexer/indexer"
)

func main() {
	overwriteFlag := flag.Int64("overwrite", 0, "overwrite block number")

	// Parse flags.
	flag.Parse()

	// zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	conf := config.LoadConfig()

	// Initialize blockchain client
	l1client, err := ethclient.NewClient(conf.L1RPC)
	if err != nil {
		if conf.Chain != indexer.Kadsea.Name && conf.Chain != indexer.KadseaTestnet.Name {
			log.Fatal().Msgf("Error initializing L1 RPC client: %s", err)
		}
	}
	l2client, err := ethclient.NewClient(conf.L2RPC)
	if err != nil {
		log.Fatal().Msgf("Error initializing L2 RPC client: %s", err)
	}
	var debugClient *ethclient.Client
	if conf.DebugRPC != nil {
		debugClient, err = ethclient.NewClient(*conf.DebugRPC)
		if err != nil {
			log.Fatal().Msgf("Error initializing Debug RPC client: %s", err)
		}
	}

	// Initialize database client
	dbClient, err := db.NewClient(conf.Pgdsn)
	if err != nil {
		log.Fatal().Msgf("Error initializing Postgres client: %s", err)
	}

	//check if felid not exist
	dbClient.CreateFieldsL1BlockNumberForL1Batches()

	// Initialize indexer
	indexer := indexer.NewIndexer(indexer.GetChain(conf.Chain), l1client, l2client, debugClient, dbClient)

	if overwriteFlag != nil && *overwriteFlag > 0 {
		indexer.FixOverwriteBlock(*overwriteFlag)
		log.Info().Msgf("🌈 [Fix Overwrite Block] Finished")
		os.Exit(0)
	}

	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	// Notify the sigCh channel when the program receives the interrupt (Ctrl+C) or termination signal.
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-sigCh
		cancel()
	}()

	// Start indexing
	indexer.Start(ctx,
		conf.Worker,
		conf.L1ForceStartBlock,
		conf.L2ForceStartBlock,
		conf.CheckMisMatchedBlocks,
		conf.L1BridgeForceStartBlock,
		conf.InscriptionForceStartBlock,
		conf.RecoveryStartBlock,
	)
}
