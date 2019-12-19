/*
Copyright © 2019 AWS Controller authors

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

package e2e

import (
	// Purely so that we can get proper test coverage
	_ "go.awsctrl.io/manager/controllers/apigateway"
	_ "go.awsctrl.io/manager/controllers/cloud9"
	_ "go.awsctrl.io/manager/controllers/cloudformation"
	_ "go.awsctrl.io/manager/controllers/controllermanager"
	_ "go.awsctrl.io/manager/controllers/ecr"
	_ "go.awsctrl.io/manager/controllers/self"
)
