package info

import (
	"strings"
	"encoding/json"
	"fmt"
	"errors"
	"github.com/go-redis/redis"
)

const (
	SERVER = "Server"
	CLIENTS = "Clients"
	MEMORY = "Memory"
	PERS = "Persistence"
	STATS = "Stats"
	REP = "Replication"
	CPU = "CPU"
	CLU = "Cluster"
	KEYSPACE = "Keyspace"
)

type Server struct {
	Redis_version string `json:"redis_version"`
	Redis_git_sha1 string `json:"redis_git_sha_1"`
	Redis_git_dirty string `json:"redis_git_dirty"`
	Redis_build_id string `json:"redis_build_id"`
	Redis_mode string `json:"redis_mode"`
	Os string `json:"os"`
	Arch_bits string `json:"arch_bits"`
	Multiplexing_api string `json:"multiplexing_api"`
	Atomicvar_api string `json:"atomicvar_api"`
	Gcc_version string `json:"gcc_version"`
	Process_id string `json:"process_id"`
	Run_id string `json:"run_id"`
	Tcp_port string `json:"tcp_port"`
	Uptime_in_seconds string `json:"uptime_in_seconds"`
	Uptime_in_days string `json:"uptime_in_days"`
	Hz string `json:"hz"`
	Lru_clock string `json:"lru_clock"`
	Executable string `json:"executable"`
	Config_file string `json:"config_file"`
}

type Clients struct{
	Connected_clients string `json:"connected_clients"`
	Client_longest_output_list string `json:"client_longest_output_list"`
	Client_biggest_input_buf string `json:"client_biggest_input_buf"`
	Blocked_clients string `json:"blocked_clients"`
}

type Memory struct{
	Used_memory string `json:"used_memory"`
	Used_memory_human string `json:"used_memory_human"`
	Used_memory_rss string `json:"used_memory_rss"`
	Used_memory_rss_human string `json:"used_memory_rss_human"`
	Used_memory_peak string `json:"used_memory_peak"`
	Used_memory_peak_human string `json:"used_memory_peak_human"`
	Used_memory_peak_perc string `json:"used_memory_peak_perc"`
	Used_memory_overhead string `json:"used_memory_overhead"`
	Used_memory_startup string `json:"used_memory_startup"`
	Used_memory_dataset string `json:"used_memory_dataset"`
	Used_memory_dataset_perc string `json:"used_memory_dataset_perc"`
	Total_system_memory string `json:"total_system_memory"`
	Total_system_memory_human string `json:"total_system_memory_human"`
	Used_memory_lua string `json:"used_memory_lua"`
	Used_memory_lua_human string `json:"used_memory_lua_human"`
	Maxmemory string `json:"maxmemory"`
	Maxmemory_human string `json:"maxmemory_human"`
	Maxmemory_policy string `json:"maxmemory_policy"`
	Mem_fragmentation_ratio string `json:"mem_fragmentation_ratio"`
	Mem_allocator string `json:"mem_allocator"`
	Active_defrag_running string `json:"active_defrag_running"`
	Lazyfree_pending_objects string `json:"lazyfree_pending_objects"`
}

type Persistence struct{
	Loading string `json:"loading"`
	Rdb_changes_since_last_save string `json:"rdb_changes_since_last_save"`
	Rdb_bgsave_in_progress string `json:"rdb_bgsave_in_progress"`
	Rdb_last_save_time string`json:"rdb_last_save_time"`
	Rdb_last_bgsave_status string `json:"rdb_last_bgsave_status"`
	Rdb_last_bgsave_time_sec string `json:"rdb_last_bgsave_time_sec"`
	Rdb_current_bgsave_time_sec string `json:"rdb_current_bgsave_time_sec"`
	Rdb_last_cow_size string `json:"rdb_last_cow_size"`
	Aof_enabled string `json:"aof_enabled"`
	Aof_rewrite_in_progress string `json:"aof_rewrite_in_progress"`
	Aof_rewrite_scheduled string `json:"aof_rewrite_scheduled"`
	Aof_last_rewrite_time_sec string `json:"aof_last_rewrite_time_sec"`
	Aof_current_rewrite_time_sec string `json:"aof_current_rewrite_time_sec"`
	Aof_last_bgrewrite_status string `json:"aof_last_bgrewrite_status"`
	Aof_last_write_status string `json:"aof_last_write_status"`
	Aof_last_cow_size string `json:"aof_last_cow_size"`
}

