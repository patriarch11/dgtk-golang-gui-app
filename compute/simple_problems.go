package compute

import "github.com/pymaxion/geographiclib-go/geodesic"

func ResolveDirectGeodesicProblem(lat1, lon1, azimuth, distance float64) (lat2, lon2 float64) {
	res := geodesic.WGS84.Direct(lat1, lon1, azimuth, distance)
	lat2, lon2 = res.Lat2, res.Lon2
	return
}

func ResolveInverseGeodesicProblem(lat1, lon1, lat2, lon2 float64) (azi1, azi2, distance float64) {
	res := geodesic.WGS84.Inverse(lat1, lon1, lat2, lon2)
	azi1, azi2, distance = res.Azi1, res.Azi2, res.S12
	return
}
