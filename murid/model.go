package murid

type Murid struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	TimeCreated string `db:"time_create"`
}

type Pendidikan struct {
	Id         int    `db:"id"`
	IdMurid    int    `db:"id_murid"`
	Status     string `db:"status"`
	TimeCrated string `db:"time_crate"`
}
