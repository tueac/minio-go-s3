package cmd

import (
	"minio/bucket"
	"minio/object"
	"os"
)

func Option() {
	args := os.Args

	if len(args) < 10 && args[1] == "bucket" {
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

		if args[2] == "put" {
			object.Putobject()
			return
		}

		if args[2] == "putdir" {
			object.PutDir()
			return
		}
	}

}
