// Copyright 2019 Authors of protobuf-gadgets
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docmaster

import (
	"github.com/dunbit/protobuf-gadgets/pkg/protoc"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// RenderTemplates ...
func RenderTemplates(options *Options, data *protoc.Result) ([]*plugin_go.CodeGeneratorResponse_File, error) {
	files := make([]*plugin_go.CodeGeneratorResponse_File, 2)

	return files, nil
}
