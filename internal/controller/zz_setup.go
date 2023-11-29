/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	api "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/api"
	apikey "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/apikey"
	apikeyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/apikeyattachment"
	customdomain "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/customdomain"
	ipstrategy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/ipstrategy"
	service "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/service"
	servicerelease "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/servicerelease"
	strategyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/strategyattachment"
	usageplan "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/usageplan"
	usageplanattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/apigateway/usageplanattachment"
	attachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/attachment"
	lifecyclehook "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/lifecyclehook"
	notification "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/notification"
	scalingconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/scalingconfig"
	scalinggroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/scalinggroup"
	scalingpolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/scalingpolicy"
	schedule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/schedule"
	audit "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/audit/audit"
	group "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/group"
	groupmembership "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/groupmembership"
	grouppolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/grouppolicyattachment"
	oidcsso "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/oidcsso"
	policy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/policy"
	role "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/role"
	rolepolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/rolepolicyattachment"
	rolesso "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/rolesso"
	samlprovider "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/samlprovider"
	user "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/user"
	userpolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/userpolicyattachment"
	taskset "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cat/taskset"
	snapshot "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/snapshot"
	snapshotpolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/snapshotpolicy"
	snapshotpolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/snapshotpolicyattachment"
	storage "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/storage"
	storageattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/storageattachment"
	storageset "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/storageset"
	attachmentccn "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ccn/attachment"
	bandwidthlimit "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ccn/bandwidthlimit"
	ccn "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ccn/ccn"
	instance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cdh/instance"
	domain "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cdn/domain"
	urlpurge "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cdn/urlpurge"
	urlpush "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cdn/urlpush"
	accessgroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cfs/accessgroup"
	accessrule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cfs/accessrule"
	filesystem "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cfs/filesystem"
	albserverattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/albserverattachment"
	attachmentclb "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/attachment"
	customizedconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/customizedconfig"
	instanceclb "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/instance"
	lb "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/lb"
	listener "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/listener"
	listenerrule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/listenerrule"
	logset "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/logset"
	logtopic "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/logtopic"
	redirection "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/redirection"
	snatip "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/snatip"
	targetgroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/targetgroup"
	targetgroupattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/targetgroupattachment"
	targetgroupinstanceattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/targetgroupinstanceattachment"
	config "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cls/config"
	configattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cls/configattachment"
	configextra "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cls/configextra"
	cosshipper "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cls/cosshipper"
	index "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cls/index"
	logsetcls "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cls/logset"
	machinegroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cls/machinegroup"
	topic "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cls/topic"
	containercluster "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/containercluster/containercluster"
	containerclusterinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/containercluster/containerclusterinstance"
	bucket "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cos/bucket"
	bucketdomaincertificateattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cos/bucketdomaincertificateattachment"
	bucketobject "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cos/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cos/bucketpolicy"
	image "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/image"
	instancecvm "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/instance"
	instanceset "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/instanceset"
	keypair "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/keypair"
	placementgroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/placementgroup"
	reservedinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/reservedinstance"
	cluster "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cynosdb/cluster"
	readonlyinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cynosdb/readonlyinstance"
	cchttppolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/cchttppolicy"
	cchttpspolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/cchttpspolicy"
	ccpolicyv2 "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/ccpolicyv2"
	ddospolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/ddospolicy"
	ddospolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/ddospolicyattachment"
	ddospolicycase "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/ddospolicycase"
	ddospolicyv2 "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/ddospolicyv2"
	eip "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/eip"
	l4rule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/l4rule"
	l4rulev2 "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/l4rulev2"
	l7rule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/l7rule"
	l7rulev2 "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dayu/l7rulev2"
	dcx "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dc/dcx"
	gateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dc/gateway"
	gatewayccnroute "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dc/gatewayccnroute"
	dcdbaccount "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dcdb/dcdbaccount"
	hourdbinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dcdb/hourdbinstance"
	securitygroupattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dcdb/securitygroupattachment"
	domaininstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dnspod/domaininstance"
	record "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/dnspod/record"
	instanceelasticsearch "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/elasticsearch/instance"
	emrcluster "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/emr/emrcluster"
	eni "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/eni/eni"
	eniattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/eni/eniattachment"
	certificate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/certificate"
	httpdomain "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/httpdomain"
	httprule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/httprule"
	layer4listener "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/layer4listener"
	layer7listener "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/layer7listener"
	proxy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/proxy"
	realserver "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/realserver"
	securitypolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/securitypolicy"
	securityrule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/gaap/securityrule"
	acl "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/kafka/acl"
	instancekafka "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/kafka/instance"
	topickafka "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/kafka/topic"
	userkafka "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/kafka/user"
	externalkey "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/kms/externalkey"
	key "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/kms/key"
	instancelighthouse "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/lighthouse/instance"
	dedicatedclusterdbinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mariadb/dedicatedclusterdbinstance"
	hourdbinstancemariadb "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mariadb/hourdbinstance"
	logfileretentionperiod "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mariadb/logfileretentionperiod"
	mariadbaccount "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mariadb/mariadbaccount"
	parameters "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mariadb/parameters"
	securitygroups "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mariadb/securitygroups"
	instancemongodb "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mongodb/instance"
	shardinginstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mongodb/shardinginstance"
	standbyinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mongodb/standbyinstance"
	alarmnotice "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/alarmnotice"
	alarmpolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/alarmpolicy"
	bindingreceiver "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/bindingreceiver"
	grafanainstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/grafanainstance"
	grafanaintegration "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/grafanaintegration"
	grafananotificationchannel "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/grafananotificationchannel"
	grafanaplugin "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/grafanaplugin"
	grafanassoaccount "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/grafanassoaccount"
	policybindingobject "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/policybindingobject"
	policygroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/policygroup"
	tmpalertrule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmpalertrule"
	tmpcvmagent "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmpcvmagent"
	tmpexporterintegration "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmpexporterintegration"
	tmpinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmpinstance"
	tmprecordingrule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmprecordingrule"
	tmpscrapejob "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmpscrapejob"
	tmptkealertpolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmptkealertpolicy"
	tmptkeclusteragent "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmptkeclusteragent"
	tmptkeconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmptkeconfig"
	tmptkeglobalnotification "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmptkeglobalnotification"
	tmptkerecordruleyaml "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmptkerecordruleyaml"
	tmptketemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmptketemplate"
	tmptketemplateattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/monitor/tmptketemplateattachment"
	account "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mysql/account"
	accountprivilege "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mysql/accountprivilege"
	backuppolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mysql/backuppolicy"
	instancemysql "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mysql/instance"
	privilege "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mysql/privilege"
	readonlyinstancemysql "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/mysql/readonlyinstance"
	instancepostgresql "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/postgresql/instance"
	readonlyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/postgresql/readonlyattachment"
	readonlygroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/postgresql/readonlygroup"
	readonlyinstancepostgresql "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/postgresql/readonlyinstance"
	recordprivatedns "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/privatedns/record"
	zone "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/privatedns/zone"
	providerconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/providerconfig"
	backupconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/redis/backupconfig"
	instanceredis "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/redis/instance"
	function "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/scf/function"
	layer "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/scf/layer"
	namespace "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/scf/namespace"
	domainses "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ses/domain"
	emailaddress "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ses/emailaddress"
	template "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ses/template"
	sign "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sms/sign"
	templatesms "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sms/template"
	accountsqlserver "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sqlserver/account"
	accountdbattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sqlserver/accountdbattachment"
	basicinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sqlserver/basicinstance"
	db "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sqlserver/db"
	instancesqlserver "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sqlserver/instance"
	publishsubscribe "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sqlserver/publishsubscribe"
	readonlyinstancesqlserver "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/sqlserver/readonlyinstance"
	certificatessl "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ssl/certificate"
	freecertificate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ssl/freecertificate"
	paycertificate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ssl/paycertificate"
	secret "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ssm/secret"
	secretversion "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ssm/secretversion"
	clustertcaplus "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcaplus/cluster"
	idl "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcaplus/idl"
	table "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcaplus/table"
	tablegroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcaplus/tablegroup"
	clusterattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcm/clusterattachment"
	mesh "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcm/mesh"
	instancetcr "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/instance"
	namespacetcr "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/namespace"
	repository "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/repository"
	token "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/token"
	vpcattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/vpcattachment"
	instancetdmq "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tdmq/instance"
	namespacetdmq "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tdmq/namespace"
	namespaceroleattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tdmq/namespaceroleattachment"
	roletdmq "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tdmq/role"
	topictdmq "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tdmq/topic"
	appconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tem/appconfig"
	application "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tem/application"
	environment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tem/environment"
	gatewaytem "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tem/gateway"
	logconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tem/logconfig"
	scalerule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tem/scalerule"
	workload "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tem/workload"
	applicationproxy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/applicationproxy"
	applicationproxyrule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/applicationproxyrule"
	ddospolicyteo "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/ddospolicy"
	dnsrecord "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/dnsrecord"
	dnssec "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/dnssec"
	origingroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/origingroup"
	origingrouprity "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/origingrouprity"
	ruleengine "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/ruleengine"
	ruleenginepriority "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/ruleenginepriority"
	securitypolicyteo "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/securitypolicy"
	zoneteo "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/zone"
	zonesetting "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/teo/zonesetting"
	addonattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tke/addonattachment"
	authattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tke/authattachment"
	clustertke "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tke/cluster"
	clusterattachmenttke "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tke/clusterattachment"
	clusterendpoint "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tke/clusterendpoint"
	nodepool "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tke/nodepool"
	scaleworker "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tke/scaleworker"
	adaptivedynamicstreamingtemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/adaptivedynamicstreamingtemplate"
	imagespritetemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/imagespritetemplate"
	proceduretemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/proceduretemplate"
	snapshotbytimeoffsettemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/snapshotbytimeoffsettemplate"
	subapplication "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/subapplication"
	superplayerconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/superplayerconfig"
	addresstemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/addresstemplate"
	addresstemplategroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/addresstemplategroup"
	dnat "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/dnat"
	eipvpc "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/eip"
	eipassociation "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/eipassociation"
	havip "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/havip"
	natgateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/natgateway"
	natgatewaysnat "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/natgatewaysnat"
	protocoltemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/protocoltemplate"
	protocoltemplategroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/protocoltemplategroup"
	routeentry "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routeentry"
	routetable "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routetable"
	routetableentry "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routetableentry"
	securitygroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygroup"
	securitygroupliterule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygroupliterule"
	securitygrouprule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygrouprule"
	subnet "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/subnet"
	vpc "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpc"
	vpcacl "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcacl"
	vpcbandwidthpackage "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcbandwidthpackage"
	vpcbandwidthpackageattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcbandwidthpackageattachment"
	vpnconnection "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnconnection"
	vpncustomergateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpncustomergateway"
	vpngateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpngateway"
	vpngatewayroute "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpngatewayroute"
	vpnsslclient "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnsslclient"
	vpnsslserver "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnsslserver"
	vpvaclattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpvaclattachment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.Setup,
		apikey.Setup,
		apikeyattachment.Setup,
		customdomain.Setup,
		ipstrategy.Setup,
		service.Setup,
		servicerelease.Setup,
		strategyattachment.Setup,
		usageplan.Setup,
		usageplanattachment.Setup,
		attachment.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		scalingconfig.Setup,
		scalinggroup.Setup,
		scalingpolicy.Setup,
		schedule.Setup,
		audit.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		oidcsso.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		rolesso.Setup,
		samlprovider.Setup,
		user.Setup,
		userpolicyattachment.Setup,
		taskset.Setup,
		snapshot.Setup,
		snapshotpolicy.Setup,
		snapshotpolicyattachment.Setup,
		storage.Setup,
		storageattachment.Setup,
		storageset.Setup,
		attachmentccn.Setup,
		bandwidthlimit.Setup,
		ccn.Setup,
		instance.Setup,
		domain.Setup,
		urlpurge.Setup,
		urlpush.Setup,
		accessgroup.Setup,
		accessrule.Setup,
		filesystem.Setup,
		albserverattachment.Setup,
		attachmentclb.Setup,
		customizedconfig.Setup,
		instanceclb.Setup,
		lb.Setup,
		listener.Setup,
		listenerrule.Setup,
		logset.Setup,
		logtopic.Setup,
		redirection.Setup,
		snatip.Setup,
		targetgroup.Setup,
		targetgroupattachment.Setup,
		targetgroupinstanceattachment.Setup,
		config.Setup,
		configattachment.Setup,
		configextra.Setup,
		cosshipper.Setup,
		index.Setup,
		logsetcls.Setup,
		machinegroup.Setup,
		topic.Setup,
		containercluster.Setup,
		containerclusterinstance.Setup,
		bucket.Setup,
		bucketdomaincertificateattachment.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		image.Setup,
		instancecvm.Setup,
		instanceset.Setup,
		keypair.Setup,
		placementgroup.Setup,
		reservedinstance.Setup,
		cluster.Setup,
		readonlyinstance.Setup,
		cchttppolicy.Setup,
		cchttpspolicy.Setup,
		ccpolicyv2.Setup,
		ddospolicy.Setup,
		ddospolicyattachment.Setup,
		ddospolicycase.Setup,
		ddospolicyv2.Setup,
		eip.Setup,
		l4rule.Setup,
		l4rulev2.Setup,
		l7rule.Setup,
		l7rulev2.Setup,
		dcx.Setup,
		gateway.Setup,
		gatewayccnroute.Setup,
		dcdbaccount.Setup,
		hourdbinstance.Setup,
		securitygroupattachment.Setup,
		domaininstance.Setup,
		record.Setup,
		instanceelasticsearch.Setup,
		emrcluster.Setup,
		eni.Setup,
		eniattachment.Setup,
		certificate.Setup,
		httpdomain.Setup,
		httprule.Setup,
		layer4listener.Setup,
		layer7listener.Setup,
		proxy.Setup,
		realserver.Setup,
		securitypolicy.Setup,
		securityrule.Setup,
		acl.Setup,
		instancekafka.Setup,
		topickafka.Setup,
		userkafka.Setup,
		externalkey.Setup,
		key.Setup,
		instancelighthouse.Setup,
		dedicatedclusterdbinstance.Setup,
		hourdbinstancemariadb.Setup,
		logfileretentionperiod.Setup,
		mariadbaccount.Setup,
		parameters.Setup,
		securitygroups.Setup,
		instancemongodb.Setup,
		shardinginstance.Setup,
		standbyinstance.Setup,
		alarmnotice.Setup,
		alarmpolicy.Setup,
		bindingreceiver.Setup,
		grafanainstance.Setup,
		grafanaintegration.Setup,
		grafananotificationchannel.Setup,
		grafanaplugin.Setup,
		grafanassoaccount.Setup,
		policybindingobject.Setup,
		policygroup.Setup,
		tmpalertrule.Setup,
		tmpcvmagent.Setup,
		tmpexporterintegration.Setup,
		tmpinstance.Setup,
		tmprecordingrule.Setup,
		tmpscrapejob.Setup,
		tmptkealertpolicy.Setup,
		tmptkeclusteragent.Setup,
		tmptkeconfig.Setup,
		tmptkeglobalnotification.Setup,
		tmptkerecordruleyaml.Setup,
		tmptketemplate.Setup,
		tmptketemplateattachment.Setup,
		account.Setup,
		accountprivilege.Setup,
		backuppolicy.Setup,
		instancemysql.Setup,
		privilege.Setup,
		readonlyinstancemysql.Setup,
		instancepostgresql.Setup,
		readonlyattachment.Setup,
		readonlygroup.Setup,
		readonlyinstancepostgresql.Setup,
		recordprivatedns.Setup,
		zone.Setup,
		providerconfig.Setup,
		backupconfig.Setup,
		instanceredis.Setup,
		function.Setup,
		layer.Setup,
		namespace.Setup,
		domainses.Setup,
		emailaddress.Setup,
		template.Setup,
		sign.Setup,
		templatesms.Setup,
		accountsqlserver.Setup,
		accountdbattachment.Setup,
		basicinstance.Setup,
		db.Setup,
		instancesqlserver.Setup,
		publishsubscribe.Setup,
		readonlyinstancesqlserver.Setup,
		certificatessl.Setup,
		freecertificate.Setup,
		paycertificate.Setup,
		secret.Setup,
		secretversion.Setup,
		clustertcaplus.Setup,
		idl.Setup,
		table.Setup,
		tablegroup.Setup,
		clusterattachment.Setup,
		mesh.Setup,
		instancetcr.Setup,
		namespacetcr.Setup,
		repository.Setup,
		token.Setup,
		vpcattachment.Setup,
		instancetdmq.Setup,
		namespacetdmq.Setup,
		namespaceroleattachment.Setup,
		roletdmq.Setup,
		topictdmq.Setup,
		appconfig.Setup,
		application.Setup,
		environment.Setup,
		gatewaytem.Setup,
		logconfig.Setup,
		scalerule.Setup,
		workload.Setup,
		applicationproxy.Setup,
		applicationproxyrule.Setup,
		ddospolicyteo.Setup,
		dnsrecord.Setup,
		dnssec.Setup,
		origingroup.Setup,
		origingrouprity.Setup,
		ruleengine.Setup,
		ruleenginepriority.Setup,
		securitypolicyteo.Setup,
		zoneteo.Setup,
		zonesetting.Setup,
		addonattachment.Setup,
		authattachment.Setup,
		clustertke.Setup,
		clusterattachmenttke.Setup,
		clusterendpoint.Setup,
		nodepool.Setup,
		scaleworker.Setup,
		adaptivedynamicstreamingtemplate.Setup,
		imagespritetemplate.Setup,
		proceduretemplate.Setup,
		snapshotbytimeoffsettemplate.Setup,
		subapplication.Setup,
		superplayerconfig.Setup,
		addresstemplate.Setup,
		addresstemplategroup.Setup,
		dnat.Setup,
		eipvpc.Setup,
		eipassociation.Setup,
		havip.Setup,
		natgateway.Setup,
		natgatewaysnat.Setup,
		protocoltemplate.Setup,
		protocoltemplategroup.Setup,
		routeentry.Setup,
		routetable.Setup,
		routetableentry.Setup,
		securitygroup.Setup,
		securitygroupliterule.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		vpc.Setup,
		vpcacl.Setup,
		vpcbandwidthpackage.Setup,
		vpcbandwidthpackageattachment.Setup,
		vpnconnection.Setup,
		vpncustomergateway.Setup,
		vpngateway.Setup,
		vpngatewayroute.Setup,
		vpnsslclient.Setup,
		vpnsslserver.Setup,
		vpvaclattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
