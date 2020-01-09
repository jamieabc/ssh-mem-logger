package parser

import "github.com/jamieabc/ssh-mem-logger/pkg/lua"

type Parser interface {
	Parse() (Config, error)
}

type Config struct {
	Servers []Server `lua:"servers"`
}

type Server struct {
	UserName string `lua:"user_name"`
	KeyPath  string `lua:"key_path"`
	IP       string `lua:"ip"`
	Port     int    `lua:"port"`
	Name     string `lua:"name"`
}

type parser struct {
	filePath string
}

func (p *parser) Parse() (Config, error) {
	conf := &Config{}
	err := lua.Parse(p.filePath, conf)
	return *conf, err
}

func NewParser(path string) Parser {
	return &parser{
		filePath: path,
	}
}
