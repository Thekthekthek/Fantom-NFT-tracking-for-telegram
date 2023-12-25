package abis

const MarketPlaceABI = `[
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_officialNFTDiscount",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_marketplaceWhitelist",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_collectionRoyalties",
        "type": "address"
      },
      { "internalType": "address", "name": "_tokens", "type": "address" },
      { "internalType": "address", "name": "_paintRouter", "type": "address" },
      { "internalType": "address", "name": "_wftm", "type": "address" },
      {
        "internalType": "address",
        "name": "_implBuyEditListing",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_implMakeAndEditOffers",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_implBidsAndAcceptOffers",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_implListAndComplete",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_implBatchTransferNFT",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_feePercentage",
        "type": "address"
      },
      { "internalType": "uint256", "name": "startId", "type": "uint256" }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "_factory",
        "type": "address"
      }
    ],
    "name": "AddedVaultFactory",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "nfts",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "tokenIds",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "amountBatches",
        "type": "uint256[]"
      }
    ],
    "name": "CancelledSale",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "bidder",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "bid",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "nextMinimum",
        "type": "uint256"
      }
    ],
    "name": "NewBid",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "nft",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "from",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "quantity",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "price",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "expires",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "searchKeywords",
        "type": "string"
      }
    ],
    "name": "NewCollectionOffer",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "nft",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "tokenIds",
        "type": "uint256[]"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "from",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "quantity",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "price",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint128[]",
        "name": "prices",
        "type": "uint128[]"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "expires",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "string[]",
        "name": "searchKeywords",
        "type": "string[]"
      }
    ],
    "name": "NewFilteredCollectionOffer",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "nfts",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "tokenIds",
        "type": "uint256[]"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "from",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "quantity",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "price",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "expires",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "searchKeywords",
        "type": "string"
      }
    ],
    "name": "NewOffer",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "nfts",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "tokenIds",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "amountBatches",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "price",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "duration",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "isAuction",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "antisnipe",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "flashAuction",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "isNSFW",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "searchKeywords",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "addresses",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "seller",
        "type": "address"
      }
    ],
    "name": "NewSale",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "marketplaceId",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "address[][]",
        "name": "nfts",
        "type": "address[][]"
      },
      {
        "indexed": false,
        "internalType": "uint256[][]",
        "name": "tokenIds",
        "type": "uint256[][]"
      },
      {
        "indexed": false,
        "internalType": "uint256[][]",
        "name": "amountBatches",
        "type": "uint256[][]"
      },
      {
        "indexed": false,
        "internalType": "uint128[]",
        "name": "prices",
        "type": "uint128[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "durations",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "bool[]",
        "name": "isAuctions",
        "type": "bool[]"
      },
      {
        "indexed": false,
        "internalType": "bool[]",
        "name": "isAntisnipes",
        "type": "bool[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "amounts",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "bool[]",
        "name": "isNSFWs",
        "type": "bool[]"
      },
      {
        "indexed": false,
        "internalType": "string[]",
        "name": "searchKeywords",
        "type": "string[]"
      },
      {
        "indexed": false,
        "internalType": "address[][]",
        "name": "addresses",
        "type": "address[][]"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "sellers",
        "type": "address[]"
      }
    ],
    "name": "NewSaleBatch",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "nfts",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "tokenIds",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "amountBatches",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "price",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "duration",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "isAuction",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "antisnipe",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "flashAuction",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "isNSFW",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "searchKeywords",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "addresses",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "seller",
        "type": "address"
      }
    ],
    "name": "NewTempListing",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "nft",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "tokenId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "quantity",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      }
    ],
    "name": "OfferAccepted",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      }
    ],
    "name": "OfferCompleted",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      }
    ],
    "name": "OfferExpired",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      }
    ],
    "name": "OfferRemoved",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "previousOwner",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "OwnershipTransferred",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "_marketplaceId",
        "type": "uint256"
      }
    ],
    "name": "RemoveExpiredListing",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "_factory",
        "type": "address"
      }
    ],
    "name": "RemovedVaultFactory",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      }
    ],
    "name": "SaleFinished",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "nfts",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "tokenIds",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "amountBatches",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "price",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "buyer",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "seller",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      }
    ],
    "name": "Sold",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "endTime",
        "type": "uint256"
      }
    ],
    "name": "UpdateEndTime",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "offerId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "nft",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "tokenId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "quantity",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "newPrice",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "expires",
        "type": "uint256"
      }
    ],
    "name": "UpdateOffer",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint128",
        "name": "price",
        "type": "uint128"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "nfts",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "tokenIds",
        "type": "uint256[]"
      }
    ],
    "name": "UpdatePrice",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address[]",
        "name": "nfts",
        "type": "address[]"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "tokenIds",
        "type": "uint256[]"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "newAmount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "newAmountRemaining",
        "type": "uint256"
      }
    ],
    "name": "UpdateQuantity",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "marketplaceId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "startTime",
        "type": "uint256"
      }
    ],
    "name": "UpdateStartTime",
    "type": "event"
  },
  {
    "inputs": [],
    "name": "VERSION",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_offerId", "type": "uint256" },
      { "internalType": "address", "name": "_nft", "type": "address" },
      { "internalType": "uint256", "name": "_tokenId", "type": "uint256" },
      { "internalType": "uint128", "name": "_quantity", "type": "uint128" },
      { "internalType": "string", "name": "_searchKeywords", "type": "string" }
    ],
    "name": "acceptOffer",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256[]", "name": "_offerIds", "type": "uint256[]" },
      { "internalType": "address[]", "name": "_nfts", "type": "address[]" },
      {
        "internalType": "uint256[]",
        "name": "_tokenIdOrMarketplaceIds",
        "type": "uint256[]"
      },
      {
        "internalType": "bool[]",
        "name": "_isMarketplaceIds",
        "type": "bool[]"
      },
      {
        "internalType": "uint128[]",
        "name": "_quantities",
        "type": "uint128[]"
      },
      {
        "internalType": "string[]",
        "name": "_searchKeywords",
        "type": "string[]"
      }
    ],
    "name": "acceptOfferBatch",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_offerId", "type": "uint256" },
      { "internalType": "uint128", "name": "_quantity", "type": "uint128" },
      { "internalType": "uint256", "name": "_marketplaceId", "type": "uint256" }
    ],
    "name": "acceptOnSaleOffer",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address[]", "name": "_nfts", "type": "address[]" },
      { "internalType": "uint256[]", "name": "_tokenIds", "type": "uint256[]" },
      {
        "internalType": "uint256[]",
        "name": "_amountBatches",
        "type": "uint256[]"
      },
      { "internalType": "uint128", "name": "_price", "type": "uint128" },
      { "internalType": "uint64[]", "name": "_times", "type": "uint64[]" },
      { "internalType": "uint128", "name": "_amount", "type": "uint128" },
      { "internalType": "bool[]", "name": "_flags", "type": "bool[]" },
      { "internalType": "string", "name": "_searchKeywords", "type": "string" },
      { "internalType": "address[]", "name": "_addresses", "type": "address[]" }
    ],
    "name": "addSale",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address[][]", "name": "_nfts", "type": "address[][]" },
      {
        "internalType": "uint256[][]",
        "name": "_tokenIds",
        "type": "uint256[][]"
      },
      {
        "internalType": "uint256[][]",
        "name": "_amountBatches",
        "type": "uint256[][]"
      },
      { "internalType": "uint128[]", "name": "_prices", "type": "uint128[]" },
      { "internalType": "uint64[][]", "name": "_times", "type": "uint64[][]" },
      { "internalType": "uint128[]", "name": "_amounts", "type": "uint128[]" },
      { "internalType": "bool[][]", "name": "_flags", "type": "bool[][]" },
      {
        "internalType": "string[]",
        "name": "_searchKeywords",
        "type": "string[]"
      },
      {
        "internalType": "address[][]",
        "name": "_addresses",
        "type": "address[][]"
      }
    ],
    "name": "addSaleBatch",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "_vaultFactory", "type": "address" }
    ],
    "name": "addValidVaultFactory",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "auctionEndTimeIncreaseCutOff",
    "outputs": [{ "internalType": "uint64", "name": "", "type": "uint64" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "auctionGracePeriodForCancelling",
    "outputs": [{ "internalType": "uint64", "name": "", "type": "uint64" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_marketplaceId",
        "type": "uint256"
      },
      { "internalType": "uint128", "name": "_amount", "type": "uint128" }
    ],
    "name": "buy",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256[]",
        "name": "_marketplaceIds",
        "type": "uint256[]"
      },
      { "internalType": "uint128[]", "name": "_amounts", "type": "uint128[]" },
      { "internalType": "bool", "name": "_allowFailures", "type": "bool" }
    ],
    "name": "buyBatch",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_offerId", "type": "uint256" }
    ],
    "name": "cancelOffer",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256[]", "name": "_offerIds", "type": "uint256[]" }
    ],
    "name": "cancelOfferBatch",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_marketplaceId", "type": "uint256" }
    ],
    "name": "cancelSale",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256[]",
        "name": "_marketplaceIds",
        "type": "uint256[]"
      }
    ],
    "name": "cancelSaleBatch",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_marketplaceId", "type": "uint256" }
    ],
    "name": "completeMarketplaceEntry",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "currentMarketplaceId",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "devFeeAddress",
    "outputs": [{ "internalType": "address", "name": "", "type": "address" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_offerId", "type": "uint256" },
      { "internalType": "address", "name": "_nft", "type": "address" },
      { "internalType": "uint256", "name": "_tokenId", "type": "uint256" },
      { "internalType": "uint128", "name": "_newQuantity", "type": "uint128" },
      {
        "internalType": "uint128",
        "name": "_newOfferPrice",
        "type": "uint128"
      },
      {
        "internalType": "uint256",
        "name": "_newOfferDuration",
        "type": "uint256"
      }
    ],
    "name": "editOffer",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256[]", "name": "_offerIds", "type": "uint256[]" },
      { "internalType": "address[]", "name": "_nfts", "type": "address[]" },
      { "internalType": "uint256[]", "name": "_tokenIds", "type": "uint256[]" },
      {
        "internalType": "uint128[]",
        "name": "_newQuantities",
        "type": "uint128[]"
      },
      {
        "internalType": "uint128[]",
        "name": "_newOfferPrices",
        "type": "uint128[]"
      },
      {
        "internalType": "uint256[]",
        "name": "_newOfferDurations",
        "type": "uint256[]"
      }
    ],
    "name": "editOfferBatch",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_marketplaceId",
        "type": "uint256"
      },
      { "internalType": "uint128", "name": "_price", "type": "uint128" }
    ],
    "name": "editPrice",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_marketplaceId",
        "type": "uint256"
      },
      { "internalType": "uint128", "name": "_newPrice", "type": "uint128" },
      { "internalType": "uint128", "name": "_newQuantity", "type": "uint128" }
    ],
    "name": "editPriceAndQuantity",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256[]",
        "name": "_marketplaceId",
        "type": "uint256[]"
      },
      {
        "internalType": "uint128[]",
        "name": "_newPrices",
        "type": "uint128[]"
      },
      {
        "internalType": "uint128[]",
        "name": "_newQuantities",
        "type": "uint128[]"
      }
    ],
    "name": "editPriceAndQuantityBatch",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256[]",
        "name": "_marketplaceIds",
        "type": "uint256[]"
      },
      { "internalType": "uint128[]", "name": "_prices", "type": "uint128[]" }
    ],
    "name": "editPriceBatch",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_marketplaceId",
        "type": "uint256"
      },
      { "internalType": "uint128", "name": "_newQuantity", "type": "uint128" }
    ],
    "name": "editQuantity",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256[]",
        "name": "_marketplaceIds",
        "type": "uint256[]"
      },
      {
        "internalType": "uint128[]",
        "name": "_newQuantities",
        "type": "uint128[]"
      }
    ],
    "name": "editQuantityBatch",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256[]", "name": "_offerIds", "type": "uint256[]" }
    ],
    "name": "expireOffers",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "flashAuctionEndTimeIncreaseCutOff",
    "outputs": [{ "internalType": "uint64", "name": "", "type": "uint64" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "flashAuctionGracePeriodForCancelling",
    "outputs": [{ "internalType": "uint64", "name": "", "type": "uint64" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_offerId", "type": "uint256" }
    ],
    "name": "getOffer",
    "outputs": [
      {
        "components": [
          { "internalType": "address", "name": "nft", "type": "address" },
          { "internalType": "address[]", "name": "nfts", "type": "address[]" },
          {
            "internalType": "uint256[]",
            "name": "tokenIds",
            "type": "uint256[]"
          },
          { "internalType": "uint128", "name": "quantity", "type": "uint128" },
          {
            "internalType": "uint128",
            "name": "quantityRemaining",
            "type": "uint128"
          },
          { "internalType": "uint128", "name": "price", "type": "uint128" },
          {
            "internalType": "uint128[]",
            "name": "prices",
            "type": "uint128[]"
          },
          { "internalType": "address", "name": "from", "type": "address" },
          {
            "internalType": "enum PaintSwapMarketplaceV3StateBase.OfferType",
            "name": "offerType",
            "type": "uint8"
          },
          { "internalType": "uint256", "name": "expires", "type": "uint256" },
          { "internalType": "uint256", "name": "saleId", "type": "uint256" }
        ],
        "internalType": "struct PaintSwapMarketplaceV3StateBase.Offer",
        "name": "",
        "type": "tuple"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_marketplaceId", "type": "uint256" }
    ],
    "name": "getSaleDetails",
    "outputs": [
      {
        "components": [
          { "internalType": "address[]", "name": "nfts", "type": "address[]" },
          {
            "internalType": "uint256[]",
            "name": "tokenIds",
            "type": "uint256[]"
          },
          {
            "internalType": "uint256[]",
            "name": "amountBatches",
            "type": "uint256[]"
          },
          { "internalType": "uint256", "name": "startTime", "type": "uint256" },
          { "internalType": "uint256", "name": "endTime", "type": "uint256" },
          { "internalType": "uint128", "name": "price", "type": "uint128" },
          { "internalType": "uint128", "name": "maxBid", "type": "uint128" },
          { "internalType": "address", "name": "maxBidder", "type": "address" },
          { "internalType": "address", "name": "seller", "type": "address" },
          { "internalType": "bool", "name": "isAuction", "type": "bool" },
          { "internalType": "bool", "name": "complete", "type": "bool" },
          { "internalType": "bool", "name": "antisnipe", "type": "bool" },
          { "internalType": "bool", "name": "flashAuction", "type": "bool" },
          { "internalType": "uint128", "name": "amount", "type": "uint128" },
          {
            "internalType": "uint128",
            "name": "amountRemaining",
            "type": "uint128"
          },
          {
            "internalType": "address",
            "name": "paymentToken",
            "type": "address"
          },
          { "internalType": "address", "name": "vault", "type": "address" }
        ],
        "internalType": "struct PaintSwapMarketplaceV3StateBase.FullDetails",
        "name": "",
        "type": "tuple"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address[]", "name": "_nfts", "type": "address[]" },
      { "internalType": "uint256[]", "name": "_tokenIds", "type": "uint256[]" },
      { "internalType": "address", "name": "_user", "type": "address" }
    ],
    "name": "hasExistingActiveOffers",
    "outputs": [{ "internalType": "bool[]", "name": "", "type": "bool[]" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address[]", "name": "_nfts", "type": "address[]" },
      { "internalType": "uint256[]", "name": "_tokenIds", "type": "uint256[]" },
      { "internalType": "uint128", "name": "_quantity", "type": "uint128" },
      { "internalType": "uint128[]", "name": "_prices", "type": "uint128[]" },
      { "internalType": "uint256", "name": "_duration", "type": "uint256" },
      {
        "internalType": "string[]",
        "name": "_searchKeywords",
        "type": "string[]"
      }
    ],
    "name": "makeAltOffer",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_marketplaceId",
        "type": "uint256"
      },
      { "internalType": "uint128", "name": "_bid", "type": "uint128" }
    ],
    "name": "makeBid",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256[]",
        "name": "_marketplaceIds",
        "type": "uint256[]"
      },
      { "internalType": "uint128[]", "name": "_bids", "type": "uint128[]" }
    ],
    "name": "makeBidBatch",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "_nft", "type": "address" },
      { "internalType": "uint128", "name": "_quantity", "type": "uint128" },
      { "internalType": "uint128", "name": "_price", "type": "uint128" },
      { "internalType": "uint256", "name": "_duration", "type": "uint256" },
      { "internalType": "string", "name": "_searchKeywords", "type": "string" }
    ],
    "name": "makeCollectionOffer",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "_nft", "type": "address" },
      { "internalType": "uint256[]", "name": "_tokenIds", "type": "uint256[]" },
      { "internalType": "uint128", "name": "_quantity", "type": "uint128" },
      { "internalType": "uint128[]", "name": "_prices", "type": "uint128[]" },
      { "internalType": "uint256", "name": "_duration", "type": "uint256" },
      {
        "internalType": "string[]",
        "name": "_searchKeywords",
        "type": "string[]"
      }
    ],
    "name": "makeFilteredCollectionOffer",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address[]", "name": "_nfts", "type": "address[]" },
      {
        "internalType": "uint256[][]",
        "name": "_tokenIds",
        "type": "uint256[][]"
      },
      {
        "internalType": "uint128[]",
        "name": "_quantities",
        "type": "uint128[]"
      },
      {
        "internalType": "uint128[][]",
        "name": "_prices",
        "type": "uint128[][]"
      },
      {
        "internalType": "uint256[]",
        "name": "_durations",
        "type": "uint256[]"
      },
      {
        "internalType": "string[][]",
        "name": "_searchKeywords",
        "type": "string[][]"
      }
    ],
    "name": "makeFilteredCollectionOfferBatch",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_marketplaceId",
        "type": "uint256"
      },
      { "internalType": "uint128", "name": "_quantity", "type": "uint128" },
      { "internalType": "uint128", "name": "_price", "type": "uint128" },
      { "internalType": "uint256", "name": "_duration", "type": "uint256" }
    ],
    "name": "makeOffer",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256[]",
        "name": "_marketplaceIds",
        "type": "uint256[]"
      },
      {
        "internalType": "uint128[]",
        "name": "_quantities",
        "type": "uint128[]"
      },
      { "internalType": "uint128[]", "name": "_prices", "type": "uint128[]" },
      { "internalType": "uint256[]", "name": "_durations", "type": "uint256[]" }
    ],
    "name": "makeOfferBatch",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "_nft", "type": "address" },
      { "internalType": "uint256", "name": "_tokenId", "type": "uint256" },
      { "internalType": "uint128", "name": "_quantity", "type": "uint128" },
      { "internalType": "uint128", "name": "_price", "type": "uint128" },
      { "internalType": "uint256", "name": "_duration", "type": "uint256" },
      { "internalType": "string", "name": "_searchKeywords", "type": "string" }
    ],
    "name": "makeOfferSingle",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address[]", "name": "_nfts", "type": "address[]" },
      { "internalType": "uint256[]", "name": "_tokenIds", "type": "uint256[]" },
      {
        "internalType": "uint128[]",
        "name": "_quantities",
        "type": "uint128[]"
      },
      { "internalType": "uint128[]", "name": "_prices", "type": "uint128[]" },
      {
        "internalType": "uint256[]",
        "name": "_durations",
        "type": "uint256[]"
      },
      {
        "internalType": "string[]",
        "name": "_searchKeywords",
        "type": "string[]"
      }
    ],
    "name": "makeOfferSingleBatch",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "maxActiveOfferCount",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "maxAuctionDuration",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "maxDuration",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "maxOfferDuration",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "maxStartTimeIncrement",
    "outputs": [{ "internalType": "uint64", "name": "", "type": "uint64" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "minAuctionDuration",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "minDuration",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "minFlashAuctionDuration",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "minIncreasedBidPercent",
    "outputs": [{ "internalType": "uint128", "name": "", "type": "uint128" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "minOfferDuration",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "minVaultVersion",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_marketplaceId", "type": "uint256" }
    ],
    "name": "nextMinimumBid",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "address", "name": "", "type": "address" }],
    "name": "offerCounts",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "offerId",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [{ "internalType": "address", "name": "", "type": "address" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "_vaultFactory", "type": "address" }
    ],
    "name": "removeInvalidVaultFactory",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "renounceOwnership",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "bytes", "name": "_data", "type": "bytes" }],
    "name": "reserve1",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "bytes", "name": "_data", "type": "bytes" }],
    "name": "reserve2",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "bytes", "name": "_data", "type": "bytes" }],
    "name": "reserve3",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "bytes", "name": "_data", "type": "bytes" }],
    "name": "reserve4",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "bytes", "name": "_data", "type": "bytes" }],
    "name": "reserveOnlyOwner",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address[]", "name": "_nfts", "type": "address[]" },
      { "internalType": "uint256[]", "name": "_tokenIds", "type": "uint256[]" },
      { "internalType": "address[]", "name": "_tos", "type": "address[]" },
      { "internalType": "uint256[]", "name": "_amounts", "type": "uint256[]" },
      { "internalType": "uint256", "name": "_chainId", "type": "uint256" },
      {
        "internalType": "address[]",
        "name": "_destinations",
        "type": "address[]"
      },
      { "internalType": "bytes", "name": "_data", "type": "bytes" }
    ],
    "name": "safeNFTTransferBulk",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "components": [
          { "internalType": "address", "name": "nft", "type": "address" },
          {
            "internalType": "uint256[]",
            "name": "tokenIds",
            "type": "uint256[]"
          },
          {
            "internalType": "uint256[]",
            "name": "amounts",
            "type": "uint256[]"
          },
          { "internalType": "address", "name": "to", "type": "address" }
        ],
        "internalType": "struct PaintSwapMarketplaceV3StateBase.BulkTransferInfo[]",
        "name": "_nftsInfo",
        "type": "tuple[]"
      },
      { "internalType": "uint256", "name": "_chainId", "type": "uint256" },
      { "internalType": "address", "name": "_destination", "type": "address" },
      { "internalType": "bytes", "name": "_data", "type": "bytes" }
    ],
    "name": "safeNFTTransferBulkOrdered",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "salesFeeAddress",
    "outputs": [{ "internalType": "address", "name": "", "type": "address" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint128",
        "name": "_maxBundleNumber",
        "type": "uint128"
      },
      { "internalType": "bool", "name": "_enable", "type": "bool" }
    ],
    "name": "setBundles",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_minDuration", "type": "uint256" },
      { "internalType": "uint256", "name": "_maxDuration", "type": "uint256" },
      {
        "internalType": "uint256",
        "name": "_minAuctionDuration",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_minFlashAuctionDuration",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_maxAuctionDuration",
        "type": "uint256"
      }
    ],
    "name": "setDurations",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "bool", "name": "_enable", "type": "bool" }],
    "name": "setEnableSelling",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "_dev", "type": "address" },
      { "internalType": "address", "name": "_salesFee", "type": "address" }
    ],
    "name": "setFeeAddresses",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "_feePercentage", "type": "address" }
    ],
    "name": "setFeePercentage",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_implBatchTransferNFT",
        "type": "address"
      }
    ],
    "name": "setImplBatchTransferNFT",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_implBidsAndAcceptOffers",
        "type": "address"
      }
    ],
    "name": "setImplBidsAndAcceptOffers",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_implBuyEditListing",
        "type": "address"
      }
    ],
    "name": "setImplBuyEditListing",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_implListAndComplete",
        "type": "address"
      }
    ],
    "name": "setImplListAndComplete",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_maxActiveOfferCount",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_minVaultVersion",
        "type": "uint256"
      }
    ],
    "name": "setMaxActiveOfferCountAndMinVaultVersion",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_minDuration", "type": "uint256" },
      { "internalType": "uint256", "name": "_maxDuration", "type": "uint256" }
    ],
    "name": "setOfferDuration",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "uint256", "name": "_maxRoyalty", "type": "uint256" },
      {
        "internalType": "uint64",
        "name": "_maxStartTimeIncrement",
        "type": "uint64"
      },
      {
        "internalType": "address",
        "name": "_officialNFTDiscount",
        "type": "address"
      },
      { "internalType": "uint256", "name": "_callGasLimit", "type": "uint256" },
      { "internalType": "address", "name": "_paintRouter", "type": "address" }
    ],
    "name": "setVarious",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint64",
        "name": "_auctionGracePeriodForCancelling",
        "type": "uint64"
      },
      {
        "internalType": "uint64",
        "name": "_flashAuctionGracePeriodForCancelling",
        "type": "uint64"
      },
      {
        "internalType": "uint64",
        "name": "_auctionEndTimeIncreaseCutOff",
        "type": "uint64"
      },
      {
        "internalType": "uint64",
        "name": "_flashAuctionEndTimeIncreaseCutOff",
        "type": "uint64"
      },
      {
        "internalType": "uint128",
        "name": "_minIncreasedBidPercent",
        "type": "uint128"
      },
      {
        "internalType": "bool",
        "name": "_modifyInitialStartTime",
        "type": "bool"
      },
      { "internalType": "bool", "name": "_modifyPostStartTime", "type": "bool" }
    ],
    "name": "setVariousTimes",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_implMakeAndEditOffers",
        "type": "address"
      }
    ],
    "name": "setimplMakeAndEditOffers",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "", "type": "address" },
      { "internalType": "address", "name": "", "type": "address" },
      { "internalType": "uint256", "name": "", "type": "uint256" }
    ],
    "name": "singleNFTListings",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "newOwner", "type": "address" }
    ],
    "name": "transferOwnership",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [{ "internalType": "address", "name": "", "type": "address" }],
    "name": "userAltOffersMade",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "", "type": "address" },
      { "internalType": "uint256", "name": "", "type": "uint256" }
    ],
    "name": "userBundleOffersMade",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "", "type": "address" },
      { "internalType": "address", "name": "", "type": "address" }
    ],
    "name": "userCollectionOffersMade",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "", "type": "address" },
      { "internalType": "address", "name": "", "type": "address" }
    ],
    "name": "userFilteredCollectionOffersMade",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      { "internalType": "address", "name": "", "type": "address" },
      { "internalType": "address", "name": "", "type": "address" },
      { "internalType": "uint256", "name": "", "type": "uint256" }
    ],
    "name": "userOffersMade",
    "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }],
    "stateMutability": "view",
    "type": "function"
  },
  { "stateMutability": "payable", "type": "receive" }
]
`
