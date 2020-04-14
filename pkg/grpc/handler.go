package grpc

import (
	"context"
	"errors"
	endpoint "github.com/LensPlatform/micro/pkg/endpoint"
	pb "github.com/LensPlatform/micro/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeCreateUserHandler creates the handler logic
func makeCreateUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateUserEndpoint, decodeCreateUserRequest, encodeCreateUserResponse, options...)
}

// decodeCreateUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateUser request.
// TODO implement the decoder
func decodeCreateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeCreateUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) CreateUser(ctx context1.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	_, rep, err := g.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserReply), nil
}

// makeCreateProfileHandler creates the handler logic
func makeCreateProfileHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateProfileEndpoint, decodeCreateProfileRequest, encodeCreateProfileResponse, options...)
}

// decodeCreateProfileResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateProfile request.
// TODO implement the decoder
func decodeCreateProfileRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeCreateProfileResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateProfileResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) CreateProfile(ctx context1.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileReply, error) {
	_, rep, err := g.createProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateProfileReply), nil
}

// makeCreateSubscriptionHandler creates the handler logic
func makeCreateSubscriptionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateSubscriptionEndpoint, decodeCreateSubscriptionRequest, encodeCreateSubscriptionResponse, options...)
}

// decodeCreateSubscriptionResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateSubscription request.
// TODO implement the decoder
func decodeCreateSubscriptionRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeCreateSubscriptionResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateSubscriptionResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) CreateSubscription(ctx context1.Context, req *pb.CreateSubscriptionRequest) (*pb.CreateSubscriptionReply, error) {
	_, rep, err := g.createSubscription.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateSubscriptionReply), nil
}

// makeCreateSubscriptionsHandler creates the handler logic
func makeCreateSubscriptionsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateSubscriptionsEndpoint, decodeCreateSubscriptionsRequest, encodeCreateSubscriptionsResponse, options...)
}

// decodeCreateSubscriptionsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateSubscriptions request.
// TODO implement the decoder
func decodeCreateSubscriptionsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeCreateSubscriptionsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateSubscriptionsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) CreateSubscriptions(ctx context1.Context, req *pb.CreateSubscriptionsRequest) (*pb.CreateSubscriptionsReply, error) {
	_, rep, err := g.createSubscriptions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateSubscriptionsReply), nil
}

// makeUpdateUserHandler creates the handler logic
func makeUpdateUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserEndpoint, decodeUpdateUserRequest, encodeUpdateUserResponse, options...)
}

// decodeUpdateUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUser request.
// TODO implement the decoder
func decodeUpdateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUser(ctx context1.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	_, rep, err := g.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserReply), nil
}

// makeUpdateUserProfileHandler creates the handler logic
func makeUpdateUserProfileHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserProfileEndpoint, decodeUpdateUserProfileRequest, encodeUpdateUserProfileResponse, options...)
}

// decodeUpdateUserProfileResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUserProfile request.
// TODO implement the decoder
func decodeUpdateUserProfileRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateUserProfileResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserProfileResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUserProfile(ctx context1.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileReply, error) {
	_, rep, err := g.updateUserProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserProfileReply), nil
}

// makeUpdateUserSubscriptionHandler creates the handler logic
func makeUpdateUserSubscriptionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserSubscriptionEndpoint, decodeUpdateUserSubscriptionRequest, encodeUpdateUserSubscriptionResponse, options...)
}

// decodeUpdateUserSubscriptionResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUserSubscription request.
// TODO implement the decoder
func decodeUpdateUserSubscriptionRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateUserSubscriptionResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserSubscriptionResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUserSubscription(ctx context1.Context, req *pb.UpdateUserSubscriptionRequest) (*pb.UpdateUserSubscriptionReply, error) {
	_, rep, err := g.updateUserSubscription.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserSubscriptionReply), nil
}

// makeDeleteUserHandler creates the handler logic
func makeDeleteUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteUserEndpoint, decodeDeleteUserRequest, encodeDeleteUserResponse, options...)
}

// decodeDeleteUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteUser request.
// TODO implement the decoder
func decodeDeleteUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteUser(ctx context1.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	_, rep, err := g.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserReply), nil
}

// makeDeleteUserProfileHandler creates the handler logic
func makeDeleteUserProfileHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteUserProfileEndpoint, decodeDeleteUserProfileRequest, encodeDeleteUserProfileResponse, options...)
}

// decodeDeleteUserProfileResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteUserProfile request.
// TODO implement the decoder
func decodeDeleteUserProfileRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteUserProfileResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteUserProfileResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteUserProfile(ctx context1.Context, req *pb.DeleteUserProfileRequest) (*pb.DeleteUserProfileReply, error) {
	_, rep, err := g.deleteUserProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserProfileReply), nil
}

// makeDeleteSubscriptionHandler creates the handler logic
func makeDeleteSubscriptionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteSubscriptionEndpoint, decodeDeleteSubscriptionRequest, encodeDeleteSubscriptionResponse, options...)
}

// decodeDeleteSubscriptionResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteSubscription request.
// TODO implement the decoder
func decodeDeleteSubscriptionRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteSubscriptionResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteSubscriptionResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteSubscription(ctx context1.Context, req *pb.DeleteSubscriptionRequest) (*pb.DeleteSubscriptionReply, error) {
	_, rep, err := g.deleteSubscription.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteSubscriptionReply), nil
}

// makeGetUserHandler creates the handler logic
func makeGetUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserEndpoint, decodeGetUserRequest, encodeGetUserResponse, options...)
}

// decodeGetUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUser request.
// TODO implement the decoder
func decodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUser(ctx context1.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	_, rep, err := g.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserReply), nil
}

// makeGetUsersHandler creates the handler logic
func makeGetUsersHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUsersEndpoint, decodeGetUsersRequest, encodeGetUsersResponse, options...)
}

// decodeGetUsersResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUsers request.
// TODO implement the decoder
func decodeGetUsersRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUsersResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUsersResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUsers(ctx context1.Context, req *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	_, rep, err := g.getUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUsersReply), nil
}

// makeGetUsersByAccountTypeHandler creates the handler logic
func makeGetUsersByAccountTypeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUsersByAccountTypeEndpoint, decodeGetUsersByAccountTypeRequest, encodeGetUsersByAccountTypeResponse, options...)
}

// decodeGetUsersByAccountTypeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUsersByAccountType request.
// TODO implement the decoder
func decodeGetUsersByAccountTypeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUsersByAccountTypeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUsersByAccountTypeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUsersByAccountType(ctx context1.Context, req *pb.GetUsersByAccountTypeRequest) (*pb.GetUsersByAccountTypeReply, error) {
	_, rep, err := g.getUsersByAccountType.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUsersByAccountTypeReply), nil
}

// makeGetUsersByIntentHandler creates the handler logic
func makeGetUsersByIntentHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUsersByIntentEndpoint, decodeGetUsersByIntentRequest, encodeGetUsersByIntentResponse, options...)
}

// decodeGetUsersByIntentResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUsersByIntent request.
// TODO implement the decoder
func decodeGetUsersByIntentRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUsersByIntentResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUsersByIntentResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUsersByIntent(ctx context1.Context, req *pb.GetUsersByIntentRequest) (*pb.GetUsersByIntentReply, error) {
	_, rep, err := g.getUsersByIntent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUsersByIntentReply), nil
}

// makeGetUserProfileHandler creates the handler logic
func makeGetUserProfileHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserProfileEndpoint, decodeGetUserProfileRequest, encodeGetUserProfileResponse, options...)
}

// decodeGetUserProfileResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserProfile request.
// TODO implement the decoder
func decodeGetUserProfileRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUserProfileResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserProfileResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUserProfile(ctx context1.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileReply, error) {
	_, rep, err := g.getUserProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserProfileReply), nil
}

// makeGetUserProfilesHandler creates the handler logic
func makeGetUserProfilesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserProfilesEndpoint, decodeGetUserProfilesRequest, encodeGetUserProfilesResponse, options...)
}

// decodeGetUserProfilesResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserProfiles request.
// TODO implement the decoder
func decodeGetUserProfilesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUserProfilesResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserProfilesResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUserProfiles(ctx context1.Context, req *pb.GetUserProfilesRequest) (*pb.GetUserProfilesReply, error) {
	_, rep, err := g.getUserProfiles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserProfilesReply), nil
}

// makeGetUserProfilesByTypeHandler creates the handler logic
func makeGetUserProfilesByTypeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserProfilesByTypeEndpoint, decodeGetUserProfilesByTypeRequest, encodeGetUserProfilesByTypeResponse, options...)
}

// decodeGetUserProfilesByTypeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserProfilesByType request.
// TODO implement the decoder
func decodeGetUserProfilesByTypeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUserProfilesByTypeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserProfilesByTypeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUserProfilesByType(ctx context1.Context, req *pb.GetUserProfilesByTypeRequest) (*pb.GetUserProfilesByTypeReply, error) {
	_, rep, err := g.getUserProfilesByType.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserProfilesByTypeReply), nil
}

// makeGetUserProfilesByNationalityHandler creates the handler logic
func makeGetUserProfilesByNationalityHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserProfilesByNationalityEndpoint, decodeGetUserProfilesByNationalityRequest, encodeGetUserProfilesByNationalityResponse, options...)
}

// decodeGetUserProfilesByNationalityResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserProfilesByNationality request.
// TODO implement the decoder
func decodeGetUserProfilesByNationalityRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUserProfilesByNationalityResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserProfilesByNationalityResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUserProfilesByNationality(ctx context1.Context, req *pb.GetUserProfilesByNationalityRequest) (*pb.GetUserProfilesByNationalityReply, error) {
	_, rep, err := g.getUserProfilesByNationality.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserProfilesByNationalityReply), nil
}

// makeGetUserSubscriptionsHandler creates the handler logic
func makeGetUserSubscriptionsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserSubscriptionsEndpoint, decodeGetUserSubscriptionsRequest, encodeGetUserSubscriptionsResponse, options...)
}

// decodeGetUserSubscriptionsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserSubscriptions request.
// TODO implement the decoder
func decodeGetUserSubscriptionsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetUserSubscriptionsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserSubscriptionsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetUserSubscriptions(ctx context1.Context, req *pb.GetUserSubscriptionsRequest) (*pb.GetUserSubscriptionsReply, error) {
	_, rep, err := g.getUserSubscriptions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserSubscriptionsReply), nil
}

// makeCreateTeamHandler creates the handler logic
func makeCreateTeamHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateTeamEndpoint, decodeCreateTeamRequest, encodeCreateTeamResponse, options...)
}

// decodeCreateTeamResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateTeam request.
// TODO implement the decoder
func decodeCreateTeamRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeCreateTeamResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateTeamResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) CreateTeam(ctx context1.Context, req *pb.CreateTeamRequest) (*pb.CreateTeamReply, error) {
	_, rep, err := g.createTeam.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateTeamReply), nil
}

// makeCreatTeamProfileHandler creates the handler logic
func makeCreatTeamProfileHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreatTeamProfileEndpoint, decodeCreatTeamProfileRequest, encodeCreatTeamProfileResponse, options...)
}

// decodeCreatTeamProfileResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreatTeamProfile request.
// TODO implement the decoder
func decodeCreatTeamProfileRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeCreatTeamProfileResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreatTeamProfileResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) CreatTeamProfile(ctx context1.Context, req *pb.CreatTeamProfileRequest) (*pb.CreatTeamProfileReply, error) {
	_, rep, err := g.creatTeamProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreatTeamProfileReply), nil
}

// makeCreateTeamSubscriptionHandler creates the handler logic
func makeCreateTeamSubscriptionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateTeamSubscriptionEndpoint, decodeCreateTeamSubscriptionRequest, encodeCreateTeamSubscriptionResponse, options...)
}

// decodeCreateTeamSubscriptionResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateTeamSubscription request.
// TODO implement the decoder
func decodeCreateTeamSubscriptionRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeCreateTeamSubscriptionResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateTeamSubscriptionResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) CreateTeamSubscription(ctx context1.Context, req *pb.CreateTeamSubscriptionRequest) (*pb.CreateTeamSubscriptionReply, error) {
	_, rep, err := g.createTeamSubscription.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateTeamSubscriptionReply), nil
}

// makeAddMemberToTeamHandler creates the handler logic
func makeAddMemberToTeamHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddMemberToTeamEndpoint, decodeAddMemberToTeamRequest, encodeAddMemberToTeamResponse, options...)
}

// decodeAddMemberToTeamResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AddMemberToTeam request.
// TODO implement the decoder
func decodeAddMemberToTeamRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeAddMemberToTeamResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddMemberToTeamResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) AddMemberToTeam(ctx context1.Context, req *pb.AddMemberToTeamRequest) (*pb.AddMemberToTeamReply, error) {
	_, rep, err := g.addMemberToTeam.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddMemberToTeamReply), nil
}

