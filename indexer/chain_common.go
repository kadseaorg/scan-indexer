package indexer

import "github.com/rs/zerolog/log"

type Chain struct {
	Name                  string
	RollupContractAddress string
	DiamondProxy          string
	WithdrawalFinalizer   string
}

var (
	ArbitrumOne = Chain{
		Name:                  "arbitrum-one",
		RollupContractAddress: "0x1c479675ad559DC151F6Ec7ed3FbF8ceE79582B6",
	}
	ScrollSepolia = Chain{
		Name:                  "scroll-sepolia",
		RollupContractAddress: "0x2D567EcE699Eabe5afCd141eDB7A4f2D0D6ce8a0",
	}
	Scroll = Chain{
		Name:                  "scroll",
		RollupContractAddress: "0xa13BAF47339d63B743e7Da8741db5456DAc1E556",
	}
	ZkSyncEra = Chain{
		Name:                  "zksync-era",
		RollupContractAddress: "",
		DiamondProxy:          "0x32400084C286CF3E17e7B677ea9583e60a000324",
		WithdrawalFinalizer:   "0xf8A16864D8De145A266a534174305f881ee2315e",
	}
	ZkSyncEraSepolia = Chain{
		Name:                  "zksync-era-sepolia",
		RollupContractAddress: "",
		DiamondProxy:          "0x9a6de0f62Aa270A8bCB1e2610078650D539B1Ef9",
		WithdrawalFinalizer:   "0xe566fDf458C6f9Cf77E7F96C3dDF21030Bf7f0ec",
	}
	Linea = Chain{
		Name:                  "linea",
		RollupContractAddress: "",
	}
	Base = Chain{
		Name:                  "base",
		RollupContractAddress: "",
	}
	MantaPacific = Chain{
		Name:                  "manta-pacific",
		RollupContractAddress: "",
	}
	BsquaredTestnet = Chain{
		Name:                  "bsquared-testnet",
		RollupContractAddress: "0xDdee8ddfA81F5E36373637240038DCCC14529BF7",
	}
	PolygonZkEVM = Chain{
		Name:                  "polygon-zkevm",
		RollupContractAddress: "0x5132a183e9f3cb7c848b0aac5ae0c4f0491b7ab2",
	}
	OKX1Sepolia = Chain{
		Name:                  "okx1-sepolia",
		RollupContractAddress: "0x6662621411A8DACC3cA7049C8BddABaa9a999ce3",
	}
	OroTestnet = Chain{
		Name:                  "oro-testnet",
		RollupContractAddress: "",
	}
	Kadsea = Chain{
		Name: "kadsea",
	}
	KadseaTestnet = Chain{
		Name: "kadsea-testnet",
	}
)

func GetChain(chainName string) Chain {
	var chain Chain
	switch chainName {
	case ArbitrumOne.Name:
		chain = ArbitrumOne
	case ScrollSepolia.Name:
		chain = ScrollSepolia
	case Scroll.Name:
		chain = Scroll
	case ZkSyncEra.Name:
		chain = ZkSyncEra
	case ZkSyncEraSepolia.Name:
		chain = ZkSyncEraSepolia
	case Linea.Name:
		chain = Linea
	case Base.Name:
		chain = Base
	case MantaPacific.Name:
		chain = MantaPacific
	case BsquaredTestnet.Name:
		chain = BsquaredTestnet
	case PolygonZkEVM.Name:
		chain = PolygonZkEVM
	case OKX1Sepolia.Name:
		chain = OKX1Sepolia
	case OroTestnet.Name:
		chain = OroTestnet
	case Kadsea.Name:
		chain = Kadsea
	case KadseaTestnet.Name:
		chain = KadseaTestnet
	default:
		log.Fatal().Msgf("Unsupported chain: %s", chainName)
	}
	return chain
}
