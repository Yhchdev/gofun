package oss

import (
	"io"
)

type OSSInterface interface {
	//上传文件流
	PutObject(ossObject string, reader io.Reader) (string, error)
	//上传本地文件
	PutObjectFromFile(ossObject string, localFile string) (string, error)
	//下载对象到流
	GetObject(ossObject string) ([]byte, error)
	//下载并保存对象到本地指定路径
	GetObjectToFile(ossObject string, localFile string) error
	//拷贝对象
	CopyObject(objectName, destObjectName string) error
	//拷贝文件夹下的所有对象
	CopyList(sourceObject, destObject string) error
	//删除单个对象
	RemoveObject(ossObject string) error

	//ListObjects(bucketName, prefix string, recursive bool, doneCh chan struct{}) <-chan ObjectInfo
	////从url格式
	//GetObjectFromUrl(url string) ([]byte, error)
	//GetObjectToFileFromUrl(url string, localFile string) error
}
