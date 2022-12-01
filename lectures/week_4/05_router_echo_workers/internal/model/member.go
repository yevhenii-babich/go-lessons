package model

import (
	"echo-with-workers/internal/db"
	"encoding/json"
	"sort"
	"time"
)

type (
	MyDate int64
	Member struct {
		Number    int64  `json:"number"`
		Name      string `json:"name"`
		Position  string `json:"position"`
		CreatedAt MyDate `json:"createdAt"`
	}
)

func NewMember(member int64, name, position string) *Member {
	return &Member{
		Number:    member,
		Name:      name,
		Position:  position,
		CreatedAt: MyDate(time.Now().Unix()),
	}
}

func (m *Member) Save(tx *db.JsonDB[Member]) error {
	err := tx.InsertInto(m, m.Number, &m.Number)
	return err
}

func (m *Member) Load(tx *db.JsonDB[Member], number int64) error {
	return tx.Select(number, m)
}

type Members []Member

func (m *Members) Load(tx *db.JsonDB[Member], order string) error {
	records := tx.GetAll()
	sort.Slice(records, func(i, j int) bool {
		switch order {
		case "desc":
			i, j = j, i
		case "asc":
		default:

		}
		return records[i].Number < records[j].Number
	})
	*m = records
	return nil
}

func (m MyDate) MarshalJSON() ([]byte, error) {
	dt := time.Unix(int64(m), 0)
	out := `"` + dt.Format(time.RFC822) + `"`
	return []byte(out), nil
}

func (m *MyDate) UnmarshalJSON(in []byte) error {
	var s string
	err := json.Unmarshal(in, &s)
	if err != nil {
		return err
	}
	t, err := time.Parse(time.RFC822, s)
	if err != nil {
		return err
	}
	*m = MyDate(t.Unix())
	return nil
}
