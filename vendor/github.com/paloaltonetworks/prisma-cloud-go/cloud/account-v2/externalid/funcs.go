package externalid

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func GetExternalId(c pc.PrismaCloudClient, req ExternalIdReq) (ExternalId, error) {
	c.Log(pc.LogAction, "(post) performing external id")

	var resp ExternalId

	_, err := c.Communicate("POST", Suffix, nil, req, &resp)
	return resp, err
}

func GetStorageUUID(c pc.PrismaCloudClient, req StorageUUID) (StorageUUID, error) {
	c.Log(pc.LogAction, "(post) getting storage_uuid")

	var resp StorageUUID

	_, err := c.Communicate("POST", StorageUUIDSuffix, nil, req, &resp)
	return resp, err
}
