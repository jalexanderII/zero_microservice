package server

import (
	"bytes"
	"context"
	"io/ioutil"
	"path"
	"strconv"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type fileServiceServer struct {
	fileServicePB.UnimplementedFileServiceServer
	DB *mongo.Client
	l  hclog.Logger
}

func NewFileServiceServer(db *mongo.Client, l hclog.Logger) *fileServiceServer {
	return &fileServiceServer{DB: db, l: l}
}

func (s fileServiceServer) Upload(ctx context.Context, in *fileServicePB.FileUploadRequest) (*fileServicePB.FileUploadResponse, error) {
	s.l.Debug("UploadContent")
	filename := path.Base(in.GetFilePath())

	data, err := ioutil.ReadFile(in.GetFilePath())
	if err != nil {
		s.l.Error("[Error] cannot real file", "error", err)
		return nil, err
	}

	bucket, err := gridfs.NewBucket(s.DB.Database(config.CONTENTDBNAME))
	if err != nil {
		s.l.Error("[Error] cannot get mongo bucket", "error", err)
		return nil, err
	}
	opts := options.GridFSUpload()
	opts.SetMetadata(in.GetMetadata())
	uploadStream, err := bucket.OpenUploadStream(filename, opts)
	if err != nil {
		s.l.Error("[Error] opening upload stream", "error", err)
		return nil, err
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		s.l.Error("[Error] writing to upload stream", "error", err)
		return nil, err
	}
	s.l.Info("Write file to DB was successful.", "File size", fileSize)
	return &fileServicePB.FileUploadResponse{Name: filename, Status: fileServicePB.STATUS_SUCCESS}, nil
}

func (s fileServiceServer) Download(ctx context.Context, in *fileServicePB.FileDownloadRequest) (*fileServicePB.FileDownloadResponse, error) {
	s.l.Debug("DownloadContent")
	db := s.DB.Database(config.CONTENTDBNAME)
	fsFiles := db.Collection("fs.files")

	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{"filename": in.GetFileName()}).Decode(&results)
	if err != nil {
		s.l.Error("[Error] finding db Collection", "error", err)
		return nil, err
	}
	md := MongoMetaDataToPB(results)
	s.l.Info("Fetched metaData from file:", "MetaData", md)

	bucket, _ := gridfs.NewBucket(db)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(in.GetFileName(), &buf)
	if err != nil {
		s.l.Error("[Error] getting data from download stream", "error", err)
		return nil, err
	}

	s.l.Info("File size to download:", "File size", dStream)
	err = ioutil.WriteFile(config.DOWNLOADPATH+in.GetFileName(), buf.Bytes(), 0600)
	if err != nil {
		s.l.Error("[Error] saving downloaded file to local storage", "error", err)
		return nil, err
	}
	return &fileServicePB.FileDownloadResponse{Metadata: md, FileSize: strconv.FormatInt(dStream, 10)}, nil
}

func MongoMetaDataToPB(results bson.M) *fileServicePB.MetaData {
	md := new(fileServicePB.MetaData)
	for _, record := range results {
		if rec, ok := record.(primitive.M); ok {
			for key, val := range rec {
				switch key {
				case "name":
					if str, ok := val.(string); ok {
						md.Name = str
					}
				case "sourceid":
					if i, ok := val.(int32); ok {
						md.SourceId = i
					}
				case "contenttype":
					if i, ok := val.(int32); ok {
						md.ContentType = fileServicePB.ContentType(i)
					}
				case "contentsource":
					if str, ok := val.(string); ok {
						md.ContentSource = str
					}
				}
			}
		}
	}
	return md
}
