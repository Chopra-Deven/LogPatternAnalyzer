/*
 * Copyright (c) Motadata 2024.  All rights reserved.
 */

package utils

import (
	"encoding/json"
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type (
	MotadataString string

	MotadataINT int

	MotadataUINT uint

	MotadataUINT8 uint8

	MotadataUINT16 uint16

	MotadataUINT32 uint32

	MotadataUINT64 uint64

	MotadataFloat32 float32

	MotadataFloat64 float64

	MotadataStringList []string

	MotadataIntList []int

	MotadataMap map[string]interface{}

	MotadataStringMap map[string]string

	MotadataIntMap map[int]string

	MotadataStringIntMap map[string]int

	MotadataIntFloatMap map[int]float64

	MotadataKB float64

	MotadataMB float64

	MotadataGB float64

	MotadataTime int

	MotadataTimeString string
)

//Convert to UINT Methods

func (value MotadataFloat64) ToUINT() MotadataUINT {

	return MotadataUINT(value)
}

func (value MotadataINT) ToUINT() MotadataUINT {

	return MotadataUINT(value)
}

//Convert to UINT Methods

func (value MotadataUINT) ToInt() int {

	return int(value)
}

//Convert to UINT8 Methods

func (value MotadataUINT8) ToInt() int {

	return int(value)
}

//Convert to UINT16 Methods

func (value MotadataFloat64) ToUINT16() MotadataUINT16 {

	return MotadataUINT16(value)
}

func (value MotadataUINT16) ToUInt16() uint16 {

	return uint16(value)
}

func (value MotadataUINT16) ToInt() int {

	return int(value)
}

func (value MotadataINT) ToInt64() int64 {

	return int64(value)
}

//Convert to string Methods

func (value MotadataINT) ToString() MotadataString {

	return MotadataString(strconv.Itoa(int(value)))
}

func (value MotadataINT) ToNativeString() string {

	return strconv.Itoa(int(value))
}

func (value MotadataUINT) ToString() MotadataString {

	return MotadataString(strconv.Itoa(int(value)))
}

func (value MotadataUINT8) ToString() MotadataString {

	return MotadataString(strconv.Itoa(int(value)))
}

func (value MotadataUINT16) ToString() MotadataString {

	return MotadataString(strconv.Itoa(int(value)))
}

func (value MotadataUINT32) ToString() MotadataString {

	return MotadataString(strconv.Itoa(int(value)))
}

func (value MotadataUINT64) ToString() MotadataString {

	return MotadataString(strconv.Itoa(int(value)))
}

func (value MotadataFloat32) ToString() MotadataString {

	return MotadataString(strconv.Itoa(int(value)))
}

func (value MotadataFloat64) ToString() MotadataString {

	return MotadataString(strconv.Itoa(int(value)))
}

func (value MotadataString) ToString() string {

	return strings.TrimSpace(string(value))
}

func (value MotadataTime) ToString() (result MotadataString) {

	days := value / 86400

	hours := (value - (days * 86400)) / 3600

	minutes := (value - (days * 86400) - (hours * 3600)) / 60

	value = value - (days * 86400) - (hours * 3600) - (minutes * 60)

	if days > 0 {

		result += MotadataString(fmt.Sprintf(" %d day%s", days, "s"))

	} else {

		result += MotadataString(fmt.Sprintf(" %d day%s", days, BlankString))

	}
	if hours > 0 {

		result += MotadataString(fmt.Sprintf(" %d hour%s", hours, "s"))

	} else {

		result += MotadataString(fmt.Sprintf(" %d hour%s", hours, BlankString))

	}
	if minutes > 0 {

		result += MotadataString(fmt.Sprintf(" %d minute%s", minutes, "s"))

	} else {

		result += MotadataString(fmt.Sprintf(" %d minute%s", minutes, BlankString))

	}
	if value > 0 {

		result += MotadataString(fmt.Sprintf(" %d second%s", value, "s"))

	} else {

		result += MotadataString(fmt.Sprintf(" %d second%s", value, BlankString))

	}

	return
}

// Convert to Int Methods

func (value MotadataUINT) ToINT() MotadataINT {

	return MotadataINT(value)
}

func (value MotadataUINT8) ToINT() MotadataINT {

	return MotadataINT(value)
}

func (value MotadataUINT16) ToINT() MotadataINT {

	return MotadataINT(value)
}

func (value MotadataUINT32) ToINT() MotadataINT {

	return MotadataINT(value)
}

func (value MotadataUINT64) ToINT() MotadataINT {

	return MotadataINT(value)
}

func (value MotadataFloat32) ToINT() MotadataINT {

	return MotadataINT(value)
}

func (value MotadataFloat64) ToINT() MotadataINT {

	return MotadataINT(value)
}

func (value MotadataString) ToINT() MotadataINT {

	result, _ := strconv.ParseInt(strings.TrimSpace(string(value)), 10, 64)

	return MotadataINT(result)

}

func (value MotadataString) ToInt() int {

	result, _ := strconv.ParseInt(strings.TrimSpace(string(value)), 10, 64)

	return int(result)

}

func (value MotadataINT) ToInt() int {

	return int(value)

}

//Convert to Float64 methods

func (value MotadataFloat64) ToFloat64() MotadataFloat64 {

	return MotadataFloat64(math.Round(float64(value*100)) / 100)
}

func (value MotadataINT) ToFloat64() MotadataFloat64 {

	return MotadataFloat64(math.Round(float64(value*100)) / 100)
}

func (value MotadataUINT) ToFloat64() MotadataFloat64 {

	return MotadataFloat64(math.Round(float64(value*100)) / 100)
}

func (value MotadataUINT8) ToFloat64() MotadataFloat64 {

	return MotadataFloat64(math.Round(float64(value*100)) / 100)
}

func (value MotadataUINT16) ToFloat64() MotadataFloat64 {

	return MotadataFloat64(math.Round(float64(value*100)) / 100)
}

func (value MotadataUINT32) ToFloat64() MotadataFloat64 {

	return MotadataFloat64(math.Round(float64(value*100)) / 100)
}

func (value MotadataUINT64) ToFloat64() MotadataFloat64 {

	return MotadataFloat64(math.Round(float64(value*100)) / 100)
}

func (value MotadataFloat32) ToFloat64() MotadataFloat64 {

	return MotadataFloat64(math.Round(float64(value*100)) / 100)
}

func (value MotadataString) ToFloat64() MotadataFloat64 {

	result, _ := strconv.ParseFloat(strings.TrimSpace(string(value)), 64)

	return MotadataFloat64(math.Round(result*100) / 100)
}

//Convert to Bytes

func (value MotadataMB) ToBytes() MotadataFloat64 {

	return MotadataFloat64(value * 1024 * 1024)
}

func (value MotadataGB) ToBytes() MotadataFloat64 {

	return MotadataFloat64(value * 1024 * 1024 * 1024)
}

func (value MotadataKB) ToBytes() MotadataFloat64 {

	return MotadataFloat64(value * 1024)
}

//Convert to json

func (value MotadataMap) ToJSON() MotadataString {

	result, _ := json.Marshal(value)

	return MotadataString(result)
}

//String Methods

func (value MotadataString) TrimSpace() MotadataString {

	return MotadataString(strings.TrimSpace(string(value)))
}

func (value MotadataString) ToLower() (result MotadataString) {

	result = MotadataString(strings.ToLower(value.ToString()))

	return
}

func (value MotadataString) ToUpper() (result MotadataString) {

	result = MotadataString(strings.ToUpper(value.ToString()))

	return
}

func (value MotadataString) ToTitle() (result MotadataString) {

	result = MotadataString(strings.ToTitle(value.ToString()))

	return
}

func (value MotadataString) ToLowerNative() (result string) {

	result = strings.ToLower(value.ToString())

	return
}

func (value MotadataString) Fields() (result []MotadataString) {

	tokens := strings.Fields(value.ToString())

	for index := range tokens {

		result = append(result, MotadataString(strings.TrimSpace(tokens[index])))
	}

	return
}

func (value MotadataStringList) Join(delimiter string) (result MotadataString) {

	result = MotadataString(strings.Join(value, delimiter))

	return
}

func (value MotadataString) Split(delimiter string) (result []MotadataString) {

	tokens := strings.Split(value.ToString(), delimiter)

	for index := range tokens {

		if tokens[index] != BlankString && tokens[index] != NewLineRegexPattern {

			result = append(result, MotadataString(strings.TrimSpace(tokens[index])))

		}
	}

	return
}

func (value MotadataString) SplitNWithEmptyEntries(delimiter string, n int) (values []MotadataString) {

	tokens := strings.SplitN(value.ToString(), delimiter, n)

	for index := range tokens {

		values = append(values, MotadataString(strings.TrimSpace(tokens[index])))

	}

	return
}

func (value MotadataString) TrimSuffix(suffix string) MotadataString {

	return MotadataString(strings.TrimSuffix(value.ToString(), suffix))
}

func (value MotadataString) Trim(delimiter string) MotadataString {

	return MotadataString(strings.Trim(value.ToString(), delimiter))
}

func (value MotadataString) TrimRight(delimiter string) MotadataString {

	return MotadataString(strings.TrimRight(value.ToString(), delimiter))

}

func (value MotadataString) SplitWithEmptyEntries(delimiter string) (result []MotadataString) {

	tokens := strings.Split(value.ToString(), delimiter)

	for index := range tokens {

		result = append(result, MotadataString(strings.TrimSpace(tokens[index])))

	}

	return
}

func (value MotadataString) SplitStringExcludeQuotes(delimiter string) (values []MotadataString) {

	quotes := false

	tokens := strings.Split(value.ToString(), delimiter)

	for i, j := 0, -1; i < len(tokens); i++ {

		if quotes {

			values[j] = values[j] + "," + MotadataString(strings.TrimSpace(tokens[i]))

		} else {

			values = append(values, MotadataString(strings.TrimSpace(tokens[i])))

			j++

		}

		if strings.Count(tokens[i], `"`) == 1 {

			quotes = !quotes

		}

	}

	return
}

func (value MotadataString) ReplaceAll(old MotadataString, new MotadataString) MotadataString {

	return MotadataString(strings.ReplaceAll(ToString(value), ToString(old), ToString(new)))

}

func (value MotadataString) Replace(old MotadataString, new MotadataString, n int) MotadataString {

	return MotadataString(strings.Replace(value.ToString(), ToString(old), ToString(new), n))

}

func (value MotadataString) HasPrefix(prefix string) bool {

	return strings.HasPrefix(value.ToString(), prefix)

}

func (value MotadataString) HasSuffix(suffix string) bool {

	return strings.HasSuffix(value.ToString(), suffix)

}

func (value MotadataString) Strip() MotadataString {

	return MotadataString(strings.Trim(strings.TrimSpace(value.ToString()), NewLineSeparator))
}

func (value MotadataString) SplitBySpace() MotadataStringList {

	result := MotadataStringList{}

	tokens := strings.Split(string(value), SpaceSeparator)

	for index := range tokens {

		if tokens[index] != BlankString {

			result = append(result, strings.TrimSpace(tokens[index]))
		}
	}

	return result
}

func (value MotadataString) SplitN(delimiter string, n int) (result []MotadataString) {

	tokens := strings.SplitN(value.ToString(), delimiter, n)

	for index := range tokens {

		if tokens[index] != BlankString {

			result = append(result, MotadataString(strings.TrimSpace(tokens[index])))

		}
	}

	return
}

func (elements MotadataStringList) Contains(value MotadataString) bool {

	for _, element := range elements {

		if MotadataString(element) == value {

			return true
		}
	}
	return false
}

func (value MotadataStringList) IsNotEmpty() bool {

	if value != nil && len(value) > 0 {

		return true
	}

	return false
}

func (value MotadataString) IsDigit() bool {

	if _, err := strconv.Atoi(strings.Replace(string(value), ".", BlankString, 1)); err == nil {

		return true
	}

	return false
}

func (value MotadataMap) IsNotEmpty() bool {

	if value != nil && len(value) > 0 {

		return true
	}

	return false
}

func (value MotadataMap) GetStringValues() []MotadataString {

	var values []MotadataString

	for _, val := range value {

		values = append(values, val.(MotadataString))

	}

	return values

}

func (value MotadataStringMap) IsNotEmpty() bool {

	if value != nil && len(value) > 0 {

		return true
	}

	return false
}

func (value MotadataString) IsNotEmpty() bool {

	if value != BlankString && len(value.TrimSpace()) > 0 {

		return true
	}

	return false
}

func (value MotadataString) Contains(substr string) bool {

	return strings.Contains(strings.TrimSpace(value.ToString()), substr)
}

func IsNotEmpty(values []interface{}) bool {

	if values != nil && len(values) > 0 {

		return true
	}

	return false
}

func IsNotEmptyMapSlice(values []MotadataMap) bool {

	if values != nil && len(values) > 0 {

		return true
	}

	return false
}

func IsNotEmptyStringSlice(values []MotadataString) bool {

	if values != nil && len(values) > 0 {

		return true
	}

	return false
}

func (context MotadataMap) Contains(key string) (result bool) {

	if _, found := context[key]; found {

		return true

	}

	return
}

func (context MotadataStringMap) Contains(key string) (result bool) {

	if _, found := context[key]; found {

		return true

	}

	return
}

func (context MotadataIntMap) Contains(key int) (result bool) {

	if _, found := context[key]; found {

		return true

	}

	return
}

func (context MotadataIntMap) GetIntMapValues() (result MotadataStringList) {

	for _, value := range context {

		result = append(result, value)
	}

	return
}

func (context MotadataMap) Delete(key string) {

	if context.Contains(key) {

		delete(context, key)
	}
}

func (context MotadataStringMap) Delete(key string) {

	if context.Contains(key) {

		delete(context, key)
	}

}
func (context MotadataMap) Merge(result MotadataMap) MotadataMap {

	for key, value := range result {

		context[key] = value
	}

	return context
}

func (context MotadataMap) Copy() (result MotadataMap) {

	result = make(MotadataMap)

	for key, value := range context {

		result[key] = value
	}

	return
}

func (context MotadataMap) DeepCopy() MotadataMap {

	result := make(MotadataMap)

	for key, value := range context {

		if reflect.TypeOf(value).String() == "[]interface {}" {

			var values []interface{}

			for _, value := range value.([]interface{}) {

				values = append(values, ToMap(value).Copy())

			}

			result[key] = values

		} else {

			result[key] = value
		}

	}

	return result
}

func DeepCopy(context cmap.ConcurrentMap[string, interface{}]) cmap.ConcurrentMap[string, interface{}] {

	result := cmap.New[interface{}]()

	for key, value := range context.Items() {

		if reflect.TypeOf(value).String() == "[]interface {}" {

			var values []interface{}

			for _, value := range value.([]interface{}) {

				values = append(values, ToMap(value).Copy())

			}

			result.Set(key, value)

		} else {

			result.Set(key, value)
		}

	}

	return result
}

func (context MotadataMap) GetMapKeyByValue(value string) string {

	for key, val := range context {

		if value == val {

			return key
		}
	}

	return ""
}

func (context MotadataMap) GetKeys() MotadataStringList {

	result := MotadataStringList{}

	if context.IsNotEmpty() {

		for key := range context {

			result = append(result, key)
		}
	}

	return result
}

func (context MotadataMap) ToStringMap() MotadataStringMap {

	result := make(MotadataStringMap)

	for key, value := range context {

		result[fmt.Sprintf("%v", key)] = fmt.Sprintf("%v", value)
	}

	return result
}

func (context MotadataMap) ToMap() MotadataMap {

	for key, value := range context {

		context[key] = MotadataMap(value.(map[string]interface{}))
	}

	return context
}

//For Converting Date and Time

func (value MotadataTime) Format() MotadataString {

	return MotadataString(time.Unix(int64(value), 0).Format(TimeFormat))

}

func (value MotadataTimeString) Format() MotadataString {

	return MotadataString(time.Now().Format(string(value)))

}

func (value MotadataString) MatchFound(patterns ...MotadataString) (result bool) {

	for _, pattern := range patterns {

		if reg, _ := regexp.MatchString(string(pattern), string(value)); reg {

			return true

		}
	}

	return
}

func (context MotadataMap) GetINTValue(key string) (result MotadataINT) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = value.(MotadataINT)

		} else if reflect.TypeOf(value).Name() == "uint" {

			result = MotadataINT(value.(uint))

		} else if reflect.TypeOf(value).Name() == "uint8" {

			result = MotadataINT(value.(uint8))

		} else if reflect.TypeOf(value).Name() == "uint16" {

			result = MotadataINT(value.(uint16))

		} else if reflect.TypeOf(value).Name() == "uint32" {

			result = MotadataINT(value.(uint32))

		} else if reflect.TypeOf(value).Name() == "uint64" {

			result = MotadataINT(value.(uint64))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataINT(value.(int))

		} else if reflect.TypeOf(value).Name() == "int8" {

			result = MotadataINT(value.(int8))

		} else if reflect.TypeOf(value).Name() == "int16" {

			result = MotadataINT(value.(int16))

		} else if reflect.TypeOf(value).Name() == "int32" {

			result = MotadataINT(value.(int32))

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = MotadataINT(value.(int64))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataINT(value.(float64))

		} else if reflect.TypeOf(value).Name() == "MotadataFloat64" {

			result = MotadataINT(value.(MotadataFloat64))
		}
	}

	return
}

