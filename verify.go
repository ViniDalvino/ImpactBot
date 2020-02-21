package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onReady2(discord *discordgo.Session, ready *discordgo.Ready) {
	go func() {
		prev := ""
		total := 0
		donatorCount := 0
		for {
			st, err := discord.GuildMembers(IMPACT_SERVER, prev, 1000)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("Fetched", len(st), "more members, total so far is", total)
			if len(st) == 0 {
				log.Println("No more members!")
				break
			}
			prev = st[len(st)-1].User.ID
			for _, member := range st {
				total++
				memberVerificationCheck(member)
				if hasRole(member, Donator) {
					donatorCount++
				}
			}
		}
		log.Println("Processed", total, "members")
		log.Println("There are", donatorCount, "donators")
	}()
}

func memberVerificationCheck(member *discordgo.Member) {
	if len(member.Roles) > 0 && !hasRole(member, Verified) {
		log.Println("Member", member.User.ID, "had roles not including verified")
		err := discord.GuildMemberRoleAdd(IMPACT_SERVER, member.User.ID, Verified.ID)
		if err != nil {
			log.Println(err)
		}
	}
}
