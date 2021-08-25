#ifndef macro_func_cb_h__
#define macro_func_cb_h__

#include <string>

typedef void(__stdcall *func_cb)(unsigned long long usr_data,
	const std::string& type,
	const std::string& key,
	const std::string& value,
	const std::string& value2);

typedef void(__stdcall *func_cb_cli)(
	const int sess_client,
	const unsigned long long usr_data,
	const std::string& type,
	const std::string& key,
	const std::string& value,
	const std::string& value2);

typedef void(__stdcall *func_cb_mic)(unsigned long long usr_data, const std::string& data,
	int dev_id, int enc_type, int channel, int sample, int bitrate);

typedef void(__stdcall *func_cb_file)(
	const int sess_client, 
	const unsigned long long usr_data,
	const int id, 
	long long file_size, 
	const std::string& full_path, 
	const std::string& key, 
	const std::string& buf);

typedef void(__stdcall *func_cb_h26x)(
	const int sess_client, const unsigned long long usr_data,
	const std::string& sender,
	int sender_type, int ch_idx, int strm,
	const std::string& buf,
	int key_frm, int width, int height, int fps, int enc_tp);

typedef void(__stdcall *func_cb_pcm)(
	const int sess_client, 
	const unsigned long long usr_data,
	const std::string& sender,
	int sender_type, 
	int ch_idx, 
	int strm, 
	const std::string& buf,
	int channel, 
	int bitrate, 
	int sample, 
	int flag_sync);


#endif

