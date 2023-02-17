---
title: "Heap Sort"
description: ""
lead: ""
date: 2023-02-11T21:47:45+08:00
lastmod: 2023-02-11T21:47:45+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "heap_sort-76bd868b47ab4c7584131c1c2273174a"
weight: 999
toc: true
---
## Tree representation of array
The array(A) can be represented with a binary tree with A[0] as root, A[1] as left child of root, A[2] as right child
of root...
![img.png](images/array_in_tree.png)

## Heapify
The first step of heap sort is to build a heap from an array by running `heapify` recursively.

The `heapify`(max) process is to retain the follow variant. For $tree(root) $
$$
\begin{cases}
&root.val >= any(tree(root.left)) \\\\
&root.val >= any(tree(root.right)) \\\\
\end{cases}
$$

```text
Heapify(root, size):
    largest = root
    if root.left < size :
        largest = indexOf(max(root.val, root.left.val))
    if root.right < size :
        largest = indexOf(max(root.val, root.right.val))
    if largest != root:
        swapValue(root, largest)
        Heapify(largest)

BuildHeap(A):
    N = A.length
    for i = size / 2 to 0:
        Heapify(i, 2*i, 2*i + 1)
```

## Pop and swap
Now that for a fully heapfied array, pops the max element and swaps it to begin/end, and run heapify
against the remain part of the array.
```text
Sort(A):
    for i = A.length - 1 to 0:
        swap(A, 0, i)
        heapify(0, i)
```

## Code
{{< details "Heapsort" >}}
```java
public static void heapify(int[] A, int root, int size) {
      int largest = root;
      int left = 2 * root + 1, right = left + 1;
      if (left < size && A[left] > A[largest]) {
          largest = left;
      }
      if (right < size && A[right] > A[largest]) {
          largest = right;
      }

      if (root != largest) {
          swap(A, root, largest);
          heapify(A, largest, size);
      }
}

public static void buildHeap(int[] A) {
    int size = A.length;
    for (int i = size / 2 - 1; i >= 0; i--) {
        heapify(A, i, size);
    }
}

public static void heapSort(int[] A) {
    buildHeap(A);
    int size = A.length;
    for (int i = size - 1; i >= 0; i--) {
        swap(A, 0, i);
        heapify(A, 0, --size);
    }
}
```
{{< /details >}}
