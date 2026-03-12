package types

import (
	"encoding/json"
	"strings"
)

// FlexibleText — поле text в экспорте Telegram может быть строкой или массивом (строки + объекты с полем text).
type FlexibleText string

func (t *FlexibleText) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*t = FlexibleText(s)
		return nil
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	var parts []string
	for _, item := range arr {
		var s string
		if err := json.Unmarshal(item, &s); err == nil {
			parts = append(parts, s)
			continue
		}
		var obj struct {
			Text string `json:"text"`
		}
		if err := json.Unmarshal(item, &obj); err == nil {
			parts = append(parts, obj.Text)
		}
	}
	*t = FlexibleText(strings.Join(parts, ""))
	return nil
}

type TextEntity struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Message struct {
	Id              int          `json:"id"`
	Type            string       `json:"type"`
	Date            string       `json:"date"`
	DateUnixTipe    string       `json:"date_unixtime"`
	From            string       `json:"from,omitempty"`
	FromId          string       `json:"from_id"`
	Text            FlexibleText `json:"text,omitempty"`
	MediaType       string       `json:"media_type,omitempty"`
	StickerEmoji    string       `json:"sticker_emoji,omitempty"`
	DurationSeconds int          `json:"duration_seconds,omitempty"`
	TextEntities    []TextEntity `json:"text_entities,omitempty"`
}

type Chat struct {
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Id       int       `json:"id"`
	Messages []Message `json:"messages"`
}
