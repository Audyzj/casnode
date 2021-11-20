// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package discuzx

import (
	"github.com/astaxie/beego"
	"github.com/casbin/casnode/object"
)

var dbName = "ultrax"
var ossEndpoint = ""
var ossAccessKeyId = ""
var ossAccessKeySecret = ""
var ossBucketName = "casnode"
var cdnDomain = "https://cdn.casbin.com/"
var discuzxDomain = "https://forum.casbin.com/"

var CasdoorOrganization = ""
var CasdoorApplication = ""

func init() {
	object.InitConfig()

	CasdoorOrganization = beego.AppConfig.String("casdoorOrganization")
	CasdoorApplication = beego.AppConfig.String("casdoorApplication")
}
