package grpc

import (
	"car-service/internal/broker"
	"car-service/internal/model"
	"car-service/internal/service"
	pb "car-service/internal/transport/grpc/vehiclepb"
	"car-service/internal/validation"
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VehicleGRPCServer struct {
	pb.UnimplementedVehicleServiceServer
	svc *service.VehicleService
}

func NewVehicleGRPCServer(s *service.VehicleService) *VehicleGRPCServer {
	return &VehicleGRPCServer{svc: s}
}

func (s *VehicleGRPCServer) CreateVehicle(ctx context.Context, req *pb.CreateVehicleRequest) (*pb.VehicleResponse, error) {
	v := &model.Vehicle{
		Make: req.Make,
		Mark: req.Mark,
		Year: int(req.Year),
	}
	if err := s.svc.Create(v); err != nil {
		return nil, err
	}

	if err := broker.PublishVehicleCreated(*v); err != nil {
		logrus.Errorf("failed to publish Kafka event from gRPC: %v", err)
	}

	if err := validation.ValidateVehicle(v); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid vehicle data")
	}

	return &pb.VehicleResponse{
		Vehicle: &pb.Vehicle{
			Id:   uint64(v.ID),
			Make: v.Make,
			Mark: v.Mark,
			Year: int32(v.Year),
		},
	}, nil
}

func (s *VehicleGRPCServer) GetVehicleByID(ctx context.Context, req *pb.GetVehicleRequest) (*pb.VehicleResponse, error) {
	v, err := s.svc.GetByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.VehicleResponse{
		Vehicle: &pb.Vehicle{
			Id:   uint64(v.ID),
			Make: v.Make,
			Mark: v.Mark,
			Year: int32(v.Year),
		},
	}, nil
}

func (s *VehicleGRPCServer) ListVehicle(ctx context.Context, _ *pb.Empty) (*pb.VehicleListResponse, error) {
	vehicles, err := s.svc.ListAll()
	if err != nil {
		return nil, err
	}

	var result []*pb.Vehicle
	for _, v := range vehicles {
		result = append(result, &pb.Vehicle{
			Id:   uint64(v.ID),
			Make: v.Make,
			Mark: v.Mark,
			Year: int32(v.Year),
		})
	}

	return &pb.VehicleListResponse{
		Vehicles: result,
	}, nil
}
