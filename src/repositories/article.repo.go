package repositories

import (
	"example.com/v1/src/models"
	"example.com/v1/src/drivers"
	"example.com/v1/src/config"
)
var configDb = config.GetDbConfig();
var conn , err = drivers.Connect(configDb);

func Index()[]models.Articles{
	var articleInfo []models.Articles
	conn.Find(&articleInfo).Limit(10)
	return articleInfo
}