func (context MotadataMap) GetUINTValue(key string) (result MotadataUINT) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "uint" {

			result = MotadataUINT(value.(uint))

		} else if reflect.TypeOf(value).Name() == "uint8" {

			result = MotadataUINT(value.(uint8))

		} else if reflect.TypeOf(value).Name() == "uint16" {

			result = MotadataUINT(value.(uint16))

		} else if reflect.TypeOf(value).Name() == "uint32" {

			result = MotadataUINT(value.(uint32))

		} else if reflect.TypeOf(value).Name() == "uint64" {

			result = MotadataUINT(value.(uint64))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataUINT(value.(int))

		} else if reflect.TypeOf(value).Name() == "int8" {

			result = MotadataUINT(value.(int8))

		} else if reflect.TypeOf(value).Name() == "int16" {

			result = MotadataUINT(value.(int16))

		} else if reflect.TypeOf(value).Name() == "int32" {

			result = MotadataUINT(value.(int32))

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = MotadataUINT(value.(int64))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataUINT(value.(float64))

		} else if reflect.TypeOf(value).Name() == "MotadataUINT" {

			result = value.(MotadataUINT)
		}
	}

	return
}

func GetUINTValue(context cmap.ConcurrentMap[string, interface{}], key string) (result MotadataUINT) {

	if value, found := context.Get(key); found {

		if reflect.TypeOf(value).Name() == "uint" {

			result = MotadataUINT(value.(uint))

		} else if reflect.TypeOf(value).Name() == "uint8" {

			result = MotadataUINT(value.(uint8))

		} else if reflect.TypeOf(value).Name() == "uint16" {

			result = MotadataUINT(value.(uint16))

		} else if reflect.TypeOf(value).Name() == "uint32" {

			result = MotadataUINT(value.(uint32))

		} else if reflect.TypeOf(value).Name() == "uint64" {

			result = MotadataUINT(value.(uint64))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataUINT(value.(int))

		} else if reflect.TypeOf(value).Name() == "int8" {

			result = MotadataUINT(value.(int8))

		} else if reflect.TypeOf(value).Name() == "int16" {

			result = MotadataUINT(value.(int16))

		} else if reflect.TypeOf(value).Name() == "int32" {

			result = MotadataUINT(value.(int32))

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = MotadataUINT(value.(int64))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataUINT(value.(float64))

		} else if reflect.TypeOf(value).Name() == "MotadataUINT" {

			result = value.(MotadataUINT)
		}
	}

	return
}

