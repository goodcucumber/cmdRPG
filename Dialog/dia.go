package Dialog

import "../Tool"
import "../Keyboard"
import "../Data"

type Dialog struct {
	Log      []string
	Pos      int
	Skipable bool
}

func (d *Dialog) Go() int {
	d.Pos = 0
	Data.DispDialog = d.Log[d.Pos]
	Tool.Update()
	for {
		keyCode := Tool.Wait()
		if keyCode == Keyboard.KUP {
			d.Pos--
		}
		if keyCode == Keyboard.KDOWN {
			d.Pos++
		}
		if keyCode == Keyboard.KENTER {
			d.Pos++
		}
		if d.Skipable {
			if keyCode == Keyboard.KESC {
				d.Pos = len(d.Log)
			}
		}
		if d.Pos < 0 {
			d.Pos = 0
		}
		if d.Pos >= len(d.Log) {
			Data.DispDialog = ""
			Tool.Update()
			return 0
		}
		Data.DispDialog = d.Log[d.Pos]
		Tool.Update()
	}
}
