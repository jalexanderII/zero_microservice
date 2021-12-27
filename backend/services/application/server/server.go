package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	"github.com/jalexanderII/zero_microservice/backend/services/application/database/genDB"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
)

type applicationServer struct {
	applicationPB.UnimplementedApplicationServer
	DB                *applicationDB.ApplicationDB
	FileServiceClient fileServicePB.FileServiceClient
	l                 hclog.Logger
}

func NewApplicationServer(db *applicationDB.ApplicationDB, f fileServicePB.FileServiceClient, l hclog.Logger) *applicationServer {
	return &applicationServer{DB: db, FileServiceClient: f, l: l}
}

func (s applicationServer) Apply(ctx context.Context, in *applicationPB.ApplicationRequest) (*applicationPB.ApplicationResponse, error) {
	s.l.Debug("Applying")

	application, err := s.CreateApplication(ctx, &applicationPB.CreateApplicationRequest{Application: in})
	if err != nil {
		s.l.Error("Error applying", "error", err)
		return nil, err
	}
	ar := &applicationPB.ApplicationResponse{
		ReferenceId:    &applicationPB.UUID{Value: uuid.NewString()},
		Status:         string(genDB.ApplicationStatusPENDING),
		ApplicationRef: application.Id,
	}

	response, err := s.CreateResponse(ctx, &applicationPB.CreateApplicationResponse{Application: ar})
	if err != nil {
		s.l.Error("Error submitting application", "error", err)
		return nil, err
	}

	return response, nil
}

func (s applicationServer) Upload(ctx context.Context, in *applicationPB.FileUploadRequest) (*applicationPB.FileUploadResponse, error) {
	s.l.Debug("Upload")
	md := &fileServicePB.MetaData{
		Name:          in.GetMetadata().GetName(),
		SourceId:      in.GetMetadata().GetSourceId(),
		ContentType:   fileServicePB.ContentType(in.GetMetadata().GetContentType()),
		ContentSource: in.GetMetadata().GetName(),
	}

	upload, err := s.FileServiceClient.Upload(ctx, &fileServicePB.FileUploadRequest{Metadata: md, FilePath: in.GetFilePath()})
	if err != nil {
		s.l.Error("[FileServiceClient] Error uploading content", "error", err)
		return nil, err
	}

	application, err := s.GetApplication(ctx, &applicationPB.GetApplicationRequest{Id: md.SourceId})
	if err != nil {
		return nil, err
	}

	err = s.DB.UpdateAttachments(ctx, genDB.UpdateAttachmentsParams{
		ApplicationRequestID: md.SourceId,
		Attachments:          append(application.Attachments, md.Name),
	})
	if err != nil {
		s.l.Error("[DB] Error updating application", "error", err)
		return nil, err
	}

	return &applicationPB.FileUploadResponse{Name: upload.GetName(), Status: applicationPB.STATUS(upload.Status)}, nil
}
