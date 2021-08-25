#ifndef macro_common_h__
#define macro_common_h__


#define	NET_XML_MAX_LEN		1024
#define NET_DATA_HEADER_LEN	4

#define USE_BASE_PORT		10000
#define USE_MAX_PORT		35000

#define SERVER_RANDOM_PORT		\
    (USE_BASE_PORT+(int)((float)(USE_MAX_PORT-USE_BASE_PORT)*rand()/(RAND_MAX+2.6)))

#define make_strm_ch(ch_idx, strm_mode)		( (((strm_mode & 0x03) << 6) & 0xc0) | (ch_idx & 0x3f) )

//(low,up) #define Random (rand()%(up-low+1)) + low - 1 
//[low,up) #define Random (rand()%(up-low)) + low 
//(low,up] #define Random (rand()%(up-low))+ low + 1 
//[low,up] #define Random (rand()%(up-low+1)) + low 
#define MACRO_RAND(low,up) ((rand()%(up-low+1)) + low)

#define MAX_A_FRAME_LENGTH						(900 * 1024)

#define MACRO_CONNECT_TIMEOUT_SEC	5
#define MACRO_INTERVAL_SECS  25
#define MACRO_INTERVAL_SECS_EX  (MACRO_INTERVAL_SECS + 5)

#define MACRO_KIND_BC_NONE	0
#define MACRO_KIND_BC_FILE	1
#define MACRO_KIND_BC_VOICE	2
#define MACRO_KIND_BC_CARD	3
#define MACRO_KIND_BC_USR	4
#define MACRO_KIND_TK		5
#define MACRO_KIND_MT		6
#define MACRO_KIND_MONITOR	7
#define MACRO_KIND_MP3_BC_USR   9

#define MACRO_TRANS_PROTO_TCP				1
#define MACRO_TRANS_PROTO_UDP				2
#define MACRO_TRANS_PROTO_UDP_MULTICAST		3
#define MACRO_TRANS_PROTO_RTP				4
#define MACRO_TRANS_PROTO_RTP_MULTICAST		5

#define MACRO_IO_CH_GPIO0_1	1
#define MACRO_IO_CH_GPIO0_2	2
#define MACRO_IO_CH_GPIO6_3	3
#define MACRO_IO_CH_GPIO6_4	4

#define MACRO_MD_TP_VID		1 //only video
#define MACRO_MD_TP_AUD		2 //only audio
#define MACRO_MD_TP_AV		3 //video && audio

#define MACRO_USR_SIG_TYPE_MD		1	//motion detection
#define MACRO_USR_SIG_TYPE_IO_IN	2	//io input
#define MACRO_USR_SIG_TYPE_DB_ALM	3	//sound db alarm

#define MACRO_TSK_STOP_CODE		(-9)

#define MACRO_AUD_CHANNEL	1
#define MACRO_AUD_BITRATE	16
#define MACRO_AUD_SAMPLE_8K		8000
#define MACRO_AUD_SAMPLE_16K	16000
#define MACRO_AUD_SAMPLE_32K	32000
#define MACRO_AUD_SAMPLE_44K	44100

#define MACRO_SYNC_TP_NO	0
#define MACRO_SYNC_TP_FRM	1
#define MACRO_SYNC_TP_SEC	2

#define GPIO6_4 1
#define GPIO6_3 2
#define GPIO0_1 3	//1是控制大功放
#define GPIO0_2 4	//2是控制小功放

#define STEM_MAIN 0
#define STEM_SUB  1

#define MAX_A_FRAME_LENGTH		(900 * 1024)
#define MACRO_EVENT_READING		0x01	/**< error encountered while reading */
#define MACRO_EVENT_WRITING		0x02	/**< error encountered while writing */
#define MACRO_EVENT_DISP_END	0x04	/**< eof dispatch reached */
#define MACRO_EVENT_RESERVED	0x08	/**< event reserved */
#define MACRO_EVENT_EOF			0x10	/**< eof file reached */
#define MACRO_EVENT_ERROR		0x20	/**< unrecoverable error encountered */
#define MACRO_EVENT_TIMEOUT		0x40	/**< user-specified timeout reached */
#define MACRO_EVENT_CONNECTED	0x80	/**< connect operation finished. */

#define MACRO_EVENT_READING_S	"net_read"
#define MACRO_EVENT_WRITING_S	"net_write"
#define MACRO_EVENT_DISP_END_S	"svc_end"
#define MACRO_EVENT_RESERVED_S	"reserved"
#define MACRO_EVENT_EOF_S		"net_eof"
#define MACRO_EVENT_ERROR_S		"net_err"
#define MACRO_EVENT_TIMEOUT_S	"net_timeout"
#define MACRO_EVENT_CONNECTED_S	"net_conn"

#define MACRO_CH_GPIO0_1	"GPIO0_1"
#define MACRO_CH_GPIO0_2	"GPIO0_2"
#define MACRO_CH_GPIO6_3	"GPIO6_3"
#define MACRO_CH_GPIO6_4	"GPIO6_4"

#define MACRO_VID_ENC_H264				"h264"
#define MACRO_AUD_FMT_PCM				"pcm"
#define MACRO_AUD_FMT_G711				"g711"
#define MACRO_AUD_FMT_G711_U			"g711"
#define MACRO_AUD_FMT_G711_A			"g711a"
#define MACRO_AUD_FMT_G722				"g722"
#define MACRO_AUD_FMT_AAC				"aac"
#define MACRO_AUD_FMT_MP3				"mp3"
#define MACRO_PROTO_TCP					"tcp"
#define MACRO_PROTO_UDP					"udp"
#define MACRO_PROTO_UDP_MULTICAST		"udp_multicast"
#define MACRO_PROTO_RTP					"rtp"
#define MACRO_PROTO_RTP_MULTICAST		"rtp_multicast"
#define MACRO_TSK_TP

#define MACRO_SUB_TYPE_PRE_VIEW         "preview"
#define MACRO_SUB_TYPE_ALARM_VD         "alarmvideo"
#define MACRO_SUB_TYPE_MP3_BROAD        "mp3broad"
#define MACRO_SUB_TYPE_VOICE_BROAD		"voicebroad"
#define MACRO_SUB_TYPE_VOICE            "voice"
#define MACRO_SUB_TYPE_ALARM_TK         "alarmtalk"

#define MACRO_NAME_MAX_LENGTH           30
#include <string>
#include <list>
#include <tuple>
#include <map>

typedef std::tuple<int, std::string>    TUPLE_TERM;
typedef std::list<TUPLE_TERM>           LIST_TERM;
typedef std::list<std::string>          LIST_USR;
typedef std::list<std::string>          LIST_INFO;
typedef std::list<int>                  LIST_SESS;
typedef std::map<int, int>				MAP_II;
typedef std::map<int, void*>            MAP_IV;
typedef std::map<unsigned int, void*>   MAP_UIV;
typedef std::map<std::string, void*>    MAP_SV;
typedef std::tuple<int, std::string>    TUPLE_TERM;
typedef std::map<std::string, std::string> MAP_SS;
typedef std::map<int, std::string>      MAP_STO;

