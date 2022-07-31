// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

type Billing struct {
	BUuid         uuid.UUID             `json:"b_uuid"`
	BPrev         uuid.NullUUID         `json:"b_prev"`
	BType         sql.NullInt32         `json:"b_type"`
	BCreated      sql.NullTime          `json:"b_created"`
	BPaid         sql.NullTime          `json:"b_paid"`
	BLastReminder sql.NullTime          `json:"b_last_reminder"`
	BOrg          uuid.NullUUID         `json:"b_org"`
	BDetail       pqtype.NullRawMessage `json:"b_detail"`
	BPaidDetail   pqtype.NullRawMessage `json:"b_paid_detail"`
}

type Exam struct {
	QUuid     uuid.UUID             `json:"q_uuid"`
	EType     sql.NullInt32         `json:"e_type"`
	ELevel    sql.NullInt32         `json:"e_level"`
	EName     sql.NullString        `json:"e_name"`
	QList     pqtype.NullRawMessage `json:"q_list"`
	CreatedAt sql.NullTime          `json:"created_at"`
}

type Hiree struct {
	HUuid       uuid.UUID             `json:"h_uuid"`
	HUUuid      uuid.NullUUID         `json:"h_u_uuid"`
	HDetails    pqtype.NullRawMessage `json:"h_details"`
	HJobDetails pqtype.NullRawMessage `json:"h_job_details"`
	HSkills     pqtype.NullRawMessage `json:"h_skills"`
}

type Job struct {
	JUuid     uuid.UUID             `json:"j_uuid"`
	JobName   sql.NullString        `json:"job_name"`
	JobType   sql.NullInt32         `json:"job_type"`
	JobStatus sql.NullInt32         `json:"job_status"`
	JobDesc   pqtype.NullRawMessage `json:"job_desc"`
	JobSkill  pqtype.NullRawMessage `json:"job_skill"`
}

type Organization struct {
	OUuid             uuid.UUID             `json:"o_uuid"`
	CountryCode       sql.NullInt32         `json:"country_code"`
	OrgName           string                `json:"org_name"`
	OrgType           sql.NullString        `json:"org_type"`
	OrgCat            sql.NullInt32         `json:"org_cat"`
	JobList           pqtype.NullRawMessage `json:"job_list"`
	AdminID           uuid.NullUUID         `json:"admin_id"`
	OrgBillingDetails pqtype.NullRawMessage `json:"org_billing_details"`
	LastBilling       uuid.NullUUID         `json:"last_billing"`
}

type Question struct {
	QUuid     uuid.UUID       `json:"q_uuid"`
	QType     int32           `json:"q_type"`
	QLevel    int32           `json:"q_level"`
	Question  json.RawMessage `json:"question"`
	CreatedAt sql.NullTime    `json:"created_at"`
}

type Skill struct {
	SUuid     uuid.UUID             `json:"s_uuid"`
	SkillName sql.NullString        `json:"skill_name"`
	SkillType sql.NullInt32         `json:"skill_type"`
	ExamList  pqtype.NullRawMessage `json:"exam_list"`
	CreatedAt sql.NullTime          `json:"created_at"`
}

type User struct {
	UUuid      uuid.UUID             `json:"u_uuid"`
	UFirstName string                `json:"u_first_name"`
	ULastName  string                `json:"u_last_name"`
	UType      int32                 `json:"u_type"`
	UOrg       uuid.NullUUID         `json:"u_org"`
	UDetail    pqtype.NullRawMessage `json:"u_detail"`
}

type Usersgroup struct {
	UgUuid uuid.UUID             `json:"ug_uuid"`
	UgName sql.NullString        `json:"ug_name"`
	UgType sql.NullInt32         `json:"ug_type"`
	UgOrg  uuid.NullUUID         `json:"ug_org"`
	UgList pqtype.NullRawMessage `json:"ug_list"`
}
