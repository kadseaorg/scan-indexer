package ethclient

import (
	"fmt"
	"testing"
)

func TestParseWithdrawalFinalizationInput(t *testing.T) {
	// Example input data (This needs to be a real encoded input data for the finalizeWithdrawals function)
	// https://etherscan.io/tx/0x61a6819c6fc9be6936e31e90aff7bbd588527f606eca7cf8226e2c89f2b70242
	inputData := "0x32bfc64d0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000360000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000008800000000000000000000000000000000000000000000000000000000000000b000000000000000000000000000000000000000000000000000000000000000d800000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000004a30a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008800000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000493e000000000000000000000000000000000000000000000000000000000000000386c0960f9d4db80f722360112ce2dae67e1fe30960a3a3ae8000000000000000000000000000000000000000000000000000029c4481a1e00000000000000000000000000000000000000000000000000000000000000000000000000000000092fca41f8f9906f31834604255e076b0a8bea7d30fdca29e1e8b5b6f7462b87f8c3d03eebfd83049991ea3d3e358b6712e7aa2e2e63dc2d4b438987cec28ac8d0e3697c7f33c31a9b0f0aeb8542287d0d21e8c4cf82163d0c44c7a98aa11aa111199cc5812543ddceeddd0fc82807646a4899444240db2c0d2f20c3cceb5f51fae4733f281f18ba3ea8775dd62d2fcd84011c8c938f16ea5790fd29a03bf8db891798a1fd9c8fbb818c98cff190daa7cc10b6e5ac9716b4a2649f7c2ebcef227266d7c5983afe44cf15ea8cf565b34c6c31ff0cb4dd744524f7842b942d08770db04e5ee349086985f74b73971ce9dfe76bbed95c84906c5dffd96504e1e5396cac506ecb5465659b3a927143f6d724f91d8d9c4bdb2463aee111d9aa869874db000000000000000000000000000000000000000000000000000000000004a30b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000bb00000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000493e0000000000000000000000000000000000000000000000000000000000000004c11a2ccc1f5a67a31386afbf3b6055b99da58f322210f0861a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000000000000000025432e1fe0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000949698e212b546bc4b1d3b852bd55434841392ce831f9f189de349e1867609a85ab064b75a131ce052508bd8545c8a907e0940d16bf971ce93279e4ab4d482ccc53e4ab376a2aa503f40ad22b4d90490b09c7c627aa675b561678933d1c431dc4199cc5812543ddceeddd0fc82807646a4899444240db2c0d2f20c3cceb5f51fae4733f281f18ba3ea8775dd62d2fcd84011c8c938f16ea5790fd29a03bf8db891798a1fd9c8fbb818c98cff190daa7cc10b6e5ac9716b4a2649f7c2ebcef227266d7c5983afe44cf15ea8cf565b34c6c31ff0cb4dd744524f7842b942d08770db04e5ee349086985f74b73971ce9dfe76bbed95c84906c5dffd96504e1e5396cac506ecb5465659b3a927143f6d724f91d8d9c4bdb2463aee111d9aa869874db000000000000000000000000000000000000000000000000000000000004a30b000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000d100000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000493e000000000000000000000000000000000000000000000000000000000000000386c0960f96fad0bd9ff7049b96fda299ab31e1cc6fe0ab68f000000000000000000000000000000000000000000000000000109a9d4bd280000000000000000000000000000000000000000000000000000000000000000000000000000000009b02984694f7b7de1abab724e9293da307136e00bdafeee48f3ccefa1791eab6eab064b75a131ce052508bd8545c8a907e0940d16bf971ce93279e4ab4d482ccc53e4ab376a2aa503f40ad22b4d90490b09c7c627aa675b561678933d1c431dc4199cc5812543ddceeddd0fc82807646a4899444240db2c0d2f20c3cceb5f51fae4733f281f18ba3ea8775dd62d2fcd84011c8c938f16ea5790fd29a03bf8db891798a1fd9c8fbb818c98cff190daa7cc10b6e5ac9716b4a2649f7c2ebcef227266d7c5983afe44cf15ea8cf565b34c6c31ff0cb4dd744524f7842b942d08770db04e5ee349086985f74b73971ce9dfe76bbed95c84906c5dffd96504e1e5396cac506ecb5465659b3a927143f6d724f91d8d9c4bdb2463aee111d9aa869874db000000000000000000000000000000000000000000000000000000000004a30b0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000011800000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000493e000000000000000000000000000000000000000000000000000000000000000386c0960f9a5bbf1ed1560bafc25ab80f3e223c817837cedf70000000000000000000000000000000000000000000000000005be8b360b6000000000000000000000000000000000000000000000000000000000000000000000000000000000090d1b0793c750582e3ffa9a772bfdceb5d3301622c11eff59f5bafd52c8ab5e1cb39606a28977d020cfa398fe6e0890c7fc805a66d6242c1a9453abbe30cabc4753e4ab376a2aa503f40ad22b4d90490b09c7c627aa675b561678933d1c431dc4199cc5812543ddceeddd0fc82807646a4899444240db2c0d2f20c3cceb5f51fae4733f281f18ba3ea8775dd62d2fcd84011c8c938f16ea5790fd29a03bf8db891798a1fd9c8fbb818c98cff190daa7cc10b6e5ac9716b4a2649f7c2ebcef227266d7c5983afe44cf15ea8cf565b34c6c31ff0cb4dd744524f7842b942d08770db04e5ee349086985f74b73971ce9dfe76bbed95c84906c5dffd96504e1e5396cac506ecb5465659b3a927143f6d724f91d8d9c4bdb2463aee111d9aa869874db000000000000000000000000000000000000000000000000000000000004a30b0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000018e00000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000493e000000000000000000000000000000000000000000000000000000000000000386c0960f9874a07cea7cf852b24e459e5b2f85a994b9c250d00000000000000000000000000000000000000000000000006bc5889734c90690000000000000000000000000000000000000000000000000000000000000000000000000000000982ac977774f778895cd70fe3225dec54957547ca50ff1e9ccd34d91e8d3d2cd8b39606a28977d020cfa398fe6e0890c7fc805a66d6242c1a9453abbe30cabc4753e4ab376a2aa503f40ad22b4d90490b09c7c627aa675b561678933d1c431dc4199cc5812543ddceeddd0fc82807646a4899444240db2c0d2f20c3cceb5f51fae4733f281f18ba3ea8775dd62d2fcd84011c8c938f16ea5790fd29a03bf8db891798a1fd9c8fbb818c98cff190daa7cc10b6e5ac9716b4a2649f7c2ebcef227266d7c5983afe44cf15ea8cf565b34c6c31ff0cb4dd744524f7842b942d08770db04e5ee349086985f74b73971ce9dfe76bbed95c84906c5dffd96504e1e5396cac506ecb5465659b3a927143f6d724f91d8d9c4bdb2463aee111d9aa869874db000000000000000000000000000000000000000000000000000000000004a30c0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000015d00000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000493e000000000000000000000000000000000000000000000000000000000000000386c0960f98ae0bcea0cd0be10a3f4ec0b835c42afc124bc1d00000000000000000000000000000000000000000000000000086ba116fad7f8000000000000000000000000000000000000000000000000000000000000000000000000000000096f95f1f9edd45a94a531f18c66b081e767b79500dc01070ebfa3b0c457e864bd9eadca69ae34d774d08ff11ba3fe2f48f115e94e28bc649832d19e16bc6b3a50532cd850a71219e8e4bfe71c7eb7d9e36ddb0b3da7f331365af7620d41d85c34199cc5812543ddceeddd0fc82807646a4899444240db2c0d2f20c3cceb5f51fae4733f281f18ba3ea8775dd62d2fcd84011c8c938f16ea5790fd29a03bf8db891798a1fd9c8fbb818c98cff190daa7cc10b6e5ac9716b4a2649f7c2ebcef227266d7c5983afe44cf15ea8cf565b34c6c31ff0cb4dd744524f7842b942d08770db04e5ee349086985f74b73971ce9dfe76bbed95c84906c5dffd96504e1e5396cac506ecb5465659b3a927143f6d724f91d8d9c4bdb2463aee111d9aa869874db000000000000000000000000000000000000000000000000000000000004a30c0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000029f00000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000493e000000000000000000000000000000000000000000000000000000000000000386c0960f988dc26f26b3c9ef47fb413e55c1d8d62d6b6e23a0000000000000000000000000000000000000000000000006124fee993bc00000000000000000000000000000000000000000000000000000000000000000000000000000000000931ca51d11623c2d4b7edce161de14fd23df756948d66714ba50cd26f26523b011eea1d7ce82f360f2d472344099263feb8ededa958f77e97a93836c38c15ba2e532cd850a71219e8e4bfe71c7eb7d9e36ddb0b3da7f331365af7620d41d85c34199cc5812543ddceeddd0fc82807646a4899444240db2c0d2f20c3cceb5f51fae4733f281f18ba3ea8775dd62d2fcd84011c8c938f16ea5790fd29a03bf8db891798a1fd9c8fbb818c98cff190daa7cc10b6e5ac9716b4a2649f7c2ebcef227266d7c5983afe44cf15ea8cf565b34c6c31ff0cb4dd744524f7842b942d08770db04e5ee349086985f74b73971ce9dfe76bbed95c84906c5dffd96504e1e5396cac506ecb5465659b3a927143f6d724f91d8d9c4bdb2463aee111d9aa869874db"

	result, err := ParseWithdrawalFinalizationInput(inputData)
	if err != nil {
		t.Errorf("Failed to parse input data: %v", err)
	} else {
		for _, item := range result {
			fmt.Printf("L2BlockNumber: %v L2MessageIndex: %v L2TxNumberInBlock: %v\n", item.L2BlockNumber.Int64(), item.L2MessageIndex.Int64(), item.L2TxNumberInBlock)
		}
	}
}

