func multiply(num1 string, num2 string) string {
	bignum1, _ := new(big.Int).SetString(num1, 0)
	bignum2, _ := new(big.Int).SetString(num2, 0)
	res := bignum1.Mul(bignum1, bignum2)
    return res.String() 
}