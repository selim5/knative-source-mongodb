/*
Copyright 2020 Google LLC

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

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

const (
	// MongoDbConditionReady has status True when the MongoDbSource is ready to send events.
	MongoDbConditionReady = apis.ConditionReady

	// MongoDbConditionSinkProvided has status True when the MongoDbSource has been configured with a sink target.
	MongoDbConditionSinkProvided apis.ConditionType = "SinkProvided"

	// MongoDbConditionConnectionEstablished has status True when the MongoDbSource has correct credentials and
	// was able to successfully connect to the correct database or collection.
	MongoDbConditionConnectionEstablished apis.ConditionType = "ConnectionEstablished"

	// MongoDbConditionDeployed has status True when the MongoDbSource has had it's deployment created.
	MongoDbConditionDeployed apis.ConditionType = "Deployed"
)

// MongoDbCondSet holds NewLivingConditionSet.
var MongoDbCondSet = apis.NewLivingConditionSet(
	MongoDbConditionSinkProvided,
	MongoDbConditionConnectionEstablished,
	MongoDbConditionDeployed,
)

// GetConditionSet retrieves the condition set for this resource. Implements the KRShaped interface.
func (*MongoDbSource) GetConditionSet() apis.ConditionSet {
	return MongoDbCondSet
}

// GetStatus retrieves the duck status for this resource. Implements the KRShaped interface.
func (m *MongoDbSource) GetStatus() *duckv1.Status {
	return &m.Status.Status
}

// GetCondition returns the condition currently associated with the given type, or nil.
func (m *MongoDbSourceStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return MongoDbCondSet.Manage(m).GetCondition(t)
}

// InitializeConditions sets relevant unset conditions to Unknown state.
func (m *MongoDbSourceStatus) InitializeConditions() {
	MongoDbCondSet.Manage(m).InitializeConditions()
}

// MarkSink sets the condition that the source has a sink configured.
func (m *MongoDbSourceStatus) MarkSink(uri *apis.URL) {
	m.SinkURI = uri
	if !uri.IsEmpty() {
		MongoDbCondSet.Manage(m).MarkTrue(MongoDbConditionSinkProvided)
	} else {
		MongoDbCondSet.Manage(m).MarkUnknown(MongoDbConditionSinkProvided, "SinkEmpty", "Sink has resolved to empty.%s", "")
	}
}

// MarkNoSink sets the condition that the source does not have a sink configured.
func (m *MongoDbSourceStatus) MarkNoSink(reason, messageFormat string, messageA ...interface{}) {
	MongoDbCondSet.Manage(m).MarkFalse(MongoDbConditionSinkProvided, reason, messageFormat, messageA...)
}

// MarkConnectionSuccess sets the condition that the source has correct credentials and that the specified database or collection is found.
func (m *MongoDbSourceStatus) MarkConnectionSuccess() {
	MongoDbCondSet.Manage(m).MarkTrue(MongoDbConditionConnectionEstablished)
}

// MarkConnectionFailed sets the condition that the source has incorrect credentials or that the specified database or collection is not found.
func (m *MongoDbSourceStatus) MarkConnectionFailed(err error) {
	MongoDbCondSet.Manage(m).MarkFalse(MongoDbConditionConnectionEstablished, "Connection failed: incorrect credentials or database or collection not found", err.Error())
}

// deploymentIsAvailable determines if the provided deployment is available. Note that if it cannot
// determine the Deployment's availability, it returns `def` (short for default). From https://github.com/knative/eventing/blob/master/pkg/apis/duck/lifecycle_helper.go .
func deploymentIsAvailable(d *appsv1.DeploymentStatus, def bool) bool {
	// Check if the Deployment is available.
	for _, cond := range d.Conditions {
		if cond.Type == appsv1.DeploymentAvailable {
			return cond.Status == "True"
		}
	}
	return def
}

// PropagateDeploymentAvailability uses the availability of the provided Deployment to determine if
// MongoDbConditionDeployed should be marked as true or false.
func (m *MongoDbSourceStatus) PropagateDeploymentAvailability(d *appsv1.Deployment) {
	if deploymentIsAvailable(&d.Status, false) {
		MongoDbCondSet.Manage(m).MarkTrue(MongoDbConditionDeployed)
	} else {
		// I don't know how to propagate the status well, so just give the name of the Deployment
		// for now.
		MongoDbCondSet.Manage(m).MarkFalse(MongoDbConditionDeployed, "DeploymentUnavailable", "The Deployment '%s' is unavailable.", d.Name)
	}
}

// IsReady returns true if the resource is ready overall.
func (m *MongoDbSourceStatus) IsReady() bool {
	return MongoDbCondSet.Manage(m).IsHappy()
}
