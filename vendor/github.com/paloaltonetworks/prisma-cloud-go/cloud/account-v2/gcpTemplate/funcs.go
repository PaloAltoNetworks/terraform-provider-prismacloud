package gcpTemplate

import (
	"fmt"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"os"
)

func GetGcpTemplate(c pc.PrismaCloudClient, req GcpTemplateReq) error {

	filename := req.FileName
	filename = filename + ".tf.json"

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	resp, err := c.Communicate("POST", path, nil, req, nil)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return fmt.Errorf("Invalid path: %s", filename)
	}

	_, err = file.Write(resp)
	if err != nil {
		return err
	}
	file.Close()
	return err

}