func (context MotadataMap) GetUINT8Value(key string) (result MotadataUINT8) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataUINT8(value.(float64))
		} else if reflect.TypeOf(context[key]).Name() == "int" {

			result = MotadataUINT8(context[key].(int))
		}
	}

	return
}

func (context MotadataMap) GetUINT16Value(key string) (result MotadataUINT16) {

	if context.Contains(key) {

		if reflect.TypeOf(context[key]).Name() == "float64" {

			result = MotadataUINT16(context[key].(float64))
		} else if reflect.TypeOf(context[key]).Name() == "int" {

			result = MotadataUINT16(context[key].(int))
		}
	}

	return
}

func GetUINT16Value(context cmap.ConcurrentMap[string, interface{}], key string) (result MotadataUINT16) {

	if value, found := context.Get(key); found {

		if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataUINT16(value.(float64))
		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataUINT16(value.(int))
		}
	}

	return
}

func (context MotadataMap) GetMotadataStringValue(key string) (result MotadataString) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "MotadataString" {

			result = value.(MotadataString)

		} else if reflect.TypeOf(value).Name() == "string" {

			result = MotadataString(value.(string))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataString(strconv.Itoa(value.(int)))

		} else if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = MotadataString(strconv.Itoa(value.(MotadataINT).ToInt()))

		} else if reflect.TypeOf(value).String() == "[]uint8" {

			result = MotadataString(value.([]uint8))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataString(strconv.Itoa(int(value.(float64))))

		} else if reflect.TypeOf(value).String() == "*string" {

			result = MotadataString(*value.(*string))

		} else if reflect.TypeOf(value).Name() == "MotadataFloat64" {

			result = MotadataString(strconv.Itoa(int(value.(MotadataFloat64))))

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = MotadataString(strconv.Itoa(int(value.(int64))))

		}
	}

	return
}

func (context MotadataMap) GetFloat64Value(key string) (result MotadataFloat64) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "float64" {

			result = MotadataFloat64(value.(float64))

		} else if reflect.TypeOf(value).Name() == "MotadataFloat64" {

			result = value.(MotadataFloat64)

		} else if reflect.TypeOf(value).Name() == "int" {

			result = MotadataFloat64(value.(int))

		} else if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = MotadataFloat64(value.(MotadataINT))
		}
	}

	return
}

func (context MotadataMap) GetStringValue(key string) (result string) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "string" {

			result = value.(string)

		} else if reflect.TypeOf(value).Name() == "MotadataString" {

			result = string(value.(MotadataString))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = strconv.Itoa(value.(int))

		} else if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = strconv.Itoa(value.(MotadataINT).ToInt())

		} else if reflect.TypeOf(value).String() == "[]uint8" {

			result = string(value.([]uint8))

		} else if reflect.TypeOf(value).String() == "uint" {

			result = strconv.Itoa(ToInt(value.(uint)))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = strconv.Itoa(int(value.(float64)))

		} else if reflect.TypeOf(value).String() == "*string" {

			result = *value.(*string)

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = strconv.Itoa(int(value.(int64)))

		}
	}

	return
}

func GetStringValue(conMap cmap.ConcurrentMap[string, interface{}], key string) (result string) {

	if value, found := conMap.Get(key); found {

		if reflect.TypeOf(value).Name() == "string" {

			result = value.(string)

		} else if reflect.TypeOf(value).Name() == "MotadataString" {

			result = string(value.(MotadataString))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = strconv.Itoa(value.(int))

		} else if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = strconv.Itoa(value.(MotadataINT).ToInt())

		} else if reflect.TypeOf(value).String() == "[]uint8" {

			result = string(value.([]uint8))

		} else if reflect.TypeOf(value).String() == "uint" {

			result = strconv.Itoa(ToInt(value.(uint)))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = strconv.Itoa(int(value.(float64)))

		} else if reflect.TypeOf(value).String() == "*string" {

			result = *value.(*string)

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = strconv.Itoa(int(value.(int64)))

		}
	}

	return result
}

