def isprime(num,i=2):
	if num==i:
		return True
	elif num<2:
		return False
	elif(num%i==0):
		return False
	return isprime(num,i+1)

print("Please enter the range to find prime number")
start=int(input("Enter the start range\t"))
end=int(input("Enter the end range\t"))

if(start>end):
	id=-2
	rng=-1
	if(end%2==0):
		end=end-1
else:
	id=+2
	rng=+1
	if(start%2==0):
		start=start+1
print("Prime num between the given range")
for i in range(start,end+rng,id):
	if(isprime(i)):
		print(i)
