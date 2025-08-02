package api

import (
	"io"
	"log"
	"time"

	"github.com/khosbilegt/llama-drover/internal/coordinator"
	"github.com/khosbilegt/llama-drover/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

// CoordinatorGRPCService defines the gRPC service interface
type CoordinatorGRPCService interface {
	Connect(stream CoordinatorGRPC_ConnectServer) error
}

// CoordinatorGRPC_ConnectServer represents the bidirectional streaming interface
type CoordinatorGRPC_ConnectServer interface {
	Send(*model.ConnectResponse) error
	Recv() (*model.ConnectRequest, error)
	grpc.ServerStream
}

// CoordinatorGRPCServer implements the gRPC service
type CoordinatorGRPCServer struct {
	// Embed UnimplementedCoordinatorGRPCServer for forward compatibility
}

// Connect handles bidirectional streaming between coordinator and nodes
func (s *CoordinatorGRPCServer) Connect(stream CoordinatorGRPC_ConnectServer) error {
	log.Println("New gRPC connection established")

	for {
		// Receive request from client
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client closed the connection")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving from stream: %v", err)
			return status.Errorf(codes.Internal, "failed to receive request: %v", err)
		}

		// Process the request based on type
		response, err := s.processRequest(req)
		if err != nil {
			log.Printf("Error processing request: %v", err)
			// Send error response
			errorResp := &model.ConnectResponse{
				Type:      req.GetType(),
				Timestamp: time.Now().Unix(),
				Response: &model.ConnectResponse_Error{
					Error: &model.ErrorResponse{
						Message: err.Error(),
					},
				},
			}
			if sendErr := stream.Send(errorResp); sendErr != nil {
				log.Printf("Error sending error response: %v", sendErr)
				return status.Errorf(codes.Internal, "failed to send error response: %v", sendErr)
			}
			continue
		}

		// Send successful response
		if err := stream.Send(response); err != nil {
			log.Printf("Error sending response: %v", err)
			return status.Errorf(codes.Internal, "failed to send response: %v", err)
		}
	}
}

// processRequest handles different types of requests
func (s *CoordinatorGRPCServer) processRequest(req *model.ConnectRequest) (*model.ConnectResponse, error) {
	timestamp := time.Now().Unix()

	switch req.GetType() {
	case model.RequestType_REGISTER:
		return s.handleRegister(req.GetRegister(), timestamp)
	case model.RequestType_UNREGISTER:
		return s.handleUnregister(req.GetUnregister(), timestamp)
	case model.RequestType_PROMPT:
		return s.handlePrompt(req.GetPrompt(), timestamp)
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unsupported request type: %v", req.GetType())
	}
}

// handleRegister processes node registration requests
func (s *CoordinatorGRPCServer) handleRegister(req *model.RegisterRequest, timestamp int64) (*model.ConnectResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "register request is nil")
	}

	log.Printf("Registering node: %s", req.NodeId)

	// Create node in database
	node := model.Node{
		ID:     req.NodeId,
		Name:   req.NodeId, // Use node ID as name for now
		Status: "ACTIVE",
	}

	_, err := coordinator.CreateNode(node)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register node: %v", err)
	}

	return &model.ConnectResponse{
		Type:      model.RequestType_REGISTER,
		Timestamp: timestamp,
		Response: &model.ConnectResponse_Register{
			Register: &model.MessageResponse{
				NodeId:  req.NodeId,
				Message: "Node registered successfully",
			},
		},
	}, nil
}

// handleUnregister processes node unregistration requests
func (s *CoordinatorGRPCServer) handleUnregister(req *model.UnregisterRequest, timestamp int64) (*model.ConnectResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "unregister request is nil")
	}

	log.Printf("Unregistering node: %s", req.NodeId)

	err := coordinator.DeleteNode(req.NodeId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unregister node: %v", err)
	}

	return &model.ConnectResponse{
		Type:      model.RequestType_UNREGISTER,
		Timestamp: timestamp,
		Response: &model.ConnectResponse_Unregister{
			Unregister: &model.UnregisterResponse{
				NodeId: req.NodeId,
			},
		},
	}, nil
}

// handlePrompt processes prompt requests and delegates to appropriate nodes
func (s *CoordinatorGRPCServer) handlePrompt(req *model.PromptRequest, timestamp int64) (*model.ConnectResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "prompt request is nil")
	}

	log.Printf("Processing prompt: %s", req.Prompt)

	// TODO: Implement actual prompt processing logic
	// This should:
	// 1. Find the best available node
	// 2. Forward the prompt to that node
	// 3. Return the response

	// For now, return a mock response
	return &model.ConnectResponse{
		Type:      model.RequestType_PROMPT,
		Timestamp: timestamp,
		Response: &model.ConnectResponse_Prompt{
			Prompt: &model.PromptResponse{
				Response:   "Mock response for: " + req.Prompt,
				IsComplete: true,
			},
		},
	}, nil
}

// NewCoordinatorGRPCServer creates a new gRPC server instance
func NewCoordinatorGRPCServer() *CoordinatorGRPCServer {
	return &CoordinatorGRPCServer{}
}
