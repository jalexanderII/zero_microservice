package server

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"

	"github.com/jalexanderII/zero_microservice/backend/services/listings/database/genDB"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UploadContent is a streaming RPC to upload content
func (s listingsServer) UploadContent(stream listingsPB.Listings_UploadContentServer) error {
	s.l.Debug("UploadContent")
	req, err := stream.Recv()
	if err != nil {
		s.l.Error("[Error] cannot receive image info", "error", err)
		return err
	}

	contentType := req.GetInfo().GetContentType()
	sourceID := req.GetInfo().GetSourceId()
	sourceType := req.GetInfo().GetContentSource()
	s.l.Info("receive an upload request for content with type", "source_id", sourceID, "content_type", contentType)

	switch sourceType {
	case listingsPB.ContentInfo_APARTMENT:
		apartment, err := s.DB.GetApartment(stream.Context(), sourceID)
		if err != nil {
			s.l.Error("[Error] cannot find apartment", "error", err)
			return err
		}
		if apartment.Name == "" {
			s.l.Error("[Error] apartment id doesn't exist", "source_id", sourceID)
			return fmt.Errorf("apartment id doesn't exist")
		}
	case listingsPB.ContentInfo_BUILDING:
		building, err := s.DB.GetBuilding(stream.Context(), sourceID)
		if err != nil {
			s.l.Error("[Error] cannot find building", "error", err)
			return err
		}
		if building.Name == "" {
			s.l.Error("[Error] building id doesn't exist", "source_id", sourceID)
			return fmt.Errorf("building id doesn't exist")
		}
	default:
		s.l.Error("[Error] incorrect source type or other error", "source_type", sourceType, "error", err)
		return fmt.Errorf("incorrect source type or other error")
	}

	contentData := bytes.Buffer{}
	contentSize := 0

	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		s.l.Info("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			s.l.Info("no more data")
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)
		s.l.Info("received a chunk", "size", size)

		contentSize += size
		if contentSize > maxImageSize {
			s.l.Error("[Error] image is too large", "content_size", contentSize, "max_size", maxImageSize)
			return fmt.Errorf("image is too large")
		}

		_, err = contentData.Write(chunk)
		if err != nil {
			s.l.Error("[Error] cannot write chunk data", "error", err)
			return err
		}
	}

	contentID, contentInfo, err := s.ContentStore.Save(sourceID, contentType, contentData)
	if err != nil {
		s.l.Error("[Error] cannot save content to the store", "error", err)
		return err
	}

	_, err = s.DB.UploadContent(stream.Context(), genDB.UploadContentParams{
		ContentID:     contentID,
		Filename:      sql.NullString{String: contentInfo.Path, Valid: true},
		ContentType:   genDB.ContentType(contentInfo.ContentType),
		ContentSource: genDB.ContentSource(sourceType),
		SourceID:      sourceID,
	})
	if err != nil {
		s.l.Error("[DB] Error uploading content", "error", err)
		return err
	}

	switch sourceType {
	case listingsPB.ContentInfo_APARTMENT:
		err = s.DB.AppendContentApartment(stream.Context(), genDB.AppendContentApartmentParams{
			ApartmentID: sourceID,
			ArrayAppend: contentInfo.Path,
		})
		if err != nil {
			s.l.Error("[Error] error appending content to apartment", "error", err)
			return err
		}
	case listingsPB.ContentInfo_BUILDING:
		err = s.DB.AppendContentBuilding(stream.Context(), genDB.AppendContentBuildingParams{
			BuildingID:  sourceID,
			ArrayAppend: contentInfo.Path,
		})
		if err != nil {
			s.l.Error("[Error] error appending content to building", "error", err)
			return err
		}
	}

	res := &listingsPB.UploadContentResponse{
		ContentId: contentID,
		Size:      uint32(contentSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		s.l.Error("[Error] cannot send response", "error", err)
		return err
	}

	s.l.Info("saved content", "content_id", contentID, "content_size", contentSize)
	return nil
}
