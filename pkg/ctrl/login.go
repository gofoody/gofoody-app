package ctrl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofoody/gofoody-app/pkg/config"
	"github.com/gofoody/gofoody-app/templates"
)

type LoginCtrl interface {
	Home(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Signin(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type loginCtrl struct {
	cfg *config.Config
}

func NewLoginCtrl(cfg *config.Config) LoginCtrl {
	return &loginCtrl{cfg: cfg}
}

func (c *loginCtrl) Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (c *loginCtrl) Login(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "templates/login.html", nil)
}

func (c *loginCtrl) Signin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	pword := r.PostFormValue("pword")

	payload := fmt.Sprintf(`{"username":"%s"}`, username)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.cfg.GetAPIGatewayURL(), "/api/login"), strings.NewReader(payload))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	req.Header.Set("Authorization", pword)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func (c *loginCtrl) Signup(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "templates/signup.html", nil)
}

func (c *loginCtrl) Register(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "templates/home.html", nil)
}
