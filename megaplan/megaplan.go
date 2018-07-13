package megaplan

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	authURL              = "/BumsCommonApiV01/User/authorize.api"
	listEventURL         = "/BumsTimeApiV01/Event/list.api"
	listEventCategoryURL = "/BumsTimeApiV01/Event/categories.api"
	cardEmployee         = "/BumsStaffApiV01/Employee/card.api"
)

type Megaplan struct {
	address   string
	accessId  string
	secretKey string
	client    *http.Client
}

type auths struct {
	UserId       int
	EmployeeId   int
	ContractorId string
	AccessId     string
	SecretKey    string
}

type status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Responce struct {
	Status status      `json:"status"`
	Data   interface{} `json:"data"`
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (m *Megaplan) ConnectMegaplan(URL, Login, Password string) error {
	m.address = URL
	params := url.Values{}
	params.Set("Login", Login)
	params.Set("Password", getMD5Hash(Password))
	auth := auths{}
	resp, err := m.do(authURL, params, &auth)
	if err != nil {
		return err
	}
	jresp := resp

	//auth := jresp.Data.(map[string]interface{})
	fmt.Println("jresp: ", jresp.Data)
	m.accessId = auth.AccessId   //["AccessId"].(string)
	m.secretKey = auth.SecretKey //["SecretKey"].(string)
	return nil
}

func (m *Megaplan) ListEvent(TodoListId int, Limit int, Offset int, OnlyActual bool, Search string) []Event {
	param := url.Values{}
	if TodoListId != 0 {
		param.Add("TodoListId", string(TodoListId))
	}
	param.Add("Limit", strconv.Itoa(Limit))
	param.Add("Offset", strconv.Itoa(Offset))
	param.Add("OnlyActual", strconv.FormatBool(OnlyActual))
	events := []Event{}
	_, err := m.do(listEventURL, param, &events)
	//jresp := resp
	if err != nil {
		fmt.Println("err: ", err)
		return events
	}
	//fmt.Println("status", jresp.Status)
	//fmt.Println("jresp.data", events)
	return events
}

func (m *Megaplan) do(URL string, param url.Values, pinterface interface{}) (Responce, error) {
	jresp := Responce{}
	jresp.Data = pinterface
	bodyText := m.dorun(URL, param)
	err := json.Unmarshal(bodyText, &jresp)
	/*fmt.Println("---------------------------------------------------------------")
	fmt.Println(string(bodyText[:]))
	fmt.Println("---------------------------------------------------------------")*/
	if err != nil {
		return jresp, err
	}
	return jresp, nil
}

func (m *Megaplan) dorun(URL string, param url.Values) []byte {
	if m.client == nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		m.client = &http.Client{Transport: tr}
	}
	hostname, _ := url.Parse(m.address)
	quest := hostname.Hostname() + URL
	if param != nil {
		quest += "?" + param.Encode()
	}
	fmt.Println(quest)
	req, err := http.NewRequest("GET", "https://"+quest, nil)
	//fmt.Println("quest: ", quest)

	if err != nil {
		return nil
	}
	current_time := time.Now().Local()
	rfc := current_time.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	req.Header.Add("Date", rfc)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Authorization", (m.accessId + ":" + m.createSignature(quest, rfc)))
	resp, err := m.client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	fmt.Println("-------------------------------------------------------")
	fmt.Println("bodyText: ", string(bodyText))
	fmt.Println("-------------------------------------------------------")
	return bodyText

}

func (m *Megaplan) ListEventCategory() Responce {
	cat := map[string]string{}
	jresp, err := m.do(listEventCategoryURL, nil, &cat)
	if err != nil {
		return jresp
	}
	return jresp
}

func (m *Megaplan) createSignature(URL, Date string) string {
	//hostname, _ := url.Parse(m.address)
	preString := "GET\n\n\n" + Date + "\n" + URL
	//fmt.Println(preString)
	return getSignature(preString, m.secretKey)
}

func getSignature(input, key string) string {
	key_for_sign := []byte(key)
	h := hmac.New(sha1.New, key_for_sign)
	h.Write([]byte(input))
	s := hex.EncodeToString(h.Sum(nil))
	signature := base64.StdEncoding.EncodeToString([]byte(s))
	return signature
}

func (m *Megaplan) GetCardEmployee(id int) Employee {
	param := url.Values{}
	param.Add("Id", string(id))
	employee := Employee{}
	_, err := m.do(cardEmployee, param, &employee)
	//jresp := resp
	if err != nil {
		fmt.Println("err: ", err)
	}
	return employee
}

func Test() {
	Sign := "GET\n\n\nTue, 09 Dec 2014 10:29:11 +0300\nexample.megatest.local/BumsCrmApiV01/Contractor/list.api?FilterId=all&Limit=1&Phone=1"
	SecretKey := "fd57A98113F7Eb562e34F5Fa1c1fDc362dbdE103"
	Hash := "NzQzMGZkMGI1OWYyZTQyNGMzMWVhZTMxMDBiZTk2ODRlMGM3ZTY3NQ=="
	MySign := getSignature(Sign, SecretKey)
	current_time := time.Now().Local()
	rfc := current_time.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	fmt.Println("rfc:   ", rfc)
	fmt.Println("Hash:   ", Hash)
	fmt.Println("MySign: ", MySign)
}
