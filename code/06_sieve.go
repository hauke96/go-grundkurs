func main() {
	var compound [100]bool
	for i := 2; i*i < len(compound); i++ {
		if !compound[i] {
			for j := i * i; j < len(compound); j += i {
				compound[j] = true
			}
		}
	}
	var primes []int
	for i := 2; i < len(compound); i++ {
		if !compound[i] {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes)
}
