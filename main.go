package main

import "./Core"
import "./Map"
import "./Dialog"
import "./Event"
import "./Data"
import "./Talk"
import "./Tool"
import "./Keyboard"
import "./Color"
import "./BGM"
import "fmt"
import "time"

func main() {
	//map 1
	m1 := new(Map.Map)
	m1.Map = make([]string, 10)
	/////////////0,1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6
	m1.Map[0] = "**********************************" //0
	m1.Map[1] = "**                            ''**" //1
	m1.Map[2] = "**  **  **************************" //2
	m1.Map[3] = "**  **                          **" //3
	m1.Map[4] = "**      ********  **   Good     **" //4
	m1.Map[5] = "********************     day~   **" //5
	m1.Map[6] = "**        **      **            **" //6
	m1.Map[7] = "****  **  **  **  **            **" //7
	m1.Map[8] = "**<>  **      **            --->  " //8
	m1.Map[9] = "**********************************" //9
	m1.Title = "Have Fun~"
	//
	m1.Events = make([]*Event.Event, 3)
	m1.Events[0] = new(Event.Event)
	m1.Events[1] = new(Event.Event)
	m1.Events[2] = new(Event.Event)
	m1.Events[0].X = 15
	m1.Events[0].Y = 1
	m1.Events[0].Valid = true
	m1.Events[0].Run = func(e *Event.Event) int {
		d := new(Dialog.Dialog)
		d.Skipable = false
		d.Log = make([]string, 2)
		d.Log[0] = "你好, 我是机智的萤萤~"
		d.Log[1] = Talk.F()
		d.Go()
		return 0
	}
	m1.Events[1].X = 1
	m1.Events[1].Y = 8
	m1.Events[1].Valid = true
	m1.Events[1].Var = make([]int, 1)
	m1.Events[1].Var[0] = 0
	m1.Events[1].Run = func(e *Event.Event) int {
		d := new(Dialog.Dialog)
		d.Skipable = false
		d.Log = make([]string, 2)
		d.Log[0] = "Hi~"
		if e.Var[0] == 0 {
			d.Log[1] = "You have *never* been here!"
		} else if e.Var[0] == 1 {
			d.Log[1] = "You have been here *once*!"
		} else if e.Var[0] == 2 {
			d.Log[1] = "You have been here *twice*!"
		} else {
			d.Log[1] = "You have been here *" + string(e.Var[0]) + " times*!"
		}
		d.Go()
		e.Var[0]++
		if e.Var[0]%2 == 0 {
			Tool.Move(3, 3)
		} else {
			Tool.Move(15, 4)
		}
		return 0
	}
	m1.Events[2].X = 16
	m1.Events[2].Y = 8
	m1.Events[2].Valid = true
	m1.Events[2].Run = func(e *Event.Event) int {
		Data.CurrentMap = Data.Maps[1]
		Tool.Move(0, 8)
		BGM.Stop()
		BGM.BGM("b1.mp3")
		BGM.Vol(0.5)
		return 0
	}
	/*


	*/
	//map2
	m2 := new(Map.Map)
	m2.Map = make([]string, 10)
	m2.Title = "Oh! No! I'm stuck!?"
	/////////////0///////////////////1///////////////////2//////////
	/////////////0,1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6,7,8,9,0,1,2,3,4,5
	m2.Map[0] = "****************************************************" //0
	m2.Map[1] = "**@    @    @@    @@@@   @@@@   @    @            **" //1
	m2.Map[2] = "**@    @    @@    @   @@ @   @@ @    @ *********  **" //2
	m2.Map[3] = "**@    @   @  @   @    @ @    @ @    @ ##*******  **" //3
	m2.Map[4] = "**@@@@@@   @  @   @    @ @    @  @  @ #**#**      **" //4
	m2.Map[5] = "**@    @  @    @  @@@@@@ @@@@@@   @@ *[]****  **  **" //5
	m2.Map[6] = "**@    @  @@@@@@  @      @        @@ *    **  **  **" //6
	m2.Map[7] = "**@    @ @      @ @      @        @@ *  ****  **##**" //7
	m2.Map[8] = "  @    @ @      @ @      @        @@ *    **  **()**" //8
	m2.Map[9] = "****************************************  **  ******" //9
	//
	m2.Events = make([]*Event.Event, 5)
	//e0: to map 1
	m2.Events[0] = new(Event.Event)
	m2.Events[0].X = 0
	m2.Events[0].Y = 8
	m2.Events[0].Valid = true
	m2.Events[0].Run = func(e *Event.Event) int {
		Data.CurrentMap = Data.Maps[0]
		Tool.Move(16, 8)
		BGM.Stop()
		BGM.BGM("b0.aac")
		BGM.Vol(2)
		return 0
	}
	//e1: dialog
	m2.Events[1] = new(Event.Event)
	m2.Events[1].X, m2.Events[1].Y = 24, 8
	m2.Events[1].Valid = true
	m2.Events[1].Var = make([]int, 1)
	m2.Events[1].Run = func(e *Event.Event) int {
		e.Var[0] = 0
		d := new(Dialog.Dialog)
		d.Log = make([]string, 2)
		d.Skipable = false
		d.Log[0] = "呐, 萤萤觉得这个游戏不好玩呢~"
		d.Log[1] = "你说呢?"
		d.Go()
		se1 := "> :好玩       :不好玩"
		se2 := "  :好玩     > :不好玩"
		Data.DispDialog = se1
		Tool.Update()
		e.Var[0] = 1
		for {
			k := Tool.Wait()
			if k == Keyboard.KENTER {
				break
			}
			if e.Var[0] == 1 {
				if k == Keyboard.KRIGHT {
					e.Var[0] = 2
					Data.DispDialog = se2
					Tool.Update()
					continue
				}
			}
			if e.Var[0] == 2 {
				if k == Keyboard.KLEFT {
					e.Var[0] = 1
					Data.DispDialog = se1
					Tool.Update()
					continue
				}
			}
		}
		if e.Var[0] == 2 {
			d2 := new(Dialog.Dialog)
			d2.Skipable = false
			d2.Log = make([]string, 2)
			d2.Log[0] = "对吧, 对吧~"
			d2.Log[1] = "所以, 结束了哦。"
			d2.Go()
			Core.Ctrl = 1
			return 0
		} else {
			d2 := new(Dialog.Dialog)
			d2.Skipable = false
			d2.Log = make([]string, 2)
			d2.Log[0] = "诶~, 是吗...唔..."
			d2.Log[1] = "嗯, 嗯, 萤萤也再看看好啦\\(o·ω·o)/吼吼!"
			d2.Go()
			return 0
		}
	}
	//e3: to e4
	m2.Events[2] = new(Event.Event)
	m2.Events[2].X, m2.Events[2].Y = 20, 9
	m2.Events[2].Valid = true
	m2.Events[2].Run = func(e *Event.Event) int {
		Tool.Move(22, 9)
		return 0
	}
	//e4: to e3
	m2.Events[3] = new(Event.Event)
	m2.Events[3].X, m2.Events[3].Y = 22, 9
	m2.Events[3].Valid = true
	m2.Events[3].Run = func(e *Event.Event) int {
		Tool.Move(20, 9)
		return 0
	}
	m2.Events[4] = new(Event.Event)
	m2.Events[4].X, m2.Events[4].Y = 19, 5
	m2.Events[4].Valid = true
	m2.Events[4].Run = func(e *Event.Event) int {
		BGM.Stop()
		BGM.BGM("be.mp3")
		d := new(Dialog.Dialog)
		d.Log = make([]string, 5)
		d.Log[0] = "终于通关了呢..."
		d.Log[1] = "是不是很好玩呢? 也许没什么意思吧..."
		d.Log[2] = "果然还是差得远呢"
		d.Log[3] = " 不过, 一定能作出更加、更加有趣的游戏吧"
		d.Log[4] = "所以 ..."
		d.Skipable = false
		d.Go()
		Tool.Clear()
		Color.SetColor(11)
		fmt.Println("        游戏就这样结束了呢")
		fmt.Println()
		fmt.Println()
		time.Sleep(4 * time.Second)
		Color.SetColor(9)
		fmt.Println("         感觉怎么样呢?")
		time.Sleep(6 * time.Second)
		Tool.Clear()
		time.Sleep(4 * time.Second)
		Color.SetColor(3)
		fmt.Println(" 虽然我知道, 只是这种程度的游戏的话")
		time.Sleep(4 * time.Second)
		Color.SetColor(1)
		fmt.Println("\n\n")
		fmt.Println("       是不太会有人喜欢的吧")
		fmt.Println("\n\n")
		time.Sleep(4 * time.Second)
		Color.SetColor(14)
		fmt.Println("我还是想问一句: 这个, 给你带来快乐了吧?")
		Color.SetColor(2)
		time.Sleep(5 * time.Second)
		Tool.Clear()
		fmt.Println("            结束了")
		//fmt.Println("     The End")
		//fmt.Println("      終わり")
		time.Sleep(4 * time.Second)
		fmt.Println("　　        谢谢你")
		time.Sleep(4 * time.Second)
		Keyboard.ClearBuffer()
		Color.SetColor(15)
		fmt.Println("          迷宫<>第三版")
		fmt.Println("\n\n")
		time.Sleep(4 * time.Second)
		fmt.Println("          制作::高见")
		fmt.Println("\n\n")
		time.Sleep(4 * time.Second)
		fmt.Println("           BGM::|")
		fmt.Println("                |1::东方project -  砕月")
		fmt.Println("                |2::  《刀语》  -  月刀歌")
		fmt.Println("                |3::东方project -  众神眷顾的幻想乡")
		time.Sleep(4 * time.Second)
		time.Sleep(3 * time.Second)
		Tool.Clear()
		fmt.Println("        ~A Golang project~")
		fmt.Println("\n\n")
		time.Sleep(4 * time.Second)
		fmt.Println("   included BASS -- an audio library")
		fmt.Println("\n\n")
		time.Sleep(4 * time.Second)
		fmt.Println("               Thanks")
		Keyboard.ClearBuffer()
		time.Sleep(time.Second)
		Keyboard.GetKeyCode()
		Tool.Clear()
		Core.Ctrl = 1
		Color.SetColor(7)
		return 0
	}
	/*

	*/
	Data.X = 1
	Data.Y = 1
	Data.Maps = make([]*Map.Map, 2)
	Data.Maps[0] = m1
	Data.Maps[1] = m2
	Data.CurrentMap = Data.Maps[0]
	Core.Init()
	BGM.BGM("b0.aac")
	BGM.Vol(2)
	Core.Run()
}