//manual
//sto input sound card  info
typedef std::tuple<std::string, std::string, std::string, int, int>   TUPLE_SOUND_CARD; 
typedef std::list<TUPLE_SOUND_CARD>               LIST_SOUND_CARD;
typedef std::map<std::string, LIST_SOUND_CARD>   MAP_SOUND_CARD;

//sto palylist info 
typedef std::map<std::string, std::string>      MAP_FILE;
typedef std::map<std::string, MAP_FILE>         MAP_PLAY_LIST;
typedef std::map<int, std::string>              MAP_WARNING;
typedef std::map<std::string, int>              MAP_USR_VOLUME;
typedef std::map<int, std::string>				MAP_CALL_USR;
typedef std::map<std::string, MAP_CALL_USR>		MAP_CALL_USR_BY_GUID;

//typedef enum
//{
//	USR_STATE_OFFLINE = 0,		// 离线
//	USR_STATE_ONLINE = 1,	    // 在线
//	USR_STATE_SHOUTING = 2,     // 喊话中
//	USR_STATE_TALKING = 3,		// 对讲中
//	USR_STATE_CALLEE = 4,		
//	USR_STATE_CALLER = 5,		
//	USR_STATE_BROAD = 6			//广播中
//}enUsrState;

typedef enum
{
	USR_STATE_OFFLINE = 0,		// 离线
	USR_STATE_ONLINE = 1,	    // 在线
	USR_STATE_SHOUTING = 2,     // 喊话中
	USR_STATE_TALKING = 3,		// 对讲中
	USR_STATE_CALLIN = 4,		// 呼叫中
	USR_STATE_CALLEDIN = 5,		// 被呼叫中
	USR_STATE_BROAD = 6	,		// 广播中
	USR_STATE_MONITOR = 7		// 监控中
}enUsrState;

typedef enum
{
	TSK_STATE_FREE = 0,
	TSK_STATE_AUTO = 1,
	TSK_STATE_MANUAL = 2,  
	TSK_STATE_CALL_START = 3,
	TSK_STATE_CALL_PICKUP = 4,
	TSK_STATE_CALL_NO_USR = 5,

}TSK_STATE;

typedef enum station_login_result{
    STA_LOGIN_RET_SUCCESS,
    STA_LOGIN_RET_CONN_FAILED,
    STA_LOGIN_RET_DIS_CONNECTED,
    STA_LOGIN_USR_DISCONN,
    STA_LOGIN_RET_CONN_NO_REPLY,
    STA_LOGIN_RET_INVALID_PROTOCOL,
    STA_LOGIN_RET_INVALID_USR,
    STA_LOGIN_RET_INVALID_PWD,
}STA_LOGIN_RET;

typedef enum station_call_result{
    STA_CALL_RET_FAILED = 0,
    STA_CALL_RET_SHIFT,
    STA_CALL_RET_SUCCESS,
    STA_CALL_RET_CANCEL,
    STA_CALL_RET_CALLEE,
}STA_CALL_RET;

typedef enum
{
	    PLAY_MODE_SEQUE_PLAY	= 0, // PLAY_MODE_SEQUE_PLAY    = 4, PLAY_MODE_SINGLE_PLAY = 0 //顺序播放 单曲播放
		PLAY_MODE_RANDOM_PLAY	= 1, //随机播放
        PLAY_MODE_SINGLE_CYCLE	= 2, //单曲循环
        PLAY_MODE_LIST_PLAY		= 3, //列表播放
        PLAY_MODE_SINGLE_PLAY   = 4,  //单曲播放
}PLAY_MODE;



#define XML_ATTR_RET_OK                 "ok"
#define XML_ATTR_RET_ALREADY            "already"
#define XML_ATTR_RET_ALREADY_EXIST      "already_exist"
#define XML_ATTR_RET_NO_SPACE			"no_space"
#define XML_ATTR_RET_NO_FILE			"no_file"
#define XML_ATTR_RET_NO_DIR				"no_dir"
#define XML_ATTR_RET_EMPTY_FILE			"empty_file"
#define XML_ATTR_RET_EMPTY_DIR			"empty_dir"
#define XML_ATTR_RET_FAIL               "failed"
#define XML_ATTR_RET_INVALID_USR        "invalid_usr"
#define XML_ATTR_RET_INVALID_PWD        "invalid_pwd"

#define XML_ROOT                        "root"
#define XML_ITEM                        "item"
#define XML_SUBS						"subs"
#define XML_SUB                         "sub"
#define XML_AREAS                       "areas"
#define XML_TERMS                       "terms"
#define XML_PROGS                       "progs"
#define XML_TASKS                       "tasks"
#define XML_USRS                        "usrs"
#define XML_FILES                       "files"
#define XML_FILE                        "file"
#define XML_DEVS                        "devs"
#define XML_DEV                         "dev"
#define XML_LOGS                        "logs"
#define XML_AREA                        "area"
#define XML_TERM                        "term"
#define XML_PROG                        "prog"
#define XML_TASK                        "task"
#define XML_TSKS						"tsks"
#define XML_TSK							"tsk"
#define XML_TSK_LINK					"tsk_link"
#define XML_USR                         "usr"
#define XML_LOG                         "log"
#define XML_ROW							"row"
#define XML_SRCS						"srcs"
#define XML_SRC							"src"
#define XML_DSTS						"dsts"
#define XML_DST							"dst"
#define XML_SESSS						"sesss"
#define XML_SVR							"svr"
#define XML_ATTR_ORDER                  "ord"
#define XML_ATTR_SESS                   "sess"
#define XML_ATTR_RET                    "ret"
#define XML_ATTR_USR                    "usr"
#define XML_ATTR_PWD                    "pwd"
#define XML_ATTR_OS                     "os"
#define XML_ATTR_TYPE                   "tp"
#define XML_ATTR_SUB_TYPE               "stp"
#define XML_ATTR_AUD_ENC                 "aud_enc"
#define XML_ATTR_VID_ENC                 "vid_enc"
#define XML_ATTR_NAME                   "name"
#define XML_ATTR_NICK_NAME              "nick_name"
#define XML_ATTR_NEW_NAME               "new_name"
#define XML_ATTR_SRC_TYPE				"src_tp"
#define XML_ATTR_ID                     "id"
#define XML_ATTR_PID                    "pid"
#define XML_ATTR_PATH                   "path"
#define XML_ATTR_FTM                    "ftm"
#define XML_ATTR_START_SESS             "start_sess"
#define XML_ATTR_TALK_SESS              "talk_sess"
#define XML_ATTR_STATUS                 "status"
#define XML_ATTR_TXT                    "txt"
#define XML_ATTR_RAND                   "rand"
#define XML_ATTR_FORCE                  "force"
#define XML_ATTR_POWER                  "power"
#define XML_ATTR_SDATE                  "sdate"
#define XML_ATTR_STIME                  "stm"
#define XML_ATTR_ETIME                  "etm"
#define XML_ATTR_END                    "end"
#define XML_ATTR_EDATE                  "edate"
#define XML_ATTR_LOGTYPE                "log_type"
#define XML_ATTR_DURATION               "dura"
#define XML_ATTR_WEEK                   "week"
#define XML_ATTR_TIME_LONG              "tm_long"
#define XML_ATTR_RESULT                 "res"
#define XML_ATTR_TOTAL_LONG             "tt_long"
#define XML_ATTR_DB_ALM					"db_alm"
#define XML_ATTR_DB_LVL					"db_lvl"
#define XML_ATTR_DB_FILE				"db_file"
#define XML_ATTR_PLAY_VOL				"ply_vol"
#define XML_ATTR_FREEZE					"freeze"
#define XML_ATTR_CALLER_NAME            "caller_name"
#define XML_ATTR_CALLEE_NAME            "callee_name"
#define XML_ATTR_CALLER_TYPE            "caller_type"
#define XML_ATTR_CALLER_GUID			"caller_guid"
#define XML_ATTR_CALLEE_TYPE            "callee_type"
#define XML_ATTR_CALLEE_GUID			"callee_guid"
#define XML_ATTR_CALLEE_SESS			"callee_sess"
#define XML_ATTR_SIZE                   "size"
#define XML_ATTR_TIME                   "time"
#define XML_ATTR_SVR_TM					"svr_tm"
#define XML_ATTR_FULL_PATH1             "full_path1"
#define XML_ATTR_FULL_PATH2             "full_path2"
#define XML_ATTR_IO_OUT					"io_out"
#define XML_ATTR_IO_OUT1                "io_out1"
#define XML_ATTR_IO_OUT2                "io_out2"
#define XML_ATTR_IO_IN2					"io_in2"
#define XML_ATTR_IO_IDX                 "io_idx"
#define XML_ATTR_HAS_VID                "has_vid"
#define XML_ATTR_HAS_TALK               "has_talk"
#define XML_ATTR_IP                     "ip"
#define XML_ATTR_PORT                   "port"
#define XML_ATTR_SDP                    "sdp"
#define XML_ATTR_RTSP                   "rtsp"
#define XML_ATTR_DATE                   "date"
#define XML_ATTR_TERM_CNT               "term_cnt"
#define XML_ATTR_USR_CNT                "usr_cnt"
#define XML_ATTR_REAL_TERM_CNT          "real_term_cnt"
#define XML_ATTR_REAL_USR_CNT           "real_usr_cnt"
#define XML_ATTR_TITLE                  "title"
#define XML_ATTR_VAL					"val"
#define XML_ATTR_GRP_NAME				"grp_name"
#define XML_ATTR_PROJ_GUID				"proj_guid"
#define XML_ATTR_PROJ_NAME				"proj_name"
#define XML_ATTR_PROJ_START_TM			"proj_start_time"
#define XML_ATTR_PROJ_END_TM			"proj_end_time"
#define XML_ATTR_PROJ_PRIORITY			"proj_priority"
#define XML_ATTR_USR_GUID				"usr_guid"
#define XML_ATTR_USR_GUID1				"usr_guid1"
#define XML_ATTR_USR_GUID2				"usr_guid2"
#define XML_ATTR_USR_NAME				"usr_name"
#define XML_ATTR_USR_IP                 "usr_ip"

