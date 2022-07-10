# nftdump

Nftdump is a simple tool to show the actual data a NFT stores in the blockchain.

Surprising absolutely no one, a lot of NFTs store only a plain old HTTP(s) link in the blockchain, completely missing the point and making an already bad idea even worse.

**Features:**
* supports fetching metadata URLs for ERC721 and ERC1155 tokens
* download and display metadata from HTTP(s) or IPFS
* parse metadata and show URL of the actual image

##  Build

```
go build
```

## Usage

ERC721:
```
./nftdump -erc721 <contract address> <token ID>
```

ERC1155:
```
./nftdump -erc1155 <contract address> <token ID>
```

By default, nftdump uses Cloudflare IPFS and Ethereum gateways. You can change them by specifying `-ipfs-gateway` and `-ethereum-gateway`.


## Known issues

```
Failed to parse metatadata: invalid character '<' looking for beginning of value
```

If you get this error when accessing Cloudflare's IPFS gateway, it means they have started displaying a captcha to you which of course breaks the tool.

