package handler

import (
	"LearnApiGo/internal/grpc/proto"
	"LearnApiGo/internal/models"
	"LearnApiGo/internal/services"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcAlbumHandler struct {
	proto.UnimplementedGrpcAlbumServer
	service services.IAlbums
}

func New(service services.IAlbums) *GrpcAlbumHandler {
	return &GrpcAlbumHandler{
		service: service,
	}
}

func (h *GrpcAlbumHandler) CreateAlbum(ctx context.Context, r *proto.Album) (*proto.Album, error) {
	newAlbum := &models.Album{
		Title:  r.Title,
		Artist: r.Artist,
		Price:  float64(r.Price),
	}

	err := h.service.CreateAlbum(newAlbum)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	} else {
		r.Id = newAlbum.Id
		return r, nil
	}
}

func (h *GrpcAlbumHandler) GetAlbums(ctx context.Context, r *proto.GetAlbumsRequest) (*proto.GetAlbumsResponse, error) {
	albums, err := h.service.GetAlbums(int(r.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	} else {
		var resp []*proto.Album
		for _, album := range *albums {
			resp = append(resp, &proto.Album{
				Id:     album.Id,
				Title:  album.Title,
				Artist: album.Artist,
				Price:  float32(album.Price),
			})
		}
		return &proto.GetAlbumsResponse{
			Albums: resp,
		}, nil
	}
}

//func (h *GrpcAlbumHandler) GetAlbum(ctx context.Context, r *GetAlbumRequest) (*Album, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetAlbum not implemented")
//}