#define XML_ATTR_USR_PWD				"usr_pwd"
#define XML_ATTR_USR_STATUS				"usr_status"
#define XML_ATTR_USR_TSK				"usr_tsk"
#define XML_ATTR_USR_NICKNAME			"usr_nickname"
#define XML_ATTR_USR_TYPE				"usr_type"
#define XML_ATTR_USR_PLAY_FILE  		"usr_play_file"
#define XML_ATTR_USR_TALKING            "usr_talking"
#define XML_ATTR_USR_INFO               "usr_info"
#define XML_ATTR_USR_PRIORITY			"usr_priority"
#define XML_ATTR_GRP_GUID				"grp_guid"
#define XML_ATTR_USR_AO_VOL				"usr_ao_vol"
#define XML_ATTR_USR_TK_VOL				"usr_tk_vol"
#define XML_ATTR_USR_MT_VOL				"usr_mt_vol"
#define XML_ATTR_USR_HAS_VID			"usr_has_vid"
#define XML_ATTR_USR_RTSP_IP			"usr_rtsp_ip"
#define XML_ATTR_USR_RTSP_PORT			"usr_rtsp_port"
#define XML_ATTR_USR_RTSP_USR			"usr_rtsp_usr"
#define XML_ATTR_USR_RTSP_PWD			"usr_rtsp_pwd"
#define XML_ATTR_USR_RTSP_SDP			"usr_rtsp_sdp"
#define XML_ATTR_RTSP_URL				"rtsp_url"
#define XML_ATTR_USR_CALL_USR_GUID1		"usr_call_usr_guid1"
#define XML_ATTR_USR_CALL_USR_GUID2		"usr_call_usr_guid2"
#define XML_ATTR_USR_CALL_USR_GUID3		"usr_call_usr_guid3"
#define XML_ATTR_USR_CALL_USR_GUID4		"usr_call_usr_guid4"
#define XML_ATTR_USR_CALL_USR_GUID5		"usr_call_usr_guid5"

#define XML_ATTR_TSK_GUID				"tsk_guid"
#define XML_ATTR_TSK_NAME				"tsk_name"
#define XML_ATTR_TSK_PROJ_GUID			"tsk_proj_guid"
#define XML_ATTR_TSK_USR_GUID			"tsk_usr_guid"
#define XML_ATTR_TSK_SRC_TYPE			"tsk_src_type"
#define XML_ATTR_TSK_FILE_SEC			"tsk_file_sec"
#define XML_ATTR_TSK_JOB_TYPE			"tsk_job_type"
#define XML_ATTR_TSK_START_TIME			"tsk_start_time"
#define XML_ATTR_TSK_HAS_DURA			"tsk_has_dura"
#define XML_ATTR_TSK_DURATION			"tsk_duration"
#define XML_ATTR_TSK_HAS_STOP_TIME		"tsk_has_stop_time"
#define XML_ATTR_TSK_STOP_TIME			"tsk_stop_time"
#define XML_ATTR_TSK_WEEK_DATA			"tsk_week_data"
#define XML_ATTR_TSK_FLAG_RANDOM		"tsk_flag_random"
#define XML_ATTR_TSK_FLAG_DB_ALM		"tsk_flag_db_alm"
#define XML_ATTR_TSK_DB_ALM_LEVEL		"tsk_db_alm_level"
#define XML_ATTR_TSK_ALM_AUD_FILE		"tsk_alm_aud_file"
#define XML_ATTR_TSK_ONLY_POWER			"tsk_only_power"
#define XML_ATTR_TSK_FORCE_SIGNAL		"tsk_force_signal"
#define XML_ATTR_TSK_PLAY_VOLUMN		"tsk_play_volumn"
#define XML_ATTR_TSK_FREEZE				"tsk_freeze"
#define XML_ATTR_TSK_OFFLINE			"tsk_offline"
#define XML_ATTR_TSK_TRANS_PROTO		"tsk_trans_proto"
#define XML_ATTR_TSK_CH_NUM				"tsk_ch_num"
#define XML_ATTR_TSK_STRM_MODE			"tsk_strm_mode"
#define XML_ATTR_TSK_MD_TP				"tsk_md_tp"
#define XML_ATTR_TSK_LINK_GUID			"tsk_link_guid"
#define XML_ATTR_TSK_LINK_NMAE			"tsk_link_name"
#define XML_ATTR_TSK_LINK_TMP_GUID		"tsk_link_tmp_guid"
#define XML_ATTR_TSK_LINK_TMP_NMAE		"tsk_link_tmp_name"
#define XML_ATTR_TSK_LINK_USR			"tsk_link_usr"
#define XML_ATTR_USR_SIG_TYPE			"usr_sig_type"
#define XML_ATTR_USR_CH_NUM				"usr_ch_num"
#define XML_ATTR_USR_SIG_VAL			"usr_sig_val"
#define XML_ATTR_TSK_SYN_TP				"sync_tp"
#define XML_ATTR_TSK_SYNC_CNT			"sync_cnt"
#define XML_ATTR_TSK_LINK_SRC_TYPE		"tsk_link_src_type"
#define XML_ATTR_TSK_LINK_JOB_TYPE		"tsk_link_job_type"
#define XML_ATTR_TSK_LINK_START_TIME	"tsk_link_start_time"
#define XML_ATTR_TSK_LINK_HAS_DURA		"tsk_link_has_dura"
#define XML_ATTR_TSK_LINK_DURATION		"tsk_link_duration"
#define XML_ATTR_TSK_LINK_HAS_STOP_TIME	"tsk_link_has_stop_time"
#define XML_ATTR_TSK_LINK_STOP_TIME		"tsk_link_stop_time"
#define XML_ATTR_TSK_LINK_WEEK_DATA		"tsk_link_week_data"
#define XML_ATTR_TSK_LINK_TSK_GUID		"tsk_link_tsk_guid"

