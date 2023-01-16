---
title: "Retrofit Okhttp"
description: ""
lead: ""
date: 2022-09-12T18:41:01+08:00
lastmod: 2022-09-12T18:41:01+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "retrofit_okhttp-9617c554ed9ff89a379ee1321ded99c3"
weight: 999
toc: true
---
Hope that I will never bother with making http requests in Java.
```java
@Configuration
@Slf4j
public class Config {
    public Retrofit retrofit(ObjectMapper objectMapper, CustomProperties properties) {
        HttpLoggingInterceptor logging = new HttpLoggingInterceptor(message -> {
            if (log.isDebugEnabled()) {
                log.debug(message);
            }
        });
        logging.setLevel(HttpLoggingInterceptor.Level.BODY);

        OkHttpClient client = new OkHttpClient.Builder()
                .readTimeout(Duration.ofSeconds(5L))
                .writeTimeout(Duration.ofSeconds(5L))
                .connectTimeout(Duration.ofSeconds(3L))
                .addInterceptor(chain -> chain.proceed(chain.request().newBuilder()
                        .header("AUTHORIZATION-NAME", properties.authName)
                        .header("AUTHORIZATION-TOKEN", properties.authToken)
                        .header("DATA-ENVIRONMENT", properties.authEnv)
                        .build()))
                .addInterceptor(logging)
                .build();

        return new Retrofit.Builder()
                .baseUrl(properties.url)
                .client(client)
                .addConverterFactory(JacksonConverterFactory.create(objectMapper))
                .build();
    }
    
    protected interface Service {
        @POST("/api/post")
        Call<Resposne> saveAndSubmit(@Body Request req);

        @GET("/api/query")
        Call<Response> query(@Query(value = "param") String param);
    }
    
    @Bean
    public Service service(Retrofit retrofit) {
        return retrofit().create(Service.class);
    }
}
```