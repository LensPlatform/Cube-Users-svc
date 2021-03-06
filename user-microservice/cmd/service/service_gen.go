// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint "github.com/LensPlatform/micro/user-microservice/pkg/endpoint"
	http1 "github.com/LensPlatform/micro/user-microservice/pkg/http"
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	initGRPCHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"AddAdvisorToTeam":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddAdvisorToTeam", logger))},
		"AddGroupMember":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddGroupMember", logger))},
		"AddMemberToTeam":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddMemberToTeam", logger))},
		"CreatTeamProfile":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreatTeamProfile", logger))},
		"CreateGroup":                  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateGroup", logger))},
		"CreateProfile":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateProfile", logger))},
		"CreateSubscription":           {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateSubscription", logger))},
		"CreateSubscriptions":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateSubscriptions", logger))},
		"CreateTeam":                   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateTeam", logger))},
		"CreateTeamSubscription":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateTeamSubscription", logger))},
		"CreateUser":                   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateUser", logger))},
		"DeleteGroup":                  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteGroup", logger))},
		"DeleteGroupMember":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteGroupMember", logger))},
		"DeleteSubscription":           {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteSubscription", logger))},
		"DeleteTeam":                   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteTeam", logger))},
		"DeleteTeamAdmin":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteTeamAdmin", logger))},
		"DeleteTeamAdvisor":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteTeamAdvisor", logger))},
		"DeleteTeamMember":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteTeamMember", logger))},
		"DeleteTeamProfile":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteTeamProfile", logger))},
		"DeleteTeamSubscription":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteTeamSubscription", logger))},
		"DeleteUser":                   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteUser", logger))},
		"DeleteUserProfile":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteUserProfile", logger))},
		"GetGroup":                     {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroup", logger))},
		"GetGroups":                    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroups", logger))},
		"GetGroupsByNumMembers":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroupsByNumMembers", logger))},
		"GetGroupsByTags":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroupsByTags", logger))},
		"GetGroupsByType":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroupsByType", logger))},
		"GetTeam":                      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeam", logger))},
		"GetTeamAdmin":                 {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamAdmin", logger))},
		"GetTeamAdvisors":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamAdvisors", logger))},
		"GetTeamMembers":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamMembers", logger))},
		"GetTeamProfile":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamProfile", logger))},
		"GetTeamProfiles":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamProfiles", logger))},
		"GetTeamSubscriptions":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamSubscriptions", logger))},
		"GetTeams":                     {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeams", logger))},
		"GetTeamsByIndustry":           {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamsByIndustry", logger))},
		"GetTeamsByNumberOfEmployees":  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamsByNumberOfEmployees", logger))},
		"GetTeamsByTags":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamsByTags", logger))},
		"GetTeamsByType":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamsByType", logger))},
		"GetUser":                      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUser", logger))},
		"GetUserProfile":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserProfile", logger))},
		"GetUserProfiles":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserProfiles", logger))},
		"GetUserProfilesByNationality": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserProfilesByNationality", logger))},
		"GetUserProfilesByType":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserProfilesByType", logger))},
		"GetUserSubscriptions":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserSubscriptions", logger))},
		"GetUsers":                     {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUsers", logger))},
		"GetUsersByAccountType":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUsersByAccountType", logger))},
		"GetUsersByIntent":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUsersByIntent", logger))},
		"UpdateGroup":                  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateGroup", logger))},
		"UpdateGroupMember":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateGroupMember", logger))},
		"UpdateTeam":                   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateTeam", logger))},
		"UpdateTeamAdmin":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateTeamAdmin", logger))},
		"UpdateTeamAdvisor":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateTeamAdvisor", logger))},
		"UpdateTeamMember":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateTeamMember", logger))},
		"UpdateTeamProfile":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateTeamProfile", logger))},
		"UpdateTeamSubscription":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateTeamSubscription", logger))},
		"UpdateUser":                   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateUser", logger))},
		"UpdateUserProfile":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateUserProfile", logger))},
		"UpdateUserSubscription":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateUserSubscription", logger))},
	}
	return options
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"AddAdvisorToTeam":             {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "AddAdvisorToTeam", logger))},
		"AddGroupMember":               {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "AddGroupMember", logger))},
		"AddMemberToTeam":              {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "AddMemberToTeam", logger))},
		"CreatTeamProfile":             {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreatTeamProfile", logger))},
		"CreateGroup":                  {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateGroup", logger))},
		"CreateProfile":                {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateProfile", logger))},
		"CreateSubscription":           {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateSubscription", logger))},
		"CreateSubscriptions":          {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateSubscriptions", logger))},
		"CreateTeam":                   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateTeam", logger))},
		"CreateTeamSubscription":       {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateTeamSubscription", logger))},
		"CreateUser":                   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateUser", logger))},
		"DeleteGroup":                  {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteGroup", logger))},
		"DeleteGroupMember":            {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteGroupMember", logger))},
		"DeleteSubscription":           {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteSubscription", logger))},
		"DeleteTeam":                   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteTeam", logger))},
		"DeleteTeamAdmin":              {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteTeamAdmin", logger))},
		"DeleteTeamAdvisor":            {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteTeamAdvisor", logger))},
		"DeleteTeamMember":             {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteTeamMember", logger))},
		"DeleteTeamProfile":            {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteTeamProfile", logger))},
		"DeleteTeamSubscription":       {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteTeamSubscription", logger))},
		"DeleteUser":                   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteUser", logger))},
		"DeleteUserProfile":            {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteUserProfile", logger))},
		"GetGroup":                     {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetGroup", logger))},
		"GetGroups":                    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetGroups", logger))},
		"GetGroupsByNumMembers":        {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetGroupsByNumMembers", logger))},
		"GetGroupsByTags":              {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetGroupsByTags", logger))},
		"GetGroupsByType":              {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetGroupsByType", logger))},
		"GetTeam":                      {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeam", logger))},
		"GetTeamAdmin":                 {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamAdmin", logger))},
		"GetTeamAdvisors":              {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamAdvisors", logger))},
		"GetTeamMembers":               {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamMembers", logger))},
		"GetTeamProfile":               {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamProfile", logger))},
		"GetTeamProfiles":              {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamProfiles", logger))},
		"GetTeamSubscriptions":         {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamSubscriptions", logger))},
		"GetTeams":                     {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeams", logger))},
		"GetTeamsByIndustry":           {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamsByIndustry", logger))},
		"GetTeamsByNumberOfEmployees":  {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamsByNumberOfEmployees", logger))},
		"GetTeamsByTags":               {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamsByTags", logger))},
		"GetTeamsByType":               {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTeamsByType", logger))},
		"GetUser":                      {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUser", logger))},
		"GetUserProfile":               {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUserProfile", logger))},
		"GetUserProfiles":              {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUserProfiles", logger))},
		"GetUserProfilesByNationality": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUserProfilesByNationality", logger))},
		"GetUserProfilesByType":        {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUserProfilesByType", logger))},
		"GetUserSubscriptions":         {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUserSubscriptions", logger))},
		"GetUsers":                     {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUsers", logger))},
		"GetUsersByAccountType":        {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUsersByAccountType", logger))},
		"GetUsersByIntent":             {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUsersByIntent", logger))},
		"UpdateGroup":                  {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateGroup", logger))},
		"UpdateGroupMember":            {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateGroupMember", logger))},
		"UpdateTeam":                   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTeam", logger))},
		"UpdateTeamAdmin":              {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTeamAdmin", logger))},
		"UpdateTeamAdvisor":            {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTeamAdvisor", logger))},
		"UpdateTeamMember":             {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTeamMember", logger))},
		"UpdateTeamProfile":            {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTeamProfile", logger))},
		"UpdateTeamSubscription":       {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTeamSubscription", logger))},
		"UpdateUser":                   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateUser", logger))},
		"UpdateUserProfile":            {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateUserProfile", logger))},
		"UpdateUserSubscription":       {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateUserSubscription", logger))},
	}
	return options
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"CreateUser", "CreateProfile", "CreateSubscription", "CreateSubscriptions", "UpdateUser", "UpdateUserProfile", "UpdateUserSubscription", "DeleteUser", "DeleteUserProfile", "DeleteSubscription", "GetUser", "GetUsers", "GetUsersByAccountType", "GetUsersByIntent", "GetUserProfile", "GetUserProfiles", "GetUserProfilesByType", "GetUserProfilesByNationality", "GetUserSubscriptions", "CreateTeam", "CreatTeamProfile", "CreateTeamSubscription", "AddMemberToTeam", "AddAdvisorToTeam", "UpdateTeam", "UpdateTeamAdmin", "UpdateTeamMember", "UpdateTeamAdvisor", "UpdateTeamProfile", "UpdateTeamSubscription", "DeleteTeam", "DeleteTeamProfile", "DeleteTeamMember", "DeleteTeamAdvisor", "DeleteTeamAdmin", "DeleteTeamSubscription", "GetTeam", "GetTeamProfile", "GetTeamSubscriptions", "GetTeamMembers", "GetTeamAdvisors", "GetTeamAdmin", "GetTeams", "GetTeamsByType", "GetTeamsByIndustry", "GetTeamsByNumberOfEmployees", "GetTeamProfiles", "GetTeamsByTags", "CreateGroup", "AddGroupMember", "UpdateGroup", "UpdateGroupMember", "DeleteGroup", "DeleteGroupMember", "GetGroup", "GetGroups", "GetGroupsByType", "GetGroupsByNumMembers", "GetGroupsByTags"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
