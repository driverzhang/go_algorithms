package string

/*
 * String2Int
 * string => int
 * "1234" => 1234
 */
func String2Int(s string) int {
	var n int
	for _, v := range s {
		v -= '0' // 字符'1' - 字符'0' 对应 ascii码对应十进制为 49 - 48 = 1
		if v > 9 {
			return 0
		}
		n = n*10 + int(v)
	}
	return n
}

/*
 * ReverseString
 * string 反转输出
 * "A man, a plan, a canal: Panama" => "amanaP :lanac a ,nalp a ,nam A"
 */
func ReverseString(s string) string {
	var b []byte
	// n := len(s) // len(string) = the number of bytes
	// for i := 0; i < n; i++ {
	// 	fmt.Printf("%d %d\n", i, s[i])
	// }
	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}

	return string(b)
}

// 如果有中文字符呢?
// golang中string底层是通过byte数组实现的。中文字符在 unicode 下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
// byte 等同于int8，常用来处理ascii字符; 只能操作简单的字符，不支持中文操作。
// rune 等同于int32,常用来处理unicode或utf-8字符; 能操作任何字符。
func ReverseAllString(s string) string {
	var b []rune
	r := []rune(s)

	// u := utf8.RuneCountInString(s)
	// log.Error(u, len(r), len(s))
	for i := len(r) - 1; i >= 0; i-- {
		b = append(b, r[i])
	}

	return string(b)
}
