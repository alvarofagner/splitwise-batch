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

package converter

// func ToExpenseByShare(expense entity.Expense, share entity.Share) (interface{}, error) {
// 	if !expense.IsValid() {
// 		return nil, fmt.Errorf("expense is invalid")
// 	}

// 	var fields []reflect.StructField
// 	s := reflect.ValueOf(expense).Elem()
// 	typeOfT := s.Type()

// 	for i := 0; i < s.NumField(); i++ {
// 		field := s.Field(i)
// 		name := typeOfT.Field(i).Name
// 		fieldType := field.Type()
// 		tag := typeOfT.Field(i).Tag

// 		fmt.Println("Name:", name)
// 		fmt.Println("Type:", fieldType)
// 		fmt.Println("Tag:", tag)
// 		fields = append(fields, reflect.StructField{
// 			Name: name,
// 			Type: fieldType,
// 			Tag:  tag,
// 		})
// 	}

// 	typ := reflect.StructOf(fields)

// 	instance := reflect.New(typ).Elem()
// 	instance.Field(0).SetString("23.20")

// 	fmt.Println("Instance:", instance)

// 	return

// 	// return &splitwise.ExpenseByShare{
// 	// 	Expense:    splitwise.Expense(expense),
// 	// 	PaidUserID: share.PaidUserID,
// 	// 	OwedUserID: share.OwedUserID,
// 	// 	PaidShare:  share.PaidShare,
// 	// 	OwedShare:  share.OwedShare,
// 	// }, nil
// }
