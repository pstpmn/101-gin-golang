package config

func GetDbConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)

	serverConfig["HOST"] = "0.0.0.0"   
	serverConfig["PORT"] = "5432"          
	serverConfig["USER"] = "root"   
	serverConfig["PASS"] = "pstpmn11"   
	serverConfig["DBNAME"] = "article" 
	serverConfig["TIMEZONE"] = "Asia/Bangkok"   
	return serverConfig
}