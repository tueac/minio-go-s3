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
	bucketname := os.Args[3]
	objectname := os.Args[4]

	file, err := os.Open(objectname)
	checkerror.CheckCodeError(err)

	defer file.Close()

	filestat, err := file.Stat()
	checkerror.CheckCodeError(err)

	uploadInfo, err := initsystem.InitClient().PutObject(context.Background(), bucketname, objectname, file, filestat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	checkerror.CheckCodeError(err)

	fmt.Printf("%s 文件成功上传到 %s 存储桶\n", objectname, bucketname)
	fmt.Println(uploadInfo.Key, uploadInfo.Size, uploadInfo.ETag)
}