//uti_bmap_term_node
//#define XML_ATTR_BMAP_GUID				"bmap_guid"
#define XML_ATTR_BMAP_TERM_GUID			"bmap_term_guid"
#define XML_ATTR_BMAP_TERM_NAME			"bmap_term_name"
#define XML_ATTR_BMAP_TERM_RTSP			"bmap_term_rtsp"
#define XML_ATTR_BMAP_TERM_TSK_TYPE		"bmap_term_tsk_type"
#define XML_ATTR_BMAP_TERM_STATUS		"bmap_term_status"
#define XML_ATTR_BMAP_TERM_LNG			"bmap_term_lng"
#define XML_ATTR_BMAP_TERM_LAT			"bmap_term_lat"
#define XML_UPDATE_USR_STATUS			"update_usr_status"

//manual tsk
#define XML_ATTR_MANUAL_PLAY_FILE_NAME  "play_list_name"
#define XML_ATTR_PLAY_FILE_NAME         "file_name"
#define XML_ATTR_PLAY_FILE_PATH         "file_path"

#define XML_ATTR_PLAY_CARD_NICK_NAME    "sound_card_nick_name"
#define XML_ATTR_PLAY_CARD_NAME         "sound_card_name"
#define XML_ATTR_PLAY_CARD_CTRL         "sound_card_ctrl"
#define XML_ATTR_PLAY_CARD_VOLUME       "sound_card_volume"
#define XML_ATTR_PLAY_CARD_QUALITY      "sound_card_quality"
#define MACRO_TASK_CFG				    "task_cfg.xml"
#define XML_CARDS                       "cards"
#define XML_CARD                        "card"
#define XML_INPUTS                      "inputs"
#define XML_INPUT                       "input"
#define XML_TXTS                        "txts"
#define XML_TXT                         "txt"
#define XML_LISTS                       "lists"
#define XML_LIST                        "list"
#define XML_ATTR_POST_EVT_SVR			"post_evt_svr"
#define XML_ATTR_POST_EVT_USR			"post_evt_usr"

#define MANUAL_TSK_PLAY_VOLUME			101
//manual tsk

//term show tsk info 
#define XML_UPDATE_TERM_TSK_INFO       "update_term_tsk_info"

#define XML_TIMIMG_TSK					"timing_tsk"
#define XML_MANUAL_TSK                  "manual_tsk"
#define XML_TK_TSK						"tk_tsk"
#define XML_MONITOR_TSK					"monitor_tsk"
#define XML_FREE_TSK					"free_tsk"
#define XML_VOICE_BROAD_TSK				"voice_broad_tsk"

#define XML_SHOW_TSK_TYPE               "show_tsk_type"
#define XML_SHOW_AUD_SRC                "show_aud_src"
#define XML_SHOW_TK_INFO                "show_tk_info"
#define XML_SHOW_INFO                   "show_info"
//

#define XML_ATTR_SRC_INFO				"src_info"
#define XML_ATTR_AUTH_INFO				"auth_info"
#define XML_ATTR_VOL					"vol"
#define XML_ATTR_CAM					"cam"
#define XML_ATTR_STRM					"strm"
#define XML_ATTR_CH						"ch"
#define XML_ATTR_RATE					"rate"
#define XML_ATTR_VALUE					"value"
#define XML_ATTR_RET_STATUS				"ret_status"
#define XML_ATTR_SERIALPORT				"serialport"
#define XML_ATTR_FMT					"fmt"
#define XML_ATTR_PROTO					"proto"
#define XML_ATTR_MULTICAST_IP			"multicast_ip"
#define XML_ATTR_SAMPLE_PER_FRM			"spf"
#define XML_ATTR_CN						"cn"
#define XML_ATTR_EN						"en"
#define XML_ATTR_BIG5					"big5"
#define XML_ATTR_WND					"wnd"
#define XML_ATTR_SEND_ID				"send_id"
#define XML_ATTR_FETCH_ID				"fetch_id"
#define XML_ATTR_NAME					"name"
#define XML_ATTR_PATH					"path"
#define XML_ATTR_SIZE					"size"
#define XML_ATTR_OVER_WRITE				"over_write"
#define XML_ATTR_DEV                    "dev"
#define XML_ATTR_TRANS                  "trans"
#define XML_ATTR_FILE                   "file"
#define XML_ATTR_TSK_ID                 "tsk_id"
#define XML_ATTR_CLIENT					"client"
#define XML_ATTR_INT_SEC				"int_sec"
#define XML_ATTR_LOCK					"lock"
#define XML_ATTR_MSG					"msg"
#define XML_ATTR_VER					"ver"
#define XML_ATTR_DAT_SVR_MAIN			"dat_svr_main"
#define XML_ATTR_DAT_SVR_BAK			"dat_svr_bak"
#define XML_ATTR_SYNC_FLAG				"sync_flag"
#define XML_ATTR_SYNC_CNT				"sync_cnt"
#define XML_ATTR_UPGRADE_PATH			"upgrade_path"
#define XML_ATTR_UPGRADE_FORCE			"upgrade_force"
#define XML_ATTR_DEV_IP					"dev_ip"
#define XML_ATTR_DEV_CH					"dev_ch"
#define XML_ATTR_STA_CH					"sta_ch"

