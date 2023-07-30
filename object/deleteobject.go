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
	bucketname := os.Args[3]
	objectname := os.Args[4]

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        "",
	}

	err := initsystem.InitClient().RemoveObject(context.Background(), bucketname, objectname, opts)
	checkerror.CheckCodeError(err)

	fmt.Printf("%s 存储桶的中 %s 文件删除成功\n", bucketname, objectname)
}
