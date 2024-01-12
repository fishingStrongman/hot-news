package tengxun

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hotinfo/app/model"
	"io"
	"net/http"
	"time"
)

const api = "https://i.news.qq.com/gw/pc_search/hotWord"

func Run() {
	ticker := time.NewTicker(5 * time.Minute)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			getInfo()
		}
	}
}
func Do(c *gin.Context) {
	getInfo()
}
func getInfo() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
	request.Header.Add("Cookie", "eas_sid=t1X6d9g1A066f0s8w8i1F667X5; pgv_pvid=6565687608; fqm_pvqid=52ef1eb5-8a6f-4225-892a-28f4159266d3; RK=G9dxSYYZMH; ptcz=ab896a6eff7abb6a845f88d684c1e74b359e745f26561f9dcca7ed629e9161bb; qq_domain_video_guid_verify=5cd8cff73065f805; _clck=248b6v|1|fh2|0; pac_uid=0_de0fac31bcdbe; iip=0; _qimei_uuid42=1810909212310040746e10d6df61ae6e145e57681f; _qimei_fingerprint=2ff9e6b26248f3363f4963ad3bc30d9d; _qimei_q36=; _qimei_h38=8f956087746e10d6df61ae6e0200000b818109; lcad_o_minduid=H8mSY6u63HVd_odIIGDWfQlCTzEmfxhK; lcad_appuser=2B14C47E068A38AF; logintype=1; wap_refresh_token=76_eNagPIhJMQbJoHr_tBSuaaRNo_RLuyO9x40Q0gyDNtUs2nP1AHBX3t52KbvgUgP3Ui04I70XDRuoV532jUOVTUfQj4Oar3-5nZJbycsRYHo; wap_encrypt_logininfo=ASuZHXPxJsxaHE13GyDl4zKJmmb%2B8%2BhuFltjqfhXW18%2BaxvOXYTThIvNY%2Fm%2BugBRQO3SHD9H1nCCgEhGFIx%2BoOJ6pTP9ap4JZ%2BOkQnkiKLBX; pgv_info=ssid=s8322238889; ts_last=news.qq.com/topboard.shtml; ts_refer=www.bing.com/; ts_uid=7516076138; lcad_Lturn=577; lcad_LKBturn=976; lcad_LPVLturn=13; lcad_LPLFturn=212; lcad_LPSJturn=618; lcad_LBSturn=267; lcad_LVINturn=435; lcad_LDERturn=747")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	//fmt.Printf("body:%s", string(body))

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 获取所有的 realtime 数据
	realtimeData := response.Hotlist

	// 打印解析后的 realtime 数据
	//fmt.Printf("Realtime data: %+v\n", realtimeData)
	data := make([]*TengXun, 0)
	now := time.Now().Unix()

	for _, list := range realtimeData {

		tmp := TengXun{
			UpdateVer:   now,
			Title:       list.Title,
			Url:         list.ShareUrl,
			Time:        list.Time,
			Source:      list.Source,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
		data = append(data, &tmp)

	}
	model.Conn.Create(data)
}
