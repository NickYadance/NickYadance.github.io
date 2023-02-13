---
title: "Cert"
description: ""
lead: ""
date: 2023-02-13T12:23:57+08:00
lastmod: 2023-02-13T12:23:57+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "cert-3acb9313c60e52beaf81fb6aae961af9"
weight: 999
toc: true
---
## Public key encryption
A message is encrypted with the intended recipient's public key.
For properly chosen and used algorithms, messages cannot in practice be decrypted by anyone who does not possess the matching private key,
who is thus presumed to be the owner of that key and so the person associated with the public key.
This can be used to ensure confidentiality of a message.
![img.png](images/bob_and_alice.png)

## Digital signatures

Digital signatures, in which a message is signed with the sender's private key and can be verified by anyone who
has access to the sender's public key. This verification proves that the sender had access to the private key,
and therefore is very likely to be the person associated with the public key. It also proves that the signature
was prepared for that exact message, since verification will fail for any other message one could devise without using the private key.

![img.png](images/digital_sig.png)

## Certificate authorities
![img.png](images/ca.png)

## TLS procedure
[TLS CSDN](https://blog.csdn.net/ustccw/article/details/76691248)
