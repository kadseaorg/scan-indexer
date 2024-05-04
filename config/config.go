package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Chain                      string
	L1RPC                      string
	L2RPC                      string
	DebugRPC                   *string
	Pgdsn                      string
	Worker                     int
	L1ForceStartBlock          *uint64 // optional, if set, start from this block number
	L2ForceStartBlock          *uint64 // optional, if set, start from this block number
	RecoveryStartBlock         *uint64 // optional, if set, recovery start from this block number
	L1BridgeForceStartBlock    *uint64 // optional, if set, start from this block number
	InscriptionForceStartBlock *uint64 // optional, if set, start from this block number
	CheckMisMatchedBlocks      bool    // optional, if set, check mismatched blocks
}

func LoadConfig() Config {
	// Set up Viper to read from environment variables
	viper.AutomaticEnv()
	// optional, read from .env file
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// Read configuration values
	cfg := Config{
		Chain:  viper.GetString("CHAIN"),
		L1RPC:  viper.GetString("L1_RPC"),
		L2RPC:  viper.GetString("L2_RPC"),
		Pgdsn:  viper.GetString("PGDSN"),
		Worker: viper.GetInt("WORKER"),
	}
	if viper.IsSet("L1_FORCE_START_BLOCK") {
		l1ForceStartBlock := viper.GetUint64("L1_FORCE_START_BLOCK")
		cfg.L1ForceStartBlock = &l1ForceStartBlock
	}
	if viper.IsSet("L2_FORCE_START_BLOCK") {
		l2ForceStartBlock := viper.GetUint64("L2_FORCE_START_BLOCK")
		cfg.L2ForceStartBlock = &l2ForceStartBlock
	}
	if viper.IsSet("RECOVERY_START_BLOCK") {
		recoveryStartBlock := viper.GetUint64("RECOVERY_START_BLOCK")
		cfg.RecoveryStartBlock = &recoveryStartBlock
	}
	if viper.IsSet("L1_BRIDGE_FORCE_START_BLOCK") {
		l1BridgeForceStartBlock := viper.GetUint64("L1_BRIDGE_FORCE_START_BLOCK")
		cfg.L1BridgeForceStartBlock = &l1BridgeForceStartBlock
	}
	if viper.IsSet("INSCRIPTION_FORCE_START_BLOCK") {
		inscriptionForceStartBlock := viper.GetUint64("INSCRIPTION_FORCE_START_BLOCK")
		cfg.InscriptionForceStartBlock = &inscriptionForceStartBlock
	}
	if viper.IsSet("CHECK_MISMATCHED_BLOCKS") {
		cfg.CheckMisMatchedBlocks = viper.GetBool("CHECK_MISMATCHED_BLOCKS")
	}
	if viper.IsSet("DEBUG_RPC") {
		debugRPC := viper.GetString("DEBUG_RPC")
		cfg.DebugRPC = &debugRPC
	}

	log.Info().Msgf("Loaded config: %+v", cfg)

	return cfg
}
