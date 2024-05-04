package contract

// Generate Token contract bindings
//go:generate solc --overwrite --bin --optimize --abi sol/IERC20.sol -o build
//go:generate abigen --abi=./build/IERC20.abi --pkg contract --type IERC20 --out=erc20.go
//go:generate solc --overwrite --bin --optimize --abi sol/IERC721.sol -o build
//go:generate abigen --abi=./build/IERC721.abi --pkg contract --type IERC721 --out=erc721.go
//go:generate solc --overwrite --bin --optimize --abi sol/IERC1155.sol -o build
//go:generate abigen --abi=./build/IERC1155.abi --pkg contract --type IERC1155 --out=erc1155.go

// Generate L1 scroll rollup bindings
// contract/sol/IScrollChain.sol
//go:generate solc --overwrite --bin --optimize --abi sol/IScrollChain.sol -o build
//go:generate abigen --abi=./build/IScrollChain.abi --pkg contract --type IScrollChain --out=scrollchain.go

// Generate L1 zkevm bindings
// contract/abi/linea-zkevm2.abi
//go:generate abigen --abi=./abi/linea-zkevmv2.abi --pkg contract --type linea --out=linea-zkevmv2.go
//go:generate abigen --abi=./abi/linea-zkevmv2-after-block-19217957.abi --pkg upgrade --type linea --out=linea-zkevmv2-new.go

// Generate L1 optimisim bindings
// contract/abi/op-L2OutputOracle.abi
//go:generate abigen --abi=./abi/op-L2OutputOracle.abi --pkg contract --type L2OutputOracle --out=op-l2outputoracle.go

// Generate L1 zksync bindings
// contract/abi/zksync-MailboxFacet.abi
//go:generate abigen --abi=./abi/zksync-MailboxFacet.abi --pkg contract --type MailboxFacet --out=zksync-mailboxfacet.go
// contract/abi/zksync-WithdrawalFinalizer.abi
//go:generate abigen --abi=./abi/zksync-WithdrawalFinalizer.abi --pkg contract --type WithdrawalFinalizer --out=zksync-withdrawalfinalizer.go
// contract/abi/zksync-DiamondProxy.abi
//go:generate abigen --abi=./abi/zksync-DiamondProxy.abi --pkg contract --type ETHWithdrawalFinalizer --out=zksync-ethWithdrawalfinalizer.go
