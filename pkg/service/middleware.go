package service

// Middleware describes a service middleware.
type Middleware func(MicroService) MicroService
