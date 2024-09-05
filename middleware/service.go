package middleware

import svc "github.com/shivdaskadam/golang-boilerplate/iface"

// ServiceMiddleware used to chain behaviors on the UserService using middleware pattern

type ServiceMiddleware func(svc.Service) svc.Service
