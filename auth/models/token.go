package models

type TokenIntrospection struct {
	Active     bool        `json:"active"`
	Username   string      `json:"username"`
	RealmRoles RealmAccess `json:"realm_access"`
	ExpiresAt  int64       `json:"exp"`
	IssuedAt   int64       `json:"iat"`
}

type RealmAccess struct {
	Roles []string `json:"roles"`
}
