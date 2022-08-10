package currytwo

// TwoArgs curries the first 2 arguments from a function with 2 arguments and 0 return values.
func TwoArgs[CurriedArg1, CurriedArg2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2)) func() {
	return func() {
		f(c1, c2)
	}
}

// TwoArgsVariadic curries the first 2 arguments from a function with 2 plus variadic arguments and 0 return values.
func TwoArgsVariadic[CurriedArg1, CurriedArg2, Var any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, ...Var)) func(...Var) {
	return func(v ...Var) {
		f(c1, c2, v...)
	}
}

// TwoArgsOneRet curries the first 2 arguments from a function with 2 arguments and 1 return values.
func TwoArgsOneRet[CurriedArg1, CurriedArg2, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2) Ret1) func() Ret1 {
	return func() Ret1 {
		return f(c1, c2)
	}
}

// TwoArgsVariadicOneRet curries the first 2 arguments from a function with 2 plus variadic arguments and 1 return values.
func TwoArgsVariadicOneRet[CurriedArg1, CurriedArg2, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, ...Var) Ret1) func(...Var) Ret1 {
	return func(v ...Var) Ret1 {
		return f(c1, c2, v...)
	}
}

// TwoArgsTwoRets curries the first 2 arguments from a function with 2 arguments and 2 return values.
func TwoArgsTwoRets[CurriedArg1, CurriedArg2, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2) (Ret1, Ret2)) func() (Ret1, Ret2) {
	return func() (Ret1, Ret2) {
		return f(c1, c2)
	}
}

// TwoArgsVariadicTwoRets curries the first 2 arguments from a function with 2 plus variadic arguments and 2 return values.
func TwoArgsVariadicTwoRets[CurriedArg1, CurriedArg2, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, ...Var) (Ret1, Ret2)) func(...Var) (Ret1, Ret2) {
	return func(v ...Var) (Ret1, Ret2) {
		return f(c1, c2, v...)
	}
}

// TwoArgsThreeRets curries the first 2 arguments from a function with 2 arguments and 3 return values.
func TwoArgsThreeRets[CurriedArg1, CurriedArg2, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2) (Ret1, Ret2, Ret3)) func() (Ret1, Ret2, Ret3) {
	return func() (Ret1, Ret2, Ret3) {
		return f(c1, c2)
	}
}

// TwoArgsVariadicThreeRets curries the first 2 arguments from a function with 2 plus variadic arguments and 3 return values.
func TwoArgsVariadicThreeRets[CurriedArg1, CurriedArg2, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, ...Var) (Ret1, Ret2, Ret3)) func(...Var) (Ret1, Ret2, Ret3) {
	return func(v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, v...)
	}
}

// TwoArgsFourRets curries the first 2 arguments from a function with 2 arguments and 4 return values.
func TwoArgsFourRets[CurriedArg1, CurriedArg2, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2) (Ret1, Ret2, Ret3, Ret4)) func() (Ret1, Ret2, Ret3, Ret4) {
	return func() (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2)
	}
}

// TwoArgsVariadicFourRets curries the first 2 arguments from a function with 2 plus variadic arguments and 4 return values.
func TwoArgsVariadicFourRets[CurriedArg1, CurriedArg2, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, v...)
	}
}

// TwoArgsFiveRets curries the first 2 arguments from a function with 2 arguments and 5 return values.
func TwoArgsFiveRets[CurriedArg1, CurriedArg2, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2) (Ret1, Ret2, Ret3, Ret4, Ret5)) func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2)
	}
}

// TwoArgsVariadicFiveRets curries the first 2 arguments from a function with 2 plus variadic arguments and 5 return values.
func TwoArgsVariadicFiveRets[CurriedArg1, CurriedArg2, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, v...)
	}
}

// ThreeArgs curries the first 2 arguments from a function with 3 arguments and 0 return values.
func ThreeArgs[CurriedArg1, CurriedArg2, Arg1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1)) func(Arg1) {
	return func(a1 Arg1) {
		f(c1, c2, a1)
	}
}

// ThreeArgsVariadic curries the first 2 arguments from a function with 3 plus variadic arguments and 0 return values.
func ThreeArgsVariadic[CurriedArg1, CurriedArg2, Arg1, Var any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, ...Var)) func(Arg1, ...Var) {
	return func(a1 Arg1, v ...Var) {
		f(c1, c2, a1, v...)
	}
}

