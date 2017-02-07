[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-processor-tag.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-processor-tag)
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-processor-tag)](http://goreportcard.com/report/intelsdi-x/snap-plugin-processor-tag)
# Snap plugin processor - tag
Snap plugin intended to process data and add tags to the data.

Please notice that tags can be added also through configuration in task manifest (for details see [documentation](https://github.com/intelsdi-x/snap/blob/master/docs/TASKS.md)) and
not every Snap publisher plugin support tagging.

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

### Installation
#### Download tag plugin binary:
You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-processor-tag/releases) page. Download the plugin from the latest release and load it into `snapteld` (`/opt/snap/plugins` is the default location for Snap packages).

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-processor-tag

Clone repo into `$GOPATH/src/github/intelsdi-x/`:
```
$ git clone https://github.com/<yourGithubID>/snap-plugin-processor-tag
```
Build the plugin by running make in repo:
```
$ make
```
This builds the plugin in `./build`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap#getting-started)

## Documentation
[Tag](https://en.wikipedia.org/wiki/Tag_URI_scheme)

### Examples
Example running psutil plugin, tag processor, and writing data into a file.

Documentation for Snap collector psutil plugin can be found [here](https://github.com/intelsdi-x/snap-plugin-collector-psutil)

In one terminal window, open the Snap daemon :
```
$ snapteld -t 0 -l 1
```
The option "-l 1" it is for setting the debugging log level and "-t 0" is for disabling plugin signing.

In another terminal window:

Download and load collector, processor and publisher plugins
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-psutil/latest/linux/x86_64/snap-plugin-collector-psutil
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-processor-tag/latest/linux/x86_64/snap-plugin-processor-tag
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snaptel plugin load snap-plugin-collector-psutil
$ snaptel plugin load snap-plugin-publisher-file
$ snaptel plugin load snap-plugin-processor-tag
```

See available metrics for your system
```
$ snaptel metric list
```

Create a task file. For example, psutil-tag-file.json:
```
{
  "version": 1,
  "schedule": {
    "type": "simple",
    "interval": "1s"
  },
  "workflow": {
    "collect": {
      "metrics": {
        "/intel/psutil/load/load1": {},
        "/intel/psutil/load/load5": {},
        "/intel/psutil/load/load15": {},
        "/intel/psutil/vm/free": {},
        "/intel/psutil/vm/used": {}
      },
      "process": [
        {
          "plugin_name": "tag",
          "config": {
            "tags": "rack1:rack,datacenter:dublin,experiment:1"
          },
          "publish": [
            {
              "plugin_name": "file",
              "config": {
                "file": "/tmp/published_tag.log"
              }
            }
          ]
        }
      ]
    }
  }
}
```
Start task:
```
$ snaptel task create -t psutil-tag-file.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See realtime output from `snaptel task watch <task_id>` (CTRL+C to exit)
```
snaptel task watch 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

This data is published to a file `/tmp/published_tag.log` per task specification

Stop task:
```
$ snaptel task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-processor-tag/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-processor-tag/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [Marcin Spoczynski](https://github.com/sandlbn)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.

This software has been contributed by MIKELANGELO, a Horizon 2020 project co-funded by the European Union. https://www.mikelangelo-project.eu/
