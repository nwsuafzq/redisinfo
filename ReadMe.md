
### 镜像上传API和版本

本程序为前端页面提供镜像上传RESTful API。
镜像上传API版本v1。

### API接口

####  GET /metrics
##### 获取redis全部信息

###### 测试
curl -i "127.0.0.1:10012/metrics"
###### 返回格式(JSON)
{"metrics":[{"active_defrag_hits":"0","active_defrag_key_hits":"0","active_defrag_key_misses":"0","active_defrag_misses":"0","active_defrag_running":"0","aof_current_rewrite_time_sec":"-1","aof_enabled":"0","aof_last_bgrewrite_status":"ok","aof_last_cow_size":"0","aof_last_rewrite_time_sec":"-1","aof_last_write_status":"ok","aof_rewrite_in_progress":"0","aof_rewrite_scheduled":"0","arch_bits":"64","atomicvar_api":"atomic-builtin","blocked_clients":"0","client_biggest_input_buf":"0","client_longest_output_list":"0","cluster_enabled":"1","config_file":"/usr/local/etc/redis.conf","connected_clients":"2","connected_slaves":"0","evicted_keys":"0","executable":"/data/redis-server","expired_keys":"0","gcc_version":"4.8.5","host":"10.1.234.36:29824","hz":"10","instantaneous_input_kbps":"0.00","instantaneous_ops_per_sec":"0","instantaneous_output_kbps":"0.02","keyspace_hits":"0","keyspace_misses":"0","latest_fork_usec":"125","lazyfree_pending_objects":"0","loading":"0","lru_clock":"4481837","master_host":"10.1.234.36","master_last_io_seconds_ago":"2","master_link_status":"up","master_port":"29092","master_repl_offset":"1985690","master_replid":"df07e82a3d7d983958d5e124dfcfb1efcc4af1ae","master_replid2":"0000000000000000000000000000000000000000","master_sync_in_progress":"0","maxmemory":"0","maxmemory_human":"0B","maxmemory_policy":"noeviction","mem_allocator":"libc","mem_fragmentation_ratio":"1.29","migrate_cached_sockets":"0","multiplexing_api":"epoll","os":"Linux 3.10.0-327.el7.x86_64 x86_64","process_id":"1","pubsub_channels":"0","pubsub_patterns":"0","rdb_bgsave_in_progress":"0","rdb_changes_since_last_save":"0","rdb_current_bgsave_time_sec":"-1","rdb_last_bgsave_status":"ok","rdb_last_bgsave_time_sec":"0","rdb_last_cow_size":"221184","rdb_last_save_time":"1531119575","redis_build_id":"a7ab6f4b4bae0753","redis_git_dirty":"0","redis_git_sha1":"00000000","redis_mode":"cluster","redis_version":"4.0.2","rejected_connections":"0","repl_backlog_active":"1","repl_backlog_first_byte_offset":"1861861","repl_backlog_histlen":"123830","repl_backlog_size":"1048576","role":"slave","run_id":"9e8256bc33405fb6b306bf0ebc286bbfd50dbd01","second_repl_offset":"-1","slave_expires_tracked_keys":"0","slave_priority":"100","slave_read_only":"1","slave_repl_offset":"1985690","sync_full":"3","sync_partial_err":"3","sync_partial_ok":"1","tcp_port":"6379","total_commands_processed":"1611835","total_connections_received":"303","total_net_input_bytes":"56314276","total_net_output_bytes":"762757570","total_system_memory":"8372051968","total_system_memory_human":"7.80G","uptime_in_days":"17","uptime_in_seconds":"1553828","used_cpu_sys":"1045.72","used_cpu_sys_children":"0.00","used_cpu_user":"822.75","used_cpu_user_children":"0.00","used_memory":"2642782","used_memory_dataset":"152127","used_memory_dataset_perc":"12.00%","used_memory_human":"2.52M","used_memory_lua":"37888","used_memory_lua_human":"37.00K","used_memory_overhead":"2490655","used_memory_peak":"6573475","used_memory_peak_human":"6.27M","used_memory_peak_perc":"40.20%","used_memory_rss":"3416064","used_memory_rss_human":"3.26M","used_memory_startup":"1375583"}],"status":0}

"status":表示请求状态0表示成功，-1代表失败

#### GET /metrics/:info
##### 获取redis指定信息
###### info应为如下值
**Server** 获取Server信息

**Clients** 获取Clients信息

**Memory** 获取Memory信息

**Persistence** 获取Persistence信息

**Stats** 获取Stats信息

**Replication** 获取Replication信息

**CPU** 获取CPU信息

**Cluster** 获取Cluster信息

**Keyspace** 获取Keyspace信息

###### 测试
curl -i "127.0.0.1:10012/metrics/Server"
###### 返回格式(JSON)
{"metrics":[{"redis_version":"4.0.2","redis_git_sha_1":"","redis_git_dirty":"0","redis_build_id":"a7ab6f4b4bae0753","redis_mode":"cluster","os":"Linux 3.10.0-327.el7.x86_64 x86_64","arch_bits":"64","multiplexing_api":"epoll","atomicvar_api":"atomic-builtin","gcc_version":"4.8.5","process_id":"1","run_id":"9e8256bc33405fb6b306bf0ebc286bbfd50dbd01","tcp_port":"6379","uptime_in_seconds":"1553996","uptime_in_days":"17","hz":"10","lru_clock":"4482005","executable":"/data/redis-server","config_file":"/usr/local/etc/redis.conf"}],"status":0}

"status":表示请求状态0表示成功，-1代表失败
### 运行镜像
**docker run -p <宿主端口号>:10012 <镜像名称>:<版本Tag> --server "10012" --auth "" ip1,ip2...ipn**

**注：--server "10012" 为服务的端口号，不填默认为10012**

**--auth "" 为redis密码**

**ip1 为redis得IP地址，可填多个,传入IP时应带端口号如"10.1.234.36:30323"，若不填择端口号为redis默认的6443端口**

**注：** 镜像中端口号为8080
