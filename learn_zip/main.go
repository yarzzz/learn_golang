package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/xerrors"
)

func main() {
	a, err := ioutil.ReadFile("./testc2inf0")
	if err != nil {
		panic(err)
	}
	az, err := Zip(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(az))
	auz, err := Unzip(az)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes.Equal(a, auz))
}

func Zip(origin []byte) ([]byte, error) {
	zf, err := os.CreateTemp("./", "*")
	if err != nil {
		return nil, err
	}
	zw := zip.NewWriter(zf)
	hw, err := zw.CreateHeader(&zip.FileHeader{
		Name:   "p1out",
		Method: zip.Deflate,
	})
	if err != nil {
		return nil, err
	}
	_, err = hw.Write(origin)
	zw.Close()
	zf.Close()
	defer os.Remove(zf.Name())
	return ioutil.ReadFile(zf.Name())
}

func Unzip(compressed []byte) ([]byte, error) {
	inBuf := bytes.NewReader(compressed)
	zr, err := zip.NewReader(inBuf, int64(inBuf.Len()))
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if f.Name == "p1out" {
			cf, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer cf.Close()
			return io.ReadAll(cf)
		}
	}
	return nil, xerrors.New("not found")
}
