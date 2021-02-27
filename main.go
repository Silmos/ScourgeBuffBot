package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Bot struct {
		Token string //`json:"token"`
	}
	Hordeunited struct {
		Hudiscordid string   //`json:"hudiscordid"`
		Huchannelid string   //`json:"huchannelid"`
		Hurendid    string   //`json:"hurendid"`
		Hudragonid  []string //`json:"hudragonid"`
		Huhakkarid  string   //`json:"huhakkarid"`
	}
	Intervalfordeletes int
	Maxmessageage      int
}

type Guilds struct {
	Guildconfigs []Guildconfig
}

type Guildconfig struct {
	Guildname   string //'json: "guildname"
	Channelid   string //`json:"channelid"`
	Guildid     string //`json:"guildid"`
	Mentionrole bool   //`json:"mentionrole"`
	Rendrole    string //`json:"rendroleid"`
	Hakkarrole  string //`json:"hakkarroleid"`
	Dragonrole  string //`json:"dragonroleid"`
}

var mentionrole string

func loadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func loadGuilds(file string) (Guilds, error) {
	var guilds Guilds
	guildsFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer guildsFile.Close()
	byteValue, _ := ioutil.ReadAll(guildsFile)
	json.Unmarshal(byteValue, &guilds)
	return guilds, err
}

func main() {

	config, _ := loadConfiguration("config.json")

	dg, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dg.AddHandler(messageHandler)
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-done
		fmt.Println("Interrupted")
		dg.Close()
		fmt.Println("Geacefully stopping Bot")
		os.Exit(0)
	}()
	checkoldmessages(dg)

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	config, _ := loadConfiguration("config.json")
	guildconfig, _ := loadGuilds("guilds.json")
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.GuildID == config.Hordeunited.Hudiscordid && m.ChannelID == config.Hordeunited.Huchannelid {
		var channels = regexp.MustCompile(config.Hordeunited.Hurendid + "|" + config.Hordeunited.Hudragonid[0] + "|" + config.Hordeunited.Hudragonid[1] + "|" + config.Hordeunited.Huhakkarid)
		if channels.MatchString(m.Content) {
			channelname := m.Content
			if len(m.Embeds) > 0 {
				for i := 0; i < len(guildconfig.Guildconfigs); i++ {
					if guildconfig.Guildconfigs[i].Mentionrole == true {
						switch {
						case strings.Contains(channelname, config.Hordeunited.Hurendid):
							mentionrole = guildconfig.Guildconfigs[i].Rendrole
						case strings.Contains(channelname, config.Hordeunited.Huhakkarid):
							mentionrole = guildconfig.Guildconfigs[i].Hakkarrole
						case strings.Contains(channelname, config.Hordeunited.Hudragonid[0]), strings.Contains(channelname, config.Hordeunited.Hudragonid[1]):
							mentionrole = guildconfig.Guildconfigs[i].Dragonrole

						}
						_, _ = s.ChannelMessageSendComplex(guildconfig.Guildconfigs[i].Channelid, &discordgo.MessageSend{Content: "<@&" + mentionrole + ">", Embed: m.Embeds[0]})
						fmt.Println("Sent Message with Role-Mention")
					} else {
						_, _ = s.ChannelMessageSendComplex(guildconfig.Guildconfigs[i].Channelid, &discordgo.MessageSend{Content: m.Content, Embed: m.Embeds[0]})
						fmt.Println("Sent Message")
					}
				}

			}

		}

	}
}

func checkoldmessages(s *discordgo.Session) {
	config, _ := loadConfiguration("config.json")
	guildconfig, _ := loadGuilds("guilds.json")
	ticker := time.NewTicker(time.Duration(config.Intervalfordeletes) * time.Minute)
	for range ticker.C {
		for i := 0; i < len(guildconfig.Guildconfigs); i++ {
			getmessage, _ := s.ChannelMessages(guildconfig.Guildconfigs[i].Channelid, 100, "", "", "")
			getmessages := []string{}
			for _, m := range getmessage {
				messagetime, _ := m.Timestamp.Parse()
				now := time.Now()
				then := now.Add(time.Duration(-config.Maxmessageage) * time.Minute)
				if m.Author.ID == s.State.User.ID && messagetime.Before(then) {
					getmessages = append(getmessages, m.ID)
				}
			}
			_ = s.ChannelMessagesBulkDelete(guildconfig.Guildconfigs[i].Channelid, getmessages)
			getmessages = nil
		}
		fmt.Println("Deleting old messages on all Servers")
	}
}
