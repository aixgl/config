package config

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type confMap map[string]map[string]string

var mem = make(map[string]*Config)

type Config struct {
	Conf confMap
	File string
}

func C(name ...string) *Config {
	key := ""
	file := ""
	switch len(name) {
	case 0:
		key = "G"
	case 1:
		key = name[0]
	case 2:
		key = name[0]
		file = name[1]
	}

	if mem[key] == nil {
		mem[key] = &Config{Conf: make(confMap), File: ""}
		if file != "" {
			mem[key].Set(file)
		}
	}
	return mem[key]
}

func (conf *Config) Set(file ...string) *Config {
	var confFile = conf.File
	if len(file) >= 1 {
		_, err := os.Stat(file[0])
		if err != nil {
			return conf
		}
		conf.File = file[0]
		confFile = file[0]
	}

	conf.Escape(confFile)
	return conf
}

func (conf *Config) GetAll() confMap {
	return conf.Conf
}

func (conf *Config) Get(fld ...string) interface{} {

	switch len(fld) {
	case 0:
		return conf.Conf
	case 1:
		sf := strings.SplitN(fld[0], ".", 2)
		if len(sf) == 1 {
			return conf.Conf[fld[0]]
		} else if len(sf) == 2 {
			if conf.Conf[sf[0]] == nil {
				return nil
			}
			return conf.Conf[sf[0]][sf[1]]
		}
	case 2:
		if conf.Conf[fld[0]] == nil {
			return nil
		}
		return conf.Conf[fld[0]][fld[1]]
	}
	return conf.Conf
}

func (conf *Config) Escape(file string) {
	fi, err := os.Open(file)
	if err != nil {
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	CurMapModule := "G"
	conf.Conf = newConfMap()
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		//Ajust first character
		b := strings.Trim(string(a), " ")
		if len(b) == 0 {
			continue
		}

		fa := b[:1]
		if fa == ";" || fa == "#" || b[:2] == "//" {
			continue
		}

		sf := strings.SplitN(b, "=", 2)
		if len(sf) == 1 && b[:1] == "[" {
			b = strings.TrimRight(b, "]")
			b = strings.TrimLeft(b, "[")
			CurMapModule = strings.Trim(b, " ")
			continue
		}

		if len(sf) == 2 {
			if conf.Conf[CurMapModule] == nil {
				conf.Conf[CurMapModule] = make(map[string]string)
			}
			sf[0] = strings.Trim(sf[0], " ")
			sf[1] = strings.Trim(sf[1], " ")
			conf.Conf[CurMapModule][sf[0]] = sf[1]
			continue
		}
	}
}

func newConfMap() confMap {
	var ret confMap = make(map[string]map[string]string)
	return ret
}
