package ddm

// 手机号 132****7986
type Mobile string

// 银行卡号 622888******5676
type BankCard string

// 身份证号 1******7
type IDCard string

// 姓名 *鸿章
// TODO:参考 https://blog.thinkeridea.com/201910/go/efficient_string_truncation.html
// Deprecated:有更好的性能选择
type IDName string

// 密码 ******
type PassWord string

// 邮箱 l***w@gmail.com
type Email string
