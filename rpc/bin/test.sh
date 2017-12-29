# cpp
rm *.cpp
rm *.h
echo "cpp"
./bin -i ./Import.b -o ./ -g cpp
./bin -i ./Example.b -o ./ -g cpp

# c sharp
rm *.cs
echo "cs"
./bin -i ./Import.b -o ./ -g cs
./bin -i ./Example.b -o ./ -g cs

# as3
rm -rf Import
rm -rf Example
echo "as3"
./bin -i ./Import.b -o ./ -g as3
./bin -i ./Example.b -o ./ -g as3

# erl
rm *.erl
rm *.hrl
echo "erl"
./bin -i ./Import.b -o ./ -g erl
./bin -i ./Example.b -o ./ -g erl

#py 
rm *.py
echo "py"
./bin -i ./Import.b -o ./ -g py
./bin -i ./Example.b -o ./ -g py
