/*  Copyright (C) 2020 by Anandpskerala@Github, < https://github.com/Anandpskerala >.
 *
 * This file is part of < https://github.com/Anandpskerala/ForwardTagRemoverBot > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/Anandpskerala/ForwardTagRemoverBot/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package commands

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"go.uber.org/zap"
)

func Start(b ext.Bot, u *gotgbot.Update) error {

	startButton := [][]ext.InlineKeyboardButton{make([]ext.InlineKeyboardButton, 2), make([]ext.InlineKeyboardButton, 1)}

	startButton[0][0] = ext.InlineKeyboardButton{
		Text: "𝗨𝗽𝗱𝗮𝘁𝗲𝘀 𝗖𝗵𝗮𝗻𝗻𝗲𝗹 🔖",
		Url:  "https://t.me/Mega_Bots_Updates",
	}

	startButton[0][1] = ext.InlineKeyboardButton{
		Text: "𝗦𝘂𝗽𝗽𝗼𝗿𝘁 𝗚𝗿𝗼𝘂𝗽 📌",
		Url:  "https://t.me/Mega_Bots_Supporters",
	}

	startButton[1][0] = ext.InlineKeyboardButton{
		Text: "My Owner🔘",
		Url:  "https://t.me/Shivaay_12",
	}

	markup := ext.InlineKeyboardMarkup{InlineKeyboard: &startButton}

	msg := b.NewSendableMessage(u.EffectiveChat.Id, fmt.Sprintf("Hi [%s](tg://user?id=%v)\nI am A Forward Tag remover Bot.Send /help To Know What I Can Do", u.EffectiveUser.FirstName, u.EffectiveUser.Id))
	msg.ReplyToMessageId = u.EffectiveMessage.MessageId
	msg.ReplyMarkup = &markup
	msg.ParseMode = parsemode.Markdown
	_, err := msg.Send()
	if err != nil {
		b.Logger.Warnw("Error in sending", zap.Error(err))
	}
	return err
}
