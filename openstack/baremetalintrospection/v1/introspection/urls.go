package introspection

import "github.com/gophercloud/gophercloud"

func listIntrospectionsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("v1/introspection")
}

func introspectionURL(client *gophercloud.ServiceClient, nodeID string) string {
	return client.ServiceURL("v1/introspection", nodeID)
}

func abortIntrospectionURL(client *gophercloud.ServiceClient, nodeID string) string {
	return client.ServiceURL("v1/introspection", nodeID, "abort")
}

func introspectionDataURL(client *gophercloud.ServiceClient, nodeID string) string {
	return client.ServiceURL("v1/introspection", nodeID, "data")
}

func introspectionUnprocessedDataURL(client *gophercloud.ServiceClient, nodeID string) string {
	return client.ServiceURL("v1/introspection", nodeID, "data", "unprocessed")
}
