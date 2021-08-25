#ifndef win_sta_core_h__
#define win_sta_core_h__

#include "../utility/uti_export.h"
#include "../common/macro_func_cb.h"
#include <string>
#include <list>

class UTI_EXPORT win_sta_core
{
private:
	static win_sta_core* inst_;
public:
	static win_sta_core* instance();

	win_sta_core();

	//预初始化，module_path表示当前程序所在路径，
	//func表示接口消息回调，usr_data是外部自定义参数，在func回调是传回给外部
	//aud_type,默认 0-表示音频采样44.1K，G711编码			1-表示音频采样32K，PCM格式
	void pre_init(const std::string& module_path, func_cb_cli func, unsigned long long usr_data = 0, int aud_type = 0);

	void pre_init_java(const std::string& module_path, int aud_type);
	void aft_fini();

	//初始化连接服务端，svr_ip表示服务端IP地址，svr_port表示服务端端口
	//usr_name表示客户端用户名，usr_pwd表示用户的密码
	//返回值 表示客户端网络连接ID（sess_client）从1开始计数，负数表示错误
	int init(const std::string& svr_ip, unsigned short svr_port,
		const std::string& usr_name, const std::string& usr_pwd);

	//关闭这个客户端的网络连接
	void fini(int sess_client);

	//关闭所有客户端网络连接
	void fini_all();

	//请求执行手动任务
	//sess_client表示网络连接ID
	//md_tp 媒体类型  1[视频] 2[音频]
	//trans_proto 传输协议 1[TCP]2[UDP]3[UDP组播]
	//list_dst_guid 表示任务的目标用户(用户GUID的形式)
	//flag_usr_data 值false表示由win_sta_core内部采集麦克风数据，true表示由用户调用send_aud_44100_1ch_g711接口，发送数据 
	int exec_manual_tsk_speak(int sess_client, int md_tp,
		int trans_proto, std::list<std::string>& list_dst_guid, int vol, bool flag_usr_data = false);

	int exec_manual_tsk_mp3_broad(int sess_client, int md_tp, int trans_proto, 
		std::list<std::string>& list_dst_guid, int vol, 
		std::list<std::string>& list_src_path, std::string& new_tsk_guid);

	int exec_manual_tsk_mp3(
		int sess_client, int md_tp, int trans_proto,
		std::list<std::string>& list_dst_guid,
		std::list<std::string>& list_src_mp3, 
		int vol, std::string& new_tsk_guid, int flag_random = 0);

	//停止任务，tsk_guid表示任务的GUID
	int stop_task(int sess_client, const std::string& tsk_guid);

	//请求监视监听其它用户的音视频，
	//callee_guid表示被请求的用户GUID
	//ch_idx 音视频通道号  0,1,2......
	//strm 表示码流类型，主要应用于视频码流, 0,1,2......
        //md_tp 媒体类型  1[视频] 2[音频] 3[音视频
	//proto 传输协议 1[TCP]2[UDP]3[UDP组播]
	//注意：目前UDP和UDP组播 只适用于音频， 如果包括视频，需选择TCP方式
	//发起任务后，会有回调消息： 回调参数type 的值是monitor_tsk_status
	//回调参数value 的值 是xml字符串，包含任务GUID，及任务状态status 等信息
	int start_monitor(int sess_client, const std::string& callee_guid, 
		int ch_idx, int strm, int md_tp, int proto, unsigned long long wnd = 0);

	//用户呼叫请求，用于发起呼叫	
	//callee_guid表示呼叫的用户GUID
	//发送usr_call_req请求服务端后，如果被呼叫用户不在线，会有回调消息：
	// 回调参数type 的值是 usr_call_req_ret（包含被呼用户状态信息，例如0表示不在线）

	//或回调参数type 的值是tk_tsk_status消息表示任务运行状态，
	//回调参数value 的值 是xml字符串，该消息中包含tsk_guid, 任务状态如下
	//    0[TSK_STATE_FREE], 1[TSK_STATE_AUTO],2[TSK_STATE_MANUAL]
	//    3[TSK_STATE_CALL_START],4[TSK_STATE_CALL_PICKUP] 或 5[TSK_STATE_CALL_NO_USR]
	int usr_call_req(int sess_client, const std::string& callee_guid, 
		int ch_idx, int strm, int md_tp, int proto, unsigned long long wnd = 0);

	//用户呼叫请求响应，用于被呼叫端 接收到呼叫消息(usr_call_req)时的回应
	//tsk_guid 表示任务GUID,由服务端发到被呼叫端消息usr_call_req中获取
	//status  0表示拒绝、1表示正忙、2表示接通
	//发送usr_call_req_ret到服务端后，服务端根据情况可能会返回usr_call_req_ret_fb消息，表示该呼叫已经被接听（already_pickup）或正在呼叫移动其它用户（calling_other）
	//新接口 wong-2018/10/10
	int usr_call_req_ret(int sess_client, int ch, int strm,
		const std::string& tsk_guid, const std::string &call_guid, int status, unsigned long long wnd = 0);

	int start_av_ret(int sess_client, const std::string& tsk_guid,
		const std::string& ret, int caller_sess, int sample, const std::string& aud_enc);