func (context MotadataMap) GetIntValue(key string) (result int) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "string" {

			output, _ := strconv.ParseInt(strings.TrimSpace(value.(string)), 10, 64)

			result = int(output)

		} else if reflect.TypeOf(value).Name() == "uint" {

			result = int(value.(uint))

		} else if reflect.TypeOf(value).Name() == "uint8" {

			result = int(value.(uint8))

		} else if reflect.TypeOf(value).Name() == "uint16" {

			result = int(value.(uint16))

		} else if reflect.TypeOf(value).Name() == "uint32" {

			result = int(value.(uint32))

		} else if reflect.TypeOf(value).Name() == "uint64" {

			result = int(value.(uint64))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = value.(int)

		} else if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = int(value.(MotadataINT))

		} else if reflect.TypeOf(value).Name() == "int8" {

			result = int(value.(int8))

		} else if reflect.TypeOf(value).Name() == "int16" {

			result = int(value.(int16))

		} else if reflect.TypeOf(value).Name() == "int32" {

			result = int(value.(int32))

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = int(value.(int64))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = int(value.(float64))
		}
	}

	return
}

func GetIntValue(conMap cmap.ConcurrentMap[string, interface{}], key string) (result int) {

	if value, found := conMap.Get(key); found {

		if reflect.TypeOf(value).Name() == "string" {

			output, _ := strconv.ParseInt(strings.TrimSpace(value.(string)), 10, 64)

			result = int(output)

		} else if reflect.TypeOf(value).Name() == "uint" {

			result = int(value.(uint))

		} else if reflect.TypeOf(value).Name() == "uint8" {

			result = int(value.(uint8))

		} else if reflect.TypeOf(value).Name() == "uint16" {

			result = int(value.(uint16))

		} else if reflect.TypeOf(value).Name() == "uint32" {

			result = int(value.(uint32))

		} else if reflect.TypeOf(value).Name() == "uint64" {

			result = int(value.(uint64))

		} else if reflect.TypeOf(value).Name() == "int" {

			result = value.(int)

		} else if reflect.TypeOf(value).Name() == "MotadataINT" {

			result = int(value.(MotadataINT))

		} else if reflect.TypeOf(value).Name() == "int8" {

			result = int(value.(int8))

		} else if reflect.TypeOf(value).Name() == "int16" {

			result = int(value.(int16))

		} else if reflect.TypeOf(value).Name() == "int32" {

			result = int(value.(int32))

		} else if reflect.TypeOf(value).Name() == "int64" {

			result = int(value.(int64))

		} else if reflect.TypeOf(value).Name() == "float64" {

			result = int(value.(float64))
		}
	}

	return
}

func (context MotadataMap) GetInt64Value(key string) (result int64) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "int64" {

			result = value.(int64)
		}
	}

	return
}

func (context MotadataMap) GetBoolValue(key string) (result bool) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "bool" {

			result = value.(bool)

		} else if reflect.TypeOf(value).Name() == "string" {

			if value.(string) == "yes" {

				result = true

			} else if value.(string) == "no" {

				result = false

			}
		}
	}

	return
}

func (context MotadataMap) GetFloatValue(key string) (result float64) {

	if context.Contains(key) {

		value := context[key]

		if reflect.TypeOf(value).Name() == "float64" {

			result = value.(float64)

		} else if reflect.TypeOf(value).Name() == "float32" {

			result = float64(value.(float32))
		}
	}

	return

}

func (context MotadataMap) GetTimeValue(key string) (result MotadataString) {

	if context.Contains(key) {

		return MotadataTime(context.GetFloat64Value(key)).ToString()
	}

	return
}

func (context MotadataMap) GetTrafficKBValue(key string) (result MotadataFloat64) {

	if context.Contains(key) {

		return MotadataKB(context.GetINTValue(key)).ToBytes()
	}

	return
}

func (context MotadataMap) GetTrafficMBValue(key string) (result MotadataFloat64) {

	if context.Contains(key) {

		return MotadataMB(context.GetINTValue(key)).ToBytes()
	}

	return
}

func (context MotadataMap) GetTrafficGBValue(key string) (result MotadataFloat64) {

	if context.Contains(key) {

		return MotadataGB(context.GetINTValue(key)).ToBytes()
	}

	return
}

func (context MotadataMap) GetMapSliceValue(key string) (result []MotadataMap) {

	if context.Contains(key) {

		return context[key].([]MotadataMap)
	}

	return
}

func (context MotadataMap) GetStringMapSliceValue(key string) (result []MotadataStringMap) {

	if context.Contains(key) {

		return context[key].([]MotadataStringMap)
	}

	return
}

