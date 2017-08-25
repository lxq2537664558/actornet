set CURR=%cd%
cd ..\..\..\..\..
set GOPATH=%cd%
cd %CURR%

go test -c -o test.exe github.com/davyxu/actornet/tests
@IF %ERRORLEVEL% NEQ 0 pause

: 基础测试

test.exe -test.v -test.run TestHelloWorld
@IF %ERRORLEVEL% NEQ 0 pause

test.exe -test.v -test.run TestEcho
@IF %ERRORLEVEL% NEQ 0 pause

test.exe -test.v -test.run TestRPC
@IF %ERRORLEVEL% NEQ 0 pause

test.exe -test.v -test.run TestSerialize
@IF %ERRORLEVEL% NEQ 0 pause

: 测试网关及后台
start test.exe -test.v -test.run TestLinkGate
start test.exe -test.v -test.run TestLinkBackend
: 防止客户端连接过快
ping -n 2 127.1>nul
test.exe -test.v -test.run TestLinkClient
@IF %ERRORLEVEL% NEQ 0 pause
ping -n 2 127.1>nul

test.exe -test.v -test.run TestAllInOneGate
@IF %ERRORLEVEL% NEQ 0 pause
ping -n 2 127.1>nul

: 跨进程调用
start test.exe -test.v -test.run TestCrossProcessCallServer
test.exe -test.v -test.run TestCrossProcessCallClient
@IF %ERRORLEVEL% NEQ 0 pause
ping -n 2 127.1>nul

: 跨进程通知
start test.exe -test.v -test.run TestCrossProcessNotifyServer
test.exe -test.v -test.run TestCrossProcessNotifyClient
@IF %ERRORLEVEL% NEQ 0 pause
ping -n 2 127.1>nul



del /q test.exe