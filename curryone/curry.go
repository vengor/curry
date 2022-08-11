package curryone

// FromOneArg curries the first 1 arguments from a function with 1 arguments and 0 return values.
func FromOneArg[CurriedArg1 any](c1 CurriedArg1, f func(CurriedArg1)) func() {
	return func() {
		f(c1)
	}
}

// FromOneArgVariadic curries the first 1 arguments from a function with 1 plus variadic arguments and 0 return values.
func FromOneArgVariadic[CurriedArg1, Var any](c1 CurriedArg1, f func(CurriedArg1, ...Var)) func(...Var) {
	return func(v ...Var) {
		f(c1, v...)
	}
}

// FromOneArgOneRet curries the first 1 arguments from a function with 1 arguments and 1 return values.
func FromOneArgOneRet[CurriedArg1, Ret1 any](c1 CurriedArg1, f func(CurriedArg1) Ret1) func() Ret1 {
	return func() Ret1 {
		return f(c1)
	}
}

// FromOneArgVariadicOneRet curries the first 1 arguments from a function with 1 plus variadic arguments and 1 return values.
func FromOneArgVariadicOneRet[CurriedArg1, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) Ret1) func(...Var) Ret1 {
	return func(v ...Var) Ret1 {
		return f(c1, v...)
	}
}

// FromOneArgTwoRets curries the first 1 arguments from a function with 1 arguments and 2 return values.
func FromOneArgTwoRets[CurriedArg1, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1) (Ret1, Ret2)) func() (Ret1, Ret2) {
	return func() (Ret1, Ret2) {
		return f(c1)
	}
}

// FromOneArgVariadicTwoRets curries the first 1 arguments from a function with 1 plus variadic arguments and 2 return values.
func FromOneArgVariadicTwoRets[CurriedArg1, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) (Ret1, Ret2)) func(...Var) (Ret1, Ret2) {
	return func(v ...Var) (Ret1, Ret2) {
		return f(c1, v...)
	}
}

// FromOneArgThreeRets curries the first 1 arguments from a function with 1 arguments and 3 return values.
func FromOneArgThreeRets[CurriedArg1, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1) (Ret1, Ret2, Ret3)) func() (Ret1, Ret2, Ret3) {
	return func() (Ret1, Ret2, Ret3) {
		return f(c1)
	}
}

// FromOneArgVariadicThreeRets curries the first 1 arguments from a function with 1 plus variadic arguments and 3 return values.
func FromOneArgVariadicThreeRets[CurriedArg1, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) (Ret1, Ret2, Ret3)) func(...Var) (Ret1, Ret2, Ret3) {
	return func(v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, v...)
	}
}

// FromOneArgFourRets curries the first 1 arguments from a function with 1 arguments and 4 return values.
func FromOneArgFourRets[CurriedArg1, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1) (Ret1, Ret2, Ret3, Ret4)) func() (Ret1, Ret2, Ret3, Ret4) {
	return func() (Ret1, Ret2, Ret3, Ret4) {
		return f(c1)
	}
}

// FromOneArgVariadicFourRets curries the first 1 arguments from a function with 1 plus variadic arguments and 4 return values.
func FromOneArgVariadicFourRets[CurriedArg1, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, v...)
	}
}

// FromOneArgFiveRets curries the first 1 arguments from a function with 1 arguments and 5 return values.
func FromOneArgFiveRets[CurriedArg1, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1) (Ret1, Ret2, Ret3, Ret4, Ret5)) func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1)
	}
}

// FromOneArgVariadicFiveRets curries the first 1 arguments from a function with 1 plus variadic arguments and 5 return values.
func FromOneArgVariadicFiveRets[CurriedArg1, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, v...)
	}
}

// FromTwoArgs curries the first 1 arguments from a function with 2 arguments and 0 return values.
func FromTwoArgs[CurriedArg1, Arg1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1)) func(Arg1) {
	return func(a1 Arg1) {
		f(c1, a1)
	}
}

// FromTwoArgsVariadic curries the first 1 arguments from a function with 2 plus variadic arguments and 0 return values.
func FromTwoArgsVariadic[CurriedArg1, Arg1, Var any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var)) func(Arg1, ...Var) {
	return func(a1 Arg1, v ...Var) {
		f(c1, a1, v...)
	}
}

