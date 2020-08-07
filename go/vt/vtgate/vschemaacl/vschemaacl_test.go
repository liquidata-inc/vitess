/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vschemaacl

import (
	"testing"

	querypb "liquidata-inc/vitess/go/vt/proto/query"
)

func TestVschemaAcl(t *testing.T) {
	redUser := querypb.VTGateCallerID{Username: "redUser"}
	yellowUser := querypb.VTGateCallerID{Username: "yellowUser"}

	// By default no users are allowed in
	if Authorized(&redUser) {
		t.Errorf("user should not be authorized")
	}
	if Authorized(&yellowUser) {
		t.Errorf("user should not be authorized")
	}

	// Test wildcard
	*AuthorizedDDLUsers = "%"
	Init()

	if !Authorized(&redUser) {
		t.Errorf("user should be authorized")
	}
	if !Authorized(&yellowUser) {
		t.Errorf("user should be authorized")
	}

	// Test user list
	*AuthorizedDDLUsers = "oneUser, twoUser, redUser, blueUser"
	Init()

	if !Authorized(&redUser) {
		t.Errorf("user should be authorized")
	}
	if Authorized(&yellowUser) {
		t.Errorf("user should not be authorized")
	}

	// Revert to baseline state for other tests
	*AuthorizedDDLUsers = ""
	Init()

	// By default no users are allowed in
	if Authorized(&redUser) {
		t.Errorf("user should not be authorized")
	}
	if Authorized(&yellowUser) {
		t.Errorf("user should not be authorized")
	}
}
