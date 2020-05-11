/*
Package prismacloud is an SDK meant to assist in interacting with the Palo
Alto Networks Prisma Cloud API.

To connect, create a client connetion with the desired params and then
initialize the connection:

    package main

    import (
        "log"
        "github.com/paloaltonetworks/prisma-cloud-go"
        "github.com/paloaltonetworks/prisma-cloud-go/compliance/standard"
    )

    func main() {
        client := &prismacloud.Client{}
        if err := c.Initialize("creds.json"); err != nil {
            log.Fatalf("Failed to connect: %s", err)
        }

        listing, err := standard.List(client)
        if err != nil {
            log.Fatalf("Failed to get compliance standards: %s", err)
        }

        log.Printf("Compliance standards:")
        for _, elm := range listing {
            log.Printf("* (%s) %s", elm.Id, elm.Name)
        }
    }

In most cases the struct and types match what the Prisma Cloud API
specifies, so you may find it useful to refer to the Prisma Cloud API
for further information: https://api.docs.prismacloud.io/reference
*/
package prismacloud
