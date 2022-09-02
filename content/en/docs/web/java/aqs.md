---
title: "AQS"
description: ""
lead: ""
date: 2022-09-02T13:15:50+08:00
lastmod: 2022-09-02T13:15:50+08:00
draft: true
images: []
menu:
  docs:
    parent: ""
    identifier: "aqs-1361780ae7d5132394755520fdcf485f"
weight: 999
toc: true
---

AQS *AbstractQueueSynchronizer*
无锁策略 #无锁
回忆下[[锁升级]]的若干内容，可以发现JVM对锁竞争的优化策略是基于**零竞争**的前提。对于实际发生竞争的场景，锁仍然是低效的。而AQS要解决的是并发场景同步的问题，不能简单的将零竞争作为前提，因此设计上需要更高效的应对竞争处理以具有扩展性。
基本操作
acquire
``` java
		  while (synchronization state does not allow acquire) {
		      enqueue current thread if not already queued;
		      possibly block current thread;
		  }
		  dequeue current thread if it was queued;
```
release
``` java
		  update synchronization state;
		  if (state may permit a blocked thread to acquire)
		      unblock one or more queued threads;
```
CLH队列
CLH队列的每个节点都会存储后继节点的状态。
新节点入队
``` java
		  do {
		      pred = tail;
		  } while ( !tail.compareAndSet(pred, node) );
```
acquire
``` java
		  // 1. 尝试获取资源
		  // @see barging FIFO
		  if (!tryAcquire(arg)) {
		  // 2. 获取不到，初始化队列（首次）并入队
		      node = create and enqueue new node;
		      pred = node effective predecessor;
		  /*
		  1. 开始自旋，当节点满足
		    a. 是头节点
		    b. 可获取到锁
		  则出队，否则
		    a. CAS设置节点状态为中断
		    b. 然后 park()
		  当节点被唤醒(unpark)时，如果
		    a. 节点为头节点
		    b. 可获取到资源
		  则出队(head = node)
		  */
		  while (pred is not head node || !tryAcquire(arg)){
		    if (pred signal bit is set)
		        park();
		    else
		        compareAndSet preds signal bit to true;
		    pred = node effective predecessor;
		  }
		  head = node;
		  }
```
release
``` java
		  // 1. 变更状态为release后，唤醒队头元素（FIFO）
		  if (tryRelease(arg) && head node signal bit is set) {
		      // 2. CAS队头信号为false
		      compareAndSet head signal bit to false;
		      unpark head successor, if one exists;
		  }
```
condition queues 条件队列
条件队列与[[monitor]]的等待队列非常类似。
内部维护一个条件队列，在获取锁的情况下，线程调用 await，线程会被放置在条件队列中并被阻塞。直到调用 signal、signalAll 唤醒线程，此后线程唤醒，会放入到 AQS 的同步队列，参与争抢锁资源。
condition queue basic await
``` java
		  create and add new node to condition queue;
		  release lock;
		  block (by invoking `park` method) until node is on lock queue;
		  re-acquire lock;
```
condition queue basic signal
``` java
		  transfer the first node from condition queue to lock queue;
```
barging FIFO
由于**tryAcquire()** 方法并没有互斥，因此 AQS 不是严格遵守FIFO的顺序。因此可能出现^^新acquire到资源的线程，抢占当前队头的线程^^。
两种方式解决AQS的不公平性问题
限制获取资源的节点为队列头节点
在队列为空时，允许进入
> In this case,multiple threads encountering an empty queue may race to be thefirst to acquire, normally without enqueuing at least one of them.
在这种情况下，遇到空队列的多个线程可能会竞争第一个获得，通常情况下，至少有一个线程不需要排队。

