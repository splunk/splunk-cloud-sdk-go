/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package config

import (
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Service struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Scheme string `yaml:"scheme"`
}

type IdpService struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Scheme string `yaml:"scheme"`
	Server string `yaml:"server"`
}

type Environment struct {
	APIService Service    `yaml:"api-service"`
	AppService Service    `yaml:"app-service"`
	IdpService IdpService `yaml:"idp-service"`
	Profile    string     `yaml:"profile"`
}

type Cfg struct {
	Profiles     map[string]map[string]string `yaml:"profiles"`
	Environments map[string]*Environment      `yaml:"environments"`
}

var config Cfg

func Load(reader io.Reader) error {
	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(&config); err != nil {
		return err
	}
	return nil
}

// Load the named config file.
func LoadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	return Load(file)
}

func Environments() map[string]*Environment {
	return config.Environments
}

func GetEnvironment(name string) (*Environment, error) {
	env, ok := config.Environments[name]
	if !ok {
		return nil, fmt.Errorf("not found: '%s'", name)
	}
	return env, nil
}

// Returns the named application profile.
func GetProfile(name string) (map[string]string, error) {
	profile, ok := config.Profiles[name]
	if !ok {
		return nil, fmt.Errorf("not found: '%s'", name)
	}
	_, ok = profile["kind"] // ensure 'kind' exists
	if !ok {
		return nil, fmt.Errorf("missing kind")
	}
	return profile, nil
}

func Profiles() map[string]map[string]string {
	return config.Profiles
}
