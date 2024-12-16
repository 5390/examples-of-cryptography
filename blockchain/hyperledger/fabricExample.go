package hyperledger

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Asset structure defines the model for the assets in the ledger
// It contains the ID (unique identifier) and Name (asset name).
type Asset struct {
	ID   string `json:"ID"`   // Asset ID
	Name string `json:"Name"` // Asset Name
}

// SmartContract struct embeds the `contractapi.Contract` struct,
// which is necessary to implement chaincode functions for interacting with the ledger.
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the ledger with predefined assets when the chaincode is first deployed.
// It creates two assets: "asset1" and "asset2" concurrently using goroutines.
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// Predefined assets
	assets := []Asset{
		{ID: "asset1", Name: "Asset One"},
		{ID: "asset2", Name: "Asset Two"},
	}

	var wg sync.WaitGroup // To wait for all goroutines to finish

	// Loop through each asset and create it concurrently using goroutines
	for _, asset := range assets {
		wg.Add(1)
		go func(asset Asset) {
			defer wg.Done()
			err := s.CreateAsset(ctx, asset.ID, asset.Name)
			if err != nil {
				fmt.Printf("Error creating asset %s: %v\n", asset.ID, err)
			}
		}(asset)
	}

	// Wait for all goroutines to finish before returning
	wg.Wait()
	return nil
}

// CreateAsset adds a new asset to the ledger.
// It checks if the asset already exists. If not, it stores the asset in the ledger.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, name string) error {
	// Create a new asset object
	asset := Asset{
		ID:   id,
		Name: name,
	}

	// Check if the asset already exists in the ledger by querying the state using the asset ID
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err) // Return error if reading from the state fails
	}

	// If the asset already exists, return an error
	if assetJSON != nil {
		return fmt.Errorf("asset %s already exists", id)
	}

	// Marshal the asset object into JSON format for storage
	assetJSONBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %v", err) // Return error if marshalling fails
	}

	// Put the asset data into the ledger with its ID as the key
	return ctx.GetStub().PutState(id, assetJSONBytes)
}

// QueryAsset retrieves an asset from the ledger by its ID.
// It returns the asset's details if found or an error if not found.
func (s *SmartContract) QueryAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	// Retrieve the asset from the ledger using its ID
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err) // Return error if reading from state fails
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("asset %s does not exist", id) // Return error if the asset doesn't exist
	}

	// Unmarshal the asset JSON data into an Asset struct
	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal asset JSON: %v", err) // Return error if unmarshalling fails
	}

	// Return the asset
	return &asset, nil
}

// UpdateAsset updates the name of an existing asset in the ledger.
// It retrieves the asset by its ID, modifies the name, and stores it back in the ledger.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, newName string) error {
	// Retrieve the existing asset from the ledger
	asset, err := s.QueryAsset(ctx, id)
	if err != nil {
		return err // Return error if the asset cannot be found
	}

	// Update the asset's name
	asset.Name = newName

	// Marshal the updated asset back into JSON format
	assetJSONBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %v", err) // Return error if marshalling fails
	}

	// Store the updated asset in the ledger
	return ctx.GetStub().PutState(id, assetJSONBytes)
}

// FabricExample initializes and starts the chaincode.
// It creates a new chaincode instance and handles any initialization errors.
func FabricExample() {
	// Create a new instance of the chaincode with the SmartContract implementation
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating asset chaincode: %v", err)
		return
	}

	// Start the chaincode and listen for transactions
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting asset chaincode: %v", err)
	}
}
