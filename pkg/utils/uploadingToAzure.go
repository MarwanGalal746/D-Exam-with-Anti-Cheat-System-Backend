package utils

import (
	"context"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"net/url"
	"os"
)

func GetAccountInfo() (string, string, string, string) {
	azrKey := os.Getenv("AZR_KEY")
	azrBlobAccountName := os.Getenv("AZR_BLOB_ACCOUNT_NAME")
	azrPrimaryBlobServiceEndpoint := os.Getenv("AZR_BLOB_SERVICE_ENDPOINT")
	azrBlobContainer := os.Getenv("AZR_BLOB_CONTAINER")

	return azrKey, azrBlobAccountName, azrPrimaryBlobServiceEndpoint, azrBlobContainer
}

// UploadBytesToBlob The below method assumes you already have the byte array ready to go
func UploadBytesToBlob(b []byte, imgName string) (string, error) {
	azrKey, accountName, endPoint, container := GetAccountInfo()           // This is our account info method
	u, _ := url.Parse(fmt.Sprint(endPoint, container, "/", imgName))       // This uses our Blob Name Generator to create individual blob urls
	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey) // Finally we create the credentials object required by the uploader
	if errC != nil {
		fmt.Println("1 ", errC.Error())
		return "", errC
	}

	// Another Azure Specific object, which combines our generated URL and credentials
	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background() // We create an empty context (https://golang.org/pkg/context/#Background)

	// Provide any needed options to UploadToBlockBlobOptions (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadToBlockBlobOptions)
	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: "image/jpg", //  Add any needed headers here
		},
	}

	// Combine all the pieces and perform the upload using UploadBufferToBlockBlob (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadBufferToBlockBlob)
	_, errU := azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
	if errU != nil {
		fmt.Println("2 ", errU.Error())
	}
	return blockBlobUrl.String(), errU
}
