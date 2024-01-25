package dao

import (
	"jade_backend/internal/dao/model"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DBCfg struct {
	DSN string
}

func ConnectDatabase() {
	var cfg DBCfg
	err := viper.Sub("Database").UnmarshalExact(&cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	Db, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	Db.AutoMigrate(&model.Jade{})

	if viper.GetString("App.RunLevel") == "debug" {
		Db = Db.Debug()
	}
}
