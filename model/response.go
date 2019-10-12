package model

import "strings"

var MainPageResponse = `سلام. خوبی ؟ خب این اولین کاریه که انجام دادی. تبریک میگم. ببین این چیزی که داری الان میبینی اگه گفتی چیه ؟ بهش میگن Response Body. توی هر مرحله من توی Response Body ها بهت میگم برای دفعه بعدی باید چیکار کنی. این چالش یه چیزیه که باید ازش یه چیزی یاد بگیری. پس چرت و پرتایی که آموزشیه رو هم یادت نره بخونی. حالا برای شروع. یدونه ریکوئست GET برا من بفرست با شماره دانشجوییت.`
var FailedGetMessage = `خب. برادر / خواهر گرامی. ظاهرا اینجا رو اشتباه کردی. قرار بود هدر برام بفرستی با شماره دانشجویی که من بشناسمت. برو تو اتاقت. قشنگ به اشتباهاتت فک کن. وقتی متوجه اشتباهاتت شدی بیا اینجا و تست کن.`
var FailedQueryParam = `یه چیزی رو داری اشتباه میزنی. ببین میخوام کمکت کنم ولی من یه سرور احمقم. فقط در این حد میتونم کمکت کنم که بهت میگم چیزی که داری بهم میدی درست نیست. درستش کن.`
var SuccessQueryParam = `آقا تبریک میگم. ایشالا بتونیم جبران کنیم براتون. خب حالا تشریف ببرید مرحله چهارم با شیرینی در خدمتیم`
var FailedPassword = `به کسی نمیگم که اینکارو کردی. به شرطی که خودت بری زود درستش کنی. پسورد و شماره دانشجوییت رو بفرست برام ببینم بچهههه!`
var SuccessPOST = `به سلامتی اینم اوکی کردی. دیگه داریم به آخرا میرسیم. اگه بدونی چه چیزایی ازت میدونم. درجا پشمات میریزه. راستی ما همه چیز رو اینجا بهت نمیدیما. من کلا دوست دارم بپیچونمت. بعضی اطلاعات رو ممکنه یجور دیگه بهت بدم. همه مدل اطلاعاتی رو دریاب. آخه نه ما خیلی شاخیم. هیچ کاری رو بی دلیل نمیکنیم`
var SuccessAuth = `پشمااام. میدونی تا اینجا اومدی یعنی چی ؟ یعنی دیگه داری یاد میگیری؟‌ ولی خداوکیلی دارم اینهمه زحمت میکشم من که یاد بگیرید. من سرما خوردم. دارم کد میزنم با مریضی. یاد بگیرید دیگه.`
var SuccessCookie = `آقا از دستم ناراحت نشو. میگمااا اینا همش سرکاری بود =))) شرمنده اگه خیلی جدی کار رو دنبال کردی. ولی خب دیگه نمیتونم اذیتتون کنم. بابا ما آخه چرا باید چالش بذاریم ؟ فوقش همون اول میگیم فلان لینکو باز کنید. بجا اینهمه چالش مثلا میگفتم "یکی از راه هایی که میشه اطلاعات رو روی کامپیوتر طرف از طریق سرور سیو کرد از طریق کوکی هاست"`

func GreetingGetMessage(name string) string {
	txt := `سلام. میبینم که تونستی به خوبی و خوشی شماره دانشجوییت رو بفرستی <b>` + name + `</b>
هاان ؟ چی شد؟ فک کردی ما اسمت رو نمیدونیم ؟ ما همه چی رو در موردت میدونیم. حتی اونا رو هم میدونیم. آره. همونا. اگه میخوای دهنم قرص بمونه. باید یه رمزی رو بهم بدی. این رمزه رو بهت میگم چجوری پیدا کنی. فک نکن میتونی بری از انجمن یا شورا یا از دوستات رمزه رو بگیری. رمزه یه چیزیه که برا هرکسی فرق میکنه. خب برای اولین مرحله رمز اگه مردی‌ ( با عرض پوزش از خواهران گرامی ) یه GET دیگه بفرست. منتها اینبار پارامتر باید براش بذاری. یعنی همراه ریکوئستت باید Query parameter بفرستی. چیا رو باید بفرستی‌؟ یکی اسمت به صورت فینگیلیش (فامیلیتو بفرس ) و شماره دانشجوییت. فقط میگم با name و sid بفرست. یجور دیگه بفرستی این سیستم ما گاوه نمیفهمه. پس بدو. بدو عمو جون`
	return `<html dir="rtl"><body><p>` + txt + `</p></body></body>`
}

func SuccessQPMessage(name string, client string) string {
	var txt string
	if client != "" {
		if client == "Curl" {
			txt = `saaaaalaaaaaaam ` + name + `.\n khoobi ? Agha tabrik migam. na baba barikala, yad gerefti request befresti.\n\n Fek kardi hala ke dari az ` + client + ` estefade mikoni yani shakhi ? age shakhi bia mano hack kon. Vali namoosan Age kolle challange ro ba curl mizani yani shakhi`
		} else {
			txt = `saaaaalaaaaaaam ` + name + `.\n khoobi ? Agha tabrik migam. na baba barikala, yad gerefti request befresti.\n\n Fek kardi hala ke dari az ` + client + ` estefade mikoni yani shakhi ? age shakhi bia mano hack kon.`
		}
	} else {
		txt = `saaaaalaaaaaaam ` + name + `.\n khoobi ? Agha tabrik migam. na baba barikala, yad gerefti request befresti`
	}
	return txt
}

func DetectClient(userAgent string) string {
	userAgent = strings.ToLower(userAgent)
	if strings.Contains(userAgent, "postman") {
		return "Postman"
	}
	if strings.Contains(userAgent, "chrome") {
		return "Google Chrome"
	}
	if strings.Contains(userAgent, "firefox") {
		return "Firefox"
	}
	if strings.Contains(userAgent, "curl") {
		return "Curl"
	}
	return ""
}
