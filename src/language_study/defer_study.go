package language_study

func Defer_study() {

	/*
	在runtime2.go下定义了defer
	type _defer struct {
				siz     int32 // includes both arguments and results
				started bool
				heap    bool
				sp      uintptr // 调用deferproc时的sp
				pc      uintptr // 调用deferproc时的ip
				fn      *funcval
				_panic  *_panic // panic that is running defer
				link    *_defer
	}

	panic.go下
	func deferproc(siz int32, fn *funcval) { // arguments of fn follow fn
		if getg().m.curg != getg() {
			// go code on the system stack can't defer
			throw("defer on system stack")
		}

		sp := getcallersp() //获取到defer的SP
		argp := uintptr(unsafe.Pointer(&fn)) + unsafe.Sizeof(fn)
		callerpc := getcallerpc() //获取调用defer函数的程序计数器

		d := newdefer(siz)
		if d._panic != nil {
			throw("deferproc: d.panic != nil after newdefer")
		}
		d.fn = fn
		d.pc = callerpc
		d.sp = sp
		switch siz {
		case 0:
			// Do nothing.
		case sys.PtrSize:
			*(*uintptr)(deferArgs(d)) = *(*uintptr)(unsafe.Pointer(argp))
		default:
			memmove(deferArgs(d), unsafe.Pointer(argp), uintptr(siz))
		}
		return0()
	}

	*/
	defer println(0x11)

	/*
		go tool objdump -s "main\.main" main

		在第33行，37行可以看到编译器将defer处理成了两个函数调用，deferproc定义一个延迟调用对象，然后在函数结束前通过deferreturn完成最终调用

		TEXT main.main(SB) /Users/muyou/go/src/awesomeProject/src/main/main.go
		  main.go:3             0x10512e0               65488b0c2530000000      MOVQ GS:0x30, CX
		  main.go:3             0x10512e9               483b6110                CMPQ 0x10(CX), SP
		  main.go:3             0x10512ed               765d                    JBE 0x105134c
		  main.go:3             0x10512ef               4883ec48                SUBQ $0x48, SP
		  main.go:3             0x10512f3               48896c2440              MOVQ BP, 0x40(SP)
		  main.go:3             0x10512f8               488d6c2440              LEAQ 0x40(SP), BP
		  main.go:14            0x10512fd               c744240808000000        MOVL $0x8, 0x8(SP)
		  main.go:14            0x1051305               488d05a4410200          LEAQ go.func.*+66(SB), AX
		  main.go:14            0x105130c               4889442420              MOVQ AX, 0x20(SP)
		  main.go:14            0x1051311               48c744243811000000      MOVQ $0x11, 0x38(SP)
		  main.go:14            0x105131a               488d442408              LEAQ 0x8(SP), AX
		  main.go:14            0x105131f               48890424                MOVQ AX, 0(SP)
		  main.go:14            0x1051323               e84820fdff              CALL runtime.deferprocStack(SB)
		  main.go:14            0x1051328               85c0                    TESTL AX, AX
		  main.go:14            0x105132a               7510                    JNE 0x105133c
		  main.go:15            0x105132c               90                      NOPL
		  main.go:15            0x105132d               e83e26fdff              CALL runtime.deferreturn(SB)
		  main.go:15            0x1051332               488b6c2440              MOVQ 0x40(SP), BP
		  main.go:15            0x1051337               4883c448                ADDQ $0x48, SP
		  main.go:15            0x105133b               c3                      RET
		  main.go:14            0x105133c               90                      NOPL
		  main.go:14            0x105133d               e82e26fdff              CALL runtime.deferreturn(SB)
		  main.go:14            0x1051342               488b6c2440              MOVQ 0x40(SP), BP
		  main.go:14            0x1051347               4883c448                ADDQ $0x48, SP
		  main.go:14            0x105134b               c3                      RET
		  main.go:3             0x105134c               e83f81ffff              CALL runtime.morestack_noctxt(SB)
		  main.go:3             0x1051351               eb8d                    JMP main.main(SB)
		  :-1                   0x1051353               cc                      INT $0x3
		  :-1                   0x1051354               cc                      INT $0x3
		  :-1                   0x1051355               cc                      INT $0x3
		  :-1                   0x1051356               cc                      INT $0x3
		  :-1                   0x1051357               cc                      INT $0x3
		  :-1                   0x1051358               cc                      INT $0x3
		  :-1                   0x1051359               cc                      INT $0x3
		  :-1                   0x105135a               cc                      INT $0x3
		  :-1                   0x105135b               cc                      INT $0x3
		  :-1                   0x105135c               cc                      INT $0x3
		  :-1                   0x105135d               cc                      INT $0x3
		  :-1                   0x105135e               cc                      INT $0x3
		  :-1                   0x105135f               cc                      INT $0x3
	*/
}
