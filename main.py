from ctypes import cdll
lib = cdll.LoadLibrary('./libadd.so')
result = lib.add(2, 3)
print(result)

lib.async_add()
lib.slow_add()
