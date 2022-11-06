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

package entity

type Expense struct {
	Cost           string
	Description    string
	Details        string
	Date           string
	RepeatInterval string
	CurrencyCode   string
	CategoryId     uint32
	GroupId        uint32
}

// Checks if expense has all mandatory fields
func (e *Expense) IsValid() bool {
	return e.Cost != "" && e.Description != "" && e.GroupId != 0
}

type Share struct {
	UserID uint64
	Share  float64
}
