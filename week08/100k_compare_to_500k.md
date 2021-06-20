##  存入10W条KV
- used_memory:8550816                       // Redis分配的内存总量(byte)，包含redis进程内部的开销和数据占用的内存
- used_memory_human:8.00M                   // Redis分配的内存总量(mb)
- used_memory_rss:8509800                   // 向操作系统申请的内存大小(byte)
- used_memory_rss_human:8.00M               // 向操作系统申请的内存大小(mb)
- used_memory_peak:8550816                  // redis的内存消耗峰值(byte)
- used_memory_peak_human:8.00M              // redis的内存消耗峰值(mb)
- used_memory_peak_perc:100.00%             // 使用内存与峰值内存的百分比(used_memory / used_memory_peak) *100%
- used_memory_overhead:6939366              // Redis维护数据集的内部机制所需的内存开销,包括所有客户端输出缓冲区、查询缓冲区、AOF重写缓冲区和主从复制的backlog
- used_memory_startup:661192                // Redis启动完成使用的内存
- used_memory_dataset:1611450               // 数据占用的内存(used_memory - used_memory_overhead)
- used_memory_dataset_perc:20.42%           // 数据占用的内存大小百分比,(used_memory_dataset / (used_memory - used_memory_startup))*100%
- allocator_allocated:9866672               // 分配器分配的内存
- allocator_active:306184192                // 分配器活跃的内存
- allocator_resident:331350016              // 分配器常驻的内存
- total_system_memory:0                     // 主机内存总量(byte)
- total_system_memory_human:0B              // 主机内存总量(mb)
- used_memory_lua:37888                     // Lua引擎存储占用的内存(byte)
- used_memory_lua_human:37.00K              // Lua引擎存储占用的内存(mb)
- used_memory_scripts:0
- used_memory_scripts_human:0B
- number_of_cached_scripts:0
- maxmemory:0                               // 配置中设置的最大可使用内存值(byte),默认0,不限制
- maxmemory_human:0B                        // 配置中设置的最大可使用内存值(mb)
- maxmemory_policy:noeviction               // 当达到maxmemory时的淘汰策略
- allocator_frag_ratio:31.03.02             // 分配器的碎片率
- allocator_frag_bytes:296317520            // 分配器的碎片大小
- allocator_rss_ratio:1.08.32               // 分配器常驻内存比例
- allocator_rss_bytes:25165824              // 分配器的常驻内存大小
- rss_overhead_ratio:0.03.57                // 常驻内存开销比例
- rss_overhead_bytes:-322840216166776832    // 常驻内存开销大小
- mem_fragmentation_ratio:1.00.53           // 碎片率(used_memory_rss / used_memory),正常(1,1.6),大于比例说明内存碎片严重
- mem_fragmentation_bytes:0                 // 内存碎片大小
- mem_not_counted_for_evict:0               // 被驱逐的内存
- mem_replication_backlog:0                 // redis复制积压缓冲区内存
- mem_clients_slaves:0                      // Redis节点客户端消耗内存
- mem_clients_normal:49950                  // Redis所有常规客户端消耗内存
- mem_aof_buffer:0                          // AOF使用内存
- mem_allocator:jemalloc-5.2.1-redis-5.1.0  // 内存分配器
- active_defrag_running:0                   // 活动碎片整理是否处于活动状态(0没有,1正在运行)
- lazyfree_pending_objects:0                // 0-不存在延迟释放的挂起对象

## 存入50W条KV

- used_memory:25219928
- used_memory_human:24.00M
- used_memory_rss:25178912
- used_memory_rss_human:24.00M
- used_memory_peak:25238344
- used_memory_peak_human:24.00M
- used_memory_peak_perc:99.93%
- used_memory_overhead:18808294
- used_memory_startup:661192
- used_memory_dataset:6411634
- used_memory_dataset_perc:26.11%
- allocator_allocated:25348376
- allocator_active:314572800
- allocator_resident:322961408
- total_system_memory:0
- total_system_memory_human:0B
- used_memory_lua:37888
- used_memory_lua_human:37.00K
- used_memory_scripts:0
- used_memory_scripts_human:0B
- number_of_cached_scripts:0
- maxmemory:0
- maxmemory_human:0B
- maxmemory_policy:noeviction
- allocator_frag_ratio:12.41
- allocator_frag_bytes:289224424
- allocator_rss_ratio:1.03
- allocator_rss_bytes:8388608
- rss_overhead_ratio:0.08
- rss_overhead_bytes:-297782496
- mem_fragmentation_ratio:1.00
- mem_fragmentation_bytes:0
- mem_not_counted_for_evict:0
- mem_replication_backlog:0
- mem_clients_slaves:0
- mem_clients_normal:49950
- mem_aof_buffer:0
- mem_allocator:jemalloc-5.2.1-redis
- active_defrag_running:0
- lazyfree_pending_objects:0

## 分析
存入值相同的10W条KV和50W条KV相比：  
存入10W条平均每条占用内存约16个比特，数据内存占用率是20.42%。  
存入50W条平均每条占用内存约13个比特，数据内存占用率是26.11%。  