// makeAddAdvisorToTeamHandler creates the handler logic
func makeAddAdvisorToTeamHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddAdvisorToTeamEndpoint, decodeAddAdvisorToTeamRequest, encodeAddAdvisorToTeamResponse, options...)
}

// decodeAddAdvisorToTeamResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AddAdvisorToTeam request.
// TODO implement the decoder
func decodeAddAdvisorToTeamRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeAddAdvisorToTeamResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddAdvisorToTeamResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) AddAdvisorToTeam(ctx context1.Context, req *pb.AddAdvisorToTeamRequest) (*pb.AddAdvisorToTeamReply, error) {
	_, rep, err := g.addAdvisorToTeam.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddAdvisorToTeamReply), nil
}

// makeUpdateTeamHandler creates the handler logic
func makeUpdateTeamHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTeamEndpoint, decodeUpdateTeamRequest, encodeUpdateTeamResponse, options...)
}

// decodeUpdateTeamResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTeam request.
// TODO implement the decoder
func decodeUpdateTeamRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateTeamResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTeamResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTeam(ctx context1.Context, req *pb.UpdateTeamRequest) (*pb.UpdateTeamReply, error) {
	_, rep, err := g.updateTeam.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateTeamReply), nil
}

// makeUpdateTeamAdminHandler creates the handler logic
func makeUpdateTeamAdminHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTeamAdminEndpoint, decodeUpdateTeamAdminRequest, encodeUpdateTeamAdminResponse, options...)
}

// decodeUpdateTeamAdminResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTeamAdmin request.
// TODO implement the decoder
func decodeUpdateTeamAdminRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateTeamAdminResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTeamAdminResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTeamAdmin(ctx context1.Context, req *pb.UpdateTeamAdminRequest) (*pb.UpdateTeamAdminReply, error) {
	_, rep, err := g.updateTeamAdmin.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateTeamAdminReply), nil
}

// makeUpdateTeamMemberHandler creates the handler logic
func makeUpdateTeamMemberHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTeamMemberEndpoint, decodeUpdateTeamMemberRequest, encodeUpdateTeamMemberResponse, options...)
}

// decodeUpdateTeamMemberResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTeamMember request.
// TODO implement the decoder
func decodeUpdateTeamMemberRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateTeamMemberResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTeamMemberResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTeamMember(ctx context1.Context, req *pb.UpdateTeamMemberRequest) (*pb.UpdateTeamMemberReply, error) {
	_, rep, err := g.updateTeamMember.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateTeamMemberReply), nil
}

// makeUpdateTeamAdvisorHandler creates the handler logic
func makeUpdateTeamAdvisorHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTeamAdvisorEndpoint, decodeUpdateTeamAdvisorRequest, encodeUpdateTeamAdvisorResponse, options...)
}

// decodeUpdateTeamAdvisorResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTeamAdvisor request.
// TODO implement the decoder
func decodeUpdateTeamAdvisorRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateTeamAdvisorResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTeamAdvisorResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTeamAdvisor(ctx context1.Context, req *pb.UpdateTeamAdvisorRequest) (*pb.UpdateTeamAdvisorReply, error) {
	_, rep, err := g.updateTeamAdvisor.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateTeamAdvisorReply), nil
}

// makeUpdateTeamProfileHandler creates the handler logic
func makeUpdateTeamProfileHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTeamProfileEndpoint, decodeUpdateTeamProfileRequest, encodeUpdateTeamProfileResponse, options...)
}

// decodeUpdateTeamProfileResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTeamProfile request.
// TODO implement the decoder
func decodeUpdateTeamProfileRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateTeamProfileResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTeamProfileResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTeamProfile(ctx context1.Context, req *pb.UpdateTeamProfileRequest) (*pb.UpdateTeamProfileReply, error) {
	_, rep, err := g.updateTeamProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateTeamProfileReply), nil
}

// makeUpdateTeamSubscriptionHandler creates the handler logic
func makeUpdateTeamSubscriptionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTeamSubscriptionEndpoint, decodeUpdateTeamSubscriptionRequest, encodeUpdateTeamSubscriptionResponse, options...)
}

// decodeUpdateTeamSubscriptionResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTeamSubscription request.
// TODO implement the decoder
func decodeUpdateTeamSubscriptionRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateTeamSubscriptionResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTeamSubscriptionResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTeamSubscription(ctx context1.Context, req *pb.UpdateTeamSubscriptionRequest) (*pb.UpdateTeamSubscriptionReply, error) {
	_, rep, err := g.updateTeamSubscription.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateTeamSubscriptionReply), nil
}

// makeDeleteTeamHandler creates the handler logic
func makeDeleteTeamHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteTeamEndpoint, decodeDeleteTeamRequest, encodeDeleteTeamResponse, options...)
}

// decodeDeleteTeamResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteTeam request.
// TODO implement the decoder
func decodeDeleteTeamRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteTeamResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteTeamResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteTeam(ctx context1.Context, req *pb.DeleteTeamRequest) (*pb.DeleteTeamReply, error) {
	_, rep, err := g.deleteTeam.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTeamReply), nil
}

// makeDeleteTeamProfileHandler creates the handler logic
func makeDeleteTeamProfileHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteTeamProfileEndpoint, decodeDeleteTeamProfileRequest, encodeDeleteTeamProfileResponse, options...)
}

// decodeDeleteTeamProfileResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteTeamProfile request.
// TODO implement the decoder
func decodeDeleteTeamProfileRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteTeamProfileResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteTeamProfileResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteTeamProfile(ctx context1.Context, req *pb.DeleteTeamProfileRequest) (*pb.DeleteTeamProfileReply, error) {
	_, rep, err := g.deleteTeamProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTeamProfileReply), nil
}

// makeDeleteTeamMemberHandler creates the handler logic
func makeDeleteTeamMemberHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteTeamMemberEndpoint, decodeDeleteTeamMemberRequest, encodeDeleteTeamMemberResponse, options...)
}

// decodeDeleteTeamMemberResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteTeamMember request.
// TODO implement the decoder
func decodeDeleteTeamMemberRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteTeamMemberResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteTeamMemberResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteTeamMember(ctx context1.Context, req *pb.DeleteTeamMemberRequest) (*pb.DeleteTeamMemberReply, error) {
	_, rep, err := g.deleteTeamMember.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTeamMemberReply), nil
}

// makeDeleteTeamAdvisorHandler creates the handler logic
func makeDeleteTeamAdvisorHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteTeamAdvisorEndpoint, decodeDeleteTeamAdvisorRequest, encodeDeleteTeamAdvisorResponse, options...)
}

// decodeDeleteTeamAdvisorResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteTeamAdvisor request.
// TODO implement the decoder
func decodeDeleteTeamAdvisorRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteTeamAdvisorResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteTeamAdvisorResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteTeamAdvisor(ctx context1.Context, req *pb.DeleteTeamAdvisorRequest) (*pb.DeleteTeamAdvisorReply, error) {
	_, rep, err := g.deleteTeamAdvisor.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTeamAdvisorReply), nil
}

// makeDeleteTeamAdminHandler creates the handler logic
func makeDeleteTeamAdminHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteTeamAdminEndpoint, decodeDeleteTeamAdminRequest, encodeDeleteTeamAdminResponse, options...)
}

// decodeDeleteTeamAdminResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteTeamAdmin request.
// TODO implement the decoder
func decodeDeleteTeamAdminRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteTeamAdminResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteTeamAdminResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteTeamAdmin(ctx context1.Context, req *pb.DeleteTeamAdminRequest) (*pb.DeleteTeamAdminReply, error) {
	_, rep, err := g.deleteTeamAdmin.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTeamAdminReply), nil
}

// makeDeleteTeamSubscriptionHandler creates the handler logic
func makeDeleteTeamSubscriptionHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteTeamSubscriptionEndpoint, decodeDeleteTeamSubscriptionRequest, encodeDeleteTeamSubscriptionResponse, options...)
}

// decodeDeleteTeamSubscriptionResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteTeamSubscription request.
// TODO implement the decoder
func decodeDeleteTeamSubscriptionRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteTeamSubscriptionResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteTeamSubscriptionResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteTeamSubscription(ctx context1.Context, req *pb.DeleteTeamSubscriptionRequest) (*pb.DeleteTeamSubscriptionReply, error) {
	_, rep, err := g.deleteTeamSubscription.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTeamSubscriptionReply), nil
}

// makeGetTeamHandler creates the handler logic
func makeGetTeamHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamEndpoint, decodeGetTeamRequest, encodeGetTeamResponse, options...)
}

// decodeGetTeamResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeam request.
// TODO implement the decoder
func decodeGetTeamRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeam(ctx context1.Context, req *pb.GetTeamRequest) (*pb.GetTeamReply, error) {
	_, rep, err := g.getTeam.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamReply), nil
}

// makeGetTeamProfileHandler creates the handler logic
func makeGetTeamProfileHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamProfileEndpoint, decodeGetTeamProfileRequest, encodeGetTeamProfileResponse, options...)
}

// decodeGetTeamProfileResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamProfile request.
// TODO implement the decoder
func decodeGetTeamProfileRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamProfileResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamProfileResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamProfile(ctx context1.Context, req *pb.GetTeamProfileRequest) (*pb.GetTeamProfileReply, error) {
	_, rep, err := g.getTeamProfile.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamProfileReply), nil
}

// makeGetTeamSubscriptionsHandler creates the handler logic
func makeGetTeamSubscriptionsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamSubscriptionsEndpoint, decodeGetTeamSubscriptionsRequest, encodeGetTeamSubscriptionsResponse, options...)
}

// decodeGetTeamSubscriptionsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamSubscriptions request.
// TODO implement the decoder
func decodeGetTeamSubscriptionsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamSubscriptionsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamSubscriptionsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamSubscriptions(ctx context1.Context, req *pb.GetTeamSubscriptionsRequest) (*pb.GetTeamSubscriptionsReply, error) {
	_, rep, err := g.getTeamSubscriptions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamSubscriptionsReply), nil
}

// makeGetTeamMembersHandler creates the handler logic
func makeGetTeamMembersHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamMembersEndpoint, decodeGetTeamMembersRequest, encodeGetTeamMembersResponse, options...)
}

// decodeGetTeamMembersResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamMembers request.
// TODO implement the decoder
func decodeGetTeamMembersRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamMembersResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamMembersResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamMembers(ctx context1.Context, req *pb.GetTeamMembersRequest) (*pb.GetTeamMembersReply, error) {
	_, rep, err := g.getTeamMembers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamMembersReply), nil
}

// makeGetTeamAdvisorsHandler creates the handler logic
func makeGetTeamAdvisorsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamAdvisorsEndpoint, decodeGetTeamAdvisorsRequest, encodeGetTeamAdvisorsResponse, options...)
}

// decodeGetTeamAdvisorsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamAdvisors request.
// TODO implement the decoder
func decodeGetTeamAdvisorsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamAdvisorsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamAdvisorsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamAdvisors(ctx context1.Context, req *pb.GetTeamAdvisorsRequest) (*pb.GetTeamAdvisorsReply, error) {
	_, rep, err := g.getTeamAdvisors.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamAdvisorsReply), nil
}

// makeGetTeamAdminHandler creates the handler logic
func makeGetTeamAdminHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamAdminEndpoint, decodeGetTeamAdminRequest, encodeGetTeamAdminResponse, options...)
}

// decodeGetTeamAdminResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamAdmin request.
// TODO implement the decoder
func decodeGetTeamAdminRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamAdminResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamAdminResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamAdmin(ctx context1.Context, req *pb.GetTeamAdminRequest) (*pb.GetTeamAdminReply, error) {
	_, rep, err := g.getTeamAdmin.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamAdminReply), nil
}

// makeGetTeamsHandler creates the handler logic
func makeGetTeamsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamsEndpoint, decodeGetTeamsRequest, encodeGetTeamsResponse, options...)
}

// decodeGetTeamsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeams request.
// TODO implement the decoder
func decodeGetTeamsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeams(ctx context1.Context, req *pb.GetTeamsRequest) (*pb.GetTeamsReply, error) {
	_, rep, err := g.getTeams.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamsReply), nil
}

// makeGetTeamsByTypeHandler creates the handler logic
func makeGetTeamsByTypeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamsByTypeEndpoint, decodeGetTeamsByTypeRequest, encodeGetTeamsByTypeResponse, options...)
}

// decodeGetTeamsByTypeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamsByType request.
// TODO implement the decoder
func decodeGetTeamsByTypeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamsByTypeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamsByTypeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamsByType(ctx context1.Context, req *pb.GetTeamsByTypeRequest) (*pb.GetTeamsByTypeReply, error) {
	_, rep, err := g.getTeamsByType.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamsByTypeReply), nil
}

// makeGetTeamsByIndustryHandler creates the handler logic
func makeGetTeamsByIndustryHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamsByIndustryEndpoint, decodeGetTeamsByIndustryRequest, encodeGetTeamsByIndustryResponse, options...)
}

// decodeGetTeamsByIndustryResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamsByIndustry request.
// TODO implement the decoder
func decodeGetTeamsByIndustryRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamsByIndustryResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamsByIndustryResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamsByIndustry(ctx context1.Context, req *pb.GetTeamsByIndustryRequest) (*pb.GetTeamsByIndustryReply, error) {
	_, rep, err := g.getTeamsByIndustry.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamsByIndustryReply), nil
}

// makeGetTeamsByNumberOfEmployeesHandler creates the handler logic
func makeGetTeamsByNumberOfEmployeesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamsByNumberOfEmployeesEndpoint, decodeGetTeamsByNumberOfEmployeesRequest, encodeGetTeamsByNumberOfEmployeesResponse, options...)
}

// decodeGetTeamsByNumberOfEmployeesResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamsByNumberOfEmployees request.
// TODO implement the decoder
func decodeGetTeamsByNumberOfEmployeesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamsByNumberOfEmployeesResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamsByNumberOfEmployeesResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamsByNumberOfEmployees(ctx context1.Context, req *pb.GetTeamsByNumberOfEmployeesRequest) (*pb.GetTeamsByNumberOfEmployeesReply, error) {
	_, rep, err := g.getTeamsByNumberOfEmployees.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamsByNumberOfEmployeesReply), nil
}

// makeGetTeamProfilesHandler creates the handler logic
func makeGetTeamProfilesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamProfilesEndpoint, decodeGetTeamProfilesRequest, encodeGetTeamProfilesResponse, options...)
}

// decodeGetTeamProfilesResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamProfiles request.
// TODO implement the decoder
func decodeGetTeamProfilesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamProfilesResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamProfilesResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamProfiles(ctx context1.Context, req *pb.GetTeamProfilesRequest) (*pb.GetTeamProfilesReply, error) {
	_, rep, err := g.getTeamProfiles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamProfilesReply), nil
}

// makeGetTeamsByTagsHandler creates the handler logic
func makeGetTeamsByTagsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamsByTagsEndpoint, decodeGetTeamsByTagsRequest, encodeGetTeamsByTagsResponse, options...)
}

// decodeGetTeamsByTagsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamsByTags request.
// TODO implement the decoder
func decodeGetTeamsByTagsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetTeamsByTagsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamsByTagsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamsByTags(ctx context1.Context, req *pb.GetTeamsByTagsRequest) (*pb.GetTeamsByTagsReply, error) {
	_, rep, err := g.getTeamsByTags.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamsByTagsReply), nil
}

// makeCreateGroupHandler creates the handler logic
func makeCreateGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateGroupEndpoint, decodeCreateGroupRequest, encodeCreateGroupResponse, options...)
}

// decodeCreateGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateGroup request.
// TODO implement the decoder
func decodeCreateGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeCreateGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) CreateGroup(ctx context1.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	_, rep, err := g.createGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateGroupReply), nil
}

// makeAddGroupMemberHandler creates the handler logic
func makeAddGroupMemberHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddGroupMemberEndpoint, decodeAddGroupMemberRequest, encodeAddGroupMemberResponse, options...)
}

// decodeAddGroupMemberResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AddGroupMember request.
// TODO implement the decoder
func decodeAddGroupMemberRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeAddGroupMemberResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddGroupMemberResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) AddGroupMember(ctx context1.Context, req *pb.AddGroupMemberRequest) (*pb.AddGroupMemberReply, error) {
	_, rep, err := g.addGroupMember.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddGroupMemberReply), nil
}

// makeUpdateGroupHandler creates the handler logic
func makeUpdateGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateGroupEndpoint, decodeUpdateGroupRequest, encodeUpdateGroupResponse, options...)
}

// decodeUpdateGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateGroup request.
// TODO implement the decoder
func decodeUpdateGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateGroup(ctx context1.Context, req *pb.UpdateGroupRequest) (*pb.UpdateGroupReply, error) {
	_, rep, err := g.updateGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateGroupReply), nil
}

// makeUpdateGroupMemberHandler creates the handler logic
func makeUpdateGroupMemberHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateGroupMemberEndpoint, decodeUpdateGroupMemberRequest, encodeUpdateGroupMemberResponse, options...)
}

// decodeUpdateGroupMemberResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateGroupMember request.
// TODO implement the decoder
func decodeUpdateGroupMemberRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeUpdateGroupMemberResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateGroupMemberResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) UpdateGroupMember(ctx context1.Context, req *pb.UpdateGroupMemberRequest) (*pb.UpdateGroupMemberReply, error) {
	_, rep, err := g.updateGroupMember.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateGroupMemberReply), nil
}

// makeDeleteGroupHandler creates the handler logic
func makeDeleteGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteGroupEndpoint, decodeDeleteGroupRequest, encodeDeleteGroupResponse, options...)
}

// decodeDeleteGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteGroup request.
// TODO implement the decoder
func decodeDeleteGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteGroup(ctx context1.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupReply, error) {
	_, rep, err := g.deleteGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteGroupReply), nil
}

// makeDeleteGroupMemberHandler creates the handler logic
func makeDeleteGroupMemberHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteGroupMemberEndpoint, decodeDeleteGroupMemberRequest, encodeDeleteGroupMemberResponse, options...)
}

// decodeDeleteGroupMemberResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteGroupMember request.
// TODO implement the decoder
func decodeDeleteGroupMemberRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeDeleteGroupMemberResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteGroupMemberResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) DeleteGroupMember(ctx context1.Context, req *pb.DeleteGroupMemberRequest) (*pb.DeleteGroupMemberReply, error) {
	_, rep, err := g.deleteGroupMember.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteGroupMemberReply), nil
}

// makeGetGroupHandler creates the handler logic
func makeGetGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupEndpoint, decodeGetGroupRequest, encodeGetGroupResponse, options...)
}

// decodeGetGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGroup request.
// TODO implement the decoder
func decodeGetGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetGroup(ctx context1.Context, req *pb.GetGroupRequest) (*pb.GetGroupReply, error) {
	_, rep, err := g.getGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupReply), nil
}

// makeGetGroupsHandler creates the handler logic
func makeGetGroupsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupsEndpoint, decodeGetGroupsRequest, encodeGetGroupsResponse, options...)
}

// decodeGetGroupsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGroups request.
// TODO implement the decoder
func decodeGetGroupsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetGroupsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGroupsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetGroups(ctx context1.Context, req *pb.GetGroupsRequest) (*pb.GetGroupsReply, error) {
	_, rep, err := g.getGroups.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupsReply), nil
}

// makeGetGroupsByTypeHandler creates the handler logic
func makeGetGroupsByTypeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupsByTypeEndpoint, decodeGetGroupsByTypeRequest, encodeGetGroupsByTypeResponse, options...)
}

// decodeGetGroupsByTypeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGroupsByType request.
// TODO implement the decoder
func decodeGetGroupsByTypeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetGroupsByTypeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGroupsByTypeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetGroupsByType(ctx context1.Context, req *pb.GetGroupsByTypeRequest) (*pb.GetGroupsByTypeReply, error) {
	_, rep, err := g.getGroupsByType.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupsByTypeReply), nil
}

// makeGetGroupsByNumMembersHandler creates the handler logic
func makeGetGroupsByNumMembersHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupsByNumMembersEndpoint, decodeGetGroupsByNumMembersRequest, encodeGetGroupsByNumMembersResponse, options...)
}

// decodeGetGroupsByNumMembersResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGroupsByNumMembers request.
// TODO implement the decoder
func decodeGetGroupsByNumMembersRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetGroupsByNumMembersResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGroupsByNumMembersResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetGroupsByNumMembers(ctx context1.Context, req *pb.GetGroupsByNumMembersRequest) (*pb.GetGroupsByNumMembersReply, error) {
	_, rep, err := g.getGroupsByNumMembers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupsByNumMembersReply), nil
}

// makeGetGroupsByTagsHandler creates the handler logic
func makeGetGroupsByTagsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupsByTagsEndpoint, decodeGetGroupsByTagsRequest, encodeGetGroupsByTagsResponse, options...)
}

// decodeGetGroupsByTagsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGroupsByTags request.
// TODO implement the decoder
func decodeGetGroupsByTagsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Decoder is not impelemented")
}

// encodeGetGroupsByTagsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGroupsByTagsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Micro' Encoder is not impelemented")
}
func (g *grpcServer) GetGroupsByTags(ctx context1.Context, req *pb.GetGroupsByTagsRequest) (*pb.GetGroupsByTagsReply, error) {
	_, rep, err := g.getGroupsByTags.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupsByTagsReply), nil
}
