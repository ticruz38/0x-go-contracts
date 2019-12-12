package main

//go:generate sh -c "abigen --abi ./artifacts/OrderValidator.json --type OrderValidator --pkg order --out contracts/order/OrderValidator.go"
//go:generate sh -c "abigen --abi ./artifacts/Exchange.json --type Exchange --pkg exchange --out contracts/exchange/Exchange.go"
//go:generate sh -c "abigen --abi ./artifacts/Forwarder.json --type Forwarder --pkg forwarder --out contracts/forwarder/Forwarder.go"
//go:generate sh -c "abigen --abi ./artifacts/CoordinatorRegistry.json --type CoordinatorRegistry --pkg coordinator --out contracts/coordinator/CoordinatorRegistry.go"
//go:generate sh -c "abigen --abi ./artifacts/DutchAuction.json --type DutchAuction --pkg dutchAuction --out contracts/dutchAuction/DutchAuction.go"
//go:generate sh -c "abigen --abi ./artifacts/EthBalanceChecker.json --type EthBalanceChecker --pkg balance --out contracts/balance/EthBalanceChecker.go"
//go:generate sh -c "abigen --abi ./artifacts/DevUtils.json --type DevUtils --pkg utils --out contracts/utils/DevUtils.go"

//go:generate sh -c "abigen --abi ./artifacts/DummyERC20Token.json --type DummyERC20Token --pkg token --out contracts/token/DummyERC20Token.go"
//go:generate sh -c "abigen --abi ./artifacts/DummyERC721Token.json --type DummyERC721Token --pkg token --out contracts/token/DummyERC721Token.go"
//go:generate sh -c "abigen --abi ./artifacts/ZRXToken.json --type ZRXToken --pkg token --out contracts/token/ZRXToken.go"
//go:generate sh -c "abigen --abi ./artifacts/WETH9.json --type WETH9 --pkg token --out contracts/token/WETH9.go"

//go:generate sh -c "abigen --abi ./artifacts/AssetProxyOwner.json --type AssetProxyOwner --pkg proxy --out contracts/proxy/AssetProxyOwner.go"
//go:generate sh -c "abigen --abi ./artifacts/ERC20Proxy.json --type ERC20Proxy --pkg proxy --out contracts/proxy/ERC20Proxy.go"
//go:generate sh -c "abigen --abi ./artifacts/ERC20Token.json --type ERC20Token --pkg proxy --out contracts/proxy/ERC20Token.go"
//go:generate sh -c "abigen --abi ./artifacts/ERC721Proxy.json --type ERC721Proxy --pkg proxy --out contracts/proxy/ERC721Proxy.go"
//go:generate sh -c "abigen --abi ./artifacts/ERC721Token.json --type ERC721Token --pkg proxy --out contracts/proxy/ERC721Token.go"
//go:generate sh -c "abigen --abi ./artifacts/ERC1155Proxy.json --type ERC1155Proxy --pkg proxy --out contracts/proxy/ERC1155Proxy.go"
//go:generate sh -c "abigen --abi ./artifacts/MultiAssetProxy.json --type MultiAssetProxy --pkg proxy --out contracts/proxy/MultiAssetProxy.go"
//go:generate sh -c "abigen --abi ./artifacts/StaticCallProxy.json --type StaticCallProxy --pkg proxy --out contracts/proxy/StaticCallProxy.go"

//go:generate sh -c "abigen --abi ./artifacts/IWallet.json --type IWallet --pkg types --out contracts/types/IWallet.go"
//go:generate sh -c "abigen --abi ./artifacts/IAssetProxy.json --type IAssetProxy --pkg types --out contracts/types/IAssetProxy.go"
//go:generate sh -c "abigen --abi ./artifacts/IValidator.json --type IValidator --pkg types --out contracts/types/IValidator.go"
//go:generate sh -c "abigen --abi ./artifacts/IWallet.json --type IWallet --pkg types --out contracts/types/IWallet.go"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ABI struct {
	ABI interface{} `json:"abi,omitempty"`
}

type Artifact struct {
	SchemaVersion  string `json:"schemaVersion,omitempty"`
	ContractName   string `json:"contractName,omitempty"`
	CompilerOutput ABI
}

func main() {
	transformArtifacts()
}

const ArtifactRepo = "../0x-monorepo/packages/contract-artifacts/artifacts"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func transformArtifacts() {
	info, err := ioutil.ReadDir("../0x-monorepo/packages/contract-artifacts/artifacts")
	check(err)
	for _, v := range info {
		b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", ArtifactRepo, v.Name()))
		check(err)
		a := new(Artifact)
		json.Unmarshal(b, a)
		abi, err := json.Marshal(a.CompilerOutput.ABI)
		// s := string(abi)
		// fmt.Println(s)
		check(err)
		filePath := fmt.Sprintf("./artifacts/%s", v.Name())
		// err = ioutil.WriteFile(filePath, abi, 0644)
		f, err := os.Create(filePath)
		check(err)
		defer f.Close()
		f.Write(abi)
	}
}
