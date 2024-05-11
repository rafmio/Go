package isprime

func IsPrime(num int) bool {
	for i := num - 1; i > 0; i-- {
		if num%i == 0 {
			return false
		}
	}

	return true
}
