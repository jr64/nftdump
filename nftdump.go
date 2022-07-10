package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jr64/nftdump/out"
	"github.com/jr64/nftdump/tokeninfo"
)

func usage() {
	fmt.Printf("Usage: %s [OPTION]... [NFT ID]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	jsonRpcApi := ""
	ipfsGateway := ""
	erc721Addr := ""
	erc1155Addr := ""

	flag.Usage = usage
	flag.StringVar(&jsonRpcApi, "ethereum-gateway", "https://cloudflare-eth.com", "Specify path to Ethereum JSON-RPC gateway.")
	flag.StringVar(&ipfsGateway, "ipfs-gateway", "https://cloudflare-ipfs.com/ipfs/", "Specify path to IPFS Gateway.")
	flag.StringVar(&erc721Addr, "erc721", "", "Specify ERC721 contract address")
	flag.StringVar(&erc1155Addr, "erc1155", "", "Specify ERC1155 contract address")

	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		usage()
		os.Exit(1)
	}

	out.Logo()

	id := new(big.Int)
	id, ok := id.SetString(args[0], 10)
	if !ok {
		out.Fatal("Failed to parse NFT ID %s", args[0])
	}

	if erc721Addr == "" && erc1155Addr == "" {
		usage()
		os.Exit(1)
	}

	client, err := ethclient.Dial(jsonRpcApi)
	if err != nil {
		out.Fatal("Failed to connect to Ethereum JSON-RPC API: %v", err)
	}

	if erc1155Addr != "" {
		tokeninfo.DumpErc1155(erc1155Addr, id, client, ipfsGateway)
	}

	if erc721Addr != "" {
		tokeninfo.DumpErc721(erc721Addr, id, client, ipfsGateway)
	}

}
