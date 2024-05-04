package indexer

import (
	"github.com/unifra20/l2scan-indexer/config"
	"github.com/unifra20/l2scan-indexer/db"
	"github.com/unifra20/l2scan-indexer/ethclient"
	"testing"
)

func TestRecovery(t *testing.T) {
	conf := config.LoadConfig()

	// Initialize blockchain client
	l1client, err := ethclient.NewClient(conf.L1RPC)
	if err != nil {
		t.Errorf("Error initializing RPC client: %s", err)
	}
	l2client, err := ethclient.NewClient(conf.L2RPC)
	if err != nil {
		t.Errorf("Error initializing RPC client: %s", err)
	}

	// Initialize database client
	dbClient, err := db.NewClient(conf.Pgdsn)
	if err != nil {
		t.Errorf("Error initializing Postgres client: %s", err)
	}
	indexer := NewIndexer(GetChain(conf.Chain), l1client, l2client, dbClient)

	indexer.L2RecoveryExecutor(2, false, "", false)

}

// findNextAbsentGreater
//func TestFindNextAbsentGreater(t *testing.T) {
//	a := []int{3, 4, 6, 7, 8, 9}
//	b := 8
//
//	res, boo := findNextAbsentSmaller(a, b)
//	t.Log("res:", res, boo)
//}
