/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package kubectl

import (
	"reflect"
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/runtime"
)

func TestSortingPrinter(t *testing.T) {
	tests := []struct {
		obj   runtime.Object
		sort  runtime.Object
		field string
		name  string
	}{
		{
			name: "in-order-already",
			obj: &api.PodList{
				Items: []api.Pod{
					{
						ObjectMeta: api.ObjectMeta{
							Name: "a",
						},
					},
					{
						ObjectMeta: api.ObjectMeta{
							Name: "b",
						},
					},
					{
						ObjectMeta: api.ObjectMeta{
							Name: "c",
						},
					},
				},
			},
			sort: &api.PodList{
				Items: []api.Pod{
					{
						ObjectMeta: api.ObjectMeta{
							Name: "a",
						},
					},
					{
						ObjectMeta: api.ObjectMeta{
							Name: "b",
						},
					},
					{
						ObjectMeta: api.ObjectMeta{
							Name: "c",
						},
					},
				},
			},
			field: "{.ObjectMeta.Name}",
		},
		{
			name: "reverse-order",
			obj: &api.PodList{
				Items: []api.Pod{
					{
						ObjectMeta: api.ObjectMeta{
							Name: "b",
						},
					},
					{
						ObjectMeta: api.ObjectMeta{
							Name: "c",
						},
					},
					{
						ObjectMeta: api.ObjectMeta{
							Name: "a",
						},
					},
				},
			},
			sort: &api.PodList{
				Items: []api.Pod{
					{
						ObjectMeta: api.ObjectMeta{
							Name: "a",
						},
					},
					{
						ObjectMeta: api.ObjectMeta{
							Name: "b",
						},
					},
					{
						ObjectMeta: api.ObjectMeta{
							Name: "c",
						},
					},
				},
			},
			field: "{.ObjectMeta.Name}",
		},
		{
			name: "random-order-numbers",
			obj: &api.ReplicationControllerList{
				Items: []api.ReplicationController{
					{
						Spec: api.ReplicationControllerSpec{
							Replicas: 5,
						},
					},
					{
						Spec: api.ReplicationControllerSpec{
							Replicas: 1,
						},
					},
					{
						Spec: api.ReplicationControllerSpec{
							Replicas: 9,
						},
					},
				},
			},
			sort: &api.ReplicationControllerList{
				Items: []api.ReplicationController{
					{
						Spec: api.ReplicationControllerSpec{
							Replicas: 1,
						},
					},
					{
						Spec: api.ReplicationControllerSpec{
							Replicas: 5,
						},
					},
					{
						Spec: api.ReplicationControllerSpec{
							Replicas: 9,
						},
					},
				},
			},
			field: "{.Spec.Replicas}",
		},
	}
	for _, test := range tests {
		sort := &SortingPrinter{SortField: test.field}
		if err := sort.sortObj(test.obj); err != nil {
			t.Errorf("unexpected error: %v (%s)", err, test.name)
			continue
		}
		if !reflect.DeepEqual(test.obj, test.sort) {
			t.Errorf("[%s]\nexpected:\n%v\nsaw:\n%v", test.name, test.sort, test.obj)
		}
	}
}
