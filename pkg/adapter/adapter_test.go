/*
Copyright 2020 Google LLC.
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
package adapter

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	id    string = "id"
	db    string = "db"
	coll  string = "coll"
	docID string = "docID"
)

func TestMakeCloudEvent(t *testing.T) {
	tests := []struct {
		name    string
		a       *mongoDbAdapter
		data    bson.M
		wantErr bool
		wantCE  map[string]string
	}{
		// {
		// 	name: "error decoding bson",
		// 	a: &mongoDbAdapter{
		// 		namespace:      "namespace",
		// 		ceSourcePrefix: "CEPrefix/",
		// 		database:       "db",
		// 		collection:     "coll",
		// 	},
		// 	data: bson.M{
		// 		"misingNs": &bson.M{
		// 			"missing": "mising",
		// 		},
		// 		"operationType": "insert",
		// 	},
		// 	wantErr: true,
		// },
		{
			name: "unrecognizable type of change",
			a: &mongoDbAdapter{
				namespace:      "namespace",
				ceSourcePrefix: "CEPrefix/",
				database:       "db",
				collection:     "coll",
			},
			data: bson.M{
				"ns": bson.M{
					"coll": coll,
					"db":   db,
				},
				"_id": bson.M{
					"_data":       "IDofChange",
					"clusterTime": "",
				},
				"documentKey": bson.M{
					"_id": docID,
				},
				"fullDocument": bson.M{
					"_id":  docID,
					"key1": "value1",
				},
				"operationType": "NOTvalid",
			},
			wantErr: true,
		},
		{
			name: "Valid",
			a: &mongoDbAdapter{
				namespace:      "namespace",
				ceSourcePrefix: "CEPrefix/",
				database:       "db",
				collection:     "coll",
			},
			data: bson.M{
				"ns": bson.M{
					"coll": coll,
					"db":   db,
				},
				"_id": bson.M{
					"_data":       "IDofChange",
					"clusterTime": "",
				},
				"documentKey": bson.M{
					"_id": docID,
				},
				"fullDocument": bson.M{
					"_id":  docID,
					"key1": "value1",
				},
				"operationType": "insert",
			},

			wantErr: true,
			wantCE: map[string]string{
				"ID":        "CEID",
				"Source":    "CEPrefix/databases/db/collections/coll",
				"EventType": "inserted",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			receivedCE, err := test.a.makeCloudEvent(test.data)
			if err != nil {
				if !test.wantErr {
					t.Errorf("makeCloudEvent got error %q want error=%v", err, test.wantErr)
				}
			} else {
				if receivedCE.ID() != test.wantCE["ID"] {
					t.Errorf("Cloud Event created doesn't match desired Cloud Event from data. Field ID: Got=%v Want=%v", receivedCE.ID(), test.wantCE["ID"])
				}
				if receivedCE.Source() != test.wantCE["Source"] {
					t.Errorf("Cloud Event created doesn't match desired Cloud Event from data. Field Source: Got=%v Want=%v", receivedCE.Source(), test.wantCE["Source"])
				}
			}
		})
	}
}
