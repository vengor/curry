package curryone

// OneArg curries the first 1 arguments from a function with 1 arguments and 0 return values.
func OneArg[CurriedArg1 any](c1 CurriedArg1, f func(CurriedArg1)) func() {
	return func() {
		f(c1)
	}
}

// OneArgVariadic curries the first 1 arguments from a function with 1 plus variadic arguments and 0 return values.
func OneArgVariadic[CurriedArg1, Var any](c1 CurriedArg1, f func(CurriedArg1, ...Var)) func(...Var) {
	return func(v ...Var) {
		f(c1, v...)
	}
}

// OneArgOneRet curries the first 1 arguments from a function with 1 arguments and 1 return values.
func OneArgOneRet[CurriedArg1, Ret1 any](c1 CurriedArg1, f func(CurriedArg1) Ret1) func() Ret1 {
	return func() Ret1 {
		return f(c1)
	}
}

// OneArgVariadicOneRet curries the first 1 arguments from a function with 1 plus variadic arguments and 1 return values.
func OneArgVariadicOneRet[CurriedArg1, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) Ret1) func(...Var) Ret1 {
	return func(v ...Var) Ret1 {
		return f(c1, v...)
	}
}

// OneArgTwoRets curries the first 1 arguments from a function with 1 arguments and 2 return values.
func OneArgTwoRets[CurriedArg1, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1) (Ret1, Ret2)) func() (Ret1, Ret2) {
	return func() (Ret1, Ret2) {
		return f(c1)
	}
}

// OneArgVariadicTwoRets curries the first 1 arguments from a function with 1 plus variadic arguments and 2 return values.
func OneArgVariadicTwoRets[CurriedArg1, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) (Ret1, Ret2)) func(...Var) (Ret1, Ret2) {
	return func(v ...Var) (Ret1, Ret2) {
		return f(c1, v...)
	}
}

// OneArgThreeRets curries the first 1 arguments from a function with 1 arguments and 3 return values.
func OneArgThreeRets[CurriedArg1, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1) (Ret1, Ret2, Ret3)) func() (Ret1, Ret2, Ret3) {
	return func() (Ret1, Ret2, Ret3) {
		return f(c1)
	}
}

// OneArgVariadicThreeRets curries the first 1 arguments from a function with 1 plus variadic arguments and 3 return values.
func OneArgVariadicThreeRets[CurriedArg1, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) (Ret1, Ret2, Ret3)) func(...Var) (Ret1, Ret2, Ret3) {
	return func(v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, v...)
	}
}

// OneArgFourRets curries the first 1 arguments from a function with 1 arguments and 4 return values.
func OneArgFourRets[CurriedArg1, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1) (Ret1, Ret2, Ret3, Ret4)) func() (Ret1, Ret2, Ret3, Ret4) {
	return func() (Ret1, Ret2, Ret3, Ret4) {
		return f(c1)
	}
}

// OneArgVariadicFourRets curries the first 1 arguments from a function with 1 plus variadic arguments and 4 return values.
func OneArgVariadicFourRets[CurriedArg1, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, v...)
	}
}

// OneArgFiveRets curries the first 1 arguments from a function with 1 arguments and 5 return values.
func OneArgFiveRets[CurriedArg1, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1) (Ret1, Ret2, Ret3, Ret4, Ret5)) func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1)
	}
}

// OneArgVariadicFiveRets curries the first 1 arguments from a function with 1 plus variadic arguments and 5 return values.
func OneArgVariadicFiveRets[CurriedArg1, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, v...)
	}
}

// TwoArgs curries the first 1 arguments from a function with 2 arguments and 0 return values.
func TwoArgs[CurriedArg1, Arg1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1)) func(Arg1) {
	return func(a1 Arg1) {
		f(c1, a1)
	}
}

// TwoArgsVariadic curries the first 1 arguments from a function with 2 plus variadic arguments and 0 return values.
func TwoArgsVariadic[CurriedArg1, Arg1, Var any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var)) func(Arg1, ...Var) {
	return func(a1 Arg1, v ...Var) {
		f(c1, a1, v...)
	}
}

