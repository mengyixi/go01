# SET测试

### 10 bytes
```
PING_INLINE: 180505.41 requests per second
PING_BULK: 177935.95 requests per second
SET: 182149.36 requests per second
GET: 179533.22 requests per second
INCR: 178890.88 requests per second
LPUSH: 179856.11 requests per second
RPUSH: 184162.06 requests per second
LPOP: 181488.20 requests per second
RPOP: 180831.83 requests per second
SADD: 180180.17 requests per second
HSET: 176366.86 requests per second
SPOP: 181159.42 requests per second
LPUSH (needed to benchmark LRANGE): 177935.95 requests per second
LRANGE_100 (first 100 elements): 65530.80 requests per second
LRANGE_300 (first 300 elements): 28058.36 requests per second
LRANGE_500 (first 450 elements): 19557.99 requests per second
LRANGE_600 (first 600 elements): 15323.32 requests per second
MSET (10 keys): 121212.12 requests per second
```

### 20 bytes
```
PING_INLINE: 182149.36 requests per second
PING_BULK: 181488.20 requests per second
SET: 180505.41 requests per second
GET: 180831.83 requests per second
INCR: 183823.52 requests per second
LPUSH: 162074.56 requests per second
RPUSH: 183823.52 requests per second
LPOP: 179533.22 requests per second
RPOP: 181159.42 requests per second
SADD: 184162.06 requests per second
HSET: 176056.33 requests per second
SPOP: 182481.77 requests per second
LPUSH (needed to benchmark LRANGE): 180180.17 requests per second
LRANGE_100 (first 100 elements): 65189.05 requests per second
LRANGE_300 (first 300 elements): 27631.94 requests per second
LRANGE_500 (first 450 elements): 19758.94 requests per second
LRANGE_600 (first 600 elements): 15281.17 requests per second
MSET (10 keys): 122850.12 requests per second
```

### 50 bytes
```
PING_INLINE: 183150.19 requests per second
PING_BULK: 183150.19 requests per second
SET: 183823.52 requests per second
GET: 182149.36 requests per second
INCR: 181159.42 requests per second
LPUSH: 171821.30 requests per second
RPUSH: 178571.42 requests per second
LPOP: 176678.45 requests per second
RPOP: 180180.17 requests per second
SADD: 181818.17 requests per second
HSET: 172711.58 requests per second
SPOP: 183150.19 requests per second
LPUSH (needed to benchmark LRANGE): 175438.59 requests per second
LRANGE_100 (first 100 elements): 61124.69 requests per second
LRANGE_300 (first 300 elements): 27173.91 requests per second
LRANGE_500 (first 450 elements): 19623.23 requests per second
LRANGE_600 (first 600 elements): 14494.85 requests per second
MSET (10 keys): 107874.87 requests per second
```
### 100 bytes
```
PING_INLINE: 182815.36 requests per second
PING_BULK: 181488.20 requests per second
SET: 185185.17 requests per second
GET: 180505.41 requests per second
INCR: 184842.88 requests per second
LPUSH: 170648.45 requests per second
RPUSH: 177304.97 requests per second
LPOP: 177304.97 requests per second
RPOP: 178890.88 requests per second
SADD: 182481.77 requests per second
HSET: 179211.45 requests per second
SPOP: 186219.73 requests per second
LPUSH (needed to benchmark LRANGE): 176991.16 requests per second
LRANGE_100 (first 100 elements): 58823.53 requests per second
LRANGE_300 (first 300 elements): 26350.46 requests per second
LRANGE_500 (first 450 elements): 17724.21 requests per second
LRANGE_600 (first 600 elements): 13107.88 requests per second
MSET (10 keys): 105152.48 requests per second
```
### 200 bytes
```
PING_INLINE: 179533.22 requests per second
PING_BULK: 180831.83 requests per second
SET: 181488.20 requests per second
GET: 179211.45 requests per second
INCR: 180180.17 requests per second
LPUSH: 166666.66 requests per second
RPUSH: 176678.45 requests per second
LPOP: 176056.33 requests per second
RPOP: 177619.89 requests per second
SADD: 180505.41 requests per second
HSET: 176678.45 requests per second
SPOP: 181488.20 requests per second
LPUSH (needed to benchmark LRANGE): 172117.05 requests per second
LRANGE_100 (first 100 elements): 49950.05 requests per second
LRANGE_300 (first 300 elements): 19577.13 requests per second
LRANGE_500 (first 450 elements): 11816.14 requests per second
LRANGE_600 (first 600 elements): 7079.14 requests per second
MSET (10 keys): 104275.29 requests per second
```
### 1kb
```
PING_INLINE: 180505.41 requests per second
PING_BULK: 177304.97 requests per second
SET: 176056.33 requests per second
GET: 171821.30 requests per second
INCR: 175746.92 requests per second
LPUSH: 163934.42 requests per second
RPUSH: 164203.61 requests per second
LPOP: 168067.22 requests per second
RPOP: 168918.92 requests per second
SADD: 176056.33 requests per second
HSET: 176991.16 requests per second
SPOP: 178253.12 requests per second
LPUSH (needed to benchmark LRANGE): 170357.75 requests per second
LRANGE_100 (first 100 elements): 23468.67 requests per second
LRANGE_300 (first 300 elements): 4725.90 requests per second
LRANGE_500 (first 450 elements): 2608.92 requests per second
LRANGE_600 (first 600 elements): 1867.52 requests per second
MSET (10 keys): 102249.49 requests per second
```
### 5kb
```
PING_INLINE: 183150.19 requests per second
PING_BULK: 181488.20 requests per second
SET: 167224.08 requests per second
GET: 166944.92 requests per second
INCR: 182149.36 requests per second
LPUSH: 118764.84 requests per second
RPUSH: 114155.25 requests per second
LPOP: 148809.53 requests per second
RPOP: 146198.83 requests per second
SADD: 181818.17 requests per second
HSET: 166389.34 requests per second
SPOP: 184842.88 requests per second
LPUSH (needed to benchmark LRANGE): 142045.45 requests per second
LRANGE_100 (first 100 elements): 3497.48 requests per second
LRANGE_300 (first 300 elements): 950.29 requests per second
LRANGE_500 (first 450 elements): 619.74 requests per second
LRANGE_600 (first 600 elements): 478.25 requests per second
MSET (10 keys): 45662.10 requests per second
```



### 分析
200b以下会有波动，但是GET/SET性能处于同一个水平。
1kb以后GET/SET性能有明显降低。
5kb较1kb有更明显的下降。