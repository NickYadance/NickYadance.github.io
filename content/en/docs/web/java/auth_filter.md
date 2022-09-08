---
title: "Auth Filter"
description: ""
lead: ""
date: 2022-09-08T18:48:48+08:00
lastmod: 2022-09-08T18:48:48+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "auth_filter-39960cd7e3682b18c00cec6aeae171c9"
weight: 999
toc: true
---
Typical custom auth filter.
```java
@Component
@Slf4j
@Profile("!test")
public class GoogleAuthFilter extends OncePerRequestFilter {
    @Autowired
    GoogleAuthHolder googleAuthHolder;

    private void reject(HttpServletResponse response) throws IOException {
        response.sendError(HttpServletResponse.SC_UNAUTHORIZED, "unauthorized");
    }

    @SneakyThrows
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain) {
        if (request.getRequestURI().startsWith(AUTH_API_PREFIX)) {
            var op = Optional.ofNullable(request.getHeader(KEY_TOKEN))
                    .map(JWT_PATTERN::matcher)
                    .map(m -> m.matches() ? m.group(1) : null)
                    .map(tokenString -> {
                        try {
                            return VERIFIER.verify(tokenString);
                        } catch (Exception e) {
                            log.warn("GoogleAuth verify failed: {}", tokenString, e);
                            return null;
                        }
                    });
            if (op.isPresent()) {
                var payload = op.get().getPayload();
                googleAuthHolder.auth = new GoogleAuth((String) payload.get("email"),
                        (String) payload.get("sub"),
                        (String) payload.get("name"),
                        (String) payload.get("picture"),
                        (String) payload.get("locale"));
                chain.doFilter(request, response);
            } else {
                reject(response);
            }
        } else {
            chain.doFilter(request, response);
        }
    }
}
```