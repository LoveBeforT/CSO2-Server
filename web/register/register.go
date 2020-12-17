package register

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct/html"
	. "github.com/KouKouChan/CSO2-Server/configure"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

var (
	mailvcode   = make(map[string]string)
	Reglock     sync.Mutex
	MailService = EmailData{
		"",
		"",
		"",
		"",
		"CSO2-Server",
		"注册验证码",
		"",
	}
)

func OnRegister(path string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Register server suffered a fault !")
			fmt.Println(err)
			fmt.Println("Fault end!")
		}
	}()
	MailService.SenderMail = Conf.REGEmail
	MailService.SenderCode = Conf.REGPassWord
	MailService.SenderSMTP = Conf.REGSMTPaddr
	http.HandleFunc("/", OnMain)
	http.HandleFunc("/download", OnDownload)
	http.HandleFunc("/register", Register)
	fmt.Println("Web is running at", "[AnyAdapter]:"+strconv.Itoa(int(Conf.REGPort)))
	if Conf.EnableMail != 0 {
		fmt.Println("Mail Service is enabled !")
	} else {
		fmt.Println("Mail Service is disabled !")
	}
	err := http.ListenAndServe(":"+strconv.Itoa(int(Conf.REGPort)), nil)
	if err != nil {
		DebugInfo(1, "ListenAndServe:", err)
	}
}

