package currythree

// ThreeArgs curries the first 3 arguments from a function with 3 arguments and 0 return values.
func ThreeArgs[CurriedArg1, CurriedArg2, CurriedArg3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3)) func() {
	return func() {
		f(c1, c2, c3)
	}
}

// ThreeArgsVariadic curries the first 3 arguments from a function with 3 plus variadic arguments and 0 return values.
func ThreeArgsVariadic[CurriedArg1, CurriedArg2, CurriedArg3, Var any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var)) func(...Var) {
	return func(v ...Var) {
		f(c1, c2, c3, v...)
	}
}

// ThreeArgsOneRet curries the first 3 arguments from a function with 3 arguments and 1 return values.
func ThreeArgsOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) Ret1) func() Ret1 {
	return func() Ret1 {
		return f(c1, c2, c3)
	}
}

// ThreeArgsVariadicOneRet curries the first 3 arguments from a function with 3 plus variadic arguments and 1 return values.
func ThreeArgsVariadicOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) Ret1) func(...Var) Ret1 {
	return func(v ...Var) Ret1 {
		return f(c1, c2, c3, v...)
	}
}

// ThreeArgsTwoRets curries the first 3 arguments from a function with 3 arguments and 2 return values.
func ThreeArgsTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) (Ret1, Ret2)) func() (Ret1, Ret2) {
	return func() (Ret1, Ret2) {
		return f(c1, c2, c3)
	}
}

// ThreeArgsVariadicTwoRets curries the first 3 arguments from a function with 3 plus variadic arguments and 2 return values.
func ThreeArgsVariadicTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) (Ret1, Ret2)) func(...Var) (Ret1, Ret2) {
	return func(v ...Var) (Ret1, Ret2) {
		return f(c1, c2, c3, v...)
	}
}

// ThreeArgsThreeRets curries the first 3 arguments from a function with 3 arguments and 3 return values.
func ThreeArgsThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) (Ret1, Ret2, Ret3)) func() (Ret1, Ret2, Ret3) {
	return func() (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3)
	}
}

// ThreeArgsVariadicThreeRets curries the first 3 arguments from a function with 3 plus variadic arguments and 3 return values.
func ThreeArgsVariadicThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) (Ret1, Ret2, Ret3)) func(...Var) (Ret1, Ret2, Ret3) {
	return func(v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, v...)
	}
}

// ThreeArgsFourRets curries the first 3 arguments from a function with 3 arguments and 4 return values.
func ThreeArgsFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) (Ret1, Ret2, Ret3, Ret4)) func() (Ret1, Ret2, Ret3, Ret4) {
	return func() (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3)
	}
}

// ThreeArgsVariadicFourRets curries the first 3 arguments from a function with 3 plus variadic arguments and 4 return values.
func ThreeArgsVariadicFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, v...)
	}
}

// ThreeArgsFiveRets curries the first 3 arguments from a function with 3 arguments and 5 return values.
func ThreeArgsFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) (Ret1, Ret2, Ret3, Ret4, Ret5)) func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3)
	}
}

// ThreeArgsVariadicFiveRets curries the first 3 arguments from a function with 3 plus variadic arguments and 5 return values.
func ThreeArgsVariadicFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, v...)
	}
}

// FourArgs curries the first 3 arguments from a function with 4 arguments and 0 return values.
func FourArgs[CurriedArg1, CurriedArg2, CurriedArg3, Arg1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1)) func(Arg1) {
	return func(a1 Arg1) {
		f(c1, c2, c3, a1)
	}
}

// FourArgsVariadic curries the first 3 arguments from a function with 4 plus variadic arguments and 0 return values.
func FourArgsVariadic[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var)) func(Arg1, ...Var) {
	return func(a1 Arg1, v ...Var) {
		f(c1, c2, c3, a1, v...)
	}
}

// FourArgsOneRet curries the first 3 arguments from a function with 4 arguments and 1 return values.
func FourArgsOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) Ret1) func(Arg1) Ret1 {
	return func(a1 Arg1) Ret1 {
		return f(c1, c2, c3, a1)
	}
}

// FourArgsVariadicOneRet curries the first 3 arguments from a function with 4 plus variadic arguments and 1 return values.
func FourArgsVariadicOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) Ret1) func(Arg1, ...Var) Ret1 {
	return func(a1 Arg1, v ...Var) Ret1 {
		return f(c1, c2, c3, a1, v...)
	}
}

