package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
web:
  enable: true
  ip: 127.0.0.1
  port: 8080
  mode: debug
  readWriteTimeout: 60s
rpc:
  enable: true
  ip: 127.0.0.1
  port: 8088
log:
  level: debug
mysql:
  common:
    host: 127.0.0.2
    port: 3307
    username: aaaaa
    password: 123456
    database: test
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type StandardConfig struct {
	MySQL     map[string]MySQLConfig `json:"mysql" yaml:"mysql"`
}

type MySQLConnectionConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     uint16 `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

type MySQLConfig struct {
	MySQLConnectionConfig `json:",inline" yaml:",inline"`
	Database              string `json:"database" yaml:"database"`
}

func main() {
	t := StandardConfig{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%#v\n\n", t)

	// d, err := yaml.Marshal(&t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- t dump:\n%s\n\n", string(d))
	//
	// m := make(map[interface{}]interface{})
	//
	// err = yaml.Unmarshal([]byte(data), &m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m:\n%v\n\n", m)
	//
	// d, err = yaml.Marshal(&m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
