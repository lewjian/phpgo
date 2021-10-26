package phpgo

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"html"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Bin2Hex 将二进制数据转为16进制字符串
func Bin2Hex(binData []byte) string {
	return hex.EncodeToString(binData)
}

// Hex2Bin 将16进制字符串转为byte数组
func Hex2Bin(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

// Explode 用一个字符串separator分割另外一个字符串str
func Explode(separator, str string) []string {
	return strings.Split(str, separator)
}

// Implode 切片转为指定字符串连接的为字符串
func Implode(separator string, data []string) string {
	return strings.Join(data, separator)
}

// Join 切片转为指定字符串连接的为字符串
func Join(separator string, data []string) string {
	return strings.Join(data, separator)
}

// LTrim 去除字符串左边指定字符串
func LTrim(str, deleteStr string) string {
	return strings.TrimLeft(str, deleteStr)
}

// RTrim 去除字符串右边边指定字符串
func RTrim(str, deleteStr string) string {
	return strings.TrimRight(str, deleteStr)
}

// Trim 删除字符串前后空白符
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// MD5 计算md5值
func MD5(data []byte) string {
	sumData := md5.Sum(data)
	return Bin2Hex(sumData[:])
}

// Sha1 计算sha1值
func Sha1(data []byte) string {
	sumData := sha1.Sum(data)
	return Bin2Hex(sumData[:])
}

// StrContains 判断是否某一个字符串包含一个子串
func StrContains(str, subStr string) bool {
	return strings.Contains(str, subStr)
}

// StrEndsWith 检查一个字符串是否以某个子串结尾
func StrEndsWith(s, suffixStr string) bool {
	return strings.HasSuffix(s, suffixStr)
}

// StrIReplace 大小写忽略的字符串替换
func StrIReplace(searchStr, replaceStr, findInStr string) (string, error) {
	reg, err := regexp.Compile(`(?i)` + searchStr)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(findInStr, replaceStr), nil
}

const (
	StrPadLeft = iota
	StrPadRight
	StrPadBoth
)

// StrPad 使用另一个字符串填充字符串为指定长度
func StrPad(str string, padLen int, padStr string, padType int) string {
	strLen := len(str)
	if padLen < strLen {
		return str
	}
	if padType < StrPadLeft || padType > StrPadBoth {
		padType = StrPadRight
	}
	diffNum := padLen - strLen
	calculatePadStr := strings.Repeat(padStr, diffNum)
	if len(calculatePadStr) > padLen-strLen {
		calculatePadStr = calculatePadStr[:diffNum]
	}

	if padType == StrPadLeft {
		return fmt.Sprintf("%s%s", calculatePadStr, str)
	}
	if padType == StrPadRight {
		return fmt.Sprintf("%s%s", str, calculatePadStr)
	}
	leftNum := diffNum / 2
	rightNum := diffNum - leftNum
	return fmt.Sprintf("%s%s%s", strings.Repeat(padStr, leftNum)[:leftNum], str, strings.Repeat(padStr, rightNum)[:rightNum])
}

// StrRepeat 重复一个字符串
func StrRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// StrReplace 字符串替换
func StrReplace(search, repl, replaceIn string) string {
	return strings.ReplaceAll(replaceIn, search, repl)
}

// StrShuffle 随机打乱一个字符串
func StrShuffle(s string) string {
	data := []rune(s)
	sLen := len(data)
	rand.Seed(time.Now().UnixNano())
	result := make([]rune, 0, sLen)
	for sLen > 0 {
		index := rand.Intn(sLen)
		result = append(result, data[index])
		if sLen-1 != index {
			data[index] = data[sLen-1]
		}
		data = data[0 : sLen-1]
		sLen = len(data)
	}
	return string(result)
}

// StrSplit 将字符串转为切片（slice），按照splitLen个字符为一个单位。
// 原来PHP里面是按照字节切割，这里换成了按照字符切割
func StrSplit(s string, splitLen int) []string {
	if splitLen < 1 {
		splitLen = 1
	}
	data := []rune(s)
	dataLen := len(data)
	resLen := int(math.Ceil(float64(dataLen / splitLen)))
	result := make([]string, 0, resLen)
	for i := 0; i < dataLen; i += splitLen {
		if i+splitLen < dataLen {
			result = append(result, string(data[i:i+splitLen]))
		} else {
			result = append(result, string(data[i:]))
		}
	}
	return result
}

// StrStartWith 检查某个字符传s是否以prefix开始
func StrStartWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// StrCMP 字符串比较，相等返回0，s1<s2,返回-1，s1>s2，返回1
func StrCMP(s1, s2 string) int {
	return strings.Compare(s1, s2)
}

// StripTags 从字符串中去除 HTML 标签
func StripTags(src string) string {
	re, _ := regexp.Compile(`<\/?[a-zA-Z0-9-]+(\s+[a-zA-Z-]+=['"].*['"])*>`)
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re, _ = regexp.Compile(`\s{2,}`)
	src = re.ReplaceAllString(src, "\n")
	src = html.UnescapeString(src)
	return strings.TrimSpace(src)
}

// StrIPos 查找字符串首次出现的位置（不区分大小写）
func StrIPos(s, subStr string) int {
	s = strings.ToLower(s)
	subStr = strings.ToLower(subStr)
	return strings.Index(s, subStr)
}

// StrPos 查找字符串首次出现的位置（区分大小写）
func StrPos(s, subStr string) int {
	return strings.Index(s, subStr)
}

// StrRIPos 查找字符串最后出现的位置（不区分大小写）
func StrRIPos(s, subStr string) int {
	s = strings.ToLower(s)
	subStr = strings.ToLower(subStr)
	return strings.LastIndex(s, subStr)
}

// StrRPos 查找字符串最后出现的位置（区分大小写）
func StrRPos(s, subStr string) int {
	return strings.LastIndex(s, subStr)
}

// StrLen 获取字符串长度
func StrLen(s string) int {
	return len(s)
}

// StrRev 字符串翻转，按照字符翻转，不是按照字节
func StrRev(s string) string {
	if len(s) == 0 {
		return ""
	}
	data := []rune(s)
	dataLen := len(data)
	for i := 0; i < dataLen/2; i++ {
		data[i], data[dataLen-1-i] = data[dataLen-1-i], data[i]
	}
	return string(data)
}

// StrStr 查找字符串首次出现前(后)的字符串
// 如果needBefore为true表示需要substr首次出现之前的字符串，否则为之后的字符串
func StrStr(s, substr string, needBefore bool) string {
	index := strings.Index(s, substr)
	if index == -1 {
		// 不存在子字符串
		return ""
	}
	if needBefore {
		// 查找该子字符串前面部分
		return s[:index]
	}
	return s[index:]
}

// StrToLower 字符串转为小写
func StrToLower(s string) string {
	return strings.ToLower(s)
}

// StrToUpper 字符串转为大写
func StrToUpper(s string) string {
	return strings.ToUpper(s)
}

// SubStr 字符串截取，需要注意是否越界，这是按照字节截取的
func SubStr(s string, start, length int) string {
	return s[start : start+length]
}

// SubStrRune 字符串截取，需要注意是否越界，这是按照字符截取的
func SubStrRune(s string, start, length int) string {
	data := []rune(s)
	return string(data[start : start+length])
}

// UCFirst 首字母大写
func UCFirst(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(s[:1]), s[1:])
}

// LCFirst 首字母小写写
func LCFirst(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	return fmt.Sprintf("%s%s", strings.ToLower(s[:1]), s[1:])
}

// UCWords 将字符串中每个单词的首字母转换为大写
func UCWords(s string) string {
	if s == "" {
		return ""
	}
	reg, err := regexp.Compile(`[\t\r\n(\r\n)\f\v\s]+`)
	if err != nil {
		return s
	}
	res := reg.Split(s, -1)
	var destStr strings.Builder
	for _, myStr := range res {
		destStr.WriteString(UCFirst(myStr))
		destStr.WriteString(" ")
	}
	return strings.TrimSpace(destStr.String())
}

// PregMatch 使用正则匹配字符串s中第一个匹配的数据
func PregMatch(pattern, s string) ([]string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return reg.FindStringSubmatch(s), nil
}

// PregMatchAll 使用正则匹配字符串s中所有匹配的数据
func PregMatchAll(pattern, s string) ([][]string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return reg.FindAllStringSubmatch(s, -1), nil
}

// PregReplace 正则替换
func PregReplace(pattern, repl, src string) (string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(src, repl), nil
}

// PregSplit 正则字符串拆分为数组
func PregSplit(pattern, src string) ([]string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return reg.Split(src, -1), nil
}
