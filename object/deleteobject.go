package object

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"minio/checkerror"
	"minio/initsystem"
	"os"
)

func DeleteObject() {
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称与 object 的名称")
		return
	} else if len(os.Args) < 5 {
		fmt.Println("请输入 object 的名称")
		return
	}

	bucketName := os.Args[3]
	objectName := os.Args[4]

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        "",
	}

	err := initsystem.InitClient().RemoveObject(context.Background(), bucketName, objectName, opts)
	checkerror.CheckCodeError(err)

	fmt.Printf("%s 存储桶的中 %s 文件删除成功\n", bucketName, objectName)
}
