package bucket

import (
	"context"
	"fmt"
	"minio/checkerror"
	"minio/initsystem"
	"os"
)

func DeleteBucket() {
	bucketname := os.Args[3]

	err := initsystem.InitClient().RemoveBucket(context.Background(), bucketname)
	checkerror.CheckCodeError(err)

	fmt.Printf("%s 删除成功\n", bucketname)
}
