// Copyright Jeevanandam M. (https://github.com/jeevatkm, jeeva@myjeeva.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	aah "aahframe.work"
	"aahframe.work/security"
)

// Copy From https://github.com/go-aah/aah/blob/edge/view.go

//
// Security view functions
//

// TmplIsAuthenticated method returns the value of `Session.IsAuthenticated`.
func TmplIsAuthenticated(viewArgs map[string]interface{}) bool {
	if sub := getSubjectFromViewArgs(viewArgs); sub != nil {
		if sub.Session != nil {
			return sub.Session.IsAuthenticated
		}
	}
	return false
}

// TmplHasRole method returns the value of `Subject.HasRole`.
func TmplHasRole(viewArgs map[string]interface{}, role string) bool {
	if sub := getSubjectFromViewArgs(viewArgs); sub != nil {
		return sub.HasRole(role)
	}
	return false
}

// TmplHasAllRoles method returns the value of `Subject.HasAllRoles`.
func TmplHasAllRoles(viewArgs map[string]interface{}, roles ...string) bool {
	if sub := getSubjectFromViewArgs(viewArgs); sub != nil {
		return sub.HasAllRoles(roles...)
	}
	return false
}

// TmplHasAnyRole method returns the value of `Subject.HasAnyRole`.
func TmplHasAnyRole(viewArgs map[string]interface{}, roles ...string) bool {
	if sub := getSubjectFromViewArgs(viewArgs); sub != nil {
		return sub.HasAnyRole(roles...)
	}
	return false
}

// TmplIsPermitted method returns the value of `Subject.IsPermitted`.
func TmplIsPermitted(viewArgs map[string]interface{}, permission string) bool {
	if sub := getSubjectFromViewArgs(viewArgs); sub != nil {
		return sub.IsPermitted(permission)
	}
	return false
}

// TmplIsPermittedAll method returns the value of `Subject.IsPermittedAll`.
func TmplIsPermittedAll(viewArgs map[string]interface{}, permissions ...string) bool {
	if sub := getSubjectFromViewArgs(viewArgs); sub != nil {
		return sub.IsPermittedAll(permissions...)
	}
	return false
}

func getSubjectFromViewArgs(viewArgs map[string]interface{}) *security.Subject {
	if sv, found := viewArgs[aah.KeyViewArgSubject]; found {
		return sv.(*security.Subject)
	}
	return nil
}
