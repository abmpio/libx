package template

import (
	"bytes"
	"fmt"
	"text/template"

	net "github.com/abmpio/libx/net"
)

var (
	// SourceFuncs is a map of all top-level functions that generate
	// sockaddr data types.
	SourceFuncs template.FuncMap

	// SortFuncs is a map of all functions used in sorting
	SortFuncs template.FuncMap

	// FilterFuncs is a map of all functions used in sorting
	FilterFuncs template.FuncMap

	// HelperFuncs is a map of all functions used in sorting
	HelperFuncs template.FuncMap
)

func init() {
	SourceFuncs = template.FuncMap{
		// GetAllInterfaces - Returns an exhaustive set of IfAddr
		// structs available on the host.  `GetAllInterfaces` is the
		// initial input and accessible as the initial "dot" in the
		// pipeline.
		"GetAllInterfaces": net.GetAllInterfaces,

		// GetDefaultInterfaces - Returns one IfAddr for every IP that
		// is on the interface containing the default route for the
		// host.
		"GetDefaultInterfaces": net.GetDefaultInterfaces,

		// GetPrivateInterfaces - Returns one IfAddr for every IP that
		// matches RFC 6890, are attached to the interface with the
		// default route, and are forwardable IP addresses.  NOTE: RFC
		// 6890 is a more exhaustive version of RFC1918 because it spans
		// IPv4 and IPv6, however it doespermit the inclusion of likely
		// undesired addresses such as multicast, therefore our
		// definition of a "private" address also excludes
		// non-forwardable IP addresses (as defined by the IETF).
		"GetPrivateInterfaces": net.GetPrivateInterfaces,

		// GetPublicInterfaces - Returns a list of IfAddr that do not
		// match RFC 6890, are attached to the default route, and are
		// forwardable.
		"GetPublicInterfaces": net.GetPublicInterfaces,
	}

	SortFuncs = template.FuncMap{
		"sort": net.SortIfBy,
	}

	FilterFuncs = template.FuncMap{
		"exclude": net.ExcludeIfs,
		"include": net.IncludeIfs,
	}

	HelperFuncs = template.FuncMap{
		// Misc functions that operate on IfAddrs inputs
		"attr":   Attr,
		"join":   net.JoinIfAddrs,
		"limit":  net.LimitIfAddrs,
		"offset": net.OffsetIfAddrs,
		"unique": net.UniqueIfAddrsBy,

		// Misc math functions that operate on a single IfAddr input
		"math": net.IfAddrsMath,

		// Return a Private RFC 6890 IP address string that is attached
		// to the default route and a forwardable address.
		"GetPrivateIP": net.GetPrivateIP,

		// Return all Private RFC 6890 IP addresses as a space-delimited string of
		// IP addresses.  Addresses returned do not have to be on the interface with
		// a default route.
		"GetPrivateIPs": net.GetPrivateIPs,

		// Return a Public RFC 6890 IP address string that is attached
		// to the default route and a forwardable address.
		"GetPublicIP": net.GetPublicIP,

		// Return allPublic RFC 6890 IP addresses as a space-delimited string of IP
		// addresses.  Addresses returned do not have to be on the interface with a
		// default route.
		"GetPublicIPs": net.GetPublicIPs,

		// Return the first IP address of the named interface, sorted by
		// the largest network size.
		"GetInterfaceIP": net.GetInterfaceIP,

		// Return all IP addresses on the named interface, sorted by the largest
		// network size.
		"GetInterfaceIPs": net.GetInterfaceIPs,
	}
}

// Attr returns the attribute from the ifAddrRaw argument.  If the argument is
// an IfAddrs, only the first element will be evaluated for resolution.
func Attr(selectorName string, ifAddrsRaw interface{}) (string, error) {
	switch v := ifAddrsRaw.(type) {
	case net.IfAddr:
		return net.IfAttr(selectorName, v)
	case net.IfAddrs:
		return net.IfAttrs(selectorName, v)
	default:
		return "", fmt.Errorf("unable to obtain attribute %s from type %T (%v)", selectorName, ifAddrsRaw, ifAddrsRaw)
	}
}

// Parse parses input as template input using the addresses available on the
// host, then returns the string output if there are no errors.
func Parse(input string) (string, error) {
	addrs, err := net.GetAllInterfaces()
	if err != nil {
		return "", fmt.Errorf("unable to query interface addresses: {{err}}", err)
	}

	return ParseIfAddrs(input, addrs)
}

// ParseIfAddrs parses input as template input using the IfAddrs inputs, then
// returns the string output if there are no errors.
func ParseIfAddrs(input string, ifAddrs net.IfAddrs) (string, error) {
	return ParseIfAddrsTemplate(input, ifAddrs, template.New("sockaddr.Parse"))
}

// ParseIfAddrsTemplate parses input as template input using the IfAddrs inputs,
// then returns the string output if there are no errors.
func ParseIfAddrsTemplate(input string, ifAddrs net.IfAddrs, tmplIn *template.Template) (string, error) {
	// Create a template, add the function map, and parse the text.
	tmpl, err := tmplIn.Option("missingkey=error").
		Funcs(SourceFuncs).
		Funcs(SortFuncs).
		Funcs(FilterFuncs).
		Funcs(HelperFuncs).
		Parse(input)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("unable to parse template %+q: {{err}}", input), err)
	}

	var outWriter bytes.Buffer
	err = tmpl.Execute(&outWriter, ifAddrs)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("unable to execute sockaddr input %+q: {{err}}", input), err)
	}

	return outWriter.String(), nil
}
