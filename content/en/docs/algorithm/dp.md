---
title: "DP"
description: ""
lead: ""
date: 2023-02-21T21:47:11+08:00
lastmod: 2023-02-21T21:47:11+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "dp-89ac0a5f7365bd60f19a94f0992ba938"
weight: 999
toc: true
---
## Complete walkthrough of DP: cutting problem
The cutting problem is to get the max profit while cutting the stick of length n $P(n)$, into many parts with any length, each
length of part have its own price $p(i)$.

### DP and divide&conquer
The divide&conquer divides the problem into subproblems and merge their results, all the subproblems never **repeat** nor **overlap**.

The DP(dynamic programming) divides the problem into subproblems and some subproblems can **repeat** but not **overlap**.

When repeatable subproblems happen, DP utilizes the `memorized table` to get the result instantly to avoid wasting time.

The DP applies to the problems that are:
1. Having the optimal substructure.
2. Having the repeated subproblems.
3. Having no overlap subproblems.

### The explosive recursion
{{< details "Cutting problem - recursion" >}}
```java
public static int cuttingProblem_recursion(int length, int[] record) {
      int maxProfit = 0;
      for (int i = 1; i <= length; i++) {
          int profit = p[i] + cuttingProblem_recursion(length - i, record);
          if (profit > maxProfit) {
              maxProfit = profit;
              record[length] = i;
          }
      }
      return maxProfit;
}
```
{{< /details >}}

For the cutting problem, the simple approach is to try all the possible cuttings.
$$ P(n)=Max(p(i) + P(n-i)), i \in [1,n] $$

The recursive tree of the simple approach looks like this with repeated subproblems exist: $P(2), \ P(1), \ P(0)$

![img.png](images/dp_cutting_recursive_tree.png)

$T(n)$ is the number of times that `cuttingProblem_recursive` get called.
$$ T(n) = 1 + \sum_{i=0}^{n-1}{T(i)}, \ T(0)=1 $$
$$ T(n) = 2^n $$

### The graphical representation of DP: the subproblems graph
{{< details "Cutting problem - dp - up" >}}
```java
public static int cuttingProblem_dp_up(int length, int[] record) {
        int[] memorizedTable = new int[1 + length];
        for (int i = 1; i <= length; i++) {
            int maxProfit = 0;
            for (int j = 1; j <= i; j++) {
                int profit = p[j] + memorizedTable[i - j];
                if (profit > maxProfit) {
                    maxProfit = profit;
                    record[i] = j;
                }
                memorizedTable[i] = maxProfit;
            }
        }
        return memorizedTable[length];
    }

```
{{< /details >}}

{{< details "Cutting problem - dp - down" >}}
```java
public static int cuttingProblem_dp_down(int length, int[] record, int[] memorizedTable) {
        if (memorizedTable[length] > 0) {
            return memorizedTable[length];
        }

        int maxProfit = 0;
        for (int i = 1; i <= length; i++) {
            int profit = p[i] + cuttingProblem_recursive(length - i, record);
            if (profit > maxProfit) {
                maxProfit = profit;
                record[length] = i;
            }
        }
        memorizedTable[length] = maxProfit;
        return maxProfit;
    }
```
{{< /details >}}

The DP tries to **collapse** the recursive tree into `the subproblems graph` and avoid the repeated subproblems.

![img.png](images/dp_cutting_dp.png)


### The optimal substructure prove: cut&paste
It's not hard to figure out the substructure of the problem, the question is how to prove it being the optimal substructure.

In DP, the `cut&paste` approach is often used to prove the optimal substructure. Take the cutting problem for example:
$$ P(n)=Max(p(i) + P(n-i)), i \in [1,n] $$

1. Suppose $P(n)$ is the answer for cutting(n), prove $P(n-i)$ is the answer for cutting(n-i).
2. Suppose $P(n-i)$ is not the answer for cutting(n-i), then we cut $P(n-i)$ and paste $P^'(n-i)$ into the answer.
3. Now we have $P^'(n-i) + p(i)$ as answer to cutting(n) which validates $P(n)$ being the answer.

It might seem that all problems with substructure can be proved with `cut&paste`. A negative example is the longest simple
path in a graph.

![img.png](images/longest_simple_path.png)

The path $q \to r \to t$ is the answer for path(q,t), but $q \to s \to t \to r$ is the answer for path(q,r).

The reason `cut&paste` not working is that the subproblems overlapped. The $t$ appears in both answer to path(q,t) and path(q,r),
which violates the definition of the longest simple path.
