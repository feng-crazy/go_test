#ifndef win_sta_c_h__
#define win_sta_c_h__

#include "../utility/uti_export.h"

typedef void(__cdecl *func_cb_cli_c)(
	const int sess_client,
	const  unsigned long long usr_data,
	const char* type,
	const char* key,
	const char* value,
	const char* value2);

#ifdef __cplusplus
extern "C"{
#endif

	UTI_EXPORT void __cdecl win_sta_pre_init(const char* module_path, 
		func_cb_cli_c func, unsigned long long usr_data, int aud_type);
	UTI_EXPORT void __cdecl win_sta_aft_fini();
	UTI_EXPORT int __cdecl win_sta_init(const char* svr_ip, unsigned short svr_port,
		const char* usr_name, const char* usr_pwd);
	UTI_EXPORT void __cdecl win_sta_fini(int sess_client);
	UTI_EXPORT void __cdecl win_sta_fini_all();
	UTI_EXPORT int __cdecl win_sta_exec_manual_tsk_speak(int sess_client, int md_tp,
		int trans_proto, const char* list_dst_guid, int vol);
	UTI_EXPORT int __cdecl win_sta_exec_manual_tsk_mp3_broad(
		int sess_client, int md_tp, int trans_proto,
		const char* list_dst_guid, int vol,
		const char* list_src_path);
	UTI_EXPORT int __cdecl win_sta_exec_manual_tsk_mp3(
		int sess_client, int md_tp, int trans_proto,
		const char* list_dst_guid,
		const char* list_src_mp3,
		int vol, int flag_random);
	UTI_EXPORT int __cdecl win_sta_stop_task(int sess_client, const char* tsk_guid);
	UTI_EXPORT int __cdecl win_sta_start_monitor(int sess_client, const char* callee_guid,
		int ch_idx, int strm, int md_tp, int proto, unsigned long long wnd);
	UTI_EXPORT int __cdecl win_sta_usr_call_req(int sess_client, const char* callee_guid,
		int ch_idx, int strm, int md_tp, int proto, unsigned long long wnd);
	UTI_EXPORT int __cdecl win_sta_usr_call_req_ret(int sess_client, int ch, int strm, 
		const char* tsk_guid, const char* call_guid, int status, unsigned long long wnd);
	UTI_EXPORT int __cdecl win_sta_start_av_ret(int sess_client, const char* tsk_guid,
		const char* ret, int caller_sess, int sample, const char* aud_enc);
	UTI_EXPORT int __cdecl win_sta_start_av_mp3_ret(int sess_client, const char* tsk_guid,
		const char* ret, int caller_sess, int sample, const char* aud_enc);
	UTI_EXPORT int __cdecl win_sta_pre_start_aud(int sess_client, const char* callee_guid);
	UTI_EXPORT int __cdecl win_sta_task_add_or_edit(int sess_client, const char* str_xml);
	UTI_EXPORT int __cdecl win_sta_task_del(int sess_client, const char* tsk_guid);
	UTI_EXPORT int __cdecl win_sta_task_start(int sess_client, const char* tsk_guid);
	UTI_EXPORT int __cdecl win_sta_task_stop(int sess_client, const char* tsk_guid);
	UTI_EXPORT int __cdecl win_sta_req_idr(int sess_client, const char* tsk_guid);
	UTI_EXPORT void __cdecl win_sta_ctrl_usr_dec_play(bool flag_dec_play);
	UTI_EXPORT int __cdecl win_simulate_event_key(int sess_client, const char* tsk_guid);
	UTI_EXPORT int __cdecl win_sta_req_cap_pic(int sess_client, const char* usr_guid, const char* path);

	UTI_EXPORT int __cdecl win_sta_file_transfer(int sess_client, const char* full_path);
	UTI_EXPORT int __cdecl win_sta_file_transfer_cancel(int sess_client, int id);
	UTI_EXPORT int __cdecl win_sta_query_talk_log(int sess_client, 
		const char* start_date, const char* end_date, int log_type);
	UTI_EXPORT int __cdecl win_sta_query_talk_sto(int sess_client, int log_id, 
		const char* caller_name, const char* callee_name, 
		const char* call_time, int caller_type, int callee_type);

	UTI_EXPORT int __cdecl win_sta_send_aud_data(
		int sess_client, const char* buf, int buf_len);


#ifdef __cplusplus
}
#endif

#endif 