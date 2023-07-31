package object

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"minio/checkerror"
	"minio/initsystem"
	"os"
)

func Putobject() {
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称与本地文件路径")
		return
	} else if len(os.Args) < 5 {
		fmt.Println("本地文件路径")
		return
	}

	bucketName := os.Args[3]
	objectName := os.Args[4]

	file, err := os.Open(objectName)
	checkerror.CheckCodeError(err)

	defer file.Close()

	filestat, err := file.Stat()
	checkerror.CheckCodeError(err)

	uploadInfo, err := initsystem.InitClient().PutObject(context.Background(), bucketName, objectName, file, filestat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	checkerror.CheckCodeError(err)

	fmt.Printf("%s 文件成功上传到 %s 存储桶\n", objectName, bucketName)
	fmt.Println(uploadInfo.Key, uploadInfo.Size, uploadInfo.ETag)
}
