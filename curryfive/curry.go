package curryfive

// FromFiveArgs curries the first 5 arguments from a function with 5 arguments and 0 return values.
func FromFiveArgs[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5)) func() {
	return func() {
		f(c1, c2, c3, c4, c5)
	}
}

// FromFiveArgsVariadic curries the first 5 arguments from a function with 5 plus variadic arguments and 0 return values.
func FromFiveArgsVariadic[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var)) func(...Var) {
	return func(v ...Var) {
		f(c1, c2, c3, c4, c5, v...)
	}
}

// FromFiveArgsOneRet curries the first 5 arguments from a function with 5 arguments and 1 return values.
func FromFiveArgsOneRet[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) Ret1) func() Ret1 {
	return func() Ret1 {
		return f(c1, c2, c3, c4, c5)
	}
}

// FromFiveArgsVariadicOneRet curries the first 5 arguments from a function with 5 plus variadic arguments and 1 return values.
func FromFiveArgsVariadicOneRet[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) Ret1) func(...Var) Ret1 {
	return func(v ...Var) Ret1 {
		return f(c1, c2, c3, c4, c5, v...)
	}
}

// FromFiveArgsTwoRets curries the first 5 arguments from a function with 5 arguments and 2 return values.
func FromFiveArgsTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) (Ret1, Ret2)) func() (Ret1, Ret2) {
	return func() (Ret1, Ret2) {
		return f(c1, c2, c3, c4, c5)
	}
}

// FromFiveArgsVariadicTwoRets curries the first 5 arguments from a function with 5 plus variadic arguments and 2 return values.
func FromFiveArgsVariadicTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) (Ret1, Ret2)) func(...Var) (Ret1, Ret2) {
	return func(v ...Var) (Ret1, Ret2) {
		return f(c1, c2, c3, c4, c5, v...)
	}
}

// FromFiveArgsThreeRets curries the first 5 arguments from a function with 5 arguments and 3 return values.
func FromFiveArgsThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) (Ret1, Ret2, Ret3)) func() (Ret1, Ret2, Ret3) {
	return func() (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, c4, c5)
	}
}

// FromFiveArgsVariadicThreeRets curries the first 5 arguments from a function with 5 plus variadic arguments and 3 return values.
func FromFiveArgsVariadicThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) (Ret1, Ret2, Ret3)) func(...Var) (Ret1, Ret2, Ret3) {
	return func(v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, c4, c5, v...)
	}
}

// FromFiveArgsFourRets curries the first 5 arguments from a function with 5 arguments and 4 return values.
func FromFiveArgsFourRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) (Ret1, Ret2, Ret3, Ret4)) func() (Ret1, Ret2, Ret3, Ret4) {
	return func() (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, c4, c5)
	}
}

// FromFiveArgsVariadicFourRets curries the first 5 arguments from a function with 5 plus variadic arguments and 4 return values.
func FromFiveArgsVariadicFourRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, c4, c5, v...)
	}
}

// FromFiveArgsFiveRets curries the first 5 arguments from a function with 5 arguments and 5 return values.
func FromFiveArgsFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5) (Ret1, Ret2, Ret3, Ret4, Ret5)) func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, c4, c5)
	}
}

// FromFiveArgsVariadicFiveRets curries the first 5 arguments from a function with 5 plus variadic arguments and 5 return values.
func FromFiveArgsVariadicFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, c4 CurriedArg4, c5 CurriedArg5, f func(CurriedArg1, CurriedArg2, CurriedArg3, CurriedArg4, CurriedArg5, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, c4, c5, v...)
	}
}
