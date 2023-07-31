package bucket

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"minio/checkerror"
	"minio/initsystem"
	"os"
)

func CreateBucket() {
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称")
		return
	}

	bucketName := os.Args[3]

	err := initsystem.InitClient().MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{
		Region:        "",
		ObjectLocking: false,
	})
	checkerror.CheckCodeError(err)

	fmt.Printf("%s 创建成功\n", bucketName)
}
