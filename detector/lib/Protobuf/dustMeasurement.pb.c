/* Automatically generated nanopb constant definitions */
/* Generated by nanopb-0.4.0-dev at Wed Aug 22 07:15:06 2018. */

#include "dustMeasurement.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif



const pb_field_t dustSensorMeasurement_fields[5] = {
    PB_FIELD(  1, FLOAT   , REQUIRED, STATIC  , FIRST, dustSensorMeasurement, particularMatter2_5um, particularMatter2_5um, 0),
    PB_FIELD(  2, FLOAT   , REQUIRED, STATIC  , OTHER, dustSensorMeasurement, particularMatter10um, particularMatter2_5um, 0),
    PB_FIELD(  3, FLOAT   , REQUIRED, STATIC  , OTHER, dustSensorMeasurement, temperature, particularMatter10um, 0),
    PB_FIELD(  4, FLOAT   , REQUIRED, STATIC  , OTHER, dustSensorMeasurement, humidity, temperature, 0),
    PB_LAST_FIELD
};


/* @@protoc_insertion_point(eof) */
