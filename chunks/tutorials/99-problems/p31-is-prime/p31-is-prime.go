package isprime

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := num - 1; i > 1; i-- {
		if num%i == 0 {
			return false
		}
	}

	return true
}
