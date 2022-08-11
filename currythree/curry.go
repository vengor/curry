package currythree

// FromThreeArgs curries the first 3 arguments from a function with 3 arguments and 0 return values.
func FromThreeArgs[CurriedArg1, CurriedArg2, CurriedArg3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3)) func() {
	return func() {
		f(c1, c2, c3)
	}
}

// FromThreeArgsVariadic curries the first 3 arguments from a function with 3 plus variadic arguments and 0 return values.
func FromThreeArgsVariadic[CurriedArg1, CurriedArg2, CurriedArg3, Var any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var)) func(...Var) {
	return func(v ...Var) {
		f(c1, c2, c3, v...)
	}
}

// FromThreeArgsOneRet curries the first 3 arguments from a function with 3 arguments and 1 return values.
func FromThreeArgsOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) Ret1) func() Ret1 {
	return func() Ret1 {
		return f(c1, c2, c3)
	}
}

// FromThreeArgsVariadicOneRet curries the first 3 arguments from a function with 3 plus variadic arguments and 1 return values.
func FromThreeArgsVariadicOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) Ret1) func(...Var) Ret1 {
	return func(v ...Var) Ret1 {
		return f(c1, c2, c3, v...)
	}
}

// FromThreeArgsTwoRets curries the first 3 arguments from a function with 3 arguments and 2 return values.
func FromThreeArgsTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) (Ret1, Ret2)) func() (Ret1, Ret2) {
	return func() (Ret1, Ret2) {
		return f(c1, c2, c3)
	}
}

// FromThreeArgsVariadicTwoRets curries the first 3 arguments from a function with 3 plus variadic arguments and 2 return values.
func FromThreeArgsVariadicTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) (Ret1, Ret2)) func(...Var) (Ret1, Ret2) {
	return func(v ...Var) (Ret1, Ret2) {
		return f(c1, c2, c3, v...)
	}
}

// FromThreeArgsThreeRets curries the first 3 arguments from a function with 3 arguments and 3 return values.
func FromThreeArgsThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) (Ret1, Ret2, Ret3)) func() (Ret1, Ret2, Ret3) {
	return func() (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3)
	}
}

// FromThreeArgsVariadicThreeRets curries the first 3 arguments from a function with 3 plus variadic arguments and 3 return values.
func FromThreeArgsVariadicThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) (Ret1, Ret2, Ret3)) func(...Var) (Ret1, Ret2, Ret3) {
	return func(v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, v...)
	}
}

// FromThreeArgsFourRets curries the first 3 arguments from a function with 3 arguments and 4 return values.
func FromThreeArgsFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) (Ret1, Ret2, Ret3, Ret4)) func() (Ret1, Ret2, Ret3, Ret4) {
	return func() (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3)
	}
}

// FromThreeArgsVariadicFourRets curries the first 3 arguments from a function with 3 plus variadic arguments and 4 return values.
func FromThreeArgsVariadicFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, v...)
	}
}

// FromThreeArgsFiveRets curries the first 3 arguments from a function with 3 arguments and 5 return values.
func FromThreeArgsFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3) (Ret1, Ret2, Ret3, Ret4, Ret5)) func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func() (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3)
	}
}

// FromThreeArgsVariadicFiveRets curries the first 3 arguments from a function with 3 plus variadic arguments and 5 return values.
func FromThreeArgsVariadicFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, v...)
	}
}

// FromFourArgs curries the first 3 arguments from a function with 4 arguments and 0 return values.
func FromFourArgs[CurriedArg1, CurriedArg2, CurriedArg3, Arg1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1)) func(Arg1) {
	return func(a1 Arg1) {
		f(c1, c2, c3, a1)
	}
}

// FromFourArgsVariadic curries the first 3 arguments from a function with 4 plus variadic arguments and 0 return values.
func FromFourArgsVariadic[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var)) func(Arg1, ...Var) {
	return func(a1 Arg1, v ...Var) {
		f(c1, c2, c3, a1, v...)
	}
}

// FromFourArgsOneRet curries the first 3 arguments from a function with 4 arguments and 1 return values.
func FromFourArgsOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) Ret1) func(Arg1) Ret1 {
	return func(a1 Arg1) Ret1 {
		return f(c1, c2, c3, a1)
	}
}

// FromFourArgsVariadicOneRet curries the first 3 arguments from a function with 4 plus variadic arguments and 1 return values.
func FromFourArgsVariadicOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) Ret1) func(Arg1, ...Var) Ret1 {
	return func(a1 Arg1, v ...Var) Ret1 {
		return f(c1, c2, c3, a1, v...)
	}
}

// FromFourArgsTwoRets curries the first 3 arguments from a function with 4 arguments and 2 return values.
func FromFourArgsTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) (Ret1, Ret2)) func(Arg1) (Ret1, Ret2) {
	return func(a1 Arg1) (Ret1, Ret2) {
		return f(c1, c2, c3, a1)
	}
}

// FromFourArgsVariadicTwoRets curries the first 3 arguments from a function with 4 plus variadic arguments and 2 return values.
func FromFourArgsVariadicTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) (Ret1, Ret2)) func(Arg1, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2) {
		return f(c1, c2, c3, a1, v...)
	}
}

