# engine_uilt
## 搜索引擎工具：   
### 布隆过滤器  
使用类似Murmur hash的方法做过滤器哈希映射     
### mmap接口
简单实现了几个接口:   
(1) 分词文件somthing映射到内存   
(2) 文件数组行形式读取   
(3) 字符串分割  

### DB接口
(1) `MysqlInit()`用于mysql连接初始化   
(2) `Get_docid_value`根据给的docid查询具体doc   
### skiplist
key，value = (分词msg，mysqlid)   

最大层高48
性能测试：   
cpu: AMD Ryzen 7 5800H with Radeon Graphics   
全部的分词索引数据一共**12193028**个key-value对   
插入索引时间：**1min**左右   
查询key时间：**4us**左右
根据docid查询mysql时间：**0.2s**左右   
Todo：   
(1)参考`redis`高效的`randlevel`算法   
(2)更高效的查询算法   