func TestParseETHWithdrawalFinalizationInput(t *testing.T) {
	// Example input data (This needs to be a real encoded input data for the finalizeWithdrawals function)
	// https://etherscan.io/tx/0x61a6819c6fc9be6936e31e90aff7bbd588527f606eca7cf8226e2c89f2b70242
	inputData := "0x6c0960f90000000000000000000000000000000000000000000000000000000000061291000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000001a900000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000386c0960f972566d958675ff24bc06be3dd8ec0aa38b3b7a0b00000000000000000000000000000000000000000000000000151c37af4160000000000000000000000000000000000000000000000000000000000000000000000000000000000bf504d9d3fe266b80d29e0ec249a323800b0310cb42025ff4b6fc09a3d7e1d0fe72c1d829514f4d8afed8a77d4f21a750f155c61dcb0b3b6ed894fcb4771a061bf77b57bc1b6232a19636d619d59a6509aa38f3ad088d0daf248e220674757f2110f0882cd188ba10143931a975b5ed13a135f23afeaa37b5ff872931923422972e6f1ba9d25b971653529ed41a8678f71fbdffb059f1cc258ef5b4c4ce5ab2041798a1fd9c8fbb818c98cff190daa7cc10b6e5ac9716b4a2649f7c2ebcef227266d7c5983afe44cf15ea8cf565b34c6c31ff0cb4dd744524f7842b942d08770db04e5ee349086985f74b73971ce9dfe76bbed95c84906c5dffd96504e1e5396cac506ecb5465659b3a927143f6d724f91d8d9c4bdb2463aee111d9aa869874db124b05ec272cecd7538fdafe53b6628d31188ffb6f345139aac3c3c1fd2e470fc3be9cbd19304d84cca3d045e06b8db3acd68c304fc9cd4cbffe6d18036cb13f"

	result, err := ParseETHWithdrawalFinalizationInput(inputData)
	if err != nil {
		t.Errorf("Failed to parse input data: %v", err)
	} else {
		fmt.Printf("L2BlockNumber: %v L2MessageIndex: %v L2TxNumberInBlock: %v\n", result.L2BlockNumber.Int64(), result.L2MessageIndex.Int64(), result.L2TxNumberInBlock)

	}
}
