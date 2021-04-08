package history

import (
	"net/url"
	"strconv"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List lists saved or recent RQL search queries.
func List(c pc.PrismaCloudClient, filter string, limit int) ([]NameId, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	query := url.Values{}
	query.Add("filter", filter)
	if limit != 0 {
		query.Add("limit", strconv.Itoa(limit))
	}

	var ans []NameId
	_, err := c.Communicate("GET", Suffix, query, nil, &ans)

	return ans, err
}

// Identify returns the ID for the given account group.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, name)

	listing, err := List(c, Recent, 0)
	if err != nil {
		return "", err
	}

	for _, o := range listing {
		if o.Model.Name == name {
			return o.Model.Id, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// Get returns an historic RQL search query.
func Get(c pc.PrismaCloudClient, id string) (Query, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Query
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Save saves an historic RQL search query to the saved searches list.
func Save(c pc.PrismaCloudClient, req SavedSearch) error {
	c.Log(pc.LogAction, "(create) saved search: %s", req.Id)

	// Sanity check the time range.
	if err := req.TimeRange.SetType(); err != nil {
		return err
	}

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, req.Id)

	_, err := c.Communicate("POST", path, nil, req, nil)
	return err
}

// Delete removes an existing saved search query.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s: %s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}
