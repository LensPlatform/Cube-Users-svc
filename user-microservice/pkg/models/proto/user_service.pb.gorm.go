// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/user_service.proto

package model

import context "context"

import gorm1 "github.com/jinzhu/gorm"

import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type UsersServiceDefaultServer struct {
	DB *gorm1.DB
}

// GetUserByID ...
func (m *UsersServiceDefaultServer) GetUserByID(ctx context.Context, in *Uint32UserTypeRequest) (*UserResponse, error) {
	return &UserResponse{}, nil
}

// GetUserByUsername ...
func (m *UsersServiceDefaultServer) GetUserByUsername(ctx context.Context, in *StringUserTypeRequest) (*UserResponse, error) {
	return &UserResponse{}, nil
}

// GetUserByEmail ...
func (m *UsersServiceDefaultServer) GetUserByEmail(ctx context.Context, in *StringUserTypeRequest) (*UserResponse, error) {
	return &UserResponse{}, nil
}

// GetUsersByType ...
func (m *UsersServiceDefaultServer) GetUsersByType(ctx context.Context, in *GetUsersSearchRequest) (*UserResponse, error) {
	return &UserResponse{}, nil
}

// CreateUser ...
func (m *UsersServiceDefaultServer) CreateUser(ctx context.Context, in *UserRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// UpdateUser ...
func (m *UsersServiceDefaultServer) UpdateUser(ctx context.Context, in *UserRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// DeleteUser ...
func (m *UsersServiceDefaultServer) DeleteUser(ctx context.Context, in *UserRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// DeleteProfileById ...
func (m *UsersServiceDefaultServer) DeleteProfileById(ctx context.Context, in *Uint32UserTypeRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// CreateTeam ...
func (m *UsersServiceDefaultServer) CreateTeam(ctx context.Context, in *TeamRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// UpdateTeam ...
func (m *UsersServiceDefaultServer) UpdateTeam(ctx context.Context, in *TeamRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// DeleteTeam ...
func (m *UsersServiceDefaultServer) DeleteTeam(ctx context.Context, in *TeamRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// GetTeamByID ...
func (m *UsersServiceDefaultServer) GetTeamByID(ctx context.Context, in *Uint32TeamTypeRequest) (*TeamResponse, error) {
	return &TeamResponse{}, nil
}

// GetTeamByName ...
func (m *UsersServiceDefaultServer) GetTeamByName(ctx context.Context, in *StringTeamTypeRequest) (*TeamResponse, error) {
	return &TeamResponse{}, nil
}

// GetTeamByEmail ...
func (m *UsersServiceDefaultServer) GetTeamByEmail(ctx context.Context, in *StringTeamTypeRequest) (*TeamResponse, error) {
	return &TeamResponse{}, nil
}

// GetTeamsByType ...
func (m *UsersServiceDefaultServer) GetTeamsByType(ctx context.Context, in *StringTeamTypeRequest) (*Team, error) {
	return &Team{}, nil
}

// GetTeamsByIndustry ...
func (m *UsersServiceDefaultServer) GetTeamsByIndustry(ctx context.Context, in *StringTeamTypeRequest) (*TeamResponse, error) {
	return &TeamResponse{}, nil
}

// CreateGroup ...
func (m *UsersServiceDefaultServer) CreateGroup(ctx context.Context, in *GroupRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// GetGroupByID ...
func (m *UsersServiceDefaultServer) GetGroupByID(ctx context.Context, in *Uint32GroupTypeRequest) (*GroupResponse, error) {
	return &GroupResponse{}, nil
}

// GetGroupByName ...
func (m *UsersServiceDefaultServer) GetGroupByName(ctx context.Context, in *StringGroupTypeRequest) (*GroupResponse, error) {
	return &GroupResponse{}, nil
}

// GetGroupsByType ...
func (m *UsersServiceDefaultServer) GetGroupsByType(ctx context.Context, in *StringGroupTypeRequest) (*GroupResponse, error) {
	return &GroupResponse{}, nil
}

// UpdateGroup ...
func (m *UsersServiceDefaultServer) UpdateGroup(ctx context.Context, in *GroupRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}

// DeleteGroup ...
func (m *UsersServiceDefaultServer) DeleteGroup(ctx context.Context, in *GroupRequest) (*UpdateOrCreateEntityResponse, error) {
	return &UpdateOrCreateEntityResponse{}, nil
}
