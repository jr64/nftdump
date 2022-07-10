package tokeninfo

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jr64/nftdump/out"
	"github.com/metachris/eth-go-bindings/erc1155"
)

type Erc155Info struct {
	Id  *big.Int
	Uri string
}

func FetchErc1155Info(token *erc1155.Erc1155, id *big.Int) (info *Erc155Info, err error) {

	info = new(Erc155Info)

	info.Id = id

	uri, err := token.Uri(nil, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve uri: %v", err)
	}
	info.Uri = uri

	return info, nil

}

func DumpErc1155(contractAddress string, id *big.Int, client *ethclient.Client, ipfsGateway string) {
	address := common.HexToAddress(contractAddress)

	out.Info("Dumping ERC1155 NFT %d from contract %s", id, contractAddress)
	token, err := erc1155.NewErc1155(address, client)
	if err != nil {
		out.Fatal("Failed to instantiate Token contract: %v", err)
	}

	info, errx := FetchErc1155Info(token, id)
	if errx != nil {
		out.Fatal("Failed to fetch token info: %v", errx)
	}

	tbl := out.NewKeyValueTable()

	tbl.Add("Id", info.Id.Text(10))
	tbl.Add("Metadata", info.Uri)

	tbl.Print(strings.Repeat(" ", 4))

	out.Success("Done")

	DumpUri(info.Uri, info.Id, ipfsGateway)
}
