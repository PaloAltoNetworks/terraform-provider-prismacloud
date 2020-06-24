package history

// Valid values for the search filter when getting a list of RQL queries.
const (
	Recent = "recent"
	Saved  = "saved"
)

const (
	singular = "historic rql query"
	plural   = "historic rql queries"
)

var Suffix = []string{"search", "history"}
