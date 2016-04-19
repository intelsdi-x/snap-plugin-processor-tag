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
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"
)

const (
	name       = "tag"
	version    = 2
	pluginType = plugin.ProcessorPluginType
)

// Meta returns a plugin meta data
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(name, version, pluginType, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

// NewTagProcessor returns a TagProcessor struct
func NewTagProcessor() *TagProcessor {
	return &TagProcessor{}
}

// TagProcessor struct
type TagProcessor struct{}

// GetConfigPolicy returns a config policy
func (p *TagProcessor) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	config := cpolicy.NewPolicyNode()
	r1, err := cpolicy.NewStringRule("tags", true)
	handleErr(err)
	r1.Description = "Tags"
	config.Add(r1)
	cp.Add([]string{""}, config)

	return cp, nil
}

func handleErr(err error) {
	if err != nil {
		panic(err)

	}
}

func parseTags(tags string) map[string]string {
	var tagMap map[string]string
	tagMap = make(map[string]string)
	pairs := strings.Split(tags, ",")
	for _, p := range pairs {
		item := strings.Split(p, ":")
		if len(item) == 2 {
			tagMap[item[0]] = item[1]
		}
	}
	return tagMap
}

// Process process metrics
func (p *TagProcessor) Process(contentType string, content []byte, config map[string]ctypes.ConfigValue) (string, []byte, error) {
	logger := log.New()
	logger.Println("Tag processor started")
	var metrics []plugin.PluginMetricType

	dec := gob.NewDecoder(bytes.NewBuffer(content))
	if err := dec.Decode(&metrics); err != nil {
		logger.Printf("Error decoding: error=%v content=%v", err, content)
		return "", nil, err
	}
	if _, ok := config["tags"]; ok {
		tags := config["tags"].(ctypes.ConfigValueStr).Value
		for i := range metrics {
			tagMap := parseTags(tags)
			if len(metrics[i].Tags_) == 0 {
				metrics[i].Tags_ = tagMap
			} else {
				for k, v := range tagMap {
					metrics[i].Tags_[k] = v
				}

			}
		}
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(metrics)
	content = buf.Bytes()

	return contentType, content, nil
}
