package indexer

import "testing"

const (
	L1SepoliaUrl       = "http://10.3.3.234:8545"
	OKX1SepoliaL2Url   = "https://testrpc.x1.tech" //"https://x1testrpc.okx.com/"
	OKX1SepoliaLocalDB = "postgres://jiaxingsun@localhost:5432/okx1_01"
)

// indexL1BatchesForOKX1Sepolia
func TestIndexL1BatchesForOKX1Sepolia(t *testing.T) {
	indexer, err := InitINdexer(L1SepoliaUrl, OKX1SepoliaL2Url, OKX1SepoliaLocalDB, "okx1-sepolia")
	if err != nil {
		t.Fatalf("Error InitINdexer : %s", err)
	}

	batchNumbers, batchesLen, err := indexer.indexL1BatchesForOKX1Sepolia(4653167, 4653367)
	if err != nil {
		t.Fatalf("Err indexL1BatchesForOKX1Sepolia:%v", err)
	}

	t.Logf("batchesLen %d batchNumbers %v ", batchesLen, batchNumbers)
}