示例：简单互斥场景
``` java
	  /**
	   * 基于AQS的简单互斥锁
	   */
	  public class SimpleAQS extends AbstractQueuedSynchronizer {
	      @Override
	      protected boolean tryAcquire(int arg) {
	          return compareAndSetState(0, 1);
	      }

	      @Override
	      protected boolean tryRelease(int arg) {
	          return compareAndSetState(1, 0);
	      }

	      public void lock() {
	          acquire(0);
	      }

	      public void unlock() {
	          release(0);
	      }

	      public static void main(String[] args) throws IOException {
	          SimpleAQS aqs = new SimpleAQS();
	          new Thread(() -> runnable(aqs), "Thread[0]").start();
	          new Thread(() -> runnable(aqs), "Thread[1]").start();
	          System.in.read();
	      }

	      private static void runnable(SimpleAQS aqs) {
	          System.out.println(Thread.currentThread().getName() + ": lock tryAcquire");
	          aqs.lock();
	          System.out.println(Thread.currentThread().getName() + ": lock acquired!");
	          try {
	              Thread.sleep(5000);
	          } catch (InterruptedException e) {
	              e.printStackTrace();
	          }
	          aqs.unlock();
	          System.out.println(Thread.currentThread().getName() + ": lock released");
	      }
	  }
```
示例：ReentrantLock
ReentrantLock获取资源时，允许新线程或持有线程重复累加资源计数器并通过，其他情况等待。释放资源时，如果不是持有线程直接抛错，否则递减计数器并当计数器为0时清除持有线程。
获取资源
``` java
		  final boolean nonfairTryAcquire(int acquires) {
		        final Thread current = Thread.currentThread();
		        int c = getState();
		        if (c == 0) {
		            // 1. 线程首次获取资源，直接获取，并设置资源持有者为当前线程
		            if (compareAndSetState(0, acquires)) {
		                setExclusiveOwnerThread(current);
		                return true;
		            }
		        }
		        // 2. 线程非首次获取资源，变更资源状态
		        else if (current == getExclusiveOwnerThread()) {
		            int nextc = c + acquires;
		            if (nextc < 0) // overflow
		                throw new Error("Maximum lock count exceeded");
		            setState(nextc);
		            return true;
		        }
		        return false;
		    }
```
释放资源
``` java
		  protected final boolean tryRelease(int releases) {
		        int c = getState() - releases;
		        // 1. 校验当前线程是否为资源持有者
		        if (Thread.currentThread() != getExclusiveOwnerThread())
		            throw new IllegalMonitorStateException();
		        boolean free = false;
		        // 2. 所有资源持有者都被释放时，才可以释放该线程
		        if (c == 0) {
		            free = true;
		            setExclusiveOwnerThread(null);
		        }
		        setState(c);
		        return free;
		    }
```
示例：CountDownLatch
CountDownLatch的计数器的最大值为用户定义，在 *countdown* 方法中允许线程递减此值，但只在 *await* 方法中允许单个线程作为排队线程获取资源。
获取资源
``` java
		  // 1. getState() 代表当前计数器的值
		  // 2. 当计数器的值 > 0 时，会阻塞等待资源，即 CountDownLatch 的等待原理
		  protected int tryAcquireShared(int acquires) {
		        return (getState() == 0) ? 1 : -1;
		  }
```
释放资源
``` java
		  // 1. 当 state == 0 时，即计数器归零时，资源（虚拟）才能释放，等待计数器的线程会被唤醒
		  protected boolean tryReleaseShared(int releases) {
		        // Decrement count; signal when transition to zero
		        for (;;) {
		            int c = getState();
		            if (c == 0)
		                return false;
		            int nextc = c-1;
		            if (compareAndSetState(c, nextc))
		                return nextc == 0;
		        }
		    }
```
计数操作
``` java
		  // 1. 递减计数器
		  public void countDown() {
		    sync.releaseShared(1);
		  }
```
等待计数器
``` java
		  // 1. CountDownLatch 理论上只有一个排队的线程
		  public boolean await(long timeout, TimeUnit unit)
		    throws InterruptedException {
		    return sync.tryAcquireSharedNanos(1, unit.toNanos(timeout));
		  }
```
