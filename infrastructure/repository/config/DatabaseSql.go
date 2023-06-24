package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type InfoDatabaseSQL struct {
	Read struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		DriverConn string
	}
	Write struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		DriverConn string
	}
}

func (infoDB *InfoDatabaseSQL) getDriverConn(nameMap string) (err error) {
	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()
	if err != nil {
		println("error 1")
		return err
	}

	print(nameMap)
	err = mapstructure.Decode(viper.GetStringMap(nameMap), infoDB)
	if err != nil {
		println("error 2")
		return
	}

	infoDB.Read.DriverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		infoDB.Read.Username, infoDB.Read.Password, infoDB.Read.Hostname, infoDB.Read.Port, infoDB.Read.Name)
	infoDB.Write.DriverConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		infoDB.Write.Username, infoDB.Write.Password, infoDB.Write.Hostname, infoDB.Write.Port, infoDB.Write.Name)

	println(infoDB.Read.DriverConn)
	println(infoDB.Write.DriverConn)
	return nil
}
