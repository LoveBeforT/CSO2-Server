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

func OnRegister() {
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
	http.HandleFunc("/", Register)
	http.HandleFunc("/bg.jpg", OnJpg)
	http.HandleFunc("/background.jpg", OnJpg)
	http.HandleFunc("css/animate.min.css", OnJpg)
	http.HandleFunc("css/bootstrap.min.css", OnJpg)
	http.HandleFunc("css/boxicons.min.css", OnJpg)
	http.HandleFunc("css/fancybox.min.css", OnJpg)
	http.HandleFunc("css/flaticon.css", OnJpg)
	http.HandleFunc("css/magnific-popup.min.css", OnJpg)
	http.HandleFunc("css/meanmenu.min.css", OnJpg)
	http.HandleFunc("css/owl.carousel.min.css", OnJpg)
	http.HandleFunc("css/owl.video.play.html", OnJpg)
	http.HandleFunc("css/responsive.css", OnJpg)
	http.HandleFunc("css/style.css", OnJpg)
	http.HandleFunc("js/ajaxchimp.js", OnJpg)
	http.HandleFunc("js/bootstrap.js", OnJpg)
	http.HandleFunc("js/contact-form-script.js", OnJpg)
	http.HandleFunc("js/cursor.js", OnJpg)
	http.HandleFunc("js/fancybox.js", OnJpg)
	http.HandleFunc("js/form-validator.js", OnJpg)
	http.HandleFunc("js/isotope.pkgd.js", OnJpg)
	http.HandleFunc("js/jquery.js", OnJpg)
	http.HandleFunc("js/magnific-popup.js", OnJpg)
	http.HandleFunc("js/main.js", OnJpg)
	http.HandleFunc("js/meanmenu.js", OnJpg)
	http.HandleFunc("js/owl.carousel.js", OnJpg)
	http.HandleFunc("js/parallax.js", OnJpg)
	http.HandleFunc("js/popper.js", OnJpg)
	http.HandleFunc("js/wow.js", OnJpg)
	http.HandleFunc("img/awards-bg.jpg", OnJpg)
	http.HandleFunc("img/awards-shape1.png", OnJpg)
	http.HandleFunc("img/awards-shape2.png", OnJpg)
	http.HandleFunc("img/bg.jpg", OnJpg)
	http.HandleFunc("img/blog-img1.jpg", OnJpg)
	http.HandleFunc("img/blog-img2.jpg", OnJpg)
	http.HandleFunc("img/blog-img3.jpg", OnJpg)
	http.HandleFunc("img/circle-shape.png", OnJpg)
	http.HandleFunc("img/contact-bg.jpg", OnJpg)
	http.HandleFunc("img/cso2-logo.png", OnJpg)
	http.HandleFunc("img/experts-bg.jpg", OnJpg)
	http.HandleFunc("img/favicon_metro.png", OnJpg)
	http.HandleFunc("img/footer-bg.png", OnJpg)
	http.HandleFunc("img/footer-map.png", OnJpg)
	http.HandleFunc("img/main-banner-bg1.jpg", OnJpg)
	http.HandleFunc("img/main-banner-bg2.jpg", OnJpg)
	http.HandleFunc("img/main-banner-bg3.jpg", OnJpg)
	http.HandleFunc("img/main-banner-bg4.jpg", OnJpg)
	http.HandleFunc("img/main-banner-shape.png", OnJpg)
	http.HandleFunc("img/main-banner-video-bg.jpg", OnJpg)
	http.HandleFunc("img/matches-bg.jpg", OnJpg)
	http.HandleFunc("img/matches-bg1.png", OnJpg)
	http.HandleFunc("img/matches-bg2.png", OnJpg)
	http.HandleFunc("img/valve.png", OnJpg)
	http.HandleFunc("img/nexon.png", OnJpg)
	http.HandleFunc("img/partner-shape.png", OnJpg)
	http.HandleFunc("img/post-bg.jpg", OnJpg)
	http.HandleFunc("img/products-shape.png", OnJpg)
	http.HandleFunc("img/social-shape1.png", OnJpg)
	http.HandleFunc("img/social-shape2.png", OnJpg)
	http.HandleFunc("img/social-shape3.png", OnJpg)
	http.HandleFunc("img/social-shape4.png", OnJpg)
	http.HandleFunc("img/social-shape5.png", OnJpg)
	http.HandleFunc("img/social-shape6.png", OnJpg)
	http.HandleFunc("img/social-shape7.png", OnJpg)
	http.HandleFunc("img/subscribe-bg.jpg", OnJpg)
	http.HandleFunc("event/ch_zh_14112020.png", OnJpg)
	http.HandleFunc("event/eng_14112020.png", OnJpg)
	http.HandleFunc("event/global_14112020.png", OnJpg)
	http.HandleFunc("event/logo.png", OnJpg)
	http.HandleFunc("event/logo2.png", OnJpg)
	http.HandleFunc("event/random_box_09_amethyst2_normal.png", OnJpg)
	http.HandleFunc("event/random_box_17_xmas2016_normal.png", OnJpg)
	http.HandleFunc("event/random_box_mp7phoenix.png", OnJpg)
	http.HandleFunc("event/random_box_new_weapon.png", OnJpg)
	http.HandleFunc("event/random_key_box_01_ui_s.png", OnJpg)
	http.HandleFunc("event/video01.jpg", OnJpg)
	http.HandleFunc("event/video02.jpg", OnJpg)
	http.HandleFunc("fonts/boxicons.eot", OnJpg)
	http.HandleFunc("fonts/boxicons.ttf", OnJpg)
	http.HandleFunc("fonts/arial.ttf", OnJpg)
	http.HandleFunc("fonts/boxicons.ttf", OnJpg)
	http.HandleFunc("fonts/boxicons.woff", OnJpg)
	http.HandleFunc("fonts/boxicons2.woff", OnJpg)
	http.HandleFunc("fonts/boxiconsd41d.svg", OnJpg)
	http.HandleFunc("fonts/Flaticon.eot", OnJpg)
	http.HandleFunc("fonts/Flaticon.svg", OnJpg)
	http.HandleFunc("fonts/Flaticon.ttf", OnJpg)
	http.HandleFunc("fonts/Flaticon.woff", OnJpg)
	http.HandleFunc("fonts/Flaticon.woff", OnJpg)
	http.HandleFunc("fonts/Flaticond41d.eot", OnJpg)
	http.HandleFunc("fonts/Vermin-Vibes-V.eot", OnJpg)
	http.HandleFunc("fonts/Vermin-Vibes-V.svg", OnJpg)
	http.HandleFunc("fonts/Vermin-Vibes-V.woff", OnJpg)
	
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

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path, err := GetExePath()
	if err != nil {
		DebugInfo(2, err)
		return
	}
	t, err := template.ParseFiles(path + "/CSO2-Server/assert/register.html")
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
			MailService.Content = "您的验证码为：" + Vcode + "<br>" + "请勿告诉他人，如非本人操作请忽略本条邮件。请勿回复。"
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

func OnJpg(w http.ResponseWriter, r *http.Request) {
	path, err := GetExePath()
	if err != nil {
		DebugInfo(2, err)
		return
	}
	file, err := os.Open(path + "/CSO2-Server/assert/bg.jpg")
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
