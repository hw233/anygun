

taskkill/im world.exe /f 
taskkill/im login.exe /f 
taskkill/im db.exe /f 
taskkill/im scene.exe /f 
taskkill/im gateway.exe /f 

start world.exe
#@ping -n 10 127.1>nul
start login.exe
start scene.exe
start db.exe
start gateway.exe