// TwoArgsOneRet curries the first 1 arguments from a function with 2 arguments and 1 return values.
func TwoArgsOneRet[CurriedArg1, Arg1, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) Ret1) func(Arg1) Ret1 {
	return func(a1 Arg1) Ret1 {
		return f(c1, a1)
	}
}

// TwoArgsVariadicOneRet curries the first 1 arguments from a function with 2 plus variadic arguments and 1 return values.
func TwoArgsVariadicOneRet[CurriedArg1, Arg1, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) Ret1) func(Arg1, ...Var) Ret1 {
	return func(a1 Arg1, v ...Var) Ret1 {
		return f(c1, a1, v...)
	}
}

// TwoArgsTwoRets curries the first 1 arguments from a function with 2 arguments and 2 return values.
func TwoArgsTwoRets[CurriedArg1, Arg1, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) (Ret1, Ret2)) func(Arg1) (Ret1, Ret2) {
	return func(a1 Arg1) (Ret1, Ret2) {
		return f(c1, a1)
	}
}

// TwoArgsVariadicTwoRets curries the first 1 arguments from a function with 2 plus variadic arguments and 2 return values.
func TwoArgsVariadicTwoRets[CurriedArg1, Arg1, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) (Ret1, Ret2)) func(Arg1, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2) {
		return f(c1, a1, v...)
	}
}

// TwoArgsThreeRets curries the first 1 arguments from a function with 2 arguments and 3 return values.
func TwoArgsThreeRets[CurriedArg1, Arg1, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) (Ret1, Ret2, Ret3)) func(Arg1) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3) {
		return f(c1, a1)
	}
}

// TwoArgsVariadicThreeRets curries the first 1 arguments from a function with 2 plus variadic arguments and 3 return values.
func TwoArgsVariadicThreeRets[CurriedArg1, Arg1, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, a1, v...)
	}
}

// TwoArgsFourRets curries the first 1 arguments from a function with 2 arguments and 4 return values.
func TwoArgsFourRets[CurriedArg1, Arg1, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) (Ret1, Ret2, Ret3, Ret4)) func(Arg1) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1)
	}
}

// TwoArgsVariadicFourRets curries the first 1 arguments from a function with 2 plus variadic arguments and 4 return values.
func TwoArgsVariadicFourRets[CurriedArg1, Arg1, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, v...)
	}
}

// TwoArgsFiveRets curries the first 1 arguments from a function with 2 arguments and 5 return values.
func TwoArgsFiveRets[CurriedArg1, Arg1, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1)
	}
}

// TwoArgsVariadicFiveRets curries the first 1 arguments from a function with 2 plus variadic arguments and 5 return values.
func TwoArgsVariadicFiveRets[CurriedArg1, Arg1, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, v...)
	}
}

// ThreeArgs curries the first 1 arguments from a function with 3 arguments and 0 return values.
func ThreeArgs[CurriedArg1, Arg1, Arg2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2)) func(Arg1, Arg2) {
	return func(a1 Arg1, a2 Arg2) {
		f(c1, a1, a2)
	}
}

// ThreeArgsVariadic curries the first 1 arguments from a function with 3 plus variadic arguments and 0 return values.
func ThreeArgsVariadic[CurriedArg1, Arg1, Arg2, Var any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var)) func(Arg1, Arg2, ...Var) {
	return func(a1 Arg1, a2 Arg2, v ...Var) {
		f(c1, a1, a2, v...)
	}
}

// ThreeArgsOneRet curries the first 1 arguments from a function with 3 arguments and 1 return values.
func ThreeArgsOneRet[CurriedArg1, Arg1, Arg2, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) Ret1) func(Arg1, Arg2) Ret1 {
	return func(a1 Arg1, a2 Arg2) Ret1 {
		return f(c1, a1, a2)
	}
}

