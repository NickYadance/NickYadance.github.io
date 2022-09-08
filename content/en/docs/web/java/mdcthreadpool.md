---
title: "MdcThreadpool"
description: ""
lead: ""
date: 2022-09-08T18:37:05+08:00
lastmod: 2022-09-08T18:37:05+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "mdcthreadpool-2d94d625842ffdc892fb3b06fd64d984"
weight: 999
toc: true
---
Use `MdcThreadPool` and `MdcWrapper` to pass log4j MDC context cross threads in threadpool, this can help with logging requests with one request id.


```java
/**
 * JBoss, Home of Professional Open Source.
 * Copyright 2014-2022 Red Hat, Inc., and individual contributors
 * as indicated by the @author tags.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
import java.util.Collection;
import java.util.List;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Future;
import java.util.concurrent.RejectedExecutionHandler;
import java.util.concurrent.ThreadFactory;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

/**
 * @author <a href="mailto:matejonnet@gmail.com">Matej Lazar</a>
 * @see https://github.com/project-ncl/pnc/blob/master/common/src/main/java/org/jboss/pnc/common/concurrent/MDCThreadPoolExecutor.java
 */
public class MDCThreadPoolExecutor implements ExecutorService {

    protected ExecutorService executorService;

    MDCThreadPoolExecutor() {
    }

    public MDCThreadPoolExecutor(
            int corePoolSize,
            int maximumPoolSize,
            long keepAliveTime,
            TimeUnit unit,
            BlockingQueue<Runnable> workQueue) {
        this.executorService = new ThreadPoolExecutor(corePoolSize, maximumPoolSize, keepAliveTime, unit, workQueue);
    }

    public MDCThreadPoolExecutor(
            int corePoolSize,
            int maximumPoolSize,
            long keepAliveTime,
            TimeUnit unit,
            BlockingQueue<Runnable> workQueue,
            ThreadFactory threadFactory) {
        this.executorService = new ThreadPoolExecutor(
                corePoolSize,
                maximumPoolSize,
                keepAliveTime,
                unit,
                workQueue,
                threadFactory);
    }

    public MDCThreadPoolExecutor(
            int corePoolSize,
            int maximumPoolSize,
            long keepAliveTime,
            TimeUnit unit,
            BlockingQueue<Runnable> workQueue,
            RejectedExecutionHandler handler) {
        this.executorService = new ThreadPoolExecutor(
                corePoolSize,
                maximumPoolSize,
                keepAliveTime,
                unit,
                workQueue,
                handler);
    }

    @Override
    public <T> Future<T> submit(Callable<T> task) {
        return executorService.submit(MDCWrappers.wrap(task));
    }

    @Override
    public <T> Future<T> submit(Runnable task, T result) {
        return executorService.submit(MDCWrappers.wrap(task), result);
    }

    @Override
    public Future<?> submit(Runnable task) {
        return executorService.submit(MDCWrappers.wrap(task));
    }

    @Override
    public <T> List<Future<T>> invokeAll(Collection<? extends Callable<T>> tasks) throws InterruptedException {
        return executorService.invokeAll(MDCWrappers.wrapCollection(tasks));
    }

    @Override
    public <T> List<Future<T>> invokeAll(Collection<? extends Callable<T>> tasks, long timeout, TimeUnit unit)
            throws InterruptedException {
        return executorService.invokeAll(MDCWrappers.wrapCollection(tasks), timeout, unit);
    }

    @Override
    public <T> T invokeAny(Collection<? extends Callable<T>> tasks) throws InterruptedException, ExecutionException {
        return executorService.invokeAny(MDCWrappers.wrapCollection(tasks));
    }

    @Override
    public <T> T invokeAny(Collection<? extends Callable<T>> tasks, long timeout, TimeUnit unit)
            throws InterruptedException, ExecutionException, TimeoutException {
        return executorService.invokeAny(MDCWrappers.wrapCollection(tasks), timeout, unit);
    }

    @Override
    public void execute(Runnable command) {
        executorService.execute(MDCWrappers.wrap(command));
    }

    @Override
    public void shutdown() {
        executorService.shutdown();
    }

    @Override
    public List<Runnable> shutdownNow() {
        return executorService.shutdownNow();
    }

    @Override
    public boolean isShutdown() {
        return executorService.isShutdown();
    }

    @Override
    public boolean isTerminated() {
        return executorService.isTerminated();
    }

    @Override
    public boolean awaitTermination(long timeout, TimeUnit unit) throws InterruptedException {
        return executorService.awaitTermination(timeout, unit);
    }
}
```

```java
/**
 * JBoss, Home of Professional Open Source.
 * Copyright 2014-2022 Red Hat, Inc., and individual contributors
 * as indicated by the @author tags.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

        import org.slf4j.MDC;

        import java.util.ArrayList;
        import java.util.Collection;
        import java.util.Map;
        import java.util.concurrent.Callable;
        import java.util.function.Consumer;

/**
 * @author <a href="mailto:matejonnet@gmail.com">Matej Lazar</a>
 * @see https://github.com/project-ncl/pnc/blob/master/common/src/main/java/org/jboss/pnc/common/concurrent/MDCWrappers.java
 */
public class MDCWrappers {

    public static Runnable wrap(final Runnable runnable) {
        final var context = MDC.getCopyOfContextMap();
        return () -> {
            var previous = MDC.getCopyOfContextMap();
            if (context == null) {
                MDC.clear();
            } else {
                MDC.setContextMap(context);
            }
            try {
                runnable.run();
            } finally {
                if (previous == null) {
                    MDC.clear();
                } else {
                    MDC.setContextMap(previous);
                }
            }
        };
    }

    public static <T> Callable<T> wrap(final Callable<T> callable) {
        final var context = MDC.getCopyOfContextMap();
        return () -> {
            var previous = MDC.getCopyOfContextMap();
            if (context == null) {
                MDC.clear();
            } else {
                MDC.setContextMap(context);
            }
            try {
                return callable.call();
            } finally {
                if (previous == null) {
                    MDC.clear();
                } else {
                    MDC.setContextMap(previous);
                }
            }
        };
    }

    public static <T> Consumer<T> wrap(final Consumer<T> consumer) {
        final Map<String, String> context = MDC.getCopyOfContextMap();
        return (t) -> {
            Map previous = MDC.getCopyOfContextMap();
            if (context == null) {
                MDC.clear();
            } else {
                MDC.setContextMap(context);
            }
            try {
                consumer.accept(t);
            } finally {
                if (previous == null) {
                    MDC.clear();
                } else {
                    MDC.setContextMap(previous);
                }
            }
        };
    }

    public static <T> Collection<Callable<T>> wrapCollection(Collection<? extends Callable<T>> tasks) {
        Collection<Callable<T>> wrapped = new ArrayList<>();
        for (Callable<T> task : tasks) {
            wrapped.add(wrap(task));
        }
        return wrapped;
    }
}

```