package cerchir

import "fmt"

type Position struct {
	X float64
	Y float64
	Z float64
}

func Chiron(dateInSeconds float64, chiron_np []float64) Position {

	type FileRecords struct {
		rec_start_addr int
		seg_start_time float64
		seg_last_time  float64
		int_len        float64
		rec_last_addr  int
	}

	var CHIRON_FILE_RECORDS = [12]FileRecords{
		{
			// SUN
			rec_start_addr: 0,
			seg_start_time: 0,
			seg_last_time:  0,
			int_len:        0,
			rec_last_addr:  0,
			// "rsize": 0,
		},
		{
			// 1
			rec_start_addr: 8_065,
			seg_start_time: -120_450_514.89409208,
			seg_last_time:  510_321_600,
			int_len:        510_321_600,
			rec_last_addr:  54_071,
			// "rsize": 20,
			// init =  391_827_779.23949105
			// n =  500.0
		},
		{
			// 2
			rec_start_addr: 54_072,
			seg_start_time: -780_357_478,
			seg_last_time:  -120_450_514,
			int_len:        -120_450_514,
			rec_last_addr:  100_078,
			// "rsize": 20,
			// init =  -256344669.7390517
			// n =  500.0
		},
		{
			// 3
			rec_start_addr: 100_079,
			seg_start_time: -1_428_684_898,
			seg_last_time:  -780_357_478,
			int_len:        -780_357_478,
			rec_last_addr:  146_085,
			// "rsize": 20,
			// init =  -909271241.4267497
			// n =  500.0
		},
		{
			// 4
			rec_start_addr: 146_086,
			seg_start_time: -2_072_145_540,
			seg_last_time:  -1_428_684_898,
			int_len:        -1_428_684_898,
			rec_last_addr:  192_092,
			// "rsize": 20,
			// init =  -1559481315.7158456
			// n =  500.0
		},
		{
			// 5
			rec_start_addr: 192_093,
			seg_start_time: -2_719_741_734,
			seg_last_time:  -2_072_145_540,
			int_len:        -2_072_145_540,
			rec_last_addr:  238_099,
			// "rsize": 20,
			// init =  -2202930574.1405
			// n =  500.0
		},
		{
			// 6
			rec_start_addr: 238_100,
			seg_start_time: -3_155_716_800,
			seg_last_time:  -2_719_741_734,
			int_len:        -2_758_703_306,
			// intlen =  -2758703306.856402 ??????????????
			rec_last_addr: 268_464,
			// "rsize": 20,
			// init =  -2891244118.0133157
			// n =  330.0
		},
		{
			// 7
			rec_start_addr: 268_465,
			seg_start_time: 510_321_600,
			seg_last_time:  1_134_172_990,
			int_len:        1_134_172_990,
			rec_last_addr:  314_471,
			// "rsize": 20,
			// init =  1002901396.5747848
			// n =  500.0
		},
		{
			// 8
			rec_start_addr: 314_472,
			seg_start_time: 1_134_172_990,
			seg_last_time:  1_785_159_000,
			int_len:        1_785_159_000,
			rec_last_addr:  360_478,
			// "rsize": 20,
			// init =  1655649864.227839
			// n =  500.0
		},
		{
			// 9
			rec_start_addr: 360_479,
			seg_start_time: 1_785_159_000,
			seg_last_time:  2_451_382_230,
			int_len:        2_451_382_230,
			rec_last_addr:  406_485,
			// "rsize": 20,
			// init =  2313659513.647009
			// n =  500.0
		},
		{
			// 10
			rec_start_addr: 406_486,
			seg_start_time: 2_451_382_230,
			seg_last_time:  3_098_798_447,
			int_len:        3_098_798_447,
			rec_last_addr:  452_492,
			// "rsize": 20,
			// init =  2967630632.730112
			// n =  500.0
		},
		{
			// 11
			rec_start_addr: 452_493,
			seg_start_time: 3_098_798_447,
			seg_last_time:  3_187_252_800,
			int_len:        3_187_252_800,
			rec_last_addr:  458_842,
			// "rsize": 20,
			// init =  3185947430.156654
			// n =  69.0
		},
	}

	const total_summaries_number = 11
	// max_dim is always 20. ceres_np[CERES_FILE_RECORDS[i_summ]["rec_last_addr"] - 2]
	const max_dim = 20
	// DFLSIZ = (4 * max_dim) + 11  // rsize = (4 * max_dim) + 11
	const DFLSIZ = 91
	const BUFSIZ = 100

	for i_summ := 1; i_summ <= total_summaries_number; i_summ++ {
		if CHIRON_FILE_RECORDS[i_summ].seg_start_time < dateInSeconds &&
			CHIRON_FILE_RECORDS[i_summ].seg_last_time > dateInSeconds {

			start_adress := CHIRON_FILE_RECORDS[i_summ].rec_start_addr
			last_adress := CHIRON_FILE_RECORDS[i_summ].rec_last_addr

			n_of_rec := int(chiron_np[last_adress-1])

			// Number of directory epochs
			var n_of_dir int = n_of_rec / BUFSIZ

			OFFD := last_adress - n_of_dir - 2
			OFFE := OFFD - n_of_rec

			RECNO := -1

			if n_of_rec <= BUFSIZ {

				data := chiron_np[(OFFE):(OFFE + n_of_rec)]

				if dateInSeconds < data[0] || dateInSeconds > data[len(data)-1] {
					fmt.Println("we have a problem")
				}

				for i, v := range data {
					if v == dateInSeconds {
						// fmt.Println("equality , index =", i, "value = ", arr[i])
						RECNO = i
						break
					}
					if v > dateInSeconds {
						// fmt.Println("index =", i-1, "value = ", arr[i-1])
						RECNO = i - 1
						break
					}
				}

			} else {
				for dir := 0; dir < n_of_dir; dir++ {

					data := chiron_np[(OFFD + dir) : (OFFD+dir)+1][0]

					if data > dateInSeconds {
						OFFD = OFFE + (dir)*BUFSIZ
						data := chiron_np[(OFFD):(OFFD + BUFSIZ)]

						// looking for the largest array element less than  dateInSeconds
						// and get its index
						if dateInSeconds < data[0] || dateInSeconds > data[len(data)-1] {
							fmt.Println("we have a problem")
						}

						for i, v := range data {
							if v == dateInSeconds {
								// fmt.Println("equality , index =", i, "value = ", arr[i])
								RECNO = i
								break
							}
							if v > dateInSeconds {
								// fmt.Println("index =", i-1, "value = ", arr[i-1])
								RECNO = i - 1
								break
							}
						}

						RECNO = dir*BUFSIZ + RECNO
						break
					}
				}
			}

			if RECNO == -1 {
				// print("chiron final records sec = ", dateInSeconds)
				Ind := n_of_rec % BUFSIZ
				data := chiron_np[(last_adress - n_of_dir - Ind) : last_adress-n_of_dir]

				// looking for the largest array element less than  dateInSeconds
				// and get its index
				if dateInSeconds < data[0] || dateInSeconds > data[len(data)-1] {
					fmt.Println("we have a problem")
				}

				for i, v := range data {
					if v == dateInSeconds {
						// fmt.Println("equality , index =", i, "value = ", arr[i])
						RECNO = i
						break
					}
					if v > dateInSeconds {
						// fmt.Println("index =", i-1, "value = ", arr[i-1])
						RECNO = i - 1
						break
					}
				}

				RECNO = (n_of_dir)*BUFSIZ + RECNO
			}

			OFFR := start_adress - 1 + (RECNO)*DFLSIZ
			mda_record := chiron_np[(OFFR):(OFFR + DFLSIZ)]

			// print("dateInSeconds = ", dateInSeconds)
			// print(mda_record)

			TL := mda_record[0]
			G := mda_record[1 : max_dim+1]
			REFPOS := []float64{mda_record[max_dim+1], mda_record[max_dim+3], mda_record[max_dim+5]}
			REFVEL := []float64{mda_record[max_dim+2], mda_record[max_dim+4], mda_record[max_dim+6]}

			KQMAX1 := int(mda_record[4*max_dim+7])

			KQ := []float64{(mda_record[4*max_dim+8]), (mda_record[4*max_dim+9]), (mda_record[4*max_dim+10])}

			KS := KQMAX1 - 1
			MQ2 := KQMAX1 - 2
			DELTA := dateInSeconds - TL
			TP := DELTA

			// FC = [0 for i in range(max_dim)]
			var FC [max_dim]float64
			FC[0] = 1.0
			// WC = [0 for i in range(max_dim - 1)]
			var WC [max_dim - 1]float64

			for J := 1; J < MQ2+1; J++ {
				if G[J-1] == 0.0 {
					fmt.Println("SPKE21\nA value of zero was found at index {0} of the step size vector.")
				}
				FC[J] = TP / G[J-1]
				WC[J-1] = DELTA / G[J-1]
				TP = DELTA + G[J-1]
			}
			// W = [0 for i in range(max_dim + 2)]
			var W [max_dim + 2]float64

			//
			//     Collect KQMAX1 reciprocals.
			//     KS = KQMAX1 - 1     KS = KQMAX1 - 1
			// for J in range(1, KQMAX1 + 1):
			//     W[J - 1] = 1.0 / float(J)
			for J := 1; J < KQMAX1+1; J++ {
				W[J-1] = 1.0 / float64(J)
			}

			//
			//     Compute the W(K) terms needed for the position interpolation
			//     (Note,  it is assumed throughout this routine that KS, which
			//     starts out as KQMAX1-1 (the ``maximum integration'')
			//     is at least 2.
			//

			JX := 0
			KS1 := KS - 1

			for KS >= 2 {
				JX = JX + 1

				for J := 1; J < JX+1; J++ {
					W[J+KS-1] = FC[J]*W[J+KS1-1] - WC[J-1]*W[J+KS-1]
				}
				KS = KS1
				KS1 = KS1 - 1
			}
			//
			//     Perform position interpolation: (Note that KS = 1 right now.
			//     We don't know much more than that.)
			//

			STATE := []float64{0, 0, 0}

			// DTtest = np.reshape(  mda_record[max_dim + 7 : max_dim * 4 + 7], (max_dim, 3), order="F"  )
			first_arr := mda_record[max_dim+7 : max_dim*2+7]
			second_arr := mda_record[max_dim*2+7 : max_dim*3+7]
			third_arr := mda_record[max_dim*3+7 : max_dim*4+7]
			DTtest := [][]float64{first_arr, second_arr, third_arr}

			for Ii := 0; Ii < 3; Ii++ {
				KQQ := KQ[Ii]
				SUM := 0.0

				for J := KQQ; J > 0; J-- {
					// v SUM = SUM + DTtest[J - 1][Ii] * W[J - 1 + KS]
					// SUM = SUM + DTtest[int(J-1)][Ii]*W[(int(J)-1+KS)]
					SUM = SUM + DTtest[Ii][int(J-1)]*W[(int(J)-1+KS)]
				}

				STATE[Ii] = REFPOS[Ii] + DELTA*(REFVEL[Ii]+DELTA*SUM)
			}
			return Position{STATE[0], STATE[1], STATE[2]}
		}

	}

	return Position{0, 0, 0}
}

