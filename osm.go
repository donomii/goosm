package main

import (
    "compress/gzip"
    "encoding/csv"
    "encoding/xml"
    "os"
    "fmt"
    "strings"
    "regexp"
    "strconv"
    "code.google.com/p/gos2/s2"
)

type Tag struct {
    Key string `xml:"k,attr"`
    Value string `xml:"v,attr"`
}

//'id', 'created_at', 'closed_at', 'user', 'uid', 'tags', 'bbox'
type Changeset struct {
    Id string `xml:"id,attr"`
    CreatedAt string `xml:"created_at,attr"`
    ClosedAt string `xml:"closed_at,attr"`
    User string `xml:"user,attr"`
    Uid string `xml:"uid,attr"`
    Tags []Tag `xml:"tag"`
    MinLat string `xml:"min_lat,attr"`
    MinLon string `xml:"min_lon,attr"`
    MaxLat string `xml:"max_lat,attr"`
    MaxLon string `xml:"max_lon,attr"`
}

//'id', 'version', 'changeset', 'user', 'uid', 'visible', 'timestamp', 'tags', 'loc'
type Node struct {
    Id string `xml:"id,attr"`
    Version string `xml:"version,attr"`
    Changeset string `xml:"changeset,attr"`
    User string `xml:"user,attr"`
    Uid string `xml:"uid,attr"`
    Visible string `xml:"visible,attr"`
    Timestamp string `xml:"timestamp,attr"`
    Tags []Tag `xml:"tag"`
    Lat string `xml:"lat,attr"`
    Lon string `xml:"lon,attr"`
}

//'id', 'version', 'changeset', 'user', 'uid', 'visible', 'timestamp', 'tags', 'nds'
type Way struct {
    Id string `xml:"id,attr"`
    Version string `xml:"version,attr"`
    Changeset string `xml:"changeset,attr"`
    User string `xml:"user,attr"`
    Uid string `xml:"uid,attr"`
    Visible string `xml:"visible,attr"`
    Timestamp string `xml:"timestamp,attr"`
    Tags []Tag `xml:"tag"`
    Nds []Nd `xml:"nd"`
}

type Nd struct {
    Ref string `xml:"ref,attr"`
}

//'id', 'version', 'changeset', 'user', 'uid', 'visible', 'timestamp', 'tags', 'members'
type Relation struct {
    Id string `xml:"id,attr"`
    Version string `xml:"version,attr"`
    Changeset string `xml:"changeset,attr"`
    User string `xml:"user,attr"`
    Uid string `xml:"uid,attr"`
    Visible string `xml:"visible,attr"`
    Timestamp string `xml:"timestamp,attr"`
    Tags []Tag `xml:"tag"`
    Members []Member `xml:"member"`
}

type Member struct {
    Type string `xml:"type,attr"`
    Ref string `xml:"ref,attr"`
    Role string `xml:"role,attr"`
}

