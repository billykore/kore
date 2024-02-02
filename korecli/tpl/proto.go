package tpl

func ProtoTemplate() []byte {
	return []byte(`syntax = "proto3";

package kore.v1;

option go_package = "kore/libs/proto/v1;v1";

import "third_party/google/api/annotations.proto";

service {{ .StructName }} {
  rpc Greet ({{ .StructName }}Request) returns ({{ .StructName }}Reply) {
    option (google.api.http) = {
      get: "/greet"
    };
  }
}

message {{ .StructName }}Request {
  string name = 1;
}

message {{ .StructName }}Reply {
  string message = 1;
}
`)
}
