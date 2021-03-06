package variables

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/earthly/earthly/domain"
	"github.com/earthly/earthly/gitutil"
	"github.com/earthly/earthly/llbutil"
	"github.com/earthly/earthly/stringutil"

	"github.com/containerd/containerd/platforms"
	dfShell "github.com/moby/buildkit/frontend/dockerfile/shell"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
)

// ProcessNonConstantVariableFunc is a function which takes in an expression and
// turns it into a state, target intput and arg index.
type ProcessNonConstantVariableFunc func(name string, expression string) (value string, argIndex int, err error)

// Collection is a collection of variables.
type Collection struct {
	variables map[string]Variable
	// activeVariables are variables that are active right now as we have passed the point of
	// their declaration.
	activeVariables map[string]bool
	// overridingVariables represent variables that should be passed in deep to override.
	overridingVariables map[string]bool
	// globalVariables represent variables that are passed to all targets in a file.
	globalVariables map[string]bool
}

// NewCollection returns a new collection.
func NewCollection() *Collection {
	return &Collection{
		variables:           make(map[string]Variable),
		activeVariables:     make(map[string]bool),
		overridingVariables: make(map[string]bool),
		globalVariables:     make(map[string]bool),
	}
}

// ParseCommandLineBuildArgs parses a slice of constant build args and returns a new collection.
func ParseCommandLineBuildArgs(args []string, dotEnvMap map[string]string) (*Collection, error) {
	ret := NewCollection()
	for k, v := range dotEnvMap {
		ret.variables[k] = NewConstant(v)
	}
	for _, arg := range args {
		splitArg := strings.SplitN(arg, "=", 2)
		if len(splitArg) < 1 {
			return nil, fmt.Errorf("invalid build arg %s", splitArg)
		}
		key := splitArg[0]
		value := ""
		hasValue := false
		if len(splitArg) == 2 {
			value = splitArg[1]
			hasValue = true
		}
		if !hasValue {
			var found bool
			value, found = os.LookupEnv(key)
			if !found {
				return nil, fmt.Errorf("env var %s not set", key)
			}
		}
		ret.variables[key] = NewConstant(value)
		ret.overridingVariables[key] = true
	}
	return ret, nil
}

// Get returns a variable by name.
func (c *Collection) Get(name string) (Variable, bool, bool) {
	variable, found := c.variables[name]
	active := false
	if found {
		active = c.activeVariables[name]
	}
	return variable, active, found
}

// Expand expands constant build args within the given word.
func (c *Collection) Expand(word string) string {
	shlex := dfShell.NewLex('\\')
	argsMap := make(map[string]string)
	for varName := range c.activeVariables {
		variable := c.variables[varName]
		argsMap[varName] = variable.ConstantValue()
	}
	ret, err := shlex.ProcessWordWithMap(word, argsMap)
	if err != nil {
		// No effect if there is an error.
		return word
	}
	return ret
}

// AsMap returns the constant variables (active and inactive) as a map.
func (c *Collection) AsMap() map[string]string {
	ret := make(map[string]string)
	for varName, variable := range c.variables {
		ret[varName] = variable.ConstantValue()
	}
	return ret
}

// SortedActiveVariables returns the active variable names in a sorted slice.
func (c *Collection) SortedActiveVariables() []string {
	varNames := make([]string, 0, len(c.activeVariables))
	for varName := range c.activeVariables {
		varNames = append(varNames, varName)
	}
	sort.Strings(varNames)
	return varNames
}

// SortedOverridingVariables returns the overriding variable names in a sorted slice.
func (c *Collection) SortedOverridingVariables() []string {
	varNames := make([]string, 0, len(c.overridingVariables))
	for varName := range c.overridingVariables {
		varNames = append(varNames, varName)
	}
	sort.Strings(varNames)
	return varNames
}