// ThreeArgsOneRet curries the first 2 arguments from a function with 3 arguments and 1 return values.
func ThreeArgsOneRet[CurriedArg1, CurriedArg2, Arg1, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1) Ret1) func(Arg1) Ret1 {
	return func(a1 Arg1) Ret1 {
		return f(c1, c2, a1)
	}
}

// ThreeArgsVariadicOneRet curries the first 2 arguments from a function with 3 plus variadic arguments and 1 return values.
func ThreeArgsVariadicOneRet[CurriedArg1, CurriedArg2, Arg1, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, ...Var) Ret1) func(Arg1, ...Var) Ret1 {
	return func(a1 Arg1, v ...Var) Ret1 {
		return f(c1, c2, a1, v...)
	}
}

// ThreeArgsTwoRets curries the first 2 arguments from a function with 3 arguments and 2 return values.
func ThreeArgsTwoRets[CurriedArg1, CurriedArg2, Arg1, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1) (Ret1, Ret2)) func(Arg1) (Ret1, Ret2) {
	return func(a1 Arg1) (Ret1, Ret2) {
		return f(c1, c2, a1)
	}
}

// ThreeArgsVariadicTwoRets curries the first 2 arguments from a function with 3 plus variadic arguments and 2 return values.
func ThreeArgsVariadicTwoRets[CurriedArg1, CurriedArg2, Arg1, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, ...Var) (Ret1, Ret2)) func(Arg1, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2) {
		return f(c1, c2, a1, v...)
	}
}

// ThreeArgsThreeRets curries the first 2 arguments from a function with 3 arguments and 3 return values.
func ThreeArgsThreeRets[CurriedArg1, CurriedArg2, Arg1, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1) (Ret1, Ret2, Ret3)) func(Arg1) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3) {
		return f(c1, c2, a1)
	}
}

// ThreeArgsVariadicThreeRets curries the first 2 arguments from a function with 3 plus variadic arguments and 3 return values.
func ThreeArgsVariadicThreeRets[CurriedArg1, CurriedArg2, Arg1, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, a1, v...)
	}
}

// ThreeArgsFourRets curries the first 2 arguments from a function with 3 arguments and 4 return values.
func ThreeArgsFourRets[CurriedArg1, CurriedArg2, Arg1, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1) (Ret1, Ret2, Ret3, Ret4)) func(Arg1) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, a1)
	}
}

// ThreeArgsVariadicFourRets curries the first 2 arguments from a function with 3 plus variadic arguments and 4 return values.
func ThreeArgsVariadicFourRets[CurriedArg1, CurriedArg2, Arg1, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, a1, v...)
	}
}

// ThreeArgsFiveRets curries the first 2 arguments from a function with 3 arguments and 5 return values.
func ThreeArgsFiveRets[CurriedArg1, CurriedArg2, Arg1, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, a1)
	}
}

// ThreeArgsVariadicFiveRets curries the first 2 arguments from a function with 3 plus variadic arguments and 5 return values.
func ThreeArgsVariadicFiveRets[CurriedArg1, CurriedArg2, Arg1, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, a1, v...)
	}
}

// FourArgs curries the first 2 arguments from a function with 4 arguments and 0 return values.
func FourArgs[CurriedArg1, CurriedArg2, Arg1, Arg2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2)) func(Arg1, Arg2) {
	return func(a1 Arg1, a2 Arg2) {
		f(c1, c2, a1, a2)
	}
}

// FourArgsVariadic curries the first 2 arguments from a function with 4 plus variadic arguments and 0 return values.
func FourArgsVariadic[CurriedArg1, CurriedArg2, Arg1, Arg2, Var any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, ...Var)) func(Arg1, Arg2, ...Var) {
	return func(a1 Arg1, a2 Arg2, v ...Var) {
		f(c1, c2, a1, a2, v...)
	}
}

// FourArgsOneRet curries the first 2 arguments from a function with 4 arguments and 1 return values.
func FourArgsOneRet[CurriedArg1, CurriedArg2, Arg1, Arg2, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2) Ret1) func(Arg1, Arg2) Ret1 {
	return func(a1 Arg1, a2 Arg2) Ret1 {
		return f(c1, c2, a1, a2)
	}
}

// FourArgsVariadicOneRet curries the first 2 arguments from a function with 4 plus variadic arguments and 1 return values.
func FourArgsVariadicOneRet[CurriedArg1, CurriedArg2, Arg1, Arg2, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, ...Var) Ret1) func(Arg1, Arg2, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, v ...Var) Ret1 {
		return f(c1, c2, a1, a2, v...)
	}
}

