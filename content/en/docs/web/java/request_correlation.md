---
title: "Request Correlation"
description: ""
lead: ""
date: 2022-09-08T18:43:19+08:00
lastmod: 2022-09-08T18:43:19+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "request_correlation-8e01364a22262fe3aef5aaf5f43c519d"
weight: 999
toc: true
---
Correlate request before and after Spring processing.
* Add `request_id` to request context
* Add `request_id` to log4j MDC
* Add `request_id` to response

Before request
```java
@Configuration
@Slf4j
public class RequestCorrelationConfig {
    public static final String KEY_REQUEST_ID = "X-Request-Id";
    public static final String KEY_START_TIME = "KEY_START_TIME";

    public static class RequestLoggingWithTimeFilter extends CommonsRequestLoggingFilter {
        @Override
        protected boolean shouldLog(HttpServletRequest request) {
            return true;
        }

        @Override
        protected void beforeRequest(HttpServletRequest request, String message) {
            String requestId = request.getHeader(KEY_REQUEST_ID);
            if (StringUtils.isBlank(requestId)) {
                requestId = UUID.randomUUID().toString().replace("-", "");
            }
            MDC.put(KEY_REQUEST_ID, requestId);
            request.setAttribute(KEY_REQUEST_ID, requestId);
            request.setAttribute(KEY_START_TIME, System.currentTimeMillis());
            log.info(message);
        }

        @Override
        protected void afterRequest(HttpServletRequest request, String message) {
            log.info(String.format("cost: [%dms], %s",
                    Optional.ofNullable((Long) request.getAttribute(KEY_START_TIME))
                            .map(s -> System.currentTimeMillis() - s)
                            .orElse(-1L),
                    message));
            // !It's potential that MDC is NOT clear correctly.
            MDC.clear();
        }
    }

    @Bean
    public FilterRegistrationBean<RequestLoggingWithTimeFilter> registerRequestLoggingFilter() {
        var filter = new RequestLoggingWithTimeFilter();
        filter.setIncludeQueryString(true);
        filter.setIncludePayload(true);
        filter.setMaxPayloadLength(10000);
        filter.setIncludeHeaders(false);
        filter.setAfterMessagePrefix("After request: ");

        var bean = new FilterRegistrationBean<>(filter);
        bean.setFilter(filter);
        bean.setOrder(Ordered.HIGHEST_PRECEDENCE);
        return bean;
    }
}
```

After request

```java
@RestControllerAdvice
public class RequestCorrelationResponseBodyAdvice implements ResponseBodyAdvice<Response> {
    @Override
    public boolean supports(MethodParameter returnType, Class<? extends HttpMessageConverter<?>> converterType) {
        return returnType.getParameterType() == Response.class;
    }

    @Override
    public Response beforeBodyWrite(Response body, MethodParameter returnType, MediaType selectedContentType, Class<? extends HttpMessageConverter<?>> selectedConverterType, ServerHttpRequest request, ServerHttpResponse response) {
        if (body != null && request instanceof ServletServerHttpRequest) {
            String requestId = ((ServletServerHttpRequest) request).getServletRequest().getAttribute(RequestCorrelationConfig.KEY_REQUEST_ID).toString();
            return body.requestId(requestId);
        } else {
            return body;
        }
    }
}
```