func OnMain(w http.ResponseWriter, r *http.Request) {
	//检查url是否合法
	if strings.Contains(r.URL.Path, "..") {
		DebugInfo(2, "Warning ! Illegal url detected from "+r.RemoteAddr)
		return
	}
	//获取exe目录
	path, err := GetExePath()
	if err != nil {
		DebugInfo(2, err)
		return
	}
	//检索请求url
	web_dir := path + "/CSO2-Server/assert/web"
	if strings.HasPrefix(r.URL.Path, "/img/") ||
		strings.HasPrefix(r.URL.Path, "/images/") ||
		strings.HasPrefix(r.URL.Path, "/css/") ||
		strings.HasPrefix(r.URL.Path, "/js/") ||
		strings.HasPrefix(r.URL.Path, "/fonts/") ||
		strings.HasPrefix(r.URL.Path, "/event/") {
		file := web_dir + r.URL.Path
		f, err := os.Open(file)
		defer f.Close()

		if err != nil && os.IsNotExist(err) {
			DebugInfo(2, "Web file doesn't exist :", r.URL.Path)
			return
		}

		http.ServeFile(w, r, file)
		return
	}
	//发送主页面
	t, err := template.ParseFiles(path + "/CSO2-Server/assert/web/index.html")
	if err != nil {
		DebugInfo(2, err)
		return
	}
	t.Execute(w, WebToHtml{})
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path, err := GetExePath()
	if err != nil {
		DebugInfo(2, err)
		return
	}
	t, err := template.ParseFiles(path + "/CSO2-Server/assert/web/register.html")
	if err != nil {
		DebugInfo(2, err)
		return
	}
	if strings.Join(r.Form["on_click"], ", ") == "sendmail" &&
		Conf.EnableMail != 0 {
		addrtmp := strings.Join(r.Form["emailaddr"], ", ")
		wth := WebToHtml{Addr: addrtmp}
		if addrtmp == "" {
			wth.Tip = MAIL_EMPTY
		} else {
			Vcode := getrand()
			DebugInfo(2, Vcode)
			Reglock.Lock()
			MailService.TargetMail = addrtmp
			MailService.Content = "<table align='center' border='0' cellpadding='0' cellspacing='0' style='border:1px solid #b5b5b5' width='620'><tbody><tr><td style='font-family:Arial,Helvetica,sans-serif;background-color:#ffffff;font-size:12px;color:#4d4d4d;line-height:1.5'><table border='0' cellpadding='0' cellspacing='0' style='border-bottom:1px solid #f2f2f2' width='100%'><tbody><tr><td style='height:54px;padding:50px 0 0 50px;vertical-align:top'><img alt='CSO2' height='50' src='https://www.vippng.com/png/detail/80-806757_counter-strike-online-2-is-the-upcoming-second.png' class='CToWUd'></td></tr></tbody></table><table border='0' cellpadding='0' cellspacing='0' width='100%'><tbody><tr><td style='padding:40px 50px 0;color:#262626;font-family:Arial,Helvetica,sans-serif;font-weight:bold;color:#4d4d4d;font-size:12px'>Counter-Strike Online 2 Doğrulama Kodunuz<br>Your Counter-Strike Online 2 Verification Code<br>网上反恐精英2您的验证码<br>카운터-스트라이크 온라인2인증 코드 :&nbsp;</td></tr><tr><td style='padding:14px 50px 34px;font-family:Arial,Helvetica,sans-serif;font-weight:bold;color:#9fbf00;font-size:28px'>" + Vcode + "" + "</td></tr></tbody></table><table border='0' cellpadding='0' cellspacing='0' width='100%'><tbody><tr><td style='padding:0 50px;word-break:break-all;word-wrap:break-word;font-family:Arial,Helvetica,sans-serif;font-weight:normal;color:#4d4d4d;font-size:12px'><br>[EN]<br>You requested a verification code to create a Counter-Strike Online 2 account.<br><br>[TR]<br>Counter-Strike Online 2 hesabı oluşturmak için bir doğrulama kodu talep ettiniz.<br><br>[KR]<br>카운터-스트라이크 온라인2 계정을 만들기 위해 인증 코드를 요청하셨습니다.<br><br>[China]<br>您已要求输入验证码来创建反恐精英Online 2帐户。<br><br></td></tr><tr><td style='padding:0 50px 45px;font-family:Arial,Helvetica,sans-serif;font-weight:normal;color:#4d4d4d;font-size:12px'>Counter-Strike Online 2 Team</td></tr></tbody></table><table border='0' cellpadding='0' cellspacing='0' style='background-color:#f2f2f2' width='100%'><tbody><tr><td style='padding:25px 0 50px 50px;font-family:Arial,Helvetica,sans-serif;font-weight:normal;color:#4d4d4d;font-size:12px'><br>&nbsp;<a href='https://github.com/KouKouChan' style='color:#005fc1;text-decoration:none' target='_blank' data-saferedirecturl='https://github.com/KouKouChan'>Server Builder: KouKouChan</a></td></tr></tbody></table></td></tr></tbody></table>"
			Reglock.Unlock()
			if SendEmailTO(&MailService) != nil {
				wth.Tip = MAIL_ERROR
			} else {
				wth.Tip = MAIL_SENT

				Reglock.Lock()
				mailvcode[addrtmp] = Vcode
				Reglock.Unlock()
				go TimeOut(addrtmp)
			}
		}
		t.Execute(w, wth)
	} else if strings.Join(r.Form["on_click"], ", ") == "register" &&
		Conf.EnableMail != 0 {
		addrtmp := strings.Join(r.Form["emailaddr"], ", ")
		usernametmp := strings.Join(r.Form["username"], ", ")
		ingamenametmp := strings.Join(r.Form["ingamename"], ", ")
		passwordtmp := strings.Join(r.Form["password"], ", ")
		vercodetmp := strings.Join(r.Form["vercode"], ", ")
		wth := WebToHtml{UserName: usernametmp, Ingamename: ingamenametmp, Password: passwordtmp, Addr: addrtmp, VerCode: vercodetmp}
		if addrtmp == "" {
			wth.Tip = MAIL_EMPTY
			t.Execute(w, wth)
			return
		} else if usernametmp == "" {
			wth.Tip = USERNAME_EMPTY
			t.Execute(w, wth)
			return
		} else if ingamenametmp == "" {
			wth.Tip = GAMENAME_EMPTY
			t.Execute(w, wth)
			return
		} else if passwordtmp == "" {
			wth.Tip = PASSWORD_EMPTY
			t.Execute(w, wth)
			return
		} else if vercodetmp == "" {
			wth.Tip = CODE_EMPTY
			t.Execute(w, wth)
			return
		} else if !check(usernametmp) || !check(ingamenametmp) {
			wth.Tip = NAME_ERROR
			t.Execute(w, wth)
			return
		} else if IsExistsUser([]byte(usernametmp)) {
			wth.Tip = USERNAME_EXISTS
			wth.UserName = ""
			t.Execute(w, wth)
			return
		} else if IsExistsIngameName([]byte(ingamenametmp)) {
			wth.Tip = GAMENAME_EXISTS
			wth.Ingamename = ""
			t.Execute(w, wth)
			return
		} else if mailvcode[addrtmp] == vercodetmp {
			u := GetNewUser()
			u.SetUserName(usernametmp, ingamenametmp)
			u.Password = []byte(fmt.Sprintf("%x", md5.Sum([]byte(usernametmp+passwordtmp))))
			u.UserMail = addrtmp
			if tf := AddUserToDB(&u); tf != nil {
				wth.Tip = DATABASE_ERROR
				t.Execute(w, wth)
				return
			}
			wth.Tip = REGISTER_SUCCESS
			t.Execute(w, wth)
			DebugInfo(1, "User name :<", usernametmp, "> ingamename :<", ingamenametmp, "> mail :<", addrtmp, "> registered !")
		} else {
			wth.Tip = CODE_WRONG
			t.Execute(w, wth)
		}
	} else if strings.Join(r.Form["on_click"], ", ") == "register" &&
		Conf.EnableMail == 0 {
		usernametmp := strings.Join(r.Form["username"], ", ")
		ingamenametmp := strings.Join(r.Form["ingamename"], ", ")
		passwordtmp := strings.Join(r.Form["password"], ", ")
		wth := WebToHtml{UserName: usernametmp, Ingamename: ingamenametmp, Password: passwordtmp}
		if usernametmp == "" {
			wth.Tip = USERNAME_EMPTY
			t.Execute(w, wth)
			return
		} else if ingamenametmp == "" {
			wth.Tip = GAMENAME_EMPTY
			t.Execute(w, wth)
			return
		} else if passwordtmp == "" {
			wth.Tip = PASSWORD_EMPTY
			t.Execute(w, wth)
			return
		} else if !check(usernametmp) || !check(ingamenametmp) {
			wth.Tip = NAME_ERROR
			t.Execute(w, wth)
			return
		} else if IsExistsUser([]byte(usernametmp)) {
			wth.Tip = USERNAME_EXISTS
			wth.UserName = ""
			t.Execute(w, wth)
			return
		} else if IsExistsIngameName([]byte(ingamenametmp)) {
			wth.Tip = GAMENAME_EXISTS
			wth.Ingamename = ""
			t.Execute(w, wth)
			return
		} else {
			u := GetNewUser()
			u.SetUserName(usernametmp, ingamenametmp)
			u.Password = []byte(fmt.Sprintf("%x", md5.Sum([]byte(usernametmp+passwordtmp))))
			u.UserMail = "Unkown"
			if tf := AddUserToDB(&u); tf != nil {
				wth.Tip = DATABASE_ERROR
				t.Execute(w, wth)
				return
			}
			wth.Tip = REGISTER_SUCCESS
			t.Execute(w, wth)
			DebugInfo(1, "User name :<", usernametmp, "> ingamename :<", ingamenametmp, "> registered !")
		}
	} else {
		t.Execute(w, nil)
	}
}

func OnDownload(w http.ResponseWriter, r *http.Request) {
	path, err := GetExePath()
	if err != nil {
		DebugInfo(2, err)
		return
	}
	file, err := os.Open(path + "/CSO2-Server/assert/web/download.html")
	if err != nil {
		DebugInfo(2, err)
		return
	}
	buff, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil {
		DebugInfo(2, err)
		return
	}
	w.Write(buff)
}

func getrand() string {
	rand.Seed(time.Now().Unix())
	randnums := strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10))
	return randnums
}

func TimeOut(addrtmp string) {
	timer := time.NewTimer(time.Minute)
	<-timer.C

	Reglock.Lock()
	defer Reglock.Unlock()
	delete(mailvcode, addrtmp)
}

func check(str string) bool {
	for _, v := range str {
		if v == '.' || v == ' ' || v == '\'' || v == '"' || v == '\\' || v == '/' {
			return false
		}
	}
	return true
}
