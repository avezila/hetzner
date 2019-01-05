package main

import (
	"bytes"
	"os"
	"strconv"
	"text/template"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var LockPath = "./lock/"

func SendTelegram(servers []AServer) error {
	var toSend []AServer

	for _, s := range servers {
		if s.HddGbPerEurAdj > 0.95 || s.CPUBenchPerEurAdj > 0.95 || s.Average > 0.26 {
			toSend = append(toSend, s)
			continue
		}
		if s.PriceEur < 24 && s.CPUBenchmark > 9100 && s.RAM > 16 {
			toSend = append(toSend, s)
			continue
		}
	}

	if toSend == nil {
		return nil
	}

	if err := os.MkdirAll(LockPath, 0700); err != nil {
		return err
	}

	var toSend2 []AServer
	for _, s := range toSend {
		if _, e := os.Stat(LockPath + strconv.Itoa(int(s.Key))); os.IsNotExist(e) {
			toSend2 = append(toSend2, s)
			continue
		}
	}
	if toSend2 == nil {
		return nil
	}

	bot, err := tgbotapi.NewBotAPI("787991996:AAEC0_MHueKV4bETpzZb3ZRfPa2wqfgATX0")
	if err != nil {
		return err
	}

	template, err := template.New("message").Parse(`key: {{.Key}}
cpu: {{.CPUBenchmark}} {{.CPU}}
ram: {{.RAMHr}}
hdd: {{.HddHr}}
allb: {{.Average}}
cpub: {{.CPUBenchPerEurAdj}}
ramb: {{.RAMGbPerEurAdj}}
hddb: {{.HddGbPerEurAdj}}
ssdb: {{.SsdGbPerEurAdj}}
price: {{.Price}}
reduce: {{.NextReduceHr}}
{{range .Description}}{{.}} {{end}}

https://api1.nirhub.ru/s/{{.Hash}}.html
https://www.hetzner.com/sb
`)
	if err != nil {
		return err
	}
	for _, s := range toSend2 {
		buf := bytes.NewBuffer(nil)
		template.Execute(buf, s)
		msg := tgbotapi.NewMessageToChannel("@gongo_hezner", buf.String())
		bot.Send(msg)
		os.Create(LockPath + strconv.Itoa(int(s.Key)))
	}
	return nil
}
