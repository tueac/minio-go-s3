package initsystem

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"minio/checkerror"
	"os"
)

func InitClient() *minio.Client {
	s3Client, err := minio.New(EndpointAddr, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyID, SecretAccessKey, ""),
		Secure: UseSSL,
	})
	checkerror.CheckCodeError(err)

	return s3Client
}

func LenInput(inputlen int, hint string) bool {
	if len(os.Args) < inputlen {
		fmt.Println(hint)
		return false
	}
	return true
}
