package tokeninfo

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jr64/nftdump/out"
	"github.com/metachris/eth-go-bindings/erc165"
	"github.com/metachris/eth-go-bindings/erc721"
)

type Erc721Info struct {
	Id               *big.Int
	SupportsMetadata bool
	Name             string
	Owner            string
	Uri              string
}

func FetchErc721Info(token *erc721.Erc721, id *big.Int) (info *Erc721Info, err error) {

	info = new(Erc721Info)

	info.Id = id
	supportsMetadata, err := token.SupportsInterface(nil, erc165.InterfaceIdErc721Metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve supportsInterface: %v", err)
	}
	info.SupportsMetadata = supportsMetadata

	if supportsMetadata {
		name, err := token.Name(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve name: %v", err)
		}
		info.Name = name

		uri, err := token.TokenURI(nil, id)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve tokenUri: %v", err)
		}
		info.Uri = uri
	}

	owner, err := token.OwnerOf(nil, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ownerOf: %v", err)
	}
	info.Owner = owner.String()

	return info, nil
}

func DumpErc721(contractAddress string, id *big.Int, client *ethclient.Client, ipfsGateway string) {
	address := common.HexToAddress(contractAddress)

	out.Info("Dumping ERC721 NFT %d from contract %s", id, contractAddress)
	token, err := erc721.NewErc721(address, client)
	if err != nil {
		out.Fatal("Failed to instantiate Token contract: %v", err)
	}

	info, errx := FetchErc721Info(token, id)
	if errx != nil {
		out.Fatal("Failed to fetch token info: %v", errx)
	}

	if !info.SupportsMetadata {
		out.Warn("Token does not support metadata")
	}

	tbl := out.NewKeyValueTable()

	tbl.Add("Id", info.Id.Text(10))
	if info.SupportsMetadata {
		tbl.Add("Name", info.Name)
	}
	tbl.Add("Owner", info.Owner)
	tbl.Add("Metadata", info.Uri)

	tbl.Print(strings.Repeat(" ", 4))

	out.Success("Done")

	DumpUri(info.Uri, info.Id, ipfsGateway)
}
