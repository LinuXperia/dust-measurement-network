/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.0-dev at Wed Aug 22 08:54:06 2018. */

#ifndef PB_PROTOBUF_APPLICATIONINFORMATION_PB_H_INCLUDED
#define PB_PROTOBUF_APPLICATIONINFORMATION_PB_H_INCLUDED
#include <pb.h>

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _protobuf_requestSoftwareInformation_RequestInformationType {
    protobuf_requestSoftwareInformation_RequestInformationType_FIRMWARE_VERSION = 1
} protobuf_requestSoftwareInformation_RequestInformationType;
#define _protobuf_requestSoftwareInformation_RequestInformationType_MIN protobuf_requestSoftwareInformation_RequestInformationType_FIRMWARE_VERSION
#define _protobuf_requestSoftwareInformation_RequestInformationType_MAX protobuf_requestSoftwareInformation_RequestInformationType_FIRMWARE_VERSION
#define _protobuf_requestSoftwareInformation_RequestInformationType_ARRAYSIZE ((protobuf_requestSoftwareInformation_RequestInformationType)(protobuf_requestSoftwareInformation_RequestInformationType_FIRMWARE_VERSION+1))

/* Struct definitions */
typedef struct _protobuf_firmwareVersion {
    uint32_t major;
    uint32_t minor;
    uint32_t patch;
    uint32_t hotfix;
/* @@protoc_insertion_point(struct:protobuf_firmwareVersion) */
} protobuf_firmwareVersion;

typedef struct _protobuf_requestSoftwareInformation {
    protobuf_requestSoftwareInformation_RequestInformationType information;
/* @@protoc_insertion_point(struct:protobuf_requestSoftwareInformation) */
} protobuf_requestSoftwareInformation;

/* Default values for struct fields */
extern const protobuf_requestSoftwareInformation_RequestInformationType protobuf_requestSoftwareInformation_information_default;

/* Initializer values for message structs */
#define protobuf_requestSoftwareInformation_init_default {protobuf_requestSoftwareInformation_RequestInformationType_FIRMWARE_VERSION}
#define protobuf_firmwareVersion_init_default    {0, 0, 0, 0}
#define protobuf_requestSoftwareInformation_init_zero {_protobuf_requestSoftwareInformation_RequestInformationType_MIN}
#define protobuf_firmwareVersion_init_zero       {0, 0, 0, 0}

/* Field tags (for use in manual encoding/decoding) */
#define protobuf_firmwareVersion_major_tag       1
#define protobuf_firmwareVersion_minor_tag       2
#define protobuf_firmwareVersion_patch_tag       3
#define protobuf_firmwareVersion_hotfix_tag      4
#define protobuf_requestSoftwareInformation_information_tag 1

/* Struct field encoding specification for nanopb */
extern const pb_field_t protobuf_requestSoftwareInformation_fields[2];
extern const pb_field_t protobuf_firmwareVersion_fields[5];

/* Maximum encoded size of messages (where known) */
#define protobuf_requestSoftwareInformation_size 2
#define protobuf_firmwareVersion_size            24

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define APPLICATIONINFORMATION_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif