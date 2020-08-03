package main

import (
	"fmt"
	"os"

	"github.com/chainHero/heroes-service/blockchain"
	//"github.com/coreos/go-systemd/util"
)

//replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		//	OrgAdmin:   "Admin",
		//	OrgName:    "Org1",
		//	ConfigFile: "config.yaml",

		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/chainVehicule-service/fixtures/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "chainVehicule-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/chainHero/chainVehicule-service/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	}
	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
	}
}
