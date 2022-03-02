// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"aah-form-based-auth/app/database"
	"aah-form-based-auth/app/models"
	"aah-form-based-auth/app/repository"
	"fmt"
	"strings"

	aah "aahframe.work"
)

var repo repository.Repository

// AppController struct application controller
type AppController struct {
	*aah.Context
}

func (c *AppController) Before() {
	repo, _ = repository.NewSormRepository(database.SDB())
}

// Index method is application home page.
func (c *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Welcome to aah framework - Form Based Auth Example",
		},
	}

	c.Reply().Ok().HTML(data)
}

// BeforeLogin method action is interceptor of Login.
func (c *AppController) BeforeLogin() {
	if c.Subject().IsAuthenticated() {
		c.Reply().Redirect(c.RouteURL("index"))
		c.Abort()
	}
}

// Login method presents the login page.
func (c *AppController) Login() {
	c.Reply().Ok()
}

// Logout method does logout currently logged in subject (aka user).
func (c *AppController) Logout() {
	c.Subject().Logout()

	// Send it to login page or whatever the page you have to send the user
	// after logout
	c.Reply().Redirect(c.RouteURL("login"))
}

// BeforeManageUsers method is action interceptor of ManageUsers.
// func (c *AppController) BeforeManageUsers() {
// 	// Checking roles and permissions
// 	if !c.Subject().HasAnyRole("manager", "administrator") ||
// 		!c.Subject().IsPermitted("users:manage:view") {
// 		c.Reply().Forbidden().HTMLf("/access-denied.html", nil)
// 		c.Abort()
// 	}
// }

// BeforeManageUsers method is action interceptor of ManageUsers.
func (c *AppController) BeforeManageUsers() {
	// Checking roles and permissions
	if c.Subject().HasAnyRole("manager", "administrator") ||
		c.Subject().IsPermitted("users:manage:view") {
		ul, err := repo.GetAllUsers()
		if err != nil {
			c.Reply().InternalServerError().HTML(aah.Data{
				"msg": "internal server error",
			})
		}

		res := make(userListResponse, 0)
		for _, v := range ul {
			if v.Email != "admin@aahframework.org" && v != nil {
				res = append(res, formatUserInfo(v))
			}
		}
		c.AddViewArg("userlist", res)
	}
}

// ManageUsers method presents the manage user page afer verifying
// Authorization
func (c *AppController) ManageUsers() {
	// looks okay, present the page
	c.Reply().Ok().HTML(nil)
}

// EditUsers method get the managed user afer verifying
// Authorization
func (c *AppController) EditUsers(email string) {
	u, err := repo.GetUserByEmail(email)
	if err != nil {
		c.Reply().InternalServerError().HTML(aah.Data{
			"msg": "cann't find the user",
		})
	}
	// looks okay, present the page
	c.Reply().Ok().HTML(aah.Data{
		"form": formatUserInfo(u),
	})
}

type userResponse struct {
	FirstName   string
	LastName    string
	Email       string
	IsLocked    bool
	Roles       string
	Permissions string
}

type userListResponse []*userResponse

func formatUserInfo(u *models.User) *userResponse {
	data := new(userResponse)
	data.FirstName = u.FirstName
	data.LastName = u.LastName
	data.Email = u.Email
	data.IsLocked = u.IsLocked.Bool

	r := []string{}
	for _, v := range u.R.Roles {
		r = append(r, v.Name.String)
	}
	data.Roles = strings.Join(r, ", ")

	p := []string{}
	for _, v := range u.R.Permissions {
		p = append(p, v.Name.String)
	}
	data.Permissions = strings.Join(p, ", ")

	return data
}

type UserUpdateRequest struct {
	FirstName   string `bind:"firstName"`
	LastName    string `bind:"lastName"`
	Email       string `bind:"email"`
	IsLocked    bool   `bind:"isLocked"`
	Roles       string `bind:"roles"`
	Permissions string `bind:"permissions"`
}

// UpdateUsers method update the managed user afer verifying Authorization
func (c *AppController) UpdateUsers(email string, req *UserUpdateRequest) {
	c.Log().Debugf("req = %v\n", req)
	exist, _ := repo.Exist(email)
	if !exist {
		c.Reply().InternalServerError().HTML(aah.Data{
			"msg": "User not found",
		})
	}

	args := repository.UpdateUserArgs{}
	if req.FirstName != "" {
		args.FirstName = req.FirstName
	}
	if req.LastName != "" {
		args.LastName = req.LastName
	}
	if req.Email != "" {
		args.Email = req.Email
	}
	args.IsLocked = req.IsLocked

	if req.Roles != "" {
		args.Roles = strings.Split(strings.TrimSpace(req.Roles), ",")
		c.Log().Debugf("args.Roles = %#v\n", args.Roles)
	} else {
		args.Roles = []string{}
	}
	if req.Permissions != "" {
		args.Permissions = []string{req.Permissions}
		c.Log().Debugf("args.Permissions = %#v\n", args.Permissions)
	} else {
		args.Permissions = []string{}
	}

	err := repo.UpdateUser(email, args)
	if err != nil {
		c.Reply().InternalServerError().HTML(aah.Data{
			"err": "internal server error",
		})
	}

	// looks okay
	c.Reply().Ok().RedirectWithStatus(fmt.Sprintf("/edit/users.html/%s", email), 302)
}
