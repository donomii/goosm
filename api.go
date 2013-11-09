package goosm

import (
    "github.com/gorilla/mux"
    "net/http"
)

func CapabilitiesHandler(w http.ResponseWriter, r *http.Request) {

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