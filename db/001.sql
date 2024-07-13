/*
 Navicat Premium Data Transfer

 Source Server         : scan-index
 Source Server Type    : PostgreSQL
 Source Server Version : 160003 (160003)
 Source Host           : database-2.cbus4wcc6csb.ap-southeast-1.rds.amazonaws.com:5432
 Source Catalog        : postgres
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160003 (160003)
 File Encoding         : 65001

 Date: 26/06/2024 10:10:26
*/


-- ----------------------------
-- Sequence structure for account_watch_list_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."account_watch_list_id_seq";
CREATE SEQUENCE "public"."account_watch_list_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."account_watch_list_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for accounts_list_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."accounts_list_id_seq";
CREATE SEQUENCE "public"."accounts_list_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."accounts_list_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for address_balances_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."address_balances_id_seq";
CREATE SEQUENCE "public"."address_balances_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."address_balances_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for address_to_labels_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."address_to_labels_id_seq";
CREATE SEQUENCE "public"."address_to_labels_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."address_to_labels_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for bridge_stats_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."bridge_stats_id_seq";
CREATE SEQUENCE "public"."bridge_stats_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."bridge_stats_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for contracts_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."contracts_id_seq";
CREATE SEQUENCE "public"."contracts_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."contracts_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for dapp_watch_list_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."dapp_watch_list_id_seq";
CREATE SEQUENCE "public"."dapp_watch_list_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."dapp_watch_list_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for dapps_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."dapps_id_seq";
CREATE SEQUENCE "public"."dapps_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."dapps_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for external_bridges_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."external_bridges_id_seq";
CREATE SEQUENCE "public"."external_bridges_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."external_bridges_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for external_swaps_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."external_swaps_id_seq";
CREATE SEQUENCE "public"."external_swaps_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."external_swaps_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for inscription_whitelist_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."inscription_whitelist_id_seq";
CREATE SEQUENCE "public"."inscription_whitelist_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."inscription_whitelist_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for inscriptions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."inscriptions_id_seq";
CREATE SEQUENCE "public"."inscriptions_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."inscriptions_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for internal_transactions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."internal_transactions_id_seq";
CREATE SEQUENCE "public"."internal_transactions_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."internal_transactions_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for l1_batches_number_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."l1_batches_number_seq";
CREATE SEQUENCE "public"."l1_batches_number_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."l1_batches_number_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for label_to_addresses_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."label_to_addresses_id_seq";
CREATE SEQUENCE "public"."label_to_addresses_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."label_to_addresses_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for reorged_blocks_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."reorged_blocks_id_seq";
CREATE SEQUENCE "public"."reorged_blocks_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."reorged_blocks_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for signatures_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."signatures_id_seq";
CREATE SEQUENCE "public"."signatures_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."signatures_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for sync_progress_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sync_progress_id_seq";
CREATE SEQUENCE "public"."sync_progress_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."sync_progress_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for token_balances_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."token_balances_id_seq";
CREATE SEQUENCE "public"."token_balances_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."token_balances_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for token_transfers_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."token_transfers_id_seq";
CREATE SEQUENCE "public"."token_transfers_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."token_transfers_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for token_watch_list_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."token_watch_list_id_seq";
CREATE SEQUENCE "public"."token_watch_list_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."token_watch_list_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for tokens_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."tokens_id_seq";
CREATE SEQUENCE "public"."tokens_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."tokens_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for transaction_logs_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."transaction_logs_id_seq";
CREATE SEQUENCE "public"."transaction_logs_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."transaction_logs_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for transactions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."transactions_id_seq";
CREATE SEQUENCE "public"."transactions_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."transactions_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for zksync_bridge_deposit_history_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."zksync_bridge_deposit_history_id_seq";
CREATE SEQUENCE "public"."zksync_bridge_deposit_history_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."zksync_bridge_deposit_history_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for zksync_bridge_withdraw_history_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."zksync_bridge_withdraw_history_id_seq";
CREATE SEQUENCE "public"."zksync_bridge_withdraw_history_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."zksync_bridge_withdraw_history_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Table structure for account_watch_list
-- ----------------------------
DROP TABLE IF EXISTS "public"."account_watch_list";
CREATE TABLE "public"."account_watch_list" (
                                               "id" int4 NOT NULL DEFAULT nextval('account_watch_list_id_seq'::regclass),
                                               "user_id" text COLLATE "pg_catalog"."default" NOT NULL,
                                               "address" text COLLATE "pg_catalog"."default" NOT NULL,
                                               "email" text COLLATE "pg_catalog"."default",
                                               "description" text COLLATE "pg_catalog"."default",
                                               "notification_method" text COLLATE "pg_catalog"."default" NOT NULL,
                                               "track_erc20" bool NOT NULL DEFAULT false,
                                               "track_erc721" bool NOT NULL DEFAULT false,
                                               "track_erc1155" bool NOT NULL DEFAULT false,
                                               "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                               "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."account_watch_list" OWNER TO "postgres";

-- ----------------------------
-- Table structure for accounts_list
-- ----------------------------
DROP TABLE IF EXISTS "public"."accounts_list";
CREATE TABLE "public"."accounts_list" (
                                          "id" int4 NOT NULL DEFAULT nextval('accounts_list_id_seq'::regclass),
                                          "user_id" text COLLATE "pg_catalog"."default" NOT NULL,
                                          "address_list" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."accounts_list" OWNER TO "postgres";

-- ----------------------------
-- Table structure for address_balances
-- ----------------------------
DROP TABLE IF EXISTS "public"."address_balances";
CREATE TABLE "public"."address_balances" (
                                             "id" int8 NOT NULL DEFAULT nextval('address_balances_id_seq'::regclass),
                                             "address" text COLLATE "pg_catalog"."default" NOT NULL,
                                             "balance" numeric(100,0),
                                             "updated_block_number" int8,
                                             "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                             "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."address_balances" OWNER TO "postgres";

-- ----------------------------
-- Table structure for address_to_labels
-- ----------------------------
DROP TABLE IF EXISTS "public"."address_to_labels";
CREATE TABLE "public"."address_to_labels" (
                                              "id" int8 NOT NULL DEFAULT nextval('address_to_labels_id_seq'::regclass),
                                              "address" text COLLATE "pg_catalog"."default" NOT NULL,
                                              "name" text COLLATE "pg_catalog"."default",
                                              "site" text COLLATE "pg_catalog"."default",
                                              "labels" text[] COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."address_to_labels" OWNER TO "postgres";

-- ----------------------------
-- Table structure for blocks
-- ----------------------------
DROP TABLE IF EXISTS "public"."blocks";
CREATE TABLE "public"."blocks" (
                                   "number" int8 NOT NULL,
                                   "hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                   "transaction_count" int4,
                                   "internal_transaction_count" int4,
                                   "trace_checked" bool,
                                   "validator" text COLLATE "pg_catalog"."default",
                                   "difficulty" numeric(80,0),
                                   "total_difficulty" numeric(80,0),
                                   "size" int4,
                                   "nonce" text COLLATE "pg_catalog"."default",
                                   "gas_used" int8,
                                   "gas_limit" int8,
                                   "extra_data" text COLLATE "pg_catalog"."default",
                                   "parent_hash" text COLLATE "pg_catalog"."default",
                                   "sha3_uncle" text COLLATE "pg_catalog"."default",
                                   "timestamp" int8,
                                   "l1_batch_number" int8,
                                   "l1_batch_timestamp" int8,
                                   "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."blocks" OWNER TO "postgres";

-- ----------------------------
-- Table structure for bridge_stats
-- ----------------------------
DROP TABLE IF EXISTS "public"."bridge_stats";
CREATE TABLE "public"."bridge_stats" (
                                         "id" int4 NOT NULL DEFAULT nextval('bridge_stats_id_seq'::regclass),
                                         "type" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "network" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "token_symbol" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "token_address" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "token_decimals" int4 NOT NULL,
                                         "transaction_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "transaction_status" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "block_number" int8 NOT NULL,
                                         "value" numeric(80,0) NOT NULL,
                                         "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."bridge_stats" OWNER TO "postgres";

-- ----------------------------
-- Table structure for contract_verify_job
-- ----------------------------
DROP TABLE IF EXISTS "public"."contract_verify_job";
CREATE TABLE "public"."contract_verify_job" (
                                                "uid" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                                "contract_address" varchar(200) COLLATE "pg_catalog"."default",
                                                "compiler" varchar(200) COLLATE "pg_catalog"."default",
                                                "standard_json" text COLLATE "pg_catalog"."default",
                                                "status" int2,
                                                "failed_reason" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."contract_verify_job" OWNER TO "postgres";

-- ----------------------------
-- Table structure for contracts
-- ----------------------------
DROP TABLE IF EXISTS "public"."contracts";
CREATE TABLE "public"."contracts" (
                                      "id" int8 NOT NULL DEFAULT nextval('contracts_id_seq'::regclass),
                                      "name" text COLLATE "pg_catalog"."default",
                                      "address" text COLLATE "pg_catalog"."default" NOT NULL,
                                      "creator" text COLLATE "pg_catalog"."default",
                                      "creation_tx_hash" text COLLATE "pg_catalog"."default",
                                      "creation_bytecode" text COLLATE "pg_catalog"."default",
                                      "deployed_bytecode" text COLLATE "pg_catalog"."default",
                                      "abi" jsonb,
                                      "constructor_arguments" text COLLATE "pg_catalog"."default",
                                      "sourcecode" text COLLATE "pg_catalog"."default",
                                      "compiler_version" text COLLATE "pg_catalog"."default",
                                      "optimization" bool,
                                      "optimization_runs" int4,
                                      "evm_version" text COLLATE "pg_catalog"."default",
                                      "is_verified" bool,
                                      "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      "license" text COLLATE "pg_catalog"."default",
                                      "creation_timestamp" int8
)
;
ALTER TABLE "public"."contracts" OWNER TO "postgres";

-- ----------------------------
-- Table structure for dapp_watch_list
-- ----------------------------
DROP TABLE IF EXISTS "public"."dapp_watch_list";
CREATE TABLE "public"."dapp_watch_list" (
                                            "id" int4 NOT NULL DEFAULT nextval('dapp_watch_list_id_seq'::regclass),
                                            "user_id" text COLLATE "pg_catalog"."default" NOT NULL,
                                            "dapp_id" int4 NOT NULL,
                                            "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                            "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."dapp_watch_list" OWNER TO "postgres";

-- ----------------------------
-- Table structure for dapps
-- ----------------------------
DROP TABLE IF EXISTS "public"."dapps";
CREATE TABLE "public"."dapps" (
                                  "id" int4 NOT NULL DEFAULT nextval('dapps_id_seq'::regclass),
                                  "name" text COLLATE "pg_catalog"."default" NOT NULL,
                                  "team" text COLLATE "pg_catalog"."default",
                                  "logo" text COLLATE "pg_catalog"."default",
                                  "description" text COLLATE "pg_catalog"."default",
                                  "contract" text COLLATE "pg_catalog"."default",
                                  "categories" text[] COLLATE "pg_catalog"."default",
                                  "website" text COLLATE "pg_catalog"."default",
                                  "discord" text COLLATE "pg_catalog"."default",
                                  "media_url" text COLLATE "pg_catalog"."default",
                                  "telegram" text COLLATE "pg_catalog"."default",
                                  "twitter" text COLLATE "pg_catalog"."default",
                                  "youtube" text COLLATE "pg_catalog"."default",
                                  "addresses" text[] COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."dapps" OWNER TO "postgres";

-- ----------------------------
-- Table structure for external_bridges
-- ----------------------------
DROP TABLE IF EXISTS "public"."external_bridges";
CREATE TABLE "public"."external_bridges" (
                                             "id" int4 NOT NULL DEFAULT nextval('external_bridges_id_seq'::regclass),
                                             "name" text COLLATE "pg_catalog"."default" NOT NULL,
                                             "logo" text COLLATE "pg_catalog"."default",
                                             "introduction" text COLLATE "pg_catalog"."default",
                                             "tags" text COLLATE "pg_catalog"."default",
                                             "external_link" text COLLATE "pg_catalog"."default" NOT NULL,
                                             "dappId" int4 NOT NULL
)
;
ALTER TABLE "public"."external_bridges" OWNER TO "postgres";

-- ----------------------------
-- Table structure for external_swaps
-- ----------------------------
DROP TABLE IF EXISTS "public"."external_swaps";
CREATE TABLE "public"."external_swaps" (
                                           "id" int4 NOT NULL DEFAULT nextval('external_swaps_id_seq'::regclass),
                                           "name" text COLLATE "pg_catalog"."default" NOT NULL,
                                           "logo" text COLLATE "pg_catalog"."default",
                                           "introduction" text COLLATE "pg_catalog"."default",
                                           "tags" text COLLATE "pg_catalog"."default",
                                           "external_link" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."external_swaps" OWNER TO "postgres";

-- ----------------------------
-- Table structure for inscription_whitelist
-- ----------------------------
DROP TABLE IF EXISTS "public"."inscription_whitelist";
CREATE TABLE "public"."inscription_whitelist" (
                                                  "id" int8 NOT NULL DEFAULT nextval('inscription_whitelist_id_seq'::regclass),
                                                  "deploy_txhash" text COLLATE "pg_catalog"."default" NOT NULL,
                                                  "timestamp" int8 NOT NULL,
                                                  "block_number" int8 NOT NULL,
                                                  "tick" text COLLATE "pg_catalog"."default" NOT NULL,
                                                  "max_supply" int8 NOT NULL,
                                                  "limit_per_mint" int8 NOT NULL,
                                                  "standard" text COLLATE "pg_catalog"."default" NOT NULL,
                                                  "mint_json" jsonb NOT NULL,
                                                  "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                  "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."inscription_whitelist" OWNER TO "postgres";

-- ----------------------------
-- Table structure for inscriptions
-- ----------------------------
DROP TABLE IF EXISTS "public"."inscriptions";
CREATE TABLE "public"."inscriptions" (
                                         "id" int8 NOT NULL DEFAULT nextval('inscriptions_id_seq'::regclass),
                                         "transaction_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "transaction_index" int4 NOT NULL,
                                         "from_address" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "to_address" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "block_number" int8 NOT NULL,
                                         "gas_price" int8,
                                         "full_inscription" jsonb NOT NULL,
                                         "timestamp" int8 NOT NULL,
                                         "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."inscriptions" OWNER TO "postgres";

-- ----------------------------
-- Table structure for internal_transactions
-- ----------------------------
DROP TABLE IF EXISTS "public"."internal_transactions";
CREATE TABLE "public"."internal_transactions" (
                                                  "id" int8 NOT NULL DEFAULT nextval('internal_transactions_id_seq'::regclass),
                                                  "block_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                                  "block_number" text COLLATE "pg_catalog"."default" NOT NULL,
                                                  "parent_transaction_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                                  "type" text COLLATE "pg_catalog"."default",
                                                  "from_address" text COLLATE "pg_catalog"."default",
                                                  "to_address" text COLLATE "pg_catalog"."default",
                                                  "value" numeric(80,0),
                                                  "gas" numeric(80,0),
                                                  "gas_used" numeric(80,0),
                                                  "input" text COLLATE "pg_catalog"."default",
                                                  "output" text COLLATE "pg_catalog"."default",
                                                  "method" text COLLATE "pg_catalog"."default",
                                                  "timestamp" int8,
                                                  "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                  "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."internal_transactions" OWNER TO "postgres";

-- ----------------------------
-- Table structure for l1_batches
-- ----------------------------
DROP TABLE IF EXISTS "public"."l1_batches";
CREATE TABLE "public"."l1_batches" (
                                       "number" int8 NOT NULL DEFAULT nextval('l1_batches_number_seq'::regclass),
                                       "commit_tx_hash" text COLLATE "pg_catalog"."default",
                                       "committed_at" timestamp(3),
                                       "execute_tx_hash" text COLLATE "pg_catalog"."default",
                                       "executed_at" timestamp(3),
                                       "prove_tx_hash" text COLLATE "pg_catalog"."default",
                                       "proven_at" timestamp(3),
                                       "root_hash" text COLLATE "pg_catalog"."default",
                                       "status" text COLLATE "pg_catalog"."default",
                                       "l1_gas_price" int8,
                                       "l1_tx_count" int8,
                                       "l2_fair_gas_price" int8,
                                       "l2_tx_count" int8,
                                       "timestamp" int8 NOT NULL,
                                       "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                       "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                       "l1_prove_block_number" int8,
                                       "l2_block_number" int8
)
;
ALTER TABLE "public"."l1_batches" OWNER TO "postgres";

-- ----------------------------
-- Table structure for label_to_addresses
-- ----------------------------
DROP TABLE IF EXISTS "public"."label_to_addresses";
CREATE TABLE "public"."label_to_addresses" (
                                               "id" int8 NOT NULL DEFAULT nextval('label_to_addresses_id_seq'::regclass),
                                               "label" text COLLATE "pg_catalog"."default" NOT NULL,
                                               "addresses" text[] COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."label_to_addresses" OWNER TO "postgres";

-- ----------------------------
-- Table structure for mv_dapp_daily_transactions
-- ----------------------------
DROP TABLE IF EXISTS "public"."mv_dapp_daily_transactions";
CREATE TABLE "public"."mv_dapp_daily_transactions" (
                                                       "dapp_id" int4 NOT NULL,
                                                       "date" date NOT NULL,
                                                       "count" int8
)
;
ALTER TABLE "public"."mv_dapp_daily_transactions" OWNER TO "postgres";

-- ----------------------------
-- Table structure for reorged_blocks
-- ----------------------------
DROP TABLE IF EXISTS "public"."reorged_blocks";
CREATE TABLE "public"."reorged_blocks" (
                                           "id" int8 NOT NULL DEFAULT nextval('reorged_blocks_id_seq'::regclass),
                                           "number" int8,
                                           "hash" text COLLATE "pg_catalog"."default",
                                           "validator" text COLLATE "pg_catalog"."default",
                                           "transaction_count" int4,
                                           "internal_transaction_count" int4,
                                           "size" int4,
                                           "gas_used" int8,
                                           "gas_limit" int8,
                                           "extra_data" text COLLATE "pg_catalog"."default",
                                           "parent_hash" text COLLATE "pg_catalog"."default",
                                           "sha3_uncle" text COLLATE "pg_catalog"."default",
                                           "uncles" text COLLATE "pg_catalog"."default",
                                           "timestamp" int8,
                                           "depth" int4,
                                           "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                           "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."reorged_blocks" OWNER TO "postgres";

-- ----------------------------
-- Table structure for signatures
-- ----------------------------
DROP TABLE IF EXISTS "public"."signatures";
CREATE TABLE "public"."signatures" (
                                       "id" int4 NOT NULL DEFAULT nextval('signatures_id_seq'::regclass),
                                       "hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                       "name" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."signatures" OWNER TO "postgres";

-- ----------------------------
-- Table structure for sync_progress
-- ----------------------------
DROP TABLE IF EXISTS "public"."sync_progress";
CREATE TABLE "public"."sync_progress" (
                                          "id" int4 NOT NULL DEFAULT nextval('sync_progress_id_seq'::regclass),
                                          "key" text COLLATE "pg_catalog"."default",
                                          "value" int4
)
;
ALTER TABLE "public"."sync_progress" OWNER TO "postgres";

-- ----------------------------
-- Table structure for token_balances
-- ----------------------------
DROP TABLE IF EXISTS "public"."token_balances";
CREATE TABLE "public"."token_balances" (
                                           "id" int8 NOT NULL DEFAULT nextval('token_balances_id_seq'::regclass),
                                           "address" text COLLATE "pg_catalog"."default" NOT NULL,
                                           "token_address" text COLLATE "pg_catalog"."default" NOT NULL,
                                           "token_id" numeric(80,0),
                                           "token_type" text COLLATE "pg_catalog"."default" NOT NULL,
                                           "balance" numeric(80,0),
                                           "updated_block_number" int8,
                                           "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                           "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."token_balances" OWNER TO "postgres";

-- ----------------------------
-- Table structure for token_transfers
-- ----------------------------
DROP TABLE IF EXISTS "public"."token_transfers";
CREATE TABLE "public"."token_transfers" (
                                            "id" int8 NOT NULL DEFAULT nextval('token_transfers_id_seq'::regclass),
                                            "transaction_hash" text COLLATE "pg_catalog"."default",
                                            "log_index" int4,
                                            "method_id" text COLLATE "pg_catalog"."default",
                                            "token_address" text COLLATE "pg_catalog"."default",
                                            "block_number" int8,
                                            "block_hash" text COLLATE "pg_catalog"."default",
                                            "from_address" text COLLATE "pg_catalog"."default",
                                            "to_address" text COLLATE "pg_catalog"."default",
                                            "value" numeric(80,0),
                                            "amount" numeric(80,0),
                                            "token_id" numeric(80,0),
                                            "amounts" numeric(80,0)[],
                                            "token_ids" numeric(80,0)[],
                                            "token_type" text COLLATE "pg_catalog"."default",
                                            "timestamp" int8,
                                            "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                            "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."token_transfers" OWNER TO "postgres";

-- ----------------------------
-- Table structure for token_watch_list
-- ----------------------------
DROP TABLE IF EXISTS "public"."token_watch_list";
CREATE TABLE "public"."token_watch_list" (
                                             "id" int4 NOT NULL DEFAULT nextval('token_watch_list_id_seq'::regclass),
                                             "user_id" text COLLATE "pg_catalog"."default" NOT NULL,
                                             "address" text COLLATE "pg_catalog"."default" NOT NULL,
                                             "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                             "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."token_watch_list" OWNER TO "postgres";

-- ----------------------------
-- Table structure for tokens
-- ----------------------------
DROP TABLE IF EXISTS "public"."tokens";
CREATE TABLE "public"."tokens" (
                                   "id" int8 NOT NULL DEFAULT nextval('tokens_id_seq'::regclass),
                                   "name" text COLLATE "pg_catalog"."default",
                                   "symbol" text COLLATE "pg_catalog"."default",
                                   "address" text COLLATE "pg_catalog"."default" NOT NULL,
                                   "decimals" int4,
                                   "total_supply" numeric(80,0),
                                   "token_type" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."tokens" OWNER TO "postgres";

-- ----------------------------
-- Table structure for transaction_logs
-- ----------------------------
DROP TABLE IF EXISTS "public"."transaction_logs";
CREATE TABLE "public"."transaction_logs" (
                                             "id" int8 NOT NULL DEFAULT nextval('transaction_logs_id_seq'::regclass),
                                             "transaction_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                             "log_index" int4 NOT NULL,
                                             "address" text COLLATE "pg_catalog"."default" NOT NULL,
                                             "topic1" text COLLATE "pg_catalog"."default",
                                             "topic2" text COLLATE "pg_catalog"."default",
                                             "topic3" text COLLATE "pg_catalog"."default",
                                             "topic4" text COLLATE "pg_catalog"."default",
                                             "data" text COLLATE "pg_catalog"."default",
                                             "removed" bool,
                                             "block_number" int8 NOT NULL,
                                             "block_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                             "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                             "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."transaction_logs" OWNER TO "postgres";

-- ----------------------------
-- Table structure for transactions
-- ----------------------------
DROP TABLE IF EXISTS "public"."transactions";
CREATE TABLE "public"."transactions" (
                                         "id" int8 NOT NULL DEFAULT nextval('transactions_id_seq'::regclass),
                                         "hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "block_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                         "block_number" int8 NOT NULL,
                                         "from_address" text COLLATE "pg_catalog"."default",
                                         "to_address" text COLLATE "pg_catalog"."default",
                                         "value" numeric(80,0),
                                         "fee" numeric(80,0),
                                         "l1fee" numeric(80,0),
                                         "gas_used" int8,
                                         "gas_price" int8,
                                         "gas_limit" int8,
                                         "method_id" text COLLATE "pg_catalog"."default",
                                         "input" text COLLATE "pg_catalog"."default",
                                         "nonce" int4,
                                         "status" int4,
                                         "transaction_index" int4,
                                         "transaction_type" text COLLATE "pg_catalog"."default",
                                         "max_priority" numeric(65,30),
                                         "max_fee" numeric(65,30),
                                         "revert_reason" text COLLATE "pg_catalog"."default",
                                         "l1_batch_number" int8,
                                         "l1_batch_tx_index" int4,
                                         "timestamp" int8,
                                         "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."transactions" OWNER TO "postgres";

-- ----------------------------
-- Table structure for zksync_bridge_deposit_history
-- ----------------------------
DROP TABLE IF EXISTS "public"."zksync_bridge_deposit_history";
CREATE TABLE "public"."zksync_bridge_deposit_history" (
                                                          "id" int4 NOT NULL DEFAULT nextval('zksync_bridge_deposit_history_id_seq'::regclass),
                                                          "l1_tx_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                                          "l2_tx_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                                          "l1_tx_timestamp" int8 NOT NULL,
                                                          "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                          "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."zksync_bridge_deposit_history" OWNER TO "postgres";

-- ----------------------------
-- Table structure for zksync_bridge_withdraw_history
-- ----------------------------
DROP TABLE IF EXISTS "public"."zksync_bridge_withdraw_history";
CREATE TABLE "public"."zksync_bridge_withdraw_history" (
                                                           "id" int4 NOT NULL DEFAULT nextval('zksync_bridge_withdraw_history_id_seq'::regclass),
                                                           "l1_batch_number" int8 NOT NULL,
                                                           "l1_batch_tx_index" int4 NOT NULL,
                                                           "l1_tx_hash" text COLLATE "pg_catalog"."default" NOT NULL,
                                                           "l1_tx_timestamp" int8 NOT NULL,
                                                           "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                           "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."zksync_bridge_withdraw_history" OWNER TO "postgres";

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."account_watch_list_id_seq"
    OWNED BY "public"."account_watch_list"."id";
SELECT setval('"public"."account_watch_list_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."accounts_list_id_seq"
    OWNED BY "public"."accounts_list"."id";
SELECT setval('"public"."accounts_list_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."address_balances_id_seq"
    OWNED BY "public"."address_balances"."id";
SELECT setval('"public"."address_balances_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."address_to_labels_id_seq"
    OWNED BY "public"."address_to_labels"."id";
SELECT setval('"public"."address_to_labels_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."bridge_stats_id_seq"
    OWNED BY "public"."bridge_stats"."id";
SELECT setval('"public"."bridge_stats_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."contracts_id_seq"
    OWNED BY "public"."contracts"."id";
SELECT setval('"public"."contracts_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."dapp_watch_list_id_seq"
    OWNED BY "public"."dapp_watch_list"."id";
SELECT setval('"public"."dapp_watch_list_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."dapps_id_seq"
    OWNED BY "public"."dapps"."id";
SELECT setval('"public"."dapps_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."external_bridges_id_seq"
    OWNED BY "public"."external_bridges"."id";
SELECT setval('"public"."external_bridges_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."external_swaps_id_seq"
    OWNED BY "public"."external_swaps"."id";
SELECT setval('"public"."external_swaps_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."inscription_whitelist_id_seq"
    OWNED BY "public"."inscription_whitelist"."id";
SELECT setval('"public"."inscription_whitelist_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."inscriptions_id_seq"
    OWNED BY "public"."inscriptions"."id";
SELECT setval('"public"."inscriptions_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."internal_transactions_id_seq"
    OWNED BY "public"."internal_transactions"."id";
SELECT setval('"public"."internal_transactions_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."l1_batches_number_seq"
    OWNED BY "public"."l1_batches"."number";
SELECT setval('"public"."l1_batches_number_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."label_to_addresses_id_seq"
    OWNED BY "public"."label_to_addresses"."id";
SELECT setval('"public"."label_to_addresses_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."reorged_blocks_id_seq"
    OWNED BY "public"."reorged_blocks"."id";
SELECT setval('"public"."reorged_blocks_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."signatures_id_seq"
    OWNED BY "public"."signatures"."id";
SELECT setval('"public"."signatures_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sync_progress_id_seq"
    OWNED BY "public"."sync_progress"."id";
SELECT setval('"public"."sync_progress_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."token_balances_id_seq"
    OWNED BY "public"."token_balances"."id";
SELECT setval('"public"."token_balances_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."token_transfers_id_seq"
    OWNED BY "public"."token_transfers"."id";
SELECT setval('"public"."token_transfers_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."token_watch_list_id_seq"
    OWNED BY "public"."token_watch_list"."id";
SELECT setval('"public"."token_watch_list_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."tokens_id_seq"
    OWNED BY "public"."tokens"."id";
SELECT setval('"public"."tokens_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."transaction_logs_id_seq"
    OWNED BY "public"."transaction_logs"."id";
SELECT setval('"public"."transaction_logs_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."transactions_id_seq"
    OWNED BY "public"."transactions"."id";
SELECT setval('"public"."transactions_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."zksync_bridge_deposit_history_id_seq"
    OWNED BY "public"."zksync_bridge_deposit_history"."id";
SELECT setval('"public"."zksync_bridge_deposit_history_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."zksync_bridge_withdraw_history_id_seq"
    OWNED BY "public"."zksync_bridge_withdraw_history"."id";
SELECT setval('"public"."zksync_bridge_withdraw_history_id_seq"', 1, false);

-- ----------------------------
-- Indexes structure for table account_watch_list
-- ----------------------------
CREATE INDEX "account_watch_list_address_idx" ON "public"."account_watch_list" USING btree (
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "account_watch_list_user_id_address_key" ON "public"."account_watch_list" USING btree (
    "user_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "account_watch_list_user_id_idx" ON "public"."account_watch_list" USING btree (
    "user_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table account_watch_list
-- ----------------------------
ALTER TABLE "public"."account_watch_list" ADD CONSTRAINT "account_watch_list_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table accounts_list
-- ----------------------------
CREATE INDEX "accounts_list_user_id_idx" ON "public"."accounts_list" USING btree (
    "user_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "accounts_list_user_id_key" ON "public"."accounts_list" USING btree (
    "user_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table accounts_list
-- ----------------------------
ALTER TABLE "public"."accounts_list" ADD CONSTRAINT "accounts_list_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table address_balances
-- ----------------------------
CREATE UNIQUE INDEX "address_balances_address_key" ON "public"."address_balances" USING btree (
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table address_balances
-- ----------------------------
ALTER TABLE "public"."address_balances" ADD CONSTRAINT "address_balances_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table address_to_labels
-- ----------------------------
CREATE UNIQUE INDEX "address_to_labels_address_key" ON "public"."address_to_labels" USING btree (
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table address_to_labels
-- ----------------------------
ALTER TABLE "public"."address_to_labels" ADD CONSTRAINT "address_to_labels_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table blocks
-- ----------------------------
CREATE INDEX "blocks_hash_idx" ON "public"."blocks" USING btree (
    "hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "blocks_hash_key" ON "public"."blocks" USING btree (
    "hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "blocks_l1_batch_number_idx" ON "public"."blocks" USING btree (
    "l1_batch_number" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table blocks
-- ----------------------------
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_pkey" PRIMARY KEY ("number");

-- ----------------------------
-- Primary Key structure for table bridge_stats
-- ----------------------------
ALTER TABLE "public"."bridge_stats" ADD CONSTRAINT "bridge_stats_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table contract_verify_job
-- ----------------------------
ALTER TABLE "public"."contract_verify_job" ADD CONSTRAINT "contract_verify_job_pkey" PRIMARY KEY ("uid");

-- ----------------------------
-- Indexes structure for table contracts
-- ----------------------------
CREATE UNIQUE INDEX "contracts_address_key" ON "public"."contracts" USING btree (
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "contracts_creation_timestamp_idx" ON "public"."contracts" USING btree (
    "creation_timestamp" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "contracts_creator_idx" ON "public"."contracts" USING btree (
    "creator" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "contracts_is_verified_creation_timestamp_idx" ON "public"."contracts" USING btree (
    "is_verified" "pg_catalog"."bool_ops" ASC NULLS LAST,
    "creation_timestamp" "pg_catalog"."int8_ops" DESC NULLS FIRST
    );

-- ----------------------------
-- Primary Key structure for table contracts
-- ----------------------------
ALTER TABLE "public"."contracts" ADD CONSTRAINT "contracts_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table dapp_watch_list
-- ----------------------------
CREATE INDEX "dapp_watch_list_dapp_id_idx" ON "public"."dapp_watch_list" USING btree (
    "dapp_id" "pg_catalog"."int4_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "dapp_watch_list_user_id_dapp_id_key" ON "public"."dapp_watch_list" USING btree (
    "user_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "dapp_id" "pg_catalog"."int4_ops" ASC NULLS LAST
    );
CREATE INDEX "dapp_watch_list_user_id_idx" ON "public"."dapp_watch_list" USING btree (
    "user_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table dapp_watch_list
-- ----------------------------
ALTER TABLE "public"."dapp_watch_list" ADD CONSTRAINT "dapp_watch_list_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table dapps
-- ----------------------------
CREATE INDEX "dapps_categories_idx" ON "public"."dapps" USING gin (
    "categories" COLLATE "pg_catalog"."default" "pg_catalog"."array_ops"
    );

-- ----------------------------
-- Primary Key structure for table dapps
-- ----------------------------
ALTER TABLE "public"."dapps" ADD CONSTRAINT "dapps_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table external_bridges
-- ----------------------------
ALTER TABLE "public"."external_bridges" ADD CONSTRAINT "external_bridges_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table external_swaps
-- ----------------------------
ALTER TABLE "public"."external_swaps" ADD CONSTRAINT "external_swaps_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table inscription_whitelist
-- ----------------------------
CREATE INDEX "inscription_whitelist_deploy_txhash_idx" ON "public"."inscription_whitelist" USING btree (
    "deploy_txhash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "inscription_whitelist_deploy_txhash_key" ON "public"."inscription_whitelist" USING btree (
    "deploy_txhash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "inscription_whitelist_tick_idx" ON "public"."inscription_whitelist" USING btree (
    "tick" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "inscription_whitelist_tick_key" ON "public"."inscription_whitelist" USING btree (
    "tick" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table inscription_whitelist
-- ----------------------------
ALTER TABLE "public"."inscription_whitelist" ADD CONSTRAINT "inscription_whitelist_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table inscriptions
-- ----------------------------
CREATE INDEX "inscriptions_block_number_idx" ON "public"."inscriptions" USING btree (
    "block_number" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "inscriptions_from_address_idx" ON "public"."inscriptions" USING btree (
    "from_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "inscriptions_full_inscription_idx" ON "public"."inscriptions" USING gin (
    "full_inscription" "pg_catalog"."jsonb_path_ops"
    );
CREATE INDEX "inscriptions_timestamp_idx" ON "public"."inscriptions" USING btree (
    "timestamp" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "inscriptions_to_address_idx" ON "public"."inscriptions" USING btree (
    "to_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "inscriptions_transaction_hash_idx" ON "public"."inscriptions" USING btree (
    "transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "inscriptions_transaction_hash_key" ON "public"."inscriptions" USING btree (
    "transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table inscriptions
-- ----------------------------
ALTER TABLE "public"."inscriptions" ADD CONSTRAINT "inscriptions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table internal_transactions
-- ----------------------------
CREATE INDEX "internal_transactions_block_hash_idx" ON "public"."internal_transactions" USING btree (
    "block_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "internal_transactions_block_number_idx" ON "public"."internal_transactions" USING btree (
    "block_number" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "internal_transactions_from_address_idx" ON "public"."internal_transactions" USING btree (
    "from_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "internal_transactions_parent_transaction_hash_idx" ON "public"."internal_transactions" USING btree (
    "parent_transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "internal_transactions_timestamp_idx" ON "public"."internal_transactions" USING btree (
    "timestamp" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "internal_transactions_to_address_idx" ON "public"."internal_transactions" USING btree (
    "to_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "internal_transactions_to_address_timestamp_idx" ON "public"."internal_transactions" USING btree (
    "to_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "timestamp" "pg_catalog"."int8_ops" DESC NULLS FIRST
    );

-- ----------------------------
-- Primary Key structure for table internal_transactions
-- ----------------------------
ALTER TABLE "public"."internal_transactions" ADD CONSTRAINT "internal_transactions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table l1_batches
-- ----------------------------
ALTER TABLE "public"."l1_batches" ADD CONSTRAINT "l1_batches_pkey" PRIMARY KEY ("number");

-- ----------------------------
-- Indexes structure for table label_to_addresses
-- ----------------------------
CREATE UNIQUE INDEX "label_to_addresses_label_key" ON "public"."label_to_addresses" USING btree (
    "label" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table label_to_addresses
-- ----------------------------
ALTER TABLE "public"."label_to_addresses" ADD CONSTRAINT "label_to_addresses_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table mv_dapp_daily_transactions
-- ----------------------------
ALTER TABLE "public"."mv_dapp_daily_transactions" ADD CONSTRAINT "mv_dapp_daily_transactions_pkey" PRIMARY KEY ("dapp_id", "date");

-- ----------------------------
-- Primary Key structure for table reorged_blocks
-- ----------------------------
ALTER TABLE "public"."reorged_blocks" ADD CONSTRAINT "reorged_blocks_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table signatures
-- ----------------------------
CREATE INDEX "signatures_hash_idx" ON "public"."signatures" USING btree (
    "hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table signatures
-- ----------------------------
ALTER TABLE "public"."signatures" ADD CONSTRAINT "signatures_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table sync_progress
-- ----------------------------
CREATE UNIQUE INDEX "sync_progress_key_key" ON "public"."sync_progress" USING btree (
    "key" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table sync_progress
-- ----------------------------
ALTER TABLE "public"."sync_progress" ADD CONSTRAINT "sync_progress_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table token_balances
-- ----------------------------
CREATE UNIQUE INDEX "token_balances_address_token_address_token_id_token_type_key" ON "public"."token_balances" USING btree (
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "token_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "token_id" "pg_catalog"."numeric_ops" ASC NULLS LAST,
    "token_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "token_balances_token_address_idx" ON "public"."token_balances" USING btree (
    "token_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table token_balances
-- ----------------------------
ALTER TABLE "public"."token_balances" ADD CONSTRAINT "token_balances_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table token_transfers
-- ----------------------------
CREATE INDEX "token_transfers_from_address_idx" ON "public"."token_transfers" USING btree (
    "from_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "token_transfers_to_address_idx" ON "public"."token_transfers" USING btree (
    "to_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "token_transfers_token_address_idx" ON "public"."token_transfers" USING btree (
    "token_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "token_transfers_token_type_idx" ON "public"."token_transfers" USING btree (
    "token_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "token_transfers_transaction_hash_idx" ON "public"."token_transfers" USING btree (
    "transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "token_transfers_transaction_hash_log_index_key" ON "public"."token_transfers" USING btree (
    "transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "log_index" "pg_catalog"."int4_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table token_transfers
-- ----------------------------
ALTER TABLE "public"."token_transfers" ADD CONSTRAINT "token_transfers_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table token_watch_list
-- ----------------------------
CREATE INDEX "token_watch_list_address_idx" ON "public"."token_watch_list" USING btree (
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "token_watch_list_user_id_address_key" ON "public"."token_watch_list" USING btree (
    "user_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "token_watch_list_user_id_idx" ON "public"."token_watch_list" USING btree (
    "user_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table token_watch_list
-- ----------------------------
ALTER TABLE "public"."token_watch_list" ADD CONSTRAINT "token_watch_list_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table tokens
-- ----------------------------
CREATE UNIQUE INDEX "tokens_address_key" ON "public"."tokens" USING btree (
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table tokens
-- ----------------------------
ALTER TABLE "public"."tokens" ADD CONSTRAINT "tokens_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table transaction_logs
-- ----------------------------
CREATE INDEX "transaction_logs_address_idx" ON "public"."transaction_logs" USING btree (
    "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "transaction_logs_block_number_idx" ON "public"."transaction_logs" USING btree (
    "block_number" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "transaction_logs_topic1_idx" ON "public"."transaction_logs" USING btree (
    "topic1" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "transaction_logs_topic2_idx" ON "public"."transaction_logs" USING btree (
    "topic2" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "transaction_logs_topic3_idx" ON "public"."transaction_logs" USING btree (
    "topic3" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "transaction_logs_topic4_idx" ON "public"."transaction_logs" USING btree (
    "topic4" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "transaction_logs_transaction_hash_idx" ON "public"."transaction_logs" USING btree (
    "transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "transaction_logs_transaction_hash_log_index_key" ON "public"."transaction_logs" USING btree (
    "transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "log_index" "pg_catalog"."int4_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table transaction_logs
-- ----------------------------
ALTER TABLE "public"."transaction_logs" ADD CONSTRAINT "transaction_logs_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table transactions
-- ----------------------------
CREATE INDEX "transactions_block_hash_idx" ON "public"."transactions" USING btree (
    "block_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "transactions_block_number_idx" ON "public"."transactions" USING btree (
    "block_number" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "transactions_from_address_idx" ON "public"."transactions" USING btree (
    "from_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "transactions_hash_idx" ON "public"."transactions" USING btree (
    "hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "transactions_hash_key" ON "public"."transactions" USING btree (
    "hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "transactions_l1_batch_number_idx" ON "public"."transactions" USING btree (
    "l1_batch_number" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "transactions_timestamp_idx" ON "public"."transactions" USING btree (
    "timestamp" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "transactions_to_address_idx" ON "public"."transactions" USING btree (
    "to_address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table transactions
-- ----------------------------
ALTER TABLE "public"."transactions" ADD CONSTRAINT "transactions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table zksync_bridge_deposit_history
-- ----------------------------
CREATE UNIQUE INDEX "zksync_bridge_deposit_history_l1_tx_hash_key" ON "public"."zksync_bridge_deposit_history" USING btree (
    "l1_tx_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "zksync_bridge_deposit_history_l2_tx_hash_idx" ON "public"."zksync_bridge_deposit_history" USING btree (
    "l2_tx_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table zksync_bridge_deposit_history
-- ----------------------------
ALTER TABLE "public"."zksync_bridge_deposit_history" ADD CONSTRAINT "zksync_bridge_deposit_history_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table zksync_bridge_withdraw_history
-- ----------------------------
ALTER TABLE "public"."zksync_bridge_withdraw_history" ADD CONSTRAINT "zksync_bridge_withdraw_history_pkey" PRIMARY KEY ("id");


CREATE TABLE "public"."cache" (
   "number" int8 NOT NULL,
   "inserted_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
   "updated_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