var XML = `<?xml version='1.0' encoding='UTF-8'?>
<osm version="0.6" generator="OpenStreetMap HistoryDump.java" timestamp="2013-02-05T17:01:47Z" copyright="OpenStreetMap and contributors" attribution="http://www.openstreetmap.org/copyright" license="http://opendatacommons.org/licenses/odbl/1-0/">
    <bound box="-90,-180,90,180" origin="http://api.openstreetmap.org/api/0.6"/>
    <changeset id="1" created_at="2005-04-09T19:54:13Z" closed_at="2005-04-09T20:54:39Z" open="false" min_lat="51.5288506" max_lat="51.5288620" min_lon="-0.1465242" max_lon="-0.1464925" user="Steve" uid="1">
        <tag k="doowa==" v=">>>>"/>
    </changeset>
    <changeset id="28" created_at="2005-05-22T15:01:46Z" closed_at="2005-05-22T18:42:22Z" open="true" min_lat="59.9038811" max_lat="60.1498756" min_lon="10.2228117" max_lon="10.7978525" user="Petter Reinholdtsen" uid="24">
    </changeset>
    <node id="24199285" version="1" changeset="185018" lat="5.70711" lon="-4.82095" user="Coastlines-RJB" uid="5330" visible="true" timestamp="2007-01-02T12:40:39Z">
        <tag k="created_by" v="almien_coastlines" />
        <tag k="source" v="PGS" />
    </node>
    <way id="213941833" visible="true" timestamp="2013-05-29T00:31:06Z" version="2" changeset="16331880" user="iandees" uid="4732">
        <nd ref="2235863476" />
        <nd ref="2235863470" />
        <nd ref="2235863497" />
        <nd ref="2235863506" />
        <nd ref="2235863476" />
        <tag k="addr:city" v="Evanston" />
        <tag k="addr:housenumber" v="634" />
        <tag k="addr:postcode" v="60202" />
        <tag k="addr:street" v="Judson Avenue" />
        <tag k="addr:street:name" v="Judson" />
        <tag k="addr:street:type" v="Avenue" />
        <tag k="building" v="house" />
    </way>
    <relation id="122552" visible="true" timestamp="2013-08-04T22:02:25Z" version="113" changeset="17220940" user="Steven Vance" uid="100042">
        <member type="node" ref="2015352819" role="stop" />
        <member type="node" ref="2015341023" role="stop" />
        <member type="node" ref="2015331925" role="stop" />
        <member type="node" ref="2015319463" role="stop" />
        <member type="node" ref="2014360953" role="stop" />
        <member type="node" ref="736686706" role="stop" />
        <member type="node" ref="736686728" role="stop" />
        <member type="way" ref="59526464" role="" />
        <member type="way" ref="59526440" role="" />
        <tag k="colour" v="#522398" />
        <tag k="from" v="The Loop" />
        <tag k="name" v="Purple Line to Linden" />
        <tag k="network" v="CTA" />
        <tag k="operator" v="Chicago Transit Authority" />
        <tag k="ref" v="Purple" />
        <tag k="route" v="subway" />
        <tag k="to" v="Linden" />
        <tag k="type" v="route" />
        <tag k="wheelchair" v="limited" />
    </relation>
</osm>`

func tagsCSV(tags []Tag) string {
    var tagParts []string
    var escape = regexp.MustCompile("([\\=>])")

    for _,tag := range tags {
        tagParts = append(tagParts, fmt.Sprintf("\"%s\"=>\"%s\"", escape.ReplaceAllString(tag.Key, "\\$1"), escape.ReplaceAllString(tag.Value, "\\$1")))
    }

    return strings.Join(tagParts, ", ")
}

func membersCSV(members []Member) string {
    var memberParts []string

    for _,member := range members {
        memberParts = append(memberParts, fmt.Sprintf("{\"%s\", %s, \"%s\"}", member.Type, member.Ref, member.Role))
    }

    return fmt.Sprintf("{%s}", strings.Join(memberParts, ", "))
}

func ndsCSV(nds []Nd) string {
    var ndParts []string

    for _,node := range nds {
        ndParts = append(ndParts, node.Ref)
    }

    return fmt.Sprintf("{%s}", strings.Join(ndParts, ", "))
}

func latLon(lat string, lon string) string {
    if lat == "" {
        return "null"
    }

    return fmt.Sprintf("%s, %s", lat, lon)
}

func bboxCSV(changeset Changeset) string {
    if changeset.MinLon == "" {
        return "null"
    }

    return fmt.Sprintf("%s, %s, %s, %s", changeset.MinLon, changeset.MaxLat, changeset.MaxLon, changeset.MinLat)
}

func emptyCheck(s string) string {
    if s == "" {
        return "null"
    }

    return s
}

