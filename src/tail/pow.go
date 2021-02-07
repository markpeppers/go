func myPow(x float64, n int) float64 {
    return myPowHelper(x, 1, n)
}

func myPowHelper(x float64, acc float64, n int) float64 {
    if n == 0 {
        return 1 * acc
    }
    var inc int
    if n > 0 {
        acc = acc * x
        inc = -1
    } else {
        acc = acc / x
        inc = 1
    }
    return(myPowHelper(x, acc, n + inc))
}