// FourArgsTwoRets curries the first 3 arguments from a function with 4 arguments and 2 return values.
func FourArgsTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) (Ret1, Ret2)) func(Arg1) (Ret1, Ret2) {
	return func(a1 Arg1) (Ret1, Ret2) {
		return f(c1, c2, c3, a1)
	}
}

// FourArgsVariadicTwoRets curries the first 3 arguments from a function with 4 plus variadic arguments and 2 return values.
func FourArgsVariadicTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) (Ret1, Ret2)) func(Arg1, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2) {
		return f(c1, c2, c3, a1, v...)
	}
}

// FourArgsThreeRets curries the first 3 arguments from a function with 4 arguments and 3 return values.
func FourArgsThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) (Ret1, Ret2, Ret3)) func(Arg1) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, a1)
	}
}

// FourArgsVariadicThreeRets curries the first 3 arguments from a function with 4 plus variadic arguments and 3 return values.
func FourArgsVariadicThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, a1, v...)
	}
}

// FourArgsFourRets curries the first 3 arguments from a function with 4 arguments and 4 return values.
func FourArgsFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) (Ret1, Ret2, Ret3, Ret4)) func(Arg1) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, a1)
	}
}

// FourArgsVariadicFourRets curries the first 3 arguments from a function with 4 plus variadic arguments and 4 return values.
func FourArgsVariadicFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, a1, v...)
	}
}

// FourArgsFiveRets curries the first 3 arguments from a function with 4 arguments and 5 return values.
func FourArgsFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, a1)
	}
}

// FourArgsVariadicFiveRets curries the first 3 arguments from a function with 4 plus variadic arguments and 5 return values.
func FourArgsVariadicFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, a1, v...)
	}
}

// FiveArgs curries the first 3 arguments from a function with 5 arguments and 0 return values.
func FiveArgs[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2)) func(Arg1, Arg2) {
	return func(a1 Arg1, a2 Arg2) {
		f(c1, c2, c3, a1, a2)
	}
}

// FiveArgsVariadic curries the first 3 arguments from a function with 5 plus variadic arguments and 0 return values.
func FiveArgsVariadic[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var)) func(Arg1, Arg2, ...Var) {
	return func(a1 Arg1, a2 Arg2, v ...Var) {
		f(c1, c2, c3, a1, a2, v...)
	}
}

// FiveArgsOneRet curries the first 3 arguments from a function with 5 arguments and 1 return values.
func FiveArgsOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) Ret1) func(Arg1, Arg2) Ret1 {
	return func(a1 Arg1, a2 Arg2) Ret1 {
		return f(c1, c2, c3, a1, a2)
	}
}

// FiveArgsVariadicOneRet curries the first 3 arguments from a function with 5 plus variadic arguments and 1 return values.
func FiveArgsVariadicOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) Ret1) func(Arg1, Arg2, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, v ...Var) Ret1 {
		return f(c1, c2, c3, a1, a2, v...)
	}
}

// FiveArgsTwoRets curries the first 3 arguments from a function with 5 arguments and 2 return values.
func FiveArgsTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) (Ret1, Ret2)) func(Arg1, Arg2) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2) {
		return f(c1, c2, c3, a1, a2)
	}
}

// FiveArgsVariadicTwoRets curries the first 3 arguments from a function with 5 plus variadic arguments and 2 return values.
func FiveArgsVariadicTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2) {
		return f(c1, c2, c3, a1, a2, v...)
	}
}

// FiveArgsThreeRets curries the first 3 arguments from a function with 5 arguments and 3 return values.
func FiveArgsThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) (Ret1, Ret2, Ret3)) func(Arg1, Arg2) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, a1, a2)
	}
}

// FiveArgsVariadicThreeRets curries the first 3 arguments from a function with 5 plus variadic arguments and 3 return values.
func FiveArgsVariadicThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, a1, a2, v...)
	}
}

// FiveArgsFourRets curries the first 3 arguments from a function with 5 arguments and 4 return values.
func FiveArgsFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, a1, a2)
	}
}

// FiveArgsVariadicFourRets curries the first 3 arguments from a function with 5 plus variadic arguments and 4 return values.
func FiveArgsVariadicFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, a1, a2, v...)
	}
}

// FiveArgsFiveRets curries the first 3 arguments from a function with 5 arguments and 5 return values.
func FiveArgsFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, a1, a2)
	}
}

// FiveArgsVariadicFiveRets curries the first 3 arguments from a function with 5 plus variadic arguments and 5 return values.
func FiveArgsVariadicFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, a1, a2, v...)
	}
}
