// Copyright 2022 vscode
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package security

import (
	"encoding/json"
	"io/ioutil"

	"github.com/anvari1313/splitwise.go"
)

type credentials struct {
	ApiKey string `json:"apiKey"`
}

func Authenticate(credentialsFile string) (splitwise.AuthProvider, error) {
	f, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	creds := &credentials{}
	if err := json.Unmarshal([]byte(f), creds); err != nil {
		return nil, err
	}

	auth := splitwise.NewAPIKeyAuth(creds.ApiKey)
	return auth, nil
}
