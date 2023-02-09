from pyspark.sql import SparkSession

spark = SparkSession.builder.getOrCreate()
sc = spark.sparkContext

import random

M = 100
N = 10000
R = 36
B = ['b' for _ in range(0, 1000)]
mapper = lambda _: [(random.randint(0, N), B) for _ in range(0, N)]
pairs = sc\
    .parallelize(range(M), M) \
    .flatMap(mapper)
pairs.count()
pairs.groupByKey(R).count()