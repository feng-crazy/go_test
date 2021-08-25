package main

/*
#include <stdlib.h>       // for free()
#include <stdint.h>

#define MAX_CLASSIFICATION_NUM  5
struct iRF_Classification
{
	int		candidate_num;
	int		index[MAX_CLASSIFICATION_NUM];
	float	confidence[MAX_CLASSIFICATION_NUM];
};
typedef struct iRF_Bound_Tag
{
	float	score;
	struct iRF_Classification	classification;
	int		x_left_top;
	int		y_left_top;
	int		x_right_top;
	int		y_right_top;
	int		x_left_bottom;
	int		y_left_bottom;
} iRF_Bound;

#define MAX_OBJECT_NUM 3
typedef struct iRF_Object_Tag
{
	int			num;
	iRF_Bound	bound[MAX_OBJECT_NUM];
} iRF_Object;
*/
import "C"

type IRF_Classification C.struct_iRF_Classification
type IRF_Bound C.iRF_Bound
type IRF_Object C.iRF_Object

const Sizeof_IRF_Classification = C.sizeof_struct_iRF_Classification
const Sizeof_IRF_Bound = C.sizeof_iRF_Bound
const Sizeof_IRF_Object = C.sizeof_iRF_Object
