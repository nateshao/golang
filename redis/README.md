# redis大key优化  

阿里云：https://help.aliyun.com/zh/redis/user-guide/identify-and-handle-large-keys-and-hotkeys/


## 优化大Key
1. 对大Key进行压缩：在保存数据到缓存数据库之前，通过序列化或者压缩算法对大Key对应的value进行压缩，使其占用更小的内存。但如果压缩之后还是特别大，可对大Key进行拆分。
2. 

