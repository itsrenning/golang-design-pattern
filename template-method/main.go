package main

func main() {
	mdb := MongoDB{}
	DailyOperation(&mdb)

	pdb := PostgresDB{}
	DailyOperation(&pdb)
}
