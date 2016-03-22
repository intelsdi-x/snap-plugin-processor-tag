[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-processor-tag.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-processor-tag)
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-processor-tag)](http://goreportcard.com/report/intelsdi-x/snap-plugin-processor-tag)
# snap plugin processor - tag
snap plugin intended to process data and add tags to the data. Please notice that not every snap publisher plugin support tagging.

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap] (#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

### Installation
#### Download tag plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [Github Releases](https://github.com/intelsdi-x/snap/releases) page.

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
This builds the plugin in `/build/rootfs`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

## Documentation
[Tag](https://en.wikipedia.org/wiki/Tag_URI_scheme)

### Examples
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
                "/intel/mock/foo": {},
                "/intel/mock/bar": {},
                "/intel/mock/*/baz": {}
            },
            "config": {
                "/intel/mock": {
                    "user": "root",
                    "password": "secret"
                }
            },
            "process": [
                {
                    "plugin_name": "tag",
                    "config":
                        {
                            "tags": "rack1:rack,datacenter:dublin,experiment:1"
                        }
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "influx",
                            "config": {
                               "host": "INFLUXDB_IP",
                               "port": 8086,
                               "database": "snap",
                               "user": "admin",
                               "password": "admin"
                            }
                        }
                    ]
                }
            ]
        }
    }
}
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-processor-tag/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-processor-tag/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [Marcin Spoczynski](https://github.com/sandlbn)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