#define XML_ORDER_ADD_GRP				"add_grp"
#define XML_ORDER_UPDATE_USR_GRP		"update_usr_grp"
#define XML_ORDER_UPDATE_GRP			"update_grp"
#define XML_ORDER_DEL_GRP				"del_grp"
#define XML_ORDER_DEL_USR				"del_usr"
#define XML_ORDER_ADD_USR				"add_usr"
#define	XML_ORDER_ADD_TSK_LINK			"add_tsk_link"
#define XML_ORDER_UPDATE_TSK_LINK		"update_tsk_link"
#define XML_ORDER_DEL_TSK_LINK			"del_tsk_link"
#define XML_ORDER_UPDATE_USR			"update_usr"
#define XML_ORDER_USR_LOGIN             "usr_login"
#define XML_ORDER_USR_LOGIN_RET         "usr_login_ret"
#define XML_ORDER_DEV_LOGIN				"dev_login"
#define XML_ORDER_DEV_LOGIN_RET			"dev_login_ret"
#define XML_ORDER_USR_STATUS            "usr_status"
#define XML_ORDER_DEV_ONLINE            "dev_online"
#define XML_ORDER_DEV_OFFLINE           "dev_offline"
#define XML_ORDER_USR_ONLINE            "usr_online"
#define XML_ORDER_USR_OFFLINE           "usr_offline"
#define XML_ORDER_USR_NO_RESPONSE       "usr_no_response"
#define XML_ORDER_USR_DISONN            "usr_disconn"
#define XML_ORDER_PERIOD_RESP           "period_resp"
#define XML_ORDER_PERIOD_RESP_RET       "period_resp_ret"
#define XML_ORDER_PROG_REFRESH          "prog_refresh"
#define XML_ORDER_PROG_REFRESH_RET      "prog_refresh_ret"
#define XML_ORDER_TERM_USR_REFRESH      "term_usr_refresh"
#define XML_ORDER_PROG_MKDIR            "prog_mkdir"
#define XML_ORDER_PROG_MKDIR_RET        "prog_mkdir_ret"
#define XML_ORDER_PROG_RENAME           "prog_rename"
#define XML_ORDER_PROG_RENAME_RET       "prog_rename_ret"
#define XML_ORDER_PROG_DEL              "prog_del"
#define XML_ORDER_PROG_DEL_RET          "prog_del_ret"
#define XML_ORDER_TASK_DEL              "task_del"
#define XML_ORDER_TASK_DEL_RET          "task_del_ret"
#define XML_ORDER_TASK_REFRESH          "task_refresh"
#define XML_ORDER_TASK_ADD              "task_add"
#define XML_ORDER_TASK_ADD_RET          "task_add_ret"
#define XML_ORDER_TASK_EDIT             "task_edit"
#define XML_ORDER_TASK_EDIT_RET         "task_edit_ret"
#define XML_ORDER_TASK_START            "task_start"
#define XML_ORDER_TASK_START_RET        "task_start_ret"
#define XML_ORDER_TASK_STOP             "task_stop"
#define XML_ORDER_TASK_STOP_RET         "task_stop_ret"
#define XML_ORDER_TASK_STATUS           "task_status"
#define XML_ORDER_TALK_START            "talk_start"
#define XML_ORDER_TALK_STARTED          "talk_started"
#define XML_ORDER_TALK_STOP             "talk_stop"
#define XML_ORDER_TALK_OVER             "talk_over"
#define XML_ORDER_TALK_STATUS           "talk_status"
#define XML_ORDER_MSG                   "msg"
#define XML_ORDER_TERM_AV_START         "term_av_start"
#define XML_ORDER_TERM_AV_START_RET     "term_av_start_ret"
#define XML_ORDER_TERM_AV_STOP          "term_av_stop"
#define XML_ORDER_TERM_AV_TALK_START    "term_av_talk_start"
#define XML_ORDER_TERM_AV_TALK_START_RET "term_av_talk_start_ret"
#define XML_ORDER_TERM_AV_TALK_STOP     "term_av_talk_stop"
#define XML_ORDER_TERM_AV_TALK_STOP_RET "term_av_talk_stop_ret"
#define XML_ORDER_TERM_STATUS           "term_status"
#define XML_ORDER_TERM_ALARM_OUT1_STATUS "term_alarm_out1_status"
#define XML_ORDER_USR_AV_START          "usr_av_start"
#define XML_ORDER_USR_AV_STOP           "usr_av_stop"
#define XML_ORDER_USR_AV_TALK_START     "usr_av_talk_start"
#define XML_ORDER_USR_AV_TALK_STOP      "usr_av_talk_stop"
#define XML_ORDER_TIMEOUT               "timeout"
#define XML_ORDER_TERM_CALL_REQ         "term_call_req"
#define XML_ORDER_TERM_CALL_REQ_RET     "term_call_req_ret"
#define XML_ORDER_TERM_CALL_REQ_CANCEL  "term_call_req_cancel"
#define XML_ORDER_USR_CALL_REQ          "usr_call_req"
#define XML_ORDER_USR_CALL_REQ_RET      "usr_call_req_ret"
#define XML_ORDER_USR_CALL_REQ_RET_FB   "usr_call_req_ret_fb"
#define XML_ORDER_USR_CALL_REQ_CANCEL   "usr_call_req_cancel"
#define XML_ORDER_USR_CALL_REQ_TIMEOUT  "usr_call_req_timeout"
#define XML_ORDER_BC_SVR_OFFLINE        "bc_svr_offline"
#define XML_ORDER_USR_DEBUG_LOG         "usr_debug_log"
#define XML_ORDER_FILE_TRANSFER         "file_transfer"
#define XML_ORDER_FILE_TRANSFER_RET     "file_transfer_ret"
#define XML_ORDER_FILE_TRANSFER_CANCEL  "file_transfer_cancel"
#define XML_ORDER_QUERY_TALK_LOG        "query_talk_log"
#define XML_ORDER_QUERY_TALK_LOG_RET    "query_talk_log_ret"
#define XML_ORDER_QUERY_TALK_STO        "query_talk_sto"
#define XML_ORDER_QUERY_TALK_STO_RET    "query_talk_sto_ret"
#define XML_ORDER_REAL_CAP_START        "real_cap_start"
#define XML_ORDER_REAL_CAP_START_RET    "real_cap_start_ret"
#define XML_ORDER_REAL_CAP_STOP         "real_cap_stop"
#define XML_ORDER_REAL_CAP_STOP_RET     "real_cap_stop_ret"
#define XML_ORDER_TERM_IO_CTRL          "term_io_ctrl"
#define XML_ORDER_TERM_IO_CTRL_RET      "term_io_ctrl_ret"
#define XML_ORDER_TERM_IO_OUT_STATUS    "term_io_out_status"
#define XML_ORDER_GROUP_DEL_TERM        "group_del_term"
#define XML_ORDER_USR_REAL_CAP_START     "usr_real_cap_start"
#define XML_ORDER_USR_REAL_CAP_START_RET "usr_real_cap_start_ret"
#define XML_ORDER_USR_REAL_CAP_STOP      "usr_real_cap_stop"
#define XML_ORDER_USR_REAL_CAP_STOP_RET  "usr_real_cap_stop_ret"

