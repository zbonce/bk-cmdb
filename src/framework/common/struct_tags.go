/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"configcenter/src/framework/core/types"
	"fmt"
	"reflect"
)

// GetTags parse a object and get the all tags
func GetTags(target interface{}) []string {

	targetType := reflect.TypeOf(target)
	switch targetType.Kind() {
	default:
		break
	case reflect.Ptr:
		fmt.Printf("hello")
		targetType = targetType.Elem()

	}

	numField := targetType.NumField()
	tags := make([]string, 0)
	for i := 0; i < numField; i++ {
		structField := targetType.Field(i)
		if tag, ok := structField.Tag.Lookup("field"); ok {
			tags = append(tags, tag)
		}
	}
	return tags

}

// SetValueToStructByTags set the struct object field value by tags
func SetValueToStructByTags(target interface{}, values types.MapStr) error {

	targetType := reflect.TypeOf(target)
	targetValue := reflect.ValueOf(target)
	switch targetType.Kind() {
	case reflect.Ptr:
		targetType = targetType.Elem()
		targetValue = targetValue.Elem()
	}

	numField := targetType.NumField()
	for i := 0; i < numField; i++ {
		structField := targetType.Field(i)
		tag, ok := structField.Tag.Lookup("field")
		if !ok {
			continue
		}

		tagVal, ok := values[tag]
		if !ok {
			continue
		}

		fieldValue := targetValue.FieldByName(structField.Name)
		if !fieldValue.CanSet() {
			continue
		}

		switch structField.Type.Kind() {
		default:
			return fmt.Errorf("unsuport the type %v", structField.Type.Kind())
		case reflect.Bool:
			fieldValue.SetBool(tagVal.(bool))
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
			switch t := tagVal.(type) {
			default:
				return fmt.Errorf("unsuport the type ,value is (%#v)", tagVal)
			case int:
				fieldValue.SetInt(int64(t))
			case int16:
				fieldValue.SetInt(int64(t))
			case int32:
				fieldValue.SetInt(int64(t))
			case int64:
				fieldValue.SetInt(int64(t))
			case int8:
				fieldValue.SetInt(int64(t))
			case uint:
				fieldValue.SetInt(int64(t))
			case uint16:
				fieldValue.SetInt(int64(t))
			case uint32:
				fieldValue.SetInt(int64(t))
			case uint64:
				fieldValue.SetInt(int64(t))
			case uint8:
				fieldValue.SetInt(int64(t))
			}

		case reflect.Float32, reflect.Float64:
			switch t := tagVal.(type) {
			default:
				return fmt.Errorf("unsuport the type ,value is (%#v)", tagVal)
			case float32:
				fieldValue.SetFloat(float64(t))
			case float64:
				fieldValue.SetFloat(float64(t))
			}

		case reflect.String:
			switch t := tagVal.(type) {
			default:
				return fmt.Errorf("unsuport the type ,value is (%#v)", tagVal)
			case string:
				fieldValue.SetString(t)
			}

		}

	}

	return nil
}