func Ceres(dateInSeconds float64, chiron_np []float64) Position {

	type FileRecords struct {
		rec_start_addr int
		seg_start_time float64
		seg_last_time  float64
		int_len        float64
		rec_last_addr  int
	}

	var CHIRON_FILE_RECORDS = [12]FileRecords{
		{
			// SUN
			rec_start_addr: 0,
			seg_start_time: 0,
			seg_last_time:  0,
			int_len:        0,
			rec_last_addr:  0,
		},
		{
			// 1
			rec_start_addr: 8_065,
			seg_start_time: 14_308_254,
			seg_last_time:  631_108_800,
			int_len:        631_108_800,
			rec_last_addr:  54_071,
		},
		{
			// 2
			rec_start_addr: 54_072,
			seg_start_time: -651_515_663,
			seg_last_time:  14_308_254,
			int_len:        14_308_254,
			rec_last_addr:  100_078,
		},
		{
			// 3
			rec_start_addr: 100_079,
			seg_start_time: -1_303_985_603,
			seg_last_time:  -651_515_663,
			int_len:        -651_515_663,
			rec_last_addr:  146_085,
		},
		{
			// 4
			rec_start_addr: 146_086,
			seg_start_time: -1_965_211_060,
			seg_last_time:  -1_303_985_603,
			int_len:        -1_303_985_603,
			rec_last_addr:  192_092,
		},
		{
			// 5
			rec_start_addr: 192_093,
			seg_start_time: -2_619_710_503,
			seg_last_time:  -1_965_211_060,
			int_len:        -1_965_211_060,
			rec_last_addr:  238_099,
		},
		{
			// 6
			rec_start_addr: 238_100,
			seg_start_time: -3_155_716_800,
			seg_last_time:  -2_619_710_503,
			int_len:        -2_638_708_105,
			rec_last_addr:  276_285,
		},
		{
			// 7
			rec_start_addr: 276_286,
			seg_start_time: -3_155_716_800,
			seg_last_time:  -2_619_710_503,
			int_len:        -2_638_708_105,
			rec_last_addr:  276_285,
			// "rsize": 20,
			// init =  1002901396.5747848
			// n =  500.0
		},
		{
			// 8
			rec_start_addr: 322_293,
			seg_start_time: 1_274_239_976,
			seg_last_time:  1_919_182_997,
			int_len:        1_919_182_997,
			rec_last_addr:  368_299,
		},
		{
			// 9
			rec_start_addr: 368_300,
			seg_start_time: 1_919_182_997,
			seg_last_time:  2_575_120_362,
			int_len:        2_575_120_362,
			rec_last_addr:  414_306,
		},
		{
			// 10
			rec_start_addr: 414_307,
			seg_start_time: 2_575_120_362,
			seg_last_time:  3_187_252_800,
			int_len:        3_092_529_717,
			rec_last_addr:  457_920,
		},
	}

	const total_summaries_number = 10
	// max_dim is always 20. ceres_np[CERES_FILE_RECORDS[i_summ]["rec_last_addr"] - 2]
	const max_dim = 20
	// DFLSIZ = (4 * max_dim) + 11  // rsize = (4 * max_dim) + 11
	const DFLSIZ = 91
	const BUFSIZ = 100

	for i_summ := 1; i_summ <= total_summaries_number; i_summ++ {
		if CHIRON_FILE_RECORDS[i_summ].seg_start_time < dateInSeconds &&
			CHIRON_FILE_RECORDS[i_summ].seg_last_time > dateInSeconds {

			start_adress := CHIRON_FILE_RECORDS[i_summ].rec_start_addr
			last_adress := CHIRON_FILE_RECORDS[i_summ].rec_last_addr

			n_of_rec := int(chiron_np[last_adress-1])

			// Number of directory epochs
			var n_of_dir int = n_of_rec / BUFSIZ

			OFFD := last_adress - n_of_dir - 2
			OFFE := OFFD - n_of_rec

			RECNO := -1

			if n_of_rec <= BUFSIZ {

				data := chiron_np[(OFFE):(OFFE + n_of_rec)]

				if dateInSeconds < data[0] || dateInSeconds > data[len(data)-1] {
					fmt.Println("we have a problem")
				}

				for i, v := range data {
					if v == dateInSeconds {
						// fmt.Println("equality , index =", i, "value = ", arr[i])
						RECNO = i
						break
					}
					if v > dateInSeconds {
						// fmt.Println("index =", i-1, "value = ", arr[i-1])
						RECNO = i - 1
						break
					}
				}

			} else {
				for dir := 0; dir < n_of_dir; dir++ {

					data := chiron_np[(OFFD + dir) : (OFFD+dir)+1][0]

					if data > dateInSeconds {
						OFFD = OFFE + (dir)*BUFSIZ
						data := chiron_np[(OFFD):(OFFD + BUFSIZ)]

						// looking for the largest array element less than  dateInSeconds
						// and get its index
						if dateInSeconds < data[0] || dateInSeconds > data[len(data)-1] {
							fmt.Println("we have a problem")
						}

						for i, v := range data {
							if v == dateInSeconds {
								// fmt.Println("equality , index =", i, "value = ", arr[i])
								RECNO = i
								break
							}
							if v > dateInSeconds {
								// fmt.Println("index =", i-1, "value = ", arr[i-1])
								RECNO = i - 1
								break
							}
						}

						RECNO = dir*BUFSIZ + RECNO
						break
					}
				}
			}

			if RECNO == -1 {
				// print("chiron final records sec = ", dateInSeconds)
				Ind := n_of_rec % BUFSIZ
				data := chiron_np[(last_adress - n_of_dir - Ind) : last_adress-n_of_dir]

				// looking for the largest array element less than  dateInSeconds
				// and get its index
				if dateInSeconds < data[0] || dateInSeconds > data[len(data)-1] {
					fmt.Println("we have a problem")
				}

				for i, v := range data {
					if v == dateInSeconds {
						// fmt.Println("equality , index =", i, "value = ", arr[i])
						RECNO = i
						break
					}
					if v > dateInSeconds {
						// fmt.Println("index =", i-1, "value = ", arr[i-1])
						RECNO = i - 1
						break
					}
				}

				RECNO = (n_of_dir)*BUFSIZ + RECNO
			}

			OFFR := start_adress - 1 + (RECNO)*DFLSIZ
			mda_record := chiron_np[(OFFR):(OFFR + DFLSIZ)]

			// print("dateInSeconds = ", dateInSeconds)
			// print(mda_record)

			TL := mda_record[0]
			G := mda_record[1 : max_dim+1]
			REFPOS := []float64{mda_record[max_dim+1], mda_record[max_dim+3], mda_record[max_dim+5]}
			REFVEL := []float64{mda_record[max_dim+2], mda_record[max_dim+4], mda_record[max_dim+6]}

			KQMAX1 := int(mda_record[4*max_dim+7])

			KQ := []float64{(mda_record[4*max_dim+8]), (mda_record[4*max_dim+9]), (mda_record[4*max_dim+10])}

			KS := KQMAX1 - 1
			MQ2 := KQMAX1 - 2
			DELTA := dateInSeconds - TL
			TP := DELTA

			// FC = [0 for i in range(max_dim)]
			var FC [max_dim]float64
			FC[0] = 1.0
			// WC = [0 for i in range(max_dim - 1)]
			var WC [max_dim - 1]float64

			for J := 1; J < MQ2+1; J++ {
				if G[J-1] == 0.0 {
					fmt.Println("SPKE21\nA value of zero was found at index {0} of the step size vector.")
				}
				FC[J] = TP / G[J-1]
				WC[J-1] = DELTA / G[J-1]
				TP = DELTA + G[J-1]
			}
			// W = [0 for i in range(max_dim + 2)]
			var W [max_dim + 2]float64

			//
			//     Collect KQMAX1 reciprocals.
			//     KS = KQMAX1 - 1     KS = KQMAX1 - 1
			// for J in range(1, KQMAX1 + 1):
			//     W[J - 1] = 1.0 / float(J)
			for J := 1; J < KQMAX1+1; J++ {
				W[J-1] = 1.0 / float64(J)
			}

			//
			//     Compute the W(K) terms needed for the position interpolation
			//     (Note,  it is assumed throughout this routine that KS, which
			//     starts out as KQMAX1-1 (the ``maximum integration'')
			//     is at least 2.
			//

			JX := 0
			KS1 := KS - 1

			for KS >= 2 {
				JX = JX + 1

				for J := 1; J < JX+1; J++ {
					W[J+KS-1] = FC[J]*W[J+KS1-1] - WC[J-1]*W[J+KS-1]
				}
				KS = KS1
				KS1 = KS1 - 1
			}
			//
			//     Perform position interpolation: (Note that KS = 1 right now.
			//     We don't know much more than that.)
			//

			STATE := []float64{0, 0, 0}

			// DTtest = np.reshape(  mda_record[max_dim + 7 : max_dim * 4 + 7], (max_dim, 3), order="F"  )
			first_arr := mda_record[max_dim+7 : max_dim*2+7]
			second_arr := mda_record[max_dim*2+7 : max_dim*3+7]
			third_arr := mda_record[max_dim*3+7 : max_dim*4+7]
			DTtest := [][]float64{first_arr, second_arr, third_arr}

			for Ii := 0; Ii < 3; Ii++ {
				KQQ := KQ[Ii]
				SUM := 0.0

				for J := KQQ; J > 0; J-- {
					// v SUM = SUM + DTtest[J - 1][Ii] * W[J - 1 + KS]
					// SUM = SUM + DTtest[int(J-1)][Ii]*W[(int(J)-1+KS)]
					SUM = SUM + DTtest[Ii][int(J-1)]*W[(int(J)-1+KS)]
				}

				STATE[Ii] = REFPOS[Ii] + DELTA*(REFVEL[Ii]+DELTA*SUM)
			}
			return Position{STATE[0], STATE[1], STATE[2]}
		}

	}

	return Position{0, 0, 0}
}
