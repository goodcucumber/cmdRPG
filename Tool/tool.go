package Tool

import "../Keyboard"
import "time"
import "os"
import "os/exec"
import "fmt"
import "../Data"
import "../Map"

/*

*/

func Wait() int {
	time.Sleep(80 * time.Millisecond)
	Keyboard.ClearBuffer()
	for {
		if !Keyboard.Hit() {
			time.Sleep(80 * time.Millisecond)
			continue
		}
		rtn := Keyboard.GetKeyCode()
		Keyboard.ClearBuffer()
		return rtn
	}
}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Draw(s string) {
	Clear()
	fmt.Println(s)
}

func Update() {
	Clear()
	fmt.Println(Data.DispMap)
	if Data.DispStatus != "" {
		fmt.Println(Data.DispStatus)
	}
	if Data.DispDialog != "" {
		fmt.Println("==============================")
		fmt.Println(Data.DispDialog)
		fmt.Print("==============================")
	}
}

//
//
//map

func GetMap(X, Y int) string {
	var m *Map.Map = Data.CurrentMap
	if 2*X >= len(m.Map[0]) {
		return ""
	}
	if X < 0 {
		return ""
	}
	if Y >= len(m.Map) {
		return ""
	}
	if Y < 0 {
		return ""
	}
	if m.Map[Y][2*X] == Data.Wall[0] {
		return ""
	}
	if m.Map[Y][2*X+1] == Data.Wall[1] {
		return ""
	}
	all := m.Title
	all += "\n"
	for i := 0; i < len(m.Map)-1; i++ {
		all += m.Map[i]
		all += "\n"
	}
	all += m.Map[len(m.Map)-1]
	b := []byte(all)
	b[len(m.Title)+1+(len(m.Map[0])+1)*Y+X*2] = Data.Player[0]
	b[len(m.Title)+1+(len(m.Map[0])+1)*Y+X*2+1] = Data.Player[1]
	return string(b)
}
func Move(x, y int) {
	s := GetMap(x, y)
	if s != "" {
		Data.X = x
		Data.Y = y
		Data.DispMap = s
		Update()
	}
}
