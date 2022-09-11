---
title: "Exception Handler"
description: ""
lead: ""
date: 2022-09-08T18:47:33+08:00
lastmod: 2022-09-08T18:47:33+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "exception_handler-f1aedd5516f96c5b240a2f9906574d85"
weight: 999
toc: true
---
```java
@RestControllerAdvice
@Slf4j
public class GlobalExceptionHandler  {
    @ExceptionHandler({BusinessException.class})
    protected Response<Object> handleBusinessException(Exception ex, WebRequest request) throws Exception {
        log.error("business error", ex);
        if (ex instanceof BusinessException be) {
            return Response.fail(be.code, be.message);
        } else if (ex.getCause() instanceof BusinessException cbe) {
            return Response.fail(cbe.code, cbe.message);
        } else {
            throw ex;
        }
    }
}
```