package command

import (
	"hc-bot/command/commands/sundries"
	"hc-bot/command/commands/trpg"
	coc2 "hc-bot/command/commands/trpg/coc"
)

// 注册所有的命令类
var help = &sundries.Help{}
var botoff = &sundries.Botoff{}
var jrrp = &sundries.Jrrp{}
var weather = &sundries.Weather{}
var translate = &sundries.Translate{}
var name = &sundries.Name{}
var rd = &trpg.Rd{}
var rh = &trpg.Rh{Rd: rd}
var coc = &coc2.Coc{}
var cocgs = &coc2.Cocgs{}
var cocgf = &coc2.Cocgf{}
var cocst = &coc2.Cocst{}
var cocrc = &coc2.Cocrc{}
var sc = &coc2.Sc{}
var en = &coc2.En{}

// Commands 命令与命令类对应的map
var Commands = map[string]Command{
	"help":    help,
	"botoff":  botoff,
	"jrrp":    jrrp,
	"weather": weather,
	"trans":   translate,
	"name":    name,
	"rd":      rd,
	"r":       rd,
	"rh":      rh,
	"coc":     coc,
	"cocgs":   cocgs,
	"cocgf":   cocgf,
	"st":      cocst,
	"rc":      cocrc,
	"sc":      sc,
	"en":      en,
}
