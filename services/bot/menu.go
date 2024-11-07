package bot

import "gopkg.in/telebot.v4"

func MenuHandler(b *Bot) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		menu := &telebot.ReplyMarkup{ResizeKeyboard: true}

		filterBtn := menu.Text("جستجو")
		loginBtn := menu.Text("(ادمین)لاگین")
		exportBtn := menu.Text("دریافت فایل")
		bookmarkBtn := menu.Text("لیست علاقه مندی ها")

		menu.Reply(
			menu.Row(filterBtn, loginBtn),
			menu.Row(exportBtn, bookmarkBtn),
		)

		return c.Send("یکی از گزینه های زیر را انتخاب کنید", menu)
	}

}
