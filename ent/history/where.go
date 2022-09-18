// Code generated by ent, DO NOT EDIT.

package history

import (
	"mofu/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatorID), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// LastUpdate applies equality check predicate on the "last_update" field. It's identical to LastUpdateEQ.
func LastUpdate(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdate), v))
	})
}

// ContentFlag applies equality check predicate on the "content_flag" field. It's identical to ContentFlagEQ.
func ContentFlag(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentFlag), v))
	})
}

// SentFlag applies equality check predicate on the "sent_flag" field. It's identical to SentFlagEQ.
func SentFlag(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSentFlag), v))
	})
}

// MentionedCount applies equality check predicate on the "mentioned_count" field. It's identical to MentionedCountEQ.
func MentionedCount(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMentionedCount), v))
	})
}

// SendingContent applies equality check predicate on the "sending_content" field. It's identical to SendingContentEQ.
func SendingContent(v []byte) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSendingContent), v))
	})
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatorID), v))
	})
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatorID), v))
	})
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...string) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatorID), v...))
	})
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...string) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatorID), v...))
	})
}

// CreatorIDGT applies the GT predicate on the "creator_id" field.
func CreatorIDGT(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatorID), v))
	})
}

// CreatorIDGTE applies the GTE predicate on the "creator_id" field.
func CreatorIDGTE(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatorID), v))
	})
}

// CreatorIDLT applies the LT predicate on the "creator_id" field.
func CreatorIDLT(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatorID), v))
	})
}

// CreatorIDLTE applies the LTE predicate on the "creator_id" field.
func CreatorIDLTE(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatorID), v))
	})
}

// CreatorIDContains applies the Contains predicate on the "creator_id" field.
func CreatorIDContains(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreatorID), v))
	})
}

// CreatorIDHasPrefix applies the HasPrefix predicate on the "creator_id" field.
func CreatorIDHasPrefix(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreatorID), v))
	})
}

// CreatorIDHasSuffix applies the HasSuffix predicate on the "creator_id" field.
func CreatorIDHasSuffix(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreatorID), v))
	})
}

// CreatorIDEqualFold applies the EqualFold predicate on the "creator_id" field.
func CreatorIDEqualFold(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreatorID), v))
	})
}

// CreatorIDContainsFold applies the ContainsFold predicate on the "creator_id" field.
func CreatorIDContainsFold(v string) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreatorID), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// LastUpdateEQ applies the EQ predicate on the "last_update" field.
func LastUpdateEQ(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateNEQ applies the NEQ predicate on the "last_update" field.
func LastUpdateNEQ(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateIn applies the In predicate on the "last_update" field.
func LastUpdateIn(vs ...time.Time) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastUpdate), v...))
	})
}

// LastUpdateNotIn applies the NotIn predicate on the "last_update" field.
func LastUpdateNotIn(vs ...time.Time) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastUpdate), v...))
	})
}

// LastUpdateGT applies the GT predicate on the "last_update" field.
func LastUpdateGT(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateGTE applies the GTE predicate on the "last_update" field.
func LastUpdateGTE(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateLT applies the LT predicate on the "last_update" field.
func LastUpdateLT(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateLTE applies the LTE predicate on the "last_update" field.
func LastUpdateLTE(v time.Time) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastUpdate), v))
	})
}

// ContentFlagEQ applies the EQ predicate on the "content_flag" field.
func ContentFlagEQ(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentFlag), v))
	})
}

// ContentFlagNEQ applies the NEQ predicate on the "content_flag" field.
func ContentFlagNEQ(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContentFlag), v))
	})
}

// ContentFlagIn applies the In predicate on the "content_flag" field.
func ContentFlagIn(vs ...int) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContentFlag), v...))
	})
}

// ContentFlagNotIn applies the NotIn predicate on the "content_flag" field.
func ContentFlagNotIn(vs ...int) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContentFlag), v...))
	})
}

// ContentFlagGT applies the GT predicate on the "content_flag" field.
func ContentFlagGT(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContentFlag), v))
	})
}

// ContentFlagGTE applies the GTE predicate on the "content_flag" field.
func ContentFlagGTE(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContentFlag), v))
	})
}

// ContentFlagLT applies the LT predicate on the "content_flag" field.
func ContentFlagLT(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContentFlag), v))
	})
}

// ContentFlagLTE applies the LTE predicate on the "content_flag" field.
func ContentFlagLTE(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContentFlag), v))
	})
}

