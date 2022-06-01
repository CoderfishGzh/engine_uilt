# engine_uilt
## 搜索引擎工具：   
### 布隆过滤器  
使用类似Murmur hash的方法做过滤器哈希映射     
### mmap接口
简单实现了几个接口:   
(1) 分词文件somthing映射到内存   
(2) 文件数组行形式读取   
(3) 字符串分割     
### skiplist

key:分词msg   
vaule:mysql id   

最大层高48
性能测试：   
(1)串行：1千万个kv对，连续插入查询时间大概为16s，推测查询时间2s   
(2)并行：插入查询时间大概6s，推测查询时间1s
Todo：
(1)参考redis高效的randlevel算法   
(2)更高效的查询算法   
