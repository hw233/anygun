
#taskkill/im worldd.exe /f 
taskkill/im logind.exe /f 
taskkill/im dbd.exe /f 
taskkill/im gatewayd.exe /f 
#taskkill/im malld.exe /f
taskkill/im scened.exe /f
#start worldd.exe
#@ping -n 10 127.1>nul
start logind.exe
start dbd.exe
start gatewayd.exe
#start malld.exe
start scened.exe