// ContentFlagIsNil applies the IsNil predicate on the "content_flag" field.
func ContentFlagIsNil() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContentFlag)))
	})
}

// ContentFlagNotNil applies the NotNil predicate on the "content_flag" field.
func ContentFlagNotNil() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContentFlag)))
	})
}

// SentFlagEQ applies the EQ predicate on the "sent_flag" field.
func SentFlagEQ(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSentFlag), v))
	})
}

// SentFlagNEQ applies the NEQ predicate on the "sent_flag" field.
func SentFlagNEQ(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSentFlag), v))
	})
}

// SentFlagIn applies the In predicate on the "sent_flag" field.
func SentFlagIn(vs ...int) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSentFlag), v...))
	})
}

// SentFlagNotIn applies the NotIn predicate on the "sent_flag" field.
func SentFlagNotIn(vs ...int) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSentFlag), v...))
	})
}

// SentFlagGT applies the GT predicate on the "sent_flag" field.
func SentFlagGT(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSentFlag), v))
	})
}

// SentFlagGTE applies the GTE predicate on the "sent_flag" field.
func SentFlagGTE(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSentFlag), v))
	})
}

// SentFlagLT applies the LT predicate on the "sent_flag" field.
func SentFlagLT(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSentFlag), v))
	})
}

// SentFlagLTE applies the LTE predicate on the "sent_flag" field.
func SentFlagLTE(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSentFlag), v))
	})
}

// SentFlagIsNil applies the IsNil predicate on the "sent_flag" field.
func SentFlagIsNil() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSentFlag)))
	})
}

// SentFlagNotNil applies the NotNil predicate on the "sent_flag" field.
func SentFlagNotNil() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSentFlag)))
	})
}

// MentionedCountEQ applies the EQ predicate on the "mentioned_count" field.
func MentionedCountEQ(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMentionedCount), v))
	})
}

// MentionedCountNEQ applies the NEQ predicate on the "mentioned_count" field.
func MentionedCountNEQ(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMentionedCount), v))
	})
}

// MentionedCountIn applies the In predicate on the "mentioned_count" field.
func MentionedCountIn(vs ...int) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMentionedCount), v...))
	})
}

// MentionedCountNotIn applies the NotIn predicate on the "mentioned_count" field.
func MentionedCountNotIn(vs ...int) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMentionedCount), v...))
	})
}

// MentionedCountGT applies the GT predicate on the "mentioned_count" field.
func MentionedCountGT(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMentionedCount), v))
	})
}

// MentionedCountGTE applies the GTE predicate on the "mentioned_count" field.
func MentionedCountGTE(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMentionedCount), v))
	})
}

// MentionedCountLT applies the LT predicate on the "mentioned_count" field.
func MentionedCountLT(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMentionedCount), v))
	})
}

// MentionedCountLTE applies the LTE predicate on the "mentioned_count" field.
func MentionedCountLTE(v int) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMentionedCount), v))
	})
}

// SendingContentEQ applies the EQ predicate on the "sending_content" field.
func SendingContentEQ(v []byte) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSendingContent), v))
	})
}

// SendingContentNEQ applies the NEQ predicate on the "sending_content" field.
func SendingContentNEQ(v []byte) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSendingContent), v))
	})
}

// SendingContentIn applies the In predicate on the "sending_content" field.
func SendingContentIn(vs ...[]byte) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSendingContent), v...))
	})
}

// SendingContentNotIn applies the NotIn predicate on the "sending_content" field.
func SendingContentNotIn(vs ...[]byte) predicate.History {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSendingContent), v...))
	})
}

// SendingContentGT applies the GT predicate on the "sending_content" field.
func SendingContentGT(v []byte) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSendingContent), v))
	})
}

// SendingContentGTE applies the GTE predicate on the "sending_content" field.
func SendingContentGTE(v []byte) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSendingContent), v))
	})
}

// SendingContentLT applies the LT predicate on the "sending_content" field.
func SendingContentLT(v []byte) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSendingContent), v))
	})
}

// SendingContentLTE applies the LTE predicate on the "sending_content" field.
func SendingContentLTE(v []byte) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSendingContent), v))
	})
}

// SendingContentIsNil applies the IsNil predicate on the "sending_content" field.
func SendingContentIsNil() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSendingContent)))
	})
}

// SendingContentNotNil applies the NotNil predicate on the "sending_content" field.
func SendingContentNotNil() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSendingContent)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.History) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		p(s.Not())
	})
}
