package config

// Bot
// Main Menu
const (
	MainMenu                = "منوی اصلی"
	SearchButton            = "جستجو"
	FiltersButton           = "فیلتر ها"
	FavoritesButton         = "لیست علاقه‌مندی ها"
	ExportButton            = "خروجی فایل"
	AccountManagementButton = "مدیریت حساب"
	SupportButton           = "تماس با ما"
	BackButton              = "بازگشت"
)

//export buttons
const (
	ExportXLSX = "xslx(اکسل)"
	ExportCSV  = "csv"
)

// Search Menu
const (
	SourceSelectButton   = "انتخاب منبع"
	KeywordSearchButton  = "کلید واژه ها"
	SearchSettingsButton = "تنظیمات جستجو"
	StartSearchButton    = "شروع جستجو"
)

// Filters Menu
const (
	PriceFilter         = "قیمت"
	AreaFilter          = "متراژ"
	RoomsFilter         = "تعداد اتاق خواب"
	PropertyTypeFilter  = "نوع ملک"
	BuildingAgeFilter   = "سن بنا"
	FloorFilter         = "طبقه"
	StorageFilter       = "انبار"
	ElevatorFilter      = "آسانسور"
	AdDateFilter        = "تاریخ ایجاد آگهی"
	LocationFilter      = "موقعیت"
	RemoveFiltersButton = "پاک کردن فیلتر ها"
	ApplyFiltersButton  = "اعمال فیلتر"
)

// Price Range Buttons
const (
	PriceUnder500M  = "زیر ۵۰۰ میلیون تومان"
	Price500MTo700M = "از ۵۰۰ تا ۷۰۰ میلیون تومان"
	Price700MTo900M = "از ۷۰۰ تا ۹۰۰ میلیون تومان"
	Price900MTo1B   = "از ۹۰۰ میلیون تا ۱ میلیارد تومان"

	Price1BTo1_5B = "از ۱ تا ۱.۵ میلیارد تومان"
	Price1_5BTo2B = "از ۱.۵ تا ۲ میلیارد تومان"
	Price2BTo3B   = "از ۲ تا ۳ میلیارد تومان"
	Price3BTo4B   = "از ۳ تا ۴ میلیارد تومان"
	Price4BTo5B   = "از ۴ تا ۵ میلیارد تومان"
	Price5BTo7B   = "از ۵ تا ۷ میلیارد تومان"
	Price7BTo10B  = "از ۷ تا ۱۰ میلیارد تومان"

	Price10BTo15B   = "از ۱۰ تا ۱۵ میلیارد تومان"
	Price15BTo20B   = "از ۱۵ تا ۲۰ میلیارد تومان"
	Price20BTo30B   = "از ۲۰ تا ۳۰ میلیارد تومان"
	Price30BTo40B   = "از ۳۰ تا ۴۰ میلیارد تومان"
	Price40BTo50B   = "از ۴۰ تا ۵۰ میلیارد تومان"
	Price50BTo75B   = "از ۵۰ تا ۷۵ میلیارد تومان"
	Price75BTo100B  = "از ۷۵ تا ۱۰۰ میلیارد تومان"
	Price100BTo200B = "از ۱۰۰ تا ۲۰۰ میلیارد تومان"
	Price200BTo300B = "از ۲۰۰ تا ۳۰۰ میلیارد تومان"
	Price300BTo500B = "از ۳۰۰ تا ۵۰۰ میلیارد تومان"
	Price500BTo700B = "از ۵۰۰ تا ۷۰۰ میلیارد تومان"
	Price700BTo900B = "از ۷۰۰ تا ۹۰۰ میلیارد تومان"
	PriceOver900B   = "بیش از ۹۰۰ میلیارد تومان"
)

// Area Range Buttons
// Area Range Buttons - From 50 to 10000 square meters, Limited to 20
const (
	AreaUnder50  = "زیر ۵۰ متر مربع"
	Area50To75   = "از ۵۰ تا ۷۵ متر مربع"
	Area75To100  = "از ۷۵ تا ۱۰۰ متر مربع"
	Area100To150 = "از ۱۰۰ تا ۱۵۰ متر مربع"

	Area150To200  = "از ۱۵۰ تا ۲۰۰ متر مربع"
	Area200To250  = "از ۲۰۰ تا ۲۵۰ متر مربع"
	Area250To300  = "از ۲۵۰ تا ۳۰۰ متر مربع"
	Area300To400  = "از ۳۰۰ تا ۴۰۰ متر مربع"
	Area400To500  = "از ۴۰۰ تا ۵۰۰ متر مربع"
	Area500To750  = "از ۵۰۰ تا ۷۵۰ متر مربع"
	Area750To1000 = "از ۷۵۰ تا ۱۰۰۰ متر مربع"

	Area1000To1500  = "از ۱۰۰۰ تا ۱۵۰۰ متر مربع"
	Area1500To2000  = "از ۱۵۰۰ تا ۲۰۰۰ متر مربع"
	Area2000To3000  = "از ۲۰۰۰ تا ۳۰۰۰ متر مربع"
	Area3000To4000  = "از ۳۰۰۰ تا ۴۰۰۰ متر مربع"
	Area4000To5000  = "از ۴۰۰۰ تا ۵۰۰۰ متر مربع"
	Area5000To7500  = "از ۵۰۰۰ تا ۷۵۰۰ متر مربع"
	Area7500To10000 = "از ۷۵۰۰ تا ۱۰۰۰۰ متر مربع"
	AreaOver10000   = "بیش از ۱۰۰۰۰ متر مربع"
)

