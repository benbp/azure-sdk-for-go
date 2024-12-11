//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraClusterListBySubscription.json
func ExampleCassandraClustersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraClustersClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListClusters = armcosmos.ListClusters{
		// 	Value: []*armcosmos.ClusterResource{
		// 		{
		// 			Name: to.Ptr("cassandra-prod"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DocumentDB/cassandraClusters"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.ClusterResourceProperties{
		// 				AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
		// 				CassandraVersion: to.Ptr("3.11"),
		// 				ClientCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				ClusterNameOverride: to.Ptr("ClusterNameIllegalForAzureResource"),
		// 				DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
		// 				ExternalGossipCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				ExternalSeedNodes: []*armcosmos.SeedNode{
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.4"),
		// 				}},
		// 				GossipCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				HoursBetweenBackups: to.Ptr[int32](24),
		// 				ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
		// 				SeedNodes: []*armcosmos.SeedNode{
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.4"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.4"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraClusterListByResourceGroup.json
func ExampleCassandraClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraClustersClient().NewListByResourceGroupPager("cassandra-prod-rg", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListClusters = armcosmos.ListClusters{
		// 	Value: []*armcosmos.ClusterResource{
		// 		{
		// 			Name: to.Ptr("cassandra-prod"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.ClusterResourceProperties{
		// 				AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
		// 				CassandraVersion: to.Ptr("3.11"),
		// 				ClientCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				ClusterNameOverride: to.Ptr("ClusterNameIllegalForAzureResource"),
		// 				DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
		// 				ExternalGossipCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				ExternalSeedNodes: []*armcosmos.SeedNode{
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.4"),
		// 				}},
		// 				GossipCertificates: []*armcosmos.Certificate{
		// 					{
		// 						Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
		// 				}},
		// 				HoursBetweenBackups: to.Ptr[int32](24),
		// 				ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
		// 				SeedNodes: []*armcosmos.SeedNode{
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.52.221.4"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.4"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraClusterGet.json
func ExampleCassandraClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCassandraClustersClient().Get(ctx, "cassandra-prod-rg", "cassandra-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterResource = armcosmos.ClusterResource{
	// 	Name: to.Ptr("cassandra-prod"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ClusterResourceProperties{
	// 		AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
	// 		CassandraVersion: to.Ptr("3.11"),
	// 		ClientCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		ClusterNameOverride: to.Ptr("ClusterNameIllegalForAzureResource"),
	// 		DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
	// 		ExternalGossipCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		ExternalSeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.4"),
	// 		}},
	// 		GossipCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		HoursBetweenBackups: to.Ptr[int32](24),
	// 		ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
	// 		SeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.4"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.4"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraClusterDelete.json
func ExampleCassandraClustersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginDelete(ctx, "cassandra-prod-rg", "cassandra-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraClusterCreate.json
func ExampleCassandraClustersClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginCreateUpdate(ctx, "cassandra-prod-rg", "cassandra-prod", armcosmos.ClusterResource{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ClusterResourceProperties{
			AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
			CassandraVersion:     to.Ptr("3.11"),
			ClientCertificates: []*armcosmos.Certificate{
				{
					Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
				}},
			ClusterNameOverride:         to.Ptr("ClusterNameIllegalForAzureResource"),
			DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
			ExternalGossipCertificates: []*armcosmos.Certificate{
				{
					Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
				}},
			ExternalSeedNodes: []*armcosmos.SeedNode{
				{
					IPAddress: to.Ptr("10.52.221.2"),
				},
				{
					IPAddress: to.Ptr("10.52.221.3"),
				},
				{
					IPAddress: to.Ptr("10.52.221.4"),
				}},
			HoursBetweenBackups:           to.Ptr[int32](24),
			InitialCassandraAdminPassword: to.Ptr("mypassword"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterResource = armcosmos.ClusterResource{
	// 	Name: to.Ptr("cassandra-prod"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ClusterResourceProperties{
	// 		AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
	// 		CassandraVersion: to.Ptr("3.11"),
	// 		ClientCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		ClusterNameOverride: to.Ptr("ClusterNameIllegalForAzureResource"),
	// 		DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
	// 		ExternalGossipCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		ExternalSeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.4"),
	// 		}},
	// 		GossipCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		HoursBetweenBackups: to.Ptr[int32](24),
	// 		ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
	// 		SeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.4"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.4"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraClusterPatch.json
func ExampleCassandraClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginUpdate(ctx, "cassandra-prod-rg", "cassandra-prod", armcosmos.ClusterResource{
		Tags: map[string]*string{
			"owner": to.Ptr("mike"),
		},
		Properties: &armcosmos.ClusterResourceProperties{
			AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodNone),
			ExternalGossipCertificates: []*armcosmos.Certificate{
				{
					Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
				}},
			ExternalSeedNodes: []*armcosmos.SeedNode{
				{
					IPAddress: to.Ptr("10.52.221.2"),
				},
				{
					IPAddress: to.Ptr("10.52.221.3"),
				},
				{
					IPAddress: to.Ptr("10.52.221.4"),
				}},
			HoursBetweenBackups: to.Ptr[int32](12),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterResource = armcosmos.ClusterResource{
	// 	Name: to.Ptr("cassandra-prod"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ClusterResourceProperties{
	// 		AuthenticationMethod: to.Ptr(armcosmos.AuthenticationMethodCassandra),
	// 		CassandraVersion: to.Ptr("3.11"),
	// 		ClientCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		ClusterNameOverride: to.Ptr("ClusterNameIllegalForAzureResource"),
	// 		DelegatedManagementSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
	// 		ExternalGossipCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		ExternalSeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.4"),
	// 		}},
	// 		GossipCertificates: []*armcosmos.Certificate{
	// 			{
	// 				Pem: to.Ptr("-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----"),
	// 		}},
	// 		HoursBetweenBackups: to.Ptr[int32](24),
	// 		ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
	// 		SeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.52.221.4"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.4"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraCommand.json
func ExampleCassandraClustersClient_BeginInvokeCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginInvokeCommand(ctx, "cassandra-prod-rg", "cassandra-prod", armcosmos.CommandPostBody{
		Command: to.Ptr("nodetool status"),
		Host:    to.Ptr("10.0.1.12"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraClusterDeallocate.json
func ExampleCassandraClustersClient_BeginDeallocate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginDeallocate(ctx, "cassandra-prod-rg", "cassandra-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraClusterStart.json
func ExampleCassandraClustersClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraClustersClient().BeginStart(ctx, "cassandra-prod-rg", "cassandra-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ded6306d00ae294c24211e5069c1f56b15ba8ef5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBManagedCassandraStatus.json
func ExampleCassandraClustersClient_Status() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCassandraClustersClient().Status(ctx, "cassandra-prod-rg", "cassandra-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CassandraClusterPublicStatus = armcosmos.CassandraClusterPublicStatus{
	// 	ConnectionErrors: []*armcosmos.ConnectionError{
	// 	},
	// 	DataCenters: []*armcosmos.CassandraClusterPublicStatusDataCentersItem{
	// 		{
	// 			Name: to.Ptr("cassandra-westus2-vmss"),
	// 			Nodes: []*armcosmos.ComponentsM9L909SchemasCassandraclusterpublicstatusPropertiesDatacentersItemsPropertiesNodesItems{
	// 				{
	// 					Address: to.Ptr("10.0.8.10"),
	// 					CassandraProcessStatus: to.Ptr("Running"),
	// 					CPUUsage: to.Ptr[float64](0.9),
	// 					DiskFreeKB: to.Ptr[int64](1001260056),
	// 					DiskUsedKB: to.Ptr[int64](749836),
	// 					HostID: to.Ptr("8ccc00a7-9e22-4ac2-aa63-e5327aa0dc51"),
	// 					Load: to.Ptr("84.61 MiB"),
	// 					MemoryBuffersAndCachedKB: to.Ptr[int64](10817580),
	// 					MemoryFreeKB: to.Ptr[int64](35731832),
	// 					MemoryTotalKB: to.Ptr[int64](57610444),
	// 					MemoryUsedKB: to.Ptr[int64](11061032),
	// 					Rack: to.Ptr("rack1"),
	// 					Size: to.Ptr[int32](16),
	// 					State: to.Ptr(armcosmos.NodeStateNormal),
	// 					Status: to.Ptr("Up"),
	// 					Timestamp: to.Ptr("10/05/2021, 14:35:20.028"),
	// 					Tokens: []*string{
	// 						to.Ptr("-7785254003841681178"),
	// 						to.Ptr("-6756518164561476165"),
	// 						to.Ptr("-4269720020504480060"),
	// 						to.Ptr("-2697657908817852783"),
	// 						to.Ptr("-2622387897201218249"),
	// 						to.Ptr("-2177258200443022948"),
	// 						to.Ptr("-129203941752875603"),
	// 						to.Ptr("1738406920822650025"),
	// 						to.Ptr("2598053056312400240"),
	// 						to.Ptr("3227346424117996872"),
	// 						to.Ptr("3644347038875498376"),
	// 						to.Ptr("4823546474906050895"),
	// 						to.Ptr("5293156680707616627"),
	// 						to.Ptr("5485806686603256557"),
	// 						to.Ptr("6250058817756969464"),
	// 						to.Ptr("7991328665766267709")},
	// 					},
	// 					{
	// 						Address: to.Ptr("10.0.8.11"),
	// 						CassandraProcessStatus: to.Ptr("Running"),
	// 						CPUUsage: to.Ptr[float64](1.2),
	// 						DiskFreeKB: to.Ptr[int64](1001283028),
	// 						DiskUsedKB: to.Ptr[int64](726864),
	// 						HostID: to.Ptr("14092117-4f77-4ec0-8984-c5b860b88a47"),
	// 						Load: to.Ptr("59.67 MiB"),
	// 						MemoryBuffersAndCachedKB: to.Ptr[int64](10691520),
	// 						MemoryFreeKB: to.Ptr[int64](35863248),
	// 						MemoryTotalKB: to.Ptr[int64](57610444),
	// 						MemoryUsedKB: to.Ptr[int64](11055676),
	// 						Rack: to.Ptr("rack2"),
	// 						Size: to.Ptr[int32](16),
	// 						State: to.Ptr(armcosmos.NodeStateNormal),
	// 						Status: to.Ptr("Up"),
	// 						Timestamp: to.Ptr("10/05/2021, 14:35:20.028"),
	// 						Tokens: []*string{
	// 							to.Ptr("-8726238055665903210"),
	// 							to.Ptr("-6687985003871487470"),
	// 							to.Ptr("-5269140854976433359"),
	// 							to.Ptr("-3989177686905645288"),
	// 							to.Ptr("-3957362495277148220"),
	// 							to.Ptr("-2539287458896988944"),
	// 							to.Ptr("-2460716365393303466"),
	// 							to.Ptr("-1848370030729221440"),
	// 							to.Ptr("137707733677015122"),
	// 							to.Ptr("579467328507000597"),
	// 							to.Ptr("1698264534774619627"),
	// 							to.Ptr("1904235159942090722"),
	// 							to.Ptr("3312208865519999146"),
	// 							to.Ptr("4035940456270983993"),
	// 							to.Ptr("4412314431151736777"),
	// 							to.Ptr("8232565668795426078")},
	// 						},
	// 						{
	// 							Address: to.Ptr("10.0.8.12"),
	// 							CassandraProcessStatus: to.Ptr("Running"),
	// 							CPUUsage: to.Ptr[float64](0.4),
	// 							DiskFreeKB: to.Ptr[int64](1001252352),
	// 							DiskUsedKB: to.Ptr[int64](757540),
	// 							HostID: to.Ptr("e16ada14-39db-462b-9f9e-5b5f6beb8bbd"),
	// 							Load: to.Ptr("93.2 MiB"),
	// 							MemoryBuffersAndCachedKB: to.Ptr[int64](10453856),
	// 							MemoryFreeKB: to.Ptr[int64](36104980),
	// 							MemoryTotalKB: to.Ptr[int64](57610444),
	// 							MemoryUsedKB: to.Ptr[int64](11051608),
	// 							Rack: to.Ptr("rack3"),
	// 							Size: to.Ptr[int32](16),
	// 							State: to.Ptr(armcosmos.NodeStateNormal),
	// 							Status: to.Ptr("Up"),
	// 							Timestamp: to.Ptr("10/05/2021, 14:35:20.028"),
	// 							Tokens: []*string{
	// 								to.Ptr("-5679481051867296621"),
	// 								to.Ptr("-4574115287969297989"),
	// 								to.Ptr("-3444578133211470522"),
	// 								to.Ptr("-2755931580714972271"),
	// 								to.Ptr("-2304431590844389550"),
	// 								to.Ptr("-1961946736975068713"),
	// 								to.Ptr("-940120277889446704"),
	// 								to.Ptr("554469308917912318"),
	// 								to.Ptr("1030447162050118004"),
	// 								to.Ptr("2745632329542596589"),
	// 								to.Ptr("4564547712926446283"),
	// 								to.Ptr("5185613478135944116"),
	// 								to.Ptr("7280237939830623824"),
	// 								to.Ptr("7504213835759531710"),
	// 								to.Ptr("7631994478195429959"),
	// 								to.Ptr("8139769477321226157")},
	// 						}},
	// 						SeedNodes: []*string{
	// 							to.Ptr("10.0.8.10"),
	// 							to.Ptr("10.0.8.11"),
	// 							to.Ptr("10.0.8.12")},
	// 					}},
	// 					ETag: to.Ptr("A350A2CE7E91B6D5A102A5E5EC222B882D981092"),
	// 					ReaperStatus: &armcosmos.ManagedCassandraReaperStatus{
	// 						Healthy: to.Ptr(true),
	// 						RepairRunIDs: map[string]*string{
	// 						},
	// 						RepairSchedules: map[string]*string{
	// 							"00000000-0000-0001-0000-000000000000": to.Ptr("ACTIVE"),
	// 						},
	// 					},
	// 				}
}
