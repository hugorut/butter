package database

import "time"

type User struct {
	ID        uint      `gorm:"primary_key"json:"id"`
	Name      string    `gorm:"column:name"json:"name"`
	Email     string    `gorm:"column:email"json:"email"`
	Password  string    `gorm:"column:password"json:"-"`
	CreatedAt time.Time `gorm:"column:created_at"json:"createdAt"`
	Tests     []Test    `gorm:"ForeignKey:ID"json:"tests"`
}

// The table name for the user model
func (User) TableName() string {
	return "users"
}

type Image struct {
	ID        int       `gorm:"primary_key"json:"id"`
	TestID    int       `gorm:"index,column:test_id"json:"-"`
	Path      string    `gorm:"column:path"json:"path"`
	CreatedAt time.Time `gorm:"column:created_at"json:"createdAt"`
}

// The table name for the image model
func (Image) TableName() string {
	return "images"
}

type Test struct {
	ID          uint       `gorm:"PrimaryKey,column:id"json:"id"`
	UserID      uint       `gorm:"index,column:user_id"json:"-"`
	User        User       `gorm:"ForeignKey:ID"json:"user"`
	TestTypeID  uint       `gorm:"index,column:test_type_id"json:"-"`
	TestType    TestType   `json:"testType"`
	Images      []Image    `gorm:"ForeignKey:TestID"json:"images"`
	Questions   []Question `gorm:"ForeignKey:TestID"json:"questions"`
	Code        string     `gorm:"column:code"json:"code"`
	Description string     `gorm:"column:description"json:"description"`
	TestExtras  []byte     `gorm:"column:test_extras"json:"testExtras"`
	CreatedAt   time.Time  `gorm:"column:created_at"json:"created_at"`
	Active      bool       `gorm:"column:active"`
}

// The table name for the test model
func (Test) TableName() string {
	return "tests"
}

type TestType struct {
	ID   uint   `gorm:"primary_key,column:id"json:"id"`
	Name string `gorm:"column:name"json:"name"`
}

// The table name for the test type model
func (TestType) TableName() string {
	return "test_types"
}

type Question struct {
	ID          uint   `gorm:"primary_key,column:id"json:"id"`
	Description string `gorm:"column:description"json:"description"`
	TestID      uint   `json:"-"`
	Answers     []Answer
}

// The table name for the questions model
func (Question) TableName() string {
	return "questions"
}

type Answer struct {
	UserID     uint     `gorm:"column:user_id"json:"-"`
	QuestionID uint     `gorm:"column:question_id"json:"-"`
	Question   Question `json:"question"`
	Value      string   `gorm:"column:value"json:"value"`
}

// The table name for the questions model
func (Answer) TableName() string {
	return "answers"
}

type Invitee struct {
	ID            uint           `gorm:"primary_key,column:id"json:"id"`
	Email         string         `gorm:"column:email"json:"email"`
	InviteeGroups []InviteeGroup `gorm:"many2many:invite_groups_have_invitees;"`
}

// The table name for the Invitee model
func (Invitee) TableName() string {
	return "invitees"
}

type InviteeGroup struct {
	ID   uint   `gorm:"primary_key,column:id"json:"id"`
	Name string `gorm:"column:name"json:"name"`
}

// The table name for the InviteeGroup model
func (InviteeGroup) TableName() string {
	return "invite_groups"
}

type TestsHaveInviteeGroups struct {
	TestID         uint `gorm:"column:tests_id"json:"-"`
	InviteeGroupID uint `gorm:"column:invite_groups_id"json:"-"`
}

// The table name for the InviteeGroup model
func (TestsHaveInviteeGroups) TableName() string {
	return "tests_have_invite_groups"
}