// ThreeArgsVariadicOneRet curries the first 1 arguments from a function with 3 plus variadic arguments and 1 return values.
func ThreeArgsVariadicOneRet[CurriedArg1, Arg1, Arg2, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) Ret1) func(Arg1, Arg2, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, v ...Var) Ret1 {
		return f(c1, a1, a2, v...)
	}
}

// ThreeArgsTwoRets curries the first 1 arguments from a function with 3 arguments and 2 return values.
func ThreeArgsTwoRets[CurriedArg1, Arg1, Arg2, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) (Ret1, Ret2)) func(Arg1, Arg2) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2) {
		return f(c1, a1, a2)
	}
}

// ThreeArgsVariadicTwoRets curries the first 1 arguments from a function with 3 plus variadic arguments and 2 return values.
func ThreeArgsVariadicTwoRets[CurriedArg1, Arg1, Arg2, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2) {
		return f(c1, a1, a2, v...)
	}
}

// ThreeArgsThreeRets curries the first 1 arguments from a function with 3 arguments and 3 return values.
func ThreeArgsThreeRets[CurriedArg1, Arg1, Arg2, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) (Ret1, Ret2, Ret3)) func(Arg1, Arg2) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2)
	}
}

// ThreeArgsVariadicThreeRets curries the first 1 arguments from a function with 3 plus variadic arguments and 3 return values.
func ThreeArgsVariadicThreeRets[CurriedArg1, Arg1, Arg2, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, v...)
	}
}

// ThreeArgsFourRets curries the first 1 arguments from a function with 3 arguments and 4 return values.
func ThreeArgsFourRets[CurriedArg1, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2)
	}
}

// ThreeArgsVariadicFourRets curries the first 1 arguments from a function with 3 plus variadic arguments and 4 return values.
func ThreeArgsVariadicFourRets[CurriedArg1, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, v...)
	}
}

// ThreeArgsFiveRets curries the first 1 arguments from a function with 3 arguments and 5 return values.
func ThreeArgsFiveRets[CurriedArg1, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2)
	}
}

// ThreeArgsVariadicFiveRets curries the first 1 arguments from a function with 3 plus variadic arguments and 5 return values.
func ThreeArgsVariadicFiveRets[CurriedArg1, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, v...)
	}
}

// FourArgs curries the first 1 arguments from a function with 4 arguments and 0 return values.
func FourArgs[CurriedArg1, Arg1, Arg2, Arg3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3)) func(Arg1, Arg2, Arg3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) {
		f(c1, a1, a2, a3)
	}
}

// FourArgsVariadic curries the first 1 arguments from a function with 4 plus variadic arguments and 0 return values.
func FourArgsVariadic[CurriedArg1, Arg1, Arg2, Arg3, Var any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var)) func(Arg1, Arg2, Arg3, ...Var) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) {
		f(c1, a1, a2, a3, v...)
	}
}

// FourArgsOneRet curries the first 1 arguments from a function with 4 arguments and 1 return values.
func FourArgsOneRet[CurriedArg1, Arg1, Arg2, Arg3, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) Ret1) func(Arg1, Arg2, Arg3) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) Ret1 {
		return f(c1, a1, a2, a3)
	}
}

// FourArgsVariadicOneRet curries the first 1 arguments from a function with 4 plus variadic arguments and 1 return values.
func FourArgsVariadicOneRet[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) Ret1) func(Arg1, Arg2, Arg3, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) Ret1 {
		return f(c1, a1, a2, a3, v...)
	}
}

// FourArgsTwoRets curries the first 1 arguments from a function with 4 arguments and 2 return values.
func FourArgsTwoRets[CurriedArg1, Arg1, Arg2, Arg3, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) (Ret1, Ret2)) func(Arg1, Arg2, Arg3) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2) {
		return f(c1, a1, a2, a3)
	}
}

// FourArgsVariadicTwoRets curries the first 1 arguments from a function with 4 plus variadic arguments and 2 return values.
func FourArgsVariadicTwoRets[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2) {
		return f(c1, a1, a2, a3, v...)
	}
}

