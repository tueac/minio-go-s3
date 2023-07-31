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
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称与 object 的名称")
		return
	} else if len(os.Args) < 5 {
		fmt.Println("请输入 object 的名称")
		return
	}

	bucketName, fileName := os.Args[3], os.Args[4]

	object, err := initsystem.InitClient().GetObject(context.Background(), bucketName, fileName, minio.GetObjectOptions{})
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
