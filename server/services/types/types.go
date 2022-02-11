package types

import (
	"regexp"
	"strconv"
	"strings"
)

var blankRegexp *regexp.Regexp

func init() {
	blankRegexp, _ = regexp.Compile("\\s")
}

type RegisterInput struct {
	Name     string
	Password string
	Mail     string
}

func (input *RegisterInput) CheckPass() bool {
	input.Name = blankRegexp.ReplaceAllString(input.Name, "")
	input.Mail = blankRegexp.ReplaceAllString(input.Mail, "")
	return input.Name != "" && input.Mail != "" && input.Password != ""
}

type RegisterResult struct {
	// success, xxx already in use, invalid xxx, failed, already login, not empty error
	Message string
}

type LoginInput struct {
	Mail     string
	Password string
}

func (input *LoginInput) CheckPass() bool {
	input.Mail = blankRegexp.ReplaceAllString(input.Mail, "")
	return input.Mail != "" && input.Password != ""
}

type LoginResult struct {
	Message string
	UserId  string
	Role    string
}

type LogoutResult struct {
	Message string
}

type ForgetResult struct {
	Message string
}

type ResetInput struct {
	Password string
	Token    string
}

func (input *ResetInput) CheckPass() bool {
	return input.Password != "" && input.Token != ""
}

type ResetResult struct {
	Message string
}

type SubmitInput struct {
	ChallengeId string
	Flag        string
}

func (input *SubmitInput) CheckPass() bool {
	input.Flag = blankRegexp.ReplaceAllString(input.Flag, "")
	return input.Flag != ""
}

type SubmitResult struct {
	Message string
}

type BulletinPubInput struct {
	Title   string
	Content string
	Topping bool
}

func (input *BulletinPubInput) CheckPass() bool {
	input.Title = strings.TrimSpace(input.Title)
	input.Content = strings.TrimSpace(input.Content)
	return input.Title != ""
}

type BulletinPubResult struct {
	Message string
}

type UserInfoUpdateInput struct {
	UserId string
	Name   string
	Role   string
	Mail   string
	State  string
}

func (input *UserInfoUpdateInput) CheckPass() bool {
	input.Name = blankRegexp.ReplaceAllString(input.Name, "")
	input.Mail = blankRegexp.ReplaceAllString(input.Mail, "")
	return input.Name != "" && input.Mail != "" && checkRole(input.Role) && checkUserState(input.State)
}

func checkRole(str string) bool {
	return str == "member" || str == "team" || str == "admin"
}

func checkUserState(str string) bool {
	return str == "normal" || str == "disabled"
}

type UserInfoUpdateResult struct {
	Message string
}

type ChallengeMutateInput struct {
	ChallengeId  string
	Name         string
	Category     string
	Score        ScoreTypeInput
	Flag         FlagTypeInput
	Description  string
	ExternalLink []string
	Hint         []string
	State        string
}

func (input *ChallengeMutateInput) CheckPass() bool {
	input.Name = strings.TrimSpace(input.Name)
	input.Description = strings.TrimSpace(input.Description)
	return input.Name != "" && checkChallengeCategory(input.Category) && input.Score.CheckPass() && input.Flag.CheckPass() && checkChallengeState(input.State)
}

func checkChallengeCategory(str string) bool {
	return str == "WEB" || str == "RE" || str == "PWN" || str == "MISC" || str == "CRYPTO" // TODO:
}

func checkChallengeState(str string) bool {
	return str == "enabled" || str == "disabled"
}

type ScoreTypeInput struct {
	Dynamic   bool
	BaseScore string
}

func (input *ScoreTypeInput) CheckPass() bool {
	parsedScore, err := strconv.Atoi(input.BaseScore)
	if err != nil {
		return false
	}
	return parsedScore >= 0
}

type FlagTypeInput struct {
	Dynamic bool
	Value   string
}

func (input *FlagTypeInput) CheckPass() bool {
	input.Value = blankRegexp.ReplaceAllString(input.Value, "")
	return input.Value != ""
}

type ChallengeMutateResult struct {
	Message string
}

type ChallengeRemoveResult struct {
	Message string
}

type UserInfoResult struct {
	Message  string
	UserInfo UserInfo
}

type UserInfo struct {
	UserId   string
	Name     string
	Avatar   string
	Mail     string
	JoinTime string
	Score    string
	Role     string
	State    string
	Rank     int
}

type AllUserInfoResult struct {
	Message      string
	AllUserInfos []UserInfo
}

type RankResult struct {
	Message         string
	RankResultDescs []RankResultDesc
}

type ChallengeInfosResult struct {
	Message        string
	ChallengeInfos []ChallengeInfo
}

type RankResultDesc struct {
	UserId string
	Name   string
	Avatar string
	Score  int
}

type ChallengeInfo struct {
	ChallengeId  string
	Category     string
	Name         string
	Score        int
	Description  string
	ExternalLink []string
	Hint         []string
	Blood        []BloodInfo
	Done         bool
}

type BloodInfo struct {
	UserId string
	Name   string
	Avatar string
}

type ChallengeConfigsResult struct {
	Message          string
	ChallengeConfigs []ChallengeConfig
}

type ChallengeConfig struct {
	ChallengeId  string
	Name         string
	Category     string
	Score        ScoreType
	Flag         FlagType
	Description  string
	ExternalLink []string
	Hint         []string
	State        string
}

type ScoreType struct {
	Dynamic   bool
	BaseScore string
}

type FlagType struct {
	Dynamic bool
	Value   string
}

type BulletinResult struct {
	Message   string
	Bulletins []Bulletin
}

type Bulletin struct {
	Title       string
	Content     string
	PublishTime string
}

type SubmitHistoryResult struct {
	Message     string
	SubmitInfos []SubmitInfo
}

type SubmitInfo struct {
	SubmitTime    string
	ChallengeName string
	Mark          int
}