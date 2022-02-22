// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	repo "aah-form-based-auth/app/repository"

	"aahframe.work/config"
	"aahframe.work/security/authc"
	"aahframe.work/security/authz"
)

var _ authz.Authorizer = (*FormAuthorizationProvider)(nil)

// FormAuthorizationProvider struct implements `authz.Authorizer` interface.
type FormAuthorizationProvider struct {
}

// Init method initializes the FormAuthoriationProvider, this method gets called
// during server start up.
func (fa *FormAuthorizationProvider) Init(cfg *config.Config) error {

	// NOTE: Init is called on application startup

	return nil
}

// GetAuthorizationInfo method is `authz.Authorizer` interface.
//
// GetAuthorizationInfo method gets called after authentication is successful
// to get Subject's (aka User) access control information such as roles and permissions.
func (fa *FormAuthorizationProvider) GetAuthorizationInfo(authcInfo *authc.AuthenticationInfo) *authz.AuthorizationInfo {
	authorities, _ := repo.R().GetUserByEmail(authcInfo.PrimaryPrincipal().Value)

	r := []string{}
	for _, v := range authorities.R.Roles {
		r = append(r, v.Name.String)
	}
	p := []string{}
	for _, v := range authorities.R.Permissions {
		p = append(p, v.Name.String)
	}

	authzInfo := authz.NewAuthorizationInfo()
	authzInfo.AddRole(r...)
	authzInfo.AddPermissionString(p...)

	return authzInfo
}
