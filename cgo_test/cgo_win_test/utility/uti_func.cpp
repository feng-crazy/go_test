#include "uti_func.h"

int uti_func::set_xml(const std::string& str_xml, CMarkupSTL& xml){
	if (str_xml.empty())	return -1;

	if (!xml.SetDoc(str_xml.c_str()) || !xml.IsWellFormed() || !xml.FindElem())
		return -1;

	return 0;
}
