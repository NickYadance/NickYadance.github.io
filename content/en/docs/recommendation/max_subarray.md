---
title: "Max Subarray"
description: ""
lead: ""
date: 2023-01-12T16:15:36+08:00
lastmod: 2023-01-12T16:15:36+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "max_subarray-6cbea84dd08055ffe14a57259b9f1016"
weight: 999
toc: true
---
## Equivalent Stock Selling Problem
The max subarray problem can be easily transferred into stock selling problem, in which we need to find the
max profit with buying and selling for only once. Take following stock price array for example
$$[1\ 3\ 7\ 5\ 8\ 2\]$$
The equivalent input for max subarray is
$$[2\ 4\ -2\ 3\ -6\]$$
The answer is 7 for both problems.

## Divide and Conquer
As in [mergesort]({{< relref "docs/recommendation/mergesort.md" >}}), we follow the 3 procedures of divide and conquer.
* _Divide_ : binary-divide the array
* _Resolve_ : when the array holds only 1 element, return it
* _Conquer_ : the result is either the left result, or right result, or crossing dividing element

As with `recursive tree` in [mergesort]({{< relref "docs/recommendation/mergesort.md" >}}), we can get the result of $T(n)=O(n\log(n))$
$$
T(n)=
\begin{cases}
1, &n=1 \\\\
2T(n/2)+n &n>1 \\\\
\end{cases}
$$

{{< details "Divide and Conquer" >}}
```java
public static class Subarray {
  int l;
  int r;
  int value;

  Subarray(int l, int r, int value) {
    this.l = l;
    this.r = r;
    this.value = value;
  }

  Subarray() {
    this.l = this.r = this.value = 0;
  }

  @Override
  public String toString() {
    return "[%d,%d,%d]".formatted(l, r, value);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    Subarray subarray = (Subarray) o;
    return l == subarray.l && r == subarray.r && value == subarray.value;
  }

  @Override
  public int hashCode() {
    return Objects.hash(l, r, value);
  }
}

class Code {
  public static Subarray cross(int[] A, int l, int mid, int r) {
    var res = new Subarray(mid, mid + 1, A[mid]);
    int leftMax = 0, rightMax = 0;
    for (int i = mid - 1, leftSum = 0; i >= l; i--) {
      if ((leftSum += A[i]) > leftMax) {
        res.l = i;
        leftMax = leftSum;
      }
    }
    for (int j = mid + 1, rightSum = 0; j < r; j++) {
      if ((rightSum += A[j]) > rightMax) {
        res.r = j + 1;
        rightMax = rightSum;
      }
    }
    res.value += leftMax + rightMax;
    return res;
  }

  public static Subarray max_subarray(int[] A, int l, int r) {
    if (r - l > 1) {
      int mid = (r + l) / 2;
      var left = max_subarray(A, l, mid);
      var right = max_subarray(A, mid, r);
      var cross = cross(A, l, mid, r);
      return left.value > right.value ?
        (cross.value > left.value ? cross : left) :
        (cross.value > right.value ? cross : right);
    }
    return new Subarray(l, r, A[l]);
  }
}
```
{{< /details >}}

## Kadane algorithm
The [Kadane algorithm](https://en.wikipedia.org/wiki/Maximum_subarray_problem) is the linear complexity solution based on the following
loop variant
$$
\begin{cases}
V: &T(i)=max(\sum^{i}_{k=j}{A[k]}), j \in [0,i] \\\\
S: &T(i)=A[0] \\\\
L: &T(i+1)=T(i)>0\ ?\ T(i)+A[i+1]:A[i+1] \\\\
E: &Answer=max(T(i),i \in 1...n) \\\\
\end{cases}
$$
{{< details "Kadane" >}}
```java
class Subarray {
  public static Subarray max_subarray_linear(int[] A, int l, int r) {
    int max = Integer.MIN_VALUE, sum = 0;
    int left = l, right = l;
    for (int i = l, currentLeft = left; i < r; i++) {
      if (sum > 0) {
        sum += A[i];
      } else {
        sum = A[i];
        currentLeft = i;
      }

      if (sum > max) {
        max = sum;
        left = currentLeft;
        right = i + 1;
      }
    }
    return new Subarray(left, right, max);
  }
}
```
{{< /details >}}