// FourArgsThreeRets curries the first 1 arguments from a function with 4 arguments and 3 return values.
func FourArgsThreeRets[CurriedArg1, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, a3)
	}
}

// FourArgsVariadicThreeRets curries the first 1 arguments from a function with 4 plus variadic arguments and 3 return values.
func FourArgsVariadicThreeRets[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, a3, v...)
	}
}

// FourArgsFourRets curries the first 1 arguments from a function with 4 arguments and 4 return values.
func FourArgsFourRets[CurriedArg1, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, a3)
	}
}

// FourArgsVariadicFourRets curries the first 1 arguments from a function with 4 plus variadic arguments and 4 return values.
func FourArgsVariadicFourRets[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, a3, v...)
	}
}

// FourArgsFiveRets curries the first 1 arguments from a function with 4 arguments and 5 return values.
func FourArgsFiveRets[CurriedArg1, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, a3)
	}
}

// FourArgsVariadicFiveRets curries the first 1 arguments from a function with 4 plus variadic arguments and 5 return values.
func FourArgsVariadicFiveRets[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, a3, v...)
	}
}

// FiveArgs curries the first 1 arguments from a function with 5 arguments and 0 return values.
func FiveArgs[CurriedArg1, Arg1, Arg2, Arg3, Arg4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4)) func(Arg1, Arg2, Arg3, Arg4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) {
		f(c1, a1, a2, a3, a4)
	}
}

// FiveArgsVariadic curries the first 1 arguments from a function with 5 plus variadic arguments and 0 return values.
func FiveArgsVariadic[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var)) func(Arg1, Arg2, Arg3, Arg4, ...Var) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) {
		f(c1, a1, a2, a3, a4, v...)
	}
}

// FiveArgsOneRet curries the first 1 arguments from a function with 5 arguments and 1 return values.
func FiveArgsOneRet[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) Ret1) func(Arg1, Arg2, Arg3, Arg4) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) Ret1 {
		return f(c1, a1, a2, a3, a4)
	}
}

// FiveArgsVariadicOneRet curries the first 1 arguments from a function with 5 plus variadic arguments and 1 return values.
func FiveArgsVariadicOneRet[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) Ret1) func(Arg1, Arg2, Arg3, Arg4, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) Ret1 {
		return f(c1, a1, a2, a3, a4, v...)
	}
}

// FiveArgsTwoRets curries the first 1 arguments from a function with 5 arguments and 2 return values.
func FiveArgsTwoRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2)) func(Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) (Ret1, Ret2) {
		return f(c1, a1, a2, a3, a4)
	}
}

// FiveArgsVariadicTwoRets curries the first 1 arguments from a function with 5 plus variadic arguments and 2 return values.
func FiveArgsVariadicTwoRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) (Ret1, Ret2) {
		return f(c1, a1, a2, a3, a4, v...)
	}
}

// FiveArgsThreeRets curries the first 1 arguments from a function with 5 arguments and 3 return values.
func FiveArgsThreeRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, a3, a4)
	}
}

// FiveArgsVariadicThreeRets curries the first 1 arguments from a function with 5 plus variadic arguments and 3 return values.
func FiveArgsVariadicThreeRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, a3, a4, v...)
	}
}

// FiveArgsFourRets curries the first 1 arguments from a function with 5 arguments and 4 return values.
func FiveArgsFourRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, a3, a4)
	}
}

// FiveArgsVariadicFourRets curries the first 1 arguments from a function with 5 plus variadic arguments and 4 return values.
func FiveArgsVariadicFourRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, a3, a4, v...)
	}
}

// FiveArgsFiveRets curries the first 1 arguments from a function with 5 arguments and 5 return values.
func FiveArgsFiveRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, a3, a4)
	}
}

// FiveArgsVariadicFiveRets curries the first 1 arguments from a function with 5 plus variadic arguments and 5 return values.
func FiveArgsVariadicFiveRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, a3, a4, v...)
	}
}
