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
	bucketname := os.Args[3]

	err := initsystem.InitClient().MakeBucket(context.Background(), bucketname, minio.MakeBucketOptions{
		Region:        "",
		ObjectLocking: false,
	})
	checkerror.CheckCodeError(err)

	fmt.Printf("%s 创建成功\n", bucketname)
}
