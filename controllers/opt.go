package controllers

import (
	r "github.com/cuu/softradius/routers"
	"github.com/cuu/softradius/models"
	"github.com/cuu/softradius/libs"
	"fmt"
	//	"reflect"
	"github.com/astaxie/beego"

	//	sort "github.com/cuu/softradius/libs/sortutil"
	rdb "github.com/cuu/softradius/database/shelf"
	
)

// 操作日志,operate_log
type OperLog struct {
	Id   string `gorethink:"id,omitempty"`
	Name string
	Ip   string
	Time string
	Desc string
}

type AcctOnline struct {
	Id            string `gorethink:"id,omitempty"`
	MemberName    string
	NasAddr       string
	AcctSessionId string
	AcctStartTime string
	FramedIpAddr  string
	MacAddr       string
	NasPortId     string
	BillingTimes  int
	InputTotal    int
	OutputTotal   int
	StartSource   int
}

type AcctTicket struct {
	Id                  string `gorethink:"id,omitempty"`
	MemberName          string
	AcctInputGigawords  int
	AcctOutputGigawords int
	AcctInutOctets      int
	AcctOutputOctets    int
	AcctInputPackets    int
	AcctOutputPackets   int
	AcctSessionId       string
	AcctSessionTime     int
	AcctStartTime       string
	AcctStopTime        string
	AcctTerminateCause  int
	MacAddr             string
	CallingStationId    string
	FrameIdNetmask      string
	FramedIpAddr        string
	NasClass            string
	NasAddr             string
	NasPort             string
	NasPortId           string
	NasPortType         int
	ServiceType         int
	SessionTimeout      int
	StartSource         int
	StopSource          int
}

//维护管理
type OptController struct {
	BaseController
	Page  int
}

var _opt_ctl OptController

func init(){
	_ctl  := &_opt_ctl
	_cate := r.MenuOpt
	
 	_ctl.routes = append( _ctl.routes,
r.Route{Path:"/online",Name:"在线用户",Category:_cate,Is_menu :true, Order:1.2,Is_open:true, Methods:"*:OptOnline"})

 	_ctl.routes = append( _ctl.routes,
		r.Route{Path:"/test",Name:"在线用户管理",Category:_cate,Is_menu :false, Order:1.3,Is_open:true, Methods:"post:TestPost"})
	
	_ctl.routes = append( _ctl.routes,
		r.Route{Path:"/operlog", Name:"操作日志",Category:_cate,Is_menu:true, Order:1.4,Is_open:true, Methods:"*:OperLog"})
	
	_ctl.AddRoutes()
	
}

func (this *OptController) AddRoutes() {
	
	for i,v := range  this.routes {
		if v.Methods != "" {
			beego.Router(v.Path, this, v.Methods)
		}else{
			beego.Router(v.Path, this)
		}
		
		//Permits.routes is map,so no confict path key ! 
		r.Permits.Add_route(v.Path,&this.routes[i])
	}
	
}

func (this *OptController) GuuPrepare(){
	
	this.TplName = libs.GetTplName(this)
	
	//mainly init form
	tst := [][]string{{"a","b","c"},{"d","e"}}
	
	fmt.Println(tst)
	
	this.Forms["test"] = models.InfoForm("haha","/test",models.TextBox(&models.Input{Name:"username",Id:"username",Description:"Username",Required:true,Valid:models.Is_email}),
		models.Password(&models.Input{Name:"password",Id:"pwd_input",Description:"Password",Required:true}),
		models.Dropdown(&models.Select{Name:"nodes",Id:"user_node",Description:"Chose Node",Args:tst}),
		models.GroupDropdown(&models.Select{Name:"nodes",Id:"user_node",Description:"Privilige",Args:tst}),
		models.CheckBox(&models.Input{Name:"rmb",Description:"Remember me",Class:"guucheckbox"}),
		models.TextArea(&models.Input{Name:"note",Description:"Note"}),
		models.Submit(&models.Input{Name:"submit",Description:"Submit",Value:"Submit",Class:"btn btn-info"}))
	
	fmt.Println(libs.GetTplName(this))
	
}

//默认的根route 显示
func (this *OptController) TestGet() {
	this.Data["Form"] = this.Forms["test"].Render()

	this.Render()
}


func (this *OptController) TestPost(){

	tst := this.GetString("password")
	
//	this.Data["Form"] = models.AlertBox(&models.Alert{Title:"Test post",Type:"Info",Msg:fmt.Sprintf("a test msg from post %s",tst)});

	if this.Validator("test") == false{
		this.Data["Form"] = this.Forms["test"].Render()
		
	}else
	{
		this.Data["Form"] = models.AlertBox(&models.Alert{Title:"Test post",Type:"Info",Msg:fmt.Sprintf("a test msg from post %s",tst)});
	}

	this.Render()

}

func (this *OptController) Test(){
	this.Data["Form"] = "只是一个测试不同route的结果"

	this.Render()
}


func (this *OptController) OnlineUsers() []AcctOnline {
	var nods []AcctOnline
	rdb.DataBase().SkipGet2(&nods,0,1000)
	return nods
}

func (this *OptController) OptOnline() {
	this.TplName = "opt_online_list.html"
	this.Data["BasList"] = this.BasList()
	this.Data["NodeList"] = this.NodeList()
	this.Data["OnlineUsers"] = this.OnlineUsers()

	this.Data["FmtOnlineTime"] = libs.FmtOnlineTime
	
	this.Render()
}


func (this *OptController) OperLog() {
	var logs []OperLog
	rdb.DataBase().SkipGet2(&logs,0,1000)
	
	this.TplName = "opr_log_list.html"
	
	
	//this.Data[""] = logs

	this.Render()
	
}
