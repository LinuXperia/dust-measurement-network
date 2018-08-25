/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.0-dev at Wed Aug 22 08:54:06 2018. */

#ifndef PB_PROTOBUF_BATTERYSTATE_PB_H_INCLUDED
#define PB_PROTOBUF_BATTERYSTATE_PB_H_INCLUDED
#include <pb.h>

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _protobuf_batteryState_StateType {
    protobuf_batteryState_StateType_UNDEFINED = 0,
    protobuf_batteryState_StateType_OK = 1,
    protobuf_batteryState_StateType_WARNING = 2,
    protobuf_batteryState_StateType_STOP = 3,
    protobuf_batteryState_StateType_SHUTDOWN = 4
} protobuf_batteryState_StateType;
#define _protobuf_batteryState_StateType_MIN protobuf_batteryState_StateType_UNDEFINED
#define _protobuf_batteryState_StateType_MAX protobuf_batteryState_StateType_SHUTDOWN
#define _protobuf_batteryState_StateType_ARRAYSIZE ((protobuf_batteryState_StateType)(protobuf_batteryState_StateType_SHUTDOWN+1))

/* Struct definitions */
typedef struct _protobuf_batteryState {
    protobuf_batteryState_StateType state;
    float voltage;
/* @@protoc_insertion_point(struct:protobuf_batteryState) */
} protobuf_batteryState;

/* Default values for struct fields */
extern const protobuf_batteryState_StateType protobuf_batteryState_state_default;

/* Initializer values for message structs */
#define protobuf_batteryState_init_default       {protobuf_batteryState_StateType_UNDEFINED, 0}
#define protobuf_batteryState_init_zero          {_protobuf_batteryState_StateType_MIN, 0}

/* Field tags (for use in manual encoding/decoding) */
#define protobuf_batteryState_state_tag          1
#define protobuf_batteryState_voltage_tag        2

/* Struct field encoding specification for nanopb */
extern const pb_field_t protobuf_batteryState_fields[3];

/* Maximum encoded size of messages (where known) */
#define protobuf_batteryState_size               7

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define BATTERYSTATE_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif