package phpgo

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// InArrayString php in_array字符串版本
func InArrayString(s string, data []string) bool {
	for i := 0; i < len(data); i++ {
		if data[i] == s {
			return true
		}
	}
	return false
}

// InArrayInt php in_array 整数版本
func InArrayInt(n int, nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == n {
			return true
		}
	}
	return false
}

// InArray PHP对应的in_array函数，支持int，string，float的判断
func InArray(item interface{}, arrayData interface{}) bool {
	itemValue := reflect.ValueOf(item)
	arrayDataValue := reflect.ValueOf(arrayData)
	// arrayData必须类型的[]T
	if !strings.Contains(arrayDataValue.Type().String(), "[]") {
		return false
	}
	for i := 0; i < arrayDataValue.Len(); i++ {
		element := arrayDataValue.Index(i)
		if element.Type().String() != itemValue.Type().String() {
			// 元素类型不一样
			log.Println(element.Type().String(), itemValue.Type().String())
			return false
		}
		switch itemValue.Type().String() {
		case "string":
			if itemValue.String() == element.String() {
				return true
			}
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			if itemValue.Int() == element.Int() {
				return true
			}
		case "float32", "float64":
			if itemValue.Float() == element.Float() {
				return true
			}
		default:
			panic(fmt.Sprintf("inArray不支持数据结构:%s", element.Type().String()))

		}
	}
	return false
}

// ArrayUnique 切片去重，目前仅支持int和string两种类型
func ArrayUnique(arrayData interface{}) interface{} {
	dataValue := reflect.ValueOf(arrayData)
	// arrayData必须类型的[]T
	if !strings.Contains(dataValue.Type().String(), "[]") {
		return arrayData
	}
	// 牺牲内存空间，采用map来处理
	m := make(map[interface{}]struct{})
	var itemType string
	allowItemType := []string{"int", "string"}
	for i := 0; i < dataValue.Len(); i++ {
		if itemType != "" && itemType != dataValue.Index(i).Type().String() {
			// 每个元素类型不一样
			panic("必须保证是一个指定类型的切片，支持int和string")
		}
		itemType = dataValue.Index(i).Type().String()
		if !InArray(itemType, allowItemType) {
			panic("必须保证是一个指定类型的切片，支持int和string")
		}
		m[dataValue.Index(i).Interface()] = struct{}{}
	}
	if itemType == "string" {
		result := make([]string, 0, len(m))
		for key, _ := range m {
			result = append(result, key.(string))
		}
		return result
	} else {
		result := make([]int, 0, len(m))
		for key, _ := range m {
			result = append(result, key.(int))
		}
		return result
	}
}

// ArraySum 计算数组之和
func ArraySum(data []int) int64 {
	var total int64
	for _, num := range data {
		total += int64(num)
	}
	return total
}
