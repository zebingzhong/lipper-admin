package services

import "github.com/zebingzhong/lipper-admin/internal/services/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
