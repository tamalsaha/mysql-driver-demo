package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"net"
	"net/url"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	cfg, err := mysql.ParseDSN("username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value")
	if err != nil {
		panic(err)
	}
	cfg.User = "user"
	cfg.Passwd = "password"
	// cfg.DBName = "adb"
	fmt.Println(cfg.FormatDSN())
	//
	//_, err := sql.Open("mysql", "tcp(mysql-demo.mysql.svc:3306)/")
	//if err != nil {
	//	panic(err)
	//}
}

func CanonicalMySQLDSN(in string) (string, error) {
	_, err := mysql.ParseDSN(in)
	if err != nil {
		return in, nil
	}
	URL, err := url.Parse(in)
	if err != nil {
		return "", err
	}

	rebuild := mysql.NewConfig()
	if URL.User != nil {
		rebuild.User = URL.User.Username()
		if pass, found := URL.User.Password(); found {
			rebuild.Passwd = pass
		}
	}
	rebuild.Net = URL.Scheme
	rebuild.Addr = URL.Host
	rebuild.DBName = URL.Path
	rebuild.Params = map[string]string{}
	for k, v := range URL.Query() {
		rebuild.Params[k] = v[0]
	}
	return rebuild.FormatDSN(), nil
}

func ParseMySQLHost(dsn string) (string, error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return "", err
	}
	host, _, err := net.SplitHostPort(cfg.Addr)
	return host, err
}

func ParseMySQLPort(dsn string) (string, error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return "", err
	}
	_, port, err := net.SplitHostPort(cfg.Addr)
	return port, err
}
