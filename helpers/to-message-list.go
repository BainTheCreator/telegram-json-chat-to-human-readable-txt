package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	types "chatreader/types"
)

func getTextFromMessage(msg types.Message) string {
	if len(msg.TextEntities) > 0 {
		var builder strings.Builder

		for _, entity := range msg.TextEntities {
			builder.WriteString(entity.Text)
		}

		return builder.String()
	}

	trimmedText := strings.TrimSpace(string(msg.Text))

	return trimmedText
}

func getMessageContent(msg types.Message) string {
	var mediaType string = msg.MediaType

	switch mediaType {
	case "sticker":
		return msg.StickerEmoji
	case "voice_message":
		return fmt.Sprintf("Голосовое сообщение (%d сек)", msg.DurationSeconds)
	case "video_message":
		return fmt.Sprintf("Видео сообщение (%d сек)", msg.DurationSeconds)
	case "video_note":
		return fmt.Sprintf("Кружок (%d сек)", msg.DurationSeconds)
	case "photo":
		var caption = getTextFromMessage(msg)

		if caption != "" {
			return fmt.Sprintf("Медия(фото): %s", caption)
		}
		return "Медия(фото)"
	case "video_file":
		var caption = getTextFromMessage(msg)
		if caption != "" {
			return fmt.Sprintf("[Медия(видео)]: %s", caption)
		}
		return "[Медия(видео)]"
	case "animation":
	case "gif":
		return "[GIF]"
	case "audio_file":
	case "voice":
		return "[Аудио]"
	case "document":
		return "[Документ]"
	}

	return getTextFromMessage(msg)
}

func toMessage(msg types.Message) string {
	unixTime, err := strconv.ParseInt(msg.DateUnixTipe, 10, 64)
	if err != nil {
		fmt.Println("Ошибка конвертации Unix времени:", err)
	}

	dateObj := time.Unix(unixTime, 0)
	humanDate := dateObj.Format("02.01.2006 15:04:05")

	from := msg.From
	if msg.From == "" {
		from = "Unknown"
	}

	contentWithoutTrim := getMessageContent(msg)
	content := strings.TrimSpace(contentWithoutTrim)

	if content == "" {
		content = "пустое сообщение"
	}

	return fmt.Sprintf("%s %s:\n%s\n", from, humanDate, content)
}

func ToMessageList(msgs []types.Message, options types.ToMessageListOptions) []string {
	var result []string

	for _, msg := range msgs {
		toMessage := toMessage(msg)
		if options.SkipEmpty && toMessage == "" {
			continue
		}
		result = append(result, toMessage)
	}

	return result
}
