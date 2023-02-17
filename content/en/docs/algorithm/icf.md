---
title: "Icf"
description: ""
lead: ""
date: 2022-10-17T17:21:21+08:00
lastmod: 2022-10-17T17:21:21+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "icf-ee223455803203ed38ce9450feaabdfb"
weight: 999
toc: true
---
> Chinese special
## 矩阵
* 转置矩阵
把矩阵A的行和列互相交换所产生的矩阵称为A的转置矩阵($A^T$)
* 矩阵相乘
如A是m×n矩阵，B是n×p矩阵，它们的乘积C是一个m×p矩阵
$$ C = AB, c_{ij}=\sum^n_{r=1}{a_{ir}b_{rj}} $$
* 理解矩阵相乘的意义
$$ 3x + 7y = 49 $$
$$ 2x + 4y = 54 $$
的矩阵表现形式为
$$ \begin{pmatrix} 3 & 7 \\\\ 2 & 4 \end{pmatrix} \\ * \begin{pmatrix} x \\\\ y \end{pmatrix} = \begin{pmatrix} 49 \\\\ 54 \end{pmatrix}$$
* 特征值和特征向量
n×n的方块矩阵A的一个特征值和对应特征向量是满足以下方程的标量以及非零向量。其中v为特征向量，$\lambda$为特征值。
$$ Av=\lambda v $$
## 余弦相似度
设坐标点 $(x_1,y_1), (x_2,y_2)$
$$ cos(\theta)=\frac {a^2+b^2-c^2}{2ab}, c=\sqrt {(x_1-x_2)^2+(y_1-y_2)^2} $$
$$ cos(\theta)=\frac {x_1 * x_2+y_1 * y_2}{\sqrt {x_1^2+y_1^2}*\sqrt{x_2^2+y_2^2}} $$
推广到多个点
$$ cos(\theta)=\frac {\sum{x_i * y_i}}{\sqrt{\sum{x_i^2}} * \sqrt{\sum{y_i^2}}} = \frac {\vec{a} * \vec{b}}{|\vec{a}| * |\vec{b}|}$$
## 物品相似度
评分矩阵
$$ \begin{pmatrix} 0 & i0 & i1 & i2 & i3 \\\\ u_0 & r_{00} & r_{01} & r_{02} & r_{03} \\\\ u_1 & r_{10} & r_{11} & r_{12} & r_{13} \\\\ u_2 & r_{20} & r_{21} & r_{22} & r_{2 3} \end{pmatrix} $$
设 $\vec{r_i}$ 为物品i的评分向量(评分矩阵的纵轴)，物品相似度为评分向量的余弦相似度
$$ w_{i,j}=\frac {\vec{r_i} * \vec{r_j}}{|\vec{r_i}| * |\vec{r_j}|} $$
## 线性回归
一元线性回归
$$ y=\beta_0 + \beta_1 x $$
设误差为 $\mu$
$$ y=\beta_0 + \beta_1 x + \mu $$
给定一组观测值 $xi,yi (i=1,2\cdots ,n)$ ，设残差(真实值和预测值之差)为 $e$ ，残差平方和为 $Q$ ，$Q$ 可以作为损失函数
$$ e=y - \hat{y} $$
$$ Q=\sum^n_1{e^2}=\sum^n_1{(y_i - \hat{y_i})^2}=\sum^n_1{(y_i-(\beta_0 + \beta_1 x))^2} $$
最小二乘法求 $Q$ 最小时 $\beta_0, \beta_1$ 的值
## ICF (basic)
设 $\vec{w_i}$ 为物品i的相似度向量(相似矩阵的横轴)，$\vec{r_u}$ 为用户的评分向量(评分矩阵的横轴)
$$ score_{(u,i)}=\frac {\vec{w_i} * \vec{r_u}}{Sum_{\vec{w_i}}} $$

## Reference
[使用余弦相似度算法计算文本相似度](https://zhuanlan.zhihu.com/p/43396514)

[comprehensive-guide-on-item-based-recommendation-systems](https://towardsdatascience.com/comprehensive-guide-on-item-based-recommendation-systems-d67e40e2b75d)

[矩阵的本质](https://www.zhihu.com/question/22047061)

[线性回归](https://zhuanlan.zhihu.com/p/72513104)