// FromTwoArgsOneRet curries the first 1 arguments from a function with 2 arguments and 1 return values.
func FromTwoArgsOneRet[CurriedArg1, Arg1, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) Ret1) func(Arg1) Ret1 {
	return func(a1 Arg1) Ret1 {
		return f(c1, a1)
	}
}

// FromTwoArgsVariadicOneRet curries the first 1 arguments from a function with 2 plus variadic arguments and 1 return values.
func FromTwoArgsVariadicOneRet[CurriedArg1, Arg1, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) Ret1) func(Arg1, ...Var) Ret1 {
	return func(a1 Arg1, v ...Var) Ret1 {
		return f(c1, a1, v...)
	}
}

// FromTwoArgsTwoRets curries the first 1 arguments from a function with 2 arguments and 2 return values.
func FromTwoArgsTwoRets[CurriedArg1, Arg1, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) (Ret1, Ret2)) func(Arg1) (Ret1, Ret2) {
	return func(a1 Arg1) (Ret1, Ret2) {
		return f(c1, a1)
	}
}

// FromTwoArgsVariadicTwoRets curries the first 1 arguments from a function with 2 plus variadic arguments and 2 return values.
func FromTwoArgsVariadicTwoRets[CurriedArg1, Arg1, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) (Ret1, Ret2)) func(Arg1, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2) {
		return f(c1, a1, v...)
	}
}

// FromTwoArgsThreeRets curries the first 1 arguments from a function with 2 arguments and 3 return values.
func FromTwoArgsThreeRets[CurriedArg1, Arg1, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) (Ret1, Ret2, Ret3)) func(Arg1) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3) {
		return f(c1, a1)
	}
}

// FromTwoArgsVariadicThreeRets curries the first 1 arguments from a function with 2 plus variadic arguments and 3 return values.
func FromTwoArgsVariadicThreeRets[CurriedArg1, Arg1, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, a1, v...)
	}
}

// FromTwoArgsFourRets curries the first 1 arguments from a function with 2 arguments and 4 return values.
func FromTwoArgsFourRets[CurriedArg1, Arg1, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) (Ret1, Ret2, Ret3, Ret4)) func(Arg1) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1)
	}
}

// FromTwoArgsVariadicFourRets curries the first 1 arguments from a function with 2 plus variadic arguments and 4 return values.
func FromTwoArgsVariadicFourRets[CurriedArg1, Arg1, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, v...)
	}
}

// FromTwoArgsFiveRets curries the first 1 arguments from a function with 2 arguments and 5 return values.
func FromTwoArgsFiveRets[CurriedArg1, Arg1, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1)
	}
}

// FromTwoArgsVariadicFiveRets curries the first 1 arguments from a function with 2 plus variadic arguments and 5 return values.
func FromTwoArgsVariadicFiveRets[CurriedArg1, Arg1, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, v...)
	}
}

// FromThreeArgs curries the first 1 arguments from a function with 3 arguments and 0 return values.
func FromThreeArgs[CurriedArg1, Arg1, Arg2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2)) func(Arg1, Arg2) {
	return func(a1 Arg1, a2 Arg2) {
		f(c1, a1, a2)
	}
}

// FromThreeArgsVariadic curries the first 1 arguments from a function with 3 plus variadic arguments and 0 return values.
func FromThreeArgsVariadic[CurriedArg1, Arg1, Arg2, Var any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var)) func(Arg1, Arg2, ...Var) {
	return func(a1 Arg1, a2 Arg2, v ...Var) {
		f(c1, a1, a2, v...)
	}
}

// FromThreeArgsOneRet curries the first 1 arguments from a function with 3 arguments and 1 return values.
func FromThreeArgsOneRet[CurriedArg1, Arg1, Arg2, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) Ret1) func(Arg1, Arg2) Ret1 {
	return func(a1 Arg1, a2 Arg2) Ret1 {
		return f(c1, a1, a2)
	}
}

// FromThreeArgsVariadicOneRet curries the first 1 arguments from a function with 3 plus variadic arguments and 1 return values.
func FromThreeArgsVariadicOneRet[CurriedArg1, Arg1, Arg2, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) Ret1) func(Arg1, Arg2, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, v ...Var) Ret1 {
		return f(c1, a1, a2, v...)
	}
}

