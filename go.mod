module personal

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Unknwon/com v0.0.0-20181010210213-41959bdd855f
	github.com/Unknwon/i18n v0.0.0-20171114194641-b64d33658966 // indirect
	github.com/go-macaron/binding v0.0.0-20170611065819-ac54ee249c27
	github.com/go-macaron/csrf v0.0.0-20180914142414-18bf065728b4 // indirect
	github.com/go-macaron/i18n v0.0.0-20160612092837-ef57533c3b0f
	github.com/go-macaron/inject v0.0.0-20160627170012-d8a0b8677191
	github.com/go-macaron/renders v0.0.0-20151026045949-c889aa4911a1
	github.com/go-macaron/session v0.0.0-20181205003248-092353907b3c // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1 // indirect
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/oxtoacart/bpool v0.0.0-20150712133111-4e1c5567d7c2 // indirect
	github.com/rs/cors v1.6.0 // indirect
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d // indirect
	github.com/smartystreets/goconvey v0.0.0-20180222194500-ef6db91d284a // indirect
	golang.org/x/crypto v0.0.0-20181112202954-3d3f9f413869
	golang.org/x/text v0.3.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/ini.v1 v1.39.0
	gopkg.in/macaron.v1 v1.3.1
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/yaml.v2 v2.2.1 // indirect
)

replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