func GetStringMapSliceValue(context cmap.ConcurrentMap[string, interface{}], key string) (result []MotadataStringMap) {

	if value, found := context.Get(key); found {

		return value.([]MotadataStringMap)

		/*typpe := reflect.TypeOf(value).Name()

		_ = typpe

		result = make([]MotadataStringMap, len(value.([]interface{})))

		for index, m := range value.([]interface{}) {

			result[index] = m.(MotadataStringMap)
		}*/
	}

	return
}

func (context MotadataMap) GetMapValue(key string) (result MotadataMap) {

	if context.Contains(key) {

		if reflect.TypeOf(context[key]).Name() == "MotadataMap" {

			result = context[key].(MotadataMap)

		} else if reflect.TypeOf(context[key]).String() == "map[string]interface {}" {

			result = context[key].(map[string]interface{})
		}

	}

	return
}

func GetMapValue(context cmap.ConcurrentMap[string, interface{}], key string) (result MotadataMap) {

	if value, found := context.Get(key); found {

		if reflect.TypeOf(value).Name() == "MotadataMap" {

			result = value.(MotadataMap)

		} else if reflect.TypeOf(value).String() == "map[string]interface {}" {

			result = value.(map[string]interface{})
		}
	}

	return
}

func (context MotadataMap) GetStringMapValue(key string) (result MotadataStringMap) {

	if context.Contains(key) {

		if reflect.TypeOf(context[key]).Name() == "MotadataStringMap" {

			result = context[key].(MotadataStringMap)

		} else if reflect.TypeOf(context[key]).String() == "map[string]string {}" {

			result = context[key].(map[string]string)
		}
	}

	return
}

func (context MotadataMap) GetIntFloatMapValue(key string) (result MotadataIntFloatMap) {

	if context.Contains(key) {

		if reflect.TypeOf(context[key]).Name() == "MotadataIntFloatMap" {

			result = context[key].(MotadataIntFloatMap)

		}
	}

	return
}

func (context MotadataMap) GetSliceValue(key string) (result []interface{}) {

	if context.Contains(key) {

		return context[key].([]interface{})
	}

	return
}

func GetSliceValue(context cmap.ConcurrentMap[string, interface{}], key string) (result []interface{}) {

	if value, found := context.Get(key); found {

		return value.([]interface{})
	}

	return
}

func (context MotadataMap) GetListValue(key string) (result MotadataStringList) {

	if context.Contains(key) {

		return context[key].(MotadataStringList)
	}

	return
}

func (context MotadataIntFloatMap) GetMaxKey() int {

	result := 0

	for key := range context {

		if result < key {

			result = key

		}
	}
	return result
}

func (context MotadataIntFloatMap) IsNotEmpty() bool {

	if context != nil && len(context) > 0 {

		return true
	}

	return false
}

func ToStringSlice(value interface{}) (result []MotadataString) {

	if value != nil {

		result = value.([]MotadataString)
	}

	return
}

func ToMap(value interface{}) (result MotadataMap) {

	if value != nil {

		if reflect.TypeOf(value).Name() == "MotadataMap" {

			result = value.(MotadataMap)

		} else if reflect.TypeOf(value).String() == "map[string]interface {}" {

			result = value.(map[string]interface{})
		}

	}

	return
}

func ToObjectSlice(value interface{}) (result []interface{}) {

	if value != nil {

		if reflect.TypeOf(value).String() == "[]interface {}" {

			result = value.([]interface{})

		}
	}
	return
}

func ToMotadataString(value interface{}) (result MotadataString) {

	if reflect.TypeOf(value).Name() == "MotadataString" {

		result = value.(MotadataString)

	} else if reflect.TypeOf(value).Name() == "string" {

		result = MotadataString(value.(string))

	} else if reflect.TypeOf(value).Name() == "int" {

		result = MotadataString(strconv.Itoa(value.(int)))

	} else if reflect.TypeOf(value).String() == "[]uint8" {

		result = MotadataString(value.([]uint8))

	} else if reflect.TypeOf(value).Name() == "float64" {

		result = MotadataString(strconv.Itoa(int(value.(float64))))

	} else if reflect.TypeOf(value).Name() == "bool" {

		result = MotadataString(strconv.FormatBool(value.(bool)))
	}

	return

}

func ToString(value interface{}) (result string) {

	if reflect.TypeOf(value).Name() == "string" {

		result = value.(string)

	} else if reflect.TypeOf(value).Name() == "MotadataString" {

		result = string(value.(MotadataString))

	} else if reflect.TypeOf(value).Name() == "int" {

		result = strconv.Itoa(value.(int))

	} else if reflect.TypeOf(value).String() == "[]uint8" {

		result = string(value.([]uint8))

	} else if reflect.TypeOf(value).String() == "[]byte" {

		result = string(value.([]byte))

	} else if reflect.TypeOf(value).Name() == "float64" {

		result = strconv.Itoa(int(value.(float64)))

	} else if reflect.TypeOf(value).Name() == "bool" {

		result = strconv.FormatBool(value.(bool))
	} else if reflect.TypeOf(value).Name() == "MotadataINT" {

		result = strconv.Itoa(int(value.(MotadataINT)))
	} else if reflect.TypeOf(value).Name() == "MotadataFloat64" {

		result = strconv.Itoa(int(value.(MotadataFloat64)))

	}

	return

}