// FromThreeArgsTwoRets curries the first 1 arguments from a function with 3 arguments and 2 return values.
func FromThreeArgsTwoRets[CurriedArg1, Arg1, Arg2, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) (Ret1, Ret2)) func(Arg1, Arg2) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2) {
		return f(c1, a1, a2)
	}
}

// FromThreeArgsVariadicTwoRets curries the first 1 arguments from a function with 3 plus variadic arguments and 2 return values.
func FromThreeArgsVariadicTwoRets[CurriedArg1, Arg1, Arg2, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2) {
		return f(c1, a1, a2, v...)
	}
}

// FromThreeArgsThreeRets curries the first 1 arguments from a function with 3 arguments and 3 return values.
func FromThreeArgsThreeRets[CurriedArg1, Arg1, Arg2, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) (Ret1, Ret2, Ret3)) func(Arg1, Arg2) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2)
	}
}

// FromThreeArgsVariadicThreeRets curries the first 1 arguments from a function with 3 plus variadic arguments and 3 return values.
func FromThreeArgsVariadicThreeRets[CurriedArg1, Arg1, Arg2, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, v...)
	}
}

// FromThreeArgsFourRets curries the first 1 arguments from a function with 3 arguments and 4 return values.
func FromThreeArgsFourRets[CurriedArg1, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2)
	}
}

// FromThreeArgsVariadicFourRets curries the first 1 arguments from a function with 3 plus variadic arguments and 4 return values.
func FromThreeArgsVariadicFourRets[CurriedArg1, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, v...)
	}
}

// FromThreeArgsFiveRets curries the first 1 arguments from a function with 3 arguments and 5 return values.
func FromThreeArgsFiveRets[CurriedArg1, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2)
	}
}

// FromThreeArgsVariadicFiveRets curries the first 1 arguments from a function with 3 plus variadic arguments and 5 return values.
func FromThreeArgsVariadicFiveRets[CurriedArg1, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, v...)
	}
}

// FromFourArgs curries the first 1 arguments from a function with 4 arguments and 0 return values.
func FromFourArgs[CurriedArg1, Arg1, Arg2, Arg3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3)) func(Arg1, Arg2, Arg3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) {
		f(c1, a1, a2, a3)
	}
}

// FromFourArgsVariadic curries the first 1 arguments from a function with 4 plus variadic arguments and 0 return values.
func FromFourArgsVariadic[CurriedArg1, Arg1, Arg2, Arg3, Var any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var)) func(Arg1, Arg2, Arg3, ...Var) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) {
		f(c1, a1, a2, a3, v...)
	}
}

// FromFourArgsOneRet curries the first 1 arguments from a function with 4 arguments and 1 return values.
func FromFourArgsOneRet[CurriedArg1, Arg1, Arg2, Arg3, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) Ret1) func(Arg1, Arg2, Arg3) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) Ret1 {
		return f(c1, a1, a2, a3)
	}
}

// FromFourArgsVariadicOneRet curries the first 1 arguments from a function with 4 plus variadic arguments and 1 return values.
func FromFourArgsVariadicOneRet[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) Ret1) func(Arg1, Arg2, Arg3, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) Ret1 {
		return f(c1, a1, a2, a3, v...)
	}
}

// FromFourArgsTwoRets curries the first 1 arguments from a function with 4 arguments and 2 return values.
func FromFourArgsTwoRets[CurriedArg1, Arg1, Arg2, Arg3, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) (Ret1, Ret2)) func(Arg1, Arg2, Arg3) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2) {
		return f(c1, a1, a2, a3)
	}
}

// FromFourArgsVariadicTwoRets curries the first 1 arguments from a function with 4 plus variadic arguments and 2 return values.
func FromFourArgsVariadicTwoRets[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2) {
		return f(c1, a1, a2, a3, v...)
	}
}

// FromFourArgsThreeRets curries the first 1 arguments from a function with 4 arguments and 3 return values.
func FromFourArgsThreeRets[CurriedArg1, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, a3)
	}
}

