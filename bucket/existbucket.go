package bucket

import (
	"context"
	"fmt"
	"minio/checkerror"
	"minio/initsystem"
	"os"
)

func ExistBucket() {
	bucketname := os.Args[3]

	fonudBucket, err := initsystem.InitClient().BucketExists(context.Background(), bucketname)
	checkerror.CheckCodeError(err)

	if fonudBucket {
		fmt.Printf("%s 是存在的\n", bucketname)
	} else {
		fmt.Printf("%s 是不存在的\n", bucketname)
	}
}
