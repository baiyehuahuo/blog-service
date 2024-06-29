package setting

import "time"

type ServerSettingConfig struct {
	RunMode      string        `yaml:"RunMode"`
	HttpPort     int           `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

type AppSettingConfig struct {
	DefaultPageSize      int      `yaml:"DefaultPageSize"`
	LogSavePath          string   `yaml:"LogSavePath"`
	LogFileName          string   `yaml:"LogFileName"`
	UploadImageMaxSize   int      `yaml:"UploadImageMaxSize"`
	UploadImageAllowExts []string `yaml:"UploadImageAllowExts"`
	MaxPageSize          int      `yaml:"MaxPageSize"`
	LogFileExt           string   `yaml:"LogFileExt"`
	UploadSavePath       string   `yaml:"UploadSavePath"`
	UploadServerURL      string   `yaml:"UploadServerURL"`
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

type JWTSettingConfig struct {
	Secret string        `yaml:"Secret"`
	Issuer string        `yaml:"Issuer"`
	Expire time.Duration `yaml:"Expire"`
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
