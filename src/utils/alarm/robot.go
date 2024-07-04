package alarm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*** 飞书文本告警方法封装 ***/

// 飞书告警文本信息
func PostFeiShu(level, info string, url string) error {
	msg := LarkAlarmTextVariable{
		MsgType: "text",
	}
	msg.Content.Text = fmt.Sprintf("[%s] | %s", level, info)
	marshal, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(marshal))
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

// 飞书告警卡片信息V2
func PostCard2FeiShuV2(url string, cardId string, card_variable interface{}) error {
	var err error
	cardData := Common4CardData{
		Id:       cardId,
		Variable: card_variable,
	}
	CardContent := TranslateCardContent{
		Type: "template",
		Data: cardData,
	}
	cardContent, err := json.Marshal(CardContent)
	if err != nil {
		return err
	}

	msg := FeiShuCardMsg{
		MsgType: "interactive",
		Card:    string(cardContent),
	}
	// fmt.Println("[DEBUG] msg.Card", msg.Card)
	// fmt.Println("[DEBUG] msg", msg)
	marshal, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(marshal))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return err
	}
	// fmt.Println("feishu post card req", body)
	return nil
}
