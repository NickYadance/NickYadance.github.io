---
title: "Matrix Chain"
description: ""
lead: ""
date: 2023-03-16T15:13:53+08:00
lastmod: 2023-03-16T15:13:53+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "matrix_chain-272e95731cef14895f5afbfc025ffca6"
weight: 999
toc: true
---
## Matrix chain and the cutting problem
Tow matrix $A1, A2$ can only multiply when $column_{A1}=row_{A2}$.
$$
A_{p \times q} * A_{q \times r}=A_{p \times r},
$$

The price of this is equal to the number of times each single element get multiplied.
$$Price=p \times q \times r$$

We try to get the min price of multiply for the whole matrix chain.
$$A_1A_2...A_n$$

Now we know that we can resolve this with DP, the thing is WHY and what's the difference between this
and the [Cutting Problem]({{< relref "docs/algorithm/dp" >}}).

Remember that DP needs these
1. The optimal substructure.
2. The repeated subproblems.

The cutting problem has, more obviously, the optimal substructure because each time we cut, we got only one sub problem.
Naturally we just tried all the cutting and resolve the sub problem recursively without even thinking about the word optimal. Actually
during the loop process we already picked the best result from all the sub problems, thus the optimal one. We
can still prove it using the `cut&paste` method.

Things get a little more complicated in the matrix chain as we get two sub problems each time, the left matrix chain and the right
matrix chain. But still, we can try all the cutting point in the matrix chain and get the min price, exactly the same as in the cutting one.

So the key to find out the optimal substructure is **how to get the best result from the best results**. We can simplify the pattern of DP
with this method.

1. Define the answer as the best result.
2. Define how to get the best result from all the possible best results.
3. Utilize some containers to memorize the best result to avoid wasting time.

Now let's try to resolve the matrix chain problem following the 3 steps above.

## DP: Step by step
1. Define the answer as the best result.
The answer is the matrix chain that has the lowest multiply price.

2. Define how to get the best result from all the possible best results.
$$
row_{A_k}=p(k-1), \ column_{A_k}=p(k) \\\\
\begin{aligned}
MinPrice_{[i,j]}=Min(MinPrice_{[i,r]} + MinPrice_{[r+1,j]} + A_{p(i-1) \times p(r)} * A_{p(r) \times p(j)}) \\\\
=Min(MinPrice_{[i,r]} + MinPrice_{[r+1,j]} + p(i-1) \times p(r) \times p(j)) \\\\
\end{aligned}
$$

3. Utilize some containers to memorize the best result to avoid wasting time.
Let m[i, j] be the min price of matrix chain from A_i to A_j.

Persuade code
```text
MinPriceOfMatrixChain(p, n):
    let m = array[n][n]
    for i in 1 to n :
        m[i][i] = 0
    for i in n to 1 :
        for j in i to n :
            for r in i to j :
                m[i][j]=Min(m[i][j], m[i][r] + m[r+1][j] + p(i-1)*p(r)*p(j))
    return m[1][n]
```

Code
```java

```
