package models

type RepositoryOwner struct {
	Id             string
	Type           string
	Organization   Organization
	RepositoryName string
	Pattern        string
	Owner          string
	ParentOwner    string
}