#define XML_ORDER_SET_AI_VOL				"set_ai_vol"
#define XML_ORDER_SET_TALK_AI_VOL			"set_talk_ai_vol"
#define XML_ORDER_SET_AO_VOL				"set_ao_vol"
#define XML_ORDER_SET_AI_NORMAL				"set_ai_normal"
#define XML_ORDER_SET_AI_TALK				"set_ai_talk"
#define XML_ORDER_START_AV					"start_av"
#define XML_ORDER_START_AV_MP3				"start_av_mp3"
#define XML_ORDER_START_MONITOR				"start_monitor"
#define XML_ORDER_START_MONITOR_RET			"start_monitor_ret"
#define XML_ORDER_STOP_TASK					"stop_task"
#define XML_ORDER_CONFIRM_SESS				"confirm_sess"
#define XML_ORDER_CONFIRM_SESS_RET			"confirm_sess_ret"
#define XML_ORDER_CANCEL_CONFIRM			"cancel_confirm"
#define XML_ORDER_START_AV_RET				"start_av_ret"
#define XML_ORDER_START_AV_MP3_RET			"start_av_mp3_ret"
#define XML_ORDER_START_AV_USR_SESS			"start_av_usr_sess"
#define XML_ORDER_START_AV_USR_SESS_RET		"start_av_usr_sess_ret"
#define XML_ORDER_STOP_AV					"stop_av"
#define XML_ORDER_STOP_AV_MP3			    "stop_av_mp3"
#define XML_ORDER_REQ_IDR					"req_idr"
#define XML_ORDER_CTRL_IO_OUT				"ctrl_io_out"
#define XML_ORDER_CTRL_IO_OUT_RET			"ctrl_io_out_ret"
#define XML_ORDER_TRANSPARENT				"transparent"
#define XML_ORDER_SET_NAME					"set_name"
#define XML_ORDER_SET_CLIENT				"set_client"
#define XML_ORDER_SET_DATE					"set_date"
#define XML_ORDER_SET_DB_ALARM				"db_alarm"
#define XML_ORDER_SET_MODEL					"set_model"
#define XML_ORDER_UPDATE_AUD_PLAY			"update_aud_play"
#define XML_ORDER_STOP_AUD_PLAY				"stop_aud_play"
#define XML_ORDER_STOP_AUD_TRANS			"stop_aud_trans"
#define XML_ORDER_UPDATE_AUD_PLAY_RET		"update_aud_play_ret"
#define XML_ORDER_MONITOR_TSK_STATUS		"monitor_tsk_status"
#define XML_ORDER_BC_USR_TSK_STATUS			"bc_usr_tsk_status"
#define XML_ORDER_BC_FILE_TSK_STATUS		"bc_file_tsk_status"
#define XML_ORDER_TK_TSK_STATUS				"tk_tsk_status"
#define XML_ORDER_REQ_WRITE_FILE			"req_write_file"
#define XML_ORDER_REQ_WRITE_FILE_RET		"req_write_file_ret"
#define XML_ORDER_SEND_START				"send_start"
#define XML_ORDER_SEND_END					"send_end"
#define XML_ORDER_SEND_START_TRANS			"send_start_trans"
#define XML_ORDER_SEND_END_TRANS			"send_end_trans"
#define XML_ORDER_CHK_FILE_INFO				"chk_file_info"
#define XML_ORDER_CHK_FILE_INFO_RET			"chk_file_info_ret"
#define XML_ORDER_REQ_FETCH_FILE			"req_fetch_file"
#define XML_ORDER_REQ_FETCH_FILE_RET		"req_send_write_file_ret"
#define XML_ORDER_REQ_CAP_PIC				"req_cap_pic"
#define XML_ORDER_REQ_CAP_PIC_RET			"req_cap_pic_ret"
#define XML_ORDER_SIMULATE_EVENT			"simulate_event"
#define XML_ORDER_ALARM_EVENT				"alarm_event"
#define XML_ORDER_ALARM_EVENT_RET			"alarm_event_ret"
#define XML_ORDER_CHK_DIR_INFO				"chk_dir_info"
#define XML_ORDER_CHK_DIR_INFO_RET			"chk_dir_info_ret"
#define XML_ORDER_DEL_FILE_INFO				"del_file_info"
#define XML_ORDER_DEL_FILE_INFO_RET			"del_file_info_ret"
#define XML_ORDER_PLAY_MP3_FILE				"play_mp3_file"
#define XML_ORDER_PLAY_MP3_FILE_RET			"play_mp3_file_ret"
#define XML_ORDER_ADD_SVR_TSK				"add_svr_tsk"
#define XML_ORDER_UPDATE_SVR_TSK			"update_svr_tsk"
#define XML_ORDER_DEL_SVR_TSK				"del_svr_tsk"
#define XML_ORDER_EXEC_BC_TSK				"exec_bc_tsk"
#define XML_ORDER_STOP_BC_TSK				"stop_bc_tsk"
#define XML_ORDER_AUTO_BC_TSK				"auto_bc_tsk"
#define XML_ORDER_DEL_STA_TSK				"del_sta_tsk"
#define XML_ORDER_ADD_STA_TSK				"add_sta_tsk"
#define XML_ORDER_UPDATE_STA_TSK			"update_sta_tsk"

#define XML_ORDER_EXEC_BC_MANUAL_TSK        "exec_manual_tsk"
#define XML_ORDER_EXEC_BC_MANUAL_TSK_RET    "exec_manual_tsk_ret"
#define XML_ORDER_STOP_BC_MANUAL_TSK        "stop_manual_tsk"
#define XML_ORDER_STOP_BC_MANUAL_TSK_RET    "stop_manual_tsk_ret"

#define XML_ORDER_ADD_SVR_PROJ				"add_svr_proj" 
#define XML_ORDER_UPDATE_SVR_PROJ           "update_svr_proj"
#define XML_ORDER_DEL_SVR_PROJ				"del_svr_proj"
#define XML_ORDER_EXEC_SVR_PROJ             "exec_svr_proj"
#define XML_ORDER_NO_DEAMOM_SVR_PROJ	    "no_deamom_svr_proj"

#define XML_ORDER_ADD_TMP_LINK_TSK			"add_tmp_link_tsk"
#define XML_ORDER_UPDATE_TMP_LINK_TSK		"update_tmp_link_tsk"
#define XML_ORDER_DEL_TMP_LINK_TSK			"del_tmp_link_tsk"
#define XML_ORDER_BIND_TERM_TSK_LINK		"bind_term_tsk_link"
#define XML_ORDER_DEL_TERM_TSK_LINK         "del_term_tsk_link"
#define XML_ORDER_SYNC_FRM					"sync_frm"
#define XML_ORDER_SYNC_FRM_RET				"sync_frm_ret"
#define XML_ORDER_DAT_CLI_STATUS			"dat_cli_status"
#define XML_ORDER_DAT_SVR_MSG				"dat_svr_msg"
#define XML_ORDER_DAT_SVR_CTRL				"dat_cli_ctrl"
#define XML_ORDER_DAT_CLI_CHK_VER			"dat_cli_chk_ver"
#define XML_ORDER_DAT_CLI_CHK_VER_RET		"dat_cli_chk_ver_ret"
#define XML_ORDER_DAT_CLI_UPGRADE			"dat_cli_upgrade"
#define XML_ORDER_DAT_CLI_UPGRADE_RET		"dat_cli_upgrade_ret"
#define XML_ORDER_DAT_SVR_UPDATE_ADDR		"dat_svr_update_addr"
#define XML_ORDER_DAT_CLI_AUTH				"dat_cli_auth"
#define XML_ORDER_DAT_CLI_AUTH_RET			"dat_cli_auth_ret"
#define XML_ORDER_DAT_SVR_CONN				"dat_svr_conn"
#define XML_ORDER_DAT_SVR_DISCONN			"dat_svr_disconn"
#define XML_ORDER_AUTH_LOGIN				"auth_login"
#define XML_ORDER_REQ_UPGRADE				"req_upgrade"
#define XML_ORDER_PRE_START_AV				"pre_start_av"
#define XML_ORDER_NET_SPEED					"net_speed"
#define XML_ORDER_NET_SPEED_AUD				"net_speed_aud"
#define XML_ORDER_NET_SPEED_VID				"net_speed_vid"

