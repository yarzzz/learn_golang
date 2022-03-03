package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/download", download)
	log.Println("listen")
	err := http.ListenAndServe("0.0.0.0:8093", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func download(rw http.ResponseWriter, r *http.Request) {
	log.Println("start download")
	rw.Header().Set("Pragma", "public")
	rw.Header().Set("Expires", "0")
	rw.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")

	rw.Header().Set("Content-Type", "application/force-download")
	rw.Header().Add("Content-Type", "application/octet-stream")

	rw.Header().Add("Content-Type", "application/download")
	rw.Header().Add("content-disposition", "attachment;filename="+"f.zip")
	rw.Header().Add("Content-Transfer-Encoding", "binary")
	f, err := os.Open("/home/t/git/release-lotus/lotus-1.10.0.zip")
	if err != nil {
		log.Fatal(err)
	}
	wn, err := io.Copy(rw, f)
	if err != nil {
		return
	}
	log.Println("finish download", wn)
}
