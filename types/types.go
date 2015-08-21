package types

import (
	"github.com/chris-ramon/graphql-go/errors"
	"github.com/chris-ramon/graphql-go/language/ast"
)

type Schema interface{}

type GraphQLResult struct {
	Data   interface{}
	Errors []errors.GraphQLFormattedError
}

func (gqR *GraphQLResult) HasErrors() bool {
	return (len(gqR.Errors) > 0)
}

type GraphQLEnumType struct {
}

type GraphQLInterfaceTypeConfig struct {
	Name        string
	Fields      interface{}
	ResolveType GraphQLObjectType
	Description string
}

type GraphQLInterfaceType struct {
	Name              string
	Description       string
	TypeConfig        GraphQLInterfaceTypeConfig
	Fields            GraphQLFieldDefinitionMap
	Implementations   []GraphQLObjectType
	PossibleTypeNames map[string]bool
}

func (it *GraphQLInterfaceType) GetName() string {
	return it.Name
}
func (it *GraphQLInterfaceType) GetDescription() string {
	return it.Description
}
func (it *GraphQLInterfaceType) Coerce(value interface{}) (r interface{}) {
	return r
}
func (it *GraphQLInterfaceType) CoerceLiteral(value interface{}) (r interface{}) {
	return r
}

func (it *GraphQLInterfaceType) Constructor(config GraphQLInterfaceTypeConfig) {
	//TODO: figure out how to make next line work
	//invariant(config.name, 'Type must be named.');
	it.Name = config.Name
	it.Description = config.Description
	it.TypeConfig = config
	//it.Implementations = []GraphQLObjectType;
}

func (it *GraphQLInterfaceType) GetFields() (fields GraphQLFieldDefinitionMap) {
	//return this._fields ||
	//(this._fields = defineFieldMap(this._typeConfig.fields));
	return fields
}

func (it *GraphQLInterfaceType) GetPossibleTypes() []GraphQLObjectType {
	return it.Implementations
}

func (it *GraphQLInterfaceType) IsPossibleType(objType GraphQLObjectType) bool {
	//var possibleTypeNames = this._possibleTypeNames;
	//if (!possibleTypeNames) {
	//this._possibleTypeNames = possibleTypeNames =
	//this.getPossibleTypes().reduce(
	//(map, possibleType) => ((map[possibleType.name] = true), map),
	//{}
	//);
	//}
	//return possibleTypeNames[type.name] === true;
	return true
}

func (it *GraphQLInterfaceType) ResolveType(any interface{}) (r GraphQLObjectType) {
	//var resolver = this._typeConfig.resolveType;
	//return resolver ? resolver(value) : getTypeOf(value, this);
	return r
}

func (it *GraphQLInterfaceType) ToString() string {
	return it.Name
}

type GQLFDRParams struct {
	Source     interface{}
	Args       map[string]interface{}
	Context    interface{}
	FieldAST   interface{}
	FieldType  interface{}
	ParentType interface{}
	Schema     GraphQLSchema
}

type GraphQLFieldDefinitionResolve func(p GQLFDRParams) interface{}

type GraphQLScalarTypeConfig struct {
	Name        string
	Description string
}

func (stC *GraphQLScalarTypeConfig) Coerce(value interface{}) (r interface{}) {
	return r
}

func (stC *GraphQLScalarTypeConfig) CoerceLiteral(value ast.Value) (r interface{}) {
	return r
}

type GraphQLScalarType struct {
	Name         string
	Description  string
	ScalarConfig GraphQLScalarTypeConfig
}

func (sT *GraphQLScalarType) GetName() string {
	return sT.Name

}
func (sT *GraphQLScalarType) GetDescription() string {
	return sT.Description

}
func (sT *GraphQLScalarType) Coerce(value interface{}) (r interface{}) {
	return r

}
func (sT *GraphQLScalarType) CoerceLiteral(value interface{}) (r interface{}) {
	return r

}
func (sT *GraphQLScalarType) ToString() string {
	return sT.Name
}

type GraphQLOutputType interface {
	GetName() string
	GetDescription() string
	Coerce(value interface{}) (r interface{})
	CoerceLiteral(value interface{}) (r interface{})
	ToString() string
}

type GraphQLInputType interface {
	GetName() string
	GetDescription() string
	Coerce(value interface{}) interface{}
	CoerceLiteral(value interface{}) interface{}
	ToString() string
}

type GraphQLFieldArgument struct {
	Name         string
	Type         GraphQLInputType
	DefaultValue interface{}
	Description  string
}

type GraphQLFieldDefinition struct {
	Name              string
	Description       string
	Type              GraphQLOutputType
	Args              []GraphQLFieldArgument
	Resolve           GraphQLFieldDefinitionResolve
	DeprecationReason string
}

type GraphQLFieldDefinitionMap map[string]GraphQLFieldDefinition

type GraphQLObjectType struct {
	Name   string
	Fields GraphQLFieldDefinitionMap
}

type GraphQLList struct {
}

type GraphQLNonNull struct {
}

type GraphQLSchemaConfig struct {
	Query    GraphQLObjectType
	Mutation GraphQLObjectType
}

type GraphQLSchema struct {
	Query        GraphQLObjectType
	SchemaConfig GraphQLSchemaConfig
}

func (gq *GraphQLSchema) Constructor(config GraphQLSchemaConfig) {
	gq.SchemaConfig = config
}

func (gq *GraphQLSchema) GetQueryType() GraphQLObjectType {
	return gq.SchemaConfig.Query
}

func (gq *GraphQLSchema) GetMutationType() GraphQLObjectType {
	return gq.SchemaConfig.Mutation
}

type GraphQLString struct {
}