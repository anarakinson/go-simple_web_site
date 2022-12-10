module simple_website

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
)

require (
	internal/handlers v1.0.0
	internal/database v1.0.0
	internal/entities v1.0.0
	internal/server v1.0.0
)

replace (
	internal/handlers => ./internal/handlers
 	internal/database => ./internal/database
	internal/entities => ./internal/entities
	internal/server => ./internal/server
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang-migrate/migrate v3.5.4+incompatible // indirect
	github.com/golang-migrate/migrate/v4 v4.15.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.14.0 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956 // indirect
	golang.org/x/text v0.4.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
