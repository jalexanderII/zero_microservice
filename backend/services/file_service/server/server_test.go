package server

import (
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	fileServiceDB "github.com/jalexanderII/zero_microservice/backend/services/file_service/database"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
)

var L = hclog.Default()

func Test_fileServiceServer_Upload(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db := fileServiceDB.InitiateMongoClient()
	server := fileServiceServer{DB: db, l: L}
	in := &fileServicePB.FileUploadRequest{
		Metadata: &fileServicePB.MetaData{
			Name:          "icon.png",
			SourceId:      1,
			ContentType:   fileServicePB.ContentType_IMAGE,
			ContentSource: "APARTMENT",
		},
		FilePath: "/Users/joel/Downloads/icon.png",
	}

	upload, err := server.Upload(ctx, in)
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if upload.Status != fileServicePB.STATUS_SUCCESS {
		t.Errorf("2: Failed to upload new content: %+v", upload)
	}
}

func Test_fileServiceServer_Download(t *testing.T) {
	ctx, cancel := config.NewDBContext(5 * time.Second)
	defer cancel()

	db := fileServiceDB.InitiateMongoClient()
	server := fileServiceServer{DB: db, l: L}
	in := &fileServicePB.FileDownloadRequest{
		FileName: "icon.png",
	}

	download, err := server.Download(ctx, in)
	if err != nil {
		t.Errorf("1: An error was returned: %v", err)
	}
	if download.GetMetadata().Name != "icon.png" {
		t.Errorf("2: Failed to download content from DB: %+v", download)
	}
}
