package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

type Conf struct {
	Addr    string `yaml:"address"`
	Command string `yaml:"command"`
}

func LoadConfig(path string) (*Conf, error) {
	c, err := parseConfig(path)
	if err != nil {
		return nil, err
	}

	c, err = validateConfig(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func parseConfig(path string) (*Conf, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Conf{}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return c, nil
}

func validateConfig(c *Conf) (*Conf, error) {

	if c.Addr == "" {
		return nil, errors.New("address config is required")
	}

	return c, nil

}