// Rooms
const (
	Bedrooms0      = "بدون اتاق خواب"
	Bedrooms1      = "۱ اتاق خواب"
	Bedrooms2      = "۲ اتاق خواب"
	Bedrooms3      = "۳ اتاق خواب"
	Bedrooms4      = "۴ اتاق خواب"
	Bedrooms5      = "۵ اتاق خواب"
	Bedrooms6      = "۶ اتاق خواب"
	Bedrooms7      = "۷ اتاق خواب"
	Bedrooms8      = "۸ اتاق خواب"
	Bedrooms9      = "۹ اتاق خواب"
	Bedrooms10     = "۱۰ اتاق خواب"
	BedroomsOver10 = "بیش از ۱۰ اتاق خواب"
)

// Property Type Buttons
const (
	PropertyApartment  = "آپارتمان"
	PropertyVilla      = "ویلا"
	PropertyCommercial = "تجاری"
	PropertyOffice     = "اداری"
	PropertyLand       = "زمین"
)

// Building Age Buttons - From New to Over 10 Years
const (
	BuildingAgeNew     = "نوساز"
	BuildingAge1Year   = "۱ سال"
	BuildingAge2Years  = "۲ سال"
	BuildingAge3Years  = "۳ سال"
	BuildingAge4Years  = "۴ سال"
	BuildingAge5Years  = "۵ سال"
	BuildingAge6Years  = "۶ سال"
	BuildingAge7Years  = "۷ سال"
	BuildingAge8Years  = "۸ سال"
	BuildingAge9Years  = "۹ سال"
	BuildingAge10Years = "۱۰ سال"
	BuildingAgeOver10  = "بیش از ۱۰ سال"
)

// Floor Number Buttons - From Floor ۰ to Over 10
const (
	Floor0      = "طبقه همکف"
	Floor1      = "طبقه ۱"
	Floor2      = "طبقه ۲"
	Floor3      = "طبقه ۳"
	Floor4      = "طبقه ۴"
	Floor5      = "طبقه ۵"
	Floor6      = "طبقه ۶"
	Floor7      = "طبقه ۷"
	Floor8      = "طبقه ۸"
	Floor9      = "طبقه ۹"
	Floor10     = "طبقه ۱۰"
	FloorOver10 = "بالاتر از ۱۰ طبقه"
)

const (
	Yes     = "بله"
	No      = "خیر"
	Unknown = "مهم نیست"
)

// Account Management Menu
const (
	SetUsernameButton    = "ویرایش نام کاربری"
	ShareBookmarksButton = "اشتراک گذاری علاقه‌مندی‌ها"
	LogoutButton         = "خروج از حساب"
)

// Search Results Menu
const (
	SaveToFavoritesButton = "ذخیره در علاقه‌مندی‌ها"
	ViewDetailsButton     = "نمایش جزئیات"
	BackToSearchButton    = "بازگشت به جستجو"
)

// Export Menu
const (
	DownloadZipButton   = "دانلود فایل ZIP"
	SendEmailButton     = "ارسال به ایمیل"
	FileFormatSelection = "انتخاب فرمت فایل (CSV یا XLSX)"
)

// Crawler Admin Menu
const (
	CrawlerStatusButton    = "مشاهده وضعیت کرالر"
	PeriodicSettingsButton = "تنظیمات دوره‌ای"
	ViewLogsButton         = "مشاهده لاگ‌ها"
)

// ‌Bot Messages
const (
	WelcomeMsg  = "سلام به بات جادویی ما خوش آمدی، مسلما برات سوال که ما چطور میتونیم تو رو به خونه رویاهات نزدیک کنیم! خب خیلی ساده به جایاین که بری کلی  خونه ببینی و از بین هزاران خونه یکی رو پسند کنی و یه لیست بلند بالا برای خودت بسازی و هر روز بهشون سر بزنی  و.... فقط کافیه مشخصات خونه ایی که میخوایی رو دست این بات بسپری، این بات با دقت بالا تمام خونه هایی که میتونه شبیه به خونه مدنظرت باشه رو پیدا میکنه و هر روز بهت خبرهای خوبی میده...  خیلی خوب ، اگر برای شروع نیاز کمک داری دکمه راهنمایی رو بزن. "
	UsernameMsg = "دوست داری نام کاربریت چی باشه؟؟"
)
