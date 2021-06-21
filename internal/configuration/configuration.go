package configuration

type ConfigurationKey string

type Configuration struct {
	DBDriverName            string
	GRPCPort                string
	KafkaBrokers            []string
	FlusherMaximumChankSize int
}
