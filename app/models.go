package app

//-- auto-generated definition
//create table group_candidate
//(
//  id                   int auto_increment
//    primary key,
//  group_candidate_name varchar(16)     not null,
//  votes_number         int default '0' null
//)
//  charset = utf8;
type GroupCandidate struct {
	Id                 int    `gorm:"AUTO_INCREMENT;primary_key"`
	GroupCandidateName string `gorm:"not null"`
	VotesNumber        int    `gorm:"default:'0'"`

	tableName string `gorm:"-"`
}

type PartyCandidate struct {
	Id                 int    `gorm:"AUTO_INCREMENT;primary_key"`
	PartyCandidateName string `gorm:"not null"`
	VotesNumber        int    `gorm:"default:'0'"`

	tableName string `gorm:"-"`
}

//-- auto-generated definition
//create table user
//(
//  user_id         varchar(16)     not null
//    primary key,
//  type            int             null
//  comment '1是计票员 2是管理员',
//  party_count_num int default '0' null,
//  group_count_num int default '0' null
//)
//  charset = utf8;
type User struct {
	UserId        string `gorm:"not null;PRIMARY_KEY;"`
	Type          int    `sql:"comment:'1是计票员 2是管理员'"`
	PartyCountNum int    `gorm:"default:0"`
	GroupCountNum int    `gorm:"default:0"`

	tableName string `gorm:"-"`
}

//-- auto-generated definition
//create table vote_number
//(
//  id        int default '-1' not null
//    primary key,
//  party_num int default '0'  not null,
//  group_num int default '0'  null
//)
//  charset = utf8;
type VoteNumber struct {
	Id       int `gorm:"not null;	primary_key"`
	PartyNum int `gorm:"default:0"`
	GroupNum int `gorm:"default:0"`

	tableName string `gorm:"-"`
}

func (c *GroupCandidate) TableName() string {
	return c.tableName
}
func (c *PartyCandidate) TableName() string {
	return c.tableName
}
func (u *User) TableName() string {
	return u.tableName
}
func (v *VoteNumber) TableName() string {
	return v.tableName
}
