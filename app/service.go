package app

import (
	"vote-cli/dao"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"errors"
)

func insertCandidate(candidateType string, candidateName string) (err error) {
	db := dao.DB()
	table := db.Table(candidateType + "_candidate")
	if candidateType == "party" {
		err = table.Create(&PartyCandidate{
			PartyCandidateName: candidateName,
			VotesNumber:        0,
		}).Error
		return
	} else if candidateType == "group" {
		err = table.Create(&GroupCandidate{
			GroupCandidateName: candidateName,
			VotesNumber:        0,
		}).Error
		return
	}
	//err = table.Create(&map[string]interface{}{
	//	candidateType + "_candidate_name": candidateName,
	//	"votes_number":                    0,
	//}).Error
	return errors.New("invalid candidateType")
}

func insertUser(userType int, userId string) (err error) {
	db := dao.DB()
	table := db.Table("user")
	err = table.Create(&User{
		UserId: userId,
		Type:   userType,
	}).Error
	return
}
func initInsertVoteNumber() error {
	db := dao.DB()
	return db.Table("vote_number").Create(&VoteNumber{Id: -1, PartyNum: 0, GroupNum: 0}).Error
}
func CreateTables() {
	db := dao.DB()
	if !db.HasTable("party_candidate") {
		db.Set("gorm:table_options", "ENGINE=innodb,DEFAULT CHARSET=utf8").CreateTable(&PartyCandidate{tableName: "party_candidate"})
	}
	if !db.HasTable("group_candidate") {
		db.Set("gorm:table_options", "ENGINE=innodb,DEFAULT CHARSET=utf8").CreateTable(&GroupCandidate{tableName: "group_candidate"})
	}
	if !db.HasTable("user") {
		db.Set("gorm:table_options", "ENGINE=innodb,DEFAULT CHARSET=utf8").CreateTable(&User{tableName: "user"})
	}
	if !db.HasTable("vote_number") {
		db.Set("gorm:table_options", "ENGINE=innodb,DEFAULT CHARSET=utf8").CreateTable(&VoteNumber{tableName: "vote_number"})
	}
}
func InitInsert() (err error) {
	cnf, err := ioutil.ReadFile("conf.json")
	if err != nil {
		return
	}

	results := gjson.ParseBytes(cnf)
	partyArr := results.Get("candidate.party").Array()
	groupArr := results.Get("candidate.group").Array()

	for _, name := range partyArr {
		err = insertCandidate("party", name.String())
		if err != nil {
			return
		}
	}
	for _, name := range groupArr {
		err = insertCandidate("group", name.String())
		if err != nil {
			return
		}
	}
	ordinaryUserArr := results.Get("user.ordinary").Array()
	adminUserArr := results.Get("user.admin").Array()
	for _, name := range ordinaryUserArr {
		err = insertUser(1, name.String())
		if err != nil {
			return
		}
	}
	for _, name := range adminUserArr {
		err = insertUser(2, name.String())
		if err != nil {
			return
		}
	}
	err = initInsertVoteNumber()
	return
}
func dropTables() {
	db := dao.DB()
	db.DropTableIfExists("group_candidate")
	db.DropTableIfExists("party_candidate")
	db.DropTableIfExists("user")
	db.DropTableIfExists("vote_number")
}
