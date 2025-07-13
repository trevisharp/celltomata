package models

type WorldRule struct {
	Commands []any
}

type ConstantDefinitionExp struct {
	NodeType     string
	ConstantName string
	Value        int
}

type StateDefinitionExp struct {
	NodeType  string
	StateName string
	IsDefault bool
}

type OperationExp struct {
	NodeType      string
	OperationName string
	LeftValue     string
	Operator      string
	RightValue    string
}

type ConditionExp struct {
	NodeType        string
	ConditionName   string
	LeftRestriction *RestrictionExp
	Operator        string
	Rest            *ConditionExp
}

type RestrictionExp struct {
	NodeType        string
	RestrictionName string
	LeftOperator    string
	Value           string
	RightOperator   string
}

type RestrictionListExp struct {
	NodeType            string
	RestrictionListName string
	Restrictions        []RestrictionExp
}

type NeighborhoodExp struct {
	NodeType         string
	NeighborhoodName string
	XRestriction     string
	YRestriction     string
}

type QueryExp struct {
	NodeType          string
	QueryName         string
	AggregateFunction string
	Value             *any
}

type MatchCaseExp struct {
	NodeType  string
	Variables []string
	Case      []string
	Become    string
}
