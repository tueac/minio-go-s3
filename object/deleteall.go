package object

import (
	"context"
	"fmt"
	"log"
	"minio/initsystem"
	"os"

	"github.com/minio/minio-go/v7"
)

func DeleteAll() {
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称")
		return
	}

	bucketName := os.Args[3]
	fmt.Println(bucketName)

	// 获取存储桶中的所有对象
	objectsCh := initsystem.InitClient().ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})

	// 遍历所有对象并逐个删除
	for object := range objectsCh {
		if object.Err != nil {
			log.Println("遍历对象失败:", object.Err)
			continue
		}

		// 删除对象
		err := initsystem.InitClient().RemoveObject(context.Background(), bucketName, object.Key, minio.RemoveObjectOptions{})
		if err != nil {
			log.Println("删除对象失败:", err)
		} else {
			fmt.Println("删除对象成功:", object.Key)
		}
	}
}
