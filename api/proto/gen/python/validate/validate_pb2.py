# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: validate/validate.proto
# Protobuf Python Version: 4.25.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import descriptor_pb2 as google_dot_protobuf_dot_descriptor__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17validate/validate.proto\x12\x08validate\x1a google/protobuf/descriptor.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc8\x08\n\nFieldRules\x12\x30\n\x07message\x18\x11 \x01(\x0b\x32\x16.validate.MessageRulesR\x07message\x12,\n\x05\x66loat\x18\x01 \x01(\x0b\x32\x14.validate.FloatRulesH\x00R\x05\x66loat\x12/\n\x06\x64ouble\x18\x02 \x01(\x0b\x32\x15.validate.DoubleRulesH\x00R\x06\x64ouble\x12,\n\x05int32\x18\x03 \x01(\x0b\x32\x14.validate.Int32RulesH\x00R\x05int32\x12,\n\x05int64\x18\x04 \x01(\x0b\x32\x14.validate.Int64RulesH\x00R\x05int64\x12/\n\x06uint32\x18\x05 \x01(\x0b\x32\x15.validate.UInt32RulesH\x00R\x06uint32\x12/\n\x06uint64\x18\x06 \x01(\x0b\x32\x15.validate.UInt64RulesH\x00R\x06uint64\x12/\n\x06sint32\x18\x07 \x01(\x0b\x32\x15.validate.SInt32RulesH\x00R\x06sint32\x12/\n\x06sint64\x18\x08 \x01(\x0b\x32\x15.validate.SInt64RulesH\x00R\x06sint64\x12\x32\n\x07\x66ixed32\x18\t \x01(\x0b\x32\x16.validate.Fixed32RulesH\x00R\x07\x66ixed32\x12\x32\n\x07\x66ixed64\x18\n \x01(\x0b\x32\x16.validate.Fixed64RulesH\x00R\x07\x66ixed64\x12\x35\n\x08sfixed32\x18\x0b \x01(\x0b\x32\x17.validate.SFixed32RulesH\x00R\x08sfixed32\x12\x35\n\x08sfixed64\x18\x0c \x01(\x0b\x32\x17.validate.SFixed64RulesH\x00R\x08sfixed64\x12)\n\x04\x62ool\x18\r \x01(\x0b\x32\x13.validate.BoolRulesH\x00R\x04\x62ool\x12/\n\x06string\x18\x0e \x01(\x0b\x32\x15.validate.StringRulesH\x00R\x06string\x12,\n\x05\x62ytes\x18\x0f \x01(\x0b\x32\x14.validate.BytesRulesH\x00R\x05\x62ytes\x12)\n\x04\x65num\x18\x10 \x01(\x0b\x32\x13.validate.EnumRulesH\x00R\x04\x65num\x12\x35\n\x08repeated\x18\x12 \x01(\x0b\x32\x17.validate.RepeatedRulesH\x00R\x08repeated\x12&\n\x03map\x18\x13 \x01(\x0b\x32\x12.validate.MapRulesH\x00R\x03map\x12&\n\x03\x61ny\x18\x14 \x01(\x0b\x32\x12.validate.AnyRulesH\x00R\x03\x61ny\x12\x35\n\x08\x64uration\x18\x15 \x01(\x0b\x32\x17.validate.DurationRulesH\x00R\x08\x64uration\x12\x38\n\ttimestamp\x18\x16 \x01(\x0b\x32\x18.validate.TimestampRulesH\x00R\ttimestampB\x06\n\x04type\"\xb0\x01\n\nFloatRules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x02R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x02R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x02R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x02R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x02R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x02R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x02R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb1\x01\n\x0b\x44oubleRules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x01R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x01R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x01R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x01R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x01R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x01R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x01R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb0\x01\n\nInt32Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x05R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x05R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x05R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x05R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x05R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x05R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x05R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb0\x01\n\nInt64Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x03R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x03R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x03R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x03R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x03R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x03R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x03R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb1\x01\n\x0bUInt32Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\rR\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\rR\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\rR\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\rR\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\rR\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\rR\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\rR\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb1\x01\n\x0bUInt64Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x04R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x04R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x04R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x04R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x04R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x04R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x04R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb1\x01\n\x0bSInt32Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x11R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x11R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x11R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x11R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x11R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x11R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x11R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb1\x01\n\x0bSInt64Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x12R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x12R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x12R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x12R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x12R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x12R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x12R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb2\x01\n\x0c\x46ixed32Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x07R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x07R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x07R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x07R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x07R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x07R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x07R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb2\x01\n\x0c\x46ixed64Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x06R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x06R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x06R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x06R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x06R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x06R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x06R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb3\x01\n\rSFixed32Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x0fR\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x0fR\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x0fR\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x0fR\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x0fR\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x0fR\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x0fR\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"\xb3\x01\n\rSFixed64Rules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x10R\x05\x63onst\x12\x0e\n\x02lt\x18\x02 \x01(\x10R\x02lt\x12\x10\n\x03lte\x18\x03 \x01(\x10R\x03lte\x12\x0e\n\x02gt\x18\x04 \x01(\x10R\x02gt\x12\x10\n\x03gte\x18\x05 \x01(\x10R\x03gte\x12\x0e\n\x02in\x18\x06 \x03(\x10R\x02in\x12\x15\n\x06not_in\x18\x07 \x03(\x10R\x05notIn\x12!\n\x0cignore_empty\x18\x08 \x01(\x08R\x0bignoreEmpty\"!\n\tBoolRules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x08R\x05\x63onst\"\xd4\x05\n\x0bStringRules\x12\x14\n\x05\x63onst\x18\x01 \x01(\tR\x05\x63onst\x12\x10\n\x03len\x18\x13 \x01(\x04R\x03len\x12\x17\n\x07min_len\x18\x02 \x01(\x04R\x06minLen\x12\x17\n\x07max_len\x18\x03 \x01(\x04R\x06maxLen\x12\x1b\n\tlen_bytes\x18\x14 \x01(\x04R\x08lenBytes\x12\x1b\n\tmin_bytes\x18\x04 \x01(\x04R\x08minBytes\x12\x1b\n\tmax_bytes\x18\x05 \x01(\x04R\x08maxBytes\x12\x18\n\x07pattern\x18\x06 \x01(\tR\x07pattern\x12\x16\n\x06prefix\x18\x07 \x01(\tR\x06prefix\x12\x16\n\x06suffix\x18\x08 \x01(\tR\x06suffix\x12\x1a\n\x08\x63ontains\x18\t \x01(\tR\x08\x63ontains\x12!\n\x0cnot_contains\x18\x17 \x01(\tR\x0bnotContains\x12\x0e\n\x02in\x18\n \x03(\tR\x02in\x12\x15\n\x06not_in\x18\x0b \x03(\tR\x05notIn\x12\x16\n\x05\x65mail\x18\x0c \x01(\x08H\x00R\x05\x65mail\x12\x1c\n\x08hostname\x18\r \x01(\x08H\x00R\x08hostname\x12\x10\n\x02ip\x18\x0e \x01(\x08H\x00R\x02ip\x12\x14\n\x04ipv4\x18\x0f \x01(\x08H\x00R\x04ipv4\x12\x14\n\x04ipv6\x18\x10 \x01(\x08H\x00R\x04ipv6\x12\x12\n\x03uri\x18\x11 \x01(\x08H\x00R\x03uri\x12\x19\n\x07uri_ref\x18\x12 \x01(\x08H\x00R\x06uriRef\x12\x1a\n\x07\x61\x64\x64ress\x18\x15 \x01(\x08H\x00R\x07\x61\x64\x64ress\x12\x14\n\x04uuid\x18\x16 \x01(\x08H\x00R\x04uuid\x12@\n\x10well_known_regex\x18\x18 \x01(\x0e\x32\x14.validate.KnownRegexH\x00R\x0ewellKnownRegex\x12\x1c\n\x06strict\x18\x19 \x01(\x08:\x04trueR\x06strict\x12!\n\x0cignore_empty\x18\x1a \x01(\x08R\x0bignoreEmptyB\x0c\n\nwell_known\"\xe2\x02\n\nBytesRules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x0cR\x05\x63onst\x12\x10\n\x03len\x18\r \x01(\x04R\x03len\x12\x17\n\x07min_len\x18\x02 \x01(\x04R\x06minLen\x12\x17\n\x07max_len\x18\x03 \x01(\x04R\x06maxLen\x12\x18\n\x07pattern\x18\x04 \x01(\tR\x07pattern\x12\x16\n\x06prefix\x18\x05 \x01(\x0cR\x06prefix\x12\x16\n\x06suffix\x18\x06 \x01(\x0cR\x06suffix\x12\x1a\n\x08\x63ontains\x18\x07 \x01(\x0cR\x08\x63ontains\x12\x0e\n\x02in\x18\x08 \x03(\x0cR\x02in\x12\x15\n\x06not_in\x18\t \x03(\x0cR\x05notIn\x12\x10\n\x02ip\x18\n \x01(\x08H\x00R\x02ip\x12\x14\n\x04ipv4\x18\x0b \x01(\x08H\x00R\x04ipv4\x12\x14\n\x04ipv6\x18\x0c \x01(\x08H\x00R\x04ipv6\x12!\n\x0cignore_empty\x18\x0e \x01(\x08R\x0bignoreEmptyB\x0c\n\nwell_known\"k\n\tEnumRules\x12\x14\n\x05\x63onst\x18\x01 \x01(\x05R\x05\x63onst\x12!\n\x0c\x64\x65\x66ined_only\x18\x02 \x01(\x08R\x0b\x64\x65\x66inedOnly\x12\x0e\n\x02in\x18\x03 \x03(\x05R\x02in\x12\x15\n\x06not_in\x18\x04 \x03(\x05R\x05notIn\">\n\x0cMessageRules\x12\x12\n\x04skip\x18\x01 \x01(\x08R\x04skip\x12\x1a\n\x08required\x18\x02 \x01(\x08R\x08required\"\xb0\x01\n\rRepeatedRules\x12\x1b\n\tmin_items\x18\x01 \x01(\x04R\x08minItems\x12\x1b\n\tmax_items\x18\x02 \x01(\x04R\x08maxItems\x12\x16\n\x06unique\x18\x03 \x01(\x08R\x06unique\x12*\n\x05items\x18\x04 \x01(\x0b\x32\x14.validate.FieldRulesR\x05items\x12!\n\x0cignore_empty\x18\x05 \x01(\x08R\x0bignoreEmpty\"\xdc\x01\n\x08MapRules\x12\x1b\n\tmin_pairs\x18\x01 \x01(\x04R\x08minPairs\x12\x1b\n\tmax_pairs\x18\x02 \x01(\x04R\x08maxPairs\x12\x1b\n\tno_sparse\x18\x03 \x01(\x08R\x08noSparse\x12(\n\x04keys\x18\x04 \x01(\x0b\x32\x14.validate.FieldRulesR\x04keys\x12,\n\x06values\x18\x05 \x01(\x0b\x32\x14.validate.FieldRulesR\x06values\x12!\n\x0cignore_empty\x18\x06 \x01(\x08R\x0bignoreEmpty\"M\n\x08\x41nyRules\x12\x1a\n\x08required\x18\x01 \x01(\x08R\x08required\x12\x0e\n\x02in\x18\x02 \x03(\tR\x02in\x12\x15\n\x06not_in\x18\x03 \x03(\tR\x05notIn\"\xe9\x02\n\rDurationRules\x12\x1a\n\x08required\x18\x01 \x01(\x08R\x08required\x12/\n\x05\x63onst\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\x05\x63onst\x12)\n\x02lt\x18\x03 \x01(\x0b\x32\x19.google.protobuf.DurationR\x02lt\x12+\n\x03lte\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationR\x03lte\x12)\n\x02gt\x18\x05 \x01(\x0b\x32\x19.google.protobuf.DurationR\x02gt\x12+\n\x03gte\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationR\x03gte\x12)\n\x02in\x18\x07 \x03(\x0b\x32\x19.google.protobuf.DurationR\x02in\x12\x30\n\x06not_in\x18\x08 \x03(\x0b\x32\x19.google.protobuf.DurationR\x05notIn\"\xf3\x02\n\x0eTimestampRules\x12\x1a\n\x08required\x18\x01 \x01(\x08R\x08required\x12\x30\n\x05\x63onst\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x05\x63onst\x12*\n\x02lt\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02lt\x12,\n\x03lte\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x03lte\x12*\n\x02gt\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02gt\x12,\n\x03gte\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x03gte\x12\x15\n\x06lt_now\x18\x07 \x01(\x08R\x05ltNow\x12\x15\n\x06gt_now\x18\x08 \x01(\x08R\x05gtNow\x12\x31\n\x06within\x18\t \x01(\x0b\x32\x19.google.protobuf.DurationR\x06within*F\n\nKnownRegex\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x14\n\x10HTTP_HEADER_NAME\x10\x01\x12\x15\n\x11HTTP_HEADER_VALUE\x10\x02:<\n\x08\x64isabled\x12\x1f.google.protobuf.MessageOptions\x18\xaf\x08 \x01(\x08R\x08\x64isabled::\n\x07ignored\x12\x1f.google.protobuf.MessageOptions\x18\xb0\x08 \x01(\x08R\x07ignored::\n\x08required\x12\x1d.google.protobuf.OneofOptions\x18\xaf\x08 \x01(\x08R\x08required:J\n\x05rules\x12\x1d.google.protobuf.FieldOptions\x18\xaf\x08 \x01(\x0b\x32\x14.validate.FieldRulesR\x05rulesB\x94\x01\n\x0c\x63om.validateB\rValidateProtoP\x01Z5github.com/raptor-ml/raptor/api/proto/gen/go/validate\xa2\x02\x03VXX\xaa\x02\x08Validate\xca\x02\x08Validate\xe2\x02\x14Validate\\GPBMetadata\xea\x02\x08Validate')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'validate.validate_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\014com.validateB\rValidateProtoP\001Z5github.com/raptor-ml/raptor/api/proto/gen/go/validate\242\002\003VXX\252\002\010Validate\312\002\010Validate\342\002\024Validate\\GPBMetadata\352\002\010Validate'
  _globals['_KNOWNREGEX']._serialized_start=5909
  _globals['_KNOWNREGEX']._serialized_end=5979
  _globals['_FIELDRULES']._serialized_start=137
  _globals['_FIELDRULES']._serialized_end=1233
  _globals['_FLOATRULES']._serialized_start=1236
  _globals['_FLOATRULES']._serialized_end=1412
  _globals['_DOUBLERULES']._serialized_start=1415
  _globals['_DOUBLERULES']._serialized_end=1592
  _globals['_INT32RULES']._serialized_start=1595
  _globals['_INT32RULES']._serialized_end=1771
  _globals['_INT64RULES']._serialized_start=1774
  _globals['_INT64RULES']._serialized_end=1950
  _globals['_UINT32RULES']._serialized_start=1953
  _globals['_UINT32RULES']._serialized_end=2130
  _globals['_UINT64RULES']._serialized_start=2133
  _globals['_UINT64RULES']._serialized_end=2310
  _globals['_SINT32RULES']._serialized_start=2313
  _globals['_SINT32RULES']._serialized_end=2490
  _globals['_SINT64RULES']._serialized_start=2493
  _globals['_SINT64RULES']._serialized_end=2670
  _globals['_FIXED32RULES']._serialized_start=2673
  _globals['_FIXED32RULES']._serialized_end=2851
  _globals['_FIXED64RULES']._serialized_start=2854
  _globals['_FIXED64RULES']._serialized_end=3032
  _globals['_SFIXED32RULES']._serialized_start=3035
  _globals['_SFIXED32RULES']._serialized_end=3214
  _globals['_SFIXED64RULES']._serialized_start=3217
  _globals['_SFIXED64RULES']._serialized_end=3396
  _globals['_BOOLRULES']._serialized_start=3398
  _globals['_BOOLRULES']._serialized_end=3431
  _globals['_STRINGRULES']._serialized_start=3434
  _globals['_STRINGRULES']._serialized_end=4158
  _globals['_BYTESRULES']._serialized_start=4161
  _globals['_BYTESRULES']._serialized_end=4515
  _globals['_ENUMRULES']._serialized_start=4517
  _globals['_ENUMRULES']._serialized_end=4624
  _globals['_MESSAGERULES']._serialized_start=4626
  _globals['_MESSAGERULES']._serialized_end=4688
  _globals['_REPEATEDRULES']._serialized_start=4691
  _globals['_REPEATEDRULES']._serialized_end=4867
  _globals['_MAPRULES']._serialized_start=4870
  _globals['_MAPRULES']._serialized_end=5090
  _globals['_ANYRULES']._serialized_start=5092
  _globals['_ANYRULES']._serialized_end=5169
  _globals['_DURATIONRULES']._serialized_start=5172
  _globals['_DURATIONRULES']._serialized_end=5533
  _globals['_TIMESTAMPRULES']._serialized_start=5536
  _globals['_TIMESTAMPRULES']._serialized_end=5907
# @@protoc_insertion_point(module_scope)
