ent=list(map(int,input("Enter list elements\t").split(" ")))
newdict={}
for i in range(len(ent)):
	if ent[i] in newdict:
		if(newdict.get(ent[i])==1):
			newdict[ent[i]]+=1
			print(ent[i])
	else:
		newdict[ent[i]]=1
