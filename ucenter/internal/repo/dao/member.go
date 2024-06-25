package dao

import (
	"context"
	"gorm.io/gorm"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
	"ucenter/internal/model"
)

type MemberDAO struct {
	conn *gorms.GormConn
}

func NewMemberDAO(db *msdb.MsDB) *MemberDAO {
	return &MemberDAO{
		conn: gorms.New(db.Conn),
	}
}

func (m *MemberDAO) FindMemberById(ctx context.Context, memberId int64) (*model.Member, error) {
	session := m.conn.Session(ctx)
	var mem model.Member
	err := session.Model(&model.Member{}).
		Where("id = ?", memberId).
		First(&mem).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &mem, err
}

func (m *MemberDAO) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	session := m.conn.Session(ctx)
	var mem model.Member
	err := session.Model(&model.Member{}).
		Where("mobile_phone = ?", phone).
		First(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &mem, err
}

func (m *MemberDAO) InsertMember(ctx context.Context, mem *model.Member) error {
	session := m.conn.Session(ctx)
	return session.Create(mem).Error
}

func (m *MemberDAO) UpdateLoginCount(ctx context.Context, id int64, step int) error {
	session := m.conn.Session(ctx)
	//login_count = login_count+1
	err := session.Exec("update member set login_count = login_count+? where id=?", step, id).Error
	return err
}
