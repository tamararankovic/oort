package checkerpb

import (
	"github.com/c12s/oort/domain/checker"
	"github.com/c12s/oort/domain/model"
	"log"
)

func (x *CheckPermissionReq) MapToDomain() checker.CheckPermissionReq {
	envAttributes := make([]model.Attribute, len(x.EnvAttributes))
	for i, attr := range x.EnvAttributes {
		domainAttr, err := attr.MapToDomain()
		if err != nil {
			log.Println(err)
			continue
		}
		envAttributes[i] = domainAttr
	}
	return checker.CheckPermissionReq{
		Principal:      x.Principal.MapToDomain(),
		Resource:       x.Resource.MapToDomain(),
		PermissionName: x.PermissionName,
		Env:            envAttributes,
	}
}
