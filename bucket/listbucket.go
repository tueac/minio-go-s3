package bucket

import (
	"context"
	"fmt"
	"minio/checkerror"
	"minio/initsystem"
)

func ListBucket() {
	buckets, err := initsystem.InitClient().ListBuckets(context.Background())
	checkerror.CheckCodeError(err)

	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}
}