// AddActive adds and activates a variable in the collection.
func (c *Collection) AddActive(name string, variable Variable) {
	c.activeVariables[name] = true
	c.variables[name] = variable
}

// DeclareActive declares a variable and sets it to active state. The effective value may be
// different than the default, if the variable has been overridden.
func (c *Collection) DeclareActive(name string, defaultValue string, global bool, pncvf ProcessNonConstantVariableFunc) (Variable, error) {
	var effective Variable
	existing, found := c.variables[name]
	if found {
		effective = existing
	} else {
		v, err := c.parseBuildArgValue(name, defaultValue, pncvf)
		if err != nil {
			return Variable{}, err
		}
		effective = v
		c.variables[name] = v
	}
	if global {
		c.globalVariables[name] = true
	}
	c.activeVariables[name] = true
	return effective, nil
}

// WithResetEnvVars returns a copy of the current collection with all env vars
// removed. This operation does not modify the current collection.
func (c *Collection) WithResetEnvVars() *Collection {
	ret := NewCollection()
	for k, v := range c.variables {
		if !v.IsEnvVar() {
			ret.variables[k] = v
			if c.activeVariables[k] {
				ret.activeVariables[k] = true
			}
		}
	}
	for k := range c.overridingVariables {
		ret.overridingVariables[k] = true
	}
	for k := range c.globalVariables {
		ret.globalVariables[k] = true
	}
	return ret
}

// WithOnlyGlobals returns a copy of the current collection, keeping only the global variables.
func (c *Collection) WithOnlyGlobals() *Collection {
	ret := NewCollection()
	for k := range c.globalVariables {
		ret.globalVariables[k] = true
		ret.activeVariables[k] = true
		ret.variables[k] = c.variables[k]
	}
	return ret
}

// getProjectName returns the depricated PROJECT_NAME value
func getProjectName(s string) string {
	protocol := "unknown"
	parts := strings.SplitN(s, "://", 2)
	if len(parts) > 1 {
		protocol = parts[0]
		s = parts[1]
	}
	parts = strings.SplitN(s, "@", 2)
	if len(parts) > 1 {
		s = parts[1]
	}
	if protocol == "unknown" {
		s = strings.Replace(s, ":", "/", 1)
	}
	s = strings.TrimSuffix(s, ".git")
	parts = strings.SplitN(s, "/", 2)
	if len(parts) > 1 {
		s = parts[1]
	}
	return s
}

// WithBuiltinBuildArgs returns a new collection containing the current variables together with
// builtin args. This operation does not modify the current collection.
func (c *Collection) WithBuiltinBuildArgs(target domain.Target, platform specs.Platform, gitMeta *gitutil.GitMetadata) *Collection {
	ret := NewCollection()
	// Copy existing variables.
	for k, v := range c.variables {
		ret.variables[k] = v
	}
	for k := range c.overridingVariables {
		ret.overridingVariables[k] = true
	}
	for k := range c.globalVariables {
		ret.globalVariables[k] = true
		ret.activeVariables[k] = true
	}
	// Add the builtin build args.
	ret.variables["EARTHLY_TARGET"] = NewConstant(target.StringCanonical())
	ret.variables["EARTHLY_TARGET_PROJECT"] = NewConstant(target.ProjectCanonical())
	ret.variables["EARTHLY_TARGET_NAME"] = NewConstant(target.Target)
	ret.variables["EARTHLY_TARGET_TAG"] = NewConstant(target.Tag)
	ret.variables["EARTHLY_TARGET_TAG_DOCKER"] = NewConstant(llbutil.DockerTagSafe(target.Tag))
	ret.SetPlatformArgs(platform)

	if gitMeta != nil {
		ret.variables["EARTHLY_GIT_HASH"] = NewConstant(gitMeta.Hash)
		branch := ""
		if len(gitMeta.Branch) > 0 {
			branch = gitMeta.Branch[0]
		}
		ret.variables["EARTHLY_GIT_BRANCH"] = NewConstant(branch)
		tag := ""
		if len(gitMeta.Tags) > 0 {
			tag = gitMeta.Tags[0]
		}
		ret.variables["EARTHLY_GIT_TAG"] = NewConstant(tag)
		ret.variables["EARTHLY_GIT_ORIGIN_URL"] = NewConstant(gitMeta.RemoteURL)
		ret.variables["EARTHLY_GIT_ORIGIN_URL_SCRUBBED"] = NewConstant(stringutil.ScrubCredentials(gitMeta.RemoteURL))
		ret.variables["EARTHLY_GIT_PROJECT_NAME"] = NewConstant(getProjectName(gitMeta.RemoteURL))
	}
	return ret
}

