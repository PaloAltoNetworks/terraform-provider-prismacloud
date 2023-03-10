package supportedFeatures

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func GetSupportedFeatures(c pc.PrismaCloudClient, req SupportedFeaturesReq) (SupportedFeatures, error) {
	c.Log(pc.LogAction, "(post) performing fetch supported features")

	var resp SupportedFeatures
	path := make([]string, 0, len(Suffix)+2)
	path = append(path, Suffix...)
	path = append(path, req.CloudType)
	_, err := c.Communicate("POST", path, nil, req, &resp)
	return resp, err
}
