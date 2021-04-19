package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"io/ioutil"
	"log"
)

type AliOSS struct {
	Endpoint         string //外网Endpoint
	EndpointInternal string //内网Endpoint
	AccessKey        string
	Secret           string
	BucketName       string
	UseInternal      bool //是否使用内网Endpoint
	Bucket           *oss.Bucket
	Client           *oss.Client
}

// useInternal 是否使用内网Endpoint
func NewAliOSSClient(endpoint, endpointInternal, accessKey, secret, bucketName string, useInternal bool) (*AliOSS, error) {
	var (
		client *oss.Client
		err    error
	)

	if useInternal {
		client, err = oss.New(endpointInternal, accessKey, secret)
	} else {
		client, err = oss.New(endpoint, accessKey, secret)
	}
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	aliOSS := &AliOSS{
		Endpoint:         endpoint,
		EndpointInternal: endpointInternal,
		BucketName:       bucketName,
		AccessKey:        accessKey,
		Secret:           secret,
		UseInternal:      useInternal,
		Bucket:           bucket,
		Client:           client,
	}
	return aliOSS, err
}

// PutObject 上传文件流
//
// ossObject 对象名
// string 上传成功后带有http或https的访问路径
func (this *AliOSS) PutObject(ossObject string, reader io.Reader) (string, error) {

	err := this.Bucket.PutObject(ossObject, reader)
	log.Println(err)
	if err != nil {
		return "", err
	}
	// 公共读object 返回的url
	var url string
	if this.UseInternal {
		url = "https://" + this.BucketName + "." + this.EndpointInternal + "/" + ossObject
	} else {
		url = "https://" + this.BucketName + "." + this.Endpoint + "/" + ossObject
	}
	return url, nil
}

// PutObjectFromFile 上传本地文件
//
// string 上传成功后带有http或https的访问路径
func (this *AliOSS) PutObjectFromFile(ossObject string, localFile string) (string, error) {
	err := this.Bucket.PutObjectFromFile(ossObject, localFile)

	if err != nil {
		return "", err
	}
	return "https://" + this.BucketName + "." + this.Endpoint + "/" + ossObject, nil
}

//GetObject：下载文件到流
func (this *AliOSS) GetObject(ossObject string) ([]byte, error) {

	// 下载文件到流。
	body, err := this.Bucket.GetObject(ossObject)
	if err != nil {
		return nil, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetObjectToFile:下载并将文件保存到本地文件系统
func (this *AliOSS) GetObjectToFile(ossObject string, localFile string) error {
	// 下载文件到本地文件。
	err := this.Bucket.GetObjectToFile(ossObject, localFile)
	if err != nil {
		return err
	}
	return nil
}

// RemoveObject: 删除一个对象
func (this *AliOSS) RemoveObject(ossObject string) error {
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err := this.Bucket.DeleteObject(ossObject)
	if err != nil {
		return err
	}
	return nil
}

// CopyObject 在同一存储空间中拷贝文件
//
// destObjectName 复制后新的对象名
func (this *AliOSS) CopyObject(objectName, destObjectName string) error {
	_, err := this.Bucket.CopyObject(objectName, destObjectName)
	if err != nil {
		return err
	}
	return nil
}
