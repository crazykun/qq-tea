# qq-tea
QQ Tea 加/解密算法, 仅用于学习研究



## php版本
```
$key = "1234657890abcdef";
$str = "hello qq tea";
// 加密
$data = Tea::encrypt($key, $str);
// 转换16进制
$data1 = bin2hex($data);
// 解密
$data2 = Tea::decrypt($key, $data);
print_r($data2);
```
[qqtea](https://github.com/manyhelp/qqtea.php)


## golang版本
```
k := "1234657890abcdef" 
c := tea.NewTeaCipher([]byte(k))

//加密字符串
str := "hello qq tea go"

// 加密
result := c.Encrypt([]byte(str))
encodedStr := hex.EncodeToString(result)
fmt.Println(encodedStr)

// 解密
result = c.Decrypt(result)
fmt.Println(string(result))
```

[go-qq-tea](https://github.com/littlefish12345/go-qq-tea)

## python版本
[chinaunix](http://bbs.chinaunix.net/thread-583468-1-1.html)
代码注释非常良心


## java版本

[掘金](https://juejin.cn/post/6844903774406836237)

[博客园](https://www.cnblogs.com/raikouissen/p/3393222.html)

## javascript版本
[github](https://github.com/sun8911879/qqtea-1) `(未测试)`




