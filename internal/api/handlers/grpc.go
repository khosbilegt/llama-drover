package handlers

import (
	"log"

	pb "github.com/khosbilegt/llama-drover/internal/model"
)

type CoordinatorGRPCServer struct {
	pb.UnimplementedCoordinatorServer
}

func ptrString(s string) *string {
	return &s
}

func (s *CoordinatorGRPCServer) Connect(stream pb.Coordinator_ConnectServer) error {
	log.Println("Client connected to Coordinator service")
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Println("Error receiving from stream:", err)
			return err
		}

		log.Printf("Received request: %+v\n", req)

		switch req.Request.(type) {
		case *pb.ConnectRequest_Register:
			registerRequest := req.GetRegister()
			log.Printf("Registering node with ID: %s\n", registerRequest.GetNodeId())
			stream.Send(&pb.ConnectResponse{
				Timestamp: req.Timestamp,
				Type:      req.Type,
				Response: &pb.ConnectResponse_Register{
					Register: &pb.RegisterResponse{
						NodeId:  ptrString(registerRequest.GetNodeId()),
						Message: ptrString("Node registered successfully"),
					},
				},
			})
			log.Println("Node registered successfully")
		default:
			stream.Send(&pb.ConnectResponse{
				Timestamp: req.Timestamp,
				Response: &pb.ConnectResponse_Error{
					Error: &pb.ErrorResponse{
						Message: ptrString("Invalid request type"),
					},
				},
			})
			log.Println("Invalid request type received")
			return nil
		}

		// err = stream.Send(&pb.ConnectResponse{
		// 	Timestamp: req.Timestamp,
		// 	Response: &pb.ConnectResponse_Register{Register: &pb.RegisterResponse{
		// 		NodeId:  req.GetRegister().NodeId,
		// 		Message: ptrString("Registered successfully"),
		// 	}},
		// })

		// if err != nil {
		// 	log.Println("Error sending response:", err)
		// 	return err
		// }
	}
}
