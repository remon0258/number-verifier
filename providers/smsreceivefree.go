package providers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/upmasked/number-verifier/util"
	"regexp"
	"strings"
)

type SMSReceiveFree struct {
}

func (p SMSReceiveFree) GetNumbers() ([]string, error) {
	var (
		numbers []string
		r = regexp.MustCompile(`\+([0-9]+)`)
	)

	doc, err := util.GetGoQueryDoc(fmt.Sprintf("%s/country/usa", p.GetProvider().BaseURL))
	if err != nil {
		return nil, err
	}

	doc.Find("a.numbutton").Each(func(i int, s *goquery.Selection) {
		numbers = append(numbers, r.FindStringSubmatch(s.Text())[1])
	})

	return numbers, nil
}

func (p SMSReceiveFree) GetMessages(number string) ([]string, error) {
	var (
		messages []string
	)

	doc, err := util.GetGoQueryDoc(fmt.Sprintf("%s/info/%s", p.GetProvider().BaseURL, number))
	if err != nil {
		return nil, err
	}

	doc.Find(".messagesTable tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		message := ""

		s.Find("td").Each(func(i int, s *goquery.Selection) {
			message += strings.TrimSpace(s.Text()) + " - "
		})

		messages = append(messages, message)
		if len(messages) > 5 {
			return false
		}

		return true
	})

	return messages, nil
}

func (p SMSReceiveFree) GetProvider() Provider {
	return Provider{
		Name: "SMSReceiveFree",
		BaseURL: "https://smsreceivefree.com",
	}
}
