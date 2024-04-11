## writer problem

how to trigger it:

```
# alexandruvladutu @ 192-168-0-180 in ~/www/go-web [17:01:04]
$ curl -sS http://localhost:8080/v1/item/10
received request for item 10%

# alexandruvladutu @ 192-168-0-180 in ~/www/go-web [10:35:53]
$ curl -sS http://localhost:8080/item/10
curl: (52) Empty reply from server
```

full log:

```
stack length: 2
2024/04/11 10:35:49 SampleMiddleware
2024/04/11 10:35:49 Logging1
2024/04/11 10:35:49 Server started on port :8080
2024/04/11 10:35:53 SampleMiddleware tick
2024/04/11 10:35:53 200 GET /v1/item/10 22.667µs
2024/04/11 10:35:56 SampleMiddleware tick
runtime: goroutine stack exceeds 1000000000-byte limit
runtime: sp=0x14020160390 stack=[0x14020160000, 0x14040160000]
fatal error: stack overflow

runtime stack:
runtime.throw({0x1010f2717?, 0x200000001?})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/panic.go:1023 +0x40 fp=0x17070ed50 sp=0x17070ed20 pc=0x100f3b080
runtime.newstack()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/stack.go:1103 +0x460 fp=0x17070ef00 sp=0x17070ed50 pc=0x100f57780
runtime.morestack()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/asm_arm64.s:341 +0x70 fp=0x17070ef00 sp=0x17070ef00 pc=0x100f6eef0

goroutine 21 gp=0x14000003880 m=4 mp=0x14000063b08 [running]:
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:14 +0x4c fp=0x14020160390 sp=0x14020160390 pc=0x1010ee80c
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201603b0 sp=0x14020160390 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201603d0 sp=0x140201603b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201603f0 sp=0x140201603d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160410 sp=0x140201603f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160430 sp=0x14020160410 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160450 sp=0x14020160430 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160470 sp=0x14020160450 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160490 sp=0x14020160470 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201604b0 sp=0x14020160490 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201604d0 sp=0x140201604b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201604f0 sp=0x140201604d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160510 sp=0x140201604f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160530 sp=0x14020160510 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160550 sp=0x14020160530 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160570 sp=0x14020160550 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160590 sp=0x14020160570 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201605b0 sp=0x14020160590 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201605d0 sp=0x140201605b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201605f0 sp=0x140201605d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160610 sp=0x140201605f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160630 sp=0x14020160610 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160650 sp=0x14020160630 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160670 sp=0x14020160650 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160690 sp=0x14020160670 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201606b0 sp=0x14020160690 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201606d0 sp=0x140201606b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201606f0 sp=0x140201606d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160710 sp=0x140201606f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160730 sp=0x14020160710 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160750 sp=0x14020160730 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160770 sp=0x14020160750 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160790 sp=0x14020160770 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201607b0 sp=0x14020160790 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201607d0 sp=0x140201607b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201607f0 sp=0x140201607d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160810 sp=0x140201607f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160830 sp=0x14020160810 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160850 sp=0x14020160830 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160870 sp=0x14020160850 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160890 sp=0x14020160870 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201608b0 sp=0x14020160890 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201608d0 sp=0x140201608b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201608f0 sp=0x140201608d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160910 sp=0x140201608f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160930 sp=0x14020160910 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160950 sp=0x14020160930 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160970 sp=0x14020160950 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x14020160990 sp=0x14020160970 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x140201609b0 sp=0x14020160990 pc=0x1010ee7e4
...16777040 frames elided...
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f3d0 sp=0x1404015f3b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f3f0 sp=0x1404015f3d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f410 sp=0x1404015f3f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f430 sp=0x1404015f410 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f450 sp=0x1404015f430 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f470 sp=0x1404015f450 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f490 sp=0x1404015f470 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f4b0 sp=0x1404015f490 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f4d0 sp=0x1404015f4b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f4f0 sp=0x1404015f4d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f510 sp=0x1404015f4f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f530 sp=0x1404015f510 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f550 sp=0x1404015f530 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f570 sp=0x1404015f550 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f590 sp=0x1404015f570 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f5b0 sp=0x1404015f590 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f5d0 sp=0x1404015f5b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f5f0 sp=0x1404015f5d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f610 sp=0x1404015f5f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f630 sp=0x1404015f610 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f650 sp=0x1404015f630 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f670 sp=0x1404015f650 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f690 sp=0x1404015f670 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f6b0 sp=0x1404015f690 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f6d0 sp=0x1404015f6b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f6f0 sp=0x1404015f6d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f710 sp=0x1404015f6f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f730 sp=0x1404015f710 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f750 sp=0x1404015f730 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f770 sp=0x1404015f750 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f790 sp=0x1404015f770 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f7b0 sp=0x1404015f790 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f7d0 sp=0x1404015f7b0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f7f0 sp=0x1404015f7d0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f810 sp=0x1404015f7f0 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f830 sp=0x1404015f810 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f850 sp=0x1404015f830 pc=0x1010ee7e4
git.v6.io/go-web/middleware.(*wrappedWriter).WriteHeader(0x140000c4000, 0x194)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:15 +0x24 fp=0x1404015f870 sp=0x1404015f850 pc=0x1010ee7e4
net/http.Error({0x1011d7200, 0x140000c4000}, {0x1010f38e7, 0x12}, 0x194)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:2178 +0x184 fp=0x1404015f8e0 sp=0x1404015f870 pc=0x1010e0c44
net/http.NotFound({0x1011d7200?, 0x140000c4000?}, 0x101684e68?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:2183 +0x34 fp=0x1404015f920 sp=0x1404015f8e0 pc=0x1010e0d54
net/http.HandlerFunc.ServeHTTP(0x14000109040?, {0x1011d7200?, 0x140000c4000?}, 0x1400007e150?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:2166 +0x38 fp=0x1404015f950 sp=0x1404015f920 pc=0x1010e0a88
net/http.(*ServeMux).ServeHTTP(0x1400007e150?, {0x1011d7200, 0x140000c4000}, 0x140000b8000)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:2683 +0x1a4 fp=0x1404015f9a0 sp=0x1404015f950 pc=0x1010e25f4
git.v6.io/go-web/middleware.SampleMiddleware.func1({0x1011d7200, 0x140000c4000}, 0x140000b8000)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:23 +0xb0 fp=0x1404015fa10 sp=0x1404015f9a0 pc=0x1010ee9d0
net/http.HandlerFunc.ServeHTTP(0x72?, {0x1011d7200?, 0x140000c4000?}, 0x100f11e3c?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:2166 +0x38 fp=0x1404015fa40 sp=0x1404015fa10 pc=0x1010e0a88
git.v6.io/go-web/middleware.Logging.func1({0x1011d7170, 0x140000c0000}, 0x140000b8000)
	/Users/alexandruvladutu/www/go-web/middleware/logging.go:41 +0xac fp=0x1404015fb10 sp=0x1404015fa40 pc=0x1010eec6c
net/http.HandlerFunc.ServeHTTP(0x0?, {0x1011d7170?, 0x140000c0000?}, 0x1400006eb50?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:2166 +0x38 fp=0x1404015fb40 sp=0x1404015fb10 pc=0x1010e0a88
net/http.serverHandler.ServeHTTP({0x140000b0090?}, {0x1011d7170?, 0x140000c0000?}, 0x6?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:3137 +0xbc fp=0x1404015fb70 sp=0x1404015fb40 pc=0x1010e369c
net/http.(*conn).serve(0x140000b4000, {0x1011d7618, 0x1400007ec60})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:2039 +0x508 fp=0x1404015ffa0 sp=0x1404015fb70 pc=0x1010df938
net/http.(*Server).Serve.gowrap3()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:3285 +0x30 fp=0x1404015ffd0 sp=0x1404015ffa0 pc=0x1010e3de0
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1404015ffd0 sp=0x1404015ffd0 pc=0x100f712c4
created by net/http.(*Server).Serve in goroutine 1
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:3285 +0x3f0

goroutine 1 gp=0x140000021c0 m=nil [IO wait]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x100f1af6c?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000143aa0 sp=0x14000143a80 pc=0x100f3dde8
runtime.netpollblock(0x14000143b38?, 0xfb1c94?, 0x1?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/netpoll.go:573 +0x158 fp=0x14000143ae0 sp=0x14000143aa0 pc=0x100f378d8
internal/poll.runtime_pollWait(0x1016c8ea0, 0x72)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/netpoll.go:345 +0xa0 fp=0x14000143b10 sp=0x14000143ae0 pc=0x100f6b720
internal/poll.(*pollDesc).wait(0x1400002c080?, 0x100f1a25c?, 0x0)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/internal/poll/fd_poll_runtime.go:84 +0x28 fp=0x14000143b40 sp=0x14000143b10 pc=0x100fb0a28
internal/poll.(*pollDesc).waitRead(...)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).Accept(0x1400002c080)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/internal/poll/fd_unix.go:611 +0x250 fp=0x14000143bf0 sp=0x14000143b40 pc=0x100fb1d80
net.(*netFD).accept(0x1400002c080)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/fd_unix.go:172 +0x28 fp=0x14000143cb0 sp=0x14000143bf0 pc=0x1010395d8
net.(*TCPListener).accept(0x14000126100)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/tcpsock_posix.go:159 +0x28 fp=0x14000143ce0 sp=0x14000143cb0 pc=0x101048cd8
net.(*TCPListener).Accept(0x14000126100)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/tcpsock.go:327 +0x2c fp=0x14000143d20 sp=0x14000143ce0 pc=0x10104809c
net/http.(*onceCloseListener).Accept(0x140000b4000?)
	<autogenerated>:1 +0x30 fp=0x14000143d40 sp=0x14000143d20 pc=0x1010ed9e0
net/http.(*Server).Serve(0x1400013c000, {0x1011d7260, 0x14000126100})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:3255 +0x2a8 fp=0x14000143e70 sp=0x14000143d40 pc=0x1010e3a48
net/http.(*Server).ListenAndServe(0x1400013c000)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:3184 +0x84 fp=0x14000143ea0 sp=0x14000143e70 pc=0x1010e3764
main.main()
	/Users/alexandruvladutu/www/go-web/main.go:31 +0x254 fp=0x14000143f40 sp=0x14000143ea0 pc=0x1010ef794
runtime.main()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:271 +0x28c fp=0x14000143fd0 sp=0x14000143f40 pc=0x100f3d9bc
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000143fd0 sp=0x14000143fd0 pc=0x100f712c4

goroutine 18 gp=0x14000094380 m=nil [force gc (idle)]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000058790 sp=0x14000058770 pc=0x100f3dde8
runtime.goparkunlock(...)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:408
runtime.forcegchelper()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:326 +0xb8 fp=0x140000587d0 sp=0x14000058790 pc=0x100f3dc78
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000587d0 sp=0x140000587d0 pc=0x100f712c4
created by runtime.init.6 in goroutine 1
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:314 +0x24

goroutine 19 gp=0x14000094540 m=nil [GC sweep wait]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000058f60 sp=0x14000058f40 pc=0x100f3dde8
runtime.goparkunlock(...)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:408
runtime.bgsweep(0x140000a2000)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mgcsweep.go:278 +0xa0 fp=0x14000058fb0 sp=0x14000058f60 pc=0x100f29ed0
runtime.gcenable.gowrap1()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mgc.go:203 +0x28 fp=0x14000058fd0 sp=0x14000058fb0 pc=0x100f1e368
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000058fd0 sp=0x14000058fd0 pc=0x100f712c4
created by runtime.gcenable in goroutine 1
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mgc.go:203 +0x6c

goroutine 20 gp=0x14000094700 m=nil [GC scavenge wait]:
runtime.gopark(0x140000a2000?, 0x1011622d8?, 0x1?, 0x0?, 0x14000094700?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000059760 sp=0x14000059740 pc=0x100f3dde8
runtime.goparkunlock(...)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:408
runtime.(*scavengerState).park(0x10135f020)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mgcscavenge.go:425 +0x5c fp=0x14000059790 sp=0x14000059760 pc=0x100f278cc
runtime.bgscavenge(0x140000a2000)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mgcscavenge.go:653 +0x44 fp=0x140000597b0 sp=0x14000059790 pc=0x100f27e24
runtime.gcenable.gowrap2()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mgc.go:204 +0x28 fp=0x140000597d0 sp=0x140000597b0 pc=0x100f1e308
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000597d0 sp=0x140000597d0 pc=0x100f712c4
created by runtime.gcenable in goroutine 1
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mgc.go:204 +0xac

goroutine 2 gp=0x140000036c0 m=nil [finalizer wait]:
runtime.gopark(0x0?, 0x0?, 0xb8?, 0xc5?, 0x100f6c7e4?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:402 +0xc8 fp=0x1400005c580 sp=0x1400005c560 pc=0x100f3dde8
runtime.runfinq()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mfinal.go:194 +0x108 fp=0x1400005c7d0 sp=0x1400005c580 pc=0x100f1d438
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400005c7d0 sp=0x1400005c7d0 pc=0x100f712c4
created by runtime.createfing in goroutine 1
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/mfinal.go:164 +0x80

goroutine 22 gp=0x140000948c0 m=nil [IO wait]:
runtime.gopark(0xffffffffffffffff?, 0xffffffffffffffff?, 0x23?, 0x0?, 0x100f9e2b0?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/proc.go:402 +0xc8 fp=0x14000059d50 sp=0x14000059d30 pc=0x100f3dde8
runtime.netpollblock(0x0?, 0x0?, 0x0?)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/netpoll.go:573 +0x158 fp=0x14000059d90 sp=0x14000059d50 pc=0x100f378d8
internal/poll.runtime_pollWait(0x1016c8da8, 0x72)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/netpoll.go:345 +0xa0 fp=0x14000059dc0 sp=0x14000059d90 pc=0x100f6b720
internal/poll.(*pollDesc).wait(0x140000ae000?, 0x140000b00a1?, 0x0)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/internal/poll/fd_poll_runtime.go:84 +0x28 fp=0x14000059df0 sp=0x14000059dc0 pc=0x100fb0a28
internal/poll.(*pollDesc).waitRead(...)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).Read(0x140000ae000, {0x140000b00a1, 0x1, 0x1})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/internal/poll/fd_unix.go:164 +0x200 fp=0x14000059e90 sp=0x14000059df0 pc=0x100fb1540
net.(*netFD).Read(0x140000ae000, {0x140000b00a1?, 0x0?, 0x0?})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/fd_posix.go:55 +0x28 fp=0x14000059ee0 sp=0x14000059e90 pc=0x1010387a8
net.(*conn).Read(0x140000b2000, {0x140000b00a1?, 0x0?, 0x0?})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/net.go:179 +0x34 fp=0x14000059f30 sp=0x14000059ee0 pc=0x1010419e4
net.(*TCPConn).Read(0x0?, {0x140000b00a1?, 0x0?, 0x0?})
	<autogenerated>:1 +0x2c fp=0x14000059f60 sp=0x14000059f30 pc=0x10104c3fc
net/http.(*connReader).backgroundRead(0x140000b0090)
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:681 +0x40 fp=0x14000059fb0 sp=0x14000059f60 pc=0x1010da030
net/http.(*connReader).startBackgroundRead.gowrap2()
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:677 +0x28 fp=0x14000059fd0 sp=0x14000059fb0 pc=0x1010d9f18
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000059fd0 sp=0x14000059fd0 pc=0x100f712c4
created by net/http.(*connReader).startBackgroundRead in goroutine 21
	/opt/homebrew/Cellar/go/1.22.1/libexec/src/net/http/server.go:677 +0xc8
exit status 2
```