func main() {
    in, _ := os.Open("/dev/stdin")
    decoder := xml.NewDecoder(in)

    changesetsFile, _ := os.Create("changesets.csv.gz")
    changesetsOut := csv.NewWriter(gzip.NewWriter(changesetsFile))
    nodesFile, _ := os.Create("nodes.csv.gz")
    nodesOut := csv.NewWriter(gzip.NewWriter(nodesFile))
    waysFile, _ := os.Create("ways.csv.gz")
    waysOut := csv.NewWriter(gzip.NewWriter(waysFile))
    relationsFile, _ := os.Create("relations.csv.gz")
    relationsOut := csv.NewWriter(gzip.NewWriter(relationsFile))

    var nodes, ways, relations, changesets, prims int

    for token, err := decoder.Token(); err == nil; token, err = decoder.Token() {
        switch tag := token.(type) {
        case xml.StartElement:
            switch tag.Name.Local {
            case "changeset":
                changeset := Changeset{}
                decoder.DecodeElement(&changeset, &tag)
                // 'id', 'created_at', 'closed_at', 'user', 'uid', 'tags', 'bbox'
                changesetsOut.Write([]string{
                    changeset.Id,
                    changeset.CreatedAt,
                    emptyCheck(changeset.ClosedAt),
                    emptyCheck(changeset.User),
                    emptyCheck(changeset.Uid),
                    tagsCSV(changeset.Tags),
                    bboxCSV(changeset),
                })
                changesets++
            case "node":
                node := Node{}
                decoder.DecodeElement(&node, &tag)
                // 'id', 'version', 'changeset', 'user', 'uid', 'visible', 'timestamp', 'tags', 'loc'
                lat, _ := strconv.ParseFloat(node.Lat, 64)
                lon, _ := strconv.ParseFloat(node.Lon, 64)
                fmt.Printf("S2 for %f %f is %s\n", lat, lon, s2.CellIDFromLatLng(s2.LatLngFromDegrees(lat, lon)).String())
                nodesOut.Write([]string{
                    node.Id,
                    node.Version,
                    node.Changeset,
                    node.User,
                    node.Uid,
                    node.Visible,
                    node.Timestamp,
                    tagsCSV(node.Tags),
                    emptyCheck(latLon(node.Lat, node.Lon)),
                })
                nodes++
            case "way":
                way := Way{}
                decoder.DecodeElement(&way, &tag)
                // 'id', 'version', 'changeset', 'user', 'uid', 'visible', 'timestamp', 'tags', 'nds'
                waysOut.Write([]string{
                    way.Id,
                    way.Version,
                    way.Changeset,
                    way.User,
                    way.Uid,
                    way.Visible,
                    way.Timestamp,
                    tagsCSV(way.Tags),
                    ndsCSV(way.Nds),
                })
                ways++
            case "relation":
                relation := Relation{}
                decoder.DecodeElement(&relation, &tag)
                // 'id', 'version', 'changeset', 'user', 'uid', 'visible', 'timestamp', 'tags', 'members'
                relationsOut.Write([]string{
                    relation.Id,
                    relation.Version,
                    relation.Changeset,
                    relation.User,
                    relation.Uid,
                    relation.Visible,
                    relation.Timestamp,
                    tagsCSV(relation.Tags),
                    membersCSV(relation.Members),
                })
                relations++
            }
        }
        prims++

        if prims % 1000 == 0 {
            os.Stdout.Write([]byte("\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b"))
            os.Stdout.Write([]byte(fmt.Sprintf("%8d changesets, %10d nodes, %8d ways, %7d relations", changesets, nodes, ways, relations)))
            os.Stdout.Sync()
        }
    }

    os.Stdout.Write([]byte("\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b"))
    os.Stdout.Write([]byte(fmt.Sprintf("%8d changesets, %10d nodes, %8d ways, %7d relations", changesets, nodes, ways, relations)))
    os.Stdout.Sync()

    changesetsOut.Flush()
    changesetsFile.Close()
    nodesOut.Flush()
    nodesFile.Close()
    waysOut.Flush()
    waysFile.Close()
    relationsOut.Flush()
    relationsFile.Close()
}