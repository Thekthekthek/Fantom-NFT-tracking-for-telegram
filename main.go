package main

import (
	"abis"
	"context"
	"fmt"
	"log"
	"math/big"
	newsaleevent "newSaleEvent"
	"os"
	"strconv"
	"strings"
	"telegrambot"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var (
	RPC                 string
	Chat_ID             int
	BOT_TOKEN           string
	TOKEN_ADDRESS       string
	MARKETPLACE_ADDRESS string
)

func main() {
	config()
	clientWSS, err := ethclient.Dial("wss" + RPC)
	tgBot := telegrambot.BotInit(BOT_TOKEN, Chat_ID)
	defer tgBot.Bot.Close(tgBot.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	marketPlaceAddress := common.HexToAddress(MARKETPLACE_ADDRESS)
	marketPlaceAbi, err := abi.JSON(strings.NewReader(abis.MarketPlaceABI))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{marketPlaceAddress},
	}
	logs := make(chan types.Log)
	sub, err := clientWSS.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	if err != nil {
		return
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			var data newsaleevent.NewSaleEvent
			unpackedData, err := marketPlaceAbi.Unpack("NewSale", vLog.Data)
			if (len(unpackedData) == 0) || (err != nil) {
				continue
			}
			data = newsaleevent.NewSaleEvent{
				MarketPlaceID:  unpackedData[0].(*big.Int),
				Nfts:           unpackedData[1].([]common.Address),
				TokenIDs:       unpackedData[2].([]*big.Int),
				AmountBatches:  unpackedData[3].([]*big.Int),
				Price:          unpackedData[4].(*big.Int),
				Duration:       unpackedData[5].(*big.Int),
				IsAuction:      unpackedData[6].(bool),
				Antisnipe:      unpackedData[7].(bool),
				FlashAuction:   unpackedData[8].(bool),
				Amount:         unpackedData[9].(*big.Int),
				IsNSFW:         unpackedData[10].(bool),
				SearchKeywords: unpackedData[11].(string),
				Addresses:      unpackedData[12].([]common.Address),
				Seller:         unpackedData[13].(common.Address),
			}
			if err != nil {
				return
			}
			for idx, nft := range data.Nfts {
				if nft != common.HexToAddress(TOKEN_ADDRESS) {
					continue
				}
				nftID := data.TokenIDs[idx]
				if err != nil {
					fmt.Println(err)
					continue
				}
				price, _ := new(big.Int).Div(data.Price, big.NewInt(1_000_000_000_000_000_000)).Float64()
				tgBot.SendTelegramMessage(fmt.Sprintf("New vEqual NFT %d sale !\nPrice: %f FTM\nlink : https://paintswap.finance/marketplace/fantom/%d", nftID, price, data.MarketPlaceID.Int64()))
			}
		}
	}
}

func config() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Chat_ID, err = strconv.Atoi(os.Getenv("Chat_ID"))
	if err != nil {
		log.Fatal("Error loading Chat_ID var")
	}
	BOT_TOKEN = os.Getenv("BOT_TOKEN")
	if BOT_TOKEN == "" {
		log.Fatal("Error loading BOT_TOKEN var")
	}
	RPC = os.Getenv("RPC")
	if RPC == "" {
		log.Fatal("Error loading WSS_RPC var")
	}
	MARKETPLACE_ADDRESS = os.Getenv("MARKETPLACE_ADDRESS")
	if MARKETPLACE_ADDRESS == "" {
		log.Fatal("Error loading MARKETPLACE_ADDRESS var")
	}
	TOKEN_ADDRESS = os.Getenv("TOKEN_ADDRESS")
	if TOKEN_ADDRESS == "" {
		log.Fatal("Error loading WSS_TOKEN_ADDRESS var")
	}

}