	int start_av_mp3_ret(int sess_client, const std::string& tsk_guid,
		const std::string& ret, int caller_sess, int sample, const std::string& aud_enc);

	int pre_start_aud(int sess_client, const std::string& callee_guid = "");

	//添加任务
	//<root ord ="task_add" >
	//	<tsk>
	//	<row>
	//	<tsk_guid val ="{ 151196D6 - 33D7 - 4426 - 9E3C - E374FE12AE51 }" / >
	//	<tsk_name val ="任务名[utf8编码]" / >
	//	<tsk_usr_guid val ="{ B7BEE137 - 901C - 46C2 - A7F8 - DBA588A04774 }" / >
	//	<tsk_src_type val =" 0[空源] 1[文件]2[指定声卡]3[指定IP终端或用户]" / >
	//	<tsk_file_sec val ="50[文件总时长秒数]" / >
	//	<tsk_job_type val ="0[每天任务]1[一次任务]2[手动任务]7[每周任务]" / >
	//	<tsk_start_time val ="2017 - 07 - 31 14:00 : 00[任务开始时间]" / >
	//	<tsk_has_dura val ="0 / 1[是否指定持续时间]" / >
	//	<tsk_duration val ="30[指定持续时间总秒数]" / >
	//	<tsk_has_stop_time val ="0 / 1[是否指定任务有效截止时间]" / >
	//	<tsk_stop_time val ="2017 - 08 - 10 16:00 : 00[任务有效截止时间]" / >
	//	<tsk_week_data val =" 0110011[从周日到周六, 1表示生效，0表示失效]" / >
	//	<tsk_flag_random val ="0 / 1[是否随机标识]" / >
	//	<tsk_flag_db_alarm val ="0 / 1[是否启用喧哗报警]" / >
	//	<tsk_db_alm_level val ="50[喧哗报警分贝值]" / >
	//	<tsk_alm_aud_file val ="报警声音文件全路径" / >
	//	<tsk_only_power val ="0 / 1[是否仅控制功放电源]" / >
	//	<tsk_force_signal val ="0 / 1[是否同步控制强切信号]" / >
	//	<tsk_play_volumn val ="任务执行时输出音量大小" / >
	//	<tsk_freeze val ="0 / 1[是否冻结任务]" / >
	//	<tsk_trans_proto val ="1[tcp]2[udp]3[udp组播]" / >
	//	< / row>
	//	<srcs>
	//	<src_info val ="根据tsk_src_type值，可能是文件全路径、声卡标识、用户ID" / >
	//	......
	//	<src_info val ="根据tsk_src_type值，可能是文件全路径、声卡标识、用户ID" / >
	//	< / srcs>
	//	<dsts>
	//	<usr_guid val ="" / >
	//	......
	//	<usr_guid val ="" / >
	//	< / dsts>
	//	< / tsk>
	//	< / root>

	//编辑任务，类似添加任务，但需要指定 tsk_guid
	//<root ord = "task_edit" tsk_guid = "" ...>

	//添加或编辑用户任务， str_xml内容如上描述，执行task_xxx后，会有相应task_xxx_ret的回调消息
	//其中 回调参数key 是ord的值, 回调参数value 表示这个xml的字符串,内容如下
	//<root ord="task_add_ret/task_edit_ret/task_del_ret" sess="1" id="" ret="ok/failed_db/failed_running" />
	//<root ord="task_start_ret" sess="1" id="" ret="ok/failed/running" / >
	//<root ord="task_stop_ret" sess="1" id="" ret="ok/failed" / >

	int task_add_or_edit(int sess_client, const std::string& str_xml);
	int task_del(int sess_client, const std::string& tsk_guid);
	int task_start(int sess_client, const std::string& tsk_guid);
	int task_stop(int sess_client, const std::string& tsk_guid);
	int req_idr(int sess_client, const std::string& tsk_guid);
	void ctrl_usr_dec_play(bool flag_dec_play);
	int simulate_event_key(int sess_client, const std::string& usr_guid);
	int req_cap_pic(int sess_client, 
		const std::string& usr_guid, const std::string& path);

	int file_transfer(int sess_client, const std::string& full_path);
	int file_transfer_cancel(int sess_client, int id);

	//start_date 例如	2018-12-01 
	//end_date	 例如	2018-12-15
	//log_type 0
	int query_talk_log(int sess_client, 
		const std::string& start_date, 
		const std::string& end_date, int log_type);

	//caller_type/callee_type 1[station_user]  2[embedded-linux term] 0[unknown]
	int query_talk_sto(int sess_client, int log_id, 
		const std::string& caller_name, 
		const std::string& callee_name, 
		const std::string& call_time, 
		int caller_type, int callee_type);

	//根据初始化接口pre_init() 的 aud_type参数决定输入的数据类型
	//若aud_type参数值默认为0，此处buf参数 传送 44.1K的采样率，G711编码数据
	//若aud_type参数值设置为1，此处buf参数 传送   32K的采样率， PCM格式数据
	int send_aud_data(const std::string& buf);

	void timer_test_callback_start();
	void timer_test_callback_stop();
};

#endif
