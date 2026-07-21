## 运行结果
- ### go run ./stage0/d02_map_struct/unsafe_map
fatal error: concurrent map writes
goroutine 27 [running]:                                                    
main.main.func1(0x9)                                                       
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:17 +0x72
created by main.main in goroutine 1                                        
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 1 [semacquire]:                                                  
sync.runtime_Semacquire(0xc0000901c8?)                                     
D:/env/GO_SDK/go1.23.5/src/runtime/sema.go:71 +0x25
sync.(*WaitGroup).Wait(0x68?)
D:/env/GO_SDK/go1.23.5/src/sync/waitgroup.go:118 +0x48
main.main()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:22 +0xfc

goroutine 18 [runnable]:
main.main.func1(0x0)
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:17 +0x72
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 19 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 20 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 21 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 22 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 23 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 24 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 25 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45

goroutine 26 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x45
exit status 2
- ### go run -race ./stage0/d02_map_struct/unsafe_map
- ==================
  WARNING: DATA RACE
  Write at 0x00c000024030 by goroutine 6:
  runtime.mapassign_fast64()
  D:/env/GO_SDK/go1.23.5/src/runtime/map_fast64.go:113 +0x0
  main.main.func1()
  E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:17 +0xc4
  main.main.gowrap1()
  E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:19 +0x41

Previous write at 0x00c000024030 by goroutine 13:
runtime.mapassign_fast64()
D:/env/GO_SDK/go1.23.5/src/runtime/map_fast64.go:113 +0x0
main.main.func1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:17 +0xc4
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:19 +0x41

Goroutine 6 (running) created at:
main.main()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7e

Goroutine 13 (finished) created at:
main.main()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7e
==================
fatal error: concurrent map writes
fatal error: concurrent map writes
fatal error: concurrent map writes

goroutine 8 [running]:
main.main.func1(0x3)
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:17 +0xc5
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc00000a058?)
D:/env/GO_SDK/go1.23.5/src/runtime/sema.go:71 +0x25
sync.(*WaitGroup).Wait(0xc00000a050)
D:/env/GO_SDK/go1.23.5/src/sync/waitgroup.go:118 +0xa8
main.main()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:22 +0x1b5

goroutine 6 [runnable]:
main.main.func1(0x1)
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:17 +0xc5
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f

goroutine 7 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f

goroutine 9 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f

goroutine 10 [running]:
goroutine running on other thread; stack unavailable
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f

goroutine 13 [running]:
goroutine running on other thread; stack unavailable
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f

goroutine 14 [runnable]:
main.main.gowrap1()
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13
runtime.goexit({})
D:/env/GO_SDK/go1.23.5/src/runtime/asm_amd64.s:1700 +0x1
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f

goroutine 10 [running]:
main.main.func1(0x5)
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:17 +0xc5
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f

goroutine 13 [running]:
main.main.func1(0x8)
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:17 +0xc5
created by main.main in goroutine 1
E:/Coding/Go/UPUP/stage0/d02_map_struct/unsafe_map/main.go:13 +0x7f
exit status 2