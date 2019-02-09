package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func main() {
	versionCode := os.Args[1]
	fmt.Println("VERSION " + versionCode)
	v, _ := strconv.ParseInt(versionCode, 10, 64)
	v64 := float64(v)
	f := float64(1000000)
	minor := math.Floor(v64 / f)
	minorString := strconv.Itoa(int(minor))

	y := v64 - (minor * 1000000)
	patch := math.Floor(y / 1000)
	patchString := strconv.Itoa(int(patch))

	versionString := "1." + minorString + "." + patchString
	fmt.Println("1." + minorString + "." + patchString)

	input, err := ioutil.ReadFile("pubspec.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte("version: 1.0.0+1"), []byte("version: "+versionString+"+"+versionCode), -1)
	if err := ioutil.WriteFile("pubspec.yaml", output, 0775); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
