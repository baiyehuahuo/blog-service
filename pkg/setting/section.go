package setting

type ServerSettingConfig struct {
	RunMode      string `yaml:"RunMode"`
	HttpPort     int    `yaml:"HttpPort"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
}

type AppSettingConfig struct {
	MaxPageSize     int    `yaml:"MaxPageSize"`
	LogSavePath     string `yaml:"LogSavePath"`
	LogFileName     string `yaml:"LogFileName"`
	LogFileExt      string `yaml:"LogFileExt"`
	DefaultPageSize int    `yaml:"DefaultPageSize"`
}

type DatabaseSettingConfig struct {
	DBType       string `yaml:"DBType"`
	Password     string `yaml:"Password"`
	Host         string `yaml:"Host"`
	DBName       string `yaml:"DBName"`
	TablePrefix  string `yaml:"TablePrefix"`
	Username     string `yaml:"Username"`
	Charset      string `yaml:"Charset"`
	ParseTime    bool   `yaml:"ParseTime"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
