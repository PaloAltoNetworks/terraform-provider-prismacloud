package prismacloud

// These are the hidden fields blanked out during send and receive logging.
var SensitiveKeys = []string{"password", "private_key"}

// Control what is echoed to the user.
const (
	LogQuiet   = "quiet"
	LogAction  = "action"
	LogPath    = "path"
	LogSend    = "send"
	LogReceive = "receive"
)
