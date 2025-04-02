// Copyright  observIQ, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package factories

import (
	"github.com/observiq/bindplane-otel-collector/receiver/awss3eventreceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/awss3rehydrationreceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/azureblobrehydrationreceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/bindplaneauditlogs"
	"github.com/observiq/bindplane-otel-collector/receiver/googlecloudstoragerehydrationreceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/httpreceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/m365receiver"
	"github.com/observiq/bindplane-otel-collector/receiver/oktareceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/pluginreceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/routereceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/sapnetweaverreceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/splunksearchapireceiver"
	"github.com/observiq/bindplane-otel-collector/receiver/telemetrygeneratorreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/bigipreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudflarereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filestatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/netflowreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcplogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udplogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zookeeperreceiver"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/nopreceiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
)

var defaultReceivers = []receiver.Factory{
	activedirectorydsreceiver.NewFactory(),
	bindplaneauditlogs.NewFactory(),
	aerospikereceiver.NewFactory(),
	apachereceiver.NewFactory(),
	apachesparkreceiver.NewFactory(),
	awscloudwatchreceiver.NewFactory(),
	awscontainerinsightreceiver.NewFactory(),
	awsecscontainermetricsreceiver.NewFactory(),
	awsfirehosereceiver.NewFactory(),
	awss3receiver.NewFactory(),
	awss3rehydrationreceiver.NewFactory(),
	awss3eventreceiver.NewFactory(),
	awsxrayreceiver.NewFactory(),
	azureblobreceiver.NewFactory(),
	azureblobrehydrationreceiver.NewFactory(),
	azureeventhubreceiver.NewFactory(),
	bigipreceiver.NewFactory(),
	carbonreceiver.NewFactory(),
	cloudflarereceiver.NewFactory(),
	cloudfoundryreceiver.NewFactory(),
	collectdreceiver.NewFactory(),
	couchdbreceiver.NewFactory(),
	dockerstatsreceiver.NewFactory(),
	elasticsearchreceiver.NewFactory(),
	filelogreceiver.NewFactory(),
	flinkmetricsreceiver.NewFactory(),
	filestatsreceiver.NewFactory(),
	fluentforwardreceiver.NewFactory(),
	googlecloudpubsubreceiver.NewFactory(),
	googlecloudspannerreceiver.NewFactory(),
	googlecloudstoragerehydrationreceiver.NewFactory(),
	hostmetricsreceiver.NewFactory(),
	httpcheckreceiver.NewFactory(),
	httpreceiver.NewFactory(),
	iisreceiver.NewFactory(),
	influxdbreceiver.NewFactory(),
	jaegerreceiver.NewFactory(),
	jmxreceiver.NewFactory(),
	journaldreceiver.NewFactory(),
	k8sclusterreceiver.NewFactory(),
	k8seventsreceiver.NewFactory(),
	kafkametricsreceiver.NewFactory(),
	kafkareceiver.NewFactory(),
	kubeletstatsreceiver.NewFactory(),
	m365receiver.NewFactory(),
	memcachedreceiver.NewFactory(),
	mongodbatlasreceiver.NewFactory(),
	mongodbreceiver.NewFactory(),
	mysqlreceiver.NewFactory(),
	netflowreceiver.NewFactory(),
	nginxreceiver.NewFactory(),
	nopreceiver.NewFactory(),
	oktareceiver.NewFactory(),
	opencensusreceiver.NewFactory(),
	otlpreceiver.NewFactory(),
	pluginreceiver.NewFactory(),
	podmanreceiver.NewFactory(),
	postgresqlreceiver.NewFactory(),
	prometheusreceiver.NewFactory(),
	rabbitmqreceiver.NewFactory(),
	redisreceiver.NewFactory(),
	riakreceiver.NewFactory(),
	routereceiver.NewFactory(),
	saphanareceiver.NewFactory(),
	sapmreceiver.NewFactory(),
	sapnetweaverreceiver.NewFactory(),
	simpleprometheusreceiver.NewFactory(),
	snmpreceiver.NewFactory(),
	splunksearchapireceiver.NewFactory(),
	splunkhecreceiver.NewFactory(),
	sqlqueryreceiver.NewFactory(),
	sqlserverreceiver.NewFactory(),
	statsdreceiver.NewFactory(),
	syslogreceiver.NewFactory(),
	tcplogreceiver.NewFactory(),
	telemetrygeneratorreceiver.NewFactory(),
	udplogreceiver.NewFactory(),
	vcenterreceiver.NewFactory(),
	webhookeventreceiver.NewFactory(),
	windowseventlogreceiver.NewFactory(),
	windowsperfcountersreceiver.NewFactory(),
	zipkinreceiver.NewFactory(),
	zookeeperreceiver.NewFactory(),
}
