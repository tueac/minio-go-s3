package cmd

import (
	"fmt"
	"minio/bucket"
	"minio/object"
	"os"
)

func Option() {
	args := os.Args

	if len(args) < 10 && args[1] == "bucket" {
		if len(os.Args) < 3 {
			fmt.Println("请输入 bucket 相关命令，参照 -h 命令")
			return
		}

		if args[2] == "create" {
			bucket.CreateBucket()
			return
		}

		if args[2] == "delete" {
			bucket.DeleteBucket()
			return
		}

		if args[2] == "exist" {
			bucket.ExistBucket()
			return
		}

		if args[2] == "list" {
			bucket.ListBucket()
			return
		}
	}

	if len(args) < 10 && args[1] == "object" {
		if len(os.Args) < 3 {
			fmt.Println("请输入 object 相关命令，参照 -h 命令")
			return
		}

		if args[2] == "delete" {
			object.DeleteObject()
			return
		}

		if args[2] == "delall" {
			object.DeleteAll()
			return
		}

		// get 下载文件
		if args[2] == "get" {
			object.GetObject()
			return
		}

		if args[2] == "list" {
			object.ListObject()
			return
		}

		if args[2] == "putdir" {
			object.PutDir()
			return
		}

		if args[2] == "put" {
			object.Putobject()
			return
		}

	}
}
