module github.com/olatheander/ksync-syncthing-test

go 1.13

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/aws/aws-sdk-go v1.23.17 // indirect
	github.com/blang/semver v3.5.1+incompatible
	github.com/briandowns/spinner v1.6.1
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustinkirkland/golang-petname v0.0.0-20190613200456-11339a705ed2
	github.com/fatih/structs v1.1.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-resty/resty v1.12.0
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/imdario/mergo v0.3.7
	github.com/jpillora/overseer v0.0.0-20190427034852-ce9055846616
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/logrusorgru/aurora v0.0.0-20190803045625-94edacc10f9b
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/olekukonko/tablewriter v0.0.1
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/phayes/permbits v0.0.0-20190612203442-39d7c581d2ee
	github.com/pkg/errors v0.8.1
	github.com/sevlyar/go-daemon v0.1.5
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/syncthing/syncthing v1.2.2
	github.com/timfallmk/go-daemon v0.1.2
	github.com/timfallmk/overseer v0.0.0-20171218213912-3f6a00445e03
	github.com/vapor-ware/ksync v0.0.0-20190724201530-c5fd87d3634b
	golang.org/x/crypto v0.0.0-20190907121410-71b5226ff739
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297
	google.golang.org/grpc v1.23.0
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190905160310-fb749d2f1064
	k8s.io/apimachinery v0.0.0-20190831074630-461753078381
	k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
	k8s.io/utils v0.0.0-20190907131718-3d4f5b7dea0b // indirect
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
