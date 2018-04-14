
d = read.csv("mannenkatu2.csv", header=F)

plot(d)
lines(d)

#plot(d$V5, d$V6, asp=1)
#lines(d$V5, d$V6)

#points(1.0522000094796893,1.2955447271518423, pch="X")
#points(1.032556,1.26164, pch="*")
#points(1.211155,1.260732, pch="*")

#points(1.1773423201278104,0.45720242916684545, pch="+")

#points(1.188869, 0.110574, pch="%")
#points(1.182998, 0.705065, pch="%")

#points(read.csv("output.csv", header=F), pch="x")

q = read.csv("stream.csv", header=F)

points(tail(q, 30), pch="*")
