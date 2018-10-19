function CheckUtil() { }
//校验是否为空(先删除二边空格再验证)
CheckUtil.isNull = function (str) {
if (null == str ||  ""== str.trim()) {
  return true;
} else {
  return false;
}
};
//校验是否全是数字
CheckUtil.isDigit  = function (str) {
var patrn=/^\d+$/;
return patrn.test(str);
};
//校验是否是整数
CheckUtil.isInteger = function (str) {
var patrn=/^([+-]?)(\d+)$/;
return patrn.test(str);
};
//校验是否为正整数
CheckUtil.isPlusInteger = function (str) {
var patrn=/^([+]?)(\d+)$/;
return patrn.test(str);
};
//校验是否为负整数
CheckUtil.isMinusInteger = function (str) {
var patrn=/^-(\d+)$/;
return patrn.test(str);
};
//校验是否为浮点数
CheckUtil.isFloat=function(str){
var patrn=/^([+-]?)\d*\.\d+$/;
return patrn.test(str)||CheckUtil.isInteger(str);
};
//校验是否为正浮点数
CheckUtil.isPlusFloat=function(str){
  var patrn=/^([+]?)\d*\.\d+$/;
  return patrn.test(str)||CheckUtil.isPlusInteger(str);
};
//校验是否为负浮点数
CheckUtil.isMinusFloat=function(str){
  var patrn=/^-\d*\.\d+$/;
  return patrn.test(str)||CheckUtil.isMinusInteger(str);
};
//校验是否仅中文
CheckUtil.isChinese=function(str){
var patrn=/[\u4E00-\u9FA5\uF900-\uFA2D]+$/;
return patrn.test(str);
};
//校验是否仅ACSII字符
CheckUtil.isAcsii=function(str){
var patrn=/^[\x00-\xFF]+$/;
return patrn.test(str);
};
//校验手机号码
CheckUtil.isMobile = function (str) {
var patrn = /^0?1(([3578][0-9]{1})|(59)){1}[0-9]{8}$/;
return patrn.test(str);
};
//校验电话号码
CheckUtil.isPhone = function (str) {
var patrn = /^(0[\d]{2,3}-)?\d{6,8}(-\d{3,4})?$/;
return patrn.test(str);
};
//校验URL地址
CheckUtil.isUrl=function(str){
var patrn= /^http[s]?:\/\/[\w-]+(\.[\w-]+)+([\w-\.\/?%&=]*)?$/;
return patrn.test(str);
};
//校验电邮地址
CheckUtil.isEmail = function (str) {
var patrn = /^[\w-]+@[\w-]+(\.[\w-]+)+$/;
return patrn.test(str);
};
//校验邮编
CheckUtil.isZipCode = function (str) {
var patrn = /^\d{6}$/;
return patrn.test(str);
};
//校验合法时间
CheckUtil.isDate = function (str) {
  if(!/\d{4}(\.|\/|\-)\d{1,2}(\.|\/|\-)\d{1,2}/.test(str)){
    return false;
  }
  var r = str.match(/\d{1,4}/g);
  if(r==null){return false;};
  var d= new Date(r[0], r[1]-1, r[2]);
  return (d.getFullYear()==r[0]&&(d.getMonth()+1)==r[1]&&d.getDate()==r[2]);
};
//校验字符串：只能输入6-20个字母、数字、下划线(常用于校验用户名和密码)
CheckUtil.isString6_20=function(str){
var patrn=/^(\w){6,20}$/;
return patrn.test(str);
};
//获取字符串的字节长度,中文占两位
CheckUtil.byteLength = function(value){
var length = value.length;    
for(var i = 0; i < value.length; i++){
	if(value.charCodeAt(i) > 127){length++;}
}
return length;
}