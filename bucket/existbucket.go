package bucket

import (
	"context"
	"fmt"
	"minio/checkerror"
	"minio/initsystem"
	"os"
)

func ExistBucket() {
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称")
		return
	}

	bucketName := os.Args[3]

	fondBucket, err := initsystem.InitClient().BucketExists(context.Background(), bucketName)
	checkerror.CheckCodeError(err)

	if fondBucket {
		fmt.Printf("%s 是存在的\n", bucketName)
	} else {
		fmt.Printf("%s 是不存在的\n", bucketName)
	}
}
