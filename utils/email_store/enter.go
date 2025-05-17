package email_store

import "sync"

type EmailStoreInfo struct {
	Email string
	Code  string
}

var emailVerifyStore = sync.Map{}

func Set(id, email, code string) {
	emailVerifyStore.Store(id, EmailStoreInfo{
		Email: email,
		Code:  code,
	})
}

func Verify(id, code string) (info EmailStoreInfo, ok bool) {
	val, ok := emailVerifyStore.Load(id)
	if !ok {
		return
	}
	info, ok = val.(EmailStoreInfo)
	if !ok {
		return
	}
	if info.Code != code {
		emailVerifyStore.Delete(id)
		ok = false
		return
	}
	emailVerifyStore.Delete(id)
	return
}
