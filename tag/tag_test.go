//
// +build unit

/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2016 Intel Corporation

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

package tag

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"testing"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/ctypes"
	. "github.com/smartystreets/goconvey/convey"
)

//Random number generator
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func TestTagProcessor(t *testing.T) {
	meta := Meta()
	Convey("Meta should return metadata for the plugin", t, func() {
		Convey("So meta.Name should equal tag", func() {
			So(meta.Name, ShouldEqual, "tag")
		})
		Convey("So meta.Version should equal version", func() {
			So(meta.Version, ShouldEqual, version)
		})
		Convey("So meta.Type should be of type plugin.ProcessorPluginType", func() {
			So(meta.Type, ShouldResemble, plugin.ProcessorPluginType)
		})
	})

	proc := NewTagProcessor()
	Convey("Create tag processor", t, func() {
		Convey("So proc should not be nil", func() {
			So(proc, ShouldNotBeNil)
		})
		Convey("So proc should be of type tagProcessor", func() {
			So(proc, ShouldHaveSameTypeAs, &TagProcessor{})
		})
		Convey("proc.GetConfigPolicy should return a config policy", func() {
			configPolicy, _ := proc.GetConfigPolicy()
			Convey("So config policy should be a cpolicy.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
			})
			testConfig := make(map[string]ctypes.ConfigValue)
			testConfig["tags"] = ctypes.ConfigValueStr{Value: "test:test1"}
			cfg, errs := configPolicy.Get([]string{""}).Process(testConfig)
			Convey("So config policy should process testConfig and return a config", func() {
				So(cfg, ShouldNotBeNil)
			})
			Convey("So testConfig processing should return no errors", func() {
				So(errs.HasErrors(), ShouldBeFalse)
			})
		})
	})
}

func TestTagProcessorMetrics(t *testing.T) {
	Convey("Tag Processor tests", t, func() {
		metrics := make([]plugin.MetricType, 10)
		config := make(map[string]ctypes.ConfigValue)

		config["tags"] = ctypes.ConfigValueStr{Value: "rack:rack1,test:test1"}
		Convey("Tag on some data, when publisher Tags_ is empty", func() {
			for i := range metrics {
				time.Sleep(3)
				rand.Seed(time.Now().UTC().UnixNano())
				data := randInt(65, 90)
				metrics[i] = *plugin.NewMetricType(core.NewNamespace("foo", "bar"), time.Now(), nil, "", data)
			}
			So(metrics[0].Tags_, ShouldBeNil)
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			So(metrics[0].Tags_, ShouldBeNil)
			tagObj := NewTagProcessor()

			_, received_data, _ := tagObj.Process("snap.gob", buf.Bytes(), config)

			var metrics_new []plugin.MetricType

			//Decodes the content into pluginMetricType
			dec := gob.NewDecoder(bytes.NewBuffer(received_data))
			dec.Decode(&metrics_new)
			So(metrics_new[0].Tags_, ShouldNotBeNil)
			So(metrics_new[0].Tags_["rack"], ShouldResemble, "rack1")
			So(metrics_new[0].Tags_["test"], ShouldResemble, "test1")
			So(metrics, ShouldNotResemble, metrics_new)

		})
		Convey("Tag on some data, when publisher Tags_ is populated", func() {
			tags := map[string]string{"test1": "test1"}
			for i := range metrics {
				time.Sleep(3)
				rand.Seed(time.Now().UTC().UnixNano())
				data := randInt(65, 90)
				metrics[i] = *plugin.NewMetricType(core.NewNamespace("foo", "bar"), time.Now(), tags, "", data)
			}
			So(metrics[0].Tags_["test1"], ShouldResemble, "test1")
			So(metrics[0].Tags_, ShouldNotBeNil)
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			tagObj := NewTagProcessor()

			_, received_data, _ := tagObj.Process("snap.gob", buf.Bytes(), config)

			var metrics_new []plugin.MetricType

			//Decodes the content into pluginMetricType
			dec := gob.NewDecoder(bytes.NewBuffer(received_data))
			dec.Decode(&metrics_new)
			So(metrics_new[0].Tags_, ShouldNotBeNil)
			So(metrics_new[0].Tags_["test1"], ShouldResemble, "test1")
			So(metrics_new[0].Tags_["rack"], ShouldResemble, "rack1")
			So(metrics_new[0].Tags_["test"], ShouldResemble, "test1")
			So(metrics, ShouldNotResemble, metrics_new)

		})
		Convey("Tag on some data, when publisher Tags_ is populated and tags are overlapping", func() {
			tags := map[string]string{"test": "test2"}
			for i := range metrics {
				time.Sleep(3)
				rand.Seed(time.Now().UTC().UnixNano())
				data := randInt(65, 90)
				metrics[i] = *plugin.NewMetricType(core.NewNamespace("foo", "bar"), time.Now(), tags, "", data)
			}
			So(metrics[0].Tags_["test"], ShouldResemble, "test2")
			So(metrics[0].Tags_, ShouldNotBeNil)
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			tagObj := NewTagProcessor()

			_, received_data, _ := tagObj.Process("snap.gob", buf.Bytes(), config)

			var metrics_new []plugin.MetricType

			//Decodes the content into pluginMetricType
			dec := gob.NewDecoder(bytes.NewBuffer(received_data))
			dec.Decode(&metrics_new)
			So(metrics_new[0].Tags_, ShouldNotBeNil)
			So(metrics_new[0].Tags_["rack"], ShouldResemble, "rack1")
			So(metrics_new[0].Tags_["test"], ShouldResemble, "test1")
			So(metrics, ShouldNotResemble, metrics_new)

		})
		config["tags"] = ctypes.ConfigValueStr{Value: "rack"}
		Convey("Tag on some data, when tag config is broken", func() {
			for i := range metrics {
				time.Sleep(3)
				rand.Seed(time.Now().UTC().UnixNano())
				data := randInt(65, 90)
				metrics[i] = *plugin.NewMetricType(core.NewNamespace("foo", "bar"), time.Now(), nil, "", data)
			}
			So(metrics[0].Tags_, ShouldBeNil)
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			tagObj := NewTagProcessor()

			_, received_data, _ := tagObj.Process("snap.gob", buf.Bytes(), config)

			var metrics_new []plugin.MetricType

			//Decodes the content into pluginMetricType
			dec := gob.NewDecoder(bytes.NewBuffer(received_data))
			dec.Decode(&metrics_new)
			So(metrics_new[0].Tags_, ShouldNotBeNil)
			emptyMap := make(map[string]string)
			So(metrics_new[0].Tags_, ShouldResemble, emptyMap)
			So(metrics, ShouldNotResemble, metrics_new)

		})

	})
}
