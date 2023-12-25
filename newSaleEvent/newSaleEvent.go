package newsaleevent

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NewSaleEvent struct {
	MarketPlaceID  *big.Int
	Nfts           []common.Address
	TokenIDs       []*big.Int
	AmountBatches  []*big.Int
	Price          *big.Int
	Duration       *big.Int
	IsAuction      bool
	Antisnipe      bool
	FlashAuction   bool
	Amount         *big.Int
	IsNSFW         bool
	SearchKeywords string
	Addresses      []common.Address
	Seller         common.Address
}