func ToINT(value interface{}) (result MotadataINT) {

	return MotadataINT(ToInt(value))
}

func ToInt(value interface{}) (result int) {

	if reflect.TypeOf(value).Name() == "string" {

		output, _ := strconv.ParseInt(strings.TrimSpace(value.(string)), 10, 64)

		result = int(output)

	} else if reflect.TypeOf(value).Name() == "MotadataUINT" {

		result = int(value.(MotadataUINT))

	} else if reflect.TypeOf(value).Name() == "uint" {

		result = int(value.(uint))

	} else if reflect.TypeOf(value).Name() == "uint8" {

		result = int(value.(uint8))

	} else if reflect.TypeOf(value).Name() == "MotadataUINT16" {

		result = int(value.(MotadataUINT16))

	} else if reflect.TypeOf(value).Name() == "uint16" {

		result = int(value.(uint16))

	} else if reflect.TypeOf(value).Name() == "uint32" {

		result = int(value.(uint32))

	} else if reflect.TypeOf(value).Name() == "uint64" {

		result = int(value.(uint64))

	} else if reflect.TypeOf(value).Name() == "int" {

		result = value.(int)

	} else if reflect.TypeOf(value).Name() == "int8" {

		result = int(value.(int8))

	} else if reflect.TypeOf(value).Name() == "int16" {

		result = int(value.(int16))

	} else if reflect.TypeOf(value).Name() == "int32" {

		result = int(value.(int32))

	} else if reflect.TypeOf(value).Name() == "int64" {

		result = int(value.(int64))

	} else if reflect.TypeOf(value).Name() == "float64" {

		result = int(value.(float64))
	}

	return
}

func ToJSON(value interface{}) (result MotadataString) {

	tokens, _ := json.Marshal(value)

	result = MotadataString(tokens)

	return
}

func Contains(elements []string, value MotadataString) bool {

	for _, element := range elements {

		if MotadataString(element) == value {

			return true
		}
	}
	return false
}

func (elements MotadataStringList) GetIndexByValues(values MotadataStringList) int {

	for _, value := range values {

		if elements.Contains(MotadataString(value)) {

			for index, element := range elements {

				if element == value {

					return index
				}
			}
		}
	}

	return -1
}

func ToList(elements []string) MotadataStringList {

	return append(MotadataStringList{}, elements...)
}

func ToSlice(elements MotadataStringList) []MotadataString {

	var values []MotadataString

	for _, element := range elements {

		values = append(values, MotadataString(element))
	}

	return values
}

func StringSliceToList(elements []MotadataString) MotadataStringList {

	values := MotadataStringList{}

	for _, element := range elements {

		values = append(values, element.ToString())
	}

	return values
}

func (context MotadataMap) ContainValues(values []MotadataMap) bool {

	for _, value := range values {

		if reflect.DeepEqual(value, context) {

			return true
		}
	}

	return false
}

func ToFloat(value interface{}) (output float64) {

	if reflect.ValueOf(value).Kind().String() == "float32" {

		output = float64(value.(float32))

	} else if reflect.ValueOf(value).Kind().String() == "float64" {

		output = value.(float64)

	}

	return
}

func ToMotadataFloat(value interface{}) (output MotadataFloat64) {

	if value != nil {

		if reflect.ValueOf(value).Kind().String() == "float64" {

			output = MotadataFloat64(value.(float64))

		}
	}

	return
}

func MakeCmap() (object cmap.ConcurrentMap[string, interface{}]) {

	return cmap.New[interface{}]()
}

func SetErrors(context cmap.ConcurrentMap[string, interface{}], errors []MotadataStringMap, err MotadataStringMap) {

	if errors != nil {

		if value, found := context.Get(Errors); found {

			str := reflect.TypeOf(value).String()

			if str == "[]interface {}" {

				context.Set(Errors, append(value.([]interface{}), errors))

			} else {

				context.Set(Errors, append(GetStringMapSliceValue(context, Errors), errors...))
			}
		} else {

			context.Set(Errors, append(GetStringMapSliceValue(context, Errors), errors...))
		}
	} else {

		if value, found := context.Get(Errors); found {

			str := reflect.TypeOf(value).String()

			if str == "[]interface {}" {

				context.Set(Errors, append(value.([]interface{}), err))

			} else {

				context.Set(Errors, append(GetStringMapSliceValue(context, Errors), err))
			}
		} else {

			context.Set(Errors, append(GetStringMapSliceValue(context, Errors), err))
		}
	}

	context.Set(Message, err[Message])
}
