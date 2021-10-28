package phpgo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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
	if arrayDataValue.Type().Kind().String() != "slice" {
		return false
	}
	for i := 0; i < arrayDataValue.Len(); i++ {
		element := arrayDataValue.Index(i)
		if element.Type().String() != itemValue.Type().String() {
			// 元素类型不一样
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
	if dataValue.Type().Kind().String() != "slice" {
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

// ArrayChunkInt 整数版本的数组切割。将arrayData按照每个长度为length切割为子数组
func ArrayChunkInt(arrayData []int, length int) [][]int {
	dataLen := len(arrayData)
	if length <= 0 {
		return nil
	}
	var result [][]int
	for i := 0; i < dataLen; i += length {
		if length+i > dataLen {
			result = append(result, arrayData[i:])
		} else {
			result = append(result, arrayData[i:i+length])
		}
	}
	return result
}

// ArrayChunkString 字符串版本的数组切割。将arrayData按照每个长度为length切割为子数组
func ArrayChunkString(arrayData []string, length int) [][]string {
	dataLen := len(arrayData)
	if length <= 0 {
		return nil
	}
	var result [][]string
	for i := 0; i < dataLen; i += length {
		if length+i > dataLen {
			result = append(result, arrayData[i:])
		} else {
			result = append(result, arrayData[i:i+length])
		}
	}
	return result
}

// ArrayDiff 模拟PHP array_diff函数 计算差集
func ArrayDiff(base []interface{}, othersParams ...[]interface{}) []interface{} {
	if len(base) == 0 {
		return []interface{}{}
	}
	if len(base) > 0 && len(othersParams) == 0 {
		return base
	}
	var tmp = make(map[interface{}]int, len(base))
	for _, v := range base {
		tmp[v] = 1
	}
	for _, param := range othersParams {
		for _, arg := range param {
			if tmp[arg] != 0 {
				tmp[arg]++
			}
		}
	}
	var res = make([]interface{}, 0, len(tmp))
	for k, v := range tmp {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// ArrayIntersect 模拟PHP array_intersect函数 计算交集
func ArrayIntersect(base []interface{}, othersParams ...[]interface{}) []interface{} {
	if len(base) == 0 {
		return []interface{}{}
	}
	if len(base) > 0 && len(othersParams) == 0 {
		return base
	}
	var tmp = make(map[interface{}]int, len(base))
	for _, v := range base {
		tmp[v] = 1
	}
	for _, param := range othersParams {
		for _, arg := range param {
			if tmp[arg] != 0 {
				tmp[arg]++
			}
		}
	}
	var res = make([]interface{}, 0, len(tmp))
	for k, v := range tmp {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

// ArrayMerge PHP array_merge函数，数组合并，暂时支持int/string两种类型
func ArrayMerge(datas ...interface{}) (interface{}, error) {
	if len(datas) == 0 {
		return nil, fmt.Errorf("参数不可为空")
	}
	var arrayCap int
	var dataType string
	allowDataType := []string{"int", "string"}
	for _, item := range datas {
		// 类型必须为切片
		refValue := reflect.ValueOf(item)
		if refValue.Type().Kind().String() != "slice" {
			return nil, fmt.Errorf("参数必须为一个切片")
		}
		if refValue.Len() > 0 {
			if dataType != "" && dataType != refValue.Index(0).Type().String() {
				return nil, fmt.Errorf("每一个切片数组的数据类型必须一致")
			}
			dataType = refValue.Index(0).Type().String()
			if !InArrayString(dataType, allowDataType) {
				return nil, fmt.Errorf("参数类型错误，仅支持:%s", strings.Join(allowDataType, ","))
			}

		}
		arrayCap += refValue.Len()
	}
	if dataType == "string" {
		result := make([]string, 0, arrayCap)
		for _, item := range datas {
			refValue := reflect.ValueOf(item)
			if refValue.Len() > 0 {
				for i := 0; i < refValue.Len(); i++ {
					result = append(result, refValue.Index(i).String())
				}
			}
		}
		return result, nil
	} else {
		result := make([]int, 0, arrayCap)
		for _, item := range datas {
			refValue := reflect.ValueOf(item)
			if refValue.Len() > 0 {
				for i := 0; i < refValue.Len(); i++ {
					result = append(result, int(refValue.Index(i).Int()))
				}
			}
		}
		return result, nil
	}
}

// ArraySearch 搜索arrayData里面是否有item，有返回对应的index，无返回-1，只返回首次
func ArraySearch(item interface{}, arrayData interface{}) (int, error) {
	itemValue := reflect.ValueOf(item)
	arrayDataValue := reflect.ValueOf(arrayData)
	if arrayDataValue.Type().Kind().String() != "slice" {
		return 0, fmt.Errorf("arrayData参数不是切片")
	}
	for i := 0; i < arrayDataValue.Len(); i++ {
		dataItem := arrayDataValue.Index(i)
		if itemValue.Type().String() != dataItem.Type().String() {
			return 0, fmt.Errorf("两个参数的数据类型不一致")
		}
		if dataItem.Interface() == item {
			return i, nil
		}
	}
	return -1, nil
}

// ArrayProduct 计算数组的各元素的乘积
func ArrayProduct(data []int) int64 {
	var result int64
	for _, num := range data {
		result *= int64(num)
	}
	return result
}

// ArrayWalk 遍历某一个数组切片，callback返回false则停止遍历
func ArrayWalk(data interface{}, callback func(item interface{}, index int) bool) error {
	ref := reflect.ValueOf(data)
	if ref.Type().Kind().String() != "slice" {
		return fmt.Errorf("参数必须是一个切片")
	}
	for i := 0; i < ref.Len(); i++ {
		if !callback(ref.Index(i).Interface(), i) {
			break
		}
	}
	return nil
}

// JoinInt 将一个int slice转为sep分割的字符串
func JoinInt(a []int, sep string) string {
	var s strings.Builder
	sLen := len(a)
	for i, item := range a {
		s.WriteString(strconv.Itoa(item))
		if i < sLen-1 {
			s.WriteString(sep)
		}
	}
	return s.String()
}

// ArrayInt2ArrayString 将int slice转为string slice
func ArrayInt2ArrayString(in []int) []string {
	length := len(in)
	data := make([]string, length)
	for i := 0; i < length; i++ {
		data[i] = strconv.Itoa(in[i])
	}
	return data
}

// ArrayReverse 数组反转，需要传指针进来，仅支持slice，array请转为slice
func ArrayReverse(array interface{}) error {
	ref := reflect.ValueOf(array)
	if ref.Kind() != reflect.Ptr {
		return errors.New("参数必须为指针类型")
	}
	ref = ref.Elem()
	if ref.Kind() != reflect.Slice {
		return errors.New("参数必须为slice")
	}
	swap := reflect.Swapper(ref.Interface())
	for i, j := 0, ref.Len()-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return nil
}
