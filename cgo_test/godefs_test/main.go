package main

/*
#include <stdlib.h>       // for free()
#include <stdint.h>
#include <stdbool.h>

typedef struct BacNetDevice {
    uint32_t deviceId;
    uint16_t devicePort;
    char *name;
    char *description;
    char *profile;
    struct BacNetDevice *next;
} BacNetDevice;

typedef struct BacNetUnionValue{
    int8_t i8;
    uint8_t ui8;
    int16_t i16;
    uint16_t ui16;
    int32_t i32;
    uint32_t ui32;
    int64_t i64;
    uint64_t ui64;
    float f32;
    double f64;
    bool bl;
    char * str;
}BacNetUnionValue;


typedef struct BacNetValue{
    uint32_t instance;
    uint32_t type;
    BacNetUnionValue value;
    struct BacNetValue *next;
}BacNetValue;

typedef struct BacNetMultiPropertyData {
    uint32_t type;
    uint32_t property;
    uint32_t instance;
    uint32_t index;
    BacNetValue value;
    uint8_t priority;
    struct BacNetMultiPropertyData *next;
} BacNetMultiPropertyData;
*/
import "C"

type BacNetDevice C.BacNetDevice
type BacNetUnionValue C.BacNetUnionValue
type BacNetValue C.BacNetValue
type BacNetMultiPropertyData C.BacNetMultiPropertyData

const Sizeof_BacNetDevice = C.sizeof_BacNetDevice
const Sizeof_BacNetUnionValue = C.sizeof_BacNetUnionValue
const Sizeof_BacNetValue = C.sizeof_BacNetValue
const Sizeof_BacNetMultiPropertyData = C.sizeof_BacNetMultiPropertyData
