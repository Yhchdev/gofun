package oss

import (
	"gofun/conf"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func NewMinio() *MinioOSS {
	minioOSS, err := NewMinioOSSClient(config.Endpoint, config.EndpointInternal, config.Bucket, config.Accesskey, config.Secret, config.UseSSL, config.UseInternal)

	if err != nil {
		log.Fatal(err)
	}
	return minioOSS
}

func TestMinioOss_PutObject(t *testing.T) {
	minioOSS := NewMinio()

	f, _ := os.Open("../testfile/thirdback_service.sh")

	str, err := minioOSS.PutObject("1726.sh", f)
	log.Println(str)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestMinioOss_PutObjectFromFile(t *testing.T) {
	minioOSS := NewMinio()

	str, err := minioOSS.PutObjectFromFile("1727.png", "../testfile/611137.png")
	log.Println(str)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestMinioOss_GetObject(t *testing.T) {

	minioOSS := NewMinio()
	content, err := minioOSS.GetObject("1128.png")
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile("../testfile/Getobject.png", content, 0644); err != nil {
		log.Fatal(err)
	}
}

func TestMinioOSS_GetObjectToFile(t *testing.T) {
	minioOSS := NewMinio()
	if err := minioOSS.GetObjectToFile("1128.png", "../testfile/1143.png"); err != nil {
		t.Fatal(err)
	}
}

func TestMinioOSS_RemoveObject(t *testing.T) {
	minioOSS := NewMinio()
	if err := minioOSS.RemoveObject("test415"); err != nil {
		t.Fatal(err)
	}
}

func TestMinioOSS_CopyObject(t *testing.T) {
	minioOSS := NewMinio()
	if err := minioOSS.CopyObject("1114.png", "1114_backup.png"); err != nil {
		t.Fatal(err)
	}
}
