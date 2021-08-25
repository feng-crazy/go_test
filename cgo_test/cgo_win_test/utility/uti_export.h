#ifndef uti_export_h__
#define uti_export_h__

#ifdef WIN32
#define UTI_EXPORT __declspec(dllexport)
#else
#define UTI_EXPORT __attribute__((visibility("default")))
#endif
	

#endif
