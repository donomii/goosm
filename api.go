package main

import (
    "encoding/xml"
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
)

type Tag struct {
    XMLName     xml.Name    `xml:"tag" json:"-"`
    Key         string      `xml:"k,attr"`
    Value       string      `xml:"v,attr"`
}

type Changeset struct {
    XMLName     xml.Name    `xml:"changeset" json:"-"`
    Id          int32       `xml:"id,attr"`
    CreatedAt   string      `xml:"created_at,attr"`
    ClosedAt    string      `xml:"closed_at,attr"`
    User        string      `xml:"user,attr"`
    Uid         int32       `xml:"uid,attr"`
    Tags        []Tag       `xml:"tag"`
    MinLat      float32     `xml:"min_lat,attr"`
    MinLon      float32     `xml:"min_lon,attr"`
    MaxLat      float32     `xml:"max_lat,attr"`
    MaxLon      float32     `xml:"max_lon,attr"`
}

type Node struct {
    XMLName     xml.Name    `xml:"node" json:"-"`
    Id          int64       `xml:"id,attr" json:"id"`
    Version     int16       `xml:"version,attr" json:"version,attr"`
    Changeset   int32       `xml:"changeset,attr" json:"changeset,attr"`
    User        string      `xml:"user,attr" json:"user,attr"`
    Uid         int32       `xml:"uid,attr" json:"uid,attr"`
    Visible     bool        `xml:"visible,attr" json:"visible,attr"`
    Timestamp   string      `xml:"timestamp,attr" json:"timestamp,attr"`
    Tags        []Tag       `xml:"tag" json:"tag"`
    Lat         float32     `xml:"lat,attr" json:"lat,attr"`
    Lon         float32     `xml:"lon,attr" json:"lon,attr"`
}

type Way struct {
    XMLName     xml.Name    `xml:"way" json:"-"`
    Id          int64       `xml:"id,attr"`
    Version     int16       `xml:"version,attr"`
    Changeset   int32       `xml:"changeset,attr"`
    User        string      `xml:"user,attr"`
    Uid         int32       `xml:"uid,attr"`
    Visible     bool        `xml:"visible,attr"`
    Timestamp   string      `xml:"timestamp,attr"`
    Tags        []Tag       `xml:"tag"`
    Nds         []Nd        `xml:"nd"`
}

type Nd struct {
    XMLName     xml.Name    `xml:"nd" json:"-"`
    Ref         string      `xml:"ref,attr"`
}

type Relation struct {
    XMLName     xml.Name    `xml:"relation" json:"-"`
    Id          int64       `xml:"id,attr"`
    Version     int16       `xml:"version,attr"`
    Changeset   int32       `xml:"changeset,attr"`
    User        string      `xml:"user,attr"`
    Uid         int32       `xml:"uid,attr"`
    Visible     bool        `xml:"visible,attr"`
    Timestamp   string      `xml:"timestamp,attr"`
    Tags        []Tag       `xml:"tag"`
    Members     []Member    `xml:"member"`
}

type Member struct {
    XMLName     xml.Name    `xml:"member" json:"-"`
    Type        string      `xml:"type,attr"`
    Ref         string      `xml:"ref,attr"`
    Role        string      `xml:"role,attr"`
}

type Capabilities struct {
    MinimumVersion string `xml:"api>version>minimum,attr"`
    MaxmimumVersion string `xml:"api>version>maximum,attr"`
}

func CapabilitiesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/xml")
    w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<osm version="0.6" generator="goosm">
    <api>
        <version minimum="0.6" maximum="0.6"/>
        <area maximum="0.25"/>
        <waynodes maximum="2000"/>
        <changesets maximum_elements="50000"/>
        <timeout seconds="300"/>
    </api>
</osm>`))
}

func MapHandler(w http.ResponseWriter, r *http.Request) {

}

func PermissionsHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateChangesetHandler(w http.ResponseWriter, r *http.Request) {

}

func ShowChangesetHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateChangesetHandler(w http.ResponseWriter, r *http.Request) {

}

func CloseChangesetHandler(w http.ResponseWriter, r *http.Request) {

}

func DownloadChangesetHandler(w http.ResponseWriter, r *http.Request) {

}

func ExpandBboxChangesetHandler(w http.ResponseWriter, r *http.Request) {

}

func UploadChangesetHandler(w http.ResponseWriter, r *http.Request) {

}

func QueryChangesetsHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateNodeHandler(w http.ResponseWriter, r *http.Request) {

}

func GetNodeHandler(w http.ResponseWriter, r *http.Request) {
    node := Node{Id:1214, Version:1, Changeset:124, User:"iandees", Uid:4732, Visible:true, Timestamp:"foo", Tags:[]Tag{}, Lat:45.1, Lon:-90.2}
    
    if (true) {
        w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>`))
        w.Write([]byte(`<osm version="0.6" generator="goosm">`))

        enc := xml.NewEncoder(w)
        enc.Encode(node)

        w.Write([]byte("</osm>\n"))
    } else {
        enc := json.NewEncoder(w)
        enc.Encode(node)
    }
}

func UpdateNodeHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteNodeHandler(w http.ResponseWriter, r *http.Request) {

}

func NodeFullHistoryHandler(w http.ResponseWriter, r *http.Request) {

}

func NodeSingleHistoryHandler(w http.ResponseWriter, r *http.Request) {

}

func NodeRelationMembershipHandler(w http.ResponseWriter, r *http.Request) {

}

func NodeWayMembershipHandler(w http.ResponseWriter, r *http.Request) {

}

func NodeMultifetchHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateWayHandler(w http.ResponseWriter, r *http.Request) {

}

func GetWayHandler(w http.ResponseWriter, r *http.Request) {
    way := Way{Id:1214, Version:1, Changeset:124, User:"iandees", Uid:4732, Visible:true, Timestamp:"foo"}
    
    w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<osm version="0.6" generator="goosm">`))

    enc := json.NewEncoder(w)
    enc.Encode(way)

    w.Write([]byte("</osm>\n"))
}

func UpdateWayHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteWayHandler(w http.ResponseWriter, r *http.Request) {

}

func WayFullHistoryHandler(w http.ResponseWriter, r *http.Request) {

}

func WaySingleHistoryHandler(w http.ResponseWriter, r *http.Request) {

}

func WayRelationMembershipHandler(w http.ResponseWriter, r *http.Request) {

}

func WayFullHandler(w http.ResponseWriter, r *http.Request) {

}

func WayMultifetchHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateRelationHandler(w http.ResponseWriter, r *http.Request) {

}

func GetRelationHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateRelationHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteRelationHandler(w http.ResponseWriter, r *http.Request) {

}

func RelationFullHistoryHandler(w http.ResponseWriter, r *http.Request) {

}

func RelationSingleHistoryHandler(w http.ResponseWriter, r *http.Request) {

}

func RelationRelationMembershipHandler(w http.ResponseWriter, r *http.Request) {

}

func RelationFullHandler(w http.ResponseWriter, r *http.Request) {

}

func RelationMultifetchHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/capabilities", CapabilitiesHandler)

    s := r.PathPrefix("/api/0.6").Subrouter()
    s.HandleFunc("/map", MapHandler).Methods("GET")
    s.HandleFunc("/permissions", PermissionsHandler).Methods("GET")

    ch := s.PathPrefix("/changeset").Subrouter()
    ch.HandleFunc("/create", CreateChangesetHandler).Methods("PUT")
    ch.HandleFunc("/{changeset_id}", ShowChangesetHandler).Methods("GET")
    ch.HandleFunc("/{changeset_id}", UpdateChangesetHandler).Methods("PUT")
    ch.HandleFunc("/{changeset_id}/close", CloseChangesetHandler).Methods("PUT")
    ch.HandleFunc("/{changeset_id}/download", DownloadChangesetHandler).Methods("GET")
    ch.HandleFunc("/{changeset_id}/expand_bbox", ExpandBboxChangesetHandler).Methods("POST")
    ch.HandleFunc("/{changeset_id}/upload", UploadChangesetHandler).Methods("POST")
    s.HandleFunc("/changesets", QueryChangesetsHandler).Methods("GET")

    nh := s.PathPrefix("/node").Subrouter()
    nh.HandleFunc("/create", CreateNodeHandler).Methods("PUT")
    nh.HandleFunc("/{node_id}", GetNodeHandler).Methods("GET")
    nh.HandleFunc("/{node_id}", UpdateNodeHandler).Methods("PUT")
    nh.HandleFunc("/{node_id}", DeleteNodeHandler).Methods("DELETE")
    nh.HandleFunc("/{node_id}/history", NodeFullHistoryHandler).Methods("GET")
    nh.HandleFunc("/{node_id}/{version}", NodeSingleHistoryHandler).Methods("GET")
    nh.HandleFunc("/{node_id}/relations", NodeRelationMembershipHandler).Methods("GET")
    nh.HandleFunc("/{node_id}/ways", NodeWayMembershipHandler).Methods("GET")
    s.HandleFunc("/nodes", NodeMultifetchHandler).Methods("GET").Queries("nodes")

    wh := s.PathPrefix("/way").Subrouter()
    wh.HandleFunc("/create", CreateWayHandler).Methods("PUT")
    wh.HandleFunc("/{way_id}", GetWayHandler).Methods("GET")
    wh.HandleFunc("/{way_id}", UpdateWayHandler).Methods("PUT")
    wh.HandleFunc("/{way_id}", DeleteWayHandler).Methods("DELETE")
    wh.HandleFunc("/{way_id}/history", WayFullHistoryHandler).Methods("GET")
    wh.HandleFunc("/{way_id}/{version}", WaySingleHistoryHandler).Methods("GET")
    wh.HandleFunc("/{way_id}/relations", WayRelationMembershipHandler).Methods("GET")
    wh.HandleFunc("/{way_id}/full", WayFullHandler).Methods("GET")
    s.HandleFunc("/ways", WayMultifetchHandler).Methods("GET").Queries("ways")

    rh := s.PathPrefix("/relation").Subrouter()
    rh.HandleFunc("/create", CreateRelationHandler).Methods("PUT")
    rh.HandleFunc("/{relation_id}", GetRelationHandler).Methods("GET")
    rh.HandleFunc("/{relation_id}", UpdateRelationHandler).Methods("PUT")
    rh.HandleFunc("/{relation_id}", DeleteRelationHandler).Methods("DELETE")
    rh.HandleFunc("/{relation_id}/history", RelationFullHistoryHandler).Methods("GET")
    rh.HandleFunc("/{relation_id}/{version}", RelationSingleHistoryHandler).Methods("GET")
    rh.HandleFunc("/{relation_id}/relations", RelationRelationMembershipHandler).Methods("GET")
    rh.HandleFunc("/{relation_id}/full", RelationFullHandler).Methods("GET")
    s.HandleFunc("/relations", RelationMultifetchHandler).Methods("GET").Queries("relations")

    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)
}