type Stats struct{
	Total_connections_received string `json:"total_connections_received"`
	Total_commands_processed string `json:"total_commands_processed"`
	Instantaneous_ops_per_sec string `json:"instantaneous_ops_per_sec"`
	Total_net_input_bytes string `json:"total_net_input_bytes"`
	Total_net_output_bytes string `json:"total_net_output_bytes"`
	Instantaneous_input_kbps string `json:"instantaneous_input_kbps"`
	Instantaneous_output_kbps string `json:"instantaneous_output_kbps"`
	Rejected_connections string `json:"rejected_connections"`
	Sync_full string `json:"sync_full"`
	Sync_partial_ok string `json:"sync_partial_ok"`
	Sync_partial_err string `json:"sync_partial_err"`
	Expired_keys string `json:"expired_keys"`
	Evicted_keys string `json:"evicted_keys"`
	Keyspace_hits string `json:"keyspace_hits"`
	Keyspace_misses string `json:"keyspace_misses"`
	Pubsub_channels string `json:"pubsub_channels"`
	Pubsub_patterns string `json:"pubsub_patterns"`
	Latest_fork_usec string `json:"latest_fork_usec"`
	Migrate_cached_sockets string `json:"migrate_cached_sockets"`
	Slave_expires_tracked_keys string `json:"slave_expires_tracked_keys"`
	Active_defrag_hits string `json:"active_defrag_hits"`
	Active_defrag_misses string `json:"active_defrag_misses"`
	Active_defrag_key_hits string `json:"active_defrag_key_hits"`
	Active_defrag_key_misses string `json:"active_defrag_key_misses"`
}

type Replication struct{
	Role string `json:"role"`
	Connected_slaves string `json:"connected_slaves"`
	Slave0 string `json:"slave0"`
	Master_replid string `json:"master_replid"`
	Master_replid2 string `json:"master_replid_2"`
	Master_repl_offset string `json:"master_repl_offset"`
	Second_repl_offset string `json:"second_repl_offset"`
	Repl_backlog_active string `json:"repl_backlog_active"`
	Repl_backlog_size string `json:"repl_backlog_size"`
	Repl_backlog_first_byte_offset string `json:"repl_backlog_first_byte_offset"`
	Repl_backlog_histlen string `json:"repl_backlog_histlen"`
}

type Cpu struct{
	Used_cpu_sys string `json:"used_cpu_sys"`
	Used_cpu_user string `json:"used_cpu_user"`
	Used_cpu_sys_children string `json:"used_cpu_sys_children"`
	Used_cpu_user_children string `json:"used_cpu_user_children"`
}

type Cluster struct{
	Cluster_enabled string `json:"cluster_enabled"`
}

type Keyspace struct{}


func appointinfo(parm string)(interface{},error){
	var infos []interface{}
	var info interface{}
	switch parm {
	case SERVER:
		info = &Server{}
	case CLIENTS:
		info = &Clients{}
	case MEMORY:
		info = &Memory{}
	case PERS:
		info = &Persistence{}
	case STATS:
		info = &Stats{}
	case REP:
		info = &Replication{}
	case CPU:
		info = &Cpu{}
	case CLU:
		info = &Cluster{}
	case KEYSPACE:
		info = &Keyspace{}
	default:
		return nil,errors.New("parm error")
	}
	vals,err := allinfo()
	for _,val := range vals {
		if err != nil {
			fmt.Println("get info error ", err)
			return nil, err
		}
		result, err := json.Marshal(val)
		if err != nil {
			fmt.Println("json.Marshal err ", err)
			return nil, err
		}
		json.Unmarshal(result, info)
		if err != nil {
			fmt.Println("json.Unmarshal err ", err)
			return nil, err
		}
		infos = append(infos,info)
	}
	return infos,nil
}

func allinfo() ([]map[string]string,error) {
	vals := []map[string]string{}
	for i := 0; i < len(Parm); i++ {
		rc := redisClient(checkIp(Parm[i]))
		defer rc.Close()
		info := rc.Info()
		value, err := info.Result()
		if err != nil {
			fmt.Println("get infomation error ", err)
			return nil,err
		}
		val := map[string]string{}
		for _, v := range strings.Split(value, "\r\n") {
			index := strings.Index(v, ":")
			if index == -1 {
				continue
			}
			val[v[:index]] = v[index+1:]
		}
		vals = append(vals, val)
	}
	return vals,nil
}

func checkIp(ip string)string{
	if strings.Contains(ip,":"){
		return ip
	}
	return ip + ":6379"
}

func redisClient(host string) *redis.Client {
	client := &redis.Client{}
	if Secret != ""{
		client = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: Secret,
		})
	}else{
		client = redis.NewClient(&redis.Options{
			Addr:     host,
		})
	}
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis ping err ", err)
	}
	return client
}