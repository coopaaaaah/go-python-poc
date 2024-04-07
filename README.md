# go-python-poc
POC for Go/Python Interop (Python code executes Go code)

Prereq
- Python installed (currently using 3.8.11)
- Go installed (currently using 1.22.2)

Build go code with 
```
go build -buildmode=c-shared -o libadd.so libadd.go
```

Execute python code with
```
python3 main.py
```
