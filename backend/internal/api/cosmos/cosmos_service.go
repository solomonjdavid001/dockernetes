package cosmos

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
	"golang.org/x/net/context"
)

var (
	cosmosClient *azcosmos.Client
	once         sync.Once
)

func getCosmosClient() *azcosmos.Client {
	once.Do(func() {
		cosmosClient = initCosmosClient()
	})
	return cosmosClient
}

func initCosmosClient() *azcosmos.Client {
	endpoint := "https://<your-account>.documents.azure.com:443/"
	key := "<your-primary-key>"

	// Create a Cosmos client
	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		log.Fatalf("Failed to create credential: %v", err)
	}

	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create Cosmos client: %v", err)
	}
	return client
}

func ListDatabases(subscriptionId string, resourceGroupName string, accountName string) []string {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal("Failed to obtain a credential: ", err)
	}

	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory(subscriptionId, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create client factory: %v", err)
	}

	var dbNames []string
	pager := clientFactory.NewSQLResourcesClient().NewListSQLDatabasesPager(resourceGroupName, accountName, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("Failed to advance next page: %v", err)
		}
		for _, db := range page.Value {
			dbNames = append(dbNames, *db.Name)
		}
	}
	return dbNames
}

func GetAllItems() interface{} {
	client := getCosmosClient()

	database, err := client.NewDatabase("test")
	if err != nil {
		log.Fatalf("Failed to create database client: %v", err)
	}

	container, err := database.NewContainer("test")
	if err != nil {
		log.Fatalf("Failed to create container client: %v", err)
	}

	query := "SELECT * FROM c"
	partitonKey := azcosmos.NewPartitionKeyString("")

	pager := container.NewQueryItemsPager(query, partitonKey, nil)

	items := []Item{}

	for pager.More() {
		response, err := pager.NextPage(context.Background())
		if err != nil {
			log.Fatalf("Failed to get next page: %v", err)
		}

		for _, bytes := range response.Items {
			item := Item{}
			err := json.Unmarshal(bytes, &item)
			if err != nil {
				log.Fatalf("Failed to unmarshal item: %v", err)
			}
			items = append(items, item)
		}
	}
	return items
}
