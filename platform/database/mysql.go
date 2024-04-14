package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // load driver for Mysql
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/utils"
)

func MySQLConnection() (*ent.Client, error) {
	// Max_Connections := viper.GetInt("datasource.max_connections")
	// Max_Idlr_Connections := viper.GetInt("datasource.max_idle_connections")
	// Max_Lifetime_Connections := viper.GetDuration("datasource.max_lifetime_connections")
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}
	client, err := ent.Open("mysql", mysqlConnURL, ent.Debug())
	if err != nil {
		return nil, fmt.Errorf("错误,无法连接数据库, %w", err)
	}
	// db.SetMaxOpenConns(Max_Connections)             // 最大打开的连接数 不超过数据库服务自身支持的并发连接数
	// db.SetMaxIdleConns(Max_Idlr_Connections)        // 最大闲置的连接数 一般建议maxIdleConns的值为MaxOpenConns的1/2
	// db.SetConnMaxLifetime(Max_Lifetime_Connections) // 连接的最大可复用时间 不超过数据库的超时参数值。

	return client, nil
}