// SetPlatformArgs set the platform-specific built-in args to a specific platform.
func (c *Collection) SetPlatformArgs(platform specs.Platform) {
	c.variables["TARGETPLATFORM"] = NewConstant(platforms.Format(platform))
	c.variables["TARGETOS"] = NewConstant(platform.OS)
	c.variables["TARGETARCH"] = NewConstant(platform.Architecture)
	c.variables["TARGETVARIANT"] = NewConstant(platform.Variant)
}

// WithParseBuildArgs takes in a slice of build args to be parsed and returns another collection
// containing the current build args, together with the newly parsed build args. This operation does
// not modify the current collection.
func (c *Collection) WithParseBuildArgs(args []string, pncvf ProcessNonConstantVariableFunc, propagate bool) (*Collection, map[string]bool, error) {
	// First, parse.
	toAdd := make(map[string]Variable)
	haveValues := make(map[string]bool)
	for _, arg := range args {
		name, variable, hasValue, err := c.parseBuildArg(arg, pncvf)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "parse build arg %s", arg)
		}
		toAdd[name] = variable
		if hasValue {
			haveValues[name] = true
		}
	}

	// Merge into a new collection.
	// Copy existing, without env vars.
	var newC *Collection
	if propagate {
		newC = c.WithResetEnvVars()
	} else {
		newC = NewCollection()
	}
	newVars := make(map[string]bool)
	// Add the parsed ones too.
	for key, ba := range toAdd {
		if ba.IsEnvVar() {
			continue
		}
		var finalValue Variable
		if !haveValues[key] {
			existing, active, found := c.Get(key)
			if found && active {
				if existing.IsEnvVar() {
					finalValue = NewConstant(existing.ConstantValue())
				} else {
					finalValue = existing
				}
			} else {
				return nil, nil, fmt.Errorf(
					"value not specified for build arg %s and no value can be inferred", key)
			}
		} else {
			finalValue = ba
		}
		newVars[key] = true
		newC.variables[key] = finalValue
		newC.overridingVariables[key] = true
	}
	return newC, newVars, nil
}

func (c *Collection) parseBuildArg(arg string, pncvf ProcessNonConstantVariableFunc) (string, Variable, bool, error) {
	var name string
	splitArg := strings.SplitN(arg, "=", 2)
	if len(splitArg) < 1 {
		return "", Variable{}, false, fmt.Errorf("invalid build arg %s", splitArg)
	}
	name = splitArg[0]
	value := ""
	hasValue := false
	if len(splitArg) == 2 {
		value = splitArg[1]
		hasValue = true
	}
	v, err := c.parseBuildArgValue(name, value, pncvf)
	if err != nil {
		return "", Variable{}, false, err
	}
	return name, v, hasValue, nil
}

func (c *Collection) parseBuildArgValue(name, value string, pncvf ProcessNonConstantVariableFunc) (Variable, error) {
	if !strings.HasPrefix(value, "$(") {
		// Constant build arg.
		return NewConstant(value), nil
	}

	// Variable build arg - resolve value and turn it into a constant build arg.
	value, _, err := pncvf(name, value)
	if err != nil {
		return Variable{}, err
	}
	return NewConstant(value), nil
}
