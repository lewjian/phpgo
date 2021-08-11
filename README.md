# phpgo
PHP有很多方便的函数，GO里面没有或者命名和PHP不一样，为了方便PHP转GO的同学方便，特写了本项目。
> 项目里面的所有函数命名都是和PHP函数一致，仅仅是首字母大写，下划线_去掉。仅支持go mod
# 使用方法
```go
import "github.com/lewjian/phpgo"

phpgo.InArray("a", []string{"a"})
```
# 进度
- 常用string函数
- 常用array函数
- 常用文件系统函数
- 常用日期相关函数
- 三元表达式

# 函数索引
```go
CONSTANTS

const (
        DefaultDateTimeFormatTpl = "2006-01-02 15:04:05"
        DefaultDateFormatTpl     = "2006-01-02"
        DefaultTimeFormatTpl     = "15:04:05"
)
const (
        StrPadLeft = iota
        StrPadRight
        StrPadBoth
)

FUNCTIONS

func ArrayChunkInt(arrayData []int, length int) [][]int
    ArrayChunkInt 整数版本的数组切割。将arrayData按照每个长度为length切割为子数组

func ArrayChunkString(arrayData []string, length int) [][]string
    ArrayChunkString 字符串版本的数组切割。将arrayData按照每个长度为length切割为子数组

func ArrayDiff(base []interface{}, othersParams ...[]interface{}) []interface{}
    ArrayDiff 模拟PHP array_diff函数 计算差集

func ArrayInt2ArrayString(in []int) []string
    ArrayInt2ArrayString 将int slice转为string slice

func ArrayIntersect(base []interface{}, othersParams ...[]interface{}) []interface{}
    ArrayIntersect 模拟PHP array_intersect函数 计算交集

func ArrayMerge(datas ...interface{}) (interface{}, error)
    ArrayMerge PHP array_merge函数，数组合并，暂时支持int/string两种类型

func ArrayProduct(data []int) int64
    ArrayProduct 计算数组的各元素的乘积

func ArraySearch(item interface{}, arrayData interface{}) (int, error)
    ArraySearch 搜索arrayData里面是否有item，有返回对应的index，无返回-1，只返回首次

func ArraySum(data []int) int64
    ArraySum 计算数组之和

func ArrayUnique(arrayData interface{}) interface{}
    ArrayUnique 切片去重，目前仅支持int和string两种类型

func ArrayWalk(data interface{}, callback func(item interface{}, index int) bool) error
    ArrayWalk 遍历某一个数组切片，callback返回false则停止遍历

func Basename(filename string, suffix ...string) string
    Basename 返回某个文件路径的文件名，suffix如果提供了则去掉改后缀

func Bin2Hex(binData []byte) string
    Bin2Hex 将二进制数据转为16进制字符串

func ChDir(dir string) bool
    ChDir 改变当前目录

func ChMod(filename string, mode os.FileMode) bool
    ChMod 改变文件权限

func Chown(filename, username string) bool
    Chown 改变文件所属者

func Copy(src, dest string) bool
    Copy 复制文件

func Date(format string, timestamp ...int64) string
    Date PHP的date函数 timestamp 时间戳，传了时间戳则以此为准，不传默认当前时间
    用法参考PHP函数：https://www.php.net/manual/zh/function.date.php
    但是不是每一个标志都实现了，具体可以参考下面的map对照表

func Dirname(path string) string
    Dirname 返回路径中的目录部分

func Explode(separator, str string) []string
    Explode 用一个字符串separator分割另外一个字符串str

func FileATime(filename string) int64
    FileATime 获取文件上次访问时间，返回为秒的时间戳

func FileExists(filename string) bool
    FileExists 检查文件是否存在

func FileGetContents(filename string) []byte
    FileGetContents 获取文件内容

func FileMTime(filename string) int64
    FileMTime 获取文件修改时间

func FilePutContents(filename string, content []byte, fileMode os.FileMode) (int, error)
    FilePutContents 写入文件内容

func FileSize(filename string) int64
    FileSize 获取文件大小

func GetCWD() string
    GetCWD 取得当前工作目录

func GetSystemOS() string
    GetSystemOS 获取当前运行的操作系统类型

func Hex2Bin(s string) ([]byte, error)
    Hex2Bin 将16进制字符串转为byte数组

func If(isTrue bool, a, b interface{}) interface{}
    If Golang版本的三元表达式，使用方法：If(a>b,a,b).(int)

func Implode(separator string, data []string) string
    Implode 切片转为指定字符串连接的为字符串

func InArray(item interface{}, arrayData interface{}) bool
    InArray PHP对应的in_array函数，支持int，string，float的判断

func InArrayInt(n int, nums []int) bool
    InArrayInt php in_array 整数版本

func InArrayString(s string, data []string) bool
    InArrayString php in_array字符串版本

func IsDir(filename string) bool
    IsDir 判断是否为文件夹

func IsFile(filename string) bool
    IsFile 判断是否为文件

func IsReadable(filename string) bool
    IsReadable 判断是否可读

func IsWritable(filename string) bool
    IsWritable 判断文件是否可写

func Join(separator string, data []string) string
    Join 切片转为指定字符串连接的为字符串

func JoinInt(a []int, sep string) string
    JoinInt 将一个int slice转为sep分割的字符串

func LCFirst(s string) string
    LCFirst 首字母小写写

func LTrim(str, deleteStr string) string
    LTrim 去除字符串左边指定字符串

func MD5(data []byte) string
    MD5 计算md5值

func MkDir(pathname string, fileMode os.FileMode, recursive bool) bool
    MkDir 创建文件夹

func ParseDate(dateStr string, layouts ...string) (time.Time, error)
    ParseDate 解析一个日期字符串 layouts省略则采用默认2006-01-02 15:04:05

func PregMatch(pattern, s string) ([]string, error)
    PregMatch 使用正则匹配字符串s中第一个匹配的数据

func PregMatchAll(pattern, s string) ([][]string, error)
    PregMatchAll 使用正则匹配字符串s中所有匹配的数据

func PregReplace(pattern, repl, src string) (string, error)
    PregReplace 正则替换

func PregSplit(pattern, src string) ([]string, error)
    PregSplit 正则字符串拆分为数组

func RTrim(str, deleteStr string) string
    RTrim 去除字符串右边边指定字符串

func ReadDir(dir string) ([]os.DirEntry, error)
    ReadDir 读取目录，golang里面不用opendir

func Rename(src, dest string) error
    Rename 重命名

func Sha1(data []byte) string
    Sha1 计算sha1值

func StrCMP(s1, s2 string) int
    StrCMP 字符串比较，相等返回0，s1<s2,返回-1，s1>s2，返回1

func StrContains(str, subStr string) bool
    StrContains 判断是否某一个字符串包含一个子串

func StrEndsWith(s, suffixStr string) bool
    StrEndsWith 检查一个字符串是否以某个子串结尾

func StrIPos(s, subStr string) int
    StrIPos 查找字符串首次出现的位置（不区分大小写）

func StrIReplace(searchStr, replaceStr, findInStr string) (string, error)
    StrIReplace 大小写忽略的字符串替换

func StrLen(s string) int
    StrLen 获取字符串长度

func StrPad(str string, padLen int, padStr string, padType int) string
    StrPad 使用另一个字符串填充字符串为指定长度

func StrPos(s, subStr string) int
    StrPos 查找字符串首次出现的位置（区分大小写）

func StrRIPos(s, subStr string) int
    StrRIPos 查找字符串最后出现的位置（不区分大小写）

func StrRPos(s, subStr string) int
    StrRPos 查找字符串最后出现的位置（区分大小写）

func StrRepeat(s string, count int) string
    StrRepeat 重复一个字符串

func StrReplace(search, repl, replaceIn string) string
    StrReplace 字符串替换

func StrRev(s string) string
    StrRev 字符串翻转，按照字符翻转，不是按照字节

func StrShuffle(s string) string
    StrShuffle 随机打乱一个字符串

func StrSplit(s string, splitLen int) []string
    StrSplit 将字符串转为切片（slice），按照splitLen个字符为一个单位。 原来PHP里面是按照字节切割，这里换成了按照字符切割

func StrStartWith(s, prefix string) bool
    StrStartWith 检查某个字符传s是否以prefix开始

func StrStr(s, substr string, needBefore bool) string
    StrStr 查找字符串首次出现前(后)的字符串 如果needBefore为true表示需要substr首次出现之前的字符串，否则为之后的字符串

func StrToLower(s string) string
    StrToLower 字符串转为小写

func StrToTime(format string, relativeTimestamp ...int64) (time.Time, error)
    StrToTime 模拟PHP的strtotime函数 支持format有："+1 year"、"-1 year"、"+5 months"、"1
    year +5 month -1 day +10 hours -23 minutes +5 seconds"等
    relativeTimestamp是时间戳，单位秒，之所以设为可变参数，是为了和php一致，使用时可以省略，省略则默认当前时间
    relativeTimestamp传多个值没用，只用到relativeTimestamp[0]

func StrToUpper(s string) string
    StrToUpper 字符串转为大写

func StripTags(src string) string
    StripTags 从字符串中去除 HTML 标签

func SubStr(s string, start, length int) string
    SubStr 字符串截取，需要注意是否越界，这是按照字节截取的

func SubStrRune(s string, start, length int) string
    SubStrRune 字符串截取，需要注意是否越界，这是按照字符截取的

func Time() int64
    Time 获取当前时间戳，单位：秒

func ToTimestamp(dateStr string, layouts ...string) int64
    ToTimestamp 将一个日期/时间转为时间戳，出错则返回-1 layouts省略则采用默认2006-01-02 15:04:05

func Trim(str string) string
    Trim 删除字符串前后空白符

func UCFirst(s string) string
    UCFirst 首字母大写

func UCWords(s string) string
    UCWords 将字符串中每个单词的首字母转换为大写

func UnLink(filename string) bool
    UnLink 删除文件
```
