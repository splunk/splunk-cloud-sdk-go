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

package fcache

// A simple JSON file cache.

import (
	"log"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
)

// Answers if the named file exists.
func fileExists(filename string) bool {
	fileinfo, e := os.Stat(filename)
	if e == nil {
		return !fileinfo.IsDir()
	}
	if os.IsNotExist(e) {
		return false
	}
	return true
}

// Opens the named file for writing.
func openWrite(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
}

type Cache struct {
	filename string
	data     *toml.Tree
}

func newCache(filename string) *Cache {
	data := make(map[string]interface{})
	tree, err := toml.TreeFromMap(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Cache{
		filename,
		tree,
	}
}

func (cache *Cache) All() map[string]interface{} {
	data := make(map[string]interface{})
	keys := cache.data.Keys()
	for _, key := range keys {
		data[key] = cache.data.Get(key)
	}
	return data
}

func (cache *Cache) Clear() {
	data := make(map[string]interface{})
	tree, err := toml.TreeFromMap(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	cache.data = tree
	err = cache.Save()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (cache *Cache) Delete(key string) error {
	err := cache.data.Delete(key)
	if err != nil {
		return err
	}
	return cache.Save()
}

func (cache *Cache) Get(key string) interface{} {
	value := cache.data.Get(key)
	return value
}

func (cache *Cache) GetString(key string) (string, bool) {
	value := cache.Get(key)
	if value != nil {
		return cache.Get(key).(string), true
	}
	return "", false
}

func (cache *Cache) Load() error {
	if !fileExists(cache.filename) {
		return nil
	}
	tree, err := toml.LoadFile(cache.filename)
	if err != nil {
		return err
	}
	cache.data = tree
	return nil
}

func (cache *Cache) Save() error {
	file, err := openWrite(cache.filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = cache.data.WriteTo(file)
	if err != nil {
		return err
	}
	return nil
}

func (cache *Cache) Set(key string, value interface{}) {
	if !isValid(key) {
		return
	}
	switch val := value.(type) {
	case map[string]interface{}:
		tree, err := toml.TreeFromMap(val)
		if err != nil {
			log.Fatal(err.Error())
		}
		cache.data.Set(key, tree)
	default:
		cache.data.Set(key, value)
	}
	err := cache.Save()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func isValid(key string) bool {
	if strings.ToLower(key) == "password" || strings.ToLower(key) == "pass" {
		log.Fatal("Storing passords is not allowed.")
		return false
	}
	if strings.ToLower(key) == "user" {
		log.Fatal("Please try `username` as your key.")
		return false
	}
	return true
}

func Load(filename string) (*Cache, error) {
	result := newCache(filename)
	if err := result.Load(); err != nil {
		return nil, err
	}
	return result, nil
}
