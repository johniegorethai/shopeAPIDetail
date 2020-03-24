package config

type (
	// Config ...
	Config struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
		API      APIConfig      `yaml:"api"`
		Firebase FirebaseConfig `yaml:"firebase"`
	}

	// ServerConfig ...
	ServerConfig struct {
		Port string `yaml:"port"`
	}

	// DatabaseConfig ...
	DatabaseConfig struct {
		Master string `yaml:"master"`
	}
	//APIConfig ...
	APIConfig struct {
		Shope string `yaml:"shope"`
		Shope1 string `yaml:"shope1"`
		Shope2 string `yaml: "shope2"`
	}

	// FirebaseConfig ...
	FirebaseConfig struct {
		ProjectID string `yaml:"projectID"`
	}
)
