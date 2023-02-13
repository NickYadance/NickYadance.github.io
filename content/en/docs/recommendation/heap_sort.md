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
Heapify(root):
    largest = max(root, root.left, root.right)
    if largest != root:
        swapValue(root, largest)
        Heapify(largest)
```


