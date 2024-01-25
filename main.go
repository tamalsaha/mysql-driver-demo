package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"net"
	"net/url"
	"strings"
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

func CanonicalMySQLDSN(dsn string) (string, error) {
	_, err := mysql.ParseDSN(dsn)
	if err == nil {
		return dsn, nil
	}

	u, err := url.Parse(dsn)
	if err != nil {
		return "", err
	}

	rebuild := mysql.NewConfig()
	rebuild.Net = u.Scheme
	rebuild.Addr = u.Host
	rebuild.DBName = strings.TrimPrefix(u.Path, "/")
	if u.User != nil {
		rebuild.User = u.User.Username()
		if pass, found := u.User.Password(); found {
			rebuild.Passwd = pass
		}
	}
	rebuild.Params = map[string]string{}
	for k, v := range u.Query() {
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
