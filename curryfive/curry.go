package curryfive

// FiveArgs curries the first 5 arguments from a function with 5 arguments and 0 return values.
func FiveArgs[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5)) func() {
	return func() {
		f(c1, c2, c3, c4, c5)
	}
}

// FiveArgsVariadic curries the first 5 arguments from a function with 5 plus variadic arguments and 0 return values.
func FiveArgsVariadic[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var)) func(...Var) {
	return func(v ...Var) {
		f(c1, c2, c3, c4, c5, v...)
	}
}

// FiveArgsOneRet curries the first 5 arguments from a function with 5 arguments and 1 return values.
func FiveArgsOneRet[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) Ret1) func() Ret1 {
	return func() Ret1 {
		return f(c1, c2, c3, c4, c5)
	}
}

// FiveArgsVariadicOneRet curries the first 5 arguments from a function with 5 plus variadic arguments and 1 return values.
func FiveArgsVariadicOneRet[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) Ret1) func(...Var) Ret1 {
	return func(v ...Var) Ret1 {
		return f(c1, c2, c3, c4, c5, v...)
	}
}

// FiveArgsTwoRets curries the first 5 arguments from a function with 5 arguments and 2 return values.
func FiveArgsTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) (Ret1, Ret2)) func() (Ret1, Ret2) {
	return func() (Ret1, Ret2) {
		return f(c1, c2, c3, c4, c5)
	}
}

// FiveArgsVariadicTwoRets curries the first 5 arguments from a function with 5 plus variadic arguments and 2 return values.
func FiveArgsVariadicTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) (Ret1, Ret2)) func(...Var) (Ret1, Ret2) {
	return func(v ...Var) (Ret1, Ret2) {
		return f(c1, c2, c3, c4, c5, v...)
	}
}

// FiveArgsThreeRets curries the first 5 arguments from a function with 5 arguments and 3 return values.
func FiveArgsThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) (Ret1, Ret2, Ret3)) func() (Ret1, Ret2, Ret3) {
	return func() (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, c4, c5)
	}
}

// FiveArgsVariadicThreeRets curries the first 5 arguments from a function with 5 plus variadic arguments and 3 return values.
func FiveArgsVariadicThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) (Ret1, Ret2, Ret3)) func(...Var) (Ret1, Ret2, Ret3) {
	return func(v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, c4, c5, v...)
	}
}

// FiveArgsFourRets curries the first 5 arguments from a function with 5 arguments and 4 return values.
func FiveArgsFourRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) (Ret1, Ret2, Ret3, Ret4)) func() (Ret1, Ret2, Ret3, Ret4) {
	return func() (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, c4, c5)
	}
}

// FiveArgsVariadicFourRets curries the first 5 arguments from a function with 5 plus variadic arguments and 4 return values.
func FiveArgsVariadicFourRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, c4, c5, v...)
	}
}

// FiveArgsFiveRets curries the first 5 arguments from a function with 5 arguments and 5 return values.
func FiveArgsFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) (Ret1, Ret2, Ret3, Ret4, Ret5)) func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, c4, c5)
	}
}

// FiveArgsVariadicFiveRets curries the first 5 arguments from a function with 5 plus variadic arguments and 5 return values.
func FiveArgsVariadicFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, c4, c5, v...)
	}
}
