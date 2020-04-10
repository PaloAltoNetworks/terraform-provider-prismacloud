package prismacloud

type PrismaCloudClient interface {
	Initialize(string) error
	Authenticate() error
	Communicate(string, []string, interface{}, interface{}, interface{}) ([]byte, error)
	Log(string, string, ...interface{})
}
