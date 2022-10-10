<?php
include_once 'Tea.php';

$key = "1234657890abcdef";
$str = "hello qq tea";
//请求参数解密 转换为数组
$data = Tea::encrypt($key, $str);
// print_r($data);
$data1 = bin2hex($data);
print_r($data1);
echo "\n";


$str = "109bdaf488698a181308314b97d17f919e5c112f1016ff93";
$data = hex2bin($str);
$data2 = Tea::decrypt($key, $data);
print_r($data2);
