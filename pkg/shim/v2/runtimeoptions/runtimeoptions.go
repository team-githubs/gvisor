// Copyright 2018 The containerd Authors.
// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runtimeoptions

import (
	proto "github.com/gogo/protobuf/proto"
	pb "gvisor.dev/gvisor/pkg/shim/v2/runtimeoptions/api_go_proto"
)

type Options = pb.Options

func init() {
	// The generated proto file auto registers with "golang/protobuf/proto"
	// package. However, typeurl uses "golang/gogo/protobuf/proto". So registers
	// the type there too.
	proto.RegisterType((*Options)(nil), "cri.runtimeoptions.v1.Options")
}
