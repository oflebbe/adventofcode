
stop=585159 
start=134564
#start=0
#stop=999999
    
def f(level, value, lastdigit,d,group,firstgroup):
    if level == 0:
        if not d or  value > stop or value < start:
           return 0
        else:
           print(value)
           return 1
    num=0
    for dig in range(lastdigit,10):
        if not firstgroup and (lastdigit == dig) :
            if group-1 <= 0:
                continue
            else:
                group-=1
        else:
            firstgroup=False
            group = 2
        d2 = d or (lastdigit == dig)
        num += f(level-1, value*10+dig, dig, d2, group,firstgroup) 
    return  num
    
print(f(6,0,0,False,2,True))
 