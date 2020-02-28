/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// HttpResponseRule HAProxy HTTP response rule configuration (corresponds to http-response directives)
type HttpResponseRule struct {
	AclFile      string `json:"acl_file,omitempty"`
	AclKeyfmt    string `json:"acl_keyfmt,omitempty"`
	Cond         string `json:"cond,omitempty"`
	CondTest     string `json:"cond_test,omitempty"`
	HdrFormat    string `json:"hdr_format,omitempty"`
	HdrMatch     string `json:"hdr_match,omitempty"`
	HdrName      string `json:"hdr_name,omitempty"`
	Id           *int32 `json:"id"`
	LogLevel     string `json:"log_level,omitempty"`
	RedirCode    int32  `json:"redir_code,omitempty"`
	RedirOption  string `json:"redir_option,omitempty"`
	RedirType    string `json:"redir_type,omitempty"`
	RedirValue   string `json:"redir_value,omitempty"`
	SpoeEngine   string `json:"spoe_engine,omitempty"`
	SpoeGroup    string `json:"spoe_group,omitempty"`
	Status       int32  `json:"status,omitempty"`
	StatusReason string `json:"status_reason,omitempty"`
	Type         string `json:"type"`
	VarExpr      string `json:"var_expr,omitempty"`
	VarName      string `json:"var_name,omitempty"`
	VarScope     string `json:"var_scope,omitempty"`
}
