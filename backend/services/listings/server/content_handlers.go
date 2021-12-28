package server

import (
	"context"
	"fmt"

	"github.com/jalexanderII/zero_microservice/backend/services/listings/database/genDB"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

func (s listingsServer) Upload(ctx context.Context, in *listingsPB.FileUploadRequest) (*listingsPB.FileUploadResponse, error) {
	s.l.Debug("Upload")
	md := &fileServicePB.MetaData{
		Name:          in.GetMetadata().GetName(),
		SourceId:      in.GetMetadata().GetSourceId(),
		ContentType:   fileServicePB.ContentType(in.GetMetadata().GetContentType()),
		ContentSource: in.GetMetadata().GetContentSource().String(),
	}

	upload, err := s.FileServiceClient.Upload(ctx, &fileServicePB.FileUploadRequest{Metadata: md, FilePath: in.GetFilePath()})
	if err != nil {
		s.l.Error("[FileServiceClient] Error uploading content", "error", err)
		return nil, err
	}

	switch sourceType := in.GetMetadata().GetContentSource(); sourceType {
	case listingsPB.ContentSource_APARTMENT:
		err = s.DB.AppendContentApartment(ctx, genDB.AppendContentApartmentParams{
			ApartmentID: md.GetSourceId(),
			ArrayAppend: md.GetName(),
		})
		if err != nil {
			s.l.Error("[DB] Error updating apartment attachments", "error", err)
			return nil, err
		}
	case listingsPB.ContentSource_BUILDING:
		err = s.DB.AppendContentBuilding(ctx, genDB.AppendContentBuildingParams{
			BuildingID:  md.GetSourceId(),
			ArrayAppend: md.GetName(),
		})
		if err != nil {
			s.l.Error("[DB] Error updating building attachments", "error", err)
			return nil, err
		}
	default:
		s.l.Error("[Error] incorrect source type or other error", "source_type", sourceType, "error", err)
		return nil, fmt.Errorf("incorrect source type or other error")
	}
	return &listingsPB.FileUploadResponse{Name: upload.GetName(), Status: listingsPB.STATUS(upload.Status)}, nil
}
