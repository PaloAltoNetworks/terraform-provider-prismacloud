package externalid

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func GetExternalId(c pc.PrismaCloudClient, req ExternalIdReq) (ExternalId, error) {
	c.Log(pc.LogAction, "(get) performing external id")

	var resp ExternalId

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)

	_, err := c.Communicate("POST", path, nil, req, &resp)
	return resp, err
}
