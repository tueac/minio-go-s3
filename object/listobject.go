package object

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"minio/initsystem"
	"os"
)

func ListObject() {
	bucketname := os.Args[3]
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := initsystem.InitClient().ListObjects(ctx, bucketname, minio.ListObjectsOptions{
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
