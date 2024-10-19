package base

// Pipeline operator as an function
// IN -> Pipeline -> OUT
// IN -> (F1 -> F2 -> F3 -> ...) -> OUT
func Pipeline1[Input any, Output any](f1 func(Input) Output) func(Input) Output {
    return func(input Input) Output {
        return f1(input)
    }
}

func Pipeline2[Input any, I1 any, Output any](f1 func(Input) I1, f2 func(I1) Output) func(Input) Output {
    return func(input Input) Output {
        return f2(f1(input))
    }
}

func Pipeline3[Input any, I1 any, I2 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) Output) func(Input) Output {
    return func(input Input) Output {
        return f3(f2(f1(input)))
    }
}

func Pipeline4[Input any, I1 any, I2 any, I3 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) Output) func(Input) Output {
    return func(input Input) Output {
        return f4(f3(f2(f1(input))))
    }
}

func Pipeline5[Input any, I1 any, I2 any, I3 any, I4 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) Output) func(Input) Output {
    return func(input Input) Output {
        return f5(f4(f3(f2(f1(input)))))
    }
}

func Pipeline6[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) Output) func(Input) Output {
    return func(input Input) Output {
        return f6(f5(f4(f3(f2(f1(input))))))
    }
}

func Pipeline7[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) Output) func(Input) Output {
    return func(input Input) Output {
        return f7(f6(f5(f4(f3(f2(f1(input)))))))
    }
}

func Pipeline8[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) Output) func(Input) Output {
    return func(input Input) Output {
        return f8(f7(f6(f5(f4(f3(f2(f1(input))))))))
    }
}

func Pipeline9[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) Output) func(Input) Output {
    return func(input Input) Output {
        return f9(f8(f7(f6(f5(f4(f3(f2(f1(input)))))))))
    }
}

func Pipeline10[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) Output) func(Input) Output {
    return func(input Input) Output {
        return f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input))))))))))
    }
}

func Pipeline11[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) Output) func(Input) Output {
    return func(input Input) Output {
        return f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input)))))))))))
    }
}

func Pipeline12[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) Output) func(Input) Output {
    return func(input Input) Output {
        return f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input))))))))))))
    }
}

func Pipeline13[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) I12, f13 func(I12) Output) func(Input) Output {
    return func(input Input) Output {
        return f13(f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input)))))))))))))
    }
}

func Pipeline14[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) I12, f13 func(I12) I13, f14 func(I13) Output) func(Input) Output {
    return func(input Input) Output {
        return f14(f13(f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input))))))))))))))
    }
}

func Pipeline15[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) I12, f13 func(I12) I13, f14 func(I13) I14, f15 func(I14) Output) func(Input) Output {
    return func(input Input) Output {
        return f15(f14(f13(f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input)))))))))))))))
    }
}

func Pipeline16[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) I12, f13 func(I12) I13, f14 func(I13) I14, f15 func(I14) I15, f16 func(I15) Output) func(Input) Output {
    return func(input Input) Output {
        return f16(f15(f14(f13(f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input))))))))))))))))
    }
}

func Pipeline17[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) I12, f13 func(I12) I13, f14 func(I13) I14, f15 func(I14) I15, f16 func(I15) I16, f17 func(I16) Output) func(Input) Output {
    return func(input Input) Output {
        return f17(f16(f15(f14(f13(f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input)))))))))))))))))
    }
}

func Pipeline18[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) I12, f13 func(I12) I13, f14 func(I13) I14, f15 func(I14) I15, f16 func(I15) I16, f17 func(I16) I17, f18 func(I17) Output) func(Input) Output {
    return func(input Input) Output {
        return f18(f17(f16(f15(f14(f13(f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input))))))))))))))))))
    }
}

func Pipeline19[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) I12, f13 func(I12) I13, f14 func(I13) I14, f15 func(I14) I15, f16 func(I15) I16, f17 func(I16) I17, f18 func(I17) I18, f19 func(I18) Output) func(Input) Output {
    return func(input Input) Output {
        return f19(f18(f17(f16(f15(f14(f13(f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input)))))))))))))))))))
    }
}

func Pipeline20[Input any, I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, Output any](f1 func(Input) I1, f2 func(I1) I2, f3 func(I2) I3, f4 func(I3) I4, f5 func(I4) I5, f6 func(I5) I6, f7 func(I6) I7, f8 func(I7) I8, f9 func(I8) I9, f10 func(I9) I10, f11 func(I10) I11, f12 func(I11) I12, f13 func(I12) I13, f14 func(I13) I14, f15 func(I14) I15, f16 func(I15) I16, f17 func(I16) I17, f18 func(I17) I18, f19 func(I18) I19, f20 func(I19) Output) func(Input) Output {
    return func(input Input) Output {
        return f20(f19(f18(f17(f16(f15(f14(f13(f12(f11(f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(input))))))))))))))))))))
    }
}
