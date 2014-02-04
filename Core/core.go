package Core

import "../Tool"
import "../Data"
import "../Keyboard"

var Ctrl int

func init() {
	Ctrl = 0
}

func Run() {
	Tool.Update()
	for {
		if Ctrl == 1 {
			return
		}
		KCode := Tool.Wait()
		switch KCode {
		case Keyboard.KUP:
			Tool.Move(Data.X, Data.Y-1)
		case Keyboard.KDOWN:
			Tool.Move(Data.X, Data.Y+1)
		case Keyboard.KLEFT:
			Tool.Move(Data.X-1, Data.Y)
		case Keyboard.KRIGHT:
			Tool.Move(Data.X+1, Data.Y)
		case Keyboard.KESC:
			return
		}
		for i := 0; i < len(Data.CurrentMap.Events); i++ {
			if Data.CurrentMap.Events[i] == nil {
				continue
			}
			if Data.CurrentMap.Events[i].X == Data.X && Data.CurrentMap.Events[i].Y == Data.Y {
				if Data.CurrentMap.Events[i].Valid {
					Data.CurrentMap.Events[i].Run(Data.CurrentMap.Events[i])
				}
				break
			}
		}
	}
}

func Init() {
	Data.DispMap = Tool.GetMap(Data.X, Data.Y)
	Data.DispDialog = ""
	Data.DispStatus = ""
}
