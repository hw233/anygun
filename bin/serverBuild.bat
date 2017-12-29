call "%VS90COMNTOOLS%vsvars32.bat"

cmake -G "NMake Makefiles" ../Server/

nmake install