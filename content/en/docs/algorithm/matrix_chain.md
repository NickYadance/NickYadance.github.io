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
