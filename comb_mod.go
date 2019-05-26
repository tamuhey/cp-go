package combmod

const MAX = 1000000
const MOD = int(1e9 + 7)

var fac [MAX]int
var inv [MAX]int
var facinv [MAX]int
var mod int

func init() {
	fac[0] = 1
	fac[1] = 1
	inv[0] = 1
	inv[1] = 1
	facinv[0] = 1
	facinv[1] = 1

	for i := 2; i < MAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		facinv[i] = facinv[i-1] * inv[i] % MOD
	}
}

func combmod(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (facinv[k] * facinv[n-k] % MOD) % MOD
}
