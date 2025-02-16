/*
Copyright 2019 LitmusChaos Authors

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

// To create logs for debugging or detailing, please follow this syntax.
// use function log.Info
// in parameters give the name of the log / error (string) ,
// with the variable name for the value(string)
// and then the value to log (any datatype)
// All values should be in key : value pairs only
// For eg. : log.Info("name_of_the_log","variable_name_for_the_value",value, ......)
// For eg. : log.Error(err,"error_statement","variable_name",value)
// For eg. : log.Printf
//("error statement %q other variables %s/%s",targetValue, object.Namespace, object.Name)
// For eg. : log.Errorf
//("unable to reconcile object %s/%s: %v", object.Namespace, object.Name, err)
// This logger uses a structured logging schema in JSON format, which will / can be used further
// to access the values in the logger.

package types

import (
	litmuschaosv1alpha1 "github.com/litmuschaos/chaos-operator/api/litmuschaos/v1alpha1"
	"github.com/litmuschaos/chaos-operator/pkg/utils"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var (

	// Log with default name ie: controller_chaosengine
	Log = log.Log.WithName("controller_chaosengine")

	// DefaultChaosRunnerImage contains the default value of runner resource
	DefaultChaosRunnerImage = "litmuschaos/chaos-runner:latest"

	// ResultCRDName contains name of the chaosresult CRD
	ResultCRDName = "chaosresults.litmuschaos.io"
)

//EngineInfo Related information
type EngineInfo struct {
	Instance       *litmuschaosv1alpha1.ChaosEngine
	AppInfo        litmuschaosv1alpha1.ApplicationParams
	Selectors      *litmuschaosv1alpha1.Selector
	Targets        string
	VolumeOpts     utils.VolumeOpts
	AppExperiments []string
}
