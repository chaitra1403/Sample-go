import time

def isprime(num,i=2):
	if num==i:
		return True
	elif num<2:
		return False
	elif(num%i==0):
		return False
	return isprime(num,i+1)

s=int(input())
e=int(input())
start=time.time()
if(s%2==0):
	s=s+1
for i in range(s,e+1,+2):
	if(isprime(i)):
		print(i)
end=time.time()
ss=round(start,2)
ee=round(end,2)
print(round(ee-ss,6),"optimal")
