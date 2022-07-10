package tokeninfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/jr64/nftdump/out"
)

func ipfsToGatewayUrl(u *url.URL, ipfsGateway string) (string, error) {
	uri := u.Host
	if u.Path != "" {
		uri = path.Join(u.Host, u.Path)
	}
	uri = strings.TrimPrefix(uri, "ipfs/")

	gatewayU, err := url.Parse(ipfsGateway)
	if err != nil {
		return "", fmt.Errorf("Failed to parse gateway uri: %v", err)
	}

	gatewayU.Path = path.Join(gatewayU.Path, uri)

	return gatewayU.String(), nil
}

func DumpUri(uri string, id *big.Int, ipfsGateway string) {
	u, err := url.Parse(uri)
	if err != nil {
		out.Fatal("Failed to parse uri: %v", err)
	}
	if u.Scheme == "https" || u.Scheme == "http" {
		out.Info("Metadata is stored on a HTTP(s) server")
	} else if u.Scheme == "ipfs" {
		out.Info("Metadata is stored in IPFS")

		uri, err = ipfsToGatewayUrl(u, ipfsGateway)

		if err != nil {
			out.Fatal("Failed to convert IPFS location to gateway URL: %v", err)
		}

	} else {
		out.Info("Metadata stored at %s uses unsupported protocol %s", uri, u.Scheme)
	}

	uri = strings.ReplaceAll(uri, "{id}", fmt.Sprintf("%064x", id))
	out.Info("Downloading metadata from %s", uri)

	resp, err := http.Get(uri)
	if err != nil {
		out.Fatal("Failed to fetch metatadata: %v", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		out.Fatal("Failed to download metatadata: %v", err)
	}

	out.Success("Download complete")

	var v interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		out.Fatal("Failed to parse metatadata: %v", err)
	}

	pre := strings.Repeat(" ", 4)
	fmt.Print(pre)
	res, err := json.MarshalIndent(v, pre, "  ")
	if err != nil {
		out.Fatal("Failed to format metatadata: %v", err)
	}
	fmt.Println(string(res))

	if j, ok := v.(map[string]interface{}); ok {
		imgLoc, ok := j["image"].(string)
		if ok && imgLoc != "" {
			out.Success("Image location from metadata: %s", imgLoc)

			imgLogU, err := url.Parse(imgLoc)
			if err == nil {
				if imgLogU.Scheme == "ipfs" {
					gatewayLink, err := ipfsToGatewayUrl(imgLogU, ipfsGateway)
					if err == nil {
						out.Info("Image location converted to gateway: %s", gatewayLink)
					}
				}
			}
		}

	}
}
