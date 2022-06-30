package config

func GetWebServConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)
	serverConfig["HOST"] = "0.0.0.0"   
	serverConfig["PORT"] = "4000"          
	return serverConfig
}
