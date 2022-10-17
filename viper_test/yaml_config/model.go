package yaml_config

type Config struct {
	Project    string     `json:"project"`
	RunEnv     string     `json:"runEnv"`
	LogLevel   string     `json:"logLevel"`
	RPCServer  RPCServer  `json:"rpcServer"`
	Register   MConfig    `json:"register"`
	Common     MConfig    `json:"common"`
	Postgresql Postgresql `json:"postgresql"`
	Bridge     Bridge     `json:"bridge"`
}

type RPCServer struct {
	AccessUrl     string `json:"accessUrl"`
	DispatcherUrl string `json:"dispatcherUrl"`
}

type MqttConfig struct {
	Id         uint   `json:"id" yaml:"id"`
	URL        string `json:"url" yaml:"url"`
	User       string `json:"user" yaml:"user"`
	Passwd     string `json:"passwd" yaml:"passwd"`
	CA         string `json:"ca" yaml:"ca"`
	Cert       string `json:"cert" yaml:"cert"`
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
}

type MqConfig struct {
	Id        uint   `json:"id" yaml:"id"`
	URL       string `json:"url"`
	QueueName string `json:"queueName"`
	QueueKey  string `json:"queueKey"`
}

type MConfig struct {
	Enable bool       `json:"enable"`
	Mqtt   MqttConfig `json:"mqtt"`
	Mq     MqConfig   `json:"mq"`
}

type Postgresql struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	Dbname   string `json:"Dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Timezone string `json:"timezone"`
}

type Bridge struct {
	Enable bool       `json:"enable"`
	Mqtt   MqttConfig `json:"mqtt"`
}
