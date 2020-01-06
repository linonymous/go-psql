package geolocation

type Geolocation struct {
	ID        int     `gorm:"id;primary_key;auto_increment"`
	MCC       int     `gorm:"mcc"`
	MNC       int     `gorm:"mnc"`
	LAC       int     `gorm:"lac"`
	CellID    int     `gorm:"cell_id"`
	Latitude  float32 `gorm:"latitude"`
	Longitude float32 `gorm:"longitude"`
	Range     int     `gorm:"range"`
	Radio     string  `gorm:"radio"`
}
