package murid

type Murid struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	TimeCreated string `db:"time_create"`
}

type Pendidikan struct {
	Id         int    `db:"id"`
	IdMurid    int    `db:"id_murid"`
	Status     string `db:"status"`
	TimeCrated string `db:"time_crate"`
}
