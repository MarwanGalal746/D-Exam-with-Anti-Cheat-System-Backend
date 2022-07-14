package main

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/handlers"
)

func main() {
	handlers.Start()
}

//func GetAccountInfo() (string, string, string, string) {
//	azrKey := "udFvBYTwhL1l27aUGOMtb8ojS+7k9IsIXZsjM4YMXCX9Z1dZ2dhUxh+vnn7f03ZTxjwyefEvCWlT+ASt8MPTYw=="
//	azrBlobAccountName := "dexamstorage1"
//	//azrPrimaryBlobServiceEndpoint := fmt.Sprintf("https://%s.blob.core.windows.net/", azrBlobAccountName)
//	azrPrimaryBlobServiceEndpoint := "https://dexamstorage1.blob.core.windows.net/"
//	azrBlobContainer := "dexamstorage1"
//
//	return azrKey, azrBlobAccountName, azrPrimaryBlobServiceEndpoint, azrBlobContainer
//}
//
//func GetBlobName() string {
//	t := time.Now()
//	uuid, _ := uuid.NewV4()
//
//	return fmt.Sprintf("%s-%v.jpg", t.Format("20060102"), uuid)
//}
//
//// The below method assumes you already have the byte array ready to go
//func UploadBytesToBlob(b []byte) (string, error) {
//	azrKey, accountName, endPoint, container := GetAccountInfo()           // This is our account info method
//	u, _ := url.Parse(fmt.Sprint(endPoint, container, "/", GetBlobName())) // This uses our Blob Name Generator to create individual blob urls
//	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey) // Finally we create the credentials object required by the uploader
//	if errC != nil {
//		fmt.Println("1 ", errC.Error())
//		return "", errC
//	}
//
//	// Another Azure Specific object, which combines our generated URL and credentials
//	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))
//
//	ctx := context.Background() // We create an empty context (https://golang.org/pkg/context/#Background)
//
//	// Provide any needed options to UploadToBlockBlobOptions (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadToBlockBlobOptions)
//	o := azblob.UploadToBlockBlobOptions{
//		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
//			ContentType: "image/jpg", //  Add any needed headers here
//		},
//	}
//
//	// Combine all the pieces and perform the upload using UploadBufferToBlockBlob (https://godoc.org/github.com/Azure/azure-storage-blob-go/azblob#UploadBufferToBlockBlob)
//	_, errU := azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
//	if errU != nil {
//		fmt.Println("2 ", errU.Error())
//	}
//	return blockBlobUrl.String(), errU
//}
