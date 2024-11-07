package bot

import "gopkg.in/telebot.v4"

func ExportHandler(b *Bot) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		exportMenu := &telebot.ReplyMarkup{ResizeKeyboard: true}

		xslxBtn := exportMenu.Text("xslx(اکسل)")
		csvBtn := exportMenu.Text("csv")

		exportMenu.Reply(
			exportMenu.Row(xslxBtn, csvBtn),
		)

		return c.Send(" یکی از گزینه های زیر را برای دریافت فایل انتخاب کنید", exportMenu)

	}
}
