ent=list(map(int,input("Enter list elements\t").split(" ")))
newdict={}
count=0
for i in range(len(ent)):
	if ent[i] in newdict:
		count=1
		if(newdict.get(ent[i])==1):
			newdict[ent[i]]+=1
			print(ent[i],end=" ")
	else:
		newdict[ent[i]]=1
if count==1:
	print("Are the repeated elements\t")
else:
	print("No repeted elements")
