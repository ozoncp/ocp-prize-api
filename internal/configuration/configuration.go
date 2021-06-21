package configuration

type ConfigurationKey string

type Configuration struct {
	DBDriverName            string
	DBHost                  string
	DBPort                  int
	DBLogin                 string
	DBPassword              string
	GRPCPort                string
	KafkaBrokers            []string
	FlusherMaximumChankSize int
}