// FromFourArgsVariadicThreeRets curries the first 1 arguments from a function with 4 plus variadic arguments and 3 return values.
func FromFourArgsVariadicThreeRets[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, a3, v...)
	}
}

// FromFourArgsFourRets curries the first 1 arguments from a function with 4 arguments and 4 return values.
func FromFourArgsFourRets[CurriedArg1, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, a3)
	}
}

// FromFourArgsVariadicFourRets curries the first 1 arguments from a function with 4 plus variadic arguments and 4 return values.
func FromFourArgsVariadicFourRets[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, a3, v...)
	}
}

// FromFourArgsFiveRets curries the first 1 arguments from a function with 4 arguments and 5 return values.
func FromFourArgsFiveRets[CurriedArg1, Arg1, Arg2, Arg3, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, a3)
	}
}

// FromFourArgsVariadicFiveRets curries the first 1 arguments from a function with 4 plus variadic arguments and 5 return values.
func FromFourArgsVariadicFiveRets[CurriedArg1, Arg1, Arg2, Arg3, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, a3, v...)
	}
}

// FromFiveArgs curries the first 1 arguments from a function with 5 arguments and 0 return values.
func FromFiveArgs[CurriedArg1, Arg1, Arg2, Arg3, Arg4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4)) func(Arg1, Arg2, Arg3, Arg4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) {
		f(c1, a1, a2, a3, a4)
	}
}

// FromFiveArgsVariadic curries the first 1 arguments from a function with 5 plus variadic arguments and 0 return values.
func FromFiveArgsVariadic[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var)) func(Arg1, Arg2, Arg3, Arg4, ...Var) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) {
		f(c1, a1, a2, a3, a4, v...)
	}
}

// FromFiveArgsOneRet curries the first 1 arguments from a function with 5 arguments and 1 return values.
func FromFiveArgsOneRet[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) Ret1) func(Arg1, Arg2, Arg3, Arg4) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) Ret1 {
		return f(c1, a1, a2, a3, a4)
	}
}

// FromFiveArgsVariadicOneRet curries the first 1 arguments from a function with 5 plus variadic arguments and 1 return values.
func FromFiveArgsVariadicOneRet[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) Ret1) func(Arg1, Arg2, Arg3, Arg4, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) Ret1 {
		return f(c1, a1, a2, a3, a4, v...)
	}
}

// FromFiveArgsTwoRets curries the first 1 arguments from a function with 5 arguments and 2 return values.
func FromFiveArgsTwoRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2)) func(Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) (Ret1, Ret2) {
		return f(c1, a1, a2, a3, a4)
	}
}

// FromFiveArgsVariadicTwoRets curries the first 1 arguments from a function with 5 plus variadic arguments and 2 return values.
func FromFiveArgsVariadicTwoRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1, Ret2 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) (Ret1, Ret2) {
		return f(c1, a1, a2, a3, a4, v...)
	}
}

// FromFiveArgsThreeRets curries the first 1 arguments from a function with 5 arguments and 3 return values.
func FromFiveArgsThreeRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, a3, a4)
	}
}

// FromFiveArgsVariadicThreeRets curries the first 1 arguments from a function with 5 plus variadic arguments and 3 return values.
func FromFiveArgsVariadicThreeRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, a1, a2, a3, a4, v...)
	}
}

// FromFiveArgsFourRets curries the first 1 arguments from a function with 5 arguments and 4 return values.
func FromFiveArgsFourRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, a3, a4)
	}
}

// FromFiveArgsVariadicFourRets curries the first 1 arguments from a function with 5 plus variadic arguments and 4 return values.
func FromFiveArgsVariadicFourRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, a1, a2, a3, a4, v...)
	}
}

// FromFiveArgsFiveRets curries the first 1 arguments from a function with 5 arguments and 5 return values.
func FromFiveArgsFiveRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3, Arg4) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, a3, a4)
	}
}

// FromFiveArgsVariadicFiveRets curries the first 1 arguments from a function with 5 plus variadic arguments and 5 return values.
func FromFiveArgsVariadicFiveRets[CurriedArg1, Arg1, Arg2, Arg3, Arg4, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, f func(CurriedArg1, Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, Arg3, Arg4, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, a3 Arg3, a4 Arg4, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, a1, a2, a3, a4, v...)
	}
}
