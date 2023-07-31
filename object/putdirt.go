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
	if len(os.Args) < 4 {
		fmt.Println("请输入 bucket 的名称与本地文件夹目录路径")
		return
	} else if len(os.Args) < 5 {
		fmt.Println("本地文件夹目录路径")
		return
	}

	bucketName := os.Args[3]
	dirName := os.Args[4]

	// 遍历文件夹并逐个上传文件
	walkerr := filepath.Walk(dirName, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过文件夹本身
		if info.IsDir() {
			return nil
		}

		// 上传文件到 MinIO 存储桶
		_, err = initsystem.InitClient().FPutObject(context.Background(), bucketName, path, path, minio.PutObjectOptions{})
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
