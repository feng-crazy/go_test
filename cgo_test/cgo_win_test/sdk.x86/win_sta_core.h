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

	//Ԥ��ʼ����module_path��ʾ��ǰ��������·����
	//func��ʾ�ӿ���Ϣ�ص���usr_data���ⲿ�Զ����������func�ص��Ǵ��ظ��ⲿ
	//aud_type,Ĭ�� 0-��ʾ��Ƶ����44.1K��G711����			1-��ʾ��Ƶ����32K��PCM��ʽ
	void pre_init(const std::string& module_path, func_cb_cli func, unsigned long long usr_data = 0, int aud_type = 0);

	void pre_init_java(const std::string& module_path, int aud_type);
	void aft_fini();

	//��ʼ�����ӷ���ˣ�svr_ip��ʾ�����IP��ַ��svr_port��ʾ����˶˿�
	//usr_name��ʾ�ͻ����û�����usr_pwd��ʾ�û�������
	//����ֵ ��ʾ�ͻ�����������ID��sess_client����1��ʼ������������ʾ����
	int init(const std::string& svr_ip, unsigned short svr_port,
		const std::string& usr_name, const std::string& usr_pwd);

	//�ر�����ͻ��˵���������
	void fini(int sess_client);

	//�ر����пͻ�����������
	void fini_all();

	//����ִ���ֶ�����
	//sess_client��ʾ��������ID
	//md_tp ý������  1[��Ƶ] 2[��Ƶ]
	//trans_proto ����Э�� 1[TCP]2[UDP]3[UDP�鲥]
	//list_dst_guid ��ʾ�����Ŀ���û�(�û�GUID����ʽ)
	//flag_usr_data ֵfalse��ʾ��win_sta_core�ڲ��ɼ���˷����ݣ�true��ʾ���û�����send_aud_44100_1ch_g711�ӿڣ��������� 
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

	//ֹͣ����tsk_guid��ʾ�����GUID
	int stop_task(int sess_client, const std::string& tsk_guid);

	//������Ӽ��������û�������Ƶ��
	//callee_guid��ʾ��������û�GUID
	//ch_idx ����Ƶͨ����  0,1,2......
	//strm ��ʾ�������ͣ���ҪӦ������Ƶ����, 0,1,2......
        //md_tp ý������  1[��Ƶ] 2[��Ƶ] 3[����Ƶ
	//proto ����Э�� 1[TCP]2[UDP]3[UDP�鲥]
	//ע�⣺ĿǰUDP��UDP�鲥 ֻ��������Ƶ�� ���������Ƶ����ѡ��TCP��ʽ
	//��������󣬻��лص���Ϣ�� �ص�����type ��ֵ��monitor_tsk_status
	//�ص�����value ��ֵ ��xml�ַ�������������GUID��������״̬status ����Ϣ
	int start_monitor(int sess_client, const std::string& callee_guid, 
		int ch_idx, int strm, int md_tp, int proto, unsigned long long wnd = 0);

	//�û������������ڷ������	
	//callee_guid��ʾ���е��û�GUID
	//����usr_call_req�������˺�����������û������ߣ����лص���Ϣ��
	// �ص�����type ��ֵ�� usr_call_req_ret�����������û�״̬��Ϣ������0��ʾ�����ߣ�

	//��ص�����type ��ֵ��tk_tsk_status��Ϣ��ʾ��������״̬��
	//�ص�����value ��ֵ ��xml�ַ���������Ϣ�а���tsk_guid, ����״̬����
	//    0[TSK_STATE_FREE], 1[TSK_STATE_AUTO],2[TSK_STATE_MANUAL]
	//    3[TSK_STATE_CALL_START],4[TSK_STATE_CALL_PICKUP] �� 5[TSK_STATE_CALL_NO_USR]
	int usr_call_req(int sess_client, const std::string& callee_guid, 
		int ch_idx, int strm, int md_tp, int proto, unsigned long long wnd = 0);

	//�û�����������Ӧ�����ڱ����ж� ���յ�������Ϣ(usr_call_req)ʱ�Ļ�Ӧ
	//tsk_guid ��ʾ����GUID,�ɷ���˷��������ж���Ϣusr_call_req�л�ȡ
	//status  0��ʾ�ܾ���1��ʾ��æ��2��ʾ��ͨ
	//����usr_call_req_ret������˺󣬷���˸���������ܻ᷵��usr_call_req_ret_fb��Ϣ����ʾ�ú����Ѿ���������already_pickup�������ں����ƶ������û���calling_other��
	//�½ӿ� wong-2018/10/10
	int usr_call_req_ret(int sess_client, int ch, int strm,
		const std::string& tsk_guid, const std::string &call_guid, int status, unsigned long long wnd = 0);

	int start_av_ret(int sess_client, const std::string& tsk_guid,
		const std::string& ret, int caller_sess, int sample, const std::string& aud_enc);

	int start_av_mp3_ret(int sess_client, const std::string& tsk_guid,
		const std::string& ret, int caller_sess, int sample, const std::string& aud_enc);

	int pre_start_aud(int sess_client, const std::string& callee_guid = "");

	//�������
	//<root ord ="task_add" >
	//	<tsk>
	//	<row>
	//	<tsk_guid val ="{ 151196D6 - 33D7 - 4426 - 9E3C - E374FE12AE51 }" / >
	//	<tsk_name val ="������[utf8����]" / >
	//	<tsk_usr_guid val ="{ B7BEE137 - 901C - 46C2 - A7F8 - DBA588A04774 }" / >
	//	<tsk_src_type val =" 0[��Դ] 1[�ļ�]2[ָ������]3[ָ��IP�ն˻��û�]" / >
	//	<tsk_file_sec val ="50[�ļ���ʱ������]" / >
	//	<tsk_job_type val ="0[ÿ������]1[һ������]2[�ֶ�����]7[ÿ������]" / >
	//	<tsk_start_time val ="2017 - 07 - 31 14:00 : 00[����ʼʱ��]" / >
	//	<tsk_has_dura val ="0 / 1[�Ƿ�ָ������ʱ��]" / >
	//	<tsk_duration val ="30[ָ������ʱ��������]" / >
	//	<tsk_has_stop_time val ="0 / 1[�Ƿ�ָ��������Ч��ֹʱ��]" / >
	//	<tsk_stop_time val ="2017 - 08 - 10 16:00 : 00[������Ч��ֹʱ��]" / >
	//	<tsk_week_data val =" 0110011[�����յ�����, 1��ʾ��Ч��0��ʾʧЧ]" / >
	//	<tsk_flag_random val ="0 / 1[�Ƿ������ʶ]" / >
	//	<tsk_flag_db_alarm val ="0 / 1[�Ƿ�������������]" / >
	//	<tsk_db_alm_level val ="50[���������ֱ�ֵ]" / >
	//	<tsk_alm_aud_file val ="���������ļ�ȫ·��" / >
	//	<tsk_only_power val ="0 / 1[�Ƿ�����ƹ��ŵ�Դ]" / >
	//	<tsk_force_signal val ="0 / 1[�Ƿ�ͬ������ǿ���ź�]" / >
	//	<tsk_play_volumn val ="����ִ��ʱ���������С" / >
	//	<tsk_freeze val ="0 / 1[�Ƿ񶳽�����]" / >
	//	<tsk_trans_proto val ="1[tcp]2[udp]3[udp�鲥]" / >
	//	< / row>
	//	<srcs>
	//	<src_info val ="����tsk_src_typeֵ���������ļ�ȫ·����������ʶ���û�ID" / >
	//	......
	//	<src_info val ="����tsk_src_typeֵ���������ļ�ȫ·����������ʶ���û�ID" / >
	//	< / srcs>
	//	<dsts>
	//	<usr_guid val ="" / >
	//	......
	//	<usr_guid val ="" / >
	//	< / dsts>
	//	< / tsk>
	//	< / root>

	//�༭��������������񣬵���Ҫָ�� tsk_guid
	//<root ord = "task_edit" tsk_guid = "" ...>

	//��ӻ�༭�û����� str_xml��������������ִ��task_xxx�󣬻�����Ӧtask_xxx_ret�Ļص���Ϣ
	//���� �ص�����key ��ord��ֵ, �ص�����value ��ʾ���xml���ַ���,��������
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

	//start_date ����	2018-12-01 
	//end_date	 ����	2018-12-15
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

	//���ݳ�ʼ���ӿ�pre_init() �� aud_type���������������������
	//��aud_type����ֵĬ��Ϊ0���˴�buf���� ���� 44.1K�Ĳ����ʣ�G711��������
	//��aud_type����ֵ����Ϊ1���˴�buf���� ����   32K�Ĳ����ʣ� PCM��ʽ����
	int send_aud_data(const std::string& buf);

	void timer_test_callback_start();
	void timer_test_callback_stop();
};

#endif