// FromFourArgsThreeRets curries the first 3 arguments from a function with 4 arguments and 3 return values.
func FromFourArgsThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) (Ret1, Ret2, Ret3)) func(Arg1) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, a1)
	}
}

// FromFourArgsVariadicThreeRets curries the first 3 arguments from a function with 4 plus variadic arguments and 3 return values.
func FromFourArgsVariadicThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, a1, v...)
	}
}

// FromFourArgsFourRets curries the first 3 arguments from a function with 4 arguments and 4 return values.
func FromFourArgsFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) (Ret1, Ret2, Ret3, Ret4)) func(Arg1) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, a1)
	}
}

// FromFourArgsVariadicFourRets curries the first 3 arguments from a function with 4 plus variadic arguments and 4 return values.
func FromFourArgsVariadicFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, a1, v...)
	}
}

// FromFourArgsFiveRets curries the first 3 arguments from a function with 4 arguments and 5 return values.
func FromFourArgsFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, a1)
	}
}

// FromFourArgsVariadicFiveRets curries the first 3 arguments from a function with 4 plus variadic arguments and 5 return values.
func FromFourArgsVariadicFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, a1, v...)
	}
}

// FromFiveArgs curries the first 3 arguments from a function with 5 arguments and 0 return values.
func FromFiveArgs[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2)) func(Arg1, Arg2) {
	return func(a1 Arg1, a2 Arg2) {
		f(c1, c2, c3, a1, a2)
	}
}

// FromFiveArgsVariadic curries the first 3 arguments from a function with 5 plus variadic arguments and 0 return values.
func FromFiveArgsVariadic[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var)) func(Arg1, Arg2, ...Var) {
	return func(a1 Arg1, a2 Arg2, v ...Var) {
		f(c1, c2, c3, a1, a2, v...)
	}
}

// FromFiveArgsOneRet curries the first 3 arguments from a function with 5 arguments and 1 return values.
func FromFiveArgsOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) Ret1) func(Arg1, Arg2) Ret1 {
	return func(a1 Arg1, a2 Arg2) Ret1 {
		return f(c1, c2, c3, a1, a2)
	}
}

// FromFiveArgsVariadicOneRet curries the first 3 arguments from a function with 5 plus variadic arguments and 1 return values.
func FromFiveArgsVariadicOneRet[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) Ret1) func(Arg1, Arg2, ...Var) Ret1 {
	return func(a1 Arg1, a2 Arg2, v ...Var) Ret1 {
		return f(c1, c2, c3, a1, a2, v...)
	}
}

// FromFiveArgsTwoRets curries the first 3 arguments from a function with 5 arguments and 2 return values.
func FromFiveArgsTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) (Ret1, Ret2)) func(Arg1, Arg2) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2) {
		return f(c1, c2, c3, a1, a2)
	}
}

// FromFiveArgsVariadicTwoRets curries the first 3 arguments from a function with 5 plus variadic arguments and 2 return values.
func FromFiveArgsVariadicTwoRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1, Ret2 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) (Ret1, Ret2)) func(Arg1, Arg2, ...Var) (Ret1, Ret2) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2) {
		return f(c1, c2, c3, a1, a2, v...)
	}
}

// FromFiveArgsThreeRets curries the first 3 arguments from a function with 5 arguments and 3 return values.
func FromFiveArgsThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) (Ret1, Ret2, Ret3)) func(Arg1, Arg2) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, a1, a2)
	}
}

// FromFiveArgsVariadicThreeRets curries the first 3 arguments from a function with 5 plus variadic arguments and 3 return values.
func FromFiveArgsVariadicThreeRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1, Ret2, Ret3 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3) {
		return f(c1, c2, c3, a1, a2, v...)
	}
}

// FromFiveArgsFourRets curries the first 3 arguments from a function with 5 arguments and 4 return values.
func FromFiveArgsFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, a1, a2)
	}
}

// FromFiveArgsVariadicFourRets curries the first 3 arguments from a function with 5 plus variadic arguments and 4 return values.
func FromFiveArgsVariadicFourRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4) {
		return f(c1, c2, c3, a1, a2, v...)
	}
}

// FromFiveArgsFiveRets curries the first 3 arguments from a function with 5 arguments and 5 return values.
func FromFiveArgsFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, a1, a2)
	}
}

// FromFiveArgsVariadicFiveRets curries the first 3 arguments from a function with 5 plus variadic arguments and 5 return values.
func FromFiveArgsVariadicFiveRets[CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, Var, Ret1, Ret2, Ret3, Ret4, Ret5 any](c1 CurriedArg1, c2 CurriedArg2, c3 CurriedArg3, f func(CurriedArg1, CurriedArg2, CurriedArg3, Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5)) func(Arg1, Arg2, ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
	return func(a1 Arg1, a2 Arg2, v ...Var) (Ret1, Ret2, Ret3, Ret4, Ret5) {
		return f(c1, c2, c3, a1, a2, v...)
	}
}
