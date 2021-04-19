package oss

import (
	minio "github.com/minio/minio-go"
	"io"
	"io/ioutil"
	"path"
)

type MinioOSS struct {
	Endpoint         string //外网Endpoint
	EndpointInternal string //内网Endpoint
	BucketName       string
	AccessKey        string
	Secret           string
	Region           string
	UseSSL           bool
	UseInternal      bool //使用内网Endpoint
	Client           *minio.Client
}

// useInternal 是否使用内网Endpoint
func NewMinioOSSClient(endpoint, endpointInternal, bucketName, accessKey, secret string, ssl, useInternal bool) (*MinioOSS, error) {
	var (
		client *minio.Client
		err    error
	)
	if useInternal {
		client, err = minio.New(endpointInternal, accessKey, secret, ssl)
	} else {
		client, err = minio.New(endpoint, accessKey, secret, ssl)
	}
	if err != nil {
		return nil, err
	}

	minioOSS := &MinioOSS{
		Endpoint:         endpoint,
		EndpointInternal: endpointInternal,
		BucketName:       bucketName,
		AccessKey:        accessKey,
		Secret:           secret,
		UseSSL:           ssl,
		UseInternal:      useInternal,
		Client:           client,
	}

	return minioOSS, nil
}

// PutObject 上传文件流
//
// ossObject 对象名
// string 上传成功后带有http或https的访问路径
func (this *MinioOSS) PutObject(ossObject string, reader io.Reader) (string, error) {

	fileSuffix := path.Ext(ossObject)
	if "" == fileSuffix {
		fileSuffix = ".noSuffix"
		ossObject += fileSuffix
	}
	contentType := "application/" + string([]byte(fileSuffix)[1:])

	_, err := this.Client.PutObject(this.BucketName, ossObject, reader, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}
	var url string
	if this.UseSSL {
		url = "https://"
	} else {
		url = "http://"
	}
	if this.UseInternal {
		url += this.EndpointInternal + "/"
	} else {
		url += this.Endpoint + "/"
	}
	url += this.BucketName + "/" + ossObject

	return url, nil
}

// PutObjectFromFile上传本地文件
//
// string 上传成功后带有http或https的访问路径
func (this *MinioOSS) PutObjectFromFile(ossObject string, localFile string) (string, error) {

	fileSuffix := path.Ext(localFile)
	if "" == fileSuffix {
		fileSuffix = ".noSuffix"
		ossObject += fileSuffix
	}
	contentType := "application/" + string([]byte(fileSuffix)[1:])

	_, err := this.Client.FPutObject(this.BucketName, ossObject, localFile, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	var url string
	if this.UseSSL {
		if this.UseInternal {
			url = "https://" + this.EndpointInternal + "/" + this.BucketName + "/" + ossObject
		} else {
			url = "https://" + this.Endpoint + "/" + this.BucketName + "/" + ossObject
		}
	} else {
		if this.UseInternal {
			url = "http://" + this.EndpointInternal + "/" + this.BucketName + "/" + ossObject
		} else {
			url = "http://" + this.Endpoint + "/" + this.BucketName + "/" + ossObject
		}
	}

	return url, nil

}

// GetObject:下载文件到流
func (this *MinioOSS) GetObject(ossObject string) ([]byte, error) {

	object, err := this.Client.GetObject(this.BucketName, ossObject, minio.GetObjectOptions{})

	data, err := ioutil.ReadAll(object)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetObjectToFile:下载并将文件保存到本地文件系统
func (this *MinioOSS) GetObjectToFile(ossObject string, localFile string) error {
	err := this.Client.FGetObject(this.BucketName, ossObject, localFile, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

//RemoveObject：删除单个对象
func (this *MinioOSS) RemoveObject(ossObject string) error {
	err := this.Client.RemoveObject(this.BucketName, ossObject)
	if err != nil {
		return err
	}
	return nil
}

//CopyObject:在同一存储空间中拷贝文件
//
// destObjectName 复制后新的对象名
func (this *MinioOSS) CopyObject(objectName, destObjectName string) error {

	src := minio.NewSourceInfo(this.BucketName, objectName, nil)
	dst, err := minio.NewDestinationInfo(this.BucketName, destObjectName, nil, nil)
	if err != nil {
		return err
	}

	err = this.Client.CopyObject(dst, src)
	if err != nil {
		return err
	}
	return nil
}
