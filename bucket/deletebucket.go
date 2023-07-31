package bucket

import (
	"context"
	"fmt"
	"minio/checkerror"
	"minio/initsystem"
	"os"
)

func DeleteBucket() {
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称")
		return
	}

	bucketName := os.Args[3]

	err := initsystem.InitClient().RemoveBucket(context.Background(), bucketName)
	checkerror.CheckCodeError(err)

	fmt.Printf("%s 删除成功\n", bucketName)
}
