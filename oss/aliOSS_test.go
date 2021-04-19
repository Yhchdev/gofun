package oss

import (
	"gofun/conf"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestAliOss_PutObject(t *testing.T) {

	alioss, err := NewAliOSSClient(config.OSSEndpoint, config.OSSEndpointInternal, config.OSSAccesskeyID, config.OSSAccessKeySecret, config.OSSBucket, config.OSSUseInternal)
	if err != nil {
		log.Fatalln(err)
	}
	fd, _ := os.Open("../testfile/thirdback_service.sh")
	url, err := alioss.PutObject("1538.sh", fd)
	log.Println(url)
	if err != nil {
		log.Println(err)
	}

}

func TestAliOss_PutObjectFromFile(t *testing.T) {
	alioss, err := NewAliOSSClient(config.OSSEndpoint, config.OSSEndpointInternal, config.OSSAccesskeyID, config.OSSAccessKeySecret, config.OSSBucket, config.OSSUseInternal)
	if err != nil {
		log.Fatalln(err)
	}

	url, err := alioss.PutObjectFromFile("1111.png", "../testfile/611137.png")
	log.Println(url)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestAliOSS_GetObject(t *testing.T) {

	alioss, err := NewAliOSSClient(config.OSSEndpoint, config.OSSEndpointInternal, config.OSSAccesskeyID, config.OSSAccessKeySecret, config.OSSBucket, config.OSSUseInternal)
	if err != nil {
		t.Fatal(err)
	}

	content, err := alioss.GetObject("exec20210409.sql")
	if err != nil {
		t.Fatal(err)
	}

	if err = ioutil.WriteFile("../testfile/GetObjFromAli.sql", content, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestAliOSS_GetObjectToFile(t *testing.T) {
	alioss, err := NewAliOSSClient(config.OSSEndpoint, config.OSSEndpointInternal, config.OSSAccesskeyID, config.OSSAccessKeySecret, config.OSSBucket, config.OSSUseInternal)
	if err != nil {
		t.Fatal(err)
	}

	if err = alioss.GetObjectToFile("exec20210409.sql", "../testfile/1145.sql"); err != nil {
		log.Fatal(err)
	}
}

func TestAliOSS_RemoveObject(t *testing.T) {
	alioss, err := NewAliOSSClient(config.OSSEndpoint, config.OSSEndpointInternal, config.OSSAccesskeyID, config.OSSAccessKeySecret, config.OSSBucket, config.OSSUseInternal)
	if err != nil {
		t.Fatal(err)
	}

	if err = alioss.RemoveObject("4191048.sh"); err != nil {
		t.Fatal(err)
	}
}

func TestAliOSS_CopyObject(t *testing.T) {
	alioss, err := NewAliOSSClient(config.OSSEndpoint, config.OSSEndpointInternal, config.OSSAccesskeyID, config.OSSAccessKeySecret, config.OSSBucket, config.OSSUseInternal)
	if err != nil {
		t.Fatal(err)
	}

	if err = alioss.CopyObject("1111.png", "1111_back.png"); err != nil {
		t.Fatal(err)
	}
}
