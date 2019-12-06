package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 各モジュールごとのLicenseを表すstruct
type license struct {
	Name     string
	URL      string
	FilePath string
	Content  string
}

// モジュールごとのライセンスを全て内包するstruct
type licenseJson struct {
	Licenses []license
}

// 与えられたパス名からライセンスを読み取り，ライセンスの一覧を返す
func getLicenses(path string) ([]license, error) {
	contents, err := ioutil.ReadFile("./credits.json")
	if err != nil {
		return nil, fmt.Errorf("%s does not exists.\n", path)
	}
	var licenses licenseJson
	if err = json.Unmarshal(contents, &licenses); err != nil {
		return nil, fmt.Errorf("Failed to parse %s.\n", path)
	}
	return licenses.Licenses, nil
}
