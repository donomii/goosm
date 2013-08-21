goosm
=====

Parses OSM XML and writes it to CSV files intended to be loaded into PostgreSQL.

Usage
-----

    go build osm.go
    curl -s http://planet.openstreetmap.org/planet/full-history/history-latest.osm.bz2 | bzcat | ./osm
