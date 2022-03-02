// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	repo "aah-form-based-auth/app/repository"

	aah "aahframe.work"
	"aahframe.work/config"
	"aahframe.work/security/authc"
)

var _ authc.Authenticator = (*FormAuthenticationProvider)(nil)

// FormAuthenticationProvider struct implements `authc.Authenticator` interface.
type FormAuthenticationProvider struct {
}

// Init method initializes the FormAuthenticationProvider, this method gets called
// during server start up.
func (fa *FormAuthenticationProvider) Init(cfg *config.Config) error {

	// NOTE: Init is called on application startup

	return nil
}

// GetAuthenticationInfo method is `authc.Authenticator` interface
func (fa *FormAuthenticationProvider) GetAuthenticationInfo(authcToken *authc.AuthenticationToken) (*authc.AuthenticationInfo, error) {
	user, err := repo.R().GetUserByEmail(authcToken.Identity)
	if user == nil || err != nil {
		// No subject exists, return nil and error
		return nil, authc.ErrSubjectNotExists
	}

	// User found, now create authentication info and return to the framework
	authcInfo := authc.NewAuthenticationInfo()
	authcInfo.Principals = append(authcInfo.Principals,
		&authc.Principal{
			Value:     user.Email,
			IsPrimary: true,
			Realm:     "database",
		})
	authcInfo.Credential = []byte(user.Password)
	authcInfo.IsLocked = *(user.IsLocked.Ptr())
	authcInfo.IsExpired = *(user.IsExpried.Ptr())

	return authcInfo, nil
}

func PostAuthEvent(e *aah.Event) {
	ctx := e.Data.(*aah.Context)

	// Populate session info after authentication
	user, _ := repo.R().GetUserByEmail(ctx.Subject().PrimaryPrincipal().Value)
	ctx.Session().Set("FirstName", user.FirstName)
	ctx.Session().Set("LastName", user.LastName)
	ctx.Session().Set("Email", user.Email)

	ctx.Session().Set("Roles", ctx.Subject().AuthorizationInfo.Roles())
	ctx.Session().Set("Perms", ctx.Subject().AuthorizationInfo.Permissions())

}
