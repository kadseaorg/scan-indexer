package db

import (
	"context"
	"encoding/hex"
	"encoding/json"

	"github.com/jackc/pgx/v5"
)

const inscriptionDataPrefixHex = "0x646174613a2c"
const inscriptionSyncedKey = "inscription_synced_block_number"

func (db *Client) GetLatestSyncedInscriptionBlockNumber(ctx context.Context) (uint64, error) {
	return db.GetSyncProgress(ctx, inscriptionSyncedKey)
}

func (db *Client) UpsertLastetSyncedInscriptionBlockNumber(ctx context.Context, blockNumber uint64) error {
	return db.UpdateSyncProgress(ctx, inscriptionSyncedKey, blockNumber)
}

func (db *Client) ExtractInscriptions(startBlockNumber uint64, endBlockNumber uint64) (int, error) {
	rows, err := db.pool.Query(context.Background(),
		`SELECT id, hash, transaction_index, from_address, to_address, block_number, gas_price, input, timestamp
		FROM transactions WHERE input LIKE $1 AND block_number BETWEEN $2 AND $3`, inscriptionDataPrefixHex+"%", startBlockNumber, endBlockNumber)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	batch := &pgx.Batch{}

	var inscriptionsCount int
	for rows.Next() {
		var (
			id               uint64
			hash             string
			transactionIndex int
			fromAddress      string
			toAddress        string
			blockNumber      uint64
			gasPrice         uint64
			input            string
			timestamp        uint64
		)

		err := rows.Scan(&id, &hash, &transactionIndex, &fromAddress, &toAddress, &blockNumber, &gasPrice, &input, &timestamp)
		if err != nil {
			return inscriptionsCount, err
		}

		fullInscription, err := HexToAscii(input)
		if err != nil {
			// log.Warn().Msgf("[🪽 Inscription Indexer] Failed to decode inscription data for transaction %s, input: %s", hash, input)
			continue
		}

		if !IsValidJSON(fullInscription) {
			// log.Warn().Msgf("[🪽 Inscription Indexer] Invalid JSON inscription data for transaction %s, input: %s, fullInscription: %s", hash, input, fullInscription)
			continue
		}

		batch.Queue(`INSERT INTO inscriptions (transaction_hash, transaction_index, from_address, to_address, block_number, gas_price, full_inscription, timestamp)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING`, hash, transactionIndex, fromAddress, toAddress, blockNumber, gasPrice, fullInscription, timestamp)

		// log.Info().Msgf("[🪽 Inscription Indexer] Inserted inscription for transaction %s, inscription: %s", hash, fullInscription)

		// increment the count of processed inscriptions
		inscriptionsCount++
	}

	if err = rows.Err(); err != nil {
		return inscriptionsCount, err
	}

	br := db.pool.SendBatch(context.Background(), batch)
	if err := br.Close(); err != nil {
		return inscriptionsCount, err
	}

	return inscriptionsCount, nil
}

func HexToAscii(hexStr string) (string, error) {
	decoded, err := hex.DecodeString(hexStr[2:])
	if err != nil {
		return "", err
	}
	return string(decoded)[6:], nil // remove "data:," prefix
}

func IsValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
