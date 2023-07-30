package object

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io/fs"
	"log"
	"minio/initsystem"
	"os"
	"path/filepath"
)

func PutDir() {
	bucketname := os.Args[3]
	dirname := os.Args[4]

	// 遍历文件夹并逐个上传文件
	walkerr := filepath.Walk(dirname, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过文件夹本身
		if info.IsDir() {
			return nil
		}

		// 上传文件到 MinIO 存储桶
		_, err = initsystem.InitClient().FPutObject(context.Background(), bucketname, path, path, minio.PutObjectOptions{})
		if err != nil {
			log.Println("上传文件失败", err)
		} else {
			fmt.Println("上传文件成功", path)
		}

		return nil
	})

	if walkerr != nil {
		log.Fatalf("遍历文件夹失败:", walkerr)
	}
}