// FourArgsTwoRets curries the first 2 arguments from a function with 4 arguments and 2 return values.
func FourArgsTwoRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2) (Ret1, Ret2)) func(Arg1, Arg2) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2) {
		return f(c1, c2, a1, a2)
	}
}

// FourArgsVariadicTwoRets curries the first 2 arguments from a function with 4 plus variadic arguments and 2 return values.
func FourArgsVariadicTwoRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2) {
		return f(c1, c2, a1, a2, v...)
	}
}

// FourArgsThreeRets curries the first 2 arguments from a function with 4 arguments and 3 return values.
func FourArgsThreeRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2) (Ret1, Ret2, Ret3)) func(Arg1, Arg2) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3) {
		return f(c1, c2, a1, a2)
	}
}

// FourArgsVariadicThreeRets curries the first 2 arguments from a function with 4 plus variadic arguments and 3 return values.
func FourArgsVariadicThreeRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, a1, a2, v...)
	}
}

// FourArgsFourRets curries the first 2 arguments from a function with 4 arguments and 4 return values.
func FourArgsFourRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, a1, a2)
	}
}

// FourArgsVariadicFourRets curries the first 2 arguments from a function with 4 plus variadic arguments and 4 return values.
func FourArgsVariadicFourRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, a1, a2, v...)
	}
}

// FourArgsFiveRets curries the first 2 arguments from a function with 4 arguments and 5 return values.
func FourArgsFiveRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, a1, a2)
	}
}

// FourArgsVariadicFiveRets curries the first 2 arguments from a function with 4 plus variadic arguments and 5 return values.
func FourArgsVariadicFiveRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, a1, a2, v...)
	}
}

// FiveArgs curries the first 2 arguments from a function with 5 arguments and 0 return values.
func FiveArgs[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3)) func(Arg1, Arg2, Arg3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) {
		f(c1, c2, a1, a2, a3)
	}
}

// FiveArgsVariadic curries the first 2 arguments from a function with 5 plus variadic arguments and 0 return values.
func FiveArgsVariadic[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Var any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, ...Var)) func(Arg1, Arg2, Arg3, ...Var) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) {
		f(c1, c2, a1, a2, a3, v...)
	}
}

// FiveArgsOneRet curries the first 2 arguments from a function with 5 arguments and 1 return values.
func FiveArgsOneRet[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3) Ret1) func(Arg1, Arg2, Arg3) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) Ret1 {
		return f(c1, c2, a1, a2, a3)
	}
}

// FiveArgsVariadicOneRet curries the first 2 arguments from a function with 5 plus variadic arguments and 1 return values.
func FiveArgsVariadicOneRet[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, ...Var) Ret1) func(Arg1, Arg2, Arg3, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) Ret1 {
		return f(c1, c2, a1, a2, a3, v...)
	}
}

// FiveArgsTwoRets curries the first 2 arguments from a function with 5 arguments and 2 return values.
func FiveArgsTwoRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3) (Ret1, Ret2)) func(Arg1, Arg2, Arg3) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2) {
		return f(c1, c2, a1, a2, a3)
	}
}

// FiveArgsVariadicTwoRets curries the first 2 arguments from a function with 5 plus variadic arguments and 2 return values.
func FiveArgsVariadicTwoRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2) {
		return f(c1, c2, a1, a2, a3, v...)
	}
}

// FiveArgsThreeRets curries the first 2 arguments from a function with 5 arguments and 3 return values.
func FiveArgsThreeRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3) {
		return f(c1, c2, a1, a2, a3)
	}
}

// FiveArgsVariadicThreeRets curries the first 2 arguments from a function with 5 plus variadic arguments and 3 return values.
func FiveArgsVariadicThreeRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, a1, a2, a3, v...)
	}
}

// FiveArgsFourRets curries the first 2 arguments from a function with 5 arguments and 4 return values.
func FiveArgsFourRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, a1, a2, a3)
	}
}

// FiveArgsVariadicFourRets curries the first 2 arguments from a function with 5 plus variadic arguments and 4 return values.
func FiveArgsVariadicFourRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, a1, a2, a3, v...)
	}
}

// FiveArgsFiveRets curries the first 2 arguments from a function with 5 arguments and 5 return values.
func FiveArgsFiveRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, a1, a2, a3)
	}
}

// FiveArgsVariadicFiveRets curries the first 2 arguments from a function with 5 plus variadic arguments and 5 return values.
func FiveArgsVariadicFiveRets[CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, f func(CurriedArg1, CurriedArg2, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, a1, a2, a3, v...)
	}
}