#define MACRO_SVR_LISTEN_PORT				1220

#define MACRO_CTS_ALARM_OUT1_STATE			911
#define MACRO_CTS_ALARM_OUT2_STATE			912

#define MACRO_TASK_TYPE_BROADCAST			1
#define MACRO_TASK_TYPE_SHOUT				2
#define MACRO_TASK_TYPE_TALK				3
#define MACRO_SUPER_DOG_MAX_TERM			5
#define MACRO_SUPER_DOG_MAX_USR				5

#define MACRO_TYPE_FILE						1
#define MACRO_TYPE_CARD						2
#define MACRO_TYPE_USR						3

//tp[1(file task) 2(sound card) 3(user)]

#define MACRO_FMT_XML						"xml"
#define MACRO_FMT_JASON						"json"


#define MACRO_AREA_UNKOWN_AREA_ID			10000
#define MACRO_DEFAULT_GRP_GUID			"{00000000-0000-0000-0000-000000000000}"
#define MACRO_USER_GRP_GUID             "{-1}"
#define MACRO_ALL_GRP_GUID              "all_grp"

//server
#define MACRO_FILE_MAIN_WINDOW          "main_window"
#define MACRO_FILE_PWD_AUTH_WIDGET      "passwd_auth_widget"
#define MACRO_FILE_PROG_DIR_OPER_WIDGET "prog_dir_oper_widget"
#define MACRO_FILE_PROG_MANAGER_WIDGET  "prog_manager_widget"
#define MACRO_FILE_PROG_TABLE_WIDGET    "prog_table_widget"
#define MACRO_FILE_REAL_CAP_WIDGET      "real_cap_widget"
#define MACRO_FILE_FOR_FULL_WIDGET      "for_full_widget"
#define MACRO_FILE_RUN_STATUS_WIDGET    "run_status_widget"
#define MACRO_FILE_TERM_TALK_WIDGET     "term_talk_widget"
#define MACRO_FILE_VOL_CTRL_WIDGET      "vol_ctrl_widget"
#define MACRO_FILE_ADV_SETTING_WIDGET   "advanced_setting_wideget"
#define MACRO_FILE_AREA_MANAGER_WIDGET  "area_manage_widget"
#define MACRO_FILE_BASIC_SETUP_WIDGET   "basic_setup_widget"
#define MACRO_FILE_MEMBER_SETUP_WIDGET  "member_setup_widget"
#define MACRO_FILE_TASK_PRIORITY_WIDGET "task_priority_widget"
#define MACRO_FILE_USR_MANAGER_WIDGET   "usr_manager_widget"
#define MACRO_FILE_TALK_LOG_WIDGET      "talk_log_widget"
#define MACRO_FILE_EDIT_FILE_WIDGET     "edit_file_widget"
#define MACRO_FILE_JOS_TASK_WIDGET      "jos_task_widget"
#define MACRO_FILE_NEW_SCHEME_WIDGET    "new_scheme_widget"
#define MACRO_FILE_RENAME_SCHEME_WIDGET "rename_scheme_widget"
#define MACRO_FILE_TASK_TIME_WIDGET     "task_time_setup_widget"
#define MACRO_FILE_TIMING_JOW_WIDGET    "timing_jow_widget"
#define MACRO_FILE_PROG_TASK_WIDGET     "prog_task_detail_widget"
#define MACRO_FILE_TIMING_PROG_WIDGET   "timing_prog_widget"
#define MACRO_FILE_UTI_AREA_MANAGER     "uti_area_manager"
#define MACRO_FILE_COMMON_INTERFACE     "common_interface"
#define MACRO_FILE_GRID_ITEM_WIDGET     "ui_grid_widget_item"

//station
#define MACRO_FILE_CALL_USR_WIDGET     "call_usr_widget"
#define MACRO_FILE_CALLER_USR_DIALOG   "caller_or_callee_dialog"
#define MACRO_FILE_USR_TALK_WIDGET     "usr_talk_widget"
#define MACRO_FILE_LOGIN_DIALOG        "login_widget"
#define MACRO_FILE_IP_LINE_WIDGET      "qiplineedit"

#define DATABASE_MANAGERR_PASSWORD      "avsz_2016"
#define MACRO_USR_SVR                   "usr_svr"
#define MACRO_TMP_SVR_GUID              "tmp_svr_guid"
#define MACRO_USR_SVR_NAME              "服务器"

#define MACRO_DB_OPEN_OK				"db_open_ok"
#define MACRO_DB_OPEN_ERR				"db_open_err"
#define MACRO_DB_CLOSE					"db_close"
#define MACRO_NET_SVC					"net_svc"
#define MACRO_INIT_OK					"init_ok"
#define MACRO_INIT_END					"init_end"
#define MACRO_INIT_ERR					"init_err"
#define MACRO_SVC_START					"svc_start"
#define MACRO_SVC_STOP					"svc_stop"
#define MACRO_INVALID_ARG				"invalid_arg"
#define MACRO_ALREADY_RUN				"already_run"
#define MACRO_ACCEPTOR_ERR				"acceptor_err"
#define MACRO_SVC_THR_ERR				"svc_thr_err"
#define MACRO_NEXT_SESS_ERR				"next_sess_err"
#define MACRO_IMPORT_LANG_DATA			"import_lang_data"
#define MACRO_TSK_STATUS				"tsk_status"
#define MACRO_ALREADY_PICKUP			"already_pickup"
#define MACRO_CALLING_OTHER				"calling_other"
#define MACRO_INVALID_DOG				"invalid_dog"
#define MACRO_TCP_CONN					"tcp_conn"
#define MACRO_TCP_CLOSED				"tcp_closed"
#define MACRO_REQ_AUTH_CNT				"req_auth_cnt"
#define MACRO_INVALID_AUTH				"invalid_auth"
#define MACRO_TERM_CONN					"term_conn"
#define MACRO_TERM_DISCONN				"term_disconn"

