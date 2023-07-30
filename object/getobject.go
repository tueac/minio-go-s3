package object

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"minio/checkerror"
	"minio/initsystem"
	"os"
)

// get 下载文件
func GetObject() {
	bucketname, filename := os.Args[3], os.Args[4]

	object, err := initsystem.InitClient().GetObject(context.Background(), bucketname, filename, minio.GetObjectOptions{})
	checkerror.CheckCodeError(err)
	defer object.Close()

	localFile, err := os.Create("local-file.jpg")
	checkerror.CheckCodeError(err)
	defer localFile.Close()

	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
}
