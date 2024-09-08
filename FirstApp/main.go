package main

import (
	"fmt"
	//"github.com/twpayne/go-geom"
)

func main() {
	fmt.Println("Hello GO")

	println("Hello GO")

	// geom.NewMultiPolygon()
}

//modül oluşturmak için go mod init komutunu çalıştırıyoruz
//projeye ek bir modül eklemek için go get github.com/twpayne/go-geom gibi bir komut çalıştırmak gerekli
//go kullanılmayan paketleri import içinde tutmaz
//go mod tidy kullanılmayan depencyleri kaldırır henüz yüklenmemiş olanlarıda indirir
