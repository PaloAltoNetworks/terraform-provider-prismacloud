package prismacloud

// These are the hidden fields blanked out during send and receive logging.
var SensitiveKeys = []string{"password", "private_key", "external_id"}

// Control what is echoed to the user.
const (
	LogQuiet   = "quiet"
	LogAction  = "action"
	LogPath    = "path"
	LogSend    = "send"
	LogReceive = "receive"
)