#define MACRO_RESULT_KEY				"result"
#define MACRO_RESULT_TRUE				"true"
#define MACRO_RESULT_FALSE				"false"
#define MACRO_LOAD_STATION_USR			"load_station_usr"
#define MACRO_LOAD_ANDROID_USR			"load_android_usr"
#define MACRO_LOAD_EMBED_USR			"load_embed_usr"
#define MACRO_LOAD_ALL_USR				"load_all_usr"
#define MACRO_LOAD_USR_BY_GRP			"load_usr_by_grp"
#define MACRO_LOAD_ALL_USR_INIT			"load_all_usr_init"
#define MACRO_LOAD_ALL_GRP				"load_all_grp"
#define MACRO_LOAD_ALL_GRP_INIT			"load_all_grp_init"
#define MACRO_LOAD_ALL_PROJ				"load_all_proj"
#define MACRO_LOAD_ALL_PROJ_INIT		"load_all_proj_init"
#define MACRO_LOAD_ALL_PROG_INIT		"load_all_prog_init"
#define MACRO_LOAD_LANG					"load_lang"
#define MACRO_LOAD_LANG_INIT			"load_lang_init"
#define MACRO_LOAD_USR_REL				"load_usr_rel"
#define MACRO_LOAD_USR_REL_INIT			"load_usr_rel_init"
#define MACRO_LOAD_NET_SVR				"load_net_svr"
#define MACRO_LOAD_NET_SVR_INIT			"load_net_svr_init"
#define MACRO_LOAD_SVR_TSK				"load_svr_tsk"
#define MACRO_LOAD_ALL_BC_TSK_BY_DB		"load_all_bc_tsk_by_db"
#define MACRO_LOAD_SVR_TSK_INIT			"load_svr_tsk_init"
#define MACRO_LOAD_STA_TSK				"load_sta_tsk"
#define MACRO_LOAD_STA_TSK_INIT			"load_sta_tsk_init"
#define MACRO_LOAD_TSK_LINK_INIT		"load_tsk_link_init"
#define MACRO_LOAD_TSK_LINK_SHOW		"load_all_show_tsk_link"
#define MACRO_LOAD_TSK_LINK				"load_all_tsk_link"
#define MACRO_LOAD_BMAP_TERM_INIT		"load_bmap_term_init"
#define MACRO_ALARM_EVENT				"alarm_event"
#define MACRO_STATUS_IO_IN				"status_io_in"

#define MACRO_DB_TABLE_USR				"usr"
#define MACRO_DB_TABLE_USR_REL			"usr_rel"
#define MACRO_DB_TABLE_TSK				"tsk"
#define MACRO_DB_TABLE_TSK_SRC			"tsk_src"
#define MACRO_DB_TABLE_TSK_DST			"tsk_dst"
#define MACRO_DB_TABLE_TSK_LINK			"tsk_link"
#define MACRO_DB_TABLE_TSK_LINK_TMP		"tsk_link_tmp"
#define MACRO_DB_TABLE_BMAP_TERM		"bmap_term"
#define MACRO_DB_TABLE_KV				"kv"
#define MACRO_DB_TABLE_GRP				"grp"
#define MACRO_DB_TABLE_PROJ				"proj"
#define MACRO_TIG_NAME_KV_KEY			"kv_key"
#define MACRO_TIG_NAME_KV_VAL			"kv_value"
#define MACRO_TIG_NAME_KV_ID			"kv_id"
#define MACRO_TIG_NAME_KV_NOTE			"kv_note"
#define MACRO_TIG_NAME_LANG_CN			"lang_cn"
#define MACRO_TIG_NAME_LANG_EN			"lang_en"
#define MACRO_TIG_NAME_LANG_BIG5		"lang_big5"
#define MACRO_TIG_NAME_LANG_FILE		"lang_file"
#define MACRO_TIG_NAME_LANG_KEY			"lang_key"

#define MACRO_KV_SYS_VER				"sys_ver"
#define MACRO_KV_SYS_PWD				"sys_pwd"
#define MACRO_KV_SYS_GUID				"sys_guid"
#define MACRO_KV_SYS_PROJ_GUID			"sys_proj_guid"
#define MACRO_KV_SYS_ENCODE_TYPE		"sys_encode_type"
#define MACRO_KV_SYS_STO_PATH			"sys_sto_path"
#define MACRO_KV_SYS_PROG_PATH          "sys_prog_path"
#define MACRO_KV_SYS_LIB_PATH			"sys_lib_path"
#define MACRO_KV_SYS_DAEMON_PROJ		"sys_daemon_proj"
#define MACRO_KV_SYS_LOCAL_IP			"sys_local_ip"

#define MACRO_JOB_TYPE_ALL				-1
#define MACRO_JOB_TYPE_EVERY_DAY		0
#define MACRO_JOB_TYPE_ONECE			1
#define MACRO_JOB_TYPE_TEMP				2
#define MACRO_JOB_TYPE_EVERY_WEEK		7
#define MACRO_KV_KEY_STO_PATH			"sto_path"
#define MACRO_KV_KEY_LIB_PATH			"lib_path"
#define MACRO_KV_KEY_SYS_VER			"sys_ver"
#define MACRO_KV_KEY_SYS_PWD			"sys_pwd"
#define MACRO_KV_KEY_SYS_GUID			"sys_guid"
#define MACRO_KV_KEY_SYS_PROJ_GUID		"sys_proj_guid"
#define MACRO_KV_KEY_ENCODE_TYPE		"encode_type"

#define MACRO_TERM_MP3_FILE_DIR         "/mnt/app/mp3"
#define MACRO_TERM_APP_FILE_DIR         "/mnt/app/app"
#define MACRO_TERM_MMC0_FILE_DIR        "/mnt/mmc0/mp3"
#define MACRO_TERM_MMC1_FILE_DIR        "/mnt/mmc1/mp3"
#define MACRO_TERM_USB_FILE_DIR         "/mnt/usb/mp3"

#define MACRO_TRAN_FILE_DIR_NAME        "/tts_dir"

#define MACRO_TSK_OWNER_SVR_TSK			"svr_tsk"
#define MACRO_TSK_OWNER_STA_TSK			"sta_tsk"

#define MACRO_GET_DEV_STATUS                "get_dev_status"
#define MACRO_GET_DEV_ONLINE                "get_dev_online"
#define MACRO_START_TSK_BC_FILE_DEV_ALL     "start_tsk_bc_file_dev_all"
#define MACRO_START_TSK_BC_FILE_DEV_SPEC    "start_tsk_bc_file_dev_spec"
#define MACRO_STOP_TSK                      "stop_tsk"
#define MACRO_ON_DEV_STATUS                 "on_dev_status"

#define MACRO_GET_DEV_STATUS_RET                "get_dev_status_ret"
#define MACRO_GET_DEV_ONLINE_RET                "get_dev_online_ret"
#define MACRO_START_TSK_BC_FILE_DEV_ALL_RET     "start_tsk_bc_file_dev_all_ret"
#define MACRO_START_TSK_BC_FILE_DEV_SPEC_RET    "start_tsk_bc_file_dev_spec_ret"
#define MACRO_STOP_TSK_RET                      "stop_tsk_ret"

#define MACRO_TASK_SRC_TYPE_NULL          0
#define MACRO_TASK_SRC_TYPE_MP3           1
#define MACRO_TASK_SRC_TYPE_TXT           2
#define MACRO_TASK_SRC_TYPE_CARD          3
#define MACRO_TASK_SRC_TYPE_USR           4

#define MACRO_TASK_STATUS_STOP            0
#define MACRO_TASK_STATUS_AUTO            1
#define MACRO_TASK_STATUS_EXEC            2

#define MACRO_PROJ_STATUS_STOP            0
#define MACRO_PROJ_STATUS_EXEC            1


#define MACRO_KEY_VEDIO_SPEED "speed"
#endif

