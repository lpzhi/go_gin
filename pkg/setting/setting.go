package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {

	PageSize int
	JwtSecret string

	ExportSavePath string
	PrefixUrl string
	RuntimeRootPath string
}


var (
	Cfg *ini.File

	//RunMode string
	//
	//HTTPPort int
	//ReadTimeout time.Duration
	//WriteTimeout time.Duration
	//
	//PageSize int
	//JwtSecret string
	//
	//ExportSavePath string
	//PrefixUrl string
)

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func init() {
	var err error
	Cfg, err = ini.Load("gin/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	SetUp()
	//LoadBase()
	//LoadServer()
	//LoadApp()
}


func SetUp()  {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("server",ServerSetting)

	ServerSetting.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	ServerSetting.WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}


//func LoadApp() {
//	sec, err := Cfg.GetSection("app")
//	if err != nil {
//		log.Fatalf("Fail to get section 'app': %v", err)
//	}
//
//	AppSetting.JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
//	AppSetting.PageSize = sec.Key("PAGE_SIZE").MustInt(10)
//}

func mapTo(section string,v interface{})  {
	err := Cfg.Section(section).MapTo(v)

	if err !=nil{
		log.Fatal("Cfg.MapTo err %v",err)
	}
}