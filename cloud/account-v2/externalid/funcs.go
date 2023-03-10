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
