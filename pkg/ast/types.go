// Copyright 2016 Marapongo, Inc. All rights reserved.

// This package contains the core Mu abstract syntax tree types.
//
// N.B. for the time being, we are leveraging the same set of types for parse trees and abstract syntax trees.  The
// reason is that minimal "extra" information is necessary between front- and back-end parts of the compiler, and so
// reusing the trees leads to less duplication in types and faster runtime performance.  As the compiler matures in
// functionality, we may want to revisit this.  The "back-end-only" parts of the data structures are easily identified
// because their fields do not map to any serializable fields (i.e., `json:"-"`).
//
// Another controversial decision is to mutate nodes in place, rather than taking the performance hit of immutability.
// This can certainly be tricky to deal with, however, it is simpler and we can revisit it down the road if needed.
// Of course, during lowering, sometimes nodes will be transformed to new types entirely, allocating entirely anew.
package ast

// Name is an identifier.  Names may be optionally fully qualified, using the delimiter `/`, or simple.  Each element
// conforms to the regex [A-Za-z_][A-Za-z0-9_]*.  For example, `marapongo/mu/stack`.
type Name string

// Ref is a dependency reference.  It is "name-like", in that it contains a Name embedded inside of it, but also carries
// a URL-like structure.  A Ref starts with an optional "protocol" (like https://, git://, etc), followed by an optional
// "base" part (like hub.mu.com/, github.com/, etc), followed by the "name" part (which is just a Name).
type Ref string

// SemVer represents a version using "semantic versioning" style.  This may include up to three distinct numbers
// delimited by `.`s: a major version number, a minor version number, and a revision number.  For example, `1.0.10`.
type SemVer string

// Node is the base of all abstract syntax tree types.
type Node struct {
}

// Workspace defines settings shared amongst many related Stacks.
type Workspace struct {
	Node

	Clusters Clusters `json:"clusters,omitempty"` // an optional set of predefined target clusters.
}

// Stack represents a collection of private and public cloud resources, a method for constructing them, and optional
// dependencies on other Stacks (by name).
type Stack struct {
	Node

	Name        Name     `json:"name,omitempty"`        // a friendly name for this node.
	Version     SemVer   `json:"version,omitempty"`     // a specific semantic version.
	Description string   `json:"description,omitempty"` // an optional friendly description.
	Author      string   `json:"author,omitempty"`      // an optional author.
	Website     string   `json:"website,omitempty"`     // an optional website for additional info.
	License     string   `json:"license,omitempty"`     // an optional license governing legal uses of this package.
	Clusters    Clusters `json:"clusters,omitempty"`    // an optional set of predefined target clusters.

	Base         Name         `json:"base,omitempty"`     // an optional base Stack type.
	Abstract     bool         `json:"abstract,omitempty"` // true if this stack is "abstract" (uninstantiable).
	Properties   Properties   `json:"properties,omitempty"`
	Dependencies Dependencies `json:"dependencies,omitempty"`
	Services     Services     `json:"services,omitempty"`

	BoundBase         *Stack            `json:"-"` // base, if available, is bound during semantic analysis.
	BoundDependencies BoundDependencies `json:"-"` // dependencies are bound during semantic analysis.
}

// Clusters is a map of target names to metadata about those targets.
type Clusters map[string]Cluster

// Cluster describes a predefined cloud runtime target, including its OS and Scheduler combination.
type Cluster struct {
	Node

	Default     bool                   `json:"default,omitempty"`     // a single target can carry default settings.
	Description string                 `json:"description,omitempty"` // a human-friendly description of this target.
	Cloud       string                 `json:"cloud,omitempty"`       // the cloud target.
	Scheduler   string                 `json:"scheduler,omitempty"`   // the cloud scheduler target.
	Options     map[string]interface{} `json:"options,omitempty"`     // any options passed to the cloud provider.

	Name string `json:"-"` // name is decorated post-parsing, since it is contextual.
}

// Propertys maps property names to metadata about those propertys.
type Properties map[string]Property

// Property describes the requirements of arguments used when constructing Stacks, etc.
type Property struct {
	Node

	Type        PropertyType `json:"type,omitempty"`        // the type of the property; required.
	Description string       `json:"description,omitempty"` // an optional friendly description of the property.
	Default     interface{}  `json:"default,omitempty"`     // an optional default value if the caller elides one.
	Optional    bool         `json:"optional,omitempty"`    // true if may be omitted (inferred if a default value).

	Name string `json:"-"` // name is decorated post-parsing, since it is contextual.
}

// PropertyType stores the name of a property's type.
type PropertyType Name

// A set of known property types.  Note that this is extensible, so names outside of this list are legal.
// TODO: support complex types (like arrays, custom JSON shapes, and so on).
const (
	PropertyTypeAny     PropertyType = "any"     // any structure.
	PropertyTypeString               = "string"  // a JSON-like string.
	PropertyTypeNumber               = "number"  // a JSON-like number (integer or floating point).
	PropertyTypeBoolean              = "boolean" // a JSON-like boolean (`true` or `false`).
	PropertyTypeService              = "service" // an untyped service reference; the runtime manifestation is a URL.
)

// Dependencies maps dependency names to the semantic version the consumer depends on.
type Dependencies map[Ref]Dependency

// Dependency is metadata describing a dependency target (for now, just its semantic version).
type Dependency SemVer

// BoundDependencies contains a map of all bound dependencies, populated during semantic analysis.
type BoundDependencies map[Ref]BoundDependency

// BoundDependency contains information about a binding.
type BoundDependency struct {
	Ref     Ref    // the reference used to bind to this dependency.
	Version SemVer // the version requested to bind to this dependency.
	Stack   *Stack // the bound stack for this dependency.
}

// Services is a list of public and private service references, keyed by name.
type Services struct {
	// These fields are expanded after parsing:
	Public  ServiceMap `json:"-"`
	Private ServiceMap `json:"-"`

	// These fields are "untyped" due to limitations in the JSON parser.  Namely, Go's parser will ignore
	// properties in the payload that it doesn't recognize as mapping to a field.  That's not what we want, especially
	// for services since they are highly extensible and the contents will differ per-type.  Therefore, we will first
	// map the services into a weakly typed map, and later on during compilation, expand them to the below fields.
	// TODO[marapongo/mu#4]: support for `json:",inline"` or the equivalent so we can eliminate these fields.
	PublicUntyped  UntypedServiceMap `json:"public,omitempty"`
	PrivateUntyped UntypedServiceMap `json:"private,omitempty"`
}

// ServiceMap is a map of service names to metadata about those services.
type ServiceMap map[Name]Service

// Service is a directive for instantiating another Stack, including its name, arguments, etc.
type Service struct {
	Node

	Type  Name                   `json:"type,omitempty"` // an explicit type; if missing, the name is used.
	Extra map[string]interface{} `json:"-"`              // all of the "extra" properties, other than what is above.

	Name   Name `json:"-"` // a friendly name; decorated post-parsing, since it is contextual.
	Public bool `json:"-"` // true if this service is publicly exposed; also decorated post-parsing.

	BoundType *Stack `json:"-"` // services are bound to stacks during semantic analysis.
}

// UntypedServiceMap is a map of service names to untyped, bags of parsed properties for those services.
type UntypedServiceMap map[Name]PropertyBag

// PropertyBag is simply a map of string property names to untyped data values.
type PropertyBag map[string]interface{}

// TODO: several more core types still need to be mapped:
//     - Schema
//     - Identity: User, Role, Group
//     - Configuration
//     - Secret
