package object

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"minio/initsystem"
	"os"
)

func ListObject() {
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称")
		return
	}

	bucketName := os.Args[3]
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := initsystem.InitClient().ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix: "",
		//Recursive: true,
	})

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		var objectsize float64 = float64(object.Size / 1000 / 1000)
		fmt.Printf("%s %.2f MB %s\n", object.Key, objectsize, object.ETag)
	}
}
