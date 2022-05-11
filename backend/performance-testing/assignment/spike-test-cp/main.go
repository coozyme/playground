package main

import (
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

//kita bisa menggunakan attack yang meningkat seperti ini untuk stress/capacity test

//melakukan attack selama 4 detik
//setiap detik frequency akan naik sesuai increaseValue
//nilai yang digunakan
/*
duration := 4 * time.Second
frequency := 5
spikeValue := 20
*/

func spikeTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	duration := 4 * time.Second
	spike := 20
	rate := vegeta.Rate{Freq: 5, Per: time.Second}

	targetter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
		Body:   []byte{},
	})

	metrics = vegetaAttackSpike(targetter, rate, duration, spike)
	return metrics
}

//setiap detik ke 2 request akan melonjak
//contoh
//detik 1 request = 5
//detik 2 request = 20
//detik 3 request = 5
//detik 4 request = 20

func vegetaAttackSpike(targeter vegeta.Targeter, rate vegeta.ConstantPacer, duration time.Duration, spikeValue int) *vegeta.Metrics {
	var metrics vegeta.Metrics
	// TODO: answer here
	attacker := vegeta.Attacker{}
	normalFreq := rate.Freq

	for i := 0; i < int(duration.Seconds()); i++ {
		if i%2 == 0 {
			rate.Freq = spikeValue
		} else {
			rate.Freq = normalFreq
		}

		for res := range attacker.Attack(targeter, rate, time.Second, "Spike test example") {
			metrics.Add(res)
		}
	}

	metrics.Close()
	return &metrics
}
