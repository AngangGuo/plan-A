package main

import (
	"fmt"
	"strings"
	"time"
)

// Old testament and new testament book names and chapters
var (
	oldNewNames = []string{
		"創世記",
		"出埃及記",
		"利未記",
		"民數記",
		"申命記",
		"約書亞記",
		"士師記",
		"路得記",
		"撒母耳記上",
		"撒母耳記下",
		"列王紀上",
		"列王紀下",
		"歷代志上",
		"歷代志下",
		"以斯拉記",
		"尼希米記",
		"以斯帖記",
		"約伯記",
		"傳道書",
		"雅歌書",
		"以賽亞書",
		"耶利米書",
		"耶利米哀歌",
		"以西結書",
		"但以理書",
		"何西阿書",
		"約珥書",
		"阿摩司書",
		"俄巴底亞書",
		"約拿書",
		"彌迦書",
		"那鴻書",
		"哈巴谷書",
		"西番雅書",
		"哈該書",
		"撒迦利亞書",
		"瑪拉基書",
		"馬太福音",
		"馬可福音",
		"路加福音",
		"約翰福音",
		"使徒行傳",
		"羅馬書",
		"哥林多前書",
		"哥林多後書",
		"加拉太書",
		"以弗所書",
		"腓立比書",
		"歌羅西書",
		"帖撒羅尼迦前書",
		"帖撒羅尼迦後書",
		"提摩太前書",
		"提摩太後書",
		"提多書",
		"腓利門書",
		"希伯來書",
		"雅各書",
		"彼得前書",
		"彼得後書",
		"約翰一書",
		"約翰二書",
		"約翰三書",
		"猶大書",
		"啟示錄",
	}
	oldNewChapters = []int{
		50,
		40,
		27,
		36,
		34,
		24,
		21,
		4,
		31,
		24,
		22,
		25,
		29,
		36,
		10,
		13,
		10,
		42,
		12,
		8,
		66,
		52,
		5,
		48,
		12,
		14,
		3,
		9,
		1,
		4,
		7,
		3,
		3,
		3,
		2,
		14,
		4,
		28,
		16,
		24,
		21,
		28,
		16,
		16,
		13,
		6,
		6,
		4,
		4,
		5,
		3,
		6,
		4,
		3,
		1,
		13,
		5,
		5,
		3,
		5,
		1,
		1,
		1,
		22,
	}
)

var psProText []string

func main() {
	initPsPro()

	day := 0
	// begin the new Bible reading plan starting from tomorrow
	startDate := time.Date(2020, 8, 30, 0, 0, 0, 0, time.UTC)
	for bookIndex, totalChapters := range oldNewChapters {
		daily := getDailyChapters(totalChapters)
		for _, ch := range daily {
			day += 1
			theDate := startDate.AddDate(0, 0, day)
			txt := fmt.Sprintf("%s(%s)\n==============\n第%d天：\n%s 第%s章\n%s\n",
				theDate.Format("2006-01-02"),
				getWeekdayCN(theDate.Weekday()),
				day,
				oldNewNames[bookIndex],
				ch,
				getPsPro(day))
			fmt.Println(txt)
		}
	}

	// finish the remaining Proverbs
	proverbsRemain := psProText[day-len(psProText):]
	start := 0
	lenPsPro := len(proverbsRemain)
	for start < lenPsPro {
		day++
		end := start + 2
		if end > lenPsPro-1 {
			end = lenPsPro - 1
		}
		theDate := startDate.AddDate(0, 0, day)
		txt := fmt.Sprintf("%s（%s）\n==============\n第%d天：\n%s - %s\n",
			theDate.Format("2006-01-02"),
			getWeekdayCN(theDate.Weekday()),
			day,
			proverbsRemain[start],
			proverbsRemain[end])
		txt = strings.Replace(txt, "章 - 箴言 第", "-", 1)
		fmt.Println(txt)
		start += 3
	}
}

// return one chapter of Psalm or Proverbs
func getPsPro(day int) string {
	n := len(psProText)
	if day > n {
		day = day - n
	}
	return psProText[day-1]
}

// return reading plan for each day according to the book chapters
func getDailyChapters(chapters int) (plan []string) {
	max := 3
	remains := chapters

	start := 1
	end := start
	for {
		switch {
		case remains == 1:
			txt := fmt.Sprintf("%d", start)
			plan = append(plan, txt)
			return
		case remains <= 3:
			end = start + remains - 1
			txt := fmt.Sprintf("%d-%d", start, end)
			plan = append(plan, txt)
			return
		case remains == 4:
			// each day 2 chapters
			end = start + 1
			txt := fmt.Sprintf("%d-%d", start, end)
			plan = append(plan, txt)
			start = end + 1
			remains -= 2
		default:
			end = start + max - 1
			txt := fmt.Sprintf("%d-%d", start, end)
			plan = append(plan, txt)
			start = end + 1
			remains -= max
		}
	}

}

func getWeekdayCN(weekday time.Weekday) string {
	weekdayCN := []string{
		"周日",
		"周一",
		"周二",
		"周三",
		"周四",
		"周五",
		"周六",
	}
	return weekdayCN[weekday]
}

// Todo: refactor: use slice for chapter 119 verse separation
// return daily reading plan for Psalm or Proverbs for specific date
func initPsPro() {
	// Psalm and Proverbs book names and chapters
	psProText = []string{}
	txt := ""

	// fill in psalm
	bookName := "詩篇"
	bookChapters := 150
	// Psalm 119 separate 6 times
	splitTimes := 6
	for currentDay := 1; currentDay <= bookChapters+splitTimes-1; currentDay++ {

		switch {
		case currentDay < 119:
			txt = fmt.Sprintf("%s 第%d章", bookName, currentDay)
		case currentDay == 119:
			txt = fmt.Sprintf("%s 第119:1-32节", bookName)
		case currentDay == 120:
			txt = fmt.Sprintf("%s 第119:33-64节", bookName)
		case currentDay == 121:
			txt = fmt.Sprintf("%s 第119:65-88节", bookName)
		case currentDay == 122:
			txt = fmt.Sprintf("%s 第119:89-120节", bookName)
		case currentDay == 123:
			txt = fmt.Sprintf("%s 第119:121-144节", bookName)
		case currentDay == 124:
			txt = fmt.Sprintf("%s 第119:145-176节", bookName)
		case currentDay > 124:
			txt = fmt.Sprintf("%s 第%d章", bookName, currentDay-splitTimes+1)
		}
		psProText = append(psProText, txt)
	}

	// fill in proverbs
	bookName = "箴言"
	bookChapters = 31
	for i := 1; i <= bookChapters; i++ {
		txt = fmt.Sprintf("%s 第%d章", bookName, i)
		psProText = append(psProText, txt)
